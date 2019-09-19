package internal

import (
	"github.com/wk8/csi-proxy/client/api"
)

type PathExistsRequest struct {
	// The path to check in the host filesystem.
	Path string
}

type PathExistsResponse struct {
	Success     bool
	CmdletError *api.CmdletError
	Exists      bool
}
