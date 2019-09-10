package v1alpha2

import (
	"context"

	server "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
)

// TODO wkpo actually generate that shit...

// TODO wkpo rename to versionedServer, and same below?
type v1alpha2Server struct {
	apiGroupServer *server.Server
	// TODO wkpo mouaif...? hard coder en fait, generated...
	version server.Version
}

// TODO wkpo next, apres generation
func (s *v1alpha2Server) ComputeDouble(ctx context.Context, v1alpha2Request *ComputeDoubleRequest) (*ComputeDoubleResponse, error) {
	request := &server.ComputeDoubleRequest{}
	if err := Convert_v1alpha2_ComputeDoubleRequest_To_server_ComputeDoubleRequest(v1alpha2Request, request); err != nil {
		return nil, err
	}
	response, err := s.apiGroupServer.ComputeDouble(ctx, request, s.version)
	if err != nil {
		return nil, err
	}
	v1alpha2Response := &ComputeDoubleResponse{}
	if err = Convert_server_ComputeDoubleResponse_To_v1alpha2_ComputeDoubleResponse(response, v1alpha2Response); err != nil {
		return nil, err
	}
	return v1alpha2Response, nil
}
