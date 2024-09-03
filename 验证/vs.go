package main

import "fmt"

type Example struct {
	code    int
	message string
}

type Example2 struct {
	code    int
	message string
}

func (e Example) String() string {
	return e.message
}

//go:generate go run main.go
func main() {
	e := Example{code: 0, message: "Hello"}
	e2 := Example2{code: 0, message: "World"}
	fmt.Printf("%v\n", e)
	fmt.Printf("%s\n", e)
	fmt.Printf("%v\n", e2)
}
