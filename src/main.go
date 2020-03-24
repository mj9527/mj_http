package main

import (
	"fmt"
	"log"
	"mj_http/src/cgi"
	"mj_http/src/config"
	_ "mj_http/src/log_files"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

// http://localhost:9001/points_mall/
//  http://localhost:9001/debug/pprof/goroutine

type PointsHandler struct {
}

func (*PointsHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("Hello points new"))
	log.Println("recv hello points ", runtime.NumGoroutine())

	ch := make(chan string)

	for i := 0; i < 1000; i++ {
		go func(ch chan string) {
			var input string
			for {
				input = <-ch
			}
			//time.Sleep(time.Millisecond * 3000)
			fmt.Println(input)
			//fmt.Println("start test go routine")
		}(ch)
	}
}

func startHttp() {
	http.Handle("/", &PointsHandler{})
	http.HandleFunc("/points_mall/balance", cgi.QueryHandler)
	http.HandleFunc("/points_mall/order", cgi.PayHandler)

	port := config.ServerConfig.ServerInfo.Port
	addr := fmt.Sprintf(":%d", port)
	log.Println("init addr ", addr)
	http.ListenAndServe(addr, nil)
}

func main() {

	startHttp()
}

// http://localhost:8080/points_mall/order?time=20191115141100&appId=TRM191001&sign=f1fe0364e725a9e0e45def0685a3cd19&orderNo=252238532_12222222&account=7Rcs3IvPq/fZtvmTXVID+g==&type=1
