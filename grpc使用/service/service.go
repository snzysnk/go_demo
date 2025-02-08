package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	user "use_grpc/service/api"
)

type UserService struct{}

// GetUserInfo 实现pb中的 LoginServiceServer接口
func (u UserService) GetUserInfo(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	return &user.LoginResponse{Token: request.Name + request.Password}, nil
}

func StartRpcServer(ctx context.Context, address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	// 加入日志拦截器
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(LogInterceptor()))
	user.RegisterLoginServiceServer(server, &UserService{})
	log.Fatal(server.Serve(listen))
}
