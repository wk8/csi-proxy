package dummy

import (
	"google.golang.org/grpc"

	"github.com/kubernetes-csi/csi-proxy/internal/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
)

const name = "dummy"

type VersionedApi interface {
	Register(grpcServer *grpc.Server)
}

type VersionedApiFactory func(*Server) VersionedApi

var versionedApiFactories = make(map[apiversion.Version]VersionedApiFactory)

func RegisterVersion(version apiversion.Version, factory VersionedApiFactory) {
	versionedApiFactories[version] = factory
}

func (s *Server) VersionedApis() []*server.VersionedApi {
	versionedApis := make([]*server.VersionedApi, len(versionedApiFactories))

	i := 0
	for version, versionedApiFactory := range versionedApiFactories {
		versionedServer := versionedApiFactory(s)
		versionedApis[i] = &server.VersionedApi{
			Group:      name,
			Version:    version,
			Registrant: versionedServer.Register,
		}
		i++
	}

	return versionedApis
}

// TODO wkpo auto generated comment!
// TODO wkpo rename to group_generated.go? or else?
