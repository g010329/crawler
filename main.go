package main

import (
	"fmt"
	"os"

	"gopkg.in/resty.v1"
)

type Job struct {
	Data *DataWrapper `json:"data"`
}

type DataWrapper struct {
	JobDetail *JobDetailWrapper `json:"jobDetail"`
}

type JobDetailWrapper struct {
	JobDescription string `json:"jobDescription"`
	Salary         string `json:"salary"`
	SalaryMin      int    `json:"salaryMin"`
	SalaryMax      int    `json:"salaryMax"`
}

func main() {
	f, err := os.Create("test.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var test *Job
	client := resty.New()
	resp, err := client.R().
		SetHeader("Referer", "https://www.104.com.tw/job/ajax/content/6hr41").
		SetResult(&test).
		Get("https://www.104.com.tw/job/ajax/content/6hr41")

	if err != nil {
		fmt.Println("  Error      :", err)
	}
	fmt.Println("  Body       :\n", resp)
	fmt.Println("  test       :\n", test.Data.JobDetail.JobDescription)
}

// resty
// https://www.it610.com/article/1409013028539695104.htm

// mahonia https://pkg.go.dev/github.com/axgle/mahonia
// \u9650 -> JS (ASCII) unicode
// https://www.it610.com/article/1409013028539695104.htm
// https://blog.csdn.net/zhizhengguan/article/details/104373271
// https://www.jianshu.com/p/9ceae300c78e

// Skill Set:
// ["react","redux","ES6"]
