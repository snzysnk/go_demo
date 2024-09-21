package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//go:generate go run main.go
func main() {
	fmt.Println(viper.Get("table"))
	viper.SetConfigName("cfg")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
	fmt.Println(viper.Get("table"))
}
