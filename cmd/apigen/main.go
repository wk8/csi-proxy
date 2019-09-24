package main

import (
	"strings"

	goflag "flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"

	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/generators"
	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/internal"
)

func main() {
	// TODO wkpo mouaif
	logrus.SetLevel(logrus.DebugLevel)

	// TODO wkpo separate function!
	genericArgs := args.Default().WithoutDefaultFlagParsing()
	genericArgs.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

	if len(genericArgs.InputDirs) == 0 {
		genericArgs.InputDirs = append(genericArgs.InputDirs, internal.CSIProxyAPIPath)
	}
	// it doesn't really make sense to consider a package in isolation, since an API group is
	// always a collection of subpackages (its versions)
	// so we consider all inputs recursively
	for i, inputDir := range genericArgs.InputDirs {
		suffix := "..."
		if !strings.HasSuffix(inputDir, suffix) {
			if !strings.HasSuffix(inputDir, "/") {
				suffix = "/" + suffix
			}
			genericArgs.InputDirs[i] = inputDir + suffix
		}
	}

	if err := genericArgs.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		logrus.Fatalf("Error: %v", err)
	}

	logrus.Infof("wkpo bordel")
}
