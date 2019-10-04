package v1alpha1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v1alpha1")

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1alpha1.RegisterDummyServer(grpcServer, s)
}

func (s *versionedAPI) ComputeDouble(context context.Context, versionedRequest *v1alpha1.ComputeDoubleRequest) (*v1alpha1.ComputeDoubleResponse, error) {
	request := &internal.ComputeDoubleRequest{}
	if err := Convert_v1alpha1_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ComputeDouble(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.ComputeDoubleResponse{}
	if err := Convert_internal_ComputeDoubleResponse_To_v1alpha1_ComputeDoubleResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
