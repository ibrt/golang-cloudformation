package cfgenz

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
	"github.com/ibrt/golang-utils/tplz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// GeneratorType describes either a top-level resource type or a structured type in the generator g.
type GeneratorType struct {
	Attributes map[string]*Attribute
	Properties map[string]*Property

	g     *Generator
	specT *cfspecz.Type
}

func newGeneratorType(gs *Generator, specT *cfspecz.Type) *GeneratorType {
	gt := &GeneratorType{
		g:     gs,
		specT: specT,
	}

	gt.Attributes = memz.TransformMapValues(specT.Attributes, func(_ string, a *cfspecz.Attribute) *Attribute {
		return newAttribute(gt, a)
	})

	gt.Properties = memz.TransformMapValues(specT.Properties, func(_ string, p *cfspecz.Property) *Property {
		return newProperty(gt, p)
	})

	return gt
}

// GetRelatedTopLevelResourceTypeName returns the top-level resource type name related to this type, that is, Name if
// IsTopLevelResourceType = true, or the portion of Name before the first "." otherwise.
func (gt *GeneratorType) GetRelatedTopLevelResourceTypeName() string {
	return gt.specT.GetRelatedTopLevelResourceTypeName()
}

// GetRelatedStructuredTypeName converts the value of a Type or ItemType field of a property or attribute in this type
// to its fully qualified structured type name.
func (gt *GeneratorType) GetRelatedStructuredTypeName(unqualifiedStructuredTypeName string) string {
	return gt.specT.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)
}

// GoName returns the Go name for this type.
func (gt *GeneratorType) GoName() string {
	structName := strings.ReplaceAll(gt.specT.Name, "::", "_")
	return strings.ReplaceAll(structName, ".", "_")
}

// GoPackageName returns the Go package name for this type.
func (gt *GeneratorType) GoPackageName() string {
	return strings.ToLower(strings.Join(
		strings.Split(gt.GetRelatedTopLevelResourceTypeName(), "::")[0:2],
		"_"))
}

// GoFileName returns the Go file name for this type.
func (gt *GeneratorType) GoFileName() string {
	return strings.ToLower(gt.GoName()) + ".go"
}

// GoImports returns the Go imports for this type.
func (gt *GeneratorType) GoImports() []string {
	ic := newImportsCollector()
	ic.collectImports(gt.g.o.TypeTemplateGetImplicitGoImports(gt)...)

	for _, ga := range gt.Attributes {
		ga.collectImports(ic)
	}

	for _, gp := range gt.Properties {
		gp.collectImports(ic)
	}

	return ic.toSlice()
}

// GoSupportBasePackage returns the base package for the support library.
func (gt *GeneratorType) GoSupportBasePackage() string {
	return gt.g.o.getGoSupportBasePackage()
}

// DocumentationURL returns the documentation URL for this type.
func (gt *GeneratorType) DocumentationURL() string {
	return gt.specT.Documentation
}

// IsTopLevelResourceType returns true if this is a top-level resource type, false if it is a structured type.
func (gt *GeneratorType) IsTopLevelResourceType() bool {
	return gt.specT.IsTopLevelResourceType
}

// Name returns the CloudFormation name for this type.
func (gt *GeneratorType) Name() string {
	return gt.specT.Name
}

func (gt *GeneratorType) generate(outDirPath string) error {
	buf, err := tplz.ExecuteGo(gt.g.o.TypeTemplate, gt)
	if err != nil {
		return errorz.Wrap(err)
	}

	outDirPath = filepath.Join(outDirPath, gt.GoPackageName())
	if err := os.MkdirAll(outDirPath, 0777); err != nil {
		return errorz.Wrap(err)
	}

	return errorz.MaybeWrap(os.WriteFile(filepath.Join(outDirPath, gt.GoFileName()), buf, 0666))
}
