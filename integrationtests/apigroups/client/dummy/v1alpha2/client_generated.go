package v1alpha2

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha2"
	"google.golang.org/grpc"
)

const groupName = "dummy"

var version = apiversion.NewVersionOrPanic("v1alpha2")

type wrapper struct {
	client     v1alpha2.DummyClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the dummy API group version v1alpha2.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*wrapper, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1alpha2.NewDummyClient(connection)
	return &wrapper{
		client:     client,
		connection: connection,
	}, nil
}

// ensures we implement all the required methods
var _ v1alpha2.DummyClient = &wrapper{}

func (w *wrapper) ComputeDouble(context context.Context, request *v1alpha2.ComputeDoubleRequest, opts ...grpc.CallOption) (*v1alpha2.ComputeDoubleResponse, error) {
	return w.client.ComputeDouble(context, request, opts...)
}
