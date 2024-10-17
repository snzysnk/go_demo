package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var ip string
var isProduct bool

//go:generate go run main.go
func main() {
	pflag.StringVar(&ip, "ip", "127.0.0.1", "ip address description")
	pflag.BoolVar(&isProduct, "product", false, "is product description")
	pflag.VisitAll(func(f *pflag.Flag) {
		fmt.Printf("%s: %v\n", f.Name, f.Value)
	})
}
