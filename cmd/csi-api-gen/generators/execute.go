package generators

import (
	"strings"

	goflag "flag"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"
	"k8s.io/klog"
)

// Execute runs csi-api-gen. It's exposed as a public function
// to be able to easily run it from integration tests.
func Execute(flagSetName string, cliArgs ...string) {
	if err := buildArgs(flagSetName, cliArgs).Execute(
		nameSystems(),
		defaultNameSystem(),
		packages,
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	klog.Infof("Generation successful!")
}

func buildArgs(flagSetName string, cliArgs []string) *args.GeneratorArgs {
	goFlagSet := goflag.NewFlagSet(flagSetName, goflag.ExitOnError)
	klog.InitFlags(goFlagSet)

	genericArgs := args.Default().WithoutDefaultFlagParsing()

	pflagFlagSet := pflag.NewFlagSet(flagSetName, pflag.ExitOnError)
	genericArgs.AddFlags(pflagFlagSet)
	pflagFlagSet.AddGoFlagSet(goFlagSet)
	pflagFlagSet.Parse(cliArgs)

	// if no package argument, default to processing canonical API groups, under csiProxyAPIPath
	if len(genericArgs.InputDirs) == 0 {
		genericArgs.InputDirs = append(genericArgs.InputDirs, csiProxyAPIPath)
	}

	// it doesn't really make sense to consider a package in isolation, since an API group is
	// always a collection of subpackages (its versions)
	// so we consider all inputs recursively
	for i, inputDir := range genericArgs.InputDirs {
		if !strings.HasSuffix(inputDir, "...") {
			genericArgs.InputDirs[i] = canonicalizePkgPath(inputDir) + "/..."
		}
	}

	return genericArgs
}
