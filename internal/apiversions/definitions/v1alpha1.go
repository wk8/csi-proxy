package definitions

import (
	"google.golang.org/grpc"

	pb "github.com/kubernetes-csi/csi-proxy/api"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
	"github.com/kubernetes-csi/csi-proxy/server/iscsi"
)

var v1alpha1Version = apiversions.NewVersion("v1alpha1")

var V1alpha1 = apiversions.Definition{
	Version: v1alpha1Version,

	BuildAndRegisterServers: func(grpcServer *grpc.Server, handler *apiversions.Handler) {
		iscsiServer := &iscsi.IscsiServer{}

		pb.RegisterIscsiCSIProxyServiceServer(grpcServer, iscsiServer)
	},
}
