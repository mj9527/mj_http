package cgi

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mj9527/points_mall"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {

}

func PayHandler(rsp http.ResponseWriter, req *http.Request) {

	GetPBInterface1(req)

	fmt.Fprintln(rsp, ComposeRequest(req))
}

func ComposeRequest(req *http.Request) string {
	raw := fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		raw += fmt.Sprintf("%s:%v\r\n", k, v[0])
	}
	raw += fmt.Sprintf("\r\n")
	body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		raw += string(body)
	}
	return raw
}

func GetPBInterface1(req *http.Request) {

	req.ParseForm()
	raw := &points_mall.PayCoinReq1{}

	decoder := schema.NewDecoder()

	err := decoder.Decode(raw, req.Form)

	log.Printf("map struct interface [%v] err[%v]\n", raw, err)
}
