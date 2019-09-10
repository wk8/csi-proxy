package internal

import (
	"github.com/kubernetes-csi/csi-proxy/internal/apiversion"
)

const (
	// PipePrefix is the prefix for Windows named pipes' names
	// TODO wkpo need public?
	PipePrefix = `\\.\\pipe\\`

	// CsiProxyNamedPipePrefix is the prefix for the named pipes the proxy creates.
	// The suffix will be the version, e.g. "\\.\\pipe\\csi-proxy-v1", "\\.\\pipe\\csi-proxy-v2alpha1", etc.
	// TODO wkpo need public?
	CsiProxyNamedPipePrefix = "csi-proxy-"
)

func PipePath(apiGroupName string, apiVersion apiversion.Version) string {
	return PipePrefix + CsiProxyNamedPipePrefix + apiGroupName + "-" + apiVersion.String()
}
