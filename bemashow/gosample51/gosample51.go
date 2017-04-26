package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36",
		"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}
	head = `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=gb2312">
    <title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
	</head>
<body>`
	foot = `
</body>
</html> `
	r           = rand.New(rand.NewSource(time.Now().UnixNano()))
	RegExpUrl   = regexp.MustCompile(`/article/.*?htm`)
	RegExpText  = regexp.MustCompile(`<div id="content">[\s\S]*<div class="art_xg">`)
	RegExpTitle = regexp.MustCompile(`<h1 class="YaHei">.*?</h1>`)
	num         = 1
	count       = 0
)

func init() {

}
func main() {
	for {
		fmt.Println(count)
		if count == 11 {
			break
		} else {
			GetUrl("http://www.jb51.net", GetBody(Url()))
		}
	}

}
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}
func Url() string {
	count++
	if count < 11 {
		return "http://www.jb51.net/list/list_246_" + strconv.Itoa(count) + ".htm"
	}
	return ""
}
func GetUrl(pre, body string) {
	url := RegExpUrl.FindAllString(body, -1)
	for _, v := range url {
		GetText(pre + v)
	}
}

func GetText(url string) {
	body := GetBody(url)
	fmt.Println(body)
	os.Exit(0)
	txt := RegExpText.FindAllString(body, -1)
	title := RegExpTitle.FindAllString(body, -1)
	alltxt := ""
	for _, vt := range txt {
		alltxt = alltxt + vt
	}
	alltl := ""
	for _, vtl := range title {
		alltl = alltl + vtl
	}
	ioutil.WriteFile("jb51/"+strconv.Itoa(num)+"gosample.html", []byte(head+alltl+alltxt+foot), os.ModeAppend)
	num++
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
