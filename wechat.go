package go_wechat

import (
	"errors"
	"log"
	"os"
)

type WexinApi struct {
	WxAccount
	WxClient
	WxServer
	Logger *log.Logger
}

type Config struct {
	AppId          string `json:"appId"`
	AppSecret      string `json:"secret"`
	ServerHost     string `json:"server"`
	EncodingAESKey string `json:"encodingAESKey"`
	Token          string `json:"token"`
}

type WxAccount struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"secret"`
}

func NewApiInstance(config Config) (*WexinApi, error) {
	// 校验配置
	if config.AppId == "" || config.AppSecret == "" {
		return nil, errors.New("appid or appsecret is nil")
	}
	api := WexinApi{
		WxAccount: WxAccount{
			AppId:     config.AppId,
			AppSecret: config.AppSecret,
		},
		WxClient: WxClient{ServerHost: config.ServerHost},
		WxServer: WxServer{EncodingAESKey: config.EncodingAESKey,
			Token: config.Token},
	}
	api.Logger = log.New(os.Stdout, "[go-wechat]", log.LstdFlags|log.Llongfile)
	api.WxClient.api = &api
	return &api, nil
}
