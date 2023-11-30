package conf

const (
	// ServerHostGeneral 通用域名
	ServerHostGeneral = "https://api.weixin.qq.com"
	//ServerHostGeneralRDR 通用异地容灾域名
	ServerHostGeneralRDR = "https://api2.weixin.qq.com"
	// ServerHostShenZhen 深圳域名
	ServerHostShenZhen = "https://sz.api.weixin.qq.com"
	// ServerHostShangHai 上海域名
	ServerHostShangHai = "https://sh.api.weixin.qq.com"
	// ServerHostHongKong 香港域名
	ServerHostHongKong = "https://hk.api.weixin.qq.com"
)

const (
	DEFAULT LangParam = ""
	ZH_CN   LangParam = "zh_CN"
	ZH_TW   LangParam = "zh_TW"
	EN      LangParam = "en"
)
