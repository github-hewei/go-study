package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// json数据类型对应go数据类型
	// Go: bool 							-> Json: booleans
	// Go: float64 							-> Json: numbers
	// Go: string 							-> Json: strings
	// Go: []interface{} 					-> Json: arrays
	// Go: map[string]interface{} 			-> Json: objects
	// Go: nil 								-> Json: null

	// 常用方法
	// json.Marshal: 将数据转换成json字符串
	// json.Unmarshal: 将json字符串解析成可用数据

	// 普通Json
	demoOne()
	fmt.Println("---------------------------")

	// Json内嵌普通Json
	demoTwo()
	fmt.Println("---------------------------")

	// Json内嵌数组Json
	demoThree()
	fmt.Println("---------------------------")

	// Json内嵌动态Key的Json
	demoFour()
}

func demoOne() {
	type Actress struct {
		Name       string
		Birthday   string
		BirthPlace string
		Opus       []string
	}

	jsonData := []byte(`{
		"name": "小红",
		"birthday": "1999-01-01",
		"birthPlace": "北京市",
		"opus": [
			"《作品一》",
			"《作品二》",
			"《作品三》"
		]
	}`)

	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("出生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for _, val := range actress.Opus {
		fmt.Println("\t", val)
	}
}

func demoTwo() {

	type Opus struct {
		Date  string
		Title string
	}

	type Actress struct {
		Name       string
		Birthday   string
		BirthPlace string
		Opus       Opus
	}

	jsonData := []byte(`{
		"name": "小红",
		"birthday": "1990-02-02",
		"birthPlace": "北京市",
		"opus": {
			"date": "2020-02",
			"title": "《2020》"
		}
	}`)

	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("出生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	fmt.Printf("\t%s:%s\n", actress.Opus.Date, actress.Opus.Title)

}

func demoThree() {

	type Opus struct {
		Date  string
		Title string
	}

	type Actress struct {
		Name       string
		Birthday   string
		BirthPlace string
		Opus       []Opus
	}

	jsonData := []byte(`{
		"name": "小红",
		"birthday": "1990-02-02",
		"birthPlace": "北京市",
		"opus": [
			{
				"date": "2020-02",
				"title": "《2020》"
			},
			{
				"date": "2020-02",
				"title": "《2020》"
			}
		]
	}`)

	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("出生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for _, val := range actress.Opus {
		fmt.Printf("\t%s:%s\n", val.Date, val.Title)
	}

}

func demoFour() {
	type Opus struct {
		Type  string
		Title string
	}

	type Actress struct {
		Name       string
		Birthday   string
		BirthPlace string
		Opus       map[string]Opus
	}

	jsonData := []byte(`{
		"name": "小红",
		"birthday": "1999-03-04",
		"birthPlace": "北京市",
		"opus": {
			"2012": {
				"type": "类型一",
				"title": "《2012》"
			},
			"2013": {
				"type": "类型一",
				"title": "《2013》"
			},
			"2014": {
				"type": "类型二",
				"title": "《2014》"
			}
		}
	}`)

	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("出生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for index, val := range actress.Opus {
		fmt.Printf("\t年份：%s，类型：%s, 标题：%s\n", index, val.Type, val.Title)
	}

}
