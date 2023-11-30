package conf

type MenuItem struct {
	Type    string     `json:"type"`
	Name    string     `json:"name"`
	URL     string     `json:"url"`
	SubMenu []MenuItem `json:"sub_button"`
}
type TopItem struct {
	Name    string     `json:"name"`
	SubMenu []MenuItem `json:"sub_button"`
}

type Menu struct {
	Buttons []TopItem `json:"button"`
}

type LangParam string

type TagList struct {
	Tags []Tag `json:"tags"`
}
type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
type OpenIdList struct {
	Openid []string `json:"openid"`
}
