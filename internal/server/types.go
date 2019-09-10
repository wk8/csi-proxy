package server

import (
	"google.golang.org/grpc"

	"github.com/kubernetes-csi/csi-proxy/internal/apiversion"
)

type VersionedApi struct {
	Group      string
	Version    apiversion.Version
	Registrant func(*grpc.Server)
}

type ApiGroup interface {
	VersionedApis() []*VersionedApi
}
