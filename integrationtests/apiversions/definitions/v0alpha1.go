package definitions

import (
	"google.golang.org/grpc"

	v0alpha1Pkg "github.com/kubernetes-csi/csi-proxy/integrationtests/apiversions/v0alpha1"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

var v0alpha1Version = apiversions.NewVersion("v0alpha1")

var v0alpha1 = apiversions.Definition{
	Version: v0alpha1Version,

	BuildAndRegisterServers: func(grpcServer *grpc.Server) {
		testServer := &v0alpha1Pkg.TestServer{Version: v0alpha1Version}

		v0alpha1Pkg.RegisterTestCSIProxyServiceServer(grpcServer, testServer)
	},
}
