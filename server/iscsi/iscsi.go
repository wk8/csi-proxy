package iscsi

import (
	"context"

	pb "github.com/kubernetes-csi/csi-proxy/api"
)

// TODO wkpo see https://blog.golang.org/generate ? add a //go:generate comment?

type IscsiServer struct {
}

// TODO wkpo NEXT! lo que pasa quand on renvoie une error?
func (s *IscsiServer) MountISCSILun(ctx context.Context, request *pb.MountISCSILunRequest) (*pb.MountISCSILunResponse, error) {
	// TODO wkpo
	return nil, nil
}

func (s *IscsiServer) ReportIScsiSendTargetPortals(ctx context.Context, request *pb.ReportIScsiSendTargetPortalsRequest) (*pb.ReportIScsiSendTargetPortalsResponse, error) {
	return nil, nil
}

func (s *IscsiServer) AddIScsiSendTargetPortal(ctx context.Context, request *pb.AddIScsiSendTargetPortalRequest) (*pb.AddIScsiSendTargetPortalResponse, error) {
	return nil, nil
}

func (s *IscsiServer) RemoveIScsiSendTargetPortal(ctx context.Context, request *pb.RemoveIScsiSendTargetPortalRequest) (*pb.RemoveIScsiSendTargetPortalResponse, error) {
	return nil, nil
}

func (s *IscsiServer) ReportIScsiTargets(ctx context.Context, request *pb.ReportIScsiTargetsRequest) (*pb.ReportIScsiTargetsResponse, error) {
	return nil, nil
}

func (s *IscsiServer) LoginIscsiTarget(ctx context.Context, request *pb.LoginIscsiTargetRequest) (*pb.LoginIscsiTargetResponse, error) {
	return nil, nil
}

func (s *IscsiServer) LogoutIScsiTarget(ctx context.Context, request *pb.LogoutIScsiTargetRequest) (*pb.LogoutIScsiTargetResponse, error) {
	return nil, nil
}

func (s *IscsiServer) GetIScsiSessionList(ctx context.Context, request *pb.GetIScsiSessionListRequest) (*pb.GetIScsiSessionListResponse, error) {
	return nil, nil
}

func (s *IscsiServer) GetDevicesForIScsiSession(ctx context.Context, request *pb.GetDevicesForIScsiSessionRequest) (*pb.GetDevicesForIScsiSessionResponse, error) {
	return nil, nil
}
