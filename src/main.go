package main

import (
	"log"
	"mj_http/src/cgi"
	"net/http"
)

func init() {
	//f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	//t.Fatalf("error opening file: %v", err)
	//	return
	//}
	//defer f.Close()
	//
	//log.SetOutput(f)
	//log.Println("This is a test log entry")
	log.SetFlags(log.Ldate | log.Lshortfile)
}

type PointsHandler struct {
}

func (*PointsHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("Hello points new"))
	log.Println("recv hello points")
}

func main() {
	http.Handle("/", &PointsHandler{})
	http.HandleFunc("/points_mall/pay", cgi.PayHandler)
	http.HandleFunc("/points_mall/query", cgi.QueryHandler)
	http.ListenAndServe(":9001", nil)
}
