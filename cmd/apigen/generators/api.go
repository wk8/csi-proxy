package generators

// TODO wkpo remove logrus from go.mod/sum...?
// TODO wkpo logrus => klog!

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"k8s.io/gengo/args"
	// TODO wkpo
	// conversiongenerator "k8s.io/gengo/examples/conversion-gen/generators/generator"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog"

	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/internal"
)

// TODO wkpo save a SHA of source files, and don't regenerate if not needed? possible?

// TODO wkpo pkg comment?

// TODO wkpo PR on gengo to allow generators to say they don't want to overwrite a file?

// TODO wkpo comments?
const (
	defaultServerBasePkg = internal.CSIProxyRootPath + "/internal/server"
	defaultClientBasePkg = internal.CSIProxyRootPath + "/client/groups"
	markerComment        = "+csi-proxy-gen"
)

// TODO wkpo comment?
var markerCommentRegex = regexp.MustCompile(`^\s*(?://)?\s*` + regexp.QuoteMeta(markerComment) + `\s*=?([^\n]*)?$`)

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	// TODO wkpo? which are used?
	return namer.NameSystems{
		"public":  namer.NewPublicNamer(0),
		"raw":     namer.NewRawNamer("", nil),
		"private": namer.NewPrivateNamer(0),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) (packages generator.Packages) {
	// find API group definitions
	groups := findAPIGroupDefinitions(context)
	klog.V(5).Infof("Found API groups: %v", groups)

	for _, group := range groups {
		packages = append(packages, generatorPackagesForGroup(group)...)
	}
	return
}

// TODO wkpo move to EOF
// TODO wkpo header , etc?
// TODO wkpo open an issue against gengo to allow returning an error here?
func generatorPackagesForGroup(group *groupDefinition) generator.Packages {
	packages := generator.Packages{
		&generator.DefaultPackage{
			PackageName: internal.SnakeCaseToPackageName(group.name),
			PackagePath: fmt.Sprintf("%s/%s", group.serverBasePkg, group.name),

			// TODO wkpo generators?
			// api_group_generated.go        => def
			// server.go (if doesn't exist)  => def + callbacks
			GeneratorList: []generator.Generator{
				generator.DefaultGen{
					OptionalName: "wkpo",
					OptionalBody: []byte("// coucou"),
				},
			},
		},

		&generator.DefaultPackage{
			PackageName: "internal",
			PackagePath: fmt.Sprintf("%s/%s/internal", group.serverBasePkg, group.name),

			// TODO wkpo generators?
			// types.go (if doesn't exist)  => def + types (from callbacks?)
			GeneratorList: []generator.Generator{
				&typesGeneratedGenerator{
					DefaultGen: generator.DefaultGen{
						OptionalName: "types_generated",
					},
					groupDefinition: group,
				},
			},
		},
	}

	for _, version := range group.versions {
		packages = append(packages,
			&generator.DefaultPackage{
				PackageName: version.Name,
				PackagePath: fmt.Sprintf("%s/%s/internal/%s", group.serverBasePkg, group.name, version.Name),

				// TODO wkpo generators?
				// conversion_generated.go => types!
				// server_generated.go
				// conversion.go (if doesn't exist)
				GeneratorList: []generator.Generator{
					generator.DefaultGen{
						OptionalName: "wkpo",
						OptionalBody: []byte("// coucou"),
					},
				},
			},

			&generator.DefaultPackage{
				PackageName: version.Name,
				PackagePath: fmt.Sprintf("%s/%s/%s", group.clientBasePkg, group.name, version.Name),

				// TODO wkpo generators?
				// client_generated.go
				GeneratorList: []generator.Generator{
					generator.DefaultGen{
						OptionalName: "wkpo",
						OptionalBody: []byte("// coucou"),
					},
				},
			},
		)
	}

	return packages
}

// findAPIGroupDefinitions iterates over the context's list of package paths,
// and builds a map mapping API group paths to their definition.
// API group definitions are either:
// * subdirectories of client/api
// * or packages whose doc.go file contains a comment containing markerComment
// If the latter, the comment can also contain optional comma-separated options:
//  * groupName: defaults to package name
//  * serverBasePkg: defaults to defaultServerBasePkg
//  * clientBasePkg: defaults to defaultClientBasePkg
// for example,
// +csi-proxy-gen=groupName:dummy,serverBasePkg:github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/server,clientBasePkg:github.com/kubernetes-csi/csi-proxy/integrationtests/apigroups/client
func findAPIGroupDefinitions(context *generator.Context) map[string]*groupDefinition {
	pkgPaths := context.Inputs

	// first, re-order the inputs by lengths, so that we always process parent packages first
	sort.Slice(pkgPaths, func(i, j int) bool {
		return len(pkgPaths[i]) < len(pkgPaths[j])
	})

	groups := make(map[string]*groupDefinition)

	for _, pkgPath := range pkgPaths {
		klog.V(5).Infof("Considering input %s", pkgPath)

		pkg := context.Universe[pkgPath]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}

		if buildAPIGroupDefinitionFromDocComment(pkgPath, pkg, groups) {
			// found a +csi-proxy-gen comment in the package's doc.go file
			continue
		}

		if strings.HasPrefix(pkgPath, internal.CSIProxyAPIPath) {
			// part of the canonical API definitions, under client/api
			buildCanonicalAPIGroupDefinition(pkgPath, pkg, groups)
			continue
		}

		// is this package a version of an API group?
		parts := strings.Split(pkgPath, "/")
		parentPkgPath := strings.Join(parts[:len(parts)-1], "/")
		if definition, present := groups[parentPkgPath]; present {
			if !apiversion.IsValidVersion(pkg.Name) {
				klog.Fatalf("Unexpected go package %q, should be of the form \"%s/<version>\", where <version> should be a valid API group version identifier",
					pkg.Name, parentPkgPath)
			}
			klog.V(5).Infof("Found version %q for API group %q", pkg.Name, parts[len(parts)-2])
			definition.addVersion(pkg)
		}
	}

	return groups
}

// buildAPIGroupDefinitionFromDocComment looks for a +csi-proxy-gen comment in the package's
// doc.go file, and if it finds one build the corresponding API definition.
// TODO wkpo use types.ExtractCommentTags and the like?
func buildAPIGroupDefinitionFromDocComment(pkgPath string, pkg *types.Package, groups map[string]*groupDefinition) bool {
	for _, commentLine := range pkg.Comments {
		if matches := markerCommentRegex.FindStringSubmatch(commentLine); matches != nil {
			definition := newGroupDefinition(pkg.Name)

			if len(matches) >= 2 {
				for _, option := range strings.Split(matches[1], ",") {
					parts := strings.Split(option, ":")
					if len(parts) != 2 {
						klog.Fatalf("Malformed option for package %q, options should be of the form \"<name>:<value>\", found %q",
							pkgPath, option)
					}

					name := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					switch name {
					case "groupName":
						definition.name = value
					case "serverBasePkg":
						definition.serverBasePkg = internal.CanonicalizePkgPath(value)
					case "clientBasePkg":
						definition.clientBasePkg = internal.CanonicalizePkgPath(value)
					default:
						klog.Fatalf("Unknown option %q for package %q", name, pkgPath)
					}
				}

			}

			klog.V(5).Infof("Found API group %q", definition.name)
			groups[pkg.Path] = definition
			return true
		}
	}

	return false
}

// buildCanonicalAPIGroupDefinition builds group definitions for API groups under client/api
// Since these API group's directories don't need to contain go files, and as such are not necessarily go packages,
// here we actually directly process versions' packages directly.
func buildCanonicalAPIGroupDefinition(pkgPath string, pkg *types.Package, groups map[string]*groupDefinition) {
	groupNameAndVersion := strings.TrimPrefix(pkgPath, internal.CSIProxyAPIPath)
	parts := strings.Split(groupNameAndVersion, "/")

	if len(parts) == 1 {
		// means it's the group's root directory, no need to process it at this time,
		// we'll get around to it when we process its versions' packages
		return
	}

	if len(parts) != 2 || !apiversion.IsValidVersion(parts[1]) {
		klog.Fatalf("Unexpected go package %q, should be of the form \"%s<api_group_name>/<version>\", where <version> should be a valid API group version identifier",
			pkgPath, internal.CSIProxyAPIPath)
	}

	groupPath := internal.CSIProxyAPIPath + parts[0]
	definition, present := groups[groupPath]
	if !present {
		klog.V(5).Infof("Found API group %q", parts[0])
		definition = newGroupDefinition(parts[0])
		groups[groupPath] = definition
	}
	definition.addVersion(pkg)

	klog.V(5).Infof("Found version %q for API group %q", parts[1], parts[0])
}
