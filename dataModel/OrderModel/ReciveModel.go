package OrderModel

type ReciveModel struct {
	Errno int    `json:"errno"`
	St    int    `json:"st"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Data  struct {
		VerifyType   string `json:"verify_type"`
		VerifyParams struct {
			SendText      string      `json:"send_text"`
			SendTextNote  string      `json:"send_text_note"`
			VerifyAccount string      `json:"verify_account"`
			DecisionConf  string      `json:"decision_conf"`
			BlockPop      interface{} `json:"block_pop"`
		} `json:"verify_params"`
		IsSend      int `json:"is_send"`
		ReceiveInfo struct {
			PostReceiver string `json:"post_receiver"`
			PostTel      string `json:"post_tel"`
			PostAddr     struct {
				Province struct {
					Name     string `json:"name"`
					ID       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"province"`
				City struct {
					Name     string `json:"name"`
					ID       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"city"`
				Town struct {
					Name     string `json:"name"`
					ID       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"town"`
				Street struct {
					Name     string `json:"name"`
					ID       string `json:"id"`
					HasChild bool   `json:"has_child"`
				} `json:"street"`
				Detail string `json:"detail"`
			} `json:"post_addr"`
			CanView           int         `json:"can_view"`
			PostTelType       int         `json:"post_tel_type"`
			ExpireTime        int         `json:"expire_time"`
			IsShowEditAddress bool        `json:"is_show_edit_address"`
			CanPostpone       bool        `json:"can_postpone"`
			ExtensionNumber   string      `json:"extension_number"`
			PostTelMask       string      `json:"post_tel_mask"`
			AddressTag        interface{} `json:"address_tag"`
			UserAccountInfos  interface{} `json:"user_account_infos"`
			UIType            string      `json:"ui_type"`
			BuyerTelInfo      struct {
				PayTel          string `json:"pay_tel"`
				PayTelType      int    `json:"pay_tel_type"`
				ExpireTime      int    `json:"expire_time"`
				CanPostpone     bool   `json:"can_postpone"`
				PayTelMask      string `json:"pay_tel_mask"`
				UIType          string `json:"ui_type"`
				CanView         int    `json:"can_view"`
				ExtensionNumber string `json:"extension_number"`
				NickName        string `json:"nick_name"`
			} `json:"buyer_tel_info"`
			ExtraMap interface{} `json:"extra_map"`
		} `json:"receive_info"`
		PreReceiveInfo    interface{} `json:"pre_receive_info"`
		NickName          string      `json:"nick_name"`
		ChosenOrderID     string      `json:"chosen_order_id"`
		RelateReceiveInfo struct {
			NickName string `json:"nick_name"`
		} `json:"relate_receive_info"`
	} `json:"data"`
	Page      int `json:"page"`
	Total     int `json:"total"`
	Size      int `json:"size"`
	TotalPage int `json:"total_page"`
}
