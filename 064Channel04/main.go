package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://admin.lmkdev.com/feedback/push"

	client := &http.Client{}

	data := bytes.NewBufferString("{\"id\":\"19957\",\"confirm\":\"yes\",\"_csrf-backend\":\"FGADS8swVXpnHoz2SJC2aOVE4U98_kgh58vkBzM4UEFTCG0tgmA8DwhqyqVl0_ke0imZBzKwB1C9p5ZWHmkZDg==\",\"nonce\":\"SNcwhZER7fAsAyHe\",\"timestamp\":1632539990000}")

	req, err := http.NewRequest("POST", url, data)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Cookie", "_csrf-backend=8e3708c873c57c473799098ad1de1da7e6e6d28a1ccdd471202214f06391b69ca%3A2%3A%7Bi%3A0%3Bs%3A13%3A%22_csrf-backend%22%3Bi%3A1%3Bs%3A32%3A%22GhnfIPiuotFS-COv7mxHNNOqZlrQ-QIO%22%3B%7D; tips_password=0; advanced-backend=d7e4l997pij4cparv5kpf8da68")

	res, err := client.Do(req)

	defer res.Body.Close()

	fmt.Println(res)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resBody))


}



