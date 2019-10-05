package generators

import (
	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/internal"
	"io"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
)

// TODO wkpo comment!

// TODO wkpo comment?
type typesGeneratedGenerator struct {
	generator.DefaultGen
	groupDefinition *groupDefinition
}

func (g *typesGeneratedGenerator) Filter(*generator.Context, *types.Type) bool {
	return false
}

func (g *typesGeneratedGenerator) Imports(*generator.Context) []string {
	return []string{
		"context",
		"google.golang.org/grpc",
		"github.com/kubernetes-csi/csi-proxy/client/apiversion",
	}
}

func (g *typesGeneratedGenerator) Init(context *generator.Context, writer io.Writer) error {
	snippetWriter := generator.NewSnippetWriter(writer, context, "$", "$")

	snippetWriter.Do(`type VersionedAPI interface {
Register(grpcServer *grpc.Server)
}

// All the functions this group's server needs to define.
type ServerInterface interface {
`, nil)

	for _, namedCallback := range g.groupDefinition.serverCallbacks {
		callback := internal.ReplaceTypesPackage(namedCallback.callback, internal.PkgPlaceholder, "")

		snippetWriter.Do(namedCallback.name+"(", nil)
		for _, param := range callback.Signature.Parameters {
			snippetWriter.Do("$.Name.String$, ", param)
		}
		// add the version parameter
		snippetWriter.Do("apiversion.Version) (", nil)
		for _, returnValue := range callback.Signature.Results {
			snippetWriter.Do("$.Name.String$, ", returnValue)
		}
		snippetWriter.Do(")\n", nil)
	}
	snippetWriter.Do("}\n", nil)

	return snippetWriter.Error()
}
