package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var query = "README.md"
var matches int
//var ds = string(os.PathSeparator)

// 递归查找指定的文件
// 耗时：time 430.2471ms
// 文件数：matches 1475

func main() {
	start := time.Now()
	search("d:/temp")
	fmt.Println("time", time.Since(start))
	fmt.Println("matches", matches)
}

func search(p string) {
	files, err := ioutil.ReadDir(p)
	if err == nil {
		for _, file := range files {
			if file.Name() == query {
				// fmt.Println(p + "/" + file.Name())
				matches++
			}
			if file.IsDir() {
				search(p + "/" + file.Name() + "/")
			}
		}
	}
}
