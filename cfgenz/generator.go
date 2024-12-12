package cfgenz

import (
	_ "embed" // embed
	"fmt"
	"os"
	"path"
	"path/filepath"
	ttpl "text/template"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"

	"github.com/ibrt/golang-cloudformation/cfschemaz"
	"github.com/ibrt/golang-cloudformation/cfspecz"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	//go:embed assets/type.tpl
	assetsTypeTPL string
)

// GeneratorOptions describes some options for the generator.
type GeneratorOptions struct {
	GoSupportPackage                 string
	GoOutputPackage                  string
	TypeTemplate                     *ttpl.Template
	TypeTemplateGetImplicitGoImports func(t *Type) []string
}

func (o *GeneratorOptions) getGoSupportBasePackage() string {
	return filepath.Base(o.GoSupportPackage)
}

func (o *GeneratorOptions) getGoSupportType(ic *importsCollector, goType string) string {
	ic.collectImports(o.GoSupportPackage)
	return fmt.Sprintf("%v.%v", o.getGoSupportBasePackage(), goType)
}

// NewDefaultGeneratorOptions initializes a new set of default generator options.
func NewDefaultGeneratorOptions() *GeneratorOptions {
	const goBasePackage = "github.com/ibrt/golang-cloudformation"
	goSupportPackage := path.Join(goBasePackage, "cfz")

	return &GeneratorOptions{
		GoSupportPackage: goSupportPackage,
		GoOutputPackage:  path.Join(goBasePackage, "cfz", "gen"),
		TypeTemplate:     ttpl.Must(ttpl.New("").Parse(assetsTypeTPL)),
		TypeTemplateGetImplicitGoImports: func(gt *Type) []string {
			return newImportsCollector().
				maybeCollectImports(gt.IsTopLevelResourceType(), goSupportPackage).
				maybeCollectImports(gt.IsTopLevelResourceType(), "encoding/json").
				toSlice()
		},
	}
}

// Generator combines g and JSON schema to generate a CloudFormation library.
type Generator struct {
	TopLevelResourceTypes map[string]*Type
	StructuredTypes       map[string]*Type

	o      *GeneratorOptions
	spec   *cfspecz.Spec
	schema *cfschemaz.Schema
}

// NewGenerator initializes a new generator.
func NewGenerator(o *GeneratorOptions, spec *cfspecz.Spec, schema *cfschemaz.Schema) *Generator {
	g := &Generator{
		o:      o,
		spec:   spec,
		schema: schema,
	}

	g.TopLevelResourceTypes = memz.TransformMapValues(
		spec.TopLevelResourceTypes,
		func(_ string, specT *cfspecz.Type) *Type {
			return newType(g, specT, schema.Resources[specT.Name])
		})

	g.StructuredTypes = memz.TransformMapValues(
		spec.StructuredTypes,
		func(_ string, specT *cfspecz.Type) *Type {
			return newType(g, specT, schema.Resources[specT.Name])
		})

	return g
}

// ResourceSpecificationVersion returns the spec version.
func (g *Generator) ResourceSpecificationVersion() string {
	return g.spec.ResourceSpecificationVersion
}

// Generate applies all the templates and writes the results out to disk.
func (g *Generator) Generate(outDirPath string) error {
	if err := os.RemoveAll(outDirPath); err != nil {
		return errorz.Wrap(err)
	}

	for _, gt := range g.StructuredTypes {
		if err := gt.generate(outDirPath); err != nil {
			return errorz.Wrap(err)
		}
	}

	for _, gt := range g.TopLevelResourceTypes {
		if err := gt.generate(outDirPath); err != nil {
			return errorz.Wrap(err)
		}
	}

	return nil
}

func (g *Generator) makeMustLookupType(t *Type, pl cfz.ProblemLocation) func(string) *Type {
	return func(unqualifiedStructuredTypeName string) *Type {
		t := g.StructuredTypes[t.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]

		errorz.Assertf(t != nil,
			"type lookup failed: from '%v' to '%v'",
			pl.GetDisplayPath(), unqualifiedStructuredTypeName)

		return t
	}
}

func (g *Generator) getPrimitiveGoType(ic *importsCollector, primitiveType string) string {
	switch primitiveType {
	case "Boolean":
		return "bool"
	case "Double":
		return "float64"
	case "Integer":
		// TODO(ibrt): Switch to int?
		return "int32"
	case "Json":
		// TODO(ibrt): Dedicated type?
		ic.collectImports("encoding/json")
		return "json.RawMessage"
	case "Long":
		return "int64"
	case "String":
		return "string"
	case "Timestamp":
		// TODO(ibrt): Dedicated type?
		return "string"
	default:
		panic(errorz.Errorf("unknown primitive type: '%v'", primitiveType))
	}
}
