package cfgenz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// GeneratorProperty describes a property of either a top-level resource or structured type in the generator spec.
type GeneratorProperty struct {
	parent         *GeneratorType
	mustLookupType func(unqualifiedStructuredTypeName string) *GeneratorType
	p              *cfspecz.Property
}

func newGeneratorProperty(gt *GeneratorType, p *cfspecz.Property) *GeneratorProperty {
	return &GeneratorProperty{
		parent:         gt,
		mustLookupType: gt.spec.makeMustLookupType(gt, p),
		p:              p,
	}
}

// GoName returns the Go name for this property.
func (gp *GeneratorProperty) GoName() string {
	s := []rune(strings.ReplaceAll(gp.p.Name, ".", "_"))
	s[0] = unicode.ToTitle(s[0])
	return string(s)
}

// GoType returns the Go type for this property.
func (gp *GeneratorProperty) GoType() string {
	return gp.goType(nil)
}

// GoGenericType returns the Go generic type for this property.
func (gp *GeneratorProperty) GoGenericType() string {
	return gp.goGenericType(nil)
}

// GoUnqualifiedOuterType returns the Go unqualified outer type for this property
// (i.e. Expression, ExpressionSlice, or ExpressionMap).
func (gp *GeneratorProperty) GoUnqualifiedOuterType() string {
	return gp.goUnqualifiedOuterType()
}

// DocumentationURL returns the documentation URL for this property.
func (gp *GeneratorProperty) DocumentationURL() string {
	return gp.p.Documentation
}

// Name returns the CloudFormation name for this property.
func (gp *GeneratorProperty) Name() string {
	return gp.p.Name
}

func (gp *GeneratorProperty) goType(ic *importsCollector) string {
	return gp.parent.spec.o.getGoSupportType(ic,
		fmt.Sprintf("%v[%v]", gp.goUnqualifiedOuterType(), gp.goGenericType(ic)))
}

func (gp *GeneratorProperty) goUnqualifiedOuterType() string {
	if gp.p.Type == "List" {
		return "ExpressionSlice"
	}
	if gp.p.Type == "Map" {
		return "ExpressionMap"
	}

	return "Expression"
}

func (gp *GeneratorProperty) goGenericType(ic *importsCollector) string {
	if gp.p.PrimitiveType != "" {
		return gp.parent.spec.getPrimitiveGoType(ic, gp.p.PrimitiveType)
	}

	if gp.p.Type == "List" || gp.p.Type == "Map" {
		if gp.p.PrimitiveItemType != "" {
			return gp.parent.spec.getPrimitiveGoType(ic, gp.p.PrimitiveItemType)
		}

		if gp.p.ItemType == "Tag" {
			return gp.parent.spec.o.getGoSupportType(ic, "Tag")
		}

		return gp.mustLookupType(gp.p.ItemType).GoName()
	}

	if gp.p.Type == "Tag" {
		return gp.parent.spec.o.getGoSupportType(ic, "Tag")
	}

	return gp.mustLookupType(gp.p.Type).GoName()
}

func (gp *GeneratorProperty) collectImports(ic *importsCollector) {
	gp.goType(ic)
}
