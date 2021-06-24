package main

//在引入的包之前加下划线表示，仅匿名导入包，不需要包里边的数据或方法
//在导入包的时候会触发包中的init初始化函数

import (
	"fmt"

	_ "github.com/github-hewei/go-study/038Package01/p2"
)

func main() {
	fmt.Println("Main")

}
