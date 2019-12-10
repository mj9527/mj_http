package main

import (
	"fmt"
	"io"
	"log"
	"mj_http/src/cgi"
	"mj_http/src/config"
	"net/http"
	"os"
)

func init() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//t.Fatalf("error opening file: %v", err)
		return
	}
	//defer f.Close()
	//
	writers := []io.Writer{
		f,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(fileAndStdoutWriter)
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
}

type PointsHandler struct {
}

func (*PointsHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("Hello points new"))
	log.Println("recv hello points")
}

func startHttp() {
	http.Handle("/", &PointsHandler{})
	http.HandleFunc("/points_mall/balance", cgi.PayHandler)
	http.HandleFunc("/points_mall/order", cgi.QueryHandler)

	port := config.ServerConfig.ServerInfo.Port
	addr := fmt.Sprintf(":%d", port)
	log.Println("init addr ", addr)
	http.ListenAndServe(addr, nil)
}

func main() {

	startHttp()
}

// http://localhost:8080/points_mall/balance?time=20191115141100&appId=TRM191001&sign=f1fe0364e725a9e0e45def0685a3cd19&orderNo=252238532_12222222&account=7Rcs3IvPq/fZtvmTXVID+g==&type=1
