package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

func main() {
	routesinit()
	fmt.Println(AllAuths)
	http.HandleFunc("/", mainHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var (
	// 定义路由配置
	// 由路径+空格+描述为键  访问方法为值
	routesconfig = map[string]func(rw http.ResponseWriter, req *http.Request){
		"/ 主页":      index,
		"/login 登陆": login,
		"/admin 管理": admin,
	}

	// 定义提供给权限配置页面使用的权限列表
	// 以访问的方法为键 备注为值
	AllAuths = map[string]string{}

	// 定义路由map
	routes = map[string]func(rw http.ResponseWriter, req *http.Request){}

	// 定义无需权限控制的页面
	noauthroutes = []string{"main.login"}
)

// 初始化路由 将路由配置MAP中的路径和描述分开并填充权限列表
func routesinit() {
	var keys []string
	var AllAuthsKey string
	for k, v := range routesconfig {
		keys = strings.Split(k, " ")
		AllAuthsKey = getfuncname(v)
		AllAuths[AllAuthsKey] = keys[1]
		routes[keys[0]] = v
	}
}

// 取得被调用方法名称 用于访问路径权限控制
func getfuncname(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// 路径权限控制
func mainHandle(rw http.ResponseWriter, req *http.Request) {
	if handler, ok := routes[req.URL.Path]; ok {
		// 取得正在访问的方法名
		callfuncname := getfuncname(handler)
		// 检查被访问路径是否不需要权限控制
		NoAccessControl := false
		for _, v := range noauthroutes {
			if callfuncname == v {
				NoAccessControl = true
				break
			}
		}
		if NoAccessControl {
			fmt.Println("guset")
			handler(rw, req)
			return
		}
		// 没有访问权限 返回 fmt.Fprintln(rw, "没有访问权限")
		handler(rw, req)
		return
	}
	fmt.Fprintln(rw, "请求页面不存在!")
}

func index(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "这是首页!")
}
func login(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "这是登陆页!")
}

func admin(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "这是管理页面!")
}
