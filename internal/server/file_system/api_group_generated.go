package filesystem

import (
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
	"github.com/kubernetes-csi/csi-proxy/internal/server/file_system/internal"
	"github.com/kubernetes-csi/csi-proxy/internal/server/file_system/internal/v1alpha1"
)

const name = "file_system"

// ensure the server defines all the required methods
var _ internal.ServerInterface = &Server{}

func (s *Server) VersionedAPIs() []*server.VersionedAPI {
	v1alpha1Server := v1alpha1.NewVersionedServer(s)

	return []*server.VersionedAPI{
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha1"),
			Registrant: v1alpha1Server.Register,
		},
	}
}
