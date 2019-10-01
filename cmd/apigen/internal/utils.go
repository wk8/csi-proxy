package internal

import (
	"strings"

	"k8s.io/gengo/types"
)

// TODO wkpo unit tests!

// TODO wkpo comments?
const (
	CSIProxyRootPath = "github.com/kubernetes-csi/csi-proxy"
	CSIProxyAPIPath  = CSIProxyRootPath + "/client/api/"

	PkgPlaceholder = "__PKG_PLACEHOLDER__"
)

// CanonicalizePkgPath ensures package paths are consistent.
func CanonicalizePkgPath(pkgPath string) string {
	return strings.TrimSuffix(pkgPath, "/")
}

// ToPackageName turns a snake case string into a go package name.
func SnakeCaseToPackageName(name string) string {
	return strings.ReplaceAll(name, "_", "")
}

// ReplaceTypesPackage return a new type, equal to t except moved from package
// pkg to package newPkg (and same for other types referenced by t).
// t itself remains unchanged.
func ReplaceTypesPackage(t *types.Type, pkg, newPkg string) *types.Type {
	return replaceTypesPackage(t, normalizePkg(pkg), normalizePkg(newPkg), make(map[*types.Type]*types.Type))
}

func normalizePkg(pkg string) string {
	if pkg != "" && !strings.HasSuffix(pkg, ".") {
		pkg += "."
	}
	return pkg
}

func replaceTypesPackage(t *types.Type, pkg, newPkg string, visited map[*types.Type]*types.Type) *types.Type {
	if t == nil {
		return nil
	}
	if result, present := visited[t]; present {
		return result
	}

	result := &types.Type{
		Name: types.Name{
			Name:    strings.ReplaceAll(t.Name.Name, pkg, newPkg),
			Package: strings.ReplaceAll(t.Name.Package, pkg, newPkg),
		},
		Kind:                      t.Kind,
		CommentLines:              t.CommentLines,
		SecondClosestCommentLines: t.SecondClosestCommentLines,
		Members:                   t.Members,
	}
	visited[t] = result

	result.Elem = replaceTypesPackage(t.Elem, pkg, newPkg, visited)
	result.Key = replaceTypesPackage(t.Key, pkg, newPkg, visited)
	result.Underlying = replaceTypesPackage(t.Underlying, pkg, newPkg, visited)

	if t.Methods != nil {
		methods := make(map[string]*types.Type)
		for k, v := range t.Methods {
			methods[k] = replaceTypesPackage(v, pkg, newPkg, visited)
		}
		result.Methods = methods
	}

	var signature *types.Signature
	if t.Signature != nil {
		signature = &types.Signature{
			Receiver:     replaceTypesPackage(t.Signature.Receiver, pkg, newPkg, visited),
			Parameters:   replaceTypesSlicePackage(t.Signature.Parameters, pkg, newPkg, visited),
			Results:      replaceTypesSlicePackage(t.Signature.Results, pkg, newPkg, visited),
			Variadic:     t.Signature.Variadic,
			CommentLines: t.Signature.CommentLines,
		}
		result.Signature = signature
	}

	return result
}

func replaceTypesSlicePackage(ts []*types.Type, pkg, newPkg string, visited map[*types.Type]*types.Type) []*types.Type {
	result := make([]*types.Type, len(ts))
	for i, t := range ts {
		result[i] = replaceTypesPackage(t, pkg, newPkg, visited)
	}
	return result
}
