package generators

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/internal"
	"github.com/wk8/go-ordered-map"
	"k8s.io/gengo/types"
	"k8s.io/klog"
)

// TODO wkpo comment?
type groupDefinition struct {
	name          string
	serverBasePkg string
	clientBasePkg string
	versions      []*apiVersion
	// serverCallbacks maps callbacks to their definitions, with all the
	// versioned types replaced with internal types.
	serverCallbacks *orderedmap.OrderedMap
}

// TODO wkpo comment?
type apiVersion struct {
	*types.Package
	// topLevelTypes maps type names to their definitions
	topLevelTypes *orderedmap.OrderedMap
	// serverCallbacks maps callbacks to their definitions
	// TODO wkpo used?
	serverCallbacks *orderedmap.OrderedMap
}

func newGroupDefinition(name string) *groupDefinition {
	return &groupDefinition{
		name:            name,
		serverBasePkg:   defaultServerBasePkg,
		clientBasePkg:   defaultClientBasePkg,
		serverCallbacks: orderedmap.New(),
	}
}

func (d *groupDefinition) addVersion(versionPkg *types.Package) {
	serverInterface, present := versionPkg.Types[d.serverInterfaceName()]
	if !present {
		klog.Fatalf("did not find interface %s in package %s", d.serverInterfaceName(), versionPkg.Path)
	}
	if serverInterface.Kind != types.Interface {
		klog.Fatalf("type %s in package %s should be an interface, it actually is a %s",
			d.serverInterfaceName(), versionPkg.Path, serverInterface.Kind)
	}

	version := &apiVersion{
		Package:         versionPkg,
		topLevelTypes:   orderedmap.New(),
		serverCallbacks: orderedmap.New(),
	}
	d.versions = append(d.versions, version)

	for callbackName, versionedCallback := range serverInterface.Methods {
		version.serverCallbacks.Set(callbackName, versionedCallback)

		serverCallback := internal.ReplaceTypesPackage(versionedCallback, versionPkg.Path, internal.PkgPlaceholder)
		klog.Infof("wkpo bordel de merde after change: %s", serverCallback)

		if previousCallbackRaw, alreadyPresent := d.serverCallbacks.Get(callbackName); alreadyPresent {
			previousCallback := previousCallbackRaw.(*types.Type)
			if serverCallback.String() != previousCallback.String() {
				errorMsg := fmt.Sprintf("Endpoint %s in API group %s inconsistent across versions:", callbackName, d.name)
				for _, vsn := range d.versions {
					if vsnCallback, present := vsn.serverCallbacks.Get(callbackName); present {
						errorMsg += fmt.Sprintf("\n  - in version %s: %s", vsn.Name, vsnCallback.(*types.Type))
					}
				}
				errorMsg += fmt.Sprintf("\nYields 2 different signatures for the internal server callback:\n%s\nand\n%s",
					previousCallback, serverCallback)
				klog.Fatalf(errorMsg)
			}
		} else {
			d.serverCallbacks.Set(callbackName, serverCallback)
		}
	}
}

// serverInterfaceName is the name of the server interface for this API group
// that we expect to find in each version's package.
func (d *groupDefinition) serverInterfaceName() string {
	return fmt.Sprintf("%sServer", strcase.ToCamel(d.name))
}

// serverPkg returns the path of the server package, e.g.
// github.com/kubernetes-csi/csi-proxy/internal/server/<api_group_name>
func (d *groupDefinition) serverPkg() string {
	return fmt.Sprintf("%s/%s", d.serverBasePkg, d.name)
}

// internalServerPkg returns the path of the internal server package, e.g.
// github.com/kubernetes-csi/csi-proxy/internal/server/<api_group_name>/internal
func (d *groupDefinition) internalServerPkg() string {
	return fmt.Sprintf("%s/%s/internal", d.serverBasePkg, d.name)
}

// versionedServerPkg returns the path of the versioned server package, e.g.
// github.com/kubernetes-csi/csi-proxy/internal/server/<api_group_name>/internal/<version>
func (d *groupDefinition) versionedServerPkg(version string) string {
	return fmt.Sprintf("%s/%s/internal/%s", d.serverBasePkg, d.name, version)
}

// versionedClientPkg returns the path of the versioned client package, e.g.
// github.com/kubernetes-csi/csi-proxy/client/groups/<api_group_name>/<version>
func (d *groupDefinition) versionedClientPkg(version string) string {
	return fmt.Sprintf("%s/%s/%s", d.clientBasePkg, d.name, version)
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
