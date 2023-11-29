package go_wechat

// WxServer 监听微信服务器
type WxServer struct {
	EncodingAESKey string `json:"encodingAESKey"`
	Token          string `json:"token"`
}
