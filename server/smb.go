package server

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/kubernetes-csi/csi-proxy/api"
)

// TODO wkpo see https://blog.golang.org/generate ? add a //go:generate comment?

// TODO wkpo en fait...!!! separer les serveurs en differents types! un

type SmbServer struct {
}

func (s *SmbServer) MountSMBShare(ctx context.Context, request *pb.MountSMBShareRequest) (*pb.MountSMBShareResponse, error) {
	return nil, nil
}

// TODO wkpo
func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterSmbCSIProxyServiceServer(grpcServer, &SmbServer{})
}
