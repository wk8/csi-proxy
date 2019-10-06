package integrationtests

import (
	"testing"

	"github.com/kubernetes-csi/csi-proxy/cmd/csi-api-gen/generators"
)

// TODO wkpo capture output for all of em!
// TODO wkpo capture panics, too...!

// This tests API generator; more specifically, its main goal is to ensure
// that the API generator works as expected when creating a new API group,
// or a new API version in an existing group.
// On top of this, the regular build checks that all checked-in generated files
// are up-to-date (i.e. consistent with the current generator).

func TestNewAPIGroup(t *testing.T) {
	// TODO wkpo!
	generators.Execute("TestNewAPIGroup",
		"--input-dirs", "github.com/kubernetes-csi/csi-proxy/integrationtests/csi-api-gen/new_group")
}

func TestNewAPIVersion(t *testing.T) {
	// TODO wkpo
}

func TestVerboseOutput(t *testing.T) {
	// TODO wkpo
}
