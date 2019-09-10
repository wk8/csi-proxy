package v1

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"google.golang.org/grpc"
)

func GetConnection() (*grpc.ClientConn, error) {
	return grpc.Dial(`\\.\pipe\csi-proxy-v1alpha1`,
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(ctx, s)
		}),
		grpc.WithInsecure())
}

// TODO wkpo auto-gen comment
