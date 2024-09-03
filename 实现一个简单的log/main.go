package main

import (
	"fmt"
	"runtime"
)

//go:generate go run main.go
func main() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(pc)
	//fmt.Println(pc, file, line, ok)
	frames := runtime.CallersFrames([]uintptr{pc})

	frame, _ := frames.Next()
	fmt.Printf("- Function: %s\n", frame.Function)
	fmt.Printf("  File: %s\n", frame.File)
	fmt.Printf("  Line: %d\n", frame.Line)

}
