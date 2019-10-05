package main

import (
	"strings"

	goflag "flag"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"
	"k8s.io/klog"

	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/generators"
	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/internal"
)

func main() {
	klog.InitFlags(nil)

	if err := buildArgs().Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		klog.Fatalf("Error: %v", err)
	}

	klog.Infof("Generation successful!")
}

func buildArgs() *args.GeneratorArgs {
	genericArgs := args.Default().WithoutDefaultFlagParsing()

	genericArgs.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

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
