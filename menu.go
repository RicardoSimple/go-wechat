package go_wechat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-wechat/conf"
	"go-wechat/wechat_tools/http"
)

// CreateMenu createMenu 创建自定义菜单
func (api *WexinApi) CreateMenu(ctx context.Context, menu *conf.Menu) error {
	// todo 校验menu是否正确
	accessToken, _ := api.GetAccessToken(ctx)
	// 发送HTTP POST请求创建自定义菜单
	createMenuURL := fmt.Sprintf(api.ServerHost+conf.CreateMenuApi+"?access_token=%s", accessToken)
	var resp conf.GeneralResp
	// 处理响应，检查是否成功创建了菜单
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(menu)
	err := http.Post(ctx, createMenuURL, nil, bf.String(), &resp, &resp)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("failed to create menu: %s", resp.ErrMsg)
	}
	return nil
}

// GetMenuConfig 查询自定义菜单
func (api *WexinApi) GetMenuConfig(ctx context.Context) (conf.GetMenuResp, error) {
	accessToken, _ := api.GetAccessToken(ctx)

	var resp conf.GetMenuResp
	url := fmt.Sprintf(api.ServerHost+conf.GetMenuConfigApi+"?access_token=%s", accessToken)

	err := http.Get(ctx, url, nil, resp, nil)
	if err != nil {
		return conf.GetMenuResp{}, err
	} else if resp.ErrCode != 0 {
		return conf.GetMenuResp{}, fmt.Errorf("failed to get menu config")
	}
	return resp, nil
}

// DeleteMenu 删除自定义菜单
func (api *WexinApi) DeleteMenu(ctx context.Context) error {
	accessToken, _ := api.GetAccessToken(ctx)
	url := fmt.Sprintf(api.ServerHost+conf.DeleteMenuApi+"?access_token=%s", accessToken)
	var resp conf.GeneralResp
	err := http.Get(ctx, url, nil, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("failed to delete")
	}
	return nil
}
