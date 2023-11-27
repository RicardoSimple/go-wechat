package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const contentType = "Content-Type"

func Get(ctx context.Context, url string, headers map[string]string, result, failureResult interface{}) error {
	return do(ctx, "GET", url, headers, nil, result, failureResult)
}

func Post(ctx context.Context, url string, headers map[string]string, data interface{}, result, failureResult interface{}) error {
	return do(ctx, "POST", url, headers, data, result, failureResult)
}

func do(ctx context.Context, method, _url string, headers map[string]string, data interface{}, result, failureResult interface{}) error {
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
