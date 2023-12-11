type T struct {
	StaticBgUrl  string `json:"static_bg_url"`
	ForceProcess bool   `json:"force_process"`
	PopupType    string `json:"popup_type"`
	RateControl  struct {
		ClientEnabled bool   `json:"client_enabled"`
		KeyTpl        string `json:"key_tpl"`
		RateScene     int    `json:"rate_scene"`
		RateSeconds   int    `json:"rate_seconds"`
		RateTimes     int    `json:"rate_times"`
	} `json:"rate_control"`
	WidgetLayout struct {
		AsyncBlockSideEffects []struct {
			Data struct {
				BizId string `json:"biz_id"`
			} `json:"data"`
			EventName string `json:"eventName"`
		} `json:"asyncBlockSideEffects"`
		BtnWidgets []struct {
			Height   int `json:"height"`
			Position struct {
				Left int `json:"left"`
				Top  int `json:"top"`
			} `json:"position"`
			SideEffects []struct {
				Data struct {
					Schema string `json:"schema,omitempty"`
				} `json:"data"`
				EventName string `json:"eventName"`
			} `json:"sideEffects"`
			Width int `json:"width"`
		} `json:"btn_widgets"`
		InfoWidget struct {
			Position struct {
				Left int `json:"left"`
				Top  int `json:"top"`
			} `json:"position"`
		} `json:"info_widget"`
		PreEffects []struct {
			Data struct {
				LottieUrl string `json:"lottie_url"`
			} `json:"data"`
			EventName string `json:"eventName"`
		} `json:"preEffects"`
		SideEffects []interface{} `json:"sideEffects"`
		Type        string        `json:"type"`
	} `json:"widget_layout"`
	ActivityId string `json:"activity_id"`
}
