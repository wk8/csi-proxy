package newgroup

import (
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/actual_output/server/new_group/internal"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/actual_output/server/new_group/internal/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/internal/server"
)

const name = "new_group"

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
