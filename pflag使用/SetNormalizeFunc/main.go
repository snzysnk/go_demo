package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"os"
	"strings"
)

//go:generate go run main.go -i="192.168.1.1"
func main() {
	var ip string
	flag := pflag.NewFlagSet("test", pflag.ExitOnError)
	flag.StringVarP(&ip, "ip-address", "i", "127.0.0.1", "请输入ip")
	flag.SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	})
	_ = flag.Parse(os.Args[1:])
	if getString, err := flag.GetString("ip_address"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(getString)
	}

	if getString, err := flag.GetString("ip-address"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(getString)
	}
}
