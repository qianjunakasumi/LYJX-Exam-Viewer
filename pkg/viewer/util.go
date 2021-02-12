package viewer

import "net/http"

// getWithToken 带 Token 的 GET 请求
func (v *Viewer) getWithToken(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("cookie", "PHPSESSID="+v.token)

	return http.DefaultClient.Do(req)
}
