package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type UserRequest struct {
	Name     string
	Password string
}

type UserResponse struct {
	Token string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:9001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	req := &UserRequest{
		Name:     "John",
		Password: "123456",
	}

	var res UserResponse
	err = client.Call("UserService.Login", req, &res)
	if err != nil {
		log.Fatal("call:", err)
	}

	fmt.Println("Token:", res.Token) //Token: John123456
}
