package main

import (
	"fmt"
	"runtime"
)

//go:generate go run main.go
func main() {
	_, file, line, ok := runtime.Caller(0)
	fmt.Println(file, line, ok)
}
