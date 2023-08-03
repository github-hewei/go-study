package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:13301)/lmktest")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(50) // 长连接数
	db.SetMaxIdleConns(10) // 闲置连接数

	if err = db.Ping(); err != nil {
		panic(err)
	}

	sqlx := "SELECT `shopid`, `name` FROM `oa_wechat_com`"
	rows, err := db.Query(sqlx)
	if err != nil {
		panic(err)
	}

	comMap := make(map[int]string, 0)
	for rows.Next() {
		var (
			shopid int
			name string
		)

		err = rows.Scan(&shopid, &name)
		if err != nil {
			panic(err)
		}
		comMap[shopid] = name
	}

	if val, ok := comMap[13]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Not Exists...")
	}

}
