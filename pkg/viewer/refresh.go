package viewer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// refreshRes 刷新响应
type refreshRes struct {
	Message string `json:"msg"`
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

// switchStudentsProfile 切换学生档案。
// 当账号下只有一个学生档案时无需切换。
// 若存在多个学生档案 且 获取的成绩信息不是您本人的情况下
// 可以使用切换档案切换为您的个人档案
func (v Viewer) switchStudentsProfile() (err error) {

	values, err := url.ParseQuery("num=" + v.number + "&name=" + v.name)
	if err != nil {
		return
	}

	resp, err := v.postWithToken("https://mic.fjjxhl.com/Jx/index.php/Home/User/switchStudent", values)
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

	if string(res) != "1" {
		return fmt.Errorf("viewer: 切换学生档案失败")
	}

	return
}
