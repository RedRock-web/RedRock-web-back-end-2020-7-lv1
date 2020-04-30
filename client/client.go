package main

import (
	"RedRock-web-back-end-2020-7-lv1/hello"
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	//address     = "47.98.57.152:1234"
	address = "127.0.0.1:1234"
)

func main() {
	//连接tcp
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//将tcp连接的客户端连接上rpc客户端
	c := hello.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	//调用服务
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Greeting: %s\n", r.GetMessage())
}
