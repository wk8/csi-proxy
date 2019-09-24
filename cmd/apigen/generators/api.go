package generators

import (
	"github.com/sirupsen/logrus"
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
	return namer.NameSystems{
		"public": namer.NewPublicNamer(0),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	// TODO wkpo break up in smaller functions?

	for _, input := range context.Inputs {
		logrus.Debugf("Considering input %s", input)

		pkg := context.Universe[input]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}

	}

	// TODO wkpo
	return nil
}
