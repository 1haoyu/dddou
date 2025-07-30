package OrderModel

type Response struct {
	St   int    `json:"st"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []Data `json:"data"`
}

type Data struct {
	OrderID                  string             `json:"order_id"`
	ShopOrderID              string             `json:"shop_order_id"`
	OrderStatus              int                `json:"order_status"`
	UserID                   string             `json:"user_id"`
	NowTs                    int64              `json:"now_ts"`
	PayType                  int                `json:"pay_type"`
	OrderType                int                `json:"order_type"`
	BType                    int                `json:"b_type"`
	CBiz                     int                `json:"c_biz"`
	Biz                      int                `json:"biz"`
	ReceiveType              int                `json:"receive_type"`
	EExpress                 int                `json:"e_express"`
	Repeat                   int                `json:"repeat"`
	IsDup                    int                `json:"is_dup"`
	PreReceiveInfoExist      int                `json:"pre_receive_info_exist"`
	HasWriteOffRecord        bool               `json:"has_write_off_record"`
	IsAlreadyModifyAmount    int                `json:"is_already_modify_amount"`
	UserIsAuth               int                `json:"user_is_auth"`
	CanModifyAmount          int                `json:"can_modify_amount"`
	ChangeAddr               int                `json:"change_addr"`
	StoreName                string             `json:"store_name"`
	WaitShipCount            int                `json:"wait_ship_count"`
	ShippedCount             int                `json:"shipped_count"`
	ProductCount             int                `json:"product_count"`
	TotalPostAmount          int                `json:"total_post_amount"`
	TotalPayAmount           int                `json:"total_pay_amount"`
	PayAmount                int                `json:"pay_amount"`
	PostAmount               int                `json:"post_amount"`
	TotalTaxAmount           int                `json:"total_tax_amount"`
	TotalIncludeTaxAmount    int                `json:"total_include_tax_amount"`
	TotalExcludingTaxAmount  int                `json:"total_excluding_tax_amount"`
	TotalGoodsAmount         int                `json:"total_goods_amount"`
	PromotionAmount          int                `json:"promotion_amount"`
	ModifyAmount             int                `json:"modify_amount"`
	ModifyPostAmount         int                `json:"modify_post_amount"`
	SkuModifyAmount          int                `json:"sku_modify_amount"`
	ShopReceiveAmount        int                `json:"shop_receive_amount"`
	PromotionPayAmount       int                `json:"promotion_pay_amount"`
	EnvelopePromotionAmount  int                `json:"envelope_promotion_amount"`
	TotalTaxAmountDesc       string             `json:"total_tax_amount_desc"`
	ActualPayAmount          int                `json:"actual_pay_amount"`
	CreateTime               int64              `json:"create_time"`
	ConfirmTime              int64              `json:"confirm_time"`
	PayTime                  int64              `json:"pay_time"`
	LogisticsTime            int64              `json:"logistics_time"`
	ReceiptTime              int64              `json:"receipt_time"`
	GroupTime                int64              `json:"group_time"`
	ExpShipTime              int64              `json:"exp_ship_time"`
	OrderTypeDesc            string             `json:"order_type_desc"`
	PayTypeDesc              string             `json:"pay_type_desc"`
	WriteOffDesc             interface{}        `json:"write_off_desc"` // Could be null or string
	BuyerWords               string             `json:"buyer_words"`
	Remark                   string             `json:"remark"`
	Star                     int                `json:"star"`
	UserNickname             string             `json:"user_nickname"`
	HasWriteOff              bool               `json:"has_write_off"`
	HasMore                  bool               `json:"has_more"`
	PreSaleDesc              string             `json:"pre_sale_desc"`
	ReceiveInfo              interface{}        `json:"receive_info"` // Could be null or object
	ReceiverInfo             ReceiverInfo       `json:"receiver_info"`
	PolicyInfo               interface{}        `json:"policy_info"` // Could be null or object
	OrderStatusInfo          OrderStatusInfo    `json:"order_status_info"`
	OperationActions         []string           `json:"operation_actions"`
	ActionMap                map[string]Action  `json:"action_map"`
	Button                   interface{}        `json:"button"`            // Could be null or object
	OrderBottomCard          interface{}        `json:"order_bottom_card"` // Could be null or object
	ProductItem              []ProductItem      `json:"product_item"`
	ShopOrderTag             []Tag              `json:"shop_order_tag"`
	PayAmountDetail          []AmountDetail     `json:"pay_amount_detail"`
	WayBillURL               string             `json:"way_bill_url"`
	CrossBorderSendType      int                `json:"cross_border_send_type"`
	OrderAmountCard          interface{}        `json:"order_amount_card"` // Could be null or object
	PayAmountDesc            string             `json:"pay_amount_desc"`
	ShopReceiveAmountDesc    string             `json:"shop_receive_amount_desc"`
	SerialNumbers            interface{}        `json:"serial_numbers"` // Could be null or array
	AddressTag               []interface{}      `json:"address_tag"`    // Could be empty array
	SupportDetail            interface{}        `json:"support_detail"` // Could be null or object
	NeedSerialNumber         bool               `json:"need_serial_number"`
	BTypeDesc                string             `json:"b_type_desc"`
	CBizDesc                 string             `json:"c_biz_desc"`
	PriceDetail              interface{}        `json:"price_detail"`     // Could be null or object
	PromotionDetail          interface{}        `json:"promotion_detail"` // Could be null or object
	PayTypeDescHover         string             `json:"pay_type_desc_hover"`
	ManualOrderType          int                `json:"manual_order_type"`
	OrderIDForShow           string             `json:"order_id_for_show"`
	OrderTagStamp            []interface{}      `json:"order_tag_stamp"` // Could be empty array
	URLMap                   map[string]URLItem `json:"url_map"`
	UserProfileTag           []Tag              `json:"user_profile_tag"`
	SupermarketOrderSerialNo string             `json:"supermarket_order_serial_no"`
	DeliverName              string             `json:"deliver_name"`
	DeliverMobile            string             `json:"deliver_mobile"`
	ReceiptTimeFmt           string             `json:"receipt_time_fmt"`
	LogisticsStatus          string             `json:"logistics_status"`
	GreetWords               string             `json:"greet_words"`
	TransferReceiverInfo     interface{}        `json:"transfer_receiver_info"` // Could be null or object
	TotalProductCount        int                `json:"total_product_count"`
	TotalPrice               int                `json:"total_price"`
	LatestLogisticInfo       LogisticInfo       `json:"latest_logistic_info"`
	CreateTimeStr            string             `json:"create_time_str"`
	AmountDetailMap          interface{}        `json:"amount_detail_map"`        // Could be null or object
	ExtraTag                 interface{}        `json:"extra_tag"`                // Could be null or object
	ShopPrivilegeInfoList    []interface{}      `json:"shop_privilege_info_list"` // Could be empty array
	GiftReceiveTimeStr       string             `json:"gift_receive_time_str"`
	RelateInfos              interface{}        `json:"relate_infos"`        // Could be null or object
	BaseCard                 interface{}        `json:"base_card"`           // Could be null or object
	ReceiverCommon           interface{}        `json:"receiver_common"`     // Could be null or object
	IsOrderInABTest          interface{}        `json:"is_order_in_ab_test"` // Could be null or object
}

type ReceiverInfo struct {
	PostReceiver      string        `json:"post_receiver"`
	PostTel           string        `json:"post_tel"`
	PostAddr          Address       `json:"post_addr"`
	CanView           int           `json:"can_view"`
	PostTelType       int           `json:"post_tel_type"`
	ExpireTime        int64         `json:"expire_time"`
	IsShowEditAddress bool          `json:"is_show_edit_address"`
	CanPostpone       bool          `json:"can_postpone"`
	ExtensionNumber   string        `json:"extension_number"`
	PostTelMask       string        `json:"post_tel_mask"`
	AddressTag        []interface{} `json:"address_tag"`
	UserAccountInfos  interface{}   `json:"user_account_infos"`
	UIType            string        `json:"ui_type"`
	BuyerTelInfo      BuyerTelInfo  `json:"buyer_tel_info"`
	ExtraMap          interface{}   `json:"extra_map"`
}

type Address struct {
	Province AddrNode `json:"province"`
	City     AddrNode `json:"city"`
	Town     AddrNode `json:"town"`
	Street   AddrNode `json:"street"`
	Detail   string   `json:"detail"`
}

type AddrNode struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	HasChild bool   `json:"has_child"`
}

type BuyerTelInfo struct {
	PayTel          string `json:"pay_tel"`
	PayTelType      int    `json:"pay_tel_type"`
	ExpireTime      int64  `json:"expire_time"`
	CanPostpone     bool   `json:"can_postpone"`
	PayTelMask      string `json:"pay_tel_mask"`
	UIType          string `json:"ui_type"`
	CanView         int    `json:"can_view"`
	ExtensionNumber string `json:"extension_number"`
	NickName        string `json:"nick_name"`
}

type OrderStatusInfo struct {
	DeadLineTime          int64  `json:"dead_line_time"`
	OrderStatusIcon       int    `json:"order_status_icon"`
	OrderStatusText       string `json:"order_status_text"`
	OrderStatusRemark     string `json:"order_status_remark"`
	OrderStatusOverRemark string `json:"order_status_over_remark"`
	ShowRule              bool   `json:"show_rule"`
	// ... (其他字段根据实际需要添加)
}

type Action struct {
	BizAction        string        `json:"biz_action"`
	BizActionDesc    string        `json:"biz_action_desc"`
	BizActionOptions ActionOptions `json:"biz_action_options"`
	BizActionName    string        `json:"biz_action_name"`
	HasActionRecord  bool          `json:"has_action_record"`
	BizActionUI      string        `json:"biz_action_ui"`
	BizActionHover   string        `json:"biz_action_hover"`
	BizActionURL     string        `json:"biz_action_url"`
	BizActionType    string        `json:"biz_action_type"`
}

type ActionOptions struct {
	ModalTitle        string      `json:"modal_title"`
	ModalContent      string      `json:"modal_content"`
	OkText            string      `json:"ok_text"`
	CancelText        string      `json:"cancel_text"`
	OkURL             string      `json:"ok_url"`
	APIParams         interface{} `json:"api_params"`
	APIMethod         string      `json:"api_method"`
	BizActionAPI      string      `json:"biz_action_api"`
	ActionOption1Text string      `json:"action_option1_text"`
	ActionOption1URL  string      `json:"action_option1_url"`
}

type ProductItem struct {
	ItemOrderID          string                 `json:"item_order_id"`
	OrderStatus          int                    `json:"order_status"`
	OrderType            string                 `json:"order_type"`
	PreSaleType          int                    `json:"pre_sale_type"`
	ProcessType          int                    `json:"process_type"`
	Biz                  int                    `json:"biz"`
	PriceHasTaxType      string                 `json:"price_has_tax_type"`
	CBiz                 int                    `json:"c_biz"`
	TotalAmount          int                    `json:"total_amount"`
	PostAmount           int                    `json:"post_amount"`
	TaxAmount            int                    `json:"tax_amount"`
	PayAmount            int                    `json:"pay_amount"`
	ProductID            string                 `json:"product_id"`
	SkuID                int                    `json:"sku_id"`
	MerchantSkuCode      string                 `json:"merchant_sku_code"`
	ProductName          string                 `json:"product_name"`
	Img                  string                 `json:"img"`
	SkuSpec              []SkuSpec              `json:"sku_spec"`
	ComboAmount          int                    `json:"combo_amount"`
	ComboNum             int                    `json:"combo_num"`
	ShippedCount         int                    `json:"shipped_count"`
	SkuWeight            int                    `json:"sku_weight"`
	ItemOrderStatusDesc  string                 `json:"item_order_status_desc"`
	PackageStatusDesc    string                 `json:"package_status_desc"`
	AftersaleService     []string               `json:"aftersale_service"`
	PolicyInfo           interface{}            `json:"policy_info"` // Could be null
	AfterSaleInfo        AfterSaleInfo          `json:"after_sale_info"`
	Tags                 []Tag                  `json:"tags"`
	PromotionDetail      interface{}            `json:"promotion_detail"`    // Could be null
	SkuBundleSubSkus     interface{}            `json:"sku_bundle_sub_skus"` // Could be null
	ItemWarehouseID      string                 `json:"item_warehouse_id"`
	ItemWarehouseName    string                 `json:"item_warehouse_name"`
	ActionMap            map[string]interface{} `json:"action_map"`     // Could be empty
	RelationOrder        interface{}            `json:"relation_order"` // Could be null
	TradeType            int                    `json:"trade_type"`
	AvailableFetchCnt    int                    `json:"available_fetch_cnt"`
	Properties           []Property             `json:"properties"`
	SkuCustomizationInfo interface{}            `json:"sku_customization_info"` // Could be null
	GivenProductType     string                 `json:"given_product_type"`
	// ... (其他字段根据实际需要添加)
}

type SkuSpec struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AfterSaleInfo struct {
	URLAfterSaleNo                int    `json:"url_after_sale_no"`
	AfterSaleText                 string `json:"after_sale_text"`
	AfterSaleRemark               string `json:"after_sale_remark"`
	AfterSaleID                   string `json:"after_sale_id"`
	HasPreSale                    bool   `json:"has_pre_sale"`
	AftersaleStatusClassColorType int    `json:"aftersale_status_class_color_type"`
	AftersaleStatusClassString    string `json:"aftersale_status_class_string"`
	AfterSaleStatusRemark         string `json:"after_sale_status_remark"`
	AfterSaleStatusRemarkHover    string `json:"after_sale_status_remark_hover"`
	AfterSaleURL                  string `json:"after_sale_url,omitempty"`
}

type Tag struct {
	Key               string      `json:"key"`
	Text              string      `json:"text"`
	HoverText         string      `json:"hover_text"`
	TagType           string      `json:"tag_type"`
	HelpDoc           string      `json:"help_doc"`
	Icon              string      `json:"icon"`
	BelongTo          string      `json:"belong_to"`
	QuestionHoverText string      `json:"question_hover_text"`
	Position          string      `json:"position"`
	URLConfigMaps     interface{} `json:"url_config_maps"` // Could be null
}

type Property struct {
	Key        string `json:"key"`
	Text       string `json:"text"`
	Action     string `json:"action"`
	ActionText string `json:"action_text"`
}

type AmountDetail struct {
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	Hover     string `json:"hover"`
	AmountInt int    `json:"amount_int"`
}

type URLItem struct {
	Text           string         `json:"text"`
	URL            string         `json:"url"`
	Type           string         `json:"type"`
	ComponentProps ComponentProps `json:"component_props"`
	RichHover      RichHover      `json:"rich_hover"`
	HoverProps     HoverProps     `json:"hover_props"`
}

type ComponentProps struct {
	Style interface{} `json:"style"` // Could be map or null
}

type RichHover struct {
	Text        string        `json:"text"`
	Placeholder []interface{} `json:"placeholder"` // Could be array of objects
	ModalTitle  string        `json:"modal_title"`
}

type HoverProps struct {
	OverlayInnerStyle interface{} `json:"overlayInnerStyle"` // Could be null
}

type LogisticInfo struct {
	TrackingNo             string      `json:"tracking_no"`
	CompanyCode            string      `json:"company_code"`
	Track                  string      `json:"track"`
	DeliverName            string      `json:"deliver_name"`
	DeliverMobile          string      `json:"deliver_mobile"`
	LoStatusText           string      `json:"lo_status_text"`
	Tags                   interface{} `json:"tags"` // Could be null
	TrackDetailDescription string      `json:"track_detail_description"`
	ShowTrack              bool        `json:"show_track"`
	LoStatusTextType       string      `json:"lo_status_text_type"`
	DescriptionTextType    string      `json:"description_text_type"`
	LoStatusTextV2         string      `json:"lo_status_text_v2"`
	ShowCallDispatch       bool        `json:"show_call_dispatch"`
}
