package main

import (
	"fmt"
	"github.com/Microsoft/go-winio"
	pb "github.com/kubernetes-csi/csi-proxy/api"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversion/definitions"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
	"github.com/kubernetes-csi/csi-proxy/internal/server/iscsi"
	"google.golang.org/grpc"
	"os/exec"
)

func main() {
	s, err := server.NewServer(definitions.V1alpha1)
	if err != nil {
		panic(err)
	}

	s.Start(nil)
}

func main3() {
	// TODO wkpo to list named pipes!
	out, err := exec.Command("powershell", `[System.IO.Directory]::GetFiles('\\.\\pipe\\')`).CombinedOutput()
	if err == nil {
		fmt.Println("success", string(out))
	} else {
		fmt.Println("error", err)
	}
}

// TODO wkpo
func main2() {
	grpcServer := grpc.NewServer()
	pb.RegisterIscsiCSIProxyServiceServer(grpcServer, &iscsi.Server{})

	listener, err := winio.ListenPipe(`\\.\pipe\csi-proxy-v1alpha1`, nil)
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
