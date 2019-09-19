package server

import (
	"google.golang.org/grpc"

	"github.com/wk8/csi-proxy/internal/apiversion"
)

type VersionedAPI struct {
	Group      string
	Version    apiversion.Version
	Registrant func(*grpc.Server)
}

type APIGroup interface {
	VersionedAPIs() []*VersionedAPI
}
