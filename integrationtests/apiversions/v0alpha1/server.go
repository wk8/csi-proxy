package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Microsoft/go-winio"
	"google.golang.org/grpc"

	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

// TODO wkpo see https://blog.golang.org/generate ? add a //go:generate comment?

type TestServer struct {
	version apiversions.Version
}

func (s *TestServer) ComputeDouble(ctx context.Context, request *ComputeDoubleRequest) (*ComputeDoubleResponse, error) {
	response := &ComputeDoubleResponse{
		Response32: 2 * request.Input32,
	}

	var err error
	if true {
		err = fmt.Errorf("wkpo bordel error")
		err = grpc.ErrServerStopped
	}
	return response, err
}

func main() {
	msg := "one arg required: c or s"
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "s":
			serverMain()
		case "c":
			clientMain()
		default:
			panic(msg)
		}
	} else {
		panic(msg)
	}
}

// TODO wkpo
func serverMain() {
	grpcServer := grpc.NewServer()
	RegisterTestCSIProxyServiceServer(grpcServer, &TestServer{})

	listener, err := winio.ListenPipe(`\\.\pipe\csi-proxy-v1alpha1`, nil)
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
