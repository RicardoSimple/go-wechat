# go-wechat

wexin sdk 对微信公众平台的集成环境

目的在于调用微信各接口都能去掉繁琐的逻辑部分,只留下传参->获取结果的过程

即将各个功能抽象成对应的方法

### 使用方法:

```go
	api, err := StartWeixinApi(
		WexinApi{
			AppId:          "your appId",
			AppSecret:      "your appsecret",
			Token:          "your config token",    // 需要与公众号配置一致，这里可不填
			ServerHost:     conf.ServerHostGeneral, //选取最近的服务器域名，默认为api.weixin.qq.com
			EncodingAESKey: "your config",
		})
	if err != nil {
		return
	}
	// 调用具体方法，传入context或其他参数
	api.GetAccessToken(ctx)
```

## 开发注意事项：

官方文档:[开发前必读 / 首页 (qq.com)](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

目前暂时是对微信公众号部分进行封装

后续可能会进行企业微信,微信支付,创建小商店等封装

微信接口返回错误码查询:[全局返回码说明](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Global_Return_Code.html)

### 开发流程

+ conf 配置/参数包
  
  + const为常量
  
  + structs为实体/模型
  
  + wechat_params为封装过程产生的变量,比如请求参数等

+ wechat_tools 工具包
  
  + http 为已封装的请求包

新增的接口需要在const中定义,例如:

```go
// 接口api
const (
	GetAccessTokenApi       = "/cgi-bin/token"
	GetStableAccessTokenApi = "/cgi-bin/stable_token"
)
```

需要新增的封装方法按照模块化的名称直接在根目录新增文件,随后写具体代码

注:每个最后暴露的封装方法需要与结构体WeixinApi绑定,比如:

```go
func (api *WexinApi) GetAccessToken(ctx context.Context) (string, error) {}
```

另外接口的访问需要传递accesstoken参数,通过上述方法即可获取,避免自己请求新的token

## 目前的功能

+ 自动接入 ☑️

+ 自定义菜单

+ 个性化菜单

+ 基础消息能力
  
  + 接收消息
  
  + 被动回复
  
  + 模板消息
  
  + 公众号订阅消息
  
  + 各种查询

+ 订阅通知

+ 客服消息

+ 微信网页开发
  
  + 网页授权

+ 素材管理

+ 用户管理

+ 账号管理

+ 数据统计

注:不同的公众号类型对应的接口权限不同,详见:[接口权限说明](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Explanation_of_interface_privileges.html)


