package main

import "fmt"

func main() {
	fmt.Println(1 & 2) // 01 10
	fmt.Println(2 & 2) // 10 10
	fmt.Println(3 & 2) // 11
	fmt.Println(4 & 2)
	fmt.Println(5 & 2)
}
