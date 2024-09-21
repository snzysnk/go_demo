package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"os"
)

//go:generate go run main.go -i="192.168.1.1"
func main() {
	var ip string
	flag := pflag.NewFlagSet("test", pflag.ExitOnError)
	flag.StringVarP(&ip, "ip", "i", "127.0.0.1", "请输入ip")
	_ = flag.Parse(os.Args[1:])
	if getString, err := flag.GetString("ip"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(getString)
	}
}
