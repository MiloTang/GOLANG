package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

//结构化存储 Api 返回的 JSON 数据，注意字段首字母必须大写
type szTong_s struct {
	ID            uint64 `json:"card_number"`
	Balance       string `json:"card_balance"`
	Balance_time  string `json:"balance_time"`
	Validity_time string `json:"card_validity"`
	Server_time   string `json:"current_time"`
}

//打印使用帮助
func prtHelp(exitCode int) {
	fmt.Println("\n查询深圳通的信息。")
	fmt.Printf("    Usage: %s <卡号>\n\n", filepath.Base(os.Args[0]))
	os.Exit(exitCode)
}

//检查用户输入的深圳通 ID 是否为 uint64， 如果不正确则直接退出
func check_szTongID() (ID string) {
	defer func() {
		//如果用户没有传入参数，即捕获 os.Args[1] 的索引错误
		if err := recover(); err != nil {
			prtHelp(1)
		}
	}()

	if len(os.Args) == 1 {
		prtHelp(0)
	} else {
		if _, err := strconv.ParseUint(os.Args[1], 10, 64); err != nil {
			prtHelp(1)
		}
	}
	return os.Args[1]
}

// 从 API 中获取深圳通信息(JSON)，并解析成 struct
func fetch_szTong(ID string) (szTong_ptr *szTong_s, err error) {
	var URL = "http://query.shenzhentong.com:8080/sztnet/qryCard.do?cardno=" + ID

	var szTong_b []byte
	if response, err := http.Get(URL); err != nil {
		return szTong_ptr, err
	} else {
		if szTong_b, err = ioutil.ReadAll(response.Body); err != nil {
			return szTong_ptr, err
		}
	}

	szTong_ptr = &szTong_s{}
	err = json.Unmarshal(szTong_b, szTong_ptr) // JSON to Struct
	return szTong_ptr, err
}

//格式化输出深圳通信息
func Print_szTong(szTong *szTong_s) {
	fmt.Println("\n深圳通查询结果：\n")

	if szTong.ID == 0 {
		fmt.Println("    <卡号无效>\n")
	} else {
		fmt.Printf("    %-20s : %d\n", "Card NO.", szTong.ID)
		fmt.Printf("    %-20s : %s\n", "Validity date", szTong.Validity_time)
		fmt.Printf("    %-20s : %s\n", "Balance", szTong.Balance)
		fmt.Printf("    %-20s : %s\n", "Last balance time", szTong.Balance_time)
		fmt.Printf("    %-20s : %s\n", "Server current time", szTong.Server_time)
	}
}

func main() {
	szTongID := check_szTongID()

	if szTong, err := fetch_szTong(szTongID); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		Print_szTong(szTong)
	}
}
