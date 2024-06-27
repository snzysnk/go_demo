package main

import (
	"cw/distribute"
	"fmt"
	"os"
)

// 启动master
// go run main.go master 127.0.0.1:9001
// 启动worker
// go run main.go worker 127.0.0.1:9001 127.0.0.1:9002
func main() {
	fmt.Println(os.Args[1], os.Args[2])
	if os.Args[1] == "master" {
		distribute.RunRpcMaster(os.Args[2])
	} else if os.Args[1] == "worker" {
		distribute.RunWorker(os.Args[2], os.Args[3])
	}
}
