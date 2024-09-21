package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

//go:generate go run main.go
func main() {
	_ = os.Setenv("TABLE", "")
	viper.AutomaticEnv()
	viper.SetConfigFile("cfg.yaml")
	fmt.Println(viper.Get("x@table.one"))
}
