package v1alpha2

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha2"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v1alpha2")

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1alpha2.RegisterDummyServer(grpcServer, s)
}

func (s *versionedAPI) ComputeDouble(context context.Context, versionedRequest *v1alpha2.ComputeDoubleRequest) (*v1alpha2.ComputeDoubleResponse, error) {
	request := &internal.ComputeDoubleRequest{}
	if err := Convert_v1alpha2_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ComputeDouble(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha2.ComputeDoubleResponse{}
	if err := Convert_internal_ComputeDoubleResponse_To_v1alpha2_ComputeDoubleResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
