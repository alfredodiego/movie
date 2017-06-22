package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sort"
	"strconv"
)

type MovieItem struct {
	Rate              string
	Cover_x           int
	Title             string
	Is_beetle_subject bool
	Url               string
	Playable          bool
	Cover             string
	Id                string
	Cover_y           int
	Is_new            bool
}

type mitem []MovieItem

func (m mitem)Len() int {
	return len(m)
}
func (m mitem)Less(i, j int) bool {
	irate, error1 := strconv.ParseFloat(m[i].Rate, 64)
	jrate, error2 := strconv.ParseFloat(m[j].Rate, 64)
	if error1 != nil {
		fmt.Println("类型转换错误:", m[i])
	}
	if error2 != nil {
		fmt.Println("类型转换错误:", m[j])
	}
	return irate > jrate
}
func (m mitem)Swap(i, j int) {
	temp := m[i]
	m[i] = m[j]
	m[j] = temp
}

type Sub struct {
	Subjects mitem
}

func main() {
	urlstr := "https://movie.douban.com/j/search_subjects?type=movie&tag=%E5%8A%A8%E4%BD%9C&sort=time&playable=on&page_limit=100&page_start=0"
	response, _ := http.Get(urlstr)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	movieliststr := []byte(body)

	var sub Sub
	error := json.Unmarshal(movieliststr, &sub)
	if error != nil {
		fmt.Println("error:", error)
	}

	mitem := sub.Subjects
	sort.Sort(mitem)
	for _, v := range mitem {
		fmt.Println(v.Rate, v.Title, v.Url)
	}
	fmt.Println("输入q退出")
	fmt.Scanf("q")
}