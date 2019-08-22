package iscsi

import (
	"context"
	"github.com/golang/protobuf/proto"

	pb "github.com/kubernetes-csi/csi-proxy/api"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

// TODO wkpo see https://blog.golang.org/generate ? add a //go:generate comment?

type Server struct {
	apiVersion apiversions.Version
	handler    *apiversions.Handler
}

func NewServer(apiVersion apiversions.Version, handler *apiversions.Handler) *Server {
	return &Server{
		apiVersion: apiVersion,
		handler:    handler,
	}
}

// TODO wkpo NEXT! lo que pasa quand on renvoie une error? see below
//PS C:\Users\wk\csi-proxy\integrationtests\apiversions\v0alpha1> go run . c
//panic: rpc error: code = Unknown desc = grpc: the server has been stopped
//
//goroutine 1 [running]:
//main.clientMain()
//C:/Users/wk/csi-proxy/integrationtests/apiversions/v0alpha1/client.go:32 +0x2c5
//main.main()
//C:/Users/wk/csi-proxy/integrationtests/apiversions/v0alpha1/server.go:40 +0x6b
func (s *Server) MountISCSILun(ctx context.Context, request *pb.MountISCSILunRequest) (*pb.MountISCSILunResponse, error) {
	rawResponse, err := s.handler.WrapServerHandler(request, s.apiVersion,
		func() (proto.Message, error) {
			return &pb.MountISCSILunResponse{}, nil
		})

	response := rawResponse.(*pb.MountISCSILunResponse)
	return response, err
}

func (s *Server) ReportIScsiSendTargetPortals(ctx context.Context, request *pb.ReportIScsiSendTargetPortalsRequest) (*pb.ReportIScsiSendTargetPortalsResponse, error) {
	return nil, nil
}

func (s *Server) AddIScsiSendTargetPortal(ctx context.Context, request *pb.AddIScsiSendTargetPortalRequest) (*pb.AddIScsiSendTargetPortalResponse, error) {
	request.Descriptor()
	return nil, nil
}

func (s *Server) RemoveIScsiSendTargetPortal(ctx context.Context, request *pb.RemoveIScsiSendTargetPortalRequest) (*pb.RemoveIScsiSendTargetPortalResponse, error) {
	request.Descriptor()
	return nil, nil
}

func (s *Server) ReportIScsiTargets(ctx context.Context, request *pb.ReportIScsiTargetsRequest) (*pb.ReportIScsiTargetsResponse, error) {
	request.ProtoMessage()
	return nil, nil
}

func (s *Server) LoginIscsiTarget(ctx context.Context, request *pb.LoginIscsiTargetRequest) (*pb.LoginIscsiTargetResponse, error) {
	return nil, nil
}

func (s *Server) LogoutIScsiTarget(ctx context.Context, request *pb.LogoutIScsiTargetRequest) (*pb.LogoutIScsiTargetResponse, error) {
	return nil, nil
}

func (s *Server) GetIScsiSessionList(ctx context.Context, request *pb.GetIScsiSessionListRequest) (*pb.GetIScsiSessionListResponse, error) {
	return nil, nil
}

func (s *Server) GetDevicesForIScsiSession(ctx context.Context, request *pb.GetDevicesForIScsiSessionRequest) (*pb.GetDevicesForIScsiSessionResponse, error) {
	return nil, nil
}
