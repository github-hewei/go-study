package main

import "fmt"

type S1 struct {
	id    int
	name  string
	price float64
}

func main() {

	// %v 值的默认格式表示
	fmt.Printf("%v\n", []int{1, 2, 3})

	// %+v 类似于 %v , 在输出结构体时会输出字段名
	fmt.Printf("%v\n", S1{1, "tom", 200.00})
	fmt.Printf("%+v\n", S1{1, "tom", 200.00})

	// %#v 值的Go语法表示
	fmt.Printf("%#v\n", S1{1, "jerry", 300.00})

	// %T 值的类型的Go语法表示
	fmt.Printf("%T\n", S1{1, "tom", 100})

	// %% 百分号
	fmt.Printf("%v%%\n", 100)

	// %t 布尔值 true 或 false
	fmt.Printf("%t\t%t\n", true, false)

	// %b 整数二进制
	fmt.Printf("%b\t%b\n", 7, 4)

	// %c 该值对应的unicode码值
	fmt.Printf("%c %c %c\n", 97, 98, 99)

	// %d 整数十进制 (后面0开头的数字Go默认它为8进制数字)
	fmt.Printf("%d\n", 0017) // 15

	// %o 整数八进制
	fmt.Printf("%o\n", 15) // 17

	// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("%q %q %q\n", 100, 20, 017)

	// %x	表示为十六进制，使用a-f
	fmt.Printf("%x %x\n", 10, 15)

	// %X	表示为十六进制，使用A-F
	fmt.Printf("%X %X %X\n", 10, 15, 26)

	// %U	表示为Unicode格式：U+1234，等价于"U+%04X"
	fmt.Printf("%U\n", 100)

	// 浮点数与复数的两个组分：
	// %b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	// %e	科学计数法，如-1234.456e+78
	// %E	科学计数法，如-1234.456E+78
	// %f	有小数部分但无指数部分，如123.456
	// %F	等价于%f
	// %g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	// %G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	// 字符串和[]byte：
	// %s	直接输出字符串或者[]byte
	fmt.Printf("%s %s\n", "Hello", []byte{'W', 'o', 'r', 'l', 'd'})

	// %q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	fmt.Printf("%q\n", "")

	// %x	每个字节用两字符十六进制数表示（使用a-f）
	// %X	每个字节用两字符十六进制数表示（使用A-F）

	// 指针：
	// %p	表示为十六进制，并加上前导的0x
	var a = new(int)
	fmt.Printf("%p\n", a)
	b := 100
	fmt.Printf("%p\n", &b)

	// 指定宽度
	// %f:    默认宽度，默认精度
	// %9f    宽度9，默认精度
	// %.2f   默认宽度，精度2
	// %9.2f  宽度9，精度2
	// %9.f   宽度9，精度0
	fmt.Printf("%f\n", 0.101)
	fmt.Printf("%.2f\n", 0.009)
	fmt.Printf("%.2f\n", 100000.009)

	// 自定义索引
	s1 := S1{100, "小明", 998.888}
	fmt.Printf("Price: %.2[3]f\tID: %[1]d\tName: %[2]s\n", s1.id, s1.name, s1.price)

}
