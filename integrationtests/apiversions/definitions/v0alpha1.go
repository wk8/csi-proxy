package definitions

import (
	"google.golang.org/grpc"
	
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

var v0alpha1 = apiversions.Definition{
	Version: apiversions.NewVersion("v0alpha1"),

	BuildAndRegisterServers: func(server *grpc.Server) {

	}
}
