package v1

import (
	"context"

	"google.golang.org/grpc"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	v1 "github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/api/dummy/v1"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server/dummy/internal"
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

func (s *versionedAPI) ComputeDouble(ctx context.Context, versionedRequest *v1.ComputeDoubleRequest) (*v1.ComputeDoubleResponse, error) {
	request := &internal.ComputeDoubleRequest{}
	if err := Convert_v1_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(versionedRequest, request); err != nil {
		return nil, err
	}
	response, err := s.apiGroupServer.ComputeDouble(ctx, request, version)
	if err != nil {
		return nil, err
	}
	versionedResponse := &v1.ComputeDoubleResponse{}
	if err = Convert_internal_ComputeDoubleResponse_To_v1_ComputeDoubleResponse(response, versionedResponse); err != nil {
		return nil, err
	}
	return versionedResponse, nil
}

func (s *versionedAPI) TellMeAPoem(ctx context.Context, versionedRequest *pb.TellMeAPoemRequest) (*pb.TellMeAPoemResponse, error) {
	request := &internal.TellMeAPoemRequest{}
	if err := convert_pb_TellMeAPoemRequest_To_internal_TellMeAPoemRequest(versionedRequest, request); err != nil {
		return nil, err
	}
	response, err := s.apiGroupServer.TellMeAPoem(ctx, request, version)
	if err != nil {
		return nil, err
	}
	versionedResponse := &pb.TellMeAPoemResponse{}
	if err = convert_internal_TellMeAPoemResponse_To_pb_TellMeAPoemResponse(response, versionedResponse); err != nil {
		return nil, err
	}
	return versionedResponse, nil
}
