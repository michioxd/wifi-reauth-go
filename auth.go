package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func CheckAuth(ctx context.Context) bool {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://free.wi-mesh.vn/status", nil)
	if err != nil {
		fmt.Println(err)
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if strings.Contains(string(body), "KẾT NỐI THÀNH CÔNG!") {
		return true
	}
	return false
}

func Logout(ctx context.Context) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://free.wi-mesh.vn/logout", nil)

	http.DefaultClient.Do(req)
}

func Login(ctx context.Context) (bool, string) {
	formData := url.Values{
		"dst":      {""},
		"dst2":     {""},
		"popup":    {"true"},
		"username": {"awing60"},
		"password": {"Awing60@2018"},
	}

	headers := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Encoding":           "gzip, deflate",
		"Accept-Language":           "en-VN,en;q=0.9",
		"Cache-Control":             "max-age=0",
		"Connection":                "keep-alive",
		"Content-Type":              "application/x-www-form-urlencoded",
		"Cookie":                    "exRun=EX_LOGIN;",
		"Origin":                    "http://free.wi-mesh.vn",
		"Referer":                   "http://free.wi-mesh.vn/login",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
	}
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://free.wi-mesh.vn/login", strings.NewReader(formData.Encode()))
	if err != nil {
		return false, err.Error()
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err.Error()
	}

	if strings.Contains(string(body), "Bạn Đã Đăng Nhập Thành Công") {
		return true, ""
	}

	return false, "Authentication failed"
}
