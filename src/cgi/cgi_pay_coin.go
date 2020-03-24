package cgi

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"mj_http/src/points_mall"
	"net/http"
	"strings"
)

func PayHandler(rsp http.ResponseWriter, req *http.Request) {

	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := points_mall.NewGreeterClient(conn)

	request := &points_mall.HelloRequest{
		Name: "hello",
	}

	response, err := t.SayHello(context.Background(), request)
	if err != nil {
		log.Println("call failed ")
	}

	log.Println("recv response", response)

	//log.Printf("context %v", req.Context())
	//
	//fmt.Fprintln(rsp, "remote address ", req.RemoteAddr)
	//GetPBInterface1(req)
	//
	//fmt.Fprintln(rsp, ComposeRequest(req))
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

func GetPBInterface1(req *http.Request) {

	//req.ParseForm()
	//raw := &points_mall.PayCoinReq1{}
	//
	//decoder := schema.NewDecoder()
	//
	//err := decoder.Decode(raw, req.Form)
	//
	//log.Printf("map struct interface [%v] err[%v]\n", raw, err)
}
