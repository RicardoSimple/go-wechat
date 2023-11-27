package conf

type AccessTokenResp struct {
	ErrorCode   int    `json:"errorcode"`
	ErrorMsg    string `json:"errormsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
