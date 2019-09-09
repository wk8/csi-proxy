package server

import (
	"context"
	"fmt"
)

// TODO wkpo comment??

type Server struct {
	// TODO wkpo
}

// TODO wkpo centraliser ca somewhere?
type Version struct {
	// TODO wkpo
}

// TODO wkpo handle le case int32? something like if version < blah cast to int32,
// OU TODO wkpo est ce que ca devrait plutot aller dans la conversion ca...??
func (s *Server) ComputeDouble(ctx context.Context, request *ComputeDoubleRequest, version Version) (*ComputeDoubleResponse, error) {
	in := request.Input64
	out := 2 * in

	if sign(in) != sign(out) {
		// overflow
		return nil, fmt.Errorf("int64 overflow with input: %d", in)
	}

	return &ComputeDoubleResponse{
		Response: out,
	}, nil
}

func sign(x int64) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func (s *Server) TellMeAPoem(ctx context.Context, request *TellMeAPoemRequest, version Version) (*TellMeAPoemResponse, error) {
	// TODO wkpo
	return nil, nil
}
