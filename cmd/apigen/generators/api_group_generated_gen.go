package generators

// TODO wkpo check all goddamn imports.....
import (
	"fmt"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
	"sort"
)

// TODO wkpo comment!

// TODO wkpo comment?
type apiGroupsGeneratedGenerator struct {
	generator.DefaultGen
	groupDefinition *groupDefinition
}

func (g *apiGroupsGeneratedGenerator) Filter(*generator.Context, *types.Type) bool {
	return false
}

func (g *apiGroupsGeneratedGenerator) Imports(*generator.Context) []string {
	imports := []string{
		"github.com/kubernetes-csi/csi-proxy/client/apiversion",
		"github.com/kubernetes-csi/csi-proxy/internal/server",
		g.groupDefinition.internalServerPkg(),
	}

	for _, version := range g.groupDefinition.versions {
		imports = append(imports, g.groupDefinition.versionedServerPkg(version.Name))
	}

	return imports
}

func (g *apiGroupsGeneratedGenerator) Init(context *generator.Context, writer io.Writer) error {
	snippetWriter := generator.NewSnippetWriter(writer, context, "$", "$")

	snippetWriter.Do(fmt.Sprintf("const name = %q", g.groupDefinition.name), nil)

	snippetWriter.Do(`

// ensure the server defines all the required methods
var _ internal.ServerInterface = &Server{}

func (s *Server) VersionedAPIs() []*server.VersionedAPI {
`, nil)

	versions := make([]string, len(g.groupDefinition.versions))
	for i, vsn := range g.groupDefinition.versions {
		versions[i] = vsn.Name
	}
	sort.Slice(versions, func(i, j int) bool {
		// TODO wkpo ce serait mieux de faire versions une []apiversion.Version non?
		return apiversion.NewVersionOrPanic(versions[i]).Compare(apiversion.NewVersionOrPanic(versions[j])) == apiversion.Lesser
	})

	for _, version := range versions {
		snippetWriter.Do(fmt.Sprintf("%sServer := %s.NewVersionedServer(s)\n", version, version), nil)
	}

	snippetWriter.Do("\n\nreturn []*server.VersionedAPI{\n", nil)
	for _, version := range versions {
		snippetWriter.Do(fmt.Sprintf(`{
				Group:      name,
				Version:    apiversion.NewVersionOrPanic(%q),
				Registrant: %sServer.Register,
			},
			`, version, version), nil)
	}
	snippetWriter.Do("\n}\n}\n", nil)

	return snippetWriter.Error()
}
