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
type versionedNamer struct {
	version *apiVersion
}

// TODO wkpo comment?
func (n *versionedNamer) Name(t *types.Type) string {
	return shortenPackagePath(t, n.version.Package)
}

// TODO wkpo comment? move to utils?
func shortenPackagePath(t *types.Type, pkg *types.Package) string {
	return strings.ReplaceAll(t.Name.String(), pkg.Path, pkg.Name)
}

// TODO wkpo comment?
type versionedVariableNamer struct {
	version *apiVersion
}

func (n *versionedVariableNamer) Name(t *types.Type) string {
	return varName(t, n.version) + " " + shortenPackagePath(t, n.version.Package)
}

// TODO wkpo comment? move to utils?
// TODO wkpo rename to versionedName?
func varName(t *types.Type, version *apiVersion) string {
	varName := shortName(t)
	if isVersionedVariable(t, version) {
		varName = "versioned" + strcase.ToCamel(varName)
	}
	return varName
}

// TODO wkpo comment? move to utils?
func shortName(t *types.Type) string {
	snake := strcase.ToSnake(t.Name.Name)
	parts := strings.Split(snake, "_")
	result := parts[len(parts)-1]
	if result == t.Name.Name {
		result = result[:3]
	}
	return result
}

// TODO wkpo comment? move to utils?
func isVersionedVariable(t *types.Type, version *apiVersion) bool {
	return strings.Contains(t.Name.Name, version.Path) ||
		strings.Contains(t.Name.Package, version.Path)
}

// TODO wkpo comment?
type rawNameNamer struct{}

func (n *rawNameNamer) Name(t *types.Type) string {
	parts := strings.Split(t.Name.String(), ".")
	return parts[len(parts)-1]
}

func (g *serverGeneratedGenerator) Namers(*generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"versionedVariable": &versionedVariableNamer{
			version: g.version,
		},
		"versioned": &versionedNamer{
			version: g.version,
		},
		"rawName": &rawNameNamer{},
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

	// write a request handler for each server callback
	for pair := g.version.serverCallbacks.Oldest(); pair != nil; pair = pair.Next() {
		callbackName := pair.Key.(string)
		callback := pair.Value.(*types.Type)

		// write the func signature
		snippetWriter.Do("func (s *versionedAPI) "+callbackName+"(", nil)
		for _, param := range callback.Signature.Parameters {
			snippetWriter.Do("$.|versionedVariable$, ", param)
		}
		snippetWriter.Do(") (", nil)
		for _, result := range callback.Signature.Results {
			snippetWriter.Do("$.|versioned$, ", result)
		}
		snippetWriter.Do(") {\n", nil)

		// then convert all versioned arguments to internal structs
		for _, param := range callback.Signature.Parameters {
			if !isVersionedVariable(param, g.version) {
				continue
			}
			// TODO wkpo check que c un pointer quand on parse initially!!
			// TODO wkpo only if versioned var...
			snippetWriter.Do(fmt.Sprintf("%s := &internal.$.|rawName${}\n", shortName(param)), param)
			snippetWriter.Do(fmt.Sprintf("if err := Convert_%s_$.|rawName$_To_internal_$.|rawName$(%s, %s); err != nil {\n",
				g.version.Name, varName(param, g.version), shortName(param)), param)
			// TODO wkpo check qu'on returne toujours pointer, err?
			// TODO wkpo simplifier le code based on ^ ...?
			snippetWriter.Do("return nil, err\n}\n", nil)
		}
		snippetWriter.Do("\n", nil)

		// call the internal server
		for i, result := range callback.Signature.Results {
			if i != 0 {
				snippetWriter.Do(", ", nil)
			}
			snippetWriter.Do(shortName(result), nil)
		}
		snippetWriter.Do(fmt.Sprintf(" := s.apiGroupServer.%s(", callbackName), nil)
		for _, param := range callback.Signature.Parameters {
			snippetWriter.Do(fmt.Sprintf("%s, ", shortName(param)), nil)
		}
		snippetWriter.Do("version)\nif err != nil {\nreturn nil, err\n}\n\n", nil)

		// convert all internal return values to versioned structs
		for _, result := range callback.Signature.Results {
			if !isVersionedVariable(result, g.version) {
				continue
			}
			snippetWriter.Do(fmt.Sprintf("%s := &%s.$.|rawName${}\n", varName(result, g.version), g.version.Name), result)
			snippetWriter.Do(fmt.Sprintf("if err := Convert_internal_$.|rawName$_To_%s_$.|rawName$(%s, %s); err != nil {\n",
				g.version.Name, shortName(result), varName(result, g.version)), result)
			// TODO wkpo assumptions here...?
			snippetWriter.Do("return nil, err\n}\n", nil)
		}

		// TODO wkpo from here

		// end of the request handler
		snippetWriter.Do("}\n\n", nil)
	}

	return snippetWriter.Error()
}
