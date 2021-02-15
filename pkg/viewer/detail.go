package viewer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"
)

type (
	// History 历史
	History struct {
		super *Viewer

		id   string
		Name string
		Date time.Time
	}

	// scoreRes 分数响应
	scoreRes struct {
		Subjects []struct {
			Subject      string `json:"subject"`
			Score        string `json:"score"`
			AverageScore string `json:"x_average"`
			MaxScore     string `json:"x_max"`
		} `json:"geke"`
	}

	// Detail 详情
	Detail struct {
		TotalScore   float64 `json:"my_allscore"`
		MaxScore     float64 `json:"zuigaofen"`
		AverageScore float64 `json:"zongfenpingjunfen"`

		Subjects map[string]Subject
	}

	// Subject 学科
	Subject struct {
		Score        float64
		AverageScore float64
		MaxScore     float64
	}
)

// GetDetail 获取详情
func (h *History) GetDetail() (d *Detail, err error) {

	values, err := url.ParseQuery("item_id=" + h.id + "&leixing=x")
	if err != nil {
		return
	}

	resp, err := h.super.postWithToken("https://mic.fjjxhl.com/Jx/index.php/Home/Newscore/ajaxxqscore", values)
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

	var subjects scoreRes
	err = json.Unmarshal(res, &d)
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}
	err = json.Unmarshal(res, &subjects)
	if err != nil {
		err = fmt.Errorf("viewer: %s", err)
		return
	}

	d.Subjects = map[string]Subject{}
	for _, v := range subjects.Subjects {
		s, err := strconv.ParseFloat(v.Score, 64)
		if err != nil {
			return nil, fmt.Errorf("viewer: %s", err)
		}

		a, err := strconv.ParseFloat(v.AverageScore, 64)
		if err != nil {
			return nil, fmt.Errorf("viewer: %s", err)
		}

		m, err := strconv.ParseFloat(v.MaxScore, 64)
		if err != nil {
			return nil, fmt.Errorf("viewer: %s", err)
		}

		d.Subjects[v.Subject] = Subject{s, a, m}
	}

	return
}
