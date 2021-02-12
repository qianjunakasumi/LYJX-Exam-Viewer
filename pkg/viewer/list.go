package viewer

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"time"
)

// History 历史
type History struct {
	ID   string    // 项目标识号
	Name string    // 项目
	Date time.Time // 时间
}

// GetExamHistory 获取历史考试
func (v *Viewer) GetExamHistory() (h []*History, err error) {

	res, err := v.getWithToken("https://mic.fjjxhl.com/42baobaobanpai/index.php/Admin/Zzy/lately")
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}

	h = []*History{}
	doc.Find("div.some-right").Each(func(_ int, sel *goquery.Selection) {

		hi := &History{}

		sel.Find("div.right-up").Each(func(i int, selection *goquery.Selection) {
			c := selection.Children()
			hi.Name = c.Eq(1).Text()
			t, err := time.Parse("2006-01-02", c.Eq(0).Text())
			if err != nil {
				err = errors.New("viewer: 解析时间失败")
				return
			}

			hi.Date = t
		})

		i, ok := sel.Find("div.right-btn.orow.tea.stzd").Attr("data")
		if !ok {
			err = errors.New("viewer: 获取考试 ID 失败")
			return
		}

		hi.ID = i
		h = append(h, hi)
	})

	return
}
