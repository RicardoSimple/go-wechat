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
