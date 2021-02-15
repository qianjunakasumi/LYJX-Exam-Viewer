package viewer

// Viewer 查看器
type Viewer struct {
	account  string // 账号
	password string // 密码
	number   string // 学籍号
	name     string // 姓名

	token string // 令牌
}

// New 返回 Viewer 查看器。
// acc：账号；pwd：密码；num：学籍号；name：姓名
func New(acc, pwd, num, name string) (*Viewer, error) {
	v := &Viewer{account: acc, password: pwd, number: num, name: name}
	return v, v.Refresh()
}
