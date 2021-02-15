package main

import (
	"log"
	"os"

	"github.com/qianjunakasumi/LYJX-Exam-Viewer/pkg/viewer"
)

// main 获取最新的考试详情
func main() {

	v, err := viewer.New(os.Getenv("acc"), os.Getenv("pwd"), os.Getenv("num"), os.Getenv("name"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("登录成功")

	/*
		err = v.SwitchStudentsProfile()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("切换学生档案成功")
	*/

	d, err := v.GetLatestDetail()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(d)
}
