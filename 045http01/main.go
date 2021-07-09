package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// *** 结构体成员变量的首字母必须大写 ***

// 请求数据json结构体
type ReqData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 响应数据json结构体
type ResData struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}

func main() {

	url := "http://lmkcs.lmkdev.com/user/login"

	contentType := "application/json"

	reqData, err := json.Marshal(ReqData{"admin", "admin"})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(reqData))

	reqBuf := bytes.NewBuffer(reqData)

	res, err := http.Post(url, contentType, reqBuf)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	resData := new(ResData)

	err = json.Unmarshal(resBody, resData)
	if err != nil {
		panic(err)
	}

	fmt.Println((*resData).Code)
	fmt.Println((*resData).Msg)
	fmt.Println((*resData).Data)
	fmt.Println((*resData).Data["token"])

}
