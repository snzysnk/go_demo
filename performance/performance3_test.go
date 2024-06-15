package performance

import (
	"fmt"
	"testing"
)

func BenchmarkFive(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += Calculate(1, 10)
		s += Calculate(2, 20)
		s += Calculate(1, 10)
		s += Calculate(2, 10)
	}
	fmt.Println(s)
}

func BenchmarkSix(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += Calculate(1, 10)
		s += Calculate(2, 20)
		s += Calculate(1, 10)
		s += Calculate(2, 10)
	}
	fmt.Println(s)
}
