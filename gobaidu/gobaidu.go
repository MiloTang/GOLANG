package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
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
	urlChannel = make(chan string, 200)
	atagRegExp = regexp.MustCompile(`<div class="spider">.*`)
	num        = 31777
	head       = `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
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
)

func init() {

}
func main() {
	go Spy("http://www.wenxue360.com/gushiwen/31777.html")
	for url := range urlChannel {
		fmt.Println("routines num = ", runtime.NumGoroutine(), "chan len = ", len(urlChannel))
		go Spy(url)
	}
	fmt.Println("a")

}
func Spy(url string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", GetRandomUserAgent())
	client := http.DefaultClient
	res, e := client.Do(req)
	if e != nil {
		fmt.Errorf("Get请求%s返回错误:%s", url, e)
		return
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)
		atag := atagRegExp.FindAllString(resStr, -1)
		for _, v := range atag {
			content := strings.Replace(strings.Replace(v, `<div class="spider">`, "", -1), "www.wenxue360.com", "www.mogupao.com", -1)
			if strings.Count(content, "")-1 > 100 {
				ioutil.WriteFile("gushiwen/"+strconv.Itoa(num)+".html", []byte(head+content+foot), os.ModeAppend)
			}
		}

		if num < 73000 {
			num++
			url := "http://www.wenxue360.com/gushiwen/" + strconv.Itoa(num) + ".html"
			urlChannel <- url
		} else {
			os.Exit(0)
		}

		//		os.Exit(0)
		//		for _, a := range atag {
		//			href, _ := GetHref(a)
		//			if strings.Contains(href, "article/details/") {
		//				fmt.Println("☆", href)
		//			} else {
		//				fmt.Println("□", href)
		//			}
		//			urlChannel <- href
		//		}
	}
}
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}
func GetHref(atag string) (href, content string) {
	inputReader := strings.NewReader(atag)
	decoder := xml.NewDecoder(inputReader)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				if strings.EqualFold(attrName, "href") || strings.EqualFold(attrName, "HREF") {
					href = attrValue
				}
			}
		// 处理元素结束（标签）
		case xml.EndElement:
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content = string([]byte(token))
		default:
			href = ""
			content = ""
		}
	}
	return href, content
}
