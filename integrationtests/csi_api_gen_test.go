package integrationtests

// TODO wkpo imports?
import (
	"github.com/kubernetes-csi/csi-proxy/cmd/csi-api-gen/generators"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

// TODO wkpo capture output for all of em!

// This tests API generator; more specifically, its main goal is to ensure
// that the API generator works as expected when creating a new API group.
// On top of this, the regular build checks that all checked-in generated files
// are up-to-date (i.e. consistent with the current generator).

func TestNewAPIGroup(t *testing.T) {
	// TODO wkpo DRY that up? and test it works?
	func() {
		defer func() {
			require.Nil(t, recover(), "panic when generating code")
		}()
		generators.Execute("TestNewAPIGroup",
			"--input-dirs", "github.com/kubernetes-csi/csi-proxy/integrationtests/csiapigen/new_group/api")
	}()

	// now check the output is exactly what we expect
	recursiveDiff(t, "csiapigen/new_group/expected_output", "csiapigen/new_group/actual_output")
}

func TestVerboseOutput(t *testing.T) {
	// TODO wkpo
}

// runGenerator runs csi-api-gen with the given CLI args, and returns stdout
func runGenerator(t *testing.T, testName string, cliArgs ...string) string {
	// TODO wkpo next from here....
	AppFs = afero.NewMemMapFs()
	logFile, err := ioutil.TempFile("", "test-csi-api-gen-"+testName)
	require.Nil(t, err)
	require.Nil(t, logFile.Close())
	defer func() {
		require.Nil(t, recover(), "panic when generating code, output:\n%s", readFile(t, logFile.Name()))
		require.Nil(t, os.Remove(logFile.Name()))
	}()

	generators.Execute(testName, append(cliArgs)...)
}
