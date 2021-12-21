package main

import "fmt"

type Personnel interface {
	Say(msg string)
	Eat(food string)
}

type Person struct {
	Name string
}

func (p *Person) Say(msg string) {
	fmt.Println(p.Name, "Say", msg)
}

func (p *Person) Eat(food string)  {
	fmt.Println(p.Name, "Eat", food)
}

func main() {

	p1 := Person{Name: "小明"}
	test1(&p1)
	test2(&p1)

	p2 := Person{Name: "小红"}
	test1(&p2)
	test2(&p2)

}

func test1(p Personnel) {
	p.Say("test1")
	p.Eat("test1")
}

func test2(p Personnel) {
	p.Say("test2")
	p.Eat("test2")
}
