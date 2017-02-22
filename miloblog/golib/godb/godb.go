package godb

import (
	"database/sql"
	"fmt"
	"reflect"
)

func Mysql() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	query, err := db.Query("select * from mysql.user")
	if err != nil {
		fmt.Println(err)
	}
	v := reflect.ValueOf(query)
	fmt.Println(v)
	db.Close()
}
