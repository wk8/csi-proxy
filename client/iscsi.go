package client

// TODO wkpo imports
import (
	"context"
	"fmt"
	"github.com/Microsoft/go-winio"
	"google.golang.org/grpc"
	"net"

	pb "github.com/kubernetes-csi/csi-proxy/api"
)

type IscsiClient struct {
}

func (*IscsiClient) wkpo() {
	conn, err := grpc.Dial(`\\.\pipe\csi-proxy-v1alpha1`,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// TODO wkpo important ca... pas oublier ca!
	defer conn.Close()

	client := pb.NewIscsiCSIProxyServiceClient(conn)
	client.MountISCSILun()

	client := NewDummyClient(conn)

	req := &ComputeDoubleRequest{
		Input32: 28,
	}
	response, err := client.ComputeDouble(context.Background(), req)

	if err != nil {
		panic(err)
	}
	fmt.Println("wkpo got response", response.Response32)
}
