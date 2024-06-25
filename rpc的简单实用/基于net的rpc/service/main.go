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

type UserAbility interface {
	Login(req *UserRequest, res *UserResponse) error
}

type UserService struct {
}

func (u UserService) Login(req *UserRequest, res *UserResponse) error {
	res.Token = req.Name + req.Password
	return nil
}

func main() {
	err := rpc.Register(new(UserService))
	listen, err := net.Listen("tcp", ":9001")
	defer listen.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go rpc.ServeConn(conn)
	}

}
