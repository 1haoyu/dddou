package ShopAccount

// {
// 	"st": 0,
// 	"msg": "success",
// 	"code": 0,
// 	"data": {
// 		"account_id": "235523422",
// 		"account_name": "亦亦女鞋店",
// 		"account_avatar": "https://p6-ecom-qualification-sign.ecombdimg.com/tos-cn-i-6vegkygxbk/34ba9b3bc8d249b08c6d588bccf83ce7~tplv-6vegkygxbk-s:750.image?lk3s=c08c0450\u0026x-expires=1784641900\u0026x-signature=HjjeiNJL5NwouuCyeTA0fQ%2F345o%3D",
// 		"is_child": true,
// 		"bus_type": 1,
// 		"uid": "2921835002084970",
// 		"auth_code": ""
// 	},
// 	"page": 0,
// 	"size": 0,
// 	"total": 0
// }

type Root struct {
	St    int    `json:"st"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Data  Data   `json:"data"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int    `json:"total"`
}

type Data struct {
	Account_id     string `json:"account_id"`
	Account_name   string `json:"account_name"`
	Account_avatar string `json:"account_avatar"`
	Is_child       bool   `json:"is_child"`
	Bus_type       int    `json:"bus_type"`
	Uid            string `json:"uid"`
	Auth_code      string `json:"auth_code"`
}
