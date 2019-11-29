package cgi

import (
	"fmt"
	"net/http"
)

func QueryHandler(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("pay handler"))
	fmt.Print("recv pay request", req)
}
