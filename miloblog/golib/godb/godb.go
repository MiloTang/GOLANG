package godb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Mysql() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	query, err := db.Query("select Host,User from mysql.user")
	if err != nil {
		fmt.Println(err)
	}
	var host, user string
	for query.Next() {
		query.Scan(&host, &user)
		fmt.Println(host, user)
	}
	db.Close()
}
