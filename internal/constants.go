package internal

const (
	// PipePrefix is the prefix for Windows named pipes' names
	PipePrefix = `\\.\\pipe\\`

	// CsiProxyNamedPipePrefix is the prefix for the named pipes the proxy creates.
	// The suffix will be the version, e.g. "\\.\\pipe\\csi-proxy-v1", "\\.\\pipe\\csi-proxy-v2alpha1", etc.
	CsiProxyNamedPipePrefix = "csi-proxy-"

	// KnownApiVersions are the versions of the CSI proxy API that the current revision of the repo
)
