package main

import "fmt"

type Person struct {
	name string
	age  uint8
	sex  byte
}

//引用传递的话，可以避免copy整个结构体数据
func (p *Person) say(s string) {
	fmt.Println(p, &p, &p.name)
	fmt.Println(p.name, "say: ", s)
}

func (p Person) say2(s string) {
	fmt.Println(p, &p.name)
	fmt.Println(p.name, "say: ", s)
}

func main() {

	var v1 = []byte("abcd")
	fmt.Println(v1)

	p1 := Person{"tom", 20, 1}
	fmt.Println(&p1.name, &p1)
	p1.say("Hello World")
	p1.say2("Hello World")

}
