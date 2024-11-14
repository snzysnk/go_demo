package main

import (
	"fmt"
	"github.com/google/uuid"
)

//go:generate go run main.go
func main() {
	uid := uuid.New().String()
	//b20605e9-aced-4bba-b9a0-638da7067598
	fmt.Println(uid)
}
