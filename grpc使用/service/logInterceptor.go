package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("LogInterceptor-start")
		handlerResult, err := handler(ctx, req)
		fmt.Println("LogInterceptor-end")
		return handlerResult, err
	}
}
