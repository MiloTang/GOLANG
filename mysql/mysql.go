package main

import (
	"database/sql"
	"fmt"
	"mysql/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err error
	db  *sql.DB
)

type data struct {
	data interface{}
}

func Connect() error {
	link := config.Username + ":" + config.Passwd + "@tcp(" + config.Ip + ":" + config.Port + ")/" + config.Database + "?charset=" + config.Charset
	db, err = sql.Open("mysql", link)
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	db.Close()
}
func Select(sqlstmt string, params ...interface{}) {
	stmt, e := db.Prepare(sqlstmt)
	if err != nil {
		panic(e)
	}
	rows, e := stmt.Query(params...)
	if err != nil {
		panic(e)
	}
	if rows == nil {
		fmt.Println(" ")
	}
	cols, e := rows.Columns()
	if err != nil {
		panic(e)
	}
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}
	var results [][]string
	for rows.Next() {
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
		results = append(results, result)
	}
	fmt.Println(results)
}
func GetValue(rows *sql.Rows) {

}
func Insert() {

}
func check_convert(ok bool) {
	if false == ok {
		panic("type cast failed!")
	}
}
func main() {
	Connect()
	defer db.Close()
	Select("select * from guests where (username=? and age=?)", "francis", 28)

}
