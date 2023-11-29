package menu

import (
	"bytes"
	"encoding/json"
	"fmt"
	go_wechat "go-wechat"
	"go-wechat/conf"
)

const (
	CreateMenuApi     = "/cgi-bin/menu/create"
	GetCurrentMenuApi = "/cgi-bin/get_current_selfmenu_info"
	DeleteMenuApi     = "/cgi-bin/menu/delete"
	GetMenuConfig     = "/cgi-bin/menu/get"
)

// CreateMenu createMenu 创建自定义菜单
func CreateMenu(api *go_wechat.WexinApi, menu *conf.Menu) error {
	// todo 校验menu是否正确

	var resp conf.GeneralResp
	// 处理响应，检查是否成功创建了菜单
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(menu)
	err := api.WxClient.PostWithAccess(CreateMenuApi, nil, bf.String(), &resp, &resp)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("failed to create menu: %s", resp.ErrMsg)
	}
	return nil
}

// GetCurrentMenu 查询自定义菜单
func GetCurrentMenu(api *go_wechat.WexinApi) (conf.GetMenuResp, error) {
	var resp conf.GetMenuResp

	err := api.WxClient.GetWithAccess(GetCurrentMenuApi, nil, resp, nil)
	if err != nil {
		return conf.GetMenuResp{}, err
	} else if resp.ErrCode != 0 {
		return conf.GetMenuResp{}, fmt.Errorf("failed to get menu config")
	}
	return resp, nil
}

// DeleteMenu 删除自定义菜单
func DeleteMenu(api *go_wechat.WexinApi) error {
	var resp conf.GeneralResp
	err := api.WxClient.GetWithAccess(DeleteMenuApi, nil, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("failed to delete")
	}
	return nil
}
