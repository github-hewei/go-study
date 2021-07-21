package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select * from oa_chain_goods")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		goods_name := ""
		err := rows.Scan(&goods_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(goods_name)
	}

}
