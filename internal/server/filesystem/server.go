package dummy

import (
	"context"

	"github.com/wk8/csi-proxy/api"
	"github.com/wk8/csi-proxy/internal/apiversion"
	"github.com/wk8/csi-proxy/internal/server/filesystem/internal"
)

type Server struct{}

func (s *Server) PathExists(ctx context.Context, request *internal.PathExistsRequest, version apiversion.Version) (*internal.PathExistsResponse, error) {
	// FIXME: actually implement this!
	return &internal.PathExistsResponse{
		Success: false,
		CmdletError: &api.CmdletError{
			CmdletName: "dummy",
			Code:       12,
			Message:    "hey there " + request.Path,
		},
	}, nil
}
