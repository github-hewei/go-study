package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Goods struct {
	id         int
	goods_name string
	goods_spec string
	producer   string
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:13301)/lmkmain")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	data1, err := fetchGoodsById(db, 10000001)
	if err != nil {
		panic(err)
	}

	data2, err := fetchGoodsById(db, 10000002)
	if err != nil {
		panic(err)
	}

	data3, err := fetchGoodsById(db, 10000003)
	if err != nil {
		panic(err)
	}

	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

}

func fetchGoodsById(db *sql.DB, id int) (d Goods, err error) {
	time.Sleep(time.Duration(time.Second))

	sql := "select id, goods_name, goods_spec, producer from `oa_chain_goods` where id = " + strconv.Itoa(id)

	err = db.QueryRow(sql).Scan(&d.id, &d.goods_name, &d.goods_spec, &d.producer)
	if err != nil {
		return
	}

	return
}
