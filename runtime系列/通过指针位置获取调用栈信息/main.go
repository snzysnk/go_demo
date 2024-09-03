package main

import (
	"fmt"
	"runtime"
)

//go:generate go run main.go
func main() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(pc) //pc是指针地址

	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next() //迭代器
	fmt.Println(frame.File, frame.Line)
}
