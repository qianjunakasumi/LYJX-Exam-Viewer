package viewer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// refreshRes 刷新响应
type refreshRes struct {
	Message string `json:"msg"` // 消息
}

// Refresh 刷新
func (v *Viewer) Refresh() (err error) {

	resp, err := http.Post(
		"https://mic.fjjxhl.com/pcnews/index.php/Home/User/parlogin",
		"application/x-www-form-urlencoded; charset=UTF-8",
		strings.NewReader("Login_phone="+v.account+"&parpwd="+v.password),
	)
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}

	var status *refreshRes
	err = json.Unmarshal(res, &status)
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}

	if m := status.Message; m != "ok" {
		return errors.New("viewer: 登录失败：" + m)
	}

	token := resp.Header.Get("set-cookie")
	if !strings.Contains(token, "PHPSESSID=") {
		return errors.New("viewer: 获取 Token 失败")
	}
	v.token = token[10:36]

	return
}
