package v1alpha1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/actual_output/server/new_group/internal"
	"github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/api/v1alpha1"
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
	v1alpha1.RegisterNewGroupServer(grpcServer, s)
}

func (s *versionedAPI) Foo(context context.Context, versionedRequest *v1alpha1.FooRequest) (*v1alpha1.FooResponse, error) {
	request := &internal.FooRequest{}
	if err := Convert_v1alpha1_FooRequest_To_internal_FooRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.Foo(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.FooResponse{}
	if err := Convert_internal_FooResponse_To_v1alpha1_FooResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
