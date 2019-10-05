package generators

// TODO wkpo check all goddamn imports.....
import (
	"io"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog"
	"strings"
)

// TODO wkpo comment!

// TODO wkpo have a complicated thing in the test, with structs referencing one another

// TODO wkpo comment?
type typesGenerator struct {
	generator.DefaultGen

	groupDefinition *groupDefinition
	version         *apiVersion

	importTracker namer.ImportTracker
}

func (g *typesGenerator) Namers(*generator.Context) namer.NameSystems {
	g.importTracker = generator.NewImportTracker()

	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.groupDefinition.internalServerPkg(), g.importTracker),
	}
}

func (g *typesGenerator) Filter(*generator.Context, *types.Type) bool {
	return false
}

func (g *typesGenerator) Imports(context *generator.Context) (imports []string) {
	return g.importTracker.ImportLines()
}

func (g *typesGenerator) Init(context *generator.Context, writer io.Writer) error {
	snippetWriter := generator.NewSnippetWriter(writer, context, "$", "$")

	protoPkgPath := g.groupDefinition.versionedAPIPkg(g.version.Name)
	protoPkg := context.Universe[protoPkgPath]
	if protoPkg == nil {
		// shouldn't happen
		klog.Fatalf("proto package %s for API group %s version %s not loaded",
			protoPkgPath, g.groupDefinition.name, g.version.Name)
	}

	// TODO wkpo make that a list, re-order first, alphabetically
	for typeName, t := range protoPkg.Types {
		// look at types that define a ProtoMessage method, these are the messages we need
		// to have internal types corresponding to
		// TODO wkpo make that a func, and check more (ie no return, no params, etc...)
		if _, isMsg := t.Methods["ProtoMessage"]; isMsg {
			g.generateStruct(typeName, t, snippetWriter)
		}
	}

	return snippetWriter.Error()
}

func (g *typesGenerator) generateStruct(typeName string, t *types.Type, snippetWriter *generator.SnippetWriter) {
	snippetWriter.Do("type "+typeName+" struct {\n", nil)

	for _, member := range t.Members {
		if strings.HasPrefix(member.Name, "XXX_") {
			// internal protobuf field
			continue
		}

		for _, commentLine := range member.CommentLines {
			commentLine = strings.TrimSpace(commentLine)
			if commentLine == "" {
				continue
			}
			snippetWriter.Do("// $.$\n", commentLine)
		}
		snippetWriter.Do(member.Name+" $.|raw$\n", member.Type)
	}

	snippetWriter.Do("}\n\n", nil)
}
