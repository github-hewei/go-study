package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// IO: 输入Input + 输出Output
	r := strings.NewReader("123456\n5678\n91234")
	br := bufio.NewReader(r)

	for {
		content, _, err := br.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(string(content))
	}

	file, err := os.OpenFile("debug.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	b := make([]byte, 0)
	buf := bytes.NewBuffer(b)
	buf.WriteString("123456")
	_, _ = io.Copy(file, buf)

	//bufio.NewReader()
	//file.Read()
	//file.Write()
}
