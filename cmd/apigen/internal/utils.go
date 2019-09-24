package internal

import (
	"strings"
)

// TODO wkpo unit tests!

// TODO wkpo comments?
const (
	CSIProxyRootPath = "github.com/kubernetes-csi/csi-proxy"
	CSIProxyAPIPath  = CSIProxyRootPath + "/client/api/"
)

// CanonicalizePkgPath ensures package paths are consistent.
func CanonicalizePkgPath(pkgPath string) string {
	return strings.TrimSuffix(pkgPath, "/")
}

// ToPackageName turns a snake case string into a go package name.
func SnakeCaseToPackageName(name string) string {
	return strings.ReplaceAll(name, "_", "")
}

// SplitLast returns the last item in the slice returned by strings.Split(s, separator).
func SplitLast(s, separator string) string {
	parts := strings.Split(s, separator)
	return parts[len(parts)-1]
}
