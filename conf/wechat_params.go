package conf

type AccessTokenResp struct {
	ErrorCode   int    `json:"errorcode"`
	ErrorMsg    string `json:"errormsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type StableAccessReq struct {
	GrantType    string `json:"grant_type"`
	AppId        string `json:"appid"`
	Secret       string `json:"secret"`
	ForceRefresh bool   `json:"force_refresh"` // true:force fresh token
}

type GeneralResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type GetMenuResp struct {
	IsMenuOpen   int  `json:"is_menu_open"`
	SelfMenuInfo Menu `json:"selfmenu_info"`
	GeneralResp
}
