package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"time"

	"github.com/bytedance/sonic"
	"github.com/playwright-community/playwright-go"

	// "main.go/dataModel/CookieModel"
	// "main.go/dataModel/ShopAccount"
	// "main.go/dataModel/ShopModel"
	"main.go/dataModel/SkuModel"
	"main.go/internal/types"
)

// 初始化Playwright服务
func InitPlaywrightService() (*types.PlaywrightService, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("启动Playwright失败: %w", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return nil, fmt.Errorf("启动浏览器失败: %w", err)
	}

	context, err := browser.NewContext()
	if err != nil {
		return nil, fmt.Errorf("创建上下文失败: %w", err)
	}

	page, err := context.NewPage()
	if err != nil {
		return nil, fmt.Errorf("创建页面失败: %w", err)
	}

	return &types.PlaywrightService{
		PW:      pw,
		Browser: browser,
		Context: context,
		Page:    page,
	}, nil
}

// 使用Playwright访问URL并拦截响应
// 修复错误2和3: 使用接口类型作为返回类型
func VisitUrlWithPlaywright(appState *types.AppState, url string) (playwright.Response, error) {
	shopInfo := appState.ShopCookieInfo

	errCls := appState.Playwright.Context.ClearCookies()
	if errCls != nil {
		return nil, fmt.Errorf("清除COOKIE失败: %v", errCls)
	}

	// 解析Cookie字符串并添加到上下文
	opck := []playwright.OptionalCookie{}
	err := sonic.UnmarshalString(shopInfo.Cookie, &opck)
	if err != nil {
		log.Fatalf("could not unmarshal cookie: %v", err)
		return nil, fmt.Errorf("添加Cookie失败: %v", err)
	}

	if err := appState.Playwright.Context.AddCookies(opck); err != nil {
		return nil, fmt.Errorf("添加Cookie失败: %v", err)
	}

	// 设置响应拦截器
	appState.Playwright.Page.OnResponse(func(response playwright.Response) {
		log.Printf("响应: %s - %d", response.URL(), response.Status())
	})

	// 导航到URL
	response, err := appState.Playwright.Page.Goto(url)
	if err != nil {
		return nil, fmt.Errorf("导航失败: %w", err)
	}

	// 等待页面加载完成
	// if err := appState.Playwright.Page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
	// 	State: playwright.LoadStateNetworkidle,
	// }); err != nil {
	// 	return nil, fmt.Errorf("等待页面加载失败: %w", err)
	// }

	return response, nil
}

func AddHttpCookie(appState *types.AppState, url string) (bool, []SkuModel.DataItem, error) {
	// 获取Playwright Cookies
	shopInfo := appState.ShopCookieInfo

	errCls := appState.Playwright.Context.ClearCookies()
	if errCls != nil {
		return false, []SkuModel.DataItem{}, fmt.Errorf("清除COOKIE失败: %v", errCls)
	}

	// 解析Cookie字符串并添加到上下文
	opck := []playwright.OptionalCookie{}
	err := sonic.UnmarshalString(shopInfo.Cookie, &opck)
	if err != nil {
		log.Fatalf("could not unmarshal cookie: %v", err)
		return false, []SkuModel.DataItem{}, fmt.Errorf("添加Cookie失败: %v", err)
	}

	if err := appState.Playwright.Context.AddCookies(opck); err != nil {
		return false, []SkuModel.DataItem{}, fmt.Errorf("添加Cookie失败: %v", err)
	}

	// 获取Playwright Cookies
	cookies, err := appState.Playwright.Context.Cookies()
	if err != nil {
		log.Fatal(err)
		return false, []SkuModel.DataItem{}, fmt.Errorf("获取浏览器Cookie失败: %v", err)
	}

	// 转换为http.Cookie
	var httpCookies []*http.Cookie
	for _, c := range cookies {
		httpCookies = append(httpCookies, &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  time.Unix(int64(c.Expires), 0),
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		})
	}

	// 使用HTTP请求
	req, _ := http.NewRequest("GET", url, nil)
	for _, cookie := range httpCookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false, []SkuModel.DataItem{}, fmt.Errorf("获取HTTP响应错误: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, []SkuModel.DataItem{}, fmt.Errorf("读取响应数据错误: %v", err)
		}
		var data SkuModel.Response
		err = json.Unmarshal(body, &data)
		if err != nil {
			return false, []SkuModel.DataItem{}, fmt.Errorf("读取响应数据错误: %v", err)
		}
		return true, data.Data, fmt.Errorf("读取响应数据错误: %v", err)
	}
	return false, []SkuModel.DataItem{}, fmt.Errorf("执行获取产品列表错误: %v", err)
}
