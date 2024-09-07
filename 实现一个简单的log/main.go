package main

import "x_log/xlog"

//go:generate go run main.go
func main() {
	xlog.Fatal("hello world")
}
