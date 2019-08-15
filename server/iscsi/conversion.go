package iscsi

import (
	"fmt"
	"math"

	iscsidsc "github.com/wk8/go-win-iscsidsc"

	pb "github.com/kubernetes-csi/csi-proxy/api"
)

// TODO: this could be auto-generated, write a lib for this?

// TODO wkpo unit tests on all this shit

func PortalToGRPC(in *iscsidsc.Portal) *pb.IscsiPortal {
	if in == nil {
		return nil
	}

	out := &pb.IscsiPortal{
		SymbolicName: in.SymbolicName,
		Address:      in.Address,
	}
	if in.Socket != nil {
		out.Socket = uint32(*in.Socket)
	}
	return out
}

func PortalFromGRPC(in *pb.IscsiPortal) (*iscsidsc.Portal, error) {
	if in == nil {
		return nil, nil
	}

	if in.Socket > math.MaxUint16 {
		return nil, fmt.Errorf("portal sockets cannot be greater than %d, got %d", math.MaxUint16, in.Socket)
	}

	var socket *uint16
	if in.Socket > 0 {
		socketInt16 := uint16(in.Socket)
		socket = &socketInt16
	}

	return &iscsidsc.Portal{
		SymbolicName: in.SymbolicName,
		Address:      in.Address,
		Socket:       socket,
	}, nil
}

func LoginOptionsToGRPC(in *iscsidsc.LoginOptions) *pb.IscsiLoginOptions {
	if in == nil {
		return nil
	}

	out := &pb.IscsiLoginOptions{
		LoginFlags: uint32(in.LoginFlags),
	}

	if in.AuthType != nil {
		out.AuthType = pb.IscsiAuthType(*in.AuthType)
	}
	if in.HeaderDigest != nil {
		out.HeaderDigest = pb.IscsiDigestType(*in.HeaderDigest)
	}
	if in.DataDigest != nil {
		out.DataDigest = pb.IscsiDigestType(*in.DataDigest)
	}
	if in.MaximumConnections != nil {
		out.MaximumConnections = *in.MaximumConnections
	}
	if in.DefaultTime2Wait != nil {
		out.DefaultTime_2Wait = *in.DefaultTime2Wait
	}
	if in.DefaultTime2Retain != nil {
		out.DefaultTime_2Retain = *in.DefaultTime2Retain
	}
	if in.Username != nil {
		out.Username = *in.Username
	}
	if in.Password != nil {
		out.Password = *in.Password
	}

	return out
}

func LoginOptionsFromGRPC(in *pb.IscsiLoginOptions) *iscsidsc.LoginOptions {
	if in == nil {
		return nil
	}

	out := &iscsidsc.LoginOptions{
		LoginFlags: iscsidsc.LoginFlags(in.LoginFlags),
	}

	if in.AuthType != pb.IscsiAuthType_NoAuth {
		authType := iscsidsc.AuthType(in.AuthType)
		out.AuthType = &authType
	}
	if in.HeaderDigest != pb.IscsiDigestType_NoDigest {
		headerDigest := iscsidsc.DigestType(in.HeaderDigest)
		out.HeaderDigest = &headerDigest
	}
	if in.DataDigest != pb.IscsiDigestType_NoDigest {
		dataDigest := iscsidsc.DigestType(in.DataDigest)
		out.DataDigest = &dataDigest
	}
	if in.MaximumConnections != 0 {
		out.MaximumConnections = &in.MaximumConnections
	}
	if in.DefaultTime_2Wait != 0 {
		out.DefaultTime2Wait = &in.DefaultTime_2Wait
	}
	if in.DefaultTime_2Retain != 0 {
		out.DefaultTime2Retain = &in.DefaultTime_2Retain
	}
	if in.Username != "" {
		out.Username = &in.Username
	}
	if in.Password != "" {
		out.Password = &in.Password
	}

	return out
}

func PortalInfoToGRPC(in *iscsidsc.PortalInfo) *pb.IscsiPortalInfo {
	if in == nil {
		return nil
	}

	return &pb.IscsiPortalInfo{
		Portal:              PortalToGRPC(&in.Portal),
		InitiatorName:       in.InitiatorName,
		InitiatorPortNumber: in.InitiatorPortNumber,
		SecurityFlags:       uint64(in.SecurityFlags),
		LoginOptions:        LoginOptionsToGRPC(&in.LoginOptions),
	}
}

func PortalInfoFromGRPC(in *pb.IscsiPortalInfo) (*iscsidsc.PortalInfo, error) {
	if in == nil {
		return nil, nil
	}

	portal, err := PortalFromGRPC(in.Portal)
	if err != nil {
		return nil, err
	}

	return &iscsidsc.PortalInfo{
		Portal:              *portal,
		InitiatorName:       in.InitiatorName,
		InitiatorPortNumber: in.InitiatorPortNumber,
		SecurityFlags:       iscsidsc.SecurityFlags(in.SecurityFlags),
		LoginOptions:        *LoginOptionsFromGRPC(in.LoginOptions),
	}, nil
}

func SessionIDToGRPC(in *iscsidsc.SessionID) *pb.IscsiID {
	if in == nil {
		return nil
	}

	return &pb.IscsiID{
		AdapterUnique:   in.AdapterUnique,
		AdapterSpecific: in.AdapterSpecific,
	}
}

func SessionIDFromGRPC(in *pb.IscsiID) *iscsidsc.SessionID {
	if in == nil {
		return nil
	}

	return &iscsidsc.SessionID{
		AdapterUnique:   in.AdapterUnique,
		AdapterSpecific: in.AdapterSpecific,
	}
}

func ConnectionIDToGRPC(in *iscsidsc.ConnectionID) *pb.IscsiID {
	if in == nil {
		return nil
	}

	return &pb.IscsiID{
		AdapterUnique:   in.AdapterUnique,
		AdapterSpecific: in.AdapterSpecific,
	}
}

func ConnectionIDFromGRPC(in *pb.IscsiID) *iscsidsc.ConnectionID {
	if in == nil {
		return nil
	}

	return &iscsidsc.ConnectionID{
		AdapterUnique:   in.AdapterUnique,
		AdapterSpecific: in.AdapterSpecific,
	}
}

func SessionInfoToGRPC(in *iscsidsc.SessionInfo) *pb.IscsiSessionInfo {
	return
}

func SessionInfoFromGRPC(in *pb.IscsiSessionInfo) *iscsidsc.SessionInfo {

}
