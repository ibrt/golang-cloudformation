package cfgenz

import (
	_ "embed" // embed
	"path"
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
	GoSupportPackage      string
	GoOutputPackage       string
	TypeTemplate          *ttpl.Template
	TypeTemplateGoImports []string
}

// NewDefaultGeneratorSpecOptions initializes a new set of default generator spec options.
func NewDefaultGeneratorSpecOptions() *GeneratorSpecOptions {
	const basePackage = "github.com/ibrt/golang-cloudformation"

	return &GeneratorSpecOptions{
		GoSupportPackage: path.Join(basePackage, "cfz"),
		GoOutputPackage:  path.Join(basePackage, "cfz", "gen"),
		TypeTemplate:     ttpl.Must(ttpl.New("").Parse(assetsTypeTPL)),
		TypeTemplateGoImports: []string{
			path.Join(basePackage, "cfz"),
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

func (gs *GeneratorSpec) makeMustLookupType(gt *GeneratorType, sc cfspecz.SpecContext) func(string) *GeneratorType {
	return func(unqualifiedStructuredTypeName string) *GeneratorType {
		t := gs.StructuredTypes[gt.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]

		errorz.Assertf(t != nil,
			"type lookup failed: from '%v' to '%v'",
			sc.GetDisplayPath(), unqualifiedStructuredTypeName)

		return t
	}
}
