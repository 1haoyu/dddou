package SkuModel

type Response struct {
	St    int        `json:"st"`
	Msg   string     `json:"msg"`
	Code  int        `json:"code"`
	Data  []DataItem `json:"data"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
	Total int        `json:"total"`
}

type DataItem struct {
	Alerts                              interface{}             `orm:"ignore" json:"alerts"`
	ProductID                           string                  `gorose:"product_id" json:"product_id"`
	ShopID                              int                     `gorose:"shop_id" json:"shop_id"`
	Name                                string                  `gorose:"name" json:"name"`
	ProductURL                          string                  `orm:"ignore" json:"product_url"`
	ProductURLForCopy                   string                  `orm:"ignore" json:"product_url_for_copy"`
	Description                         string                  `orm:"ignore" json:"description"`
	Img                                 string                  `gorose:"img" json:"img"`
	CollectNum                          int                     `orm:"ignore" json:"collect_num"`
	SellNum                             int                     `orm:"ignore" json:"sell_num"`
	MarketPrice                         int                     `gorose:"market_price" json:"market_price"`
	DiscountPrice                       int                     `gorose:"discount_price" json:"discount_price"`
	PriceLower                          int                     `gorose:"price_lower" json:"price_lower"`
	PriceHigher                         int                     `gorose:"price_higher" json:"price_higher"`
	CreateTime                          string                  `orm:"ignore" json:"create_time"`
	EndTime                             int                     `orm:"ignore" json:"end_time"`
	DescDetail                          string                  `orm:"ignore" json:"desc_detail"`
	Status                              int                     `orm:"ignore" json:"status"`
	BuyDesc                             string                  `orm:"ignore" json:"buy_desc"`
	OrderDesc                           string                  `orm:"ignore" json:"order_desc"`
	SpecID                              int64                   `orm:"ignore" json:"spec_id"`
	UpdateTime                          string                  `orm:"ignore" json:"update_time"`
	CheckStatus                         int                     `orm:"ignore" json:"check_status"`
	DraftStatus                         int                     `orm:"ignore" json:"draft_status"`
	Reason                              string                  `orm:"ignore" json:"reason"`
	ComboDesc                           string                  `orm:"ignore" json:"combo_desc"`
	Checker                             string                  `orm:"ignore" json:"checker"`
	IsShow                              int                     `orm:"ignore" json:"is_show"`
	Category                            string                  `orm:"ignore" json:"category"`
	Mobile                              string                  `orm:"ignore" json:"mobile"`
	BuyRecordSwitch                     int                     `orm:"ignore" json:"buy_record_switch"`
	ImgStyleSwitch                      int                     `orm:"ignore" json:"img_style_switch"`
	HotStyle                            int                     `orm:"ignore" json:"hot_style"`
	CosRatio                            int                     `orm:"ignore" json:"cos_ratio"`
	BizType                             int                     `orm:"ignore" json:"biz_type"`
	Usp                                 string                  `orm:"ignore" json:"usp"`
	FirstCid                            int                     `orm:"ignore" json:"first_cid"`
	SecondCid                           int                     `orm:"ignore" json:"second_cid"`
	ThirdCid                            int                     `orm:"ignore" json:"third_cid"`
	PayType                             int                     `orm:"ignore" json:"pay_type"`
	ThirdURL                            string                  `orm:"ignore" json:"third_url"`
	RecommendRemark                     string                  `orm:"ignore" json:"recommend_remark"`
	GroupID                             string                  `orm:"ignore" json:"group_id"`
	PutStatus                           int                     `orm:"ignore" json:"put_status"`
	NeedCode                            int                     `orm:"ignore" json:"need_code"`
	NeedShare                           int                     `orm:"ignore" json:"need_share"`
	Extra                               string                  `orm:"ignore" json:"extra"`
	OutProductID                        string                  `orm:"ignore" json:"out_product_id"`
	ShowPlace                           int                     `orm:"ignore" json:"show_place"`
	BrandID                             int                     `orm:"ignore" json:"brand_id"`
	SupplyStatus                        int                     `orm:"ignore" json:"supply_status"`
	ProductType                         int                     `orm:"ignore" json:"product_type"`
	DistrSellNum                        int                     `orm:"ignore" json:"distr_sell_num"`
	PresellType                         int                     `orm:"ignore" json:"presell_type"`
	Operations                          []Operation             `orm:"ignore" json:"operations"`
	DeliveryDelayDay                    int                     `orm:"ignore" json:"delivery_delay_day"`
	AuditSendTime                       int                     `orm:"ignore" json:"audit_send_time"`
	StockNum                            int                     `orm:"ignore" json:"stock_num"`
	IsSeckill                           int                     `orm:"ignore" json:"is_seckill"`
	InChannel                           int                     `orm:"ignore" json:"in_channel"`
	CommentNum                          int                     `orm:"ignore" json:"comment_num"`
	CommentPercent                      int                     `orm:"ignore" json:"comment_percent"`
	CommentGood                         int                     `orm:"ignore" json:"comment_good"`
	Tab                                 string                  `orm:"ignore" json:"tab"`
	AppealDetail                        interface{}             `orm:"ignore" json:"appeal_detail"`
	DraftTime                           int64                   `orm:"ignore" json:"draft_time"`
	PublishTime                         int64                   `orm:"ignore" json:"publish_time"`
	AuditTime                           int64                   `orm:"ignore" json:"audit_time"`
	AuditReason                         *AuditReason            `orm:"ignore" json:"audit_reason"`
	PublishID                           int64                   `orm:"ignore" json:"publish_id"`
	PublishIDStr                        string                  `orm:"ignore" json:"publish_id_str"`
	Highlight                           interface{}             `orm:"ignore" json:"highlight"`
	OfflineTime                         int64                   `orm:"ignore" json:"offline_time"`
	DeleteTime                          int                     `orm:"ignore" json:"delete_time"`
	EditSpecFlag                        int                     `orm:"ignore" json:"edit_spec_flag"`
	Promotions                          []interface{}           `orm:"ignore" json:"promotions"`
	PromotionStockNum                   int                     `orm:"ignore" json:"promotion_stock_num"`
	LockProperty                        interface{}             `orm:"ignore" json:"lock_property"`
	WithAd                              bool                    `orm:"ignore" json:"with_ad"`
	CategoryLeafID                      int                     `orm:"ignore" json:"category_leaf_id"`
	BizKind                             int                     `orm:"ignore" json:"biz_kind"`
	SourceCountryID                     int                     `orm:"ignore" json:"source_country_id"`
	OriginCountryID                     int                     `orm:"ignore" json:"origin_country_id"`
	BrandCountryID                      int                     `orm:"ignore" json:"brand_country_id"`
	LockStockNum                        int                     `orm:"ignore" json:"lock_stock_num"`
	PresellDelay                        int                     `orm:"ignore" json:"presell_delay"`
	PresellEndTime                      int                     `orm:"ignore" json:"presell_end_time"`
	OfflineType                         int                     `orm:"ignore" json:"offline_type"`
	ReturnPolicyID                      string                  `orm:"ignore" json:"return_policy_id"`
	IsAuctionProduct                    bool                    `orm:"ignore" json:"is_auction_product"`
	IsMassAuctionProduct                bool                    `orm:"ignore" json:"is_mass_auction_product"`
	IsPrize                             bool                    `orm:"ignore" json:"is_prize"`
	IsNotForSale                        bool                    `orm:"ignore" json:"is_not_for_sale"`
	NewStepProduct                      bool                    `orm:"ignore" json:"new_step_product"`
	Tags                                []interface{}           `orm:"ignore" json:"tags"`
	NotShowTags                         []interface{}           `orm:"ignore" json:"not_show_tags"`
	SellChannel                         interface{}             `orm:"ignore" json:"sell_channel"`
	DraftVideo                          *DraftVideo             `orm:"ignore" json:"draft_video"`
	IsDepositFindGoods                  bool                    `orm:"ignore" json:"is_deposit_find_goods"`
	InfoInMainPackage                   bool                    `orm:"ignore" json:"info_in_main_package"`
	IsPackageProduct                    int                     `orm:"ignore" json:"is_package_product"`
	PackageProductList                  []interface{}           `orm:"ignore" json:"package_product_list"`
	SubProductList                      []interface{}           `orm:"ignore" json:"sub_product_list"`
	AreaStockButtonCase                 int                     `orm:"ignore" json:"area_stock_button_case"`
	ProductIndicators                   interface{}             `orm:"ignore" json:"product_indicators"`
	InPromotion                         int                     `orm:"ignore" json:"in_promotion"`
	SelfSellStockNum                    int                     `orm:"ignore" json:"self_sell_stock_num"`
	ChannelStockNum                     int                     `orm:"ignore" json:"channel_stock_num"`
	ChannelProductNum                   int                     `orm:"ignore" json:"channel_product_num"`
	ChannelProductNumWithCond           int                     `orm:"ignore" json:"channel_product_num_with_cond"`
	DraftMaterial                       *DraftMaterial          `orm:"ignore" json:"draft_material"`
	HasChannel                          int                     `orm:"ignore" json:"has_channel"`
	PresellConfigLevel                  interface{}             `orm:"ignore" json:"presell_config_level"`
	PresellDeliveryType                 interface{}             `orm:"ignore" json:"presell_delivery_type"`
	ProductMaterialMap                  map[string]MediaItem    `gorose:"product_material_map" json:"product_material_map"`
	CanCombine                          interface{}             `orm:"ignore" json:"can_combine"`
	CanNotCombineReason                 interface{}             `orm:"ignore" json:"can_not_combine_reason"`
	AppealInfo                          interface{}             `orm:"ignore" json:"appeal_info"`
	HasBaseScore                        bool                    `orm:"ignore" json:"has_base_score"`
	BaseScore                           int                     `orm:"ignore" json:"base_score"`
	Pics                                []string                `gorose:"pics" json:"pics"`
	StandardBrandID                     int64                   `orm:"ignore" json:"standard_brand_id"`
	BrandName                           string                  `orm:"ignore" json:"brand_name"`
	MpuInfoCheckStatus                  int                     `orm:"ignore" json:"mpu_info_check_status"`
	CategoryDetail                      *CategoryDetail         `orm:"ignore" json:"category_detail"`
	QualityScore                        int                     `orm:"ignore" json:"quality_score"`
	QualityLevel                        string                  `orm:"ignore" json:"quality_level"`
	HasQualityScore                     bool                    `orm:"ignore" json:"has_quality_score"`
	IsDiagnoseExemption                 bool                    `orm:"ignore" json:"is_diagnose_exemption"`
	NamePrefix                          string                  `orm:"ignore" json:"name_prefix"`
	NameSuffix                          string                  `orm:"ignore" json:"name_suffix"`
	ProductFormatNew                    map[string][]FormatItem `orm:"ignore" json:"product_format_new"`
	CanSeries                           interface{}             `orm:"ignore" json:"can_series"`
	CanNotSeriesReason                  interface{}             `orm:"ignore" json:"can_not_series_reason"`
	StoreInfo                           interface{}             `orm:"ignore" json:"store_info"`
	OperatorInfo                        interface{}             `orm:"ignore" json:"operator_info"`
	PrizeActivityStatus                 interface{}             `orm:"ignore" json:"prize_activity_status"`
	PassWuyouRiskCheck                  interface{}             `orm:"ignore" json:"pass_wuyou_risk_check"`
	Skus                                interface{}             `orm:"ignore" json:"skus"`
	IsWuyou                             interface{}             `orm:"ignore" json:"is_wuyou"`
	TitleLimit                          *TitleLimit             `orm:"ignore" json:"title_limit"`
	IsProductTalentEvaluationClose      bool                    `orm:"ignore" json:"is_product_talent_evaluation_close"`
	CanPackage                          interface{}             `orm:"ignore" json:"can_package"`
	CanNotPackageReason                 interface{}             `orm:"ignore" json:"can_not_package_reason"`
	SkuPriceReadOnly                    int                     `orm:"ignore" json:"sku_price_read_only"`
	IsDcp                               bool                    `orm:"ignore" json:"is_dcp"`
	IsProductMainVideoPreGeneratedClose bool                    `orm:"ignore" json:"is_product_main_video_pre_generated_close"`
	Identity                            *Identity               `orm:"ignore" json:"identity"`
	CategoryName                        string                  `orm:"ignore" json:"category_name"`
	StockNumReadOnly                    int                     `orm:"ignore" json:"stock_num_read_only"`
	CanAigcProduce                      bool                    `orm:"ignore" json:"can_aigc_produce"`
	AigcProduceAbility                  interface{}             `orm:"ignore" json:"aigc_produce_ability"`
	ShopCategoryPath                    interface{}             `orm:"ignore" json:"shop_category_path"`
	PricePowerLevelInfo                 interface{}             `orm:"ignore" json:"price_power_level_info"`
	PriceLevelStandardInfo              interface{}             `orm:"ignore" json:"price_level_standard_info"`
	IsAggProduct                        bool                    `orm:"ignore" json:"is_agg_product"`
	MarketAdvertisementExposure         *MarketAdvertisement    `orm:"ignore" json:"market_advertisement_exposure"`
	SellNum30d                          int                     `orm:"ignore" json:"sell_num_30d"`
	SyncStoreNum                        interface{}             `orm:"ignore" json:"sync_store_num"`
	PreProductSizeTemplateInfo          *SizeTemplateInfo       `orm:"ignore" json:"pre_product_size_template_info"`
	FoodPropertyOptimize                *FoodProperty           `orm:"ignore" json:"food_property_optimize"`
	PredCategoryDetail                  interface{}             `orm:"ignore" json:"pred_category_detail"`
	TextCountEmitCodes                  interface{}             `orm:"ignore" json:"text_count_emit_codes"`
	PunishTicket                        *PunishTicket           `orm:"ignore" json:"punish_ticket"`
	ProductMaterialCnt                  *MaterialCount          `orm:"ignore" json:"product_material_cnt"`
	ProductMaterialMissCnt              *MaterialMissCount      `orm:"ignore" json:"product_material_miss_cnt"`
	TradeLimitRule                      *TradeLimitRule         `orm:"ignore" json:"trade_limit_rule"`
	EanPublishedEditDisabled            bool                    `orm:"ignore" json:"ean_published_edit_disabled"`
	NeedShowCoupon                      interface{}             `orm:"ignore" json:"need_show_coupon"`
	PickUpMethod                        interface{}             `orm:"ignore" json:"PickUpMethod"`
	RectifyCanRecover                   interface{}             `orm:"ignore" json:"rectify_can_recover"`
	DiagnoseActivityInfo                interface{}             `orm:"ignore" json:"diagnose_activity_info"`
	DiagnoseActivityInfoData            interface{}             `orm:"ignore" json:"diagnose_activity_info_data"`
	RectifyInfo                         interface{}             `orm:"ignore" json:"rectify_info"`
	SoldoutUpdateTime                   int                     `orm:"ignore" json:"soldout_update_time"`
	StockTips                           interface{}             `orm:"ignore" json:"stock_tips"`
	PayNoStockSkus                      interface{}             `orm:"ignore" json:"pay_no_stock_skus"`
	HaveAuditRejectSuggest              interface{}             `orm:"ignore" json:"have_audit_reject_suggest"`
	AuditRejectSuggestion               interface{}             `orm:"ignore" json:"audit_reject_suggestion"`
	AuditReasonDetailV2                 []interface{}           `orm:"ignore" json:"audit_reason_detail_v2"`
	MaterialOpRestrict                  interface{}             `orm:"ignore" json:"material_op_restrict"`
	CorrectionTicketID                  interface{}             `orm:"ignore" json:"correction_ticket_id"`
	MainPics34                          interface{}             `orm:"ignore" json:"main_pics_34"`
	CanBatchUpdateStock                 interface{}             `orm:"ignore" json:"can_batch_update_stock"`
	CantBatchUpdateStockReason          interface{}             `orm:"ignore" json:"cant_batch_update_stock_reason"`
	StatusDel                           int                     `gorose:"status_del" json:"status_del"`
}

type Operation struct {
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Status      string      `json:"status"`
	Reason      string      `json:"reason"`
	AppLink     string      `json:"app_link"`
	AlertDialog AlertDialog `json:"alert_dialog"`
	JumpURL     string      `json:"jump_url"`
	Hover       *Hover      `json:"hover"`
}

type AlertDialog struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	SwitchDialog bool   `json:"switch_dialog"`
}

type Hover struct {
	Content string `json:"content"`
}

type AuditReason struct {
	Name            []string `json:"name"`
	RecommendRemark []string `json:"recommend_remark"`
	DiscountPrice   []string `json:"discount_price"`
	ProductFormat   []string `json:"product_format"`
	QualityReport   []string `json:"quality_report"`
	ClassQuality    []string `json:"class_quality"`
	Common          []string `json:"common"`
	ProductPic      []string `json:"product_pic"`
	ProductPicNum   []string `json:"product_pic_num"`
	SkuPic          []string `json:"sku_pic"`
	SkuPicNum       []string `json:"sku_pic_num"`
	SkuName         []string `json:"sku_name"`
	DescriptionPic  []string `json:"description_pic"`
	DescriptionText []string `json:"description_text"`
	ProductVideo    []string `json:"product_video"`
	Brand           []string `json:"brand"`
	ProductMaterial []string `json:"product_material"`
	DeliveryMode    []string `json:"delivery_mode"`
	SkuPrice        []string `json:"sku_price"`
}

type DraftVideo struct {
	Vid       string `json:"vid"`
	Status    int    `json:"status"`
	Reason    string `json:"reason"`
	AuditTime int64  `json:"audit_time"`
}

type DraftMaterial struct {
	WhiteBackgroundImg *MaterialStatus `json:"whiteBackgroundImg"`
}

type MaterialStatus struct {
	URL    []string `json:"url"`
	Status int      `json:"status"`
	Reason string   `json:"reason"`
}

type MediaItem struct {
	MediaType   int         `json:"media_type"`
	URLs        []string    `json:"urls"`
	ResourceIDs []string    `json:"resource_ids"`
	Status      int         `json:"status"`
	Reason      string      `json:"reason"`
	MaterialID  string      `json:"material_id"`
	Approval    bool        `json:"approval"`
	ActionList  interface{} `json:"action_list"`
	Extra       string      `json:"extra"`
}

type CategoryDetail struct {
	FirstCid    int    `json:"first_cid"`
	SecondCid   int    `json:"second_cid"`
	ThirdCid    int    `json:"third_cid"`
	FourthCid   int    `json:"fourth_cid"`
	FirstCname  string `json:"first_cname"`
	SecondCname string `json:"second_cname"`
	ThirdCname  string `json:"third_cname"`
	FourthCname string `json:"fourth_cname"`
	Status      int    `json:"status"`
}

type FormatItem struct {
	Value        int         `json:"value"`
	Name         string      `json:"name"`
	DiyType      int         `json:"diy_type"`
	Tags         interface{} `json:"tags"`
	PropertyName string      `json:"PropertyName"`
}

type TitleLimit struct {
	MaxLength int `json:"max_length"`
	MinLength int `json:"min_length"`
}

type Identity struct {
	SellerType     string `json:"seller_type"`
	VerticalMarket string `json:"vertical_market"`
	Source         string `json:"source"`
}

type MarketAdvertisement struct {
	EnrollableBigPromotionInfo  interface{}    `json:"enrollable_big_promotion_info"`
	EnrollablePromotionToolInfo interface{}    `json:"enrollable_promotion_tool_info"`
	EnrollableResources         []ResourceItem `json:"enrollable_resources"`
	HitAb                       int            `json:"hit_ab"`
}

type ResourceItem struct {
	ResourceKey      string      `json:"resource_key"`
	ResourceName     string      `json:"resource_name"`
	ProfitEffectInfo string      `json:"profit_effect_info"`
	OperationLink    string      `json:"operation_link"`
	Extra            interface{} `json:"extra"`
}

type SizeTemplateInfo struct {
	SizeDescImageURL     string `json:"size_desc_image_url"`
	PreGenSizeImageURL   string `json:"pre_gen_size_image_url"`
	PreGenSizeTemplateID string `json:"pre_gen_size_template_id"`
	OptimizeText         string `json:"optimize_text"`
	OptimizeType         string `json:"optimize_type"`
}

type FoodProperty struct {
	ImageURL     interface{} `json:"image_url"`
	Ingredients  string      `json:"ingredients"`
	Nutrition    interface{} `json:"nutrition"`
	SpuInfo      interface{} `json:"spu_info"`
	OptimizeType string      `json:"optimize_type"`
}

type PunishTicket struct {
	HasPunishTicket bool   `json:"has_punish_ticket"`
	PunishTicketURL string `json:"punish_ticket_url"`
}

type MaterialCount struct {
	Pic11Cnt     int `json:"pic_11_cnt"`
	Pic34Cnt     int `json:"pic_34_cnt"`
	MainVideoCnt int `json:"main_video_cnt"`
	WhiteBgCnt   int `json:"white_bg_cnt"`
	LongPicCnt   int `json:"long_pic_cnt"`
}

type MaterialMissCount struct {
	MissPic11Cnt     int `json:"miss_pic_11_cnt"`
	MissPic34Cnt     int `json:"miss_pic_34_cnt"`
	MissMainVideoCnt int `json:"miss_main_video_cnt"`
	MissWhiteBgCnt   int `json:"miss_white_bg_cnt"`
}

type TradeLimitRule struct {
	LimitRuleItems []interface{} `json:"limit_rule_items"`
	LimitRuleCount int           `json:"limit_rule_count"`
}
