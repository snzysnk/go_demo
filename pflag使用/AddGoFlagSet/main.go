package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
)

var ip string

//go:generate go run main.go --ip="192.168.1.1"
func main() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip address description")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	fmt.Println(ip)
}
