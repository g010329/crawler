package main

import (
	"encoding/csv"
	"fmt"

	// "fmt"
	"os"
	"strings"

	"gopkg.in/resty.v1"
)

// 搜尋工作關鍵字結果list

type SearchResult struct {
	Data *DataWrapper `json:"data"`
}

type DataWrapper struct {
	List      []*JobRow `json:"list"`
	TotalPage int       `json:"totalPage"`
	PageNo    int       `json:"pageNo"`
}

type JobRow struct {
	JobName     string `json:"jobName"`
	CustName    string `json:"custName"`
	Description string `json:"description"`
	SalaryDesc  string `json:"salaryDesc"`
	SalaryLow   int    `json:"salaryLow"`
	SalaryHigh  int    `json:"salaryHigh"`
	// jobNameRaw
	// custNameRaw
	// remoteWorkType
	// optionEdu
	// jobAddrNoDesc
	// periodDesc
}

// 單一工作結果
// type Job struct {
// 	Data *DataWrapper `json:"data"`
// }

// type DataWrapper struct {
// 	JobDetail *JobDetailWrapper `json:"jobDetail"`
// 	Header    *HeaderWrapper    `json:"header"`
// }

// type JobDetailWrapper struct {
// 	JobDescription string `json:"jobDescription"`
// 	Salary         string `json:"salary"`
// 	SalaryMin      int    `json:"salaryMin"`
// 	SalaryMax      int    `json:"salaryMax"`
// }

// type HeaderWrapper struct {
// 	JobName  string `json:"jobName"`
// 	CustName string `json:"custName"`
// }

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

	f, err := os.Create("example2.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// var data *Job
	// client := resty.New()
	// client.R().
	// 	SetHeader("Referer", "https://www.104.com.tw/job/ajax/content/6hr41").
	// 	SetResult(&data).
	// 	Get("https://www.104.com.tw/job/ajax/content/6hr41")

	// isCompleteMatch1, matches1, test := checkSubstrings(data.Data.JobDetail.JobDescription, "react", "redux", "ES6", "typescript", "UI/UX")
	// fmt.Printf("Test 1: { isCompleteMatch: %t, Matches: %d }\n", isCompleteMatch1, matches1)
	// fmt.Println("命中skill set: ", test)

	// f.WriteString("\xEF\xBB\xBF")
	// w := csv.NewWriter(f)
	// csvData := [][]string{
	// 	{data.Data.Header.JobName, data.Data.Header.CustName, data.Data.JobDetail.Salary, strings.Join(test, ", ")},
	// }
	// w.WriteAll(csvData)
	// w.Flush()

	// ----------------------------------------------------------------
	var data *SearchResult

	client := resty.New()
	client.R().
		SetHeader("Referer", "https://www.104.com.tw/jobs/search/").
		SetQueryParams(map[string]string{
			"ro":            "0",
			"kwop":          "7",
			"keyword":       "前端",
			"expansionType": "area,spec,com,job,wf,wktm",
			"asc":           "0",
			"mode":          "s", // s or l
			"jobsource":     "2018indexpoc",
			"page":          "0",
		}).
		SetResult(&data).
		Get("https://www.104.com.tw/jobs/search/list")

	// totalPageNumber := data.Data.TotalPage

	// fmt.Printf("TotalPage: %d \n", totalPageNumber)

	f.WriteString("\xEF\xBB\xBF")
	var csvData [][]string

	for i := 0; i < len(data.Data.List); i++ {

		isCompleteMatch1, matches1, matchWords := checkSubstrings(data.Data.List[i].Description, "react", "redux", "ES6", "typescript", "UI/UX", "CSS", "Vue.js", "Vue", "Javascript", "RWD", "Git")
		fmt.Printf("Test 1: { isCompleteMatch: %t, Matches: %d, word: %s }\n", isCompleteMatch1, matches1, strings.Join(matchWords, ", "))

		csvData = append(csvData, []string{data.Data.List[i].JobName, data.Data.List[i].CustName, data.Data.List[i].SalaryDesc, strings.Join(matchWords, ", ")})
		// csvData = append(csvData, []string{data.Data.List[i].JobName, data.Data.List[i].CustName, data.Data.List[i].SalaryDesc})
	}

	w := csv.NewWriter(f)
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

// https://www.104.com.tw/jobs/search/list?ro=0&kwop:=7&keyword=前端&expansionType=area,spec,com,job,wf,wktm&order=1&asc=0&page=1&mode=s&jobsource=2018indexpoc
// ro: 0(不限) / 1(全職)
// keyword:"前端"
// area:'6001001000', # 限定在台北的工作 不限不填
// 'isnew':'30', # 只要最近一個月有更新的過的職缺
