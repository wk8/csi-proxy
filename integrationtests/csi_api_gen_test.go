package integrationtests

import (
	"k8s.io/gengo/args"
	"testing"

	"github.com/kubernetes-csi/csi-proxy/cmd/csi-api-gen/generators"
)

// TODO wkpo capture output for all of em!
// TODO wkpo capture panics, too...!

// This tests API generator; more specifically, its main goal is to ensure
// that the API generator works as expected when creating a new API group.
// On top of this, the regular build checks that all checked-in generated files
// are up-to-date (i.e. consistent with the current generator).

func TestNewAPIGroup(t *testing.T) {
	// TODO wkpo!
	generators.Execute("TestNewAPIGroup",
		"--input-dirs", "github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/api",
		"--output-base", `C:\go\src\`)
	// TODO wkpo next, regarder
	args.DefaultSourceTree()
	// et setter accordingly (peut etre avec reflect??)

	// TOOD wkpo check qu'on peut compiler?
}

func TestVerboseOutput(t *testing.T) {
	// TODO wkpo
}
