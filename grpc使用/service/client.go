package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	user "use_grpc/service/api"
)

func StartRpcClient(ctx context.Context, address string) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := user.NewLoginServiceClient(conn)
	info, err := client.GetUserInfo(context.Background(), &user.LoginRequest{
		Name:     "kiki",
		Password: "@123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info.Token) //kiki@123456
}
