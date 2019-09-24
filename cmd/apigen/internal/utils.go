package internal

import (
	"strings"
)

// TODO wkpo comments?
const (
	CSIProxyRootPath = "github.com/kubernetes-csi/csi-proxy"
	CSIProxyAPIPath  = CSIProxyRootPath + "/client/api/"
)

// CanonicalizePkgPath ensures package paths are consistent.
func CanonicalizePkgPath(pkgPath string) string {
	return strings.TrimSuffix(pkgPath, "/")
}
