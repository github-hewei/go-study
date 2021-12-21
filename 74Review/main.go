package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	token, err := getAccessToken(13)
	if err != nil {
		panic(err)
	}

	fmt.Println(token)


}

// 获取微信Token
func getAccessToken(sid int) (map[string]string, error) {
	url := fmt.Sprintf("http://xxx.com/api/v1/wechat/get_access_token?sid=%d", sid)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	if code, ok := ret["errcode"]; ok && code.(float64) == 100 {
		data := ret["data"].(map[string]interface{})
		r := make(map[string]string)
		r["appid"] = data["appid"].(string)
		r["token"] = data["access_token"].(string)
		return r, nil
	}

	err = errors.New("接口异常")
	if m, ok := ret["msg"]; ok && m.(string) != "" {
		err = errors.New(m.(string))
	}

	return nil, err
}
