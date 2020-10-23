package main

import (
	"context"
	"fmt"
	"log"
	"mj_http/src/cgi"
	"mj_http/src/config"
	_ "mj_http/src/log_files"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
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

type Info struct {
	a int
	b int
}

func TestInof() {
	var ls []*Info
	for i := 0; i < 10; i++ {
		item := &Info{
			a: i,
			b: i,
		}
		ls = append(ls, item)
	}

	var ls1 []*Info
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ls1 = append(ls1, ls[i])
		}
	}

	ModifyInfo(ls1)

	for _, item := range ls {
		fmt.Println("item", item)
	}
}

func ModifyInfo(ls []*Info) {
	for _, item := range ls {
		item.a = 100
		item.b = 200
	}
}
func monitor(ctx context.Context, index int) {
	for {
		select {
		//case <-ctx.Done():
		//	fmt.Println("监控退出，停止了...", index, ctx.Err())
		//	return
		default:
			fmt.Println("goroutine监控中...", index, ctx.Err())
			time.Sleep(2 * time.Second)
		}
	}
}


var m sync.Map

func UseMap() {

	//写
	m.Store("dablelv", "27")
	m.Store("cat", "28")

	//遍历
	//操作函数
	f := func(key, value interface{}) bool {
		fmt.Printf("Range: k, v = %v, %v\n", key, value)
		return true
	}
	m.Range(f)

	m.Range(func(key, value interface{}) bool {
		m.Delete(key)
		return true
	})
	fmt.Println("after delete")

	m.Range(f)

	fmt.Println("end")
}

func main() {
	//WritePkgWithPipeline("pkg_set", CMD_SET)
	//WritePkgWithPipeline("pkg_hash", CMD_HASH)
	//WritePkgWithPipeline("pkg_bitmap", CMD_BITMAP)
	//
	//ReadPkgWithPipeline("pkg_set", CMD_SET)
	//ReadPkgWithPipeline("pkg_hash", CMD_HASH)
	//ReadPkgWithPipeline("pkg_bitmap", CMD_BITMAP)
	//GetMax()
	startHttp()
}

// http://localhost:8080/points_mall/order?time=20191115141100&appId=TRM191001&sign=f1fe0364e725a9e0e45def0685a3cd19&orderNo=252238532_12222222&account=7Rcs3IvPq/fZtvmTXVID+g==&type=1
