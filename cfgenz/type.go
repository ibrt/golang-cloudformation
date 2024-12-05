package cfgenz

import (
	"strings"

	"github.com/ibrt/golang-utils/memz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// GeneratorType describes either a top-level resource type or a structured type in the generator spec.
type GeneratorType struct {
	Attributes map[string]*GeneratorAttribute
	Properties map[string]*GeneratorProperty

	spec *GeneratorSpec
	t    *cfspecz.Type
}

func newGeneratorType(gs *GeneratorSpec, t *cfspecz.Type) *GeneratorType {
	gt := &GeneratorType{
		spec: gs,
		t:    t,
	}

	gt.Attributes = memz.TransformMapValues(t.Attributes, func(_ string, a *cfspecz.Attribute) *GeneratorAttribute {
		return newGeneratorAttribute(gt, a)
	})

	gt.Properties = memz.TransformMapValues(t.Properties, func(_ string, p *cfspecz.Property) *GeneratorProperty {
		return newGeneratorProperty(gt, p)
	})

	return gt
}

// GetRelatedTopLevelResourceTypeName returns the top-level resource type name related to this type, that is, Name if
// IsTopLevelResourceType = true, or the portion of Name before the first "." otherwise.
func (t *GeneratorType) GetRelatedTopLevelResourceTypeName() string {
	return t.t.GetRelatedTopLevelResourceTypeName()
}

// GetRelatedStructuredTypeName converts the value of a Type or ItemType field of a property or attribute in this type
// to its fully qualified structured type name.
func (t *GeneratorType) GetRelatedStructuredTypeName(unqualifiedStructuredTypeName string) string {
	return t.t.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)
}

// GoName returns the Go name for this type.
func (t *GeneratorType) GoName() string {
	structName := strings.ReplaceAll(t.t.Name, "::", "_")
	return strings.ReplaceAll(structName, ".", "_")
}

// GoPackageName returns the Go package name for this type.
func (t *GeneratorType) GoPackageName() string {
	return strings.ToLower(strings.Join(
		strings.Split(t.GetRelatedTopLevelResourceTypeName(), "::")[0:2],
		"_"))
}

// GoFileName returns the Go file name for this type.
func (t *GeneratorType) GoFileName() string {
	return strings.ToLower(t.GoName()) + ".go"
}

// GoImports returns the Go imports for this type.
func (t *GeneratorType) GoImports() []string {
	ic := newImportsCollector()
	ic.collectImports(t.spec.o.TypeTemplateRequiredGoImports...)

	for _, ga := range t.Attributes {
		ga.collectImports(ic)
	}

	for _, gp := range t.Properties {
		gp.collectImports(ic)
	}

	return ic.toSlice()
}

// GoSupportBasePackage returns the base package for the support library.
func (t *GeneratorType) GoSupportBasePackage() string {
	return t.spec.o.getGoSupportBasePackage()
}

// DocumentationURL returns the documentation URL for this type.
func (t *GeneratorType) DocumentationURL() string {
	return t.t.Documentation
}

// IsTopLevelResourceType returns true if this is a top-level resource type, false if it is a structured type.
func (t *GeneratorType) IsTopLevelResourceType() bool {
	return t.t.IsTopLevelResourceType
}

// Name returns the CloudFormation name for this type.
func (t *GeneratorType) Name() string {
	return t.t.Name
}
