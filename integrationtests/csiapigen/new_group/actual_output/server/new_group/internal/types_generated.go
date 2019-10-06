package internal

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

type VersionedAPI interface {
	Register(grpcServer *grpc.Server)
}

// All the functions this group's server needs to define.
type ServerInterface interface {
	Foo(context.Context, *FooRequest, apiversion.Version) (*FooResponse, error)
}
