package viewer

// GetLatestDetail 获取最新考试详情
func (v *Viewer) GetLatestDetail() (d *Detail, err error) {

	h, err := v.GetExamHistory()
	if err != nil {
		return
	}

	d, err = h[0].GetDetail()
	if err != nil {
		return
	}

	return
}
