package internal

import (
	"github.com/kubernetes-csi/csi-proxy/client/api"
)

type PathExistsRequest struct {
	// The path to check in the host file_system.
	Path string
}

type PathExistsResponse struct {
	Success     bool
	CmdletError *api.CmdletError
	Exists      bool
}
