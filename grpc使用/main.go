package main

import (
	"context"
	"fmt"
	"os"
	"use_grpc/service"
)

// 启动 server
// go run main.go server 127.0.0.1:9001
// 启动 client
// go run main.go client 127.0.0.1:9001
func main() {
	fmt.Println(os.Args[1], os.Args[2])
	if os.Args[1] == "server" {
		service.StartRpcServer(context.Background(), os.Args[2])
	} else if os.Args[1] == "client" {
		service.StartRpcClient(context.Background(), os.Args[2])
	}
}
