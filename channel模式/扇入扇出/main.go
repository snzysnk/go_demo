package main

import (
	"fmt"
	"sync"
)

var example [][]int
var inCh chan int

func init() {
	example = append(example, []int{1, 2, 3, 4})
	example = append(example, []int{11, 12, 13, 14})
	example = append(example, []int{21, 22, 23, 24})
	inCh = make(chan int)
}

// 扇入
func into() {
	count := 0
	for v := range inCh {
		count += v
	}
	fmt.Println(count)
}

// 扇出
func out() {
	var wg sync.WaitGroup
	wg.Add(len(example))
	for _, input := range example {
		go add(inCh, input, &wg)
	}
	wg.Wait()
	close(inCh)
}

func add(inCh chan int, input []int, wg *sync.WaitGroup) {
	count := 0
	for _, v := range input {
		count += v
	}
	inCh <- count
	wg.Done()
}

func main() {
	go out()
	into() //150
}
