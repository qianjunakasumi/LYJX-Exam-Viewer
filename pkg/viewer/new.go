package viewer

// Viewer 查看器
type Viewer struct {
	account  string // 账号
	password string // 密码

	token string // 令牌
}

// New 返回 Viewer 查看器
func New(acc string, pwd string) (v *Viewer, err error) {
	v = &Viewer{account: acc, password: pwd}
	err = v.Refresh()
	return
}
