package main

import (
	"log"
	"net/http"
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
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	log.Fatal(http.ListenAndServe(":9001", nil))
}
