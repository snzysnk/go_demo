package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//go:generate go run main.go
func main() {
	_ = os.Setenv("X_TABLE", "hello world!")
	_ = os.Setenv("TABLE", "city!")
	viper.SetEnvPrefix("X")
	viper.AutomaticEnv()
	fmt.Println(viper.Get("TABLE"))
}
