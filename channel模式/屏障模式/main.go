package main

import (
	"fmt"
	"sync"
	"time"
)

func apiA(res *int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	*res = 1
	wg.Done()
}

func apiB(res *int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	*res = 2
	wg.Done()
}

func apiC(res *int, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	*res = 3
	wg.Done()
}

func main() {
	var (
		wg sync.WaitGroup
		a  int
		b  int
		c  int
	)
	wg.Add(3)
	go apiA(&a, &wg)
	go apiB(&b, &wg)
	go apiC(&c, &wg)
	wg.Wait()
	//等待所有goroutine完成
	fmt.Println(a, b, c) // 1,2,3
}
