package main

import (
	"fmt"
	"time"
)

func cut(ch chan string, desc string) {
	time.Sleep(3 * time.Second)
	ch <- desc
}

func main() {
	chApple := make(chan string)
	chBanana := make(chan string)
	go cut(chApple, "阿姨切好了苹果")
	go cut(chBanana, "阿姨切好了香蕉")
	time.Sleep(1 * time.Second)
	fmt.Println("休息一下")
	//吃阿姨切好的苹果，不吃也没事
	fmt.Println("吃" + <-chApple) //吃阿姨切好了苹果
}
