package generators

// TODO wkpo check all goddamn imports.....
import (
	"github.com/kubernetes-csi/csi-proxy/cmd/apigen/internal"
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
	"strings"
)

// TODO wkpo comment!

// TODO wkpo comment?
type serverGenerator struct {
	generator.DefaultGen
	groupDefinition *groupDefinition
}

func (g *serverGenerator) Filter(*generator.Context, *types.Type) bool {
	return false
}

func (g *serverGenerator) Imports(*generator.Context) []string {
	return []string{
		"context",
		"fmt",
		"github.com/kubernetes-csi/csi-proxy/client/apiversion",
		g.groupDefinition.internalServerPkg(),
	}
}

func (g *serverGenerator) Init(context *generator.Context, writer io.Writer) error {
	snippetWriter := generator.NewSnippetWriter(writer, context, "$", "$")

	snippetWriter.Do("type Server struct{}\n\n", nil)

	for _, namedCallback := range g.groupDefinition.serverCallbacks {
		callback := internal.ReplaceTypesPackage(namedCallback.callback, internal.PkgPlaceholder, "internal")

		snippetWriter.Do("func (s *Server) "+namedCallback.name+"(", nil)
		for _, param := range callback.Signature.Parameters {
			snippetWriter.Do("$.|short$ $.Name.String$, ", param)
		}
		// add the version parameter
		snippetWriter.Do("version apiversion.Version) (", nil)
		for _, returnValue := range callback.Signature.Results {
			snippetWriter.Do("$.Name.String$, ", returnValue)
		}

		snippetWriter.Do(") {\n// TODO: auto-generated stub\n", nil)
		snippetWriter.Do("return nil"+strings.Repeat(", nil", len(callback.Signature.Results)-1)+"}\n\n", nil)
	}

	return snippetWriter.Error()
}
