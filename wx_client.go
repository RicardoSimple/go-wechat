package go_wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const contentType = "Content-Type"

type WxClient struct {
	ServerHost string `json:"server"`
	api        *WexinApi
}

// GetWithAccess 附带access的GET请求
func (c *WxClient) GetWithAccess(url string, headers map[string]string, result, failureResult interface{}) error {
	return c.Get(url, headers, result, failureResult, true)
}

// PostWithAccess 附带access的POST请求
func (c *WxClient) PostWithAccess(url string, headers map[string]string, data interface{}, result, failureResult interface{}) error {
	return c.Post(url, headers, data, result, failureResult, true)
}

func (c *WxClient) Get(url string, headers map[string]string, result, failureResult interface{}, needAccess bool) error {
	fullUrl := c.addHostBefore(url)
	if needAccess {
		newUrl, err := c.addAccessTokenAndHost(url)
		if err != nil {
			return fmt.Errorf("token error")
		}
		fullUrl = newUrl
	}
	c.api.Logger.Println("GET " + fullUrl)
	return do("GET", fullUrl, headers, nil, result, failureResult)
}

func (c *WxClient) Post(url string, headers map[string]string, data interface{}, result, failureResult interface{}, needAccess bool) error {
	fullUrl := c.addHostBefore(url)
	if needAccess {
		newUrl, err := c.addAccessTokenAndHost(url)
		if err != nil {
			return fmt.Errorf("token error")
		}
		fullUrl = newUrl
	}
	c.api.Logger.Println("POST " + fullUrl)
	return do("POST", fullUrl, headers, data, result, failureResult)
}

func (c *WxClient) addHostBefore(url string) string {
	return c.api.ServerHost + url
}
func (c *WxClient) addAccessTokenAndHost(url string) (newUrl string, err error) {
	// 自己追加accesstoken
	token, err := c.api.GetAccessToken()
	if err != nil {
		return
	}
	if strings.Contains(url, "?") {
		newUrl = url + "&access_token=" + token
		return
	} else {
		newUrl = url + "?access_token=" + token
		return
	}
}

// 通用请求函数
func do(method, _url string, headers map[string]string, data interface{}, result, failureResult interface{}) error {
	client := &http.Client{}
	var reqBody []byte
	var err error
	// 如果请求头为空 设置默认请求头
	if headers == nil {
		headers = map[string]string{contentType: "application/json"}
	}
	if data != nil {
		switch data.(type) {
		case url.Values:
			headers[contentType] = "application/x-www-form-urlencoded"
			formData, _ := data.(url.Values)
			reqBody = []byte(formData.Encode())
		default:
			headers[contentType] = "application/json"
			reqBody, err = json.Marshal(data)
			if err != nil {
				return err
			}
		}
	}
	request, err := http.NewRequest(method, _url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	teeReader := io.TeeReader(resp.Body, &buf)
	if resp.StatusCode >= http.StatusBadRequest {
		err = json.NewDecoder(teeReader).Decode(failureResult)
	} else {
		err = json.NewDecoder(teeReader).Decode(result)
	}
	if err != nil {
		errBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("format failure: %v", string(errBody))
	}
	return nil
}
