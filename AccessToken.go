package go_wechat

import (
	"context"
	"fmt"
	"go-wechat/conf"
	"go-wechat/wechat_tools/http"
	"sync"
	"time"
)

var (
	tokenMu           sync.Mutex
	accessToken       string
	accessTokenExpire time.Time
)

// GetAccessToken 获取access_token，如果token已缓存且未过期，则直接返回缓存的token
// 否则，调用微信接口获取新的token，并更新缓存 904ms->284ms
func (api *WexinApi) GetAccessToken(ctx context.Context) (string, error) {
	// 检查缓存是否有效
	tokenMu.Lock()
	defer tokenMu.Unlock() // 使用 defer 语句确保在函数返回之前释放锁

	if accessToken != "" && accessTokenExpire.After(time.Now()) {
		return accessToken, nil
	}

	// 缓存无效，调用微信接口获取新的 token
	cfg := api
	appID := cfg.AppId
	secret := cfg.AppSecret
	var resp conf.AccessTokenResp
	url := cfg.ServerHost + conf.GetAccessTokenApi
	err := http.Get(ctx, url+"?grant_type=client_credential&appid="+appID+"&secret="+secret, nil, &resp, nil)
	if err != nil {
		return "", err
	} else if resp.ErrorMsg != "" {
		return "", fmt.Errorf(resp.ErrorMsg)
	}
	// 更新缓存
	accessToken = resp.AccessToken
	accessTokenExpire = time.Now().Add(time.Second * time.Duration(resp.ExpiresIn))

	return accessToken, nil
}

// GetStableAccessToken 获取稳定accesstoken
func (api *WexinApi) GetStableAccessToken(ctx context.Context) (string, error) {
	var req conf.StableAccessReq
	cfg := api
	req.GrantType = "client_credential"
	req.AppId = cfg.AppId
	req.Secret = cfg.AppSecret
	req.ForceRefresh = false
	var resp conf.AccessTokenResp

	fmt.Println(req)
	//{grant_type:"client_credential",appid: "wx1092b036ca86acce",secret: "67af2e7c617064586b5dbe16a2e1a0f0",force_refresh: false}
	url := cfg.ServerHost + conf.GetStableAccessTokenApi

	err := http.Post(ctx, url, nil, req, &resp, nil)
	fmt.Println(resp)

	if err != nil {
		fmt.Println("eror")
		return "", err
	} else if resp.ErrorMsg != "" {
		return "", fmt.Errorf(resp.ErrorMsg)
	}
	return resp.AccessToken, nil
}
