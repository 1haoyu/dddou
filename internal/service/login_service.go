package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/playwright-community/playwright-go"
	"main.go/dataModel/CookieModel"
	"main.go/dataModel/ShopAccount"
	"main.go/dataModel/ShopModel"
	"main.go/dataModel/UserModel"
	"main.go/internal/types"
)

const (
	loginURL         = "https://fxg.jinritemai.com/login/common"
	successURL       = "https://fxg.jinritemai.com/ffa/mshop/homepage/index"
	targetRoute      = "get_login_subject?"
	shopAccountRoute = "get_account_info?"
	tokenRoute       = "get_interact_renew?"
	shopIdRoute      = "getdid?"
)

// LoginOrAutoLogin 主登录函数
func LoginOrAutoLogin(username, password string, appState *types.AppState) (bool, error) {
	// 1. 检查数据库中是否存在用户信息
	user, err := UserModel.Api_find_by_username(username)
	if err == nil && user.LoginName == username && user.Cookie != "" {
		// 2. 如果存在有效Cookie，尝试自动登录
		unixTime := time.Now().Unix()
		if user.Expires > float64(unixTime) {
			if ok, err := autoLoginWithCookie(username, user.Cookie, appState); err == nil {
				return ok, nil
			}
		}

		log.Println("Cookie已失效或需要人工交互，尝试用户名密码登录")
	}

	// 3. 使用Playwright进行自动化登录
	ok, err := LoginByAccount(username, password, appState)
	if err != nil {
		return ok, fmt.Errorf("登录失败: %v", err)
	}

	return true, nil
}

// autoLoginWithCookie 使用Cookie自动登录并等待用户选择店铺
func autoLoginWithCookie(username, cookieStr string, appState *types.AppState) (bool, error) {
	// 初始化Playwright
	pw, err := playwright.Run()
	if err != nil {
		return false, fmt.Errorf("启动Playwright失败: %v", err)
	}
	defer pw.Stop()

	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // 必须非无头模式，需要用户交互
	})
	if err != nil {
		return false, fmt.Errorf("启动浏览器失败: %v", err)
	}
	defer browser.Close()

	// 创建上下文并添加Cookie
	context, err := browser.NewContext()
	if err != nil {
		return false, fmt.Errorf("创建上下文失败: %v", err)
	}
	defer context.Close()

	// 解析Cookie字符串并添加到上下文
	opck := []playwright.OptionalCookie{}
	err = sonic.UnmarshalString(cookieStr, &opck)
	if err != nil {
		log.Fatalf("could not unmarshal cookie: %v", err)
		return false, fmt.Errorf("添加Cookie失败: %v", err)
	}

	if err := context.AddCookies(opck); err != nil {
		return false, fmt.Errorf("添加Cookie失败: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		return false, fmt.Errorf("创建页面失败: %v", err)
	}

	// 创建通道用于接收数据
	shopChan := make(chan []ShopModel.Account, 1)
	tokenChan := make(chan map[string]string, 1) // 用于传递token参数
	loginSuccessChan := make(chan bool, 1)
	loginErrorChan := make(chan error, 1)

	// 设置路由拦截 - 捕获店铺信息和token
	err = page.Route("**/*", func(route playwright.Route) {
		url := route.Request().URL()
		log.Printf("拦截到请求: %s", url)

		// 1. 拦截店铺信息请求
		if strings.Contains(url, targetRoute) {
			log.Printf("检测到店铺信息请求: %s", url)
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
				return
			}

			// 等待响应
			response, err := page.ExpectResponse(url, func() error { return nil })
			if err != nil {
				log.Printf("等待响应失败: %v", err)
				return
			}

			body, err := response.Body()
			if err != nil {
				log.Printf("获取响应体失败: %v", err)
				return
			}

			var resp ShopModel.Response
			if err := json.Unmarshal(body, &resp); err != nil {
				log.Printf("解析JSON失败: %v", err)
				return
			}
			log.Printf("获取到 %d 个店铺信息", len(resp.Data.LoginSubjectList))
			shopChan <- resp.Data.LoginSubjectList
		} else if strings.Contains(url, tokenRoute) {
			log.Printf("检测到token请求: %s", url)
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
				return
			}

			// 解析URL参数
			params := ExtractTokenParams(url)
			if params["__token"] != "" && params["verifyFp"] != "" {
				log.Printf("获取到token参数: %+v", params)
				tokenChan <- params
			} else {
				log.Printf("未找到完整的token参数: %s", url)
			}
		} else if strings.Contains(url, shopAccountRoute) {
			log.Printf("检测到token请求: %s", url)
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
				return
			}

			// 等待响应
			response, err := page.ExpectResponse(url, func() error { return nil })
			if err != nil {
				log.Printf("等待响应失败: %v", err)
				return
			}

			body, err := response.Body()
			if err != nil {
				log.Printf("获取响应体失败: %v", err)
				return
			}

			var resp ShopAccount.Root
			if err := json.Unmarshal(body, &resp); err != nil {
				log.Printf("解析JSON失败: %v", err)
				return
			}
			appState.ShopAccountChan = resp.Data
		} else if strings.Contains(url, shopIdRoute) {
			log.Printf("检测到token请求: %s", url)
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
				return
			}
			// 解析URL参数
			params := ExtractTokenParams(url)
			if params["shopid"] != "" && params["account"] != "" && params["shopname"] != "" {
				log.Printf("获取到token参数: %+v", params)
				appState.ShopAccountChan = ShopAccount.Data{
					Account_id: params["shopid"],
				}
			} else {
				log.Printf("未找到完整的token参数: %s", url)
			}
		} else {
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
			}
		}
	})

	if err != nil {
		return false, fmt.Errorf("设置路由拦截失败: %v", err)
	}

	// 监听页面导航事件
	page.On("response", func(response playwright.Response) {
		if response.URL() == successURL && response.Status() == 200 {
			log.Println("检测到成功登录页面")
			loginSuccessChan <- true
		}
	})

	// 导航到登录页面
	log.Println("导航到登录页面...")
	if _, err := page.Goto(loginURL); err != nil {
		return false, fmt.Errorf("导航到登录页面失败: %v", err)
	}

	// 等待页面加载完成
	page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateDomcontentloaded,
	})

	// 提示用户手动操作
	log.Println("请在弹出的浏览器中选择店铺...")

	// 等待用户操作结果
	var shops []ShopModel.Account
	var tokenParams map[string]string

	select {
	case shops = <-shopChan:
		log.Printf("成功获取 %d 个店铺", len(shops))
		// 获取Cookie
		cookies, err := context.Cookies()
		if err != nil {
			return false, fmt.Errorf("获取Cookie失败: %v", err)
		}

		var exp float64 = 0
		for _, c := range cookies {
			if c.Expires == 0 {
				continue
			}
			exp = c.Expires
		}

		ck_save, errs := sonic.MarshalString(cookies)
		if errs != nil {
			log.Fatalf("unable to save cookie: %v", errs)
		}
		if uperr := UserModel.UpdateUserCookie(username, ck_save, exp); uperr != nil {
			log.Printf("更新COOKIE错误: %v\n", err)
		}

		saveOrShop(shops, username)
	case <-loginSuccessChan:
		log.Println("登录成功，但未获取到店铺数据")
	case err := <-loginErrorChan:
		return false, fmt.Errorf("登录失败: %v", err)
	case <-time.After(5 * time.Minute): // 设置较长的超时时间
		log.Println("等待用户操作超时")
		return false, fmt.Errorf("用户操作超时")
	}

	// 等待token信息
	select {
	case tokenParams = <-tokenChan:
		log.Println("成功获取token参数")
	case <-time.After(5 * time.Minute):
		log.Println("等待token信息超时")
		return false, fmt.Errorf("等待token信息超时")
	}

	// 将Cookie转换为字符串
	// 保存token信息到数据库
	if tokenParams != nil {
		// 获取更新后的Cookie
		newCookies, err := context.Cookies()
		if err != nil {
			return false, fmt.Errorf("获取Cookie失败: %v", err)
		}
		ck_save, errs := sonic.MarshalString(newCookies)
		if errs != nil {
			return false, fmt.Errorf("获取Cookie失败: %v", err)
		}
		var exp float64 = 0
		for _, c := range newCookies {
			exp = c.Expires
		}

		if err := saveTokenInfo(appState, username, ck_save, exp, tokenParams); err != nil {
			log.Printf("保存token信息失败: %v", err)
			return false, fmt.Errorf("保存token信息失败: %v", err)
		}

		return true, nil
	}

	return false, fmt.Errorf("未获取完整信息: %v", err)
}

// extractTokenParams 从URL中提取token参数
func ExtractTokenParams(url string) map[string]string {
	params := make(map[string]string)

	// 解析查询参数
	if strings.Contains(url, "?") {
		query := strings.SplitN(url, "?", 2)[1]
		keyValues := strings.Split(query, "&")

		for _, kv := range keyValues {
			parts := strings.SplitN(kv, "=", 2)
			if len(parts) == 2 {
				params[parts[0]] = parts[1]
			}
		}
	}

	return params
}

// saveTokenInfo 保存token信息到数据库
func saveTokenInfo(appState *types.AppState, username string, cookie string, exp float64, params map[string]string) error {
	token := params["__token"]
	verifyFp := params["verifyFp"]

	if token == "" || verifyFp == "" {
		return fmt.Errorf("token或verifyFp为空")
	}

	if appState.ShopAccountChan.Account_id == "" {
		return fmt.Errorf("还未获取账号信息")
	}
	shopInfo := ShopModel.Api_find_struct_by_id(appState.ShopAccountChan.Account_id)
	if shopInfo.AccountID == appState.ShopAccountChan.Account_id {

	}
	findCookie := CookieModel.Api_find_struct_by_id(appState.ShopAccountChan.Account_id)
	if findCookie.AccountID == appState.ShopAccountChan.Account_id {
		// 用户存在，更新信息
		findCookie.Cookie = cookie
		findCookie.Expires = exp
		updateData := map[string]interface{}{
			"cookie":  findCookie.Cookie,
			"expires": findCookie.Expires,
		}

		// 调用 Api_update 方法
		if err := CookieModel.Api_update(findCookie.SubjectID, updateData); err != nil {
			return fmt.Errorf("更新用户信息失败: %v", err)
		}
		return nil
	} else {
		// 创建或更新token信息
		cookieInfo := CookieModel.CookieInfo{
			SubjectID:    shopInfo.SubjectID,
			LoginName:    shopInfo.LoginName,
			AccountID:    shopInfo.AccountID,
			LoginType:    shopInfo.IdentityType,
			Cookie:       cookie,
			Token:        token,
			VerifyFp:     verifyFp,
			MemberID:     shopInfo.MemberID,
			EncodeShopID: shopInfo.EncodeShopID,
			IdentityType: shopInfo.IdentityType,
			Expires:      exp,
		}

		if err := CookieModel.Api_insert(cookieInfo); err == false {
			return fmt.Errorf("保存Cookie信息失败: %v", err)
		}

		log.Printf("成功保存token信息: %s", token)
		return nil
	}

}

// LoginByAccount 使用Playwright进行自动化登录
func LoginByAccount(username, password string, appState *types.AppState) (bool, error) {
	// 初始化Playwright
	pw, err := playwright.Run()
	if err != nil {
		return false, fmt.Errorf("启动Playwright失败: %v", err)
	}
	defer pw.Stop()

	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // 调试时设为false
	})
	if err != nil {
		return false, fmt.Errorf("启动浏览器失败: %v", err)
	}
	defer browser.Close()

	// 创建上下文和页面
	context, err := browser.NewContext()
	if err != nil {
		return false, fmt.Errorf("创建上下文失败: %v", err)
	}
	defer context.Close()

	page, err := context.NewPage()
	if err != nil {
		return false, fmt.Errorf("创建页面失败: %v", err)
	}

	// 创建通道用于接收店铺数据
	shopChan := make(chan []ShopModel.Account, 1)
	loginSuccessChan := make(chan bool, 1)
	loginErrorChan := make(chan error, 1)

	// 修正：正确使用路由拦截
	err = page.Route("**/*", func(route playwright.Route) {
		url := route.Request().URL()

		if strings.Contains(url, targetRoute) {
			log.Printf("获取的目标URL: %s", url)

			// 修正：Continue() 只返回一个错误值
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
				return
			}

			// 等待响应
			response, err := page.ExpectResponse(url, func() error { return nil })
			if err != nil {
				log.Printf("等待响应失败: %v", err)
				return
			}

			body, err := response.Body()
			if err != nil {
				log.Printf("获取响应体失败: %v", err)
				return
			}
			log.Printf("获取的目标URL数据: %s", body)

			var resp ShopModel.Response
			if err := json.Unmarshal(body, &resp); err != nil {
				log.Printf("解析JSON失败: %v", err)
				return
			}
			log.Printf("请求到店铺数量: %d", len(resp.Data.LoginSubjectList))
			shopChan <- resp.Data.LoginSubjectList
		} else {
			// 其他请求继续执行
			if err := route.Continue(); err != nil {
				log.Printf("继续请求失败: %v", err)
			}
		}
	})

	if err != nil {
		return false, fmt.Errorf("设置路由拦截失败: %v", err)
	}

	// 2. 监听页面导航事件 - 替代WaitForURL
	page.On("response", func(response playwright.Response) {
		if response.URL() == successURL && response.Status() == 200 {
			loginSuccessChan <- true
		}
	})

	// 导航到登录页面
	if _, err := page.Goto(loginURL); err != nil {
		return false, fmt.Errorf("导航到登录页面失败: %v", err)
	}

	// 等待页面加载完成
	page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateDomcontentloaded,
	})

	// 点击"邮箱登录"按钮
	emailLoginBtn := page.GetByText("邮箱登录")
	if err := emailLoginBtn.Click(); err != nil {
		return false, fmt.Errorf("点击邮箱登录按钮失败: %v", err)
	}

	// 输入邮箱
	emailInput := page.GetByPlaceholder("请输入邮箱")
	if err := emailInput.Fill(username); err != nil {
		return false, fmt.Errorf("输入邮箱失败: %v", err)
	}

	// 输入密码
	passwordInput := page.GetByPlaceholder("密码")
	if err := passwordInput.Fill(password); err != nil {
		return false, fmt.Errorf("输入密码失败: %v", err)
	}
	err = page.GetByRole(*playwright.AriaRoleCheckbox).Check()
	if err != nil {
		log.Fatalf("could not Check: %v", err)
		return false, fmt.Errorf("点击同意隐私条款: %v", err)
	}
	time.Sleep(2 * time.Second)
	// 点击登录按钮
	loginBtn := page.GetByRole(*playwright.AriaRoleButton).GetByText("登录")
	if err := loginBtn.Click(); err != nil {
		return false, fmt.Errorf("点击登录按钮失败: %v", err)
	}

	// 5. 等待登录结果 - 使用多路选择器处理不同情况
	var shops []ShopModel.Account
	select {
	case shops = <-shopChan:
		log.Printf("成功获取 %d 个店铺", len(shops))
		// 获取Cookie
		cookies, err := context.Cookies()
		if err != nil {
			return false, fmt.Errorf("获取Cookie失败: %v", err)
		}

		var exp float64 = 0
		for _, c := range cookies {
			if c.Expires == 0 {
				continue
			}
			exp = c.Expires
		}

		ck_save, errs := sonic.MarshalString(cookies)
		if errs != nil {
			log.Fatalf("unable to save cookie: %v", errs)
		}

		saveOrUpdateUser(username, password, ck_save, exp)
		saveOrShop(shops, username)
	case <-loginSuccessChan:
		log.Println("登录成功，但未获取到店铺数据")
	case err := <-loginErrorChan:
		return false, fmt.Errorf("登录失败: %v", err)
	case <-time.After(30 * time.Second):
		// 检查页面状态作为后备方案
		currentURL := page.URL()
		if strings.Contains(currentURL, "index") {
			log.Printf("超时但已跳转到目标页面: %s", currentURL)
		} else {
			// 保存截图用于调试
			if _, err := page.Screenshot(playwright.PageScreenshotOptions{
				Path: playwright.String("login_timeout.png"),
			}); err == nil {
				log.Println("已保存超时截图: login_timeout.png")
			}
			return false, fmt.Errorf("登录超时，当前URL: %s", currentURL)
		}
	}

	return true, nil
}

// saveOrUpdateUser 保存或更新用户信息
func saveOrUpdateUser(username, password, cookie string, expires float64) error {
	// 检查用户是否存在
	userInfo, err := UserModel.Api_find_by_username(username)
	if err == nil {
		if userInfo.LoginName == username {
			// 用户存在，更新信息
			userInfo.Cookie = cookie
			userInfo.Expires = expires
			updateData := map[string]interface{}{
				"login_pass": password,
				"cookie":     userInfo.Cookie,
				"expires":    userInfo.Expires,
			}

			// 调用 Api_update 方法
			if err := UserModel.Api_update(username, updateData); err != nil {
				return fmt.Errorf("更新用户信息失败: %v", err)
			}
			return nil
		} else {
			// 用户不存在，创建新用户
			newUser := UserModel.UserInfo{
				LoginName: username,
				LoginPass: password,
				Cookie:    cookie,
				Expires:   expires,
			}

			// 直接调用 Api_insert 方法
			success := UserModel.Api_insert(newUser)
			if !success {
				return fmt.Errorf("创建新用户失败")
			}
			return nil
		}
	}
	return err
}

// saveOrShop 保存或更新用户店铺
func saveOrShop(shops []ShopModel.Account, username string) error {
	// 检查用户是否存在
	for _, shop := range shops {
		shop.LoginName = username
		shopInfo := ShopModel.Api_find_struct(shop.AccountName)
		if shopInfo.AccountName == username || shopInfo.LoginName == username {
			// 更新用户表中的cookie字段
			updateData := map[string]interface{}{
				"account_id":   shop.AccountID,
				"account_name": shop.AccountName,
				"login_name":   shop.LoginName,
			}
			if err := ShopModel.Api_update(shop.SubjectID, updateData); err != nil {
				log.Printf("保存店铺失败: %v\n", err)
				return fmt.Errorf("保存店铺失败: %v", err)
			} else {
				log.Println("店铺数据已保存到数据库")
			}
		} else {
			success := ShopModel.Api_insert(shop)
			if !success {
				return fmt.Errorf("创建新用户失败")
			}
		}
	}
	return nil
}
