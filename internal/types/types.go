package types

import (
	"context"
	"sync"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/doraemonkeys/paddleocr"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/playwright-community/playwright-go"
	"main.go/dataModel/CookieModel"
	"main.go/dataModel/ShopAccount"
	"main.go/dataModel/ShopModel"
	"main.go/dataModel/SkuModel"
	"main.go/dataModel/UserModel"
)

type AppState struct {
	MyApp         fyne.App
	MyWindow      fyne.Window
	StatusLabel   *widget.Label
	LogOutput     *widget.Entry
	ImageGrid     *fyne.Container
	DownloadQueue chan struct {
		URL   string
		Index int
	}
	OCRQueue     chan OCRJob
	ImageMutex   sync.Mutex
	ImageObjects []*canvas.Image
	OCREngine    *paddleocr.Ppocr
	Context      context.Context
	CancelFunc   context.CancelFunc
	DownloadWG   sync.WaitGroup
	OCRWG        sync.WaitGroup

	CurrentUser      UserModel.UserInfo
	Shops            []ShopModel.Account
	ProductTabs      *container.AppTabs
	StatusBar        *widget.Label
	ShopListBinding  binding.UntypedList
	LoginForm        *widget.Form
	LeftPanel        *fyne.Container
	FilterFilePath   string
	FilterKeywords   []string
	ShopListPanel    *fyne.Container
	FilterPanel      *fyne.Container
	KeywordCount     *widget.Label
	TabShopMap       map[string]ShopModel.Account
	SplitContainer   *container.Split
	TopPanel         *fyne.Container
	ContentPanel     *fyne.Container
	NeedsRefresh     bool
	LastRefreshTime  time.Time
	OnReactivate     func(filtered []string)
	PaginationStates map[string]*PaginationState
	Playwright       *PlaywrightService // Playwright服务
	LogEntry         *widget.Entry      // URL输入框
	ShopAccountChan  ShopAccount.Data
	ShopCookieInfo   CookieModel.CookieInfo
}

type OCRJob struct {
	URL       string
	Index     int
	ImagePath string
}

type ProductOcr struct {
	URLS      []string
	ProductID string
}

// PlaywrightService 管理Playwright实例
type PlaywrightService struct {
	PW      *playwright.Playwright
	Browser playwright.Browser
	Context playwright.BrowserContext
	Page    playwright.Page
}

// 新增分页状态结构体
type PaginationState struct {
	CurrentPage   int
	PageSize      int
	TotalPages    int
	TotalProducts int
	Products      []SkuModel.DataItem
	PageInfo      binding.String // 添加绑定数据 - 修复崩溃问题
}

func NewAppState(pwService *PlaywrightService) *AppState {
	state := &AppState{
		DownloadQueue: make(chan struct {
			URL   string
			Index int
		}, 100),
		OCRQueue:         make(chan OCRJob, 100),
		ImageObjects:     make([]*canvas.Image, 0, 50), // 初始化切片避免nil
		TabShopMap:       make(map[string]ShopModel.Account),
		LastRefreshTime:  time.Now(),
		PaginationStates: make(map[string]*PaginationState),
		Playwright:       pwService, // 注入Playwright服务
	}

	// 初始化OCR引擎
	isEnable := true
	ocrEngine, err := paddleocr.NewPpocr(
		"./ocrexe/PaddleOCR-json.exe",
		paddleocr.OcrArgs{EnableMkldnn: &isEnable},
	)
	if err == nil {
		state.OCREngine = ocrEngine
	}
	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	state.Context = ctx
	state.CancelFunc = cancel

	return state
}
