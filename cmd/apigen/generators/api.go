package generators

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"k8s.io/gengo/args"
	gengogenerator "k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	gengotypes "k8s.io/gengo/types"

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

// TODO wkpo comment?
type groupDefinition struct {
	name          string
	serverBasePkg string
	clientBasePkg string
	versions      []*gengotypes.Package
}

func newGroupDefinition(name string, versions ...*gengotypes.Package) *groupDefinition {
	return &groupDefinition{
		name:          name,
		serverBasePkg: defaultServerBasePkg,
		clientBasePkg: defaultClientBasePkg,
		versions:      versions,
	}
}

func (d *groupDefinition) String() string {
	if d == nil {
		return "<nil>"
	}

	result := fmt.Sprintf("{name: %q", d.name)
	if d.serverBasePkg != "" && d.serverBasePkg != defaultServerBasePkg {
		result += fmt.Sprintf(", serverBasePkg: %q", d.serverBasePkg)
	}
	if d.clientBasePkg != "" && d.clientBasePkg != defaultClientBasePkg {
		result += fmt.Sprintf(", clientBasePkg: %q", d.clientBasePkg)
	}
	if len(d.versions) != 0 {
		result += ", versions: ["
		for _, version := range d.versions {
			if version == nil {
				result += "<nil> "
			} else {
				result += version.Name + " "
			}
		}
		result = result[:len(result)-1] + "]"
	}
	return result + "}"
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	// TODO wkpo? which are used?
	return namer.NameSystems{
		"public": namer.NewPublicNamer(0),
		"raw":    namer.NewRawNamer("", nil),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func Packages(context *gengogenerator.Context, arguments *args.GeneratorArgs) (packages gengogenerator.Packages) {
	// find API group definitions
	groups := findAPIGroupDefinitions(context)
	logrus.Debugf("Found API groups: %v", groups)

	for _, group := range groups {
		packages = append(packages, generatorPackagesForGroup(group)...)
	}
	return
}

// TODO wkpo move to EOF
// TODO wkpo header , etc?
// TODO wkpo open an issue against gengo to allow returning an error here?
func generatorPackagesForGroup(group *groupDefinition) gengogenerator.Packages {
	serverInterfaceName, serverCallbacks, errors := findServerCallbacksForGroup(group)
	if len(errors) != 0 {
		// TODO wkpo
		for _, err := range errors {
			logrus.Errorf("wkpo error: %v", err)
		}
		logrus.Fatalf("aborting")
	}
	// TODO wkpo pas used??
	logrus.Debugf("wkpo %s", serverInterfaceName)
	logrus.Debugf("wkpo on a trouve %v callbacks", len(serverCallbacks))

	packages := gengogenerator.Packages{
		&gengogenerator.DefaultPackage{
			PackageName: internal.SnakeCaseToPackageName(group.name),
			PackagePath: fmt.Sprintf("%s/%s", group.serverBasePkg, group.name),

			// TODO wkpo generators?
			// api_group_generated.go        => def
			// server.go (if doesn't exist)  => def + callbacks
			GeneratorList: []gengogenerator.Generator{
				gengogenerator.DefaultGen{
					OptionalName: "wkpo",
					OptionalBody: []byte("// coucou"),
				},
			},
		},

		&gengogenerator.DefaultPackage{
			PackageName: "internal",
			PackagePath: fmt.Sprintf("%s/%s/internal", group.serverBasePkg, group.name),

			// TODO wkpo generators?
			// types.go (if doesn't exist)  => def + types (from callbacks?)
			// types_generated.go           => def + callbacks
			GeneratorList: []gengogenerator.Generator{
				&typesGeneratedGenerator{
					DefaultGen: gengogenerator.DefaultGen{
						OptionalName: "wkpo_types_generated",
					},
					serverCallbacks: serverCallbacks,
				},
			},
		},
	}

	for _, version := range group.versions {
		packages = append(packages,
			&gengogenerator.DefaultPackage{
				PackageName: version.Name,
				PackagePath: fmt.Sprintf("%s/%s/internal/%s", group.serverBasePkg, group.name, version.Name),

				// TODO wkpo generators?
				// conversion_generated.go => types!
				// server_generated.go
				// conversion.go (if doesn't exist)
				GeneratorList: []gengogenerator.Generator{
					gengogenerator.DefaultGen{
						OptionalName: "wkpo",
						OptionalBody: []byte("// coucou"),
					},
				},
			},

			&gengogenerator.DefaultPackage{
				PackageName: version.Name,
				PackagePath: fmt.Sprintf("%s/%s/%s", group.clientBasePkg, group.name, version.Name),

				// TODO wkpo generators?
				// conversion_generated.go
				// server_generated.go
				// conversion.go (if doesn't exist)
				GeneratorList: []gengogenerator.Generator{
					gengogenerator.DefaultGen{
						OptionalName: "wkpo",
						OptionalBody: []byte("// coucou"),
					},
				},
			},
		)
	}

	return packages
}

// TODO wkpo move to EOF
// TODO wkpo comment?
type versionedType struct {
	*gengotypes.Type
	version *gengotypes.Package
}

// TODO wkpo move to EOF
// TODO wkpo comment? explain that the default name doesn't really offer what we need...?
type namedCallback struct {
	*gengotypes.Type
	Name string
}

// TODO wkpo move to EOF
// TODO wkpo comment?
// TODO wkpo name?
func findServerCallbacksForGroup(group *groupDefinition) (serverInterfaceName string, orderedCallbacks []*namedCallback, errors []error) {
	// that is the name of the server interface for this API group that we expect to find
	// in each version's package
	serverInterfaceName = fmt.Sprintf("%sServer", strcase.ToCamel(group.name))

	callbacks := make(map[string]*versionedType)

	for _, version := range group.versions {
		if serverInterface, present := version.Types[serverInterfaceName]; present {
			if serverInterface.Kind == gengotypes.Interface {
				errors = append(errors, mergeCallbacks(callbacks, serverInterface.Methods, version, serverInterfaceName)...)
			} else {
				errors = append(errors, fmt.Errorf("Type %s in package %s should be an interface, it actually is a %s",
					serverInterfaceName, version.Path, serverInterface.Kind))
			}
		} else {
			errors = append(errors, fmt.Errorf("Did not find interface %s in package %s",
				serverInterfaceName, version.Path))
		}
	}

	orderedCallbacks = make([]*namedCallback, len(callbacks))
	i := 0
	for name, callback := range callbacks {
		orderedCallbacks[i] = &namedCallback{
			Type: callback.Type,
			Name: name,
		}
		i++
	}
	sort.Slice(orderedCallbacks, func(i, j int) bool {
		return orderedCallbacks[i].Name < orderedCallbacks[j].Name
	})

	return
}

// TODO wkpo move to EOF
// TODO wkpo comment?
func mergeCallbacks(callbacks map[string]*versionedType, methods map[string]*gengotypes.Type, version *gengotypes.Package, interfaceName string) (errors []error) {
	for name, method := range methods {
		if previousCallback, present := callbacks[name]; present {
			if !similarSignatures(previousCallback.Signature, method.Signature) {
				// TODO wkpo add a signature string to the error message? since we need to generate it anyway...?
				errors = append(errors, fmt.Errorf("Method %s on interface %s in package %s has a different signature than in package %s",
					name, interfaceName, version.Name, previousCallback.version.Name))
			}
		} else {
			callbacks[name] = &versionedType{
				Type:    method,
				version: version,
			}
		}
	}

	return
}

// TODO wkpo move to EOF
// TODO wkpo comment?
// TODO wkpo unit tests on this
func similarSignatures(s1, s2 *gengotypes.Signature) bool {
	if s1 == nil || s2 == nil {
		return s1 == s2
	}

	return similarTypes(s1.Receiver, s2.Receiver) &&
		similarTypeLists(s1.Parameters, s2.Parameters) &&
		similarTypeLists(s1.Results, s2.Results)
}

// TODO wkpo move to EOF
// TODO wkpo comment?
func similarTypes(t1, t2 *gengotypes.Type) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}

	return t1.Kind == t2.Kind &&
		internal.SplitLast(t1.Name.Name, ".") == internal.SplitLast(t2.Name.Name, ".") &&
		similarTypes(t1.Elem, t2.Elem)
}

// TODO wkpo move to EOF
// TODO wkpo comment?
func similarTypeLists(l1, l2 []*gengotypes.Type) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, t1 := range l1 {
		if !similarTypes(t1, l2[i]) {
			return false
		}
	}
	return true
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
func findAPIGroupDefinitions(context *gengogenerator.Context) map[string]*groupDefinition {
	pkgPaths := context.Inputs

	// first, re-order the inputs by lengths, so that we always process parent packages first
	sort.Slice(pkgPaths, func(i, j int) bool {
		return len(pkgPaths[i]) < len(pkgPaths[j])
	})

	groups := make(map[string]*groupDefinition)

	for _, pkgPath := range pkgPaths {
		logrus.Debugf("Considering input %s", pkgPath)

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
				logrus.Fatalf("Unexpected go package %q, should be of the form \"%s/<version>\", where <version> should be a valid API group version identifier",
					pkg.Name, parentPkgPath)
			}
			logrus.Debugf("Found version %q for API group %q", pkg.Name, parts[len(parts)-2])
			definition.versions = append(definition.versions, pkg)
		}
	}

	return groups
}

// buildAPIGroupDefinitionFromDocComment looks for a +csi-proxy-gen comment in the package's
// doc.go file, and if it finds one build the corresponding API definition.
func buildAPIGroupDefinitionFromDocComment(pkgPath string, pkg *gengotypes.Package, groups map[string]*groupDefinition) bool {
	for _, commentLine := range pkg.Comments {
		if matches := markerCommentRegex.FindStringSubmatch(commentLine); matches != nil {
			definition := newGroupDefinition(pkg.Name)

			if len(matches) >= 2 {
				for _, option := range strings.Split(matches[1], ",") {
					parts := strings.Split(option, ":")
					if len(parts) != 2 {
						logrus.Fatalf("Malformed option for package %q, options should be of the form \"<name>:<value>\", found %q",
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
						logrus.Fatalf("Unknown option %q for package %q", name, pkgPath)
					}
				}

			}

			logrus.Debugf("Found API group %q", definition.name)
			groups[pkg.Path] = definition
			return true
		}
	}

	return false
}

// buildCanonicalAPIGroupDefinition builds group definitions for API groups under client/api
// Since these API group's directories don't need to contain go files, and as such are not necessarily go packages,
// here we actually directly process versions' packages directly.
func buildCanonicalAPIGroupDefinition(pkgPath string, pkg *gengotypes.Package, groups map[string]*groupDefinition) {
	groupNameAndVersion := strings.TrimPrefix(pkgPath, internal.CSIProxyAPIPath)
	parts := strings.Split(groupNameAndVersion, "/")

	if len(parts) == 1 {
		// means it's the group's root directory, no need to process it at this time,
		// we'll get around to it when we process its versions' packages
		return
	}

	if len(parts) != 2 || !apiversion.IsValidVersion(parts[1]) {
		logrus.Fatalf("Unexpected go package %q, should be of the form \"%s<api_group_name>/<version>\", where <version> should be a valid API group version identifier",
			pkgPath, internal.CSIProxyAPIPath)
	}

	groupPath := internal.CSIProxyAPIPath + parts[0]
	if definition, present := groups[groupPath]; present {
		definition.versions = append(definition.versions, pkg)
	} else {
		logrus.Debugf("Found API group %q", parts[0])
		groups[groupPath] = newGroupDefinition(parts[0], pkg)
	}
	logrus.Debugf("Found version %q for API group %q", parts[1], parts[0])
}
