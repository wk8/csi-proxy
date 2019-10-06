package execute

import (
	"strings"

	goflag "flag"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"
	"k8s.io/klog"

	"github.com/kubernetes-csi/csi-proxy/cmd/csi-api-gen/generators"
	"github.com/kubernetes-csi/csi-proxy/cmd/csi-api-gen/internal"
)

// Execute runs csi-api-gen. It's exposed as a public function in a separate package
// to be able to run it from integration tests.
func Execute(cliArgs []string) {
	klog.InitFlags(nil)

	if err := buildArgs(cliArgs).Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	klog.Infof("Generation successful!")
}

func buildArgs(cliArgs []string) *args.GeneratorArgs {
	genericArgs := args.Default().WithoutDefaultFlagParsing()

	genericArgs.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.CommandLine.Parse(cliArgs)

	// if no package argument, default to processing canonical API groups, under CSIProxyAPIPath
	if len(genericArgs.InputDirs) == 0 {
		genericArgs.InputDirs = append(genericArgs.InputDirs, internal.CSIProxyAPIPath)
	}

	// it doesn't really make sense to consider a package in isolation, since an API group is
	// always a collection of subpackages (its versions)
	// so we consider all inputs recursively
	for i, inputDir := range genericArgs.InputDirs {
		if !strings.HasSuffix(inputDir, "...") {
			genericArgs.InputDirs[i] = internal.CanonicalizePkgPath(inputDir) + "/..."
		}
	}

	return genericArgs
}
