package v1alpha1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/api/file_system/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server/file_system/internal"
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
	v1alpha1.RegisterFileSystemServer(grpcServer, s)
}

func (s *versionedAPI) PathExists(context context.Context, versionedRequest *v1alpha1.PathExistsRequest) (*v1alpha1.PathExistsResponse, error) {
	request := &internal.PathExistsRequest{}
	if err := Convert_v1alpha1_PathExistsRequest_To_internal_PathExistsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.PathExists(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.PathExistsResponse{}
	if err := Convert_internal_PathExistsResponse_To_v1alpha1_PathExistsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
