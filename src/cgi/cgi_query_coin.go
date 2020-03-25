package cgi

import (
	"log"
	"net/http"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func QueryHandler(rsp http.ResponseWriter, req *http.Request) {
	start := time.Now()
	rsp.Write([]byte("pay handler"))
	log.Println("recv pay request", req)

	//	SendReq()
	end := time.Now()
	delta := end.Sub(start)
	log.Printf("execture time %d", delta)

}
