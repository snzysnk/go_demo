package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

//go:generate go run main.go --ip="192.168.1.1" "8080" "test"
func main() {
	var ip = pflag.StringP("ip", "p", "127.0.0.1", "请输入ip")
	_ = pflag.CommandLine.MarkHidden("ip")
	pflag.Parse()
	fmt.Println(*ip)
}
