package cgi

import (
	"context"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/mj9527/points_mall"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
)

func PayHandler(rsp http.ResponseWriter, req *http.Request) {

	request := GetPBInterface1(req)

	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := points_mall.NewPointMallClient(conn)

	//GetPBInterface1(req)

	response, err := t.PayCoin(context.Background(), request)
	if err != nil {
		log.Println("call failed ")
	}

	log.Println("recv response", response)

	fmt.Fprintln(rsp, response)
}

func ComposeRequest(req *http.Request) string {
	//raw := fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL, req.Proto)
	//for k, v := range req.Header {
	//	raw += fmt.Sprintf("%s:%v\r\n", k, v[0])
	//	//log.Printf("[%s:%v] len %d [%v] ", k, v[0], len(v), v)
	//}
	//raw += fmt.Sprintf("\r\n")
	//body, err := ioutil.ReadAll(req.Body)
	//if err == nil {
	//	raw += string(body)
	//}
	buf := strings.Builder{}
	buf.WriteString(req.Method)
	buf.WriteString(" ")
	buf.WriteString(req.URL.RawPath)
	buf.WriteString(" ")
	buf.WriteString(req.Proto)
	buf.WriteString("\r\n")

	//raw := fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		buf.WriteString(k)
		buf.WriteString(":")
		buf.WriteString(v[0])
		buf.WriteString("\r\n")
		//log.Printf("[%s:%v] len %d [%v] ", k, v[0], len(v), v)
	}
	buf.WriteString("\r\n")

	return buf.String()
}

func GetPBInterface1(req *http.Request) *points_mall.PayCoinReq {

	req.ParseForm()
	raw := &points_mall.PayCoinReq{}
	decoder := schema.NewDecoder()
	err := decoder.Decode(raw, req.Form)

	log.Printf("map struct interface [%v] err[%v]\n", raw, err)
	return raw
}
