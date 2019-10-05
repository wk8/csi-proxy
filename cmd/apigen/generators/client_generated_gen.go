package generators

// TODO wkpo check all goddamn imports.....
import (
	"github.com/iancoleman/strcase"
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

// TODO wkpo comment!

// TODO wkpo comment?
type clientGeneratedGenerator struct {
	generator.DefaultGen
	groupDefinition *groupDefinition
	version         *apiVersion
}

func (g *clientGeneratedGenerator) Namers(*generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"shortenVersionPackage": &shortenVersionPackageNamer{
			version: g.version,
		},
	}
}

func (g *clientGeneratedGenerator) Filter(*generator.Context, *types.Type) bool {
	return false
}

func (g *clientGeneratedGenerator) Imports(*generator.Context) []string {
	return []string{
		"context",
		"net",
		"github.com/Microsoft/go-winio",
		"google.golang.org/grpc",
		"github.com/kubernetes-csi/csi-proxy/client",
		"github.com/kubernetes-csi/csi-proxy/client/apiversion",
		g.groupDefinition.versionedAPIPkg(g.version.Name),
	}
}

func (g *clientGeneratedGenerator) Init(context *generator.Context, writer io.Writer) error {
	snippetWriter := generator.NewSnippetWriter(writer, context, "$", "$")

	snippetWriter.Do(`
const groupName = "$.groupName$"

var version = apiversion.NewVersionOrPanic("$.version$")

type wrapper struct {
	client     $.version$.$.camelGroupName$Client
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the $.groupName$ API group version $.version$.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*wrapper, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := $.version$.New$.camelGroupName$Client(connection)
	return &wrapper{
		client:     client,
		connection: connection,
	}, nil
}

// ensures we implement all the required methods
var _ $.version$.$.camelGroupName$Client = &wrapper{}

	`, map[string]string{
		"camelGroupName": strcase.ToCamel(g.groupDefinition.name),
		"groupName":      g.groupDefinition.name,
		"version":        g.version.Name,
	})

	for _, namedCallback := range g.version.serverCallbacks {
		g.writeWrapperFunction(namedCallback.name, namedCallback.callback, snippetWriter)
	}

	return snippetWriter.Error()
}

func (g *clientGeneratedGenerator) writeWrapperFunction(callbackName string, callback *types.Type, snippetWriter *generator.SnippetWriter) {
	snippetWriter.Do("func (w *wrapper) $.$(", callbackName)

	for _, param := range callback.Signature.Parameters {
		snippetWriter.Do("$.|short$ $.|shortenVersionPackage$, ", param)
	}
	snippetWriter.Do("opts ...grpc.CallOption) (", nil)
	for _, returnValue := range callback.Signature.Results {
		snippetWriter.Do("$.|shortenVersionPackage$, ", returnValue)
	}
	snippetWriter.Do(") {\n", nil)

	snippetWriter.Do("return w.client.$.$(", callbackName)
	for _, param := range callback.Signature.Parameters {
		snippetWriter.Do("$.|short$, ", param)
	}
	snippetWriter.Do("opts ...)\n", nil)

	snippetWriter.Do("}\n\n", nil)
}
