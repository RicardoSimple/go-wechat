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

// GeneralResp 微信通用错误返回
type GeneralResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type GetMenuResp struct {
	IsMenuOpen   int  `json:"is_menu_open"`
	SelfMenuInfo Menu `json:"selfmenu_info"`
	GeneralResp
}

type GetUserInfoResp struct {
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Language       string `json:"language"`
	SubscribeTime  int    `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	Groupid        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
	GeneralResp
}

type GetUserListResp struct {
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
	GeneralResp
}

type SetUserRemarkReq struct {
	Openid string `json:"openid"`
	Remark string `json:"remark"`
}

type CreateTagResp struct {
	// {   "tag":{ "id":134,//标签id "name":"广东"   } }
	Tag struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tag"`
	GeneralResp
}
type CreateTagReq struct {
	Tag struct {
		Name string `json:"name"`
	} `json:"tag"`
}

type GetUserListOfTagResp struct {
	Count      int        `json:"count"`
	Data       OpenIdList `json:"data"`
	NextOpenid string     `json:"next_openid"`
	GeneralResp
}
type GetTagListOfUserReq struct {
	Tagid      int    `json:"tagid"`
	NextOpenid string `json:"next_openid"`
}
type MembersBatchReq struct {
	OpenIdList []string `json:"openid_list"`
	TagId      string   `json:"tagid"`
}

type BlackListReq struct {
	OpenIdList []string `json:"openid_list"`
}
type BlackListResp struct {
	Total int `json:"total"`
	GetUserListOfTagResp
}
