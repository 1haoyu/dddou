package UserModel

// 用户信息结构体
type UserInfo struct {
	Id          int64   `gorose:"id" json:"id"`                          // 用户ID
	LoginName   string  `gorose:"login_name" json:"login_name"`          //账号
	LoginPass   string  `gorose:"login_pass" json:"login_pass"`          //记录的登录密码
	AccountName string  `gorose:"account_name" json:"account_name"`      //账号昵称
	Cookie      string  `gorm:"type:text" gorose:"cookie" json:"cookie"` // 新增Cookie字段
	LoginType   int     `gorose:"login_type" json:"login_type"`          //登陆类型，1为手机号，3为邮箱登录，其他类型暂时不支持
	ShopCount   int     `gorose:"shop_count" json:"shop_count"`          // 店铺数量
	Expires     float64 `gorose:"expires" json:"expires"`
	CreatedAt   string  `orm:"ignore" json:"created_at"` // 创建时间
}
