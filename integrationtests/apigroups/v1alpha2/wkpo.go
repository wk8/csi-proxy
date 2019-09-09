package v1alpha2

import "google.golang.org/grpc"

// TODO wkpo remove
func wkpo() {
	grpcServer := grpc.NewServer()
	RegisterTestCSIProxyServiceServer(grpcServer, &v1alpha2Server{})
}
