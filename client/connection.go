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
	apiVersion apiversions.Version
	handler    *apiversions.Handler
}

// NewConnection creates a new client.
// If the API version is not specified, will automatically select the highest common API version supported
// by both the server and us.
func NewConnection(apiVersion *apiversions.Version, definitions ...apiversions.Definition) (*Connection, error) {
	handler := &apiversions.Handler{}
	if err := handler.Register(definitions...); err != nil {
		return nil, err
	}

	availableVersions, err := listCsiProxyVersionsFromPipes()
	if err != nil {
		return nil, err
	}

	if apiVersion == nil {
		apiVersion, err = resolveHighestCommonVersion(handler.Definitions(), availableVersions)
		if err != nil {
			return nil, err
		}
	} else if _, present := availableVersions[apiVersion.String()]; !present {
		return nil, fmt.Errorf("version %v not available; available versions: %v", apiVersion, mapKeysToSlice(availableVersions))
	}

	return &Connection{
		handler: handler,
		apiVersion: *apiVersion,
	}, nil
}

// resolveHighestCommonVersion compares available server API versions to our own, and selects the highest common one.
func resolveHighestCommonVersion(definitions []apiversions.Definition, availableVersions map[string]bool) (apiVersion *apiversions.Version, err error) {
	for i := len(definitions) - 1; i > 0; i-- {
		if _, present := availableVersions[definitions[i].Version.String()]; present {
			apiVersion = &definitions[i].Version
			break
		}
	}

	if apiVersion == nil {
		ourVersions := make([]string, len(definitions)) {
			for i, definition := range definitions {
				ourVersions[i] = definition.Version.String()
			}
		}
		err = fmt.Errorf("no matching versions: we support versions %v, while the server offers %v", mapKeysToSlice(availableVersions))
	}

	return
}

func mapKeysToSlice(m map[string]bool) []string {
	keys := make([]string, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func (c *Connection) Version() apiversions.Version {
	return c.Version()
}

// listCsiProxyVersionsFromPipes lists the raw string versions for the CSI proxy pipes that are
// listening on their named pipes.
// Returns a map for easy look ups.
func listCsiProxyVersionsFromPipes() (map[string]bool, error) {
	allNamedPipes, err := listNamedPipes()
	if err != nil {
		return nil, err
	}

	csiProxyPipesPrefix := internal.PipePrefix + internal.CsiProxyNamedPipePrefix
	versions := make(map[string]bool)

	for _, namedPipe := range allNamedPipes {
		withoutPrefix := strings.TrimPrefix(namedPipe, csiProxyPipesPrefix)
		if withoutPrefix != namedPipe && apiversions.IsValidVersionName(withoutPrefix) {
			versions[withoutPrefix] = true
		}
	}

	return versions, nil
}

// listNamedPipes list all currently listening Windows named pipes.
func listNamedPipes() ([]string, error) {
	powershellCmd := fmt.Sprintf("[System.IO.Directory]::GetFiles('%s')", internal.PipePrefix)
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
