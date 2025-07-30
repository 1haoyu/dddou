package OrderModel

type OrderItem struct {
	Errno int    `json:"errno"`
	St    int    `json:"st"`
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Data  struct {
		OrderID   string `json:"order_id"`
		OrderBase struct {
			OperationActions []string `json:"operation_actions"`
			Remark           string   `json:"remark"`
			Star             int      `json:"star"`
			Labels           []struct {
				Key        string `json:"key"`
				Label      string `json:"label"`
				Value      string `json:"value"`
				Hover      string `json:"hover"`
				HoverTitle string `json:"hover_title"`
				Type       string `json:"type"`
				ValueHover string `json:"value_hover"`
				Status     string `json:"status"`
				StatusDesc string `json:"status_desc"`
				URL        string `json:"url"`
			} `json:"labels"`
			ShopOrderTag []struct {
				Key               string      `json:"key"`
				Text              string      `json:"text"`
				HoverText         string      `json:"hover_text"`
				TagType           string      `json:"tag_type"`
				HelpDoc           string      `json:"help_doc"`
				Icon              string      `json:"icon"`
				BelongTo          string      `json:"belong_to"`
				QuestionHoverText string      `json:"question_hover_text"`
				Position          string      `json:"position"`
				URLConfigMaps     interface{} `json:"url_config_maps"`
			} `json:"shop_order_tag"`
			OrderStatusInfo struct {
				DeadLineTime           int    `json:"dead_line_time"`
				OrderStatusText        string `json:"order_status_text"`
				OrderStatusRemark      string `json:"order_status_remark"`
				OrderStatusOverRemark  string `json:"order_status_over_remark"`
				BorderURL              string `json:"border_url"`
				IllustrationURL        string `json:"illustration_url"`
				ShipTimeChangeDesc     string `json:"ship_time_change_desc"`
				ShipTimeChangeHover    string `json:"ship_time_change_hover"`
				OrderStatusShortRemark string `json:"order_status_short_remark"`
				OrderStatusStyle       struct {
				} `json:"order_status_style"`
				OrderStatusRemarkType         string        `json:"order_status_remark_type"`
				OrderStatusRemarkTypeV2       string        `json:"order_status_remark_type_v2"`
				OrderStatusOverRemarkType     string        `json:"order_status_over_remark_type"`
				OrderStatusRemarkScene        string        `json:"order_status_remark_scene"`
				EarlyPromiseDeadlineTimeDelta int           `json:"early_promise_deadline_time_delta"`
				EarlyPromiseRemark            string        `json:"early_promise_remark"`
				EarlyPromiseRemarkV2          string        `json:"early_promise_remark_v2"`
				EarlyPromiseOverRemark        string        `json:"early_promise_over_remark"`
				EarlyPromiseRemarkType        string        `json:"early_promise_remark_type"`
				EarlyPromiseRemarkTypeV2      string        `json:"early_promise_remark_type_v2"`
				EarlyPromiseOverRemarkType    string        `json:"early_promise_over_remark_type"`
				DeadLineTimeDelta             int           `json:"dead_line_time_delta"`
				OrderStatusRemarkV2           string        `json:"order_status_remark_v2"`
				OrderStatusOverRemarkV2       string        `json:"order_status_over_remark_v2"`
				EarlyPromiseDeadlineTime      int           `json:"early_promise_deadline_time"`
				CompareWith                   int           `json:"compare_with"`
				RecommendLogisticsCompanyList []interface{} `json:"recommend_logistics_company_list"`
			} `json:"order_status_info"`
			OrderStateFlow []struct {
				Text      string `json:"text"`
				Status    int    `json:"status"`
				IsCurrent bool   `json:"is_current"`
			} `json:"order_state_flow"`
			SerialNumbers     []interface{} `json:"serial_numbers"`
			IsActivityProduct bool          `json:"is_activity_product"`
			OrderPhaseText    struct {
			} `json:"order_phase_text"`
			NeedSerialNumber bool `json:"need_serial_number"`
			BaseLabels       []struct {
				Title  string `json:"title"`
				Labels []struct {
					Key        string `json:"key"`
					Label      string `json:"label"`
					Value      string `json:"value"`
					Hover      string `json:"hover"`
					HoverTitle string `json:"hover_title"`
					Type       string `json:"type"`
					ValueHover string `json:"value_hover"`
					Status     string `json:"status"`
					StatusDesc string `json:"status_desc"`
					URL        string `json:"url"`
				} `json:"labels"`
			} `json:"base_labels"`
		} `json:"order_base"`
		Receiver struct {
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
			CanView           int           `json:"can_view"`
			PostTelType       int           `json:"post_tel_type"`
			ExpireTime        int           `json:"expire_time"`
			IsShowEditAddress bool          `json:"is_show_edit_address"`
			CanPostpone       bool          `json:"can_postpone"`
			ExtensionNumber   string        `json:"extension_number"`
			PostTelMask       string        `json:"post_tel_mask"`
			AddressTag        []interface{} `json:"address_tag"`
			UserAccountInfos  interface{}   `json:"user_account_infos"`
			UIType            string        `json:"ui_type"`
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
		} `json:"receiver"`
		BuyUser struct {
			CanView int    `json:"can_view"`
			UserID  string `json:"user_id"`
			Labels  []struct {
				Key        string `json:"key"`
				Label      string `json:"label"`
				Value      string `json:"value"`
				Hover      string `json:"hover"`
				HoverTitle string `json:"hover_title"`
				Type       string `json:"type"`
				ValueHover string `json:"value_hover"`
				Status     string `json:"status"`
				StatusDesc string `json:"status_desc"`
				URL        string `json:"url"`
			} `json:"labels"`
			Title          string `json:"title"`
			UserProfileTag []struct {
				Key               string      `json:"key"`
				Text              string      `json:"text"`
				HoverText         string      `json:"hover_text"`
				TagType           string      `json:"tag_type"`
				HelpDoc           string      `json:"help_doc"`
				Icon              string      `json:"icon"`
				BelongTo          string      `json:"belong_to"`
				QuestionHoverText string      `json:"question_hover_text"`
				Position          string      `json:"position"`
				URLConfigMaps     interface{} `json:"url_config_maps"`
			} `json:"user_profile_tag"`
			BuyerWords string `json:"buyer_words"`
		} `json:"buy_user"`
		Delivery struct {
			Labels []struct {
				Key        string `json:"key"`
				Label      string `json:"label"`
				Value      string `json:"value"`
				Hover      string `json:"hover"`
				HoverTitle string `json:"hover_title"`
				Type       string `json:"type"`
				ValueHover string `json:"value_hover"`
				Status     string `json:"status"`
				StatusDesc string `json:"status_desc"`
				URL        string `json:"url"`
			} `json:"labels"`
		} `json:"delivery"`
		Other struct {
			Labels []struct {
				Key        string `json:"key"`
				Label      string `json:"label"`
				Value      string `json:"value"`
				Hover      string `json:"hover"`
				HoverTitle string `json:"hover_title"`
				Type       string `json:"type"`
				ValueHover string `json:"value_hover"`
				Status     string `json:"status"`
				StatusDesc string `json:"status_desc"`
				URL        string `json:"url"`
			} `json:"labels"`
		} `json:"other"`
		InspectionInfo interface{} `json:"inspection_info"`
		Product        struct {
			PayAmount string `json:"pay_amount"`
			Sku       []struct {
				SkuOrderID        string        `json:"sku_order_id"`
				TotalAmount       string        `json:"total_amount"`
				ProductID         string        `json:"product_id"`
				ProductName       string        `json:"product_name"`
				Img               string        `json:"img"`
				SkuPrice          string        `json:"sku_price"`
				SkuNum            int           `json:"sku_num"`
				ItemWarehouseID   string        `json:"item_warehouse_id"`
				ItemWarehouseName string        `json:"item_warehouse_name"`
				HasSnapshot       bool          `json:"has_snapshot"`
				OrderStatusDesc   string        `json:"order_status_desc"`
				PolicyInfo        interface{}   `json:"policy_info"`
				PrivilegeInfoList []interface{} `json:"privilege_info_list"`
				AfterSaleInfo     struct {
					URLAfterSaleNo                int    `json:"url_after_sale_no"`
					AfterSaleText                 string `json:"after_sale_text"`
					AfterSaleRemark               string `json:"after_sale_remark"`
					AfterSaleID                   string `json:"after_sale_id"`
					HasPreSale                    bool   `json:"has_pre_sale"`
					AftersaleStatusClassColorType int    `json:"aftersale_status_class_color_type"`
					AftersaleStatusClassString    string `json:"aftersale_status_class_string"`
					AfterSaleStatusRemark         string `json:"after_sale_status_remark"`
					AfterSaleStatusRemarkHover    string `json:"after_sale_status_remark_hover"`
				} `json:"after_sale_info"`
				Tags []struct {
					Key               string      `json:"key"`
					Text              string      `json:"text"`
					HoverText         string      `json:"hover_text"`
					TagType           string      `json:"tag_type"`
					HelpDoc           string      `json:"help_doc"`
					Icon              string      `json:"icon"`
					BelongTo          string      `json:"belong_to"`
					QuestionHoverText string      `json:"question_hover_text"`
					Position          string      `json:"position"`
					URLConfigMaps     interface{} `json:"url_config_maps"`
				} `json:"tags"`
				SkuBundleSubSkus interface{} `json:"sku_bundle_sub_skus"`
				Properties       []struct {
					Key        string `json:"key"`
					Text       string `json:"text"`
					Action     string `json:"action"`
					ActionText string `json:"action_text"`
				} `json:"properties"`
				DepositFindProduct       interface{} `json:"deposit_find_product"`
				SkuCustomizationInfo     interface{} `json:"sku_customization_info"`
				OrderStatusRemark        string      `json:"order_status_remark"`
				OrderStatusRemarkHover   string      `json:"order_status_remark_hover"`
				ItemOutWarehouseID       string      `json:"item_out_warehouse_id"`
				LowPriceInfo             interface{} `json:"low_price_info"`
				CampainDetail            interface{} `json:"campain_detail"`
				AccessorySkuInfo         interface{} `json:"accessory_sku_info"`
				RelationSkuOrderIds      interface{} `json:"relation_sku_order_ids"`
				ProductCustomPropertites interface{} `json:"product_custom_propertites"`
				RelatedOrderInfo         interface{} `json:"related_order_info"`
				SkuOrderStatusInfo       struct {
					DeadLineTime                    int           `json:"dead_line_time"`
					OrderStatusIcon                 int           `json:"order_status_icon"`
					OrderStatusText                 string        `json:"order_status_text"`
					OrderStatusRemark               string        `json:"order_status_remark"`
					OrderStatusOverRemark           string        `json:"order_status_over_remark"`
					ShowRule                        bool          `json:"show_rule"`
					OrderStatusOverRemarkV2         string        `json:"order_status_over_remark_v2"`
					IsAppointmentShip               bool          `json:"is_appointment_ship"`
					AppointmentShipTime             int           `json:"appointment_ship_time"`
					AppointmentShipTimeStr          string        `json:"appointment_ship_time_str"`
					CutDownSecond                   int           `json:"cut_down_second"`
					OrderStatusRemarkV2             string        `json:"order_status_remark_v2"`
					ShipTimeChangeDesc              string        `json:"ship_time_change_desc"`
					ShipTimeChangeHover             string        `json:"ship_time_change_hover"`
					OrderStatusRemarkType           string        `json:"order_status_remark_type"`
					OrderStatusRemarkHover          string        `json:"order_status_remark_hover"`
					OrderExtraText                  string        `json:"order_extra_text"`
					OrderExtraTextHover             string        `json:"order_extra_text_hover"`
					OrderExtraTextType              string        `json:"order_extra_text_type"`
					DeadLineTimeV3                  int           `json:"dead_line_time_v3"`
					OrderStatusRemarkV3             string        `json:"order_status_remark_v3"`
					OrderStatusOverRemarkV3         string        `json:"order_status_over_remark_v3"`
					OrderStatusRemarkTypeV3         string        `json:"order_status_remark_type_v3"`
					OrderStatusShortRemark          string        `json:"order_status_short_remark"`
					OrderStatusRemarkBackgroundType string        `json:"order_status_remark_background_type"`
					OrderStatusRemarkTypeV2         string        `json:"order_status_remark_type_v2"`
					OrderStatusOverRemarkType       string        `json:"order_status_over_remark_type"`
					OrderStatusRemarkScene          string        `json:"order_status_remark_scene"`
					EarlyPromiseDeadlineTimeDelta   int           `json:"early_promise_deadline_time_delta"`
					EarlyPromiseRemark              string        `json:"early_promise_remark"`
					EarlyPromiseRemarkV2            string        `json:"early_promise_remark_v2"`
					EarlyPromiseOverRemark          string        `json:"early_promise_over_remark"`
					EarlyPromiseRemarkType          string        `json:"early_promise_remark_type"`
					EarlyPromiseRemarkTypeV2        string        `json:"early_promise_remark_type_v2"`
					EarlyPromiseOverRemarkType      string        `json:"early_promise_over_remark_type"`
					DeadLineTimeDelta               int           `json:"dead_line_time_delta"`
					EarlyPromiseDeadlineTime        int           `json:"early_promise_deadline_time"`
					CompareWith                     int           `json:"compare_with"`
					RecommendLogisticsCompanyList   []interface{} `json:"recommend_logistics_company_list"`
				} `json:"sku_order_status_info"`
				SkuURLMap struct {
					CommissionFirstOrderWithdraw struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"CommissionFirstOrderWithdraw"`
					CommissionWithdraw struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"CommissionWithdraw"`
					FullFreeWithdraw struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"FullFreeWithdraw"`
					SignInWithdraw struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"SignInWithdraw"`
					AccountChangeDetail struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"account_change_detail"`
					AfterSaleList struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"after_sale_list"`
					AheadEarlyPromiseRemark struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color string `json:"color"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string `json:"text"`
							Placeholder []struct {
								Key   string `json:"key"`
								Type  string `json:"type"`
								Label string `json:"label"`
								Value string `json:"value"`
								Style struct {
								} `json:"style"`
								Text      string      `json:"text"`
								RichHover interface{} `json:"rich_hover"`
							} `json:"placeholder"`
							ModalTitle string `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"ahead_early_promise_remark"`
					AheadPeriodReceiveRemark struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"ahead_period_receive_remark"`
					AheadPeriodSendRemark struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"ahead_period_send_remark"`
					AheadStatusRemark struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color string `json:"color"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string `json:"text"`
							Placeholder []struct {
								Key   string `json:"key"`
								Type  string `json:"type"`
								Label string `json:"label"`
								Value string `json:"value"`
								Style struct {
								} `json:"style"`
								Text      string      `json:"text"`
								RichHover interface{} `json:"rich_hover"`
							} `json:"placeholder"`
							ModalTitle string `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"ahead_status_remark"`
					AlreadyChangePromise struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Border       string `json:"border"`
								BorderRadius string `json:"borderRadius"`
								Color        string `json:"color"`
								Padding      string `json:"padding"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"already_change_promise"`
					AppealCenterURL struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"appeal_center_url"`
					CheckNegotiation struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"check_negotiation"`
					CheckSupplyOrder struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"check_supply_order"`
					EarlySendRule struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"early_send_rule"`
					GiftInfoURL struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"gift_info_url"`
					GiftReceiveHover struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"gift_receive_hover"`
					LateDeliveryHover struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"late_delivery_hover"`
					OrderStatusEndHover struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color string `json:"color"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string        `json:"text"`
							Placeholder []interface{} `json:"placeholder"`
							ModalTitle  string        `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"order_status_end_hover"`
					PaymentFreeze struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color      string `json:"color"`
								FontWeight string `json:"font-weight"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string `json:"text"`
							Placeholder []struct {
								Key   string `json:"key"`
								Type  string `json:"type"`
								Label string `json:"label"`
								Value string `json:"value"`
								Style struct {
								} `json:"style"`
								Text      string      `json:"text"`
								RichHover interface{} `json:"rich_hover"`
							} `json:"placeholder"`
							ModalTitle string `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"payment_freeze"`
					PriorityDeliveryDetail struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"priority_delivery_detail"`
					PriorityDeliveryRemark struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color string `json:"color"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover struct {
							Text        string `json:"text"`
							Placeholder []struct {
								Key   string `json:"key"`
								Type  string `json:"type"`
								Label string `json:"label"`
								Value string `json:"value"`
								Style struct {
								} `json:"style"`
								Text      string      `json:"text"`
								RichHover interface{} `json:"rich_hover"`
							} `json:"placeholder"`
							ModalTitle string `json:"modal_title"`
						} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"priority_delivery_remark"`
					PromiseModal struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color  string `json:"color"`
								Cursor string `json:"cursor"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"promise_modal"`
					QicURL struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"qic_url"`
					QicURLForDetail struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"qic_url_for_detail"`
					RechargeLink struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"recharge_link"`
					RemoteCostRule struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"remote_cost_rule"`
					SendRule struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"send_rule"`
					SupplyChainOrderRechargeLink struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"supply_chain_order_recharge_link"`
					ToReport struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"to_report"`
					UserAppointmentReceiptTimeDesc struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color string `json:"color"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"user_appointment_receipt_time_desc"`
					UserAppointmentShipTimeDesc struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style struct {
								Color string `json:"color"`
							} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"user_appointment_ship_time_desc"`
					WuyouIcon struct {
						Text           string `json:"text"`
						URL            string `json:"url"`
						Type           string `json:"type"`
						ComponentProps struct {
							Style interface{} `json:"style"`
						} `json:"component_props"`
						RichHover  interface{} `json:"rich_hover"`
						HoverProps struct {
							OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
						} `json:"hover_props"`
					} `json:"wuyou_icon"`
				} `json:"sku_url_map"`
			} `json:"sku"`
			PriceDetail []struct {
				Key        string `json:"key"`
				Label      string `json:"label"`
				Value      string `json:"value"`
				Hover      string `json:"hover"`
				HoverTitle string `json:"hover_title"`
				Type       string `json:"type"`
				ValueHover string `json:"value_hover"`
				Status     string `json:"status"`
				StatusDesc string `json:"status_desc"`
				URL        string `json:"url"`
			} `json:"price_detail"`
			Promotion []struct {
				Key          string `json:"key"`
				Label        string `json:"label"`
				Value        string `json:"value"`
				Hover        string `json:"hover"`
				ExtraInfoMap struct {
					ActivityCreator     string `json:"activity_creator"`
					ActivityCreatorDesc string `json:"activity_creator_desc"`
					ShareCostMap        []struct {
						ShareCostAmount string `json:"share_cost_amount"`
						ShareCostDesc   string `json:"share_cost_desc"`
						ShareCostType   string `json:"share_cost_type"`
					} `json:"share_cost_map"`
				} `json:"extra_info_map"`
				HoverTitle string `json:"hover_title"`
				Type       string `json:"type"`
				ValueHover string `json:"value_hover"`
				Status     string `json:"status"`
				StatusDesc string `json:"status_desc"`
				URL        string `json:"url"`
			} `json:"promotion"`
			AmountDetailMap struct {
				DiscountAmount []struct {
					Key        string `json:"key"`
					Label      string `json:"label"`
					Value      string `json:"value"`
					Hover      string `json:"hover"`
					HoverTitle string `json:"hover_title"`
					Type       string `json:"type"`
					ValueHover string `json:"value_hover"`
					Status     string `json:"status"`
					StatusDesc string `json:"status_desc"`
					URL        string `json:"url"`
				} `json:"discount_amount"`
				PayAmount []struct {
					Key        string `json:"key"`
					Label      string `json:"label"`
					Value      string `json:"value"`
					Hover      string `json:"hover"`
					HoverTitle string `json:"hover_title"`
					Type       string `json:"type"`
					ValueHover string `json:"value_hover"`
					Status     string `json:"status"`
					StatusDesc string `json:"status_desc"`
					URL        string `json:"url"`
				} `json:"pay_amount"`
			} `json:"amount_detail_map"`
		} `json:"product"`
		Insurance                   interface{} `json:"insurance"`
		Ep                          interface{} `json:"ep"`
		WriteOff                    interface{} `json:"write_off"`
		PhysicalCard                interface{} `json:"physical_card"`
		ReceiveType                 int         `json:"receive_type"`
		SampleConsign               interface{} `json:"sample_consign"`
		AppealEntrance              string      `json:"appeal_entrance"`
		TransferReceiverInfo        interface{} `json:"transfer_receiver_info"`
		SupermarketOrderSerialNo    string      `json:"supermarket_order_serial_no"`
		TotalEstimatedIncludeRefund interface{} `json:"total_estimated_include_refund"`
		ActionMap                   struct {
			ContactBuyer struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
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
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUI     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionURL    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"contactBuyer"`
			ForbiddenUser struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
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
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUI     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionURL    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"forbiddenUser"`
			Remark struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
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
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUI     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionURL    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"remark"`
			Remit struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
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
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUI     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionURL    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"remit"`
			Reship struct {
				BizAction        string `json:"biz_action"`
				BizActionDesc    string `json:"biz_action_desc"`
				BizActionOptions struct {
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
				} `json:"biz_action_options"`
				BizActionName   string `json:"biz_action_name"`
				HasActionRecord bool   `json:"has_action_record"`
				BizActionUI     string `json:"biz_action_ui"`
				BizActionHover  string `json:"biz_action_hover"`
				BizActionURL    string `json:"biz_action_url"`
				BizActionType   string `json:"biz_action_type"`
			} `json:"reship"`
		} `json:"action_map"`
		URLMap struct {
			CommissionFirstOrderWithdraw struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"CommissionFirstOrderWithdraw"`
			CommissionWithdraw struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"CommissionWithdraw"`
			FullFreeWithdraw struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"FullFreeWithdraw"`
			SignInWithdraw struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"SignInWithdraw"`
			AccountChangeDetail struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"account_change_detail"`
			AfterSaleList struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"after_sale_list"`
			AheadEarlyPromiseRemark struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_early_promise_remark"`
			AheadPeriodReceiveRemark struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_period_receive_remark"`
			AheadPeriodSendRemark struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_period_send_remark"`
			AheadStatusRemark struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"ahead_status_remark"`
			AlreadyChangePromise struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Border       string `json:"border"`
						BorderRadius string `json:"borderRadius"`
						Color        string `json:"color"`
						Padding      string `json:"padding"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"already_change_promise"`
			AppealCenterURL struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"appeal_center_url"`
			CheckNegotiation struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"check_negotiation"`
			CheckSupplyOrder struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"check_supply_order"`
			EarlySendRule struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"early_send_rule"`
			GiftInfoURL struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"gift_info_url"`
			GiftReceiveHover struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string        `json:"text"`
					Placeholder []interface{} `json:"placeholder"`
					ModalTitle  string        `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"gift_receive_hover"`
			ImmunityDetailInfoURL struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps interface{} `json:"hover_props"`
			} `json:"immunity_detail_info_url"`
			PaymentFreeze struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color      string `json:"color"`
						FontWeight string `json:"font-weight"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"payment_freeze"`
			PriorityDeliveryDetail struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"priority_delivery_detail"`
			PriorityDeliveryRemark struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover struct {
					Text        string `json:"text"`
					Placeholder []struct {
						Key   string `json:"key"`
						Type  string `json:"type"`
						Label string `json:"label"`
						Value string `json:"value"`
						Style struct {
						} `json:"style"`
						Text      string      `json:"text"`
						RichHover interface{} `json:"rich_hover"`
					} `json:"placeholder"`
					ModalTitle string `json:"modal_title"`
				} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"priority_delivery_remark"`
			PromiseModal struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style struct {
						Color  string `json:"color"`
						Cursor string `json:"cursor"`
					} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"promise_modal"`
			QicURL struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"qic_url"`
			QicURLForDetail struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"qic_url_for_detail"`
			RechargeLink struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"recharge_link"`
			RemoteCostRule struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"remote_cost_rule"`
			SendRule struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"send_rule"`
			SupplierShopName             interface{} `json:"supplier_shop_name"`
			SupplyChainOrderRechargeLink struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"supply_chain_order_recharge_link"`
			ToReport struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"to_report"`
			UserAppointmentReceiptTimeDesc struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"user_appointment_receipt_time_desc"`
			UserAppointmentShipTimeDesc struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"user_appointment_ship_time_desc"`
			WuyouIcon struct {
				Text           string `json:"text"`
				URL            string `json:"url"`
				Type           string `json:"type"`
				ComponentProps struct {
					Style interface{} `json:"style"`
				} `json:"component_props"`
				RichHover  interface{} `json:"rich_hover"`
				HoverProps struct {
					OverlayInnerStyle interface{} `json:"overlayInnerStyle"`
				} `json:"hover_props"`
			} `json:"wuyou_icon"`
		} `json:"url_map"`
		ExtraInfo struct {
			SendRule struct {
				ValueHover     string      `json:"value_hover"`
				ValueURL       string      `json:"value_url"`
				ValueURLConfig interface{} `json:"value_url_config"`
				Hovers         interface{} `json:"hovers"`
			} `json:"sendRule"`
		} `json:"extra_info"`
		OrderNoticeList []struct {
			Announcement string   `json:"announcement"`
			Slogan       []string `json:"slogan"`
			URLMap       struct {
				RegularQuestion struct {
					Text string `json:"text"`
					URL  string `json:"url"`
				} `json:"regular_question"`
			} `json:"url_map"`
			Type string `json:"type"`
		} `json:"order_notice_list"`
		ReceiverCommon     interface{} `json:"receiver_common"`
		GiftReceiveTimeStr string      `json:"gift_receive_time_str"`
		OrderRecord        []struct {
			Operator   string        `json:"operator"`
			Action     string        `json:"action"`
			Desc       string        `json:"desc"`
			CreateTime string        `json:"create_time"`
			Role       int           `json:"role"`
			OpID       int           `json:"op_id"`
			Evidence   []interface{} `json:"evidence"`
		} `json:"order_record"`
		OrderService interface{} `json:"order_service"`
		SellerLog    interface{} `json:"seller_log"`
	} `json:"data"`
	Page      int `json:"page"`
	Total     int `json:"total"`
	Size      int `json:"size"`
	TotalPage int `json:"total_page"`
}
