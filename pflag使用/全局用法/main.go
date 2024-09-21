package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

//go:generate go run main.go --ip="192.168.1.1"
func main() {
	var ip = pflag.String("ip", "127.0.0.1", "请输入ip")
	pflag.Parse()
	fmt.Println(*ip)
}
