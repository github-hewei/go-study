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

	ch := make(chan Goods, 3)

	go fetchGoodsById(ch, db, 10000001)
	go fetchGoodsById(ch, db, 10000002)
	go fetchGoodsById(ch, db, 10000003)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

func fetchGoodsById(ch chan Goods, db *sql.DB, id int) {
	time.Sleep(time.Duration(time.Second))
	sql := "select id, goods_name, goods_spec, producer from `oa_chain_goods` where id = " + strconv.Itoa(id)
	d := Goods{}
	err := db.QueryRow(sql).Scan(&d.id, &d.goods_name, &d.goods_spec, &d.producer)
	if err != nil {
		panic(err)
	}
	ch <- d
}
