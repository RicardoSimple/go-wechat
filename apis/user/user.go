package user

import (
	"fmt"
	"go-wechat"
	"go-wechat/conf"
)

const (
	GetUserInfoApi           = "/cgi-bin/user/info"
	GetUserListApi           = "/cgi-bin/user/get"
	SetUserRemarkApi         = "/cgi-bin/user/info/updateremark"
	CreateUserTagsApi        = "/cgi-bin/tags/create"
	GetCreatedTagsApi        = "/cgi-bin/tags/get"
	UpdateTagsApi            = "/cgi-bin/tags/update"
	DeleteTagsApi            = "/cgi-bin/tags/delete"
	MembersBatchTaggingApi   = "/cgi-bin/tags/members/batchtagging"
	MembersBatchUnTaggingApi = "/cgi-bin/tags/members/batchuntagging"
	GetTagListOfUserApi      = "/cgi-bin/tags/getidlist"
	GetBlackListApi          = "/cgi-bin/tags/members/getblacklist"
	BatchBlackListApi        = "/cgi-bin/tags/members/batchblacklist"
	BatchUnBlackListApi      = "/cgi-bin/tags/members/batchunblacklist"
)

// GetUserInfo 通过OpenID来获取用户基本信息 lang为返回国家地区语言版本
func GetUserInfo(api *go_wechat.WexinApi, openId string, lang conf.LangParam) (conf.GetUserInfoResp, error) {
	var resp conf.GetUserInfoResp
	var err error
	if lang == conf.DEFAULT {
		err = api.WxClient.GetWithAccess(GetUserInfoApi+"?openid="+openId, nil, resp, nil)
	} else {
		err = api.WxClient.GetWithAccess(GetUserInfoApi+"?openid="+openId+"&lang="+string(lang), nil, resp, nil)
	}
	if err != nil {
		return conf.GetUserInfoResp{}, err
	}
	return resp, nil
}

// GetUserList 获取账号的关注者列表 openId为第一个拉取的openid，可传入空字符串，默认从头拉取
func GetUserList(api *go_wechat.WexinApi, openId string) (conf.GetUserListResp, error) {
	var resp conf.GetUserListResp
	reqUrl := GetUserListApi
	if openId != "" {
		reqUrl = reqUrl + "?next_openid=" + openId
	}
	err := api.WxClient.GetWithAccess(reqUrl, nil, resp, nil)
	if err != nil {
		return conf.GetUserListResp{}, err
	}
	return resp, nil
}

// SetUserRemark 给指定用户设置备注名
func SetUserRemark(api *go_wechat.WexinApi, openId, remark string) error {
	if openId == "" {
		return fmt.Errorf("openid is empty")
	}
	var resp conf.GeneralResp
	req := &conf.SetUserRemarkReq{
		Openid: openId,
		Remark: remark,
	}
	err := api.WxClient.PostWithAccess(SetUserRemarkApi, nil, req, resp, nil)
	if err != nil {
		return nil
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("error:" + resp.ErrMsg)
	}
	return nil
}

// CreateTag 创建标签
func CreateTag(api *go_wechat.WexinApi, tagName string) (conf.CreateTagResp, error) {
	req := &conf.CreateTagReq{Tag: struct {
		Name string `json:"name"`
	}(struct{ Name string }{Name: tagName})}
	var resp conf.CreateTagResp
	err := api.WxClient.PostWithAccess(CreateUserTagsApi, nil, req, resp, nil)
	if err != nil {
		return conf.CreateTagResp{}, err
	} else if resp.ErrCode != 0 {
		return conf.CreateTagResp{}, fmt.Errorf("error:" + resp.ErrMsg)
	}
	return resp, nil
}

// GetCreatedTags 获取公众号已创建的标签
func GetCreatedTags(api *go_wechat.WexinApi) (conf.TagList, error) {
	var resp conf.TagList
	err := api.WxClient.GetWithAccess(GetCreatedTagsApi, nil, resp, nil)
	if err != nil {
		return conf.TagList{}, err
	}
	return resp, nil
}

// UpdateTags 编辑标签
func UpdateTags(api *go_wechat.WexinApi, tagId int, tagName string) (err error) {
	var resp conf.GeneralResp
	req := &conf.CreateTagResp{Tag: struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}(struct {
		Id   int
		Name string
	}{Id: tagId, Name: tagName})}
	err = api.WxClient.PostWithAccess(UpdateTagsApi, nil, req, resp, nil)
	if err != nil {
		return
	} else if resp.ErrCode != 0 {
		err = fmt.Errorf("error :" + resp.ErrMsg)
		return
	}
	return
}

// DeleteTags 删除标签
func DeleteTags(api *go_wechat.WexinApi, tagId int) error {
	var req = struct {
		Tag struct {
			id int `json:"id"`
		} `json:"tag"`
	}{
		Tag: struct {
			id int `json:"id"`
		}{id: tagId},
	}
	var resp conf.GeneralResp
	err := api.WxClient.PostWithAccess(DeleteTagsApi, nil, req, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("error :" + resp.ErrMsg)

	}
	return nil
}

// MembersBatchTagging  批量为用户打标签
func MembersBatchTagging(api *go_wechat.WexinApi, openIdList []string, tagId string) error {
	var req = &conf.MembersBatchReq{
		OpenIdList: openIdList,
		TagId:      tagId,
	}
	var resp conf.GeneralResp
	err := api.WxClient.PostWithAccess(MembersBatchTaggingApi, nil, req, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("error :" + resp.ErrMsg)

	}
	return nil
}

// MembersBatchUnTagging 批量为用户取消标签
func MembersBatchUnTagging(api *go_wechat.WexinApi, openIdList []string, tagId string) error {
	var req = &conf.MembersBatchReq{
		OpenIdList: openIdList,
		TagId:      tagId,
	}
	var resp conf.GeneralResp
	err := api.WxClient.PostWithAccess(MembersBatchUnTaggingApi, nil, req, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("error :" + resp.ErrMsg)

	}
	return nil
}

// GetTagListOfUser 获取标签下粉丝列表
func GetTagListOfUser(api *go_wechat.WexinApi, tagId int, openId string) (resp conf.GetUserListOfTagResp, err error) {
	var req = &conf.GetTagListOfUserReq{
		Tagid:      tagId,
		NextOpenid: openId,
	}
	err = api.WxClient.PostWithAccess(GetTagListOfUserApi, nil, req, resp, nil)
	if err != nil {
		return
	} else if resp.ErrCode != 0 {
		err = fmt.Errorf("error:" + resp.ErrMsg)
		return
	}
	return
}

// GetBlackList 获取黑名单
func GetBlackList(api *go_wechat.WexinApi) (resp conf.BlackListResp, err error) {
	err = api.WxClient.GetWithAccess(GetBlackListApi, nil, resp, nil)
	if err != nil {
		return
	} else if resp.ErrCode != 0 {
		err = fmt.Errorf("error: " + resp.ErrMsg)
	}
	return
}

// BatchBlack 批量绑定黑名单
func BatchBlack(api *go_wechat.WexinApi, openIdList []string) error {
	var req = conf.BlackListReq{OpenIdList: openIdList}
	var resp conf.GeneralResp
	err := api.WxClient.PostWithAccess(BatchBlackListApi, nil, req, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("error :" + resp.ErrMsg)

	}
	return nil
}

// BatchUnBlack 批量取消黑名单
func BatchUnBlack(api *go_wechat.WexinApi, openIdList []string) error {
	var req = conf.BlackListReq{OpenIdList: openIdList}
	var resp conf.GeneralResp
	err := api.WxClient.PostWithAccess(BatchUnBlackListApi, nil, req, resp, nil)
	if err != nil {
		return err
	} else if resp.ErrCode != 0 {
		return fmt.Errorf("error :" + resp.ErrMsg)

	}
	return nil
}
