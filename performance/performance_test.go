package main

import (
	"fmt"
	"testing"
)

// 位移代替乘除
func BenchmarkOne(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s = i << 8
	}
	fmt.Println(s)
}

func BenchmarkTwo(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s = i * 256
	}
	fmt.Println(s)
}
