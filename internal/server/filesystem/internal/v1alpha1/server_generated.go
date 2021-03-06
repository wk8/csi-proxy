// Code generated by csi-proxy-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/wk8/csi-proxy/client/api/filesystem/v1alpha1"
	"github.com/wk8/csi-proxy/client/apiversion"
	"github.com/wk8/csi-proxy/internal/server/filesystem/internal"
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
	pb.RegisterFileSystemServer(grpcServer, s)
}

func (s *versionedAPI) PathExists(ctx context.Context, versionedRequest *pb.PathExistsRequest) (*pb.PathExistsResponse, error) {
	request := &internal.PathExistsRequest{}
	if err := convert_pb_PathExistsRequest_To_internal_PathExistsRequest(versionedRequest, request); err != nil {
		return nil, err
	}
	response, err := s.apiGroupServer.PathExists(ctx, request, version)
	if err != nil {
		return nil, err
	}
	versionedResponse := &pb.PathExistsResponse{}
	if err = convert_internal_PathExistsResponse_To_pb_PathExistsResponse(response, versionedResponse); err != nil {
		return nil, err
	}
	return versionedResponse, nil
}
