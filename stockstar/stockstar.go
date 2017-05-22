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

	"github.com/axgle/mahonia"
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
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {

}
func main() {
	Spy("http://quote.stockstar.com//webhandler/futures.ashx")
}
func Spy(url string) {
	resStr := GetBody(url)
	fmt.Println(mahonia.NewEncoder("gbk").ConvertString(resStr))
	os.Exit(0)
	ioutil.WriteFile("stockstar", []byte(resStr), os.ModeAppend)
}
func GetBody(url string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	bodyStr := ""
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		bodyStr = err.Error()
	} else {
		req.Header.Set("User-Agent", GetRandomUserAgent())
		req.Header.Set("Content-Type", "application/json; charset=utf-8")

		//	reg.Header.Set(`HTTP`, `2.0`)
		//	reg.Header.Set(`Accept`, `*/*`)
		//	reg.Header.Set(`Accept-Language`, `zh-cn`)
		//	reg.Header.Set(`Host`, `itunes.apple.com`)
		//	reg.Header.Set(`Connection`, `keep-alive`)
		//	reg.Header.Set(`X-Apple-Store-Front`, `143465-19,21 t:native`)
		//	reg.Header.Set(`X-Dsid`, `932530590`)
		//	reg.Header.Set(`Cookie`, `xp_ci=3z1E7umazD0Dz5SBzCwNzB7weVKgD; s_vi=[CS]v1|29F61F868501299A-60000114E000452B[CE]; Pod=20; itspod=20; xt-src=b; xt-b-ts-932530590=1408324780262; mz_at_ssl-932530590=AwUAAAFRAAER1gAAAABT8vlrAo2EAZQvwAJjChIlGtIxIKYErLQ=; mz_at0-932530590=AwQAAAFRAAER1gAAAABT8VSrdHM0dXgdzosavj4+sT0AJfhYBx4=; wosid-lite=qQmZVeBH9vj91TakAeKEZg; ns-mzf-inst=35-163-80-118-68-8171-202429-20-nk11; X-Dsid=932530590`)
		if res, e := http.DefaultClient.Do(req); e != nil {
			bodyStr = e.Error()
		} else {
			if res.StatusCode == 200 {
				body := res.Body
				defer body.Close()
				bodyByte, _ := ioutil.ReadAll(body)
				bodyStr = string(bodyByte)
			}
		}
	}
	reg := regexp.MustCompile(`(hx_json31492672869171)|(\()|(\))|(;)`)
	bodyStr = reg.ReplaceAllString(bodyStr, "")
	strs := strings.Split(bodyStr, `,`)
	for k, v := range strs {

		fmt.Println(k, "==", v)
	}
	fmt.Println(bodyStr)
	return bodyStr
}
func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}
