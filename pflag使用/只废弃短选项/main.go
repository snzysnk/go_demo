package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var ip = pflag.StringP("ip", "p", "127.0.0.1", "请输入ip")
	_ = pflag.CommandLine.MarkShorthandDeprecated("ip", "p 已弃用")
	pflag.Parse()
	fmt.Println(*ip)
}
