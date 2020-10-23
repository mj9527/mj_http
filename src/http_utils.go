/**
 * @Author: mjzheng
 * @Description:
 * @File:  http_utils.go
 * @Version: 1.0.0
 * @Date: 2020/6/17 下午4:58
 */

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SimpleGet() {
	rsp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()

	body, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(body))
}

func SimplePost() {
	reqBody := "{\"action\":20}"
	rsp, err := http.Post("http://www.baidu.com", "application/json;charset=utf-8", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rsp.Body.Close()

	body, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(body))
}
