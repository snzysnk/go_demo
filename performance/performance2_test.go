package performance

import (
	"fmt"
	"testing"
)

var cache map[int]int

func init() {
	cache = make(map[int]int)
}

func CalculateWithCache(begin, end int) int {
	var (
		res int
		key = begin + end
	)
	if m, ok := cache[key]; ok {
		return m
	}
	for i := begin; i < end; i++ {
		res += i
	}
	cache[begin+end] = res
	return res
}

func Calculate(begin, end int) int {
	var res int
	for i := begin; i < end; i++ {
		res += i
	}
	return res
}

func BenchmarkThree(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += CalculateWithCache(1, 10)
		s += CalculateWithCache(2, 20)
		s += CalculateWithCache(1, 10)
		s += CalculateWithCache(2, 10)
	}
	fmt.Println(s)
}

func BenchmarkFour(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += Calculate(1, 10)
		s += Calculate(2, 20)
		s += Calculate(1, 10)
		s += Calculate(2, 10)
	}
	fmt.Println(s)
}
