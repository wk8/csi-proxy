package client

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

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
		if !strings.HasPrefix(line, pipePrefix) {
			return nil, fmt.Errorf("unexpected output from %q, line %v %q doesn't start with %q - full output: %q", powershellCmd, i+1, line, outStr)
		}
	}

	return namedPipes, nil
}
