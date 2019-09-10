package v1

import "google.golang.org/grpc"

// TODO wkpo remove
func wkpo() {
	grpcServer := grpc.NewServer()
	RegisterDummyServer(grpcServer, &v1Server{})
}
