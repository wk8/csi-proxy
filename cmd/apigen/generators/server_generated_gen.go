package generators

// TODO wkpo check all goddamn imports.....
import (
	"fmt"
	"github.com/iancoleman/strcase"
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"strings"
)

// TODO wkpo comment!

// TODO wkpo comment?
type serverGeneratedGenerator struct {
	generator.DefaultGen
	groupDefinition *groupDefinition
	version         *apiVersion
}

// TODO wkpo comment?
type versionedVariableNamer struct {
	version *apiVersion
}

func (n *versionedVariableNamer) Name(t *types.Type) string {
	return varName(t, n.version) + " " + strings.ReplaceAll(t.Name.String(), n.version.Path, n.version.Name)
}

// TODO wkpo comment?
func varName(t *types.Type, version *apiVersion) string {
	varName := shortName(t)
	if isVersionedVariable(t, version) {
		varName = "versioned" + strcase.ToCamel(varName)
	}
	return varName
}

// TODO wkpo comment?
func shortName(t *types.Type) string {
	snake := strcase.ToSnake(t.Name.Name)
	parts := strings.Split(snake, "_")
	return parts[len(parts)-1]
}

// TODO wkpo comment?
func isVersionedVariable(t *types.Type, version *apiVersion) bool {
	return strings.Contains(t.Name.Name, version.Path) ||
		strings.Contains(t.Name.Package, version.Path)
}

func (g *serverGeneratedGenerator) Namers(*generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"versioned": &versionedVariableNamer{
			version: g.version,
		},
	}
}

func (g *serverGeneratedGenerator) Filter(*generator.Context, *types.Type) bool {
	return false
}

func (g *serverGeneratedGenerator) Imports(*generator.Context) []string {
	return []string{
		"context",
		"google.golang.org/grpc",
		"github.com/kubernetes-csi/csi-proxy/client/apiversion",
		g.groupDefinition.internalServerPkg(),
		g.groupDefinition.versionedAPIPkg(g.version.Name),
	}
}

func (g *serverGeneratedGenerator) Init(context *generator.Context, writer io.Writer) error {
	snippetWriter := generator.NewSnippetWriter(writer, context, "$", "$")

	snippetWriter.Do(fmt.Sprintf(`var version = apiversion.NewVersionOrPanic(%q)

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	%s.RegisterDummyServer(grpcServer, s) // TODO wkpo NOPE! pas dummy! group name!
}

`, g.version.Name, g.version.Name), nil)

	for pair := g.version.serverCallbacks.Oldest(); pair != nil; pair = pair.Next() {
		callbackName := pair.Key.(string)
		callback := pair.Value.(*types.Type)

		snippetWriter.Do("func (s *versionedAPI) "+callbackName+"(", nil)

		for _, param := range callback.Signature.Parameters {
			snippetWriter.Do("$.|versioned$, ", param)
		}
	}

	return snippetWriter.Error()
}
