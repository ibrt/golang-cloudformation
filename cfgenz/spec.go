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

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

var (
	//go:embed assets/type.tpl
	assetsTypeTPL string
)

// GeneratorSpecOptions describes some options for the generator spec.
type GeneratorSpecOptions struct {
	GoSupportPackage                 string
	GoOutputPackage                  string
	TypeTemplate                     *ttpl.Template
	TypeTemplateGetImplicitGoImports func(t *GeneratorType) []string
}

func (o *GeneratorSpecOptions) getGoSupportBasePackage() string {
	return filepath.Base(o.GoSupportPackage)
}

func (o *GeneratorSpecOptions) getGoSupportType(ic *importsCollector, goType string) string {
	ic.collectImports(o.GoSupportPackage)
	return fmt.Sprintf("%v.%v", o.getGoSupportBasePackage(), goType)
}

// NewDefaultGeneratorSpecOptions initializes a new set of default generator spec options.
func NewDefaultGeneratorSpecOptions() *GeneratorSpecOptions {
	const goBasePackage = "github.com/ibrt/golang-cloudformation"
	goSupportPackage := path.Join(goBasePackage, "cfz")

	return &GeneratorSpecOptions{
		GoSupportPackage: goSupportPackage,
		GoOutputPackage:  path.Join(goBasePackage, "cfz", "gen"),
		TypeTemplate:     ttpl.Must(ttpl.New("").Parse(assetsTypeTPL)),
		TypeTemplateGetImplicitGoImports: func(gt *GeneratorType) []string {
			return newImportsCollector().
				maybeCollectImports(gt.IsTopLevelResourceType(), goSupportPackage).
				maybeCollectImports(gt.IsTopLevelResourceType(), "encoding/json").
				toSlice()
		},
	}
}

// GeneratorSpec is a processed version of the CloudFormation spec ready to be generated.
type GeneratorSpec struct {
	TopLevelResourceTypes map[string]*GeneratorType
	StructuredTypes       map[string]*GeneratorType

	o *GeneratorSpecOptions
	s *cfspecz.Spec
}

// NewGeneratorSpec initializes a new generator spec.
func NewGeneratorSpec(o *GeneratorSpecOptions, s *cfspecz.Spec) *GeneratorSpec {
	gs := &GeneratorSpec{
		o: o,
		s: s,
	}

	gs.TopLevelResourceTypes = memz.TransformMapValues(
		s.ResourceTypes,
		func(_ string, t *cfspecz.Type) *GeneratorType {
			return newGeneratorType(gs, t)
		})

	gs.StructuredTypes = memz.TransformMapValues(
		s.PropertyTypes,
		func(_ string, t *cfspecz.Type) *GeneratorType {
			return newGeneratorType(gs, t)
		})

	return gs
}

// ResourceSpecificationVersion returns the CloudFormation resource specification version for this spec.
func (gs *GeneratorSpec) ResourceSpecificationVersion() string {
	return gs.s.ResourceSpecificationVersion
}

// Generate applies all the templates and writes the results out to disk.
func (gs *GeneratorSpec) Generate(outDirPath string) error {
	if err := os.RemoveAll(outDirPath); err != nil {
		return errorz.Wrap(err)
	}

	for _, gt := range gs.StructuredTypes {
		if err := gt.generate(outDirPath); err != nil {
			return errorz.Wrap(err)
		}
	}

	for _, gt := range gs.TopLevelResourceTypes {
		if err := gt.generate(outDirPath); err != nil {
			return errorz.Wrap(err)
		}
	}

	return nil
}

func (gs *GeneratorSpec) makeMustLookupType(gt *GeneratorType, sc cfspecz.SpecContext) func(string) *GeneratorType {
	return func(unqualifiedStructuredTypeName string) *GeneratorType {
		t := gs.StructuredTypes[gt.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]

		errorz.Assertf(t != nil,
			"type lookup failed: from '%v' to '%v'",
			sc.GetDisplayPath(), unqualifiedStructuredTypeName)

		return t
	}
}

func (gs *GeneratorSpec) getPrimitiveGoType(ic *importsCollector, primitiveType string) string {
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
