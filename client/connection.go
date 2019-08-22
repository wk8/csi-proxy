package client

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/kubernetes-csi/csi-proxy/internal"
	"github.com/kubernetes-csi/csi-proxy/internal/apiversions"
)

// TODO wkpo should be in internal...
// TODO wkpo reevaluate what is in internal and what's outside? only the client funcs should be outside, really.

type Connection struct {
	handler *apiversions.Handler
}

// NewConnection creates a new client.
// If the API version is not specified, will automatically select
// TODO wkpo comment?
func NewConnection(apiVersion *apiversions.Version, definitions ...apiversions.Definition) (*Client, error) {
	// TODO wkpo next from here
}

// listNamedPipes list all Windows named pipes.
func listNamedPipes() ([]string, error) {
	powershellCmd := fmt.Sprintf("[System.IO.Directory]::GetFiles('%s')", pipePrefix)
	out, err := exec.Command("powershell", powershellCmd).CombinedOutput()
	outStr := string(out)

	if err != nil {
		return nil, errors.Wrapf(err, "failed listing named pipes from %q, output: %q", powershellCmd, outStr)
	}

	namedPipes := strings.Split(strings.TrimSpace(outStr), "\r\n")
	for i, line := range namedPipes {
		if !strings.HasPrefix(line, internal.PipePrefix) {
			return nil, fmt.Errorf("unexpected output from %q, line %v %q doesn't start with %q - full output: %q", powershellCmd, i+1, line, outStr)
		}
	}

	return namedPipes, nil
}
