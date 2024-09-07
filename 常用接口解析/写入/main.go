package main

import (
	"fmt"
	"io"
)

type XWriter struct {
}

func (X XWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

var _ io.Writer = XWriter{}

//go:generate go run main.go
func main() {
	var x XWriter
	l, err := x.Write([]byte("hello world"))
	fmt.Println(l, err)
}
