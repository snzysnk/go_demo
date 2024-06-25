package main

import (
	"fmt"
	"log"
	"net"
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
	dial, err := net.Dial("tcp", "localhost:9001")
	defer dial.Close()
	if err != nil {
		log.Fatal("dialing:", err)
	}

	req := &UserRequest{
		Name:     "John",
		Password: "123456",
	}

	var res UserResponse
	client := rpc.NewClient(dial)
	err = client.Call("UserService.Login", req, &res)
	if err != nil {
		log.Fatal("call:", err)
	}

	fmt.Println("Token:", res.Token) //Token: John123456
}
