package viewer

import "time"

// History 历史
type History struct {
	id   string    // 项目标识号
	Name string    // 项目
	Date time.Time // 时间
}
