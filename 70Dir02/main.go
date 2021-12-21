package main

import (
	"fmt"
	"os"
)

func main()  {

	path := "D:/temp/"

	result, err := readDir(path, 100, 1)
	if err != nil {
		panic(err)
	}

	for _, ret := range result {
		fmt.Println(ret)
	}

}

// readDir 递归遍历目录
func readDir(path string, deep int, lv int) ([]map[string]interface{}, error) {
	handle, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, err
	}

	info, err := handle.Readdir(-1)
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	for _, fileInfo := range info {
		result = append(result, map[string]interface{}{
			"Name": fileInfo.Name(),
			"IsDir": fileInfo.IsDir(),
			"Lv": lv,
		})
		if fileInfo.IsDir() && deep > lv {
			ret, err := readDir(path + fileInfo.Name() + "/", deep, lv + 1)
			if err != nil {
				return nil, err
			}
			result = append(result, ret...)
		}
	}

	return result, nil
}
