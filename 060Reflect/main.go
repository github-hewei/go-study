package main

import (
	"fmt"
	"reflect"
)

func main() {

	c := make(chan int, 1)
	dump(1, "abc", []int{1, 2, 3}, map[string]int{"aaa": 1, "bbb": 2}, func() {}, c)

}

func dump(args ...interface{}) {
	for _, val := range args {
		rv := reflect.ValueOf(val)
		switch rv.Kind() {
		case reflect.Int:
			fmt.Println("Int", val)
		case reflect.String:
			fmt.Println("String", val)
		case reflect.Slice:
			fmt.Println("Slice", val)
		case reflect.Map:
			fmt.Println("Map", val)
		case reflect.Func:
			fmt.Println("Func", val)
		case reflect.Struct:
			fmt.Println("Struct", val)
		case reflect.Bool:
			fmt.Println("Bool", val)
		default:
			fmt.Println("Unknown", val)
		}
	}
	fmt.Println()

	for _, val := range args {
		fmt.Println(reflect.TypeOf(val).Kind(), val)
	}
}

