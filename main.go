package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"gopkg.in/resty.v1"
)

type Job struct {
	Data *DataWrapper `json:"data"`
}

type DataWrapper struct {
	JobDetail *JobDetailWrapper `json:"jobDetail"`
	Header    *HeaderWrapper    `json:"header"`
}

type JobDetailWrapper struct {
	JobDescription string `json:"jobDescription"`
	Salary         string `json:"salary"`
	SalaryMin      int    `json:"salaryMin"`
	SalaryMax      int    `json:"salaryMax"`
}

type HeaderWrapper struct {
	JobName  string `json:"jobName"`
	CustName string `json:"custName"`
}

func checkSubstrings(str string, subs ...string) (bool, int, []string) {

	matches := 0
	matchWord := []string{}
	isCompleteMatch := true

	// fmt.Printf("String: \"%s\", Substrings: %s\n", str, subs)

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matchWord = append(matchWord, sub)
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches, matchWord
}

func main() {

	f, err := os.Create("example.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data *Job
	client := resty.New()
	client.R().
		SetHeader("Referer", "https://www.104.com.tw/job/ajax/content/6hr41").
		SetResult(&data).
		Get("https://www.104.com.tw/job/ajax/content/6hr41")

	// if err != nil {
	// 	fmt.Println("  Error      :", err)
	// }
	// fmt.Println("  Body       :\n", resp)
	// fmt.Println("  data       :\n", data.Data.Header.JobName)

	isCompleteMatch1, matches1, test := checkSubstrings(data.Data.JobDetail.JobDescription, "react", "redux", "ES6", "typescript", "UI/UX")
	fmt.Printf("Test 1: { isCompleteMatch: %t, Matches: %d }\n", isCompleteMatch1, matches1)
	fmt.Println(test)

	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	csvData := [][]string{
		{data.Data.Header.JobName, data.Data.Header.CustName, data.Data.JobDetail.Salary, strings.Join(test, ", ")},
	}
	w.WriteAll(csvData)
	w.Flush()

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

// data studio

// 搜尋入口->搜尋結果的api url->爬取所有資料->比對skill set->上lambda-->經過處理，大家各自的table變成大table
