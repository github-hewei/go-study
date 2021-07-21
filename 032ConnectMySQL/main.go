package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("异常：", err)
		}
	}()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:13301)/test")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	rows, err := db.Query("show tables")
	if err != nil {
		panic(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Println(columns)

	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			panic(err)
		}
		fmt.Println(table)
	}

	//插入
	sql := "INSERT INTO table_json (content) VALUES (?)"
	str, err := json.Marshal([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	if err != nil {
		panic(err)
	}
	result, err := db.Exec(sql, str)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
