package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//go:generate go run main.go
func main() {
	fmt.Println(viper.Get("table"))
	viper.SetDefault("table", "users")
	fmt.Println(viper.Get("table"))
}
