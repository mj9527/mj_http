package cgi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {

}

func PayHandler(rsp http.ResponseWriter, req *http.Request) {
	//rsp.Write([]byte("pay handler"))
	//fmt.Fprintln(rsp, "pay handler1")
	//fmt.Print("recv query handle", req)

	//fmt.Fprintln(rsp, "method ", req.Method)
	//fmt.Fprintln(rsp, "url ", req.URL)
	//fmt.Fprintln(rsp, "body ", req.Body)
	fmt.Fprintln(rsp, "address ", req.RemoteAddr)
	fmt.Fprintln(rsp, "  ", ComposeRequest(req))

	req.ParseForm()

	fmt.Println(req.Form)

	for k, v := range req.Form {
		fmt.Printf("%v=%v\n", k, v)
	}
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
