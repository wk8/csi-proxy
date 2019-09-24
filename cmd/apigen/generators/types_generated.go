package generators

import (
	"io"

	gengogenerator "k8s.io/gengo/generator"
	gengotypes "k8s.io/gengo/types"
)

// TODO wkpo comment!

// TODO wkpo comment?
type typesGeneratedGenerator struct {
	gengogenerator.DefaultGen
	serverCallbacks []*namedCallback
}

func (g *typesGeneratedGenerator) Filter(_ *gengogenerator.Context, t *gengotypes.Type) bool {
	return false
}

func (g *typesGeneratedGenerator) Imports(_ *gengogenerator.Context) []string {
	return []string{
		"context",
		"google.golang.org/grpc",
		"github.com/kubernetes-csi/csi-proxy/client/apiversion",
	}
}

func (g *typesGeneratedGenerator) Init(context *gengogenerator.Context, writer io.Writer) error {
	snippetWriter := gengogenerator.NewSnippetWriter(writer, context, "$", "$")

	snippetWriter.Do(`type VersionedAPI interface {
	Register(grpcServer *grpc.Server)
}

// All the functions this group's server needs to define.
type ServerInterface interface {
`, nil)

	for _, callback := range g.serverCallbacks {
		snippetWriter.Do("	$.Name$", callback)
		writeTypesList(snippetWriter, callback.Signature.Parameters...)
		snippetWriter.Do(" ", nil)
		writeTypesList(snippetWriter, callback.Signature.Results...)
		snippetWriter.Do("\n", nil)
	}

	return snippetWriter.Error()
}

func writeTypesList(snippetWriter *gengogenerator.SnippetWriter, types ...*gengotypes.Type) {
	snippetWriter.Do("(", nil)
	for _, t := range types {
		snippetWriter.Do("$.Name.Name$, ", t)
		// TODO wkpo next from here, gonna need a custom namer...
		// TODO wkpo next look at the one from conversion-namer, using NameStrategy, might be decent?
	}
	snippetWriter.Do(")", nil)
}
