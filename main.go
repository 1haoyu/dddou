package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"main.go/dataModel/CookieModel"
	"main.go/dataModel/ShopModel"
	"main.go/dataModel/SkuModel"
	"main.go/dataModel/UserModel"
	"main.go/internal/downloader"
	"main.go/internal/ocr"
	"main.go/internal/service"
	"main.go/internal/types"
	"main.go/res"
	"main.go/tuuz/database"
)

// 添加状态检查快捷键
func addStateDebugShortcut(window fyne.Window, appState *types.AppState) {
	window.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if ev.Name == fyne.KeyF5 {
			refreshLeftPanel(appState)
			appState.StatusBar.SetText("手动刷新UI")
		} else if ev.Name == fyne.KeyS {
			fmt.Println("===== 应用状态快照 =====")
			fmt.Printf("当前用户: %s\n", appState.CurrentUser.LoginName)
			fmt.Printf("店铺数量: %d\n", len(appState.Shops))
			fmt.Printf("最后刷新时间: %s\n", appState.LastRefreshTime.Format("15:04:05.000"))
			fmt.Println("=======================")
		}
	})
}

// 添加日志记录函数
func appendLog(appState *types.AppState, message string) {
	fyne.DoAndWait(func() {
		// 添加带时间戳的消息
		timestamp := time.Now().Format("15:04:05.000")
		logMessage := fmt.Sprintf("[%s] %s\n", timestamp, message)

		// 追加新日志并滚动到底部
		appState.LogEntry.SetText(appState.LogEntry.Text + logMessage)
		appState.LogEntry.CursorRow = len(strings.Split(appState.LogEntry.Text, "\n"))
	})
}

func main() {
	os.Setenv("PLAYWRIGHT_BROWSERS_PATH", "./browsers")
	database.Init()
	UserModel.UserInit()
	ShopModel.ShopInit()
	CookieModel.CreateCookieInfoTable()
	SkuModel.ProductInit()
	// 创建缓存目录
	if err := os.MkdirAll("cacheimg", 0755); err != nil {
		log.Printf("创建缓存目录失败: %v", err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("店铺管理工具")
	myWindow.Resize(fyne.NewSize(1200, 800))

	// 初始化Playwright服务
	pwService, err := service.InitPlaywrightService()
	if err != nil {
		log.Fatalf("初始化Playwright失败: %v", err)
	}
	defer func() {
		if err := pwService.Browser.Close(); err != nil {
			log.Printf("关闭浏览器失败: %v", err)
		}
		// 修复错误1: 使用正确的停止方法
		if err := pwService.PW.Stop(); err != nil {
			log.Printf("停止Playwright失败: %v", err)
		}
	}()

	// 初始化应用状态
	appState := types.NewAppState(pwService)
	appState.MyWindow = myWindow // 关键：将窗口对象存入state
	appState.MyApp = myApp       // 关键：将窗口对象存入state

	appState.FilterFilePath = getDefaultFilterPath()
	// 注册调试快捷键
	addStateDebugShortcut(myWindow, appState)
	// 启动状态监听器
	startStateListener(appState)
	// 尝试加载默认过滤文件
	go loadFilterFile(appState)

	// 创建状态栏
	appState.StatusBar = widget.NewLabel("就绪")

	// 创建URL访问控件
	// 创建日志显示控件 - 替代原来的 URL 控件
	appState.LogEntry = widget.NewMultiLineEntry()
	// appState.LogEntry.SetReadOnly(true) // 设置为只读
	appState.LogEntry.Wrapping = fyne.TextWrapWord // 自动换行

	// 创建带滚动条的日志容器
	logScroll := container.NewScroll(appState.LogEntry)
	logScroll.SetMinSize(fyne.NewSize(0, 150)) // 固定高度150

	// 创建底部区域（状态栏 + URL控件）
	bottomArea := container.NewVBox(
		widget.NewLabel("运行日志:"),
		logScroll,
		widget.NewSeparator(),
		container.NewHBox(layout.NewSpacer(), appState.StatusBar),
	)

	// 创建主布局
	mainContent := createMainUI(myWindow, appState)

	// 设置整体布局
	content := container.NewBorder(
		nil,        // 顶部
		bottomArea, // 底部（包含URL控件和状态栏）
		nil,        // 左侧
		nil,        // 右侧
		mainContent,
	)
	myWindow.SetContent(content)

	// 启动时尝试自动登录
	go tryAutoLogin(appState)

	myWindow.ShowAndRun()
	defer appState.OCREngine.Close()
}

// 新增状态监听器 - 定期检查状态变化
func startStateListener(appState *types.AppState) {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond) // 每100ms检查一次
			if appState.NeedsRefresh {
				fyne.DoAndWait(func() {
					refreshLeftPanel(appState)
					appState.NeedsRefresh = false
				})
			}
		}
	}()
}

// 获取默认过滤文件路径
func getDefaultFilterPath() string {
	if runtime.GOOS == "windows" {
		return filepath.Join("filter.txt")
	}
	return filepath.Join(os.Getenv("HOME"), "filter.txt")
}

// 修改 refreshAllProductTabs 函数
func refreshAllProductTabs(appState *types.AppState) {
	if appState.ProductTabs == nil || len(appState.ProductTabs.Items) == 0 {
		return
	}

	// 遍历所有标签页并刷新
	for _, tab := range appState.ProductTabs.Items {
		// 通过标签页标题获取店铺
		shop, exists := appState.TabShopMap[tab.Text]
		if !exists {
			continue
		}

		// 重新加载商品
		go func(shop ShopModel.Account) {
			products, err := loadProductsForShop(shop, appState)
			if err != nil {
				fyne.DoAndWait(func() {
					appState.StatusBar.SetText(fmt.Sprintf("刷新 %s 商品失败: %s", shop.AccountName, err.Error()))
				})
				return
			}

			fyne.DoAndWait(func() {
				// 更新标签页内容
				tab.Content = container.NewMax(createProductTable(products))
				appState.ProductTabs.Refresh()
				appState.StatusBar.SetText(fmt.Sprintf("已刷新 %s 的商品", shop.AccountName))
			})
		}(shop)
	}
}

// 加载过滤文件
func loadFilterFile(appState *types.AppState) {
	if appState.FilterFilePath == "" {
		log.Printf("加载本地过滤文件失败: %s", appState.FilterFilePath)
		return
	}

	if _, err := os.Stat(appState.FilterFilePath); os.IsNotExist(err) {
		err := os.WriteFile(appState.FilterFilePath, []byte{}, 0644)
		if err != nil {
			log.Printf("创建过滤文件失败: %v", err)
		}
		return
	}

	content, err := os.ReadFile(appState.FilterFilePath)
	if err != nil {
		log.Printf("读取过滤文件失败: %v", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	appState.FilterKeywords = []string{}
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			appState.FilterKeywords = append(appState.FilterKeywords, trimmed)
		}
	}

	// 关键修复：启动工作协程池
	downloader.StartDownloadWorkers(appState) // 启动下载协程
	ocr.StartOCRWorkers(appState)             // 启动OCR协程

	fyne.DoAndWait(func() {
		appState.StatusBar.SetText(fmt.Sprintf("已加载 %d 个过滤关键字", len(appState.FilterKeywords)))
		// 更新关键字数量标签
		if appState.KeywordCount != nil { // 修正为 KeywordCount
			appState.KeywordCount.SetText(fmt.Sprintf("关键字数量: %d", len(appState.FilterKeywords)))
		}
		// 刷新所有已打开的商品标签页
		refreshAllProductTabs(appState)
	})
}

// 修改 createMainUI 函数 - 保存分割布局引用
func createMainUI(window fyne.Window, appState *types.AppState) fyne.CanvasObject {
	appState.MyWindow = window

	// 创建整个左侧面板
	leftPanel := createLeftPanel(appState)
	appState.LeftPanel = leftPanel

	// 右侧面板
	appState.ProductTabs = container.NewAppTabs()
	appState.ProductTabs.SetTabLocation(container.TabLocationTop)

	rightPanel := container.NewBorder(
		widget.NewLabelWithStyle("商品信息", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		nil, nil, nil,
		container.NewMax(appState.ProductTabs),
	)

	// 使用HSplit布局 - 保存引用
	split := container.NewHSplit(leftPanel, rightPanel)
	split.SetOffset(0.25)
	appState.SplitContainer = split // 保存分割布局引用

	return split
}

// 修改createFilterPanel函数 - 返回容器并保存引用
func createFilterPanel(appState *types.AppState) *fyne.Container {
	// 创建文件路径标签
	pathLabel := widget.NewLabel("过滤文件: " + appState.FilterFilePath)
	pathLabel.Wrapping = fyne.TextWrapWord

	// 创建选择文件按钮
	selectButton := widget.NewButton("选择过滤文件", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, appState.MyWindow)
				return
			}
			if reader == nil {
				return // 用户取消
			}

			// 更新文件路径
			appState.FilterFilePath = reader.URI().Path()
			pathLabel.SetText("过滤文件: " + appState.FilterFilePath)

			// 加载过滤文件
			go func() {
				loadFilterFile(appState)

				// 刷新所有已打开的商品标签页
				refreshAllProductTabs(appState)
			}()
		}, appState.MyWindow)
	})

	// 创建刷新按钮
	refreshButton := widget.NewButton("刷新过滤", func() {
		// if appState.FilterFilePath != "" {
		// 	appState.StatusBar.SetText("刷新过滤关键字...")
		// 	go func() {
		// 		loadFilterFile(appState)

		// 		// 刷新所有已打开的商品标签页
		// 		refreshAllProductTabs(appState)
		// 	}()
		// } else {
		// 	appState.StatusBar.SetText("请先选择过滤文件")
		// }

		// sku, err := SkuModel.Api_find_by_id("3766017532163653929")
		// if err == nil {
		// 	downloader.DoOcrImages(sku, appState)
		// }

		// skus, total, err := SkuModel.Api_select_all("229940601")
		// if err == nil {
		// 	log.Printf("加载商品总数: %d", total)
		// 	for _, sku := range skus {
		// 		log.Printf("检测商品信息: %s", sku.Name)
		// 		//downloader.DoOcrImages(sku, appState)
		// 	}
		// }
		// all := SkuModel.Api_count(229940601)
		// log.Printf("加载商品总数: %d", all)

		// skus, total := SkuModel.Api_select(229940601, 20, 0)
		// if total > 0 {
		// 	log.Printf("加载商品总数: %d", total)
		// 	for _, sku := range skus {
		// 		log.Printf("检测商品信息: %s", sku.Name)
		// 		//downloader.DoOcrImages(sku, appState)
		// 	}
		// }
	})

	// 创建"增加商品"按钮
	addProductsButton := widget.NewButton("增加商品", func() {
		if appState.ProductTabs.Selected() == nil {
			appState.StatusBar.SetText("请先选择一个店铺标签页")
			return
		}

		// shopName := appState.ProductTabs.Selected().Text
		shopInfo := appState.ShopCookieInfo
		if shopInfo.Token == "" || shopInfo.VerifyFp == "" {
			appState.StatusBar.SetText("未获取授权信息")
			return
		}
		unixTime := time.Now().Unix()
		if shopInfo.Expires < float64(unixTime) {
			appState.StatusBar.SetText("授权已经过期，需要重新获取")
			return
		}

		// var productUrl = "https://fxg.jinritemai.com/product/tproduct/list?page=0&pageSize=20&draft_status=0&comment_percent=&group_id=&sku_type=&tab=onSale&business_type=4&is_online=1&not_for_sale_search_type=1&from_mng=1&check_status=3&status=0&supply_status=&need_auto_rectify_info=true&need_pay_no_stock_skus=true&order_field=audit_time&sort=desc&appid=1&__token=c13037160dfe5d7ff679cc12b61502d8&_bid=ffa_goods&_lid=855297703397&verifyFp=verify_md574s92_Ywos6QbE_qhu8_4CB7_8gRt_hGLiRUCjakuW&fp=verify_md574s92_Ywos6QbE_qhu8_4CB7_8gRt_hGLiRUCjakuW&msToken=N3amr2NW37yQqRuzfmpvcJpmsNLBN7X5NeZXxt_0R75U5frDvIetba2u6NqaAZlCtEd8looeL4PGEQNipBA1nDJOXAD7NTtiRKOysiM2p2TRrbeXVWMla5BCGp-oUGg0fk-Zdzf6r34UEtIAOM_JD63BKwq3XjiEXwAHkqi2xg%3D%3D&a_bogus=xvmwBmzhDk6TfVDk5WKLfY3qIWF3YpdC0Gi0MDZMXVvrRL39HMOn9exE9s4vOxjjis%2FmIe8jy4hjTNMMx5%2FyA3vRHuDKUIcgmESDeM32so0j5H4yuy6QnGJx4vJlFeeQ5i53Ec7MqJKcFYmk09Q95kI6PEVja3Lk96EtrNqL2o8W"
		var productUrl = fmt.Sprintf("https://fxg.jinritemai.com/product/tproduct/list?page=0&pageSize=%d&draft_status=0&comment_percent=&group_id=&sku_type=&tab=onSale&business_type=4&is_online=1&not_for_sale_search_type=1&from_mng=1&check_status=3&status=0&supply_status=&need_auto_rectify_info=true&need_pay_no_stock_skus=true&order_field=audit_time&sort=desc&appid=1&__token=%s&_bid=ffa_goods&verifyFp=%s&fp=%s", 50, shopInfo.Token, shopInfo.VerifyFp, shopInfo.VerifyFp)
		appState.StatusBar.SetText(fmt.Sprintf("正在访问: %s...", productUrl))

		go func() {
			// 访问URL
			// 修复错误2和3: 使用正确的返回类型和状态码访问方式
			response, err := service.VisitUrlWithPlaywright(appState, productUrl)
			if err != nil {
				fyne.DoAndWait(func() {
					appState.StatusBar.SetText(fmt.Sprintf("访问失败: %v", err))
				})
				return
			}

			fyne.DoAndWait(func() {
				body, err := response.Body()
				if err != nil {
					log.Printf("获取响应体失败: %v", err)
					return
				}
				var data SkuModel.Response
				err = json.Unmarshal(body, &data)
				if err != nil {
					return
				}
				for _, task := range data.Data {
					success := SkuModel.Api_insert(task)
					if !success {
						log.Printf("插入商品数据错误: %s", task.ProductID)
					}
				}
				// 使用Status()方法获取状态码
				appState.StatusBar.SetText(fmt.Sprintf("访问完成! 状态码: %d", response.Status()))
			})
		}()
	})

	// 修改按钮容器，添加新按钮
	buttonContainer := container.NewHBox(
		selectButton,
		refreshButton,
		addProductsButton, // 新增按钮
	)

	// 创建关键字计数标签 - 保存引用
	keywordCount := widget.NewLabel(fmt.Sprintf("关键字数量: %d", len(appState.FilterKeywords)))
	keywordCount.TextStyle = fyne.TextStyle{Bold: true}
	appState.KeywordCount = keywordCount

	// 创建面板容器
	panel := container.NewVBox(
		widget.NewSeparator(),
		widget.NewLabelWithStyle("商品过滤", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		pathLabel,
		keywordCount,
		buttonContainer,
	)

	return panel
}

// 修改 createLoggedInPanel 函数 - 确保注销时直接刷新
func createLoggedInPanel(appState *types.AppState) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabelWithStyle("登录状态", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewLabel("用户:"),
			widget.NewLabel(appState.CurrentUser.LoginName),
		),
		container.NewHBox(
			widget.NewLabel("店铺数量:"),
			widget.NewLabel(fmt.Sprintf("%d", len(appState.Shops))),
		),
		widget.NewSeparator(),
		container.NewCenter(
			widget.NewButton("注销", func() {
				// 重置状态
				appState.CurrentUser = UserModel.UserInfo{}
				appState.Shops = nil
				appState.ProductTabs.Items = nil
				appState.ProductTabs.Refresh()
				appState.TabShopMap = make(map[string]ShopModel.Account)

				// 直接调用刷新函数
				refreshLeftPanel(appState)
				appState.StatusBar.SetText("已注销")
			}),
		),
	)
}

// 重构创建顶部面板函数 - 确保状态正确反映
func createTopPanel(appState *types.AppState) *fyne.Container {
	// 添加调试日志
	fmt.Printf("创建顶部面板: 登录状态=%t, 用户名=%s\n",
		appState.CurrentUser.LoginName != "",
		appState.CurrentUser.LoginName)

	var content fyne.CanvasObject

	if appState.CurrentUser.LoginName != "" {
		content = createLoggedInPanel(appState)
	} else {
		content = createLoginForm(appState)
	}

	return container.NewMax(content)
}

// 重构 createContentPanel 函数 - 添加详细日志
func createContentPanel(appState *types.AppState) *fyne.Container {
	// 添加详细调试日志
	fmt.Printf("创建内容面板: 登录状态=%t, 用户名=%s, 店铺数量=%d\n",
		appState.CurrentUser.LoginName != "",
		appState.CurrentUser.LoginName,
		len(appState.Shops))

	if appState.CurrentUser.LoginName != "" {
		if len(appState.Shops) > 0 {
			return createShopListPanel(appState)
		}
		return container.NewCenter(
			widget.NewLabel("没有可用的店铺"),
		)
	}
	return container.NewCenter(
		widget.NewLabel("请先登录查看店铺列表"),
	)
}

// 重构刷新函数 - 确保完全重建UI
func refreshLeftPanel(appState *types.AppState) {
	if appState.SplitContainer == nil {
		return
	}
	// 添加详细调试信息
	fmt.Printf("刷新左侧面板 - 时间: %s, 用户: %s, 店铺数量: %d\n",
		time.Now().Format("15:04:05.000"),
		appState.CurrentUser.LoginName,
		len(appState.Shops))
	// 创建新的左侧面板
	newLeftPanel := createLeftPanel(appState)
	// 添加调试背景色（登录状态不同颜色不同）
	var debugColor color.Color
	if appState.CurrentUser.LoginName != "" {
		debugColor = color.NRGBA{R: 0, G: 100, B: 0, A: 30} // 登录状态绿色半透明
	} else {
		debugColor = color.NRGBA{R: 100, G: 0, B: 0, A: 30} // 未登录状态红色半透明
	}

	debugPanel := container.NewMax(
		canvas.NewRectangle(debugColor),
		newLeftPanel,
	)

	// 替换分割布局中的左侧面板
	appState.SplitContainer.Leading = debugPanel
	appState.LeftPanel = debugPanel

	// 刷新分割布局
	appState.SplitContainer.Refresh()

	// 强制重绘整个窗口
	appState.MyWindow.Content().Refresh()
	appState.LastRefreshTime = time.Now()
}

// 重构 createLeftPanel 函数 - 确保使用正确的状态
func createLeftPanel(appState *types.AppState) *fyne.Container {
	// 创建顶部面板（用户状态/登录表单）
	topPanel := createTopPanel(appState)

	// 创建内容面板（店铺列表或提示）
	contentPanel := createContentPanel(appState)

	// 创建过滤面板
	filterPanel := createFilterPanel(appState)

	// 使用Border布局
	return container.NewBorder(
		topPanel,    // 顶部
		filterPanel, // 底部
		nil, nil,    // 左右
		contentPanel, // 中间内容
	)
}

// 修改登录按钮回调 - 确保状态正确更新
func createLoginForm(appState *types.AppState) fyne.CanvasObject {
	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	usernameEntry.PlaceHolder = "输入邮箱地址"
	passwordEntry.PlaceHolder = "输入密码"

	// 登录按钮回调
	loginButton := widget.NewButton("登录", func() {
		appState.StatusBar.SetText("登录中...")
		appendLog(appState, "登录中...")
		go func() {
			// 模拟网络延迟
			// time.Sleep(500 * time.Millisecond)
			_, err := service.LoginOrAutoLogin(usernameEntry.Text, passwordEntry.Text, appState)
			if err != nil {
				appState.StatusBar.SetText("登录失败: " + err.Error())
				return
			}
			// 获取店铺信息
			shops := ShopModel.Api_select_struct(nil)

			fyne.DoAndWait(func() {
				if len(shops) == 0 {
					appState.StatusBar.SetText("获取店铺信息为空")
					return
				}

				// 更新应用状态
				appState.Shops = shops
				appState.CurrentUser, _ = UserModel.Api_find_by_username(usernameEntry.Text)
				// 更新店铺列表绑定
				updateShopListBinding(appState) // 新增：更新绑定数据
				// 添加状态更新日志
				fmt.Printf("登录成功 - 用户: %s, 店铺数量: %d\n",
					appState.CurrentUser.LoginName,
					len(appState.Shops))
				if appState.CurrentUser.LoginName == "" {
					appState.CurrentUser.LoginName = "1"
				}
				appState.StatusBar.SetText(fmt.Sprintf("登录成功! 共 %d 个店铺", len(shops)))

				// 直接刷新UI
				refreshLeftPanel(appState)
			})
		}()
	})

	form := widget.NewForm(
		widget.NewFormItem("邮箱:", usernameEntry),
		widget.NewFormItem("密码:", passwordEntry),
	)
	appState.LoginForm = form

	return container.NewVBox(
		widget.NewLabelWithStyle("登录面板", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
		container.NewCenter(loginButton),
	)
}

// 修改自动登录函数 - 添加详细日志
func tryAutoLogin(appState *types.AppState) {
	// 获取所有用户
	users := UserModel.Api_select_struct(nil)
	if len(users) == 0 {
		fyne.DoAndWait(func() {
			appState.StatusBar.SetText("获取已经存在的账号为空")
		})
		return
	}

	// 尝试使用第一个用户自动登录
	user := users[0]

	fyne.DoAndWait(func() {
		appState.StatusBar.SetText(fmt.Sprintf("尝试自动登录: %s...", user.LoginName))
		appendLog(appState, fmt.Sprintf("尝试自动登录: %s...", user.LoginName))
	})

	// 获取用户名输入框
	if appState.LoginForm == nil || len(appState.LoginForm.Items) < 2 {
		fyne.DoAndWait(func() {
			appState.StatusBar.SetText("自动登录失败: 登录表单尚未初始化")
		})
		return
	}

	usernameItem := appState.LoginForm.Items[0]
	usernameEntry, ok := usernameItem.Widget.(*widget.Entry)
	if !ok {
		fyne.DoAndWait(func() {
			appState.StatusBar.SetText("自动登录失败: 用户名控件类型错误")
		})
		return
	}

	passwordItem := appState.LoginForm.Items[1]
	passwordEntry, ok := passwordItem.Widget.(*widget.Entry)
	if !ok {
		fyne.DoAndWait(func() {
			appState.StatusBar.SetText("自动登录失败: 密码控件类型错误")
		})
		return
	}

	// 触发登录
	fyne.DoAndWait(func() {
		usernameEntry.SetText(user.LoginName)
		passwordEntry.SetText(user.LoginPass)
		appState.StatusBar.SetText("正在自动登录...")
		appendLog(appState, "正在自动登录...")

		// 更新应用状态
		appState.CurrentUser = user
		appState.Shops = ShopModel.Api_select_struct(nil)
		// 更新店铺列表绑定
		updateShopListBinding(appState) // 新增
		// 添加自动登录日志
		fmt.Printf("自动登录成功 - 用户: %s, 店铺数量: %d\n",
			appState.CurrentUser.LoginName,
			len(appState.Shops))

		// 直接刷新UI
		refreshLeftPanel(appState)
	})
}

// 修改后的异步加载店铺头像函数
func loadShopAvatar(img *canvas.Image, url string) {
	if url == "" {
		// 使用默认头像
		fyne.DoAndWait(func() {
			img.Resource = fyne.Theme.Icon(fyne.CurrentApp().Settings().Theme(), "account")
			img.Refresh()
		})
		return
	}

	// 创建HTTP客户端（可设置超时）
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("加载头像失败: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("头像请求失败: %s", resp.Status)
		return
	}

	// 解码图片
	imgData, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Printf("解码头像失败: %v", err)
		return
	}

	// 在主线程更新UI
	fyne.DoAndWait(func() {
		img.Image = imgData
		img.Refresh()
	})
}

// 修改后的 createShopListPanel 函数
func createShopListPanel(appState *types.AppState) *fyne.Container {
	// 创建绑定数据
	if appState.ShopListBinding == nil {
		appState.ShopListBinding = binding.NewUntypedList()
	} else {
		// 确保绑定数据是最新的
		updateShopListBinding(appState)
	}

	// 创建列表控件
	list := widget.NewListWithData(
		appState.ShopListBinding,
		func() fyne.CanvasObject {
			avatar := canvas.NewImageFromResource(nil)
			avatar.SetMinSize(fyne.NewSize(40, 40))
			avatar.FillMode = canvas.ImageFillContain
			nameLabel := widget.NewLabel("")
			statusIcon := widget.NewIcon(nil)

			return container.NewHBox(
				avatar,
				container.NewVBox(nameLabel),
				layout.NewSpacer(),
				statusIcon,
			)
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			hbox, ok := obj.(*fyne.Container)
			if !ok || len(hbox.Objects) < 4 {
				return
			}

			avatar, _ := hbox.Objects[0].(*canvas.Image)
			nameContainer, _ := hbox.Objects[1].(*fyne.Container)
			nameLabel, _ := nameContainer.Objects[0].(*widget.Label)
			statusIcon, _ := hbox.Objects[3].(*widget.Icon)

			val, err := item.(binding.Untyped).Get()
			if err != nil {
				return
			}

			shop, ok := val.(ShopModel.Account)
			if !ok {
				return
			}

			nameLabel.SetText(shop.AccountName)
			if shop.CanLogin {
				statusIcon.SetResource(res.ResShuffleSvg)
			} else {
				statusIcon.SetResource(fyne.Theme.Icon(fyne.CurrentApp().Settings().Theme(), "error"))
			}
			go loadShopAvatar(avatar, shop.AccountAvatar)
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		if id < 0 || id >= len(appState.Shops) {
			return
		}

		shop := appState.Shops[id]
		shopCookie := CookieModel.Api_find_struct_by_id(shop.AccountID)
		if shopCookie.SubjectID != shop.SubjectID {
			service.LoginOrAutoLogin(shop.LoginName, "", appState)
			return
		}
		appState.ShopCookieInfo = shopCookie
		appState.StatusBar.SetText(fmt.Sprintf("加载 %s 的商品...", shop.AccountName))

		go func() {
			products, err := loadProductsForShop(shop, appState)
			if err != nil {
				fyne.DoAndWait(func() {
					appState.StatusBar.SetText("加载商品失败: " + err.Error())
				})
				return
			}

			fyne.DoAndWait(func() {
				appState.StatusBar.SetText(fmt.Sprintf("已加载 %d 个商品", len(products)))
				addOrUpdateProductTab(appState, shop, products)
			})
		}()
	}

	// 创建滚动容器 - 设置最小高度确保可滚动
	scrollContainer := container.NewScroll(list)
	scrollContainer.SetMinSize(fyne.NewSize(280, 200)) // 最小高度200确保可滚动

	// 使用Max容器确保填充空间
	return container.NewMax(
		container.NewBorder(
			widget.NewLabel("店铺列表"),
			nil, nil, nil,
			scrollContainer,
		),
	)
}

// 更新店铺列表绑定数据
func updateShopListBinding(appState *types.AppState) {
	if appState.ShopListBinding == nil {
		appState.ShopListBinding = binding.NewUntypedList()
	}

	values := make([]interface{}, len(appState.Shops))
	for i, shop := range appState.Shops {
		values[i] = shop
	}

	appState.ShopListBinding.Set(values)
}

// 应用商品过滤
func applyProductFilter(products []SkuModel.DataItem, keywords []string) []SkuModel.DataItem {
	if len(keywords) == 0 {
		return products // 没有关键字，返回所有商品
	}

	filtered := []SkuModel.DataItem{}
	for _, product := range products {
		exclude := false
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
				exclude = true
				break
			}
		}

		if !exclude {
			filtered = append(filtered, product)
		}
	}

	return filtered
}

// 修改 loadProductsForShop 函数，生成更多模拟数据
func loadProductsForShop(shop ShopModel.Account, appState *types.AppState) ([]SkuModel.DataItem, error) {
	// all := SkuModel.Api_count(229940601)
	// log.Printf("加载商品总数: %d", all)
	appendLog(appState, fmt.Sprintf("检测商品信息: %s", appState.ShopCookieInfo.AccountID))
	var pageIndex = 0
	pagination, exists := appState.PaginationStates[shop.AccountName]
	if exists {
		pageIndex = pagination.CurrentPage
	}
	products, total := SkuModel.Api_select(appState.ShopCookieInfo.AccountID, 20, pageIndex)
	if total > 0 {
		log.Printf("加载商品总数: %d", total)
		for _, sku := range products {
			appendLog(appState, fmt.Sprintf("检测商品信息: %s", sku.Name))
			//log.Printf("检测商品信息: %s", sku.Name)
			//downloader.DoOcrImages(sku, appState)
		}
	}
	// 应用过滤
	filteredProducts := applyProductFilter(products, appState.FilterKeywords)

	return filteredProducts, nil
}

// 修改 addOrUpdateProductTab 函数，添加分页支持
func addOrUpdateProductTab(appState *types.AppState, shop ShopModel.Account, products []SkuModel.DataItem) {
	tabTitle := shop.AccountName

	// 获取或创建分页状态
	pagination, exists := appState.PaginationStates[tabTitle]
	if !exists {
		// 初始化分页状态
		pagination = &types.PaginationState{
			PageSize:      10,
			CurrentPage:   1,
			TotalProducts: len(products),
			PageInfo:      binding.NewString(), // 关键修复：初始化PageInfo
		}
		// 计算总页数
		pagination.TotalPages = (pagination.TotalProducts + pagination.PageSize - 1) / pagination.PageSize
		if pagination.TotalPages == 0 {
			pagination.TotalPages = 1
		}
		// 设置初始分页信息
		pagination.PageInfo.Set(fmt.Sprintf("第 %d 页/共 %d 页", pagination.CurrentPage, pagination.TotalPages))
		appState.PaginationStates[tabTitle] = pagination
	} else {
		// 更新商品总数
		pagination.TotalProducts = len(products)
	}

	// 计算总页数
	pagination.TotalPages = (pagination.TotalProducts + pagination.PageSize - 1) / pagination.PageSize
	if pagination.TotalPages == 0 {
		pagination.TotalPages = 1
	}
	// 获取当前页数据
	currentPageProducts := getCurrentPageProducts(pagination, products)

	// 检查是否已存在该TAB
	for _, tab := range appState.ProductTabs.Items {
		if tab.Text == tabTitle {
			// 修改调用，传入店铺名称
			tab.Content = createProductListWithPagination(appState, currentPageProducts, tabTitle, products)
			// 更新映射
			appState.TabShopMap[tabTitle] = shop
			appState.ProductTabs.Refresh()
			return
		}
	}

	// 创建新TAB
	newTab := container.NewTabItem(
		shop.AccountName,
		createProductDetailView(appState, shop, products),
	)
	// 添加到映射
	appState.TabShopMap[tabTitle] = shop
	appState.ProductTabs.Append(newTab)
	appState.ProductTabs.Select(newTab)
}

// 创建商品详情视图
func createProductDetailView(appState *types.AppState, shop ShopModel.Account, products []SkuModel.DataItem) fyne.CanvasObject {
	// 主容器
	mainContainer := container.NewVBox()

	// 店铺信息面板
	shopInfo := container.NewHBox(
		widget.NewLabel(fmt.Sprintf("店铺: %s", shop.AccountName)),
		layout.NewSpacer(),
		widget.NewLabel(fmt.Sprintf("商品总数: %d", len(products))),
	)

	mainContainer.Add(shopInfo)
	mainContainer.Add(widget.NewSeparator())

	// 商品表格
	table := createProductTable(products)
	mainContainer.Add(table)

	// 分页控件
	pagination := createPaginationControls(appState, shop.AccountName, products)
	mainContainer.Add(pagination)

	return container.NewPadded(mainContainer)
}

// 修改 getCurrentPageProducts 函数
func getCurrentPageProducts(pagination *types.PaginationState, products []SkuModel.DataItem) []SkuModel.DataItem {
	start := (pagination.CurrentPage - 1) * pagination.PageSize
	if start >= len(products) {
		start = 0
	}

	end := start + pagination.PageSize
	if end > len(products) {
		end = len(products)
	}

	return products[start:end]
}

// 修改 createProductListWithPagination 函数
func createProductListWithPagination(appState *types.AppState, currentPageProducts []SkuModel.DataItem, shopName string, allProducts []SkuModel.DataItem) fyne.CanvasObject {
	// 创建表格
	table := createProductTable(currentPageProducts)

	// 创建分页控件 - 传入店铺名称
	pagination := createPaginationControls(appState, shopName, allProducts)

	// 创建布局：表格在上，分页控件在下
	return container.NewBorder(nil, pagination, nil, nil, table)
}

// 定义固定行高布局
type fixedHeightLayout struct {
	height float32
}

func (f *fixedHeightLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, o := range objects {
		o.Resize(fyne.NewSize(size.Width, f.height))
	}
}

func (f *fixedHeightLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, f.height)
}

// 修改createProductTable函数，添加图片列
func createProductTable(products []SkuModel.DataItem) fyne.CanvasObject {
	// 创建表格
	table := widget.NewTable(
		func() (int, int) {
			return len(products) + 1, 5 // 增加一列用于显示图片
		},
		func() fyne.CanvasObject {
			// 使用固定高度的容器包装HBox
			hbox := container.NewHBox()
			return container.New(&fixedHeightLayout{height: 60}, hbox) // 设置固定行高为60像素
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			// 获取包含HBox的固定高度容器
			fixedContainer := cell.(*fyne.Container)
			// 获取HBox容器
			hbox := fixedContainer.Objects[0].(*fyne.Container)
			// 清空容器
			hbox.Objects = nil

			if id.Row == 0 {
				// 表头
				switch id.Col {
				case 0:
					hbox.Add(widget.NewLabel("商品ID"))
				case 1:
					hbox.Add(widget.NewLabel("商品名称"))
				case 2:
					hbox.Add(widget.NewLabel("价格"))
				case 3:
					hbox.Add(widget.NewLabel("图片"))
				case 4:
					hbox.Add(widget.NewLabel("库存"))
				}
				return
			}

			// 数据行
			if id.Row-1 >= len(products) {
				return
			}
			product := products[id.Row-1]
			switch id.Col {
			case 0:
				hbox.Add(widget.NewLabel(product.ProductID))
			case 1:
				hbox.Add(widget.NewLabel(product.Name))
			case 2:
				hbox.Add(widget.NewLabel(fmt.Sprintf("¥%.2f", float64(product.MarketPrice)/100)))
			case 3: // 图片列
				// 最多显示4张图片
				maxDisplay := 4
				if len(product.Pics) < maxDisplay {
					maxDisplay = len(product.Pics)
				}

				for i := 0; i < maxDisplay; i++ {
					if i >= len(product.Pics) {
						break
					}

					// 使用异步图片组件
					img := NewAsyncImage(product.Pics[i])
					hbox.Add(img)
				}
			case 4:
				hbox.Add(widget.NewLabel(fmt.Sprintf("%d", product.DiscountPrice)))
			}
		},
	)

	// 设置列宽
	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 300)
	table.SetColumnWidth(2, 100)
	table.SetColumnWidth(3, 180) // 图片列需要更宽
	table.SetColumnWidth(4, 100)

	// 创建滚动容器
	scrollContainer := container.NewScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(800, 400)) // 增加宽度以适应新列
	return scrollContainer
}

// 修改 createPaginationControls 函数
func createPaginationControls(appState *types.AppState, shopName string, allProducts []SkuModel.DataItem) *fyne.Container {
	// 获取该店铺的分页状态
	pagination, exists := appState.PaginationStates[shopName]
	if !exists {
		// 如果不存在，创建默认状态
		pagination = &types.PaginationState{
			PageSize:      10,
			CurrentPage:   1,
			TotalProducts: len(allProducts),
			PageInfo:      binding.NewString(), // 关键修复：初始化PageInfo
		}
		pagination.TotalPages = (len(allProducts) + pagination.PageSize - 1) / pagination.PageSize
		if pagination.TotalPages == 0 {
			pagination.TotalPages = 1
		}
		pagination.PageInfo.Set(fmt.Sprintf("第 %d 页/共 %d 页", pagination.CurrentPage, pagination.TotalPages))
		appState.PaginationStates[shopName] = pagination
	}

	// 更新分页信息的函数
	updatePageInfo := func() {
		pagination.PageInfo.Set(fmt.Sprintf("第 %d 页/共 %d 页", pagination.CurrentPage, pagination.TotalPages))
	}

	// 使用闭包捕获当前店铺名称
	refreshForShop := func() {
		refreshCurrentProductTab(appState, shopName, allProducts)
	}

	// 上一页按钮
	prevBtn := widget.NewButton("上一页", func() {
		if pagination.CurrentPage > 1 {
			pagination.CurrentPage--
			updatePageInfo()
			refreshForShop()
		}
	})

	// 页码信息 - 使用绑定标签
	pageInfo := widget.NewLabelWithData(pagination.PageInfo)

	// 下一页按钮
	nextBtn := widget.NewButton("下一页", func() {
		if pagination.CurrentPage < pagination.TotalPages {
			pagination.CurrentPage++
			updatePageInfo()
			refreshForShop()
		}
	})

	// 跳转输入框
	jumpEntry := widget.NewEntry()
	jumpEntry.SetPlaceHolder("页码")
	jumpEntry.Validator = func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return errors.New("请输入数字")
		}
		return nil
	}
	jumpBtn := widget.NewButton("跳转", func() {
		page, err := strconv.Atoi(jumpEntry.Text)
		if err == nil && page >= 1 && page <= pagination.TotalPages {
			pagination.CurrentPage = page
			updatePageInfo()
			refreshForShop()
		}
	})

	// 页面大小选择器
	pageSizeSelect := widget.NewSelect([]string{"5", "10", "20", "50"}, nil)
	pageSizeSelect.SetSelected(fmt.Sprintf("%d", pagination.PageSize))
	pageSizeSelect.OnChanged = func(value string) {
		size, _ := strconv.Atoi(value)
		pagination.PageSize = size
		pagination.CurrentPage = 1
		// 重新计算总页数
		pagination.TotalPages = (len(allProducts) + pagination.PageSize - 1) / pagination.PageSize
		if pagination.TotalPages == 0 {
			pagination.TotalPages = 1
		}
		updatePageInfo()
		refreshForShop()
	}
	pageSizeLabel := widget.NewLabel("每页:")

	// 布局
	return container.NewHBox(
		prevBtn,
		pageSizeLabel,
		pageSizeSelect,
		pageInfo,
		nextBtn,
		jumpEntry,
		jumpBtn,
	)
}

// 修改 refreshCurrentProductTab 函数
func refreshCurrentProductTab(appState *types.AppState, shopName string, allProducts []SkuModel.DataItem) {
	// 获取当前选中的标签页
	currentTab := appState.ProductTabs.Selected()
	if currentTab == nil {
		return
	}

	// 获取该店铺的分页状态
	pagination, exists := appState.PaginationStates[shopName]
	if !exists {
		// 如果不存在，创建默认状态
		pagination = &types.PaginationState{
			PageSize:      10,
			CurrentPage:   1,
			TotalProducts: len(allProducts),
			PageInfo:      binding.NewString(),
		}
		appState.PaginationStates[shopName] = pagination
	}

	// 防止除数为零
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}

	pagination.TotalProducts = len(allProducts)

	// 计算总页数
	pagination.TotalPages = (pagination.TotalProducts + pagination.PageSize - 1) / pagination.PageSize
	if pagination.TotalPages == 0 {
		pagination.TotalPages = 1
	}

	// 确保当前页在有效范围内
	if pagination.CurrentPage > pagination.TotalPages {
		pagination.CurrentPage = pagination.TotalPages
	} else if pagination.CurrentPage < 1 {
		pagination.CurrentPage = 1
	}

	// 更新分页信息
	pagination.PageInfo.Set(fmt.Sprintf("第 %d 页/共 %d 页", pagination.CurrentPage, pagination.TotalPages))

	// 获取当前页数据
	currentPageProducts := getCurrentPageProducts(pagination, allProducts)

	// 检查内容是否真的需要更新
	currentContent := currentTab.Content
	if paginationContent, ok := currentContent.(*fyne.Container); ok {
		if len(paginationContent.Objects) > 0 {
			if tableContainer, ok := paginationContent.Objects[0].(*container.Scroll); ok {
				if existingTable, ok := tableContainer.Content.(*widget.Table); ok {
					// 获取表格的行数
					rows, _ := existingTable.Length()

					// 如果行数相同，只刷新数据
					if rows == len(currentPageProducts)+1 {
						// 使用温和刷新 - 只更新文本内容
						refreshTableData(existingTable, currentPageProducts)
						appState.ProductTabs.Refresh()
						return
					}
				}
			}
		}
	}

	// 需要完全更新内容
	currentTab.Content = createProductListWithPagination(appState, currentPageProducts, shopName, allProducts)
	appState.ProductTabs.Refresh()
}

// 温和刷新 - 只更新文本内容，不重建图片
func refreshTableData(table *widget.Table, products []SkuModel.DataItem) {
	table.Length = func() (int, int) {
		return len(products) + 1, 5
	}

	// 只刷新文本列
	table.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
		fixedContainer := template.(*fyne.Container)
		hbox := fixedContainer.Objects[0].(*fyne.Container)

		if id.Row == 0 {
			return // 表头不变
		}

		if id.Row-1 >= len(products) {
			return
		}
		product := products[id.Row-1]

		// 只更新文本列，保留图片列不变
		switch id.Col {
		case 0, 1, 2, 4:
			// 清除旧的文本控件
			var newObjects []fyne.CanvasObject
			for _, obj := range hbox.Objects {
				if _, isLabel := obj.(*widget.Label); !isLabel {
					newObjects = append(newObjects, obj)
				}
			}
			hbox.Objects = newObjects

			// 添加新的文本控件
			switch id.Col {
			case 0:
				hbox.Add(widget.NewLabel(product.ProductID))
			case 1:
				hbox.Add(widget.NewLabel(product.Name))
			case 2:
				hbox.Add(widget.NewLabel(fmt.Sprintf("¥%.2f", float64(product.MarketPrice)/100)))
			case 4:
				hbox.Add(widget.NewLabel(fmt.Sprintf("%d", product.DiscountPrice)))
			}
		}
	}

	table.Refresh()
}

// 图片加载服务
type ImageLoaderService struct {
	queue    chan *ImageLoadTask
	cache    map[string]fyne.Resource
	cacheMux sync.RWMutex
}

type ImageLoadTask struct {
	URL      string
	Callback func(fyne.Resource)
}

// 创建图片加载服务
func NewImageLoaderService(workers int) *ImageLoaderService {
	service := &ImageLoaderService{
		queue: make(chan *ImageLoadTask, 1000),
		cache: make(map[string]fyne.Resource),
	}

	// 启动工作池
	for i := 0; i < workers; i++ {
		go service.worker()
	}

	return service
}

func (s *ImageLoaderService) worker() {
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	for task := range s.queue {
		// 先检查缓存
		s.cacheMux.RLock()
		cached, exists := s.cache[task.URL]
		s.cacheMux.RUnlock()

		if exists {
			task.Callback(cached)
			continue
		}

		// 下载图片
		resp, err := client.Get(task.URL)
		if err != nil || resp.StatusCode != http.StatusOK {
			log.Printf("图片加载失败: %s, 错误: %v", task.URL, err)
			continue
		}

		// 创建资源
		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			continue
		}

		// 生成资源ID
		hash := md5.Sum([]byte(task.URL))
		resourceID := fmt.Sprintf("img_%x", hash)
		res := fyne.NewStaticResource(resourceID, data)

		// 更新缓存
		s.cacheMux.Lock()
		s.cache[task.URL] = res
		s.cacheMux.Unlock()

		// 执行回调
		task.Callback(res)
	}
}

// 添加加载任务
func (s *ImageLoaderService) LoadImage(url string, callback func(fyne.Resource)) {
	task := &ImageLoadTask{
		URL:      url,
		Callback: callback,
	}
	s.queue <- task
}

// 全局图片加载服务
var imageLoader = NewImageLoaderService(5) // 5个工作线程

// 图片显示组件
type AsyncImage struct {
	widget.BaseWidget
	url      string
	resource fyne.Resource
	image    *canvas.Image
}

func NewAsyncImage(url string) *AsyncImage {
	img := &AsyncImage{
		url:   url,
		image: canvas.NewImageFromResource(nil),
	}
	img.image.SetMinSize(fyne.NewSize(40, 40))
	img.image.FillMode = canvas.ImageFillContain

	// 设置占位符
	img.image.Resource = fyne.Theme.Icon(fyne.CurrentApp().Settings().Theme(), "question")

	img.ExtendBaseWidget(img)

	// 启动异步加载
	if url != "" {
		imageLoader.LoadImage(url, img.onImageLoaded)
	}

	return img
}

func (i *AsyncImage) onImageLoaded(res fyne.Resource) {
	// 只在资源确实加载完成时更新
	if res != nil {
		i.resource = res
		fyne.DoAndWait(func() {
			i.image.Resource = res
			i.Refresh()
		})
	}
}

func (i *AsyncImage) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(i.image)
}
