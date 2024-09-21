package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

//go:generate go run main.go
func main() {
	viper.SetConfigFile("cfg.yaml")
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
	})
	viper.WatchConfig()

	time.Sleep(20 * time.Second)
}
