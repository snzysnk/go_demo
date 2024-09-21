package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//go:generate go run main.go
func main() {
	viper.Set("table", "users")
	fmt.Println(viper.GetString("table"))
}
