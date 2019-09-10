package v1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"google.golang.org/grpc"

	pb "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	"github.com/kubernetes-csi/csi-proxy/internal"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversion"
)

const groupName = "dummy"

var version = apiversion.NewVersionOrPanic("v1")

type wrapper struct {
	client     pb.DummyClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the dummy API group version v1.
// It's the caller's responsibility to Close the client when done.
func NewClient() (pb.DummyClient, error) {
	pipePath := internal.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewDummyClient(connection)
	return &wrapper{
		client:     client,
		connection: connection,
	}, nil
}

// ComputeDouble computes the double of the input. Real smart stuff!
func (w *wrapper) ComputeDouble(ctx context.Context, in *pb.ComputeDoubleRequest, opts ...grpc.CallOption) (*pb.ComputeDoubleResponse, error) {
	return w.client.ComputeDouble(ctx, in, opts...)
}

// TellMeAPoem returns a very poetic response.
func (w *wrapper) TellMeAPoem(ctx context.Context, in *pb.TellMeAPoemRequest, opts ...grpc.CallOption) (*pb.TellMeAPoemResponse, error) {
	return w.client.TellMeAPoem(ctx, in, opts...)
}

func (w *wrapper) Close() error {
	return w.connection.Close()
}

// TODO wkpo auto-gen comment
