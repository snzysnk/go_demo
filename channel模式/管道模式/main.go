package main

import (
	"fmt"
	"strconv"
)

func jobStart() <-chan string {
	outCh := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			outCh <- "工序begin" + strconv.Itoa(i)
		}
		close(outCh)
	}()
	return outCh
}

func jobB(inCh <-chan string) <-chan string {
	outCh := make(chan string)
	go func() {
		for v := range inCh {
			outCh <- v + "-工序B"
		}
		close(outCh)
	}()
	return outCh
}

func jobC(inCh <-chan string) <-chan string {
	outCh := make(chan string)
	go func() {
		for v := range inCh {
			outCh <- v + "-工序C"
		}
		close(outCh)
	}()
	return outCh
}

func main() {
	ch := jobStart()
	ch = jobB(ch)
	ch = jobC(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
