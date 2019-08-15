package iscsi

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
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

	socket, err := uint32SocketToUint16Pointer(in.Socket)
	if err != nil {
		return nil, err
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

func ConnectionInfoToGRPC(in *iscsidsc.ConnectionInfo) *pb.IscsiConnectionInfo {
	if in == nil {
		return nil
	}

	return &pb.IscsiConnectionInfo{
		ConnectionId:     ConnectionIDToGRPC(&in.ConnectionID),
		InitiatorAddress: in.InitiatorAddress,
		TargetAddress:    in.TargetAddress,
		InitiatorSocket:  uint32(in.InitiatorSocket),
		TargetSocket:     uint32(in.TargetSocket),
		CID:              in.CID[:],
	}
}

func ConnectionInfoFromGRPC(in *pb.IscsiConnectionInfo) (*iscsidsc.ConnectionInfo, error) {
	if in == nil {
		return nil, nil
	}

	initiatorSocket, err := uint32SocketToUint16(in.InitiatorSocket)
	if err != nil {
		return nil, err
	}
	targetSocket, err := uint32SocketToUint16(in.TargetSocket)
	if err != nil {
		return nil, err
	}

	var cid [2]byte
	if len(in.CID) != 2 {
		return nil, fmt.Errorf("Error when deserializing connection ID %v: expected CID to be 2 bytes long VS actual length: %d", in.ConnectionId, len(in.CID))
	}
	copy(cid[:], in.CID)

	return &iscsidsc.ConnectionInfo{
		ConnectionID:     *ConnectionIDFromGRPC(in.ConnectionId),
		InitiatorAddress: in.InitiatorAddress,
		TargetAddress:    in.TargetAddress,
		InitiatorSocket:  initiatorSocket,
		TargetSocket:     targetSocket,
		CID:              cid,
	}, nil
}

func SessionInfoToGRPC(in *iscsidsc.SessionInfo) *pb.IscsiSessionInfo {
	if in == nil {
		return nil
	}

	var connections []*pb.IscsiConnectionInfo
	if in.Connections != nil {
		connections = make([]*pb.IscsiConnectionInfo, len(in.Connections))

		for i, connection := range in.Connections {
			connections[i] = ConnectionInfoToGRPC(&connection)
		}
	}

	return &pb.IscsiSessionInfo{
		SessionId:      SessionIDToGRPC(&in.SessionID),
		InitiatorName:  in.InitiatorName,
		TargetNodeName: in.TargetNodeName,
		TargetName:     in.TargetName,
		ISID:           in.ISID[:],
		TSID:           in.TSID[:],
		Connections:    connections,
	}
}

func SessionInfoFromGRPC(in *pb.IscsiSessionInfo) (*iscsidsc.SessionInfo, error) {
	if in == nil {
		return nil, nil
	}

	var connections []iscsidsc.ConnectionInfo
	if in.Connections != nil {
		connections = make([]iscsidsc.ConnectionInfo, len(in.Connections))

		for i, connIn := range in.Connections {
			connOut, err := ConnectionInfoFromGRPC(connIn)
			if err != nil {
				return nil, wrapSessionInfoError(in, err)
			}
			connections[i] = *connOut
		}
	}

	var isid [6]byte
	if len(in.ISID) != 6 {
		return nil, wrapSessionInfoError(in, fmt.Errorf("expected ISID to be 6 bytes long VS actual length: %d", len(in.ISID)))
	}
	copy(isid[:], in.ISID)

	var tsid [2]byte
	if len(in.TSID) != 2 {
		return nil, wrapSessionInfoError(in, fmt.Errorf("expected TSID to be 2 bytes long VS actual length: %d", len(in.TSID)))
	}
	copy(tsid[:], in.TSID)

	return &iscsidsc.SessionInfo{
		// TODO wkpo next from here
	}, nil
}

func wrapSessionInfoError(in *pb.IscsiSessionInfo, err error) error {
	return errors.Wrapf(err, "Error when deserializing session ID %v", in.SessionId)
}

func uint32SocketToUint16(in uint32) (uint16, error) {
	if in > math.MaxUint16 {
		return 0, fmt.Errorf("sockets cannot be greater than %d, got %d", math.MaxUint16, in)
	}

	return uint16(in), nil
}

func uint32SocketToUint16Pointer(in uint32) (*uint16, error) {
	if in == 0 {
		return nil, nil
	}

	out, err := uint32SocketToUint16(in)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
