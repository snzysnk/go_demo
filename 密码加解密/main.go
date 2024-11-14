package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

//go:generate go run main.go
func main() {
	hash1, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	hash2, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(hash1))
	fmt.Println(string(hash2))
	err = bcrypt.CompareHashAndPassword(hash1, []byte("123456"))
	if err != nil {
		fmt.Println("密码错误!!")
		return
	}
	fmt.Println("密码正确")
}
