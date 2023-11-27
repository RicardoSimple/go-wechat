package go_wechat

import (
	"errors"
)

type WexinApi struct {
	AppId          string `json:"appId"`
	AppSecret      string `json:"secret"`
	ServerHost     string `json:"server"`
	EncodingAESKey string `json:"encodingAESKey"`
	Token          string `json:"token"`
}

func StartWeixinApi(config WexinApi) (*WexinApi, error) {
	if config.AppId == "" || config.AppSecret == "" {
		return nil, errors.New("appid or appsecret is nil")
	}
	return &WexinApi{
		AppId:          config.AppId,
		AppSecret:      config.AppSecret,
		ServerHost:     config.ServerHost,
		EncodingAESKey: config.EncodingAESKey,
		Token:          config.Token,
	}, nil
}
