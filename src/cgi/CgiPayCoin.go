package cgi

import (
	"fmt"
	"net/http"
)

func init() {

}

func PayHandler(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("query handler"))
	fmt.Print("recv query handle", req)
}
