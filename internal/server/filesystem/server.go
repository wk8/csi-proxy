package dummy

import (
	"context"

	corev1 "k8s.io/api/core/v1"

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
			CmdletName: string(corev1.PodRunning),
			Code:       12,
			Message:    "hey there " + request.Path,
		},
	}, nil
}
