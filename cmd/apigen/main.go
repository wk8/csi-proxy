package main

import (
	goflag "flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"k8s.io/gengo/args"

	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/generators"
)

// TODO wkpo comment?
const csiProxyAPIPackage = "github.com/kubernetes-csi/csi-proxy/client/api/..."

func main() {
	// TODO wkpo mouaif
	logrus.SetLevel(logrus.DebugLevel)

	// TODO wkpo different function!
	genericArgs := args.Default().WithoutDefaultFlagParsing()
	genericArgs.AddFlags(pflag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

	if len(genericArgs.InputDirs) == 0 {
		genericArgs.InputDirs = append(genericArgs.InputDirs, csiProxyAPIPackage)
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
