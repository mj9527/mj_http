package cgi

import (
	"fmt"
	"github.com/mj9527/points_mall"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func QueryHandler(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("pay handler"))
	fmt.Println("recv pay request", req)

	SendReq()

}

func SendReq() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := points_mall.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &points_mall.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting Again: %s\n", r.GetMessage())

}
