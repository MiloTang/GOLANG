package main

import (
	"bufio"
	"container/list"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"websocket"
)

var connid int32
var conns *list.List

func ChatroomServer(ws *websocket.Conn) {
	defer ws.Close()
	if connid == 2147483647 {
		connid = 1
	} else {
		connid++
	}
	id := connid
	fmt.Printf("connection id: %d\n", id)
	item := conns.PushBack(ws)
	defer conns.Remove(item)
	name := fmt.Sprintf("游客%d", id)
	SendMessage(nil, fmt.Sprintf("欢迎 %s 加入\n", name))
	r := bufio.NewReader(ws)
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Printf("disconnected id: %d\n", id)
			SendMessage(item, fmt.Sprintf("%s offline\n", name))
			break
		}
		fmt.Printf("%s: %s", name, data)
		SendMessage(item, fmt.Sprintf("%s\t>>> %s", name, data))
	}
}
func SendMessage(self *list.Element, data string) {
	for item := conns.Front(); item != nil; item = item.Next() {
		ws, ok := item.Value.(*websocket.Conn)
		if !ok {
			panic("item not *websocket.Conn")
		}
		if item == self {
			continue
		}
		io.WriteString(ws, data)
	}
}

// 网页客户端
func Client(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chat.html")
	t.Execute(w, nil)
}
func main() {
	fmt.Printf("服务器启动")
	connid = 0
	conns = list.New()
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))
	http.Handle("/chatroom", websocket.Handler(ChatroomServer))
	http.HandleFunc("/", Client)
	err := http.ListenAndServe(":6611", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
