package internal

import (
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
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

func PipePathForApiVersion(apiVersion apiversions.Version) string {
	return PipePrefix + CsiProxyNamedPipePrefix + apiVersion.String()
}
