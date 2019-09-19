// Code generated by csi-proxy-gen. DO NOT EDIT.

// Deprecated: Do not use.
// v1alpha1 is no longer maintained, and will be removed soon.

package v1alpha1

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/wk8/csi-proxy/integrationtests/apigroups/api/dummy/v1alpha1"
	"github.com/wk8/csi-proxy/integrationtests/apigroups/server/dummy/internal"
	"github.com/wk8/csi-proxy/internal/apiversion"
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
