package main

import (
	"context"
	"fmt"
	"github.com/Microsoft/go-winio"
	"net"

	"google.golang.org/grpc"
)

// TODO wkpo
func clientMain() {
	conn, err := grpc.Dial(`\\.\pipe\csi-proxy-v1alpha1`,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := NewTestCSIProxyServiceClient(conn)

	req := &ComputeDoubleRequest{
		Input32: 28,
	}
	response, err := client.ComputeDouble(context.Background(), req)

	if err != nil {
		panic(err)
	}
	fmt.Println("wkpo got response", response.Response32)
}
