package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var ip = pflag.String("ip", "127.0.0.1", "请输入ip")
	pflag.Lookup("ip").NoOptDefVal = "0.0.0.0"
	pflag.Parse()
	fmt.Println(*ip)
}
