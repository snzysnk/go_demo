package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

//go:generate go run main.go
func main() {
	_ = os.Setenv("X_TABLE_ONE", "hello world!")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "@", "_"))
	fmt.Println(viper.Get("x@table.one"))
}
