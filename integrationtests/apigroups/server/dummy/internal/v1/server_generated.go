package v1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	v1 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v1")

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1.RegisterDummyServer(grpcServer, s)
}

func (s *versionedAPI) ComputeDouble(context context.Context, versionedRequest *v1.ComputeDoubleRequest) (*v1.ComputeDoubleResponse, error) {
	request := &internal.ComputeDoubleRequest{}
	if err := Convert_v1_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ComputeDouble(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.ComputeDoubleResponse{}
	if err := Convert_internal_ComputeDoubleResponse_To_v1_ComputeDoubleResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) TellMeAPoem(context context.Context, versionedRequest *v1.TellMeAPoemRequest) (*v1.TellMeAPoemResponse, error) {
	request := &internal.TellMeAPoemRequest{}
	if err := Convert_v1_TellMeAPoemRequest_To_internal_TellMeAPoemRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.TellMeAPoem(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1.TellMeAPoemResponse{}
	if err := Convert_internal_TellMeAPoemResponse_To_v1_TellMeAPoemResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
