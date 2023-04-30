package gapi

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func UnaryUserValidationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	fmt.Println("UnaryUserValidationInterceptor")
	return handler(ctx, req)
}
