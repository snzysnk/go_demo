package main

import (
	"fmt"
	"github.com/golang/groupcache/consistenthash"
)

func main() {
	//每个ip增加五个虚拟节点
	m := consistenthash.New(5, nil)
	m.Add("127.0.0.1")
	m.Add("127.0.0.2")
	m.Add("127.0.0.3")
	fmt.Println(m.Get("127.0.0.4"))
}
