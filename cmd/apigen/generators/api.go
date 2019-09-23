package generators

import (
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
)

// TODO wkpo pkg comment?

// TODO wkpo comment?
type genAPI struct {
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return nil
}

func DefaultNameSystem() string {
	return "wkpo"
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	return nil
}
