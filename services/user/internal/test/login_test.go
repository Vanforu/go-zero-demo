package test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	url := "http://127.0.0.1:8888/v1/login"
	data := bytes.NewReader([]byte(`{"phone": "110", "password": "123456"}`))
	// 创建POST请求
	resp, err := http.Post(url, "application/json", data)
	if err != nil {
		t.Errorf("接口调取失败，报错信息为: %s", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("接口调取失败，报错信息为:  %s", err)
		return
	}
	t.Logf("http post succeed:  %s", string(respBody))
}
