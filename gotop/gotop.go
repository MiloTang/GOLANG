package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
		"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
		"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
		"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
		"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}
	r          = rand.New(rand.NewSource(time.Now().UnixNano()))
	atagRegExp = regexp.MustCompile(`"\./buzz\?b=.*?"`)
	keyword    = regexp.MustCompile(`<a class="list-title".*>.*</a>`)
	replace    = regexp.MustCompile(`<a class="list-title".*?>`)
	head       = `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=gb2312">
    <title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
	<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
</head>
<body>
<div class="container">
<div class="row">
<div class="col-md-12">`
	foot = `</div>
</div>
</div>
</body>
</html> `
	allstr = ""
)

func init() {

}
func main() {
	Spy("http://top.baidu.com/boards?fr=topindex")
}
func Spy(url string) {
	resStr := GetBody(url)
	atag := atagRegExp.FindAllString(resStr, -1)
	subUrl := ""
	for _, v := range atag {
		subUrl = "http://top.baidu.com" + strings.Replace(v, `".`, "", -1)
		subUrl = strings.Replace(subUrl, `"`, "", -1)
		txt := keyword.FindAllString(GetBody(subUrl), -1)
		//
		for _, v1 := range txt {
			allstr = allstr + strings.Replace(replace.ReplaceAllString(v1, ""), "</a>", "", -1) + ","
		}

	}
	ioutil.WriteFile("top.html", []byte(head+allstr+foot), os.ModeAppend)
}
func GetBody(url string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	bodyStr := ""
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", GetRandomUserAgent())
	client := http.DefaultClient
	res, e := client.Do(req)
	if e != nil {
		fmt.Errorf("Get请求%s返回错误:%s", url, e)
		return bodyStr
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		bodyStr = string(bodyByte)
	}
	return bodyStr
}
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}
