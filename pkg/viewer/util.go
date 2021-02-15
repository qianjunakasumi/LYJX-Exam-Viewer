package viewer

import (
	"net/http"
	"net/url"
	"strings"
)

// getWithToken 带 Token 的 GET 请求
func (v *Viewer) getWithToken(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("cookie", "PHPSESSID="+v.token)

	return http.DefaultClient.Do(req)
}

// postWithToken 带 Token 的 POST 请求
func (v *Viewer) postWithToken(url string, data url.Values) (*http.Response, error) {

	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cookie", "PHPSESSID="+v.token)

	return http.DefaultClient.Do(req)
}
