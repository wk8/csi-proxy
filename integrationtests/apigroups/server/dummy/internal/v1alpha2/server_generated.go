// Code generated by csi-proxy-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"

	"google.golang.org/grpc"

	"github.com/wk8/csi-proxy/client/apiversion"
	pb "github.com/wk8/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha2"
	"github.com/wk8/csi-proxy/integrationtests/apigroups/server/dummy/internal"
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
	pb.RegisterDummyServer(grpcServer, s)
}

func (s *versionedAPI) ComputeDouble(ctx context.Context, versionedRequest *pb.ComputeDoubleRequest) (*pb.ComputeDoubleResponse, error) {
	request := &internal.ComputeDoubleRequest{}
	if err := convert_pb_ComputeDoubleRequest_To_internal_ComputeDoubleRequest(versionedRequest, request); err != nil {
		return nil, err
	}
	response, err := s.apiGroupServer.ComputeDouble(ctx, request, version)
	if err != nil {
		return nil, err
	}
	versionedResponse := &pb.ComputeDoubleResponse{}
	if err = convert_internal_ComputeDoubleResponse_To_pb_ComputeDoubleResponse(response, versionedResponse); err != nil {
		return nil, err
	}
	return versionedResponse, nil
}
