package gomysql

import (
	"database/sql"
	"fmt"
	"golib/goconfig"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func Connect() {
	link := goconfig.Username + ":" + goconfig.Passwd + "@tcp(" + goconfig.Host + ")/" + goconfig.Database + "?" + goconfig.Charset
	db, err = sql.Open("mysql", link)
	if err != nil {
		fmt.Println(err)
	}
}
func Select(sqlstmt string, params ...interface{}) [][]string {
	stmt, e := db.Prepare(sqlstmt)
	if e != nil {
		fmt.Println("请检查sql语句->", sqlstmt, e)
	}
	rows, e := stmt.Query(params...)
	if e != nil {
		fmt.Println("请检查参数->", params, e)
	}
	cols, e := rows.Columns()
	if e != nil {
		panic(e)
	}
	lenght := len(cols)
	rawResult := make([][]byte, lenght)
	results := make([][]string, lenght)
	dest := make([]interface{}, lenght)
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			panic(err)
		}
		for i, raw := range rawResult {
			if raw == nil {
				results[i] = nil
			} else {
				results[i] = append(results[i], string(raw))
			}
		}
	}
	stmt.Close()
	return results
}
func Insert(sqlstmt string, params ...interface{}) {
	stmt, e := db.Prepare(sqlstmt)
	if e != nil {
		fmt.Println("请检查sql语句->", sqlstmt, e)
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		fmt.Println("数据插入失败->", err)
	}
	stmt.Close()
}
func Delete(sqlstmt string, params ...interface{}) {
	stmt, e := db.Prepare(sqlstmt)
	if e != nil {
		fmt.Println("请检查sql语句->", sqlstmt, e)
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		fmt.Println("数据删除失败->", err)
	}
	stmt.Close()
}
func Update(sqlstmt string, params ...interface{}) {
	stmt, e := db.Prepare(sqlstmt)
	if e != nil {
		fmt.Println("请检查sql语句->", sqlstmt, e)
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		fmt.Println("数据更新失败->", err)
	}
	stmt.Close()
}
func Close() {
	db.Close()
}
