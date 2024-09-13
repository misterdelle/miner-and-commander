package main

import (
	"context"

	"google.golang.org/grpc"
)

// unaryAuthInterceptor is an interceptor automatically adding the auth token
// to a request.
func unaryAuthInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}
