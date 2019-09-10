package dummy

import (
	"google.golang.org/grpc"

	"github.com/kubernetes-csi/csi-proxy/internal/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
)

const name = "dummy"

type VersionedAPI interface {
	Register(grpcServer *grpc.Server)
}

type VersionedAPIFactory func(*Server) VersionedAPI

var versionedAPIFactories = make(map[apiversion.Version]VersionedAPIFactory)

func RegisterVersion(version apiversion.Version, factory VersionedAPIFactory) {
	versionedAPIFactories[version] = factory
}

func (s *Server) VersionedAPIs() []*server.VersionedAPI {
	versionedAPIs := make([]*server.VersionedAPI, len(versionedAPIFactories))

	i := 0
	for version, versionedAPIFactory := range versionedAPIFactories {
		versionedServer := versionedAPIFactory(s)
		versionedAPIs[i] = &server.VersionedAPI{
			Group:      name,
			Version:    version,
			Registrant: versionedServer.Register,
		}
		i++
	}

	return versionedAPIs
}

// TODO wkpo auto generated comment!
