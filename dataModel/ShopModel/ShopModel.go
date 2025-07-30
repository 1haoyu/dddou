package ShopModel

type Response struct {
	St    int    `json:"st"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Data  Data   `json:"data"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
}

type Data struct {
	LoginSubjectList []Account   `json:"login_subject_list"`
	PageConfig       PageConfig  `json:"page_config"`
	ZeroAction       interface{} `json:"zero_action"` // 根据JSON设置为null，使用空接口
}

type Account struct {
	SubjectID            string              `gorose:"subject_id" json:"subject_id"`
	AccountName          string              `gorose:"account_name" json:"account_name"`
	AccountAvatar        string              `gorose:"account_avatar" json:"account_avatar"`
	AccountType          AccountType         `orm:"ignore" json:"account_type"`
	IdentityType         int                 `gorose:"identity_type" json:"identity_type"`
	IdentityTypeDesc     string              `gorose:"identity_type_desc" json:"identity_type_desc"`
	BusMemberID          string              `gorose:"bus_member_id" json:"bus_member_id"`
	MemberID             string              `gorose:"member_id" json:"member_id"`
	AccountID            string              `gorose:"account_id" json:"account_id"`
	CanLogin             bool                `gorose:"can_login" json:"can_login"`
	EncodeShopID         string              `gorose:"encode_shop_id" json:"encode_shop_id"`
	LeftUpperCornerLabel []Label             `orm:"ignore" json:"left_upper_corner_label"`
	MiddleLabel          []Label             `orm:"ignore" json:"middle_label"`
	BottomLabel          []Label             `orm:"ignore" json:"bottom_label"`
	Button               Button              `orm:"ignore" json:"button"`
	SecurityCheckResult  SecurityCheckResult `orm:"ignore" json:"security_check_result"`
	LoginName            string              `gorose:"login_name" json:"login_name"` //账号
}

type AccountType struct {
	SysID            int    `json:"sys_id"`
	BusType          int    `json:"bus_type"`
	BusChildType     int    `json:"bus_child_type"`
	BusChildTypeDesc string `json:"bus_child_type_desc"`
}

type Label struct {
	Text  string `json:"text"`
	Hover string `json:"hover"`
	Show  bool   `json:"show"`
	Style Style  `json:"style"`
}

type Style struct {
	Type    int         `json:"type"`
	Default StyleDetail `json:"default"`
}

type StyleDetail struct {
	FontColor   string `json:"font_color"`
	BorderColor string `json:"border_color"`
	BgColor     string `json:"bg_color"`
}

type Button struct {
	Text   string       `json:"text"`
	Show   bool         `json:"show"`
	Type   int          `json:"type"`
	Style  ButtonStyle  `json:"style"`
	Action ButtonAction `json:"action"`
}

type ButtonStyle struct {
	Type int `json:"type"`
}

type ButtonAction struct {
	Type        int  `json:"type"`
	AutoExecute bool `json:"auto_execute"`
}

type SecurityCheckResult struct {
	Decision       int         `json:"decision"`
	VerifyType     int         `json:"verify_type"`
	FaceVerifyInfo interface{} `json:"face_verify_info"` // 根据JSON设置为空对象
}

type PageConfig struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
}
