package cfgenz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// GeneratorProperty describes a property of either a top-level resource or structured type in the generator g.
type GeneratorProperty struct {
	parent         *GeneratorType
	mustLookupType func(unqualifiedStructuredTypeName string) *GeneratorType
	specP          *cfspecz.Property
}

func newGeneratorProperty(gt *GeneratorType, specP *cfspecz.Property) *GeneratorProperty {
	return &GeneratorProperty{
		parent:         gt,
		mustLookupType: gt.g.makeMustLookupType(gt, specP),
		specP:          specP,
	}
}

// GoName returns the Go name for this property.
func (gp *GeneratorProperty) GoName() string {
	s := []rune(strings.ReplaceAll(gp.specP.Name, ".", "_"))
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
	return gp.specP.Documentation
}

// Name returns the CloudFormation name for this property.
func (gp *GeneratorProperty) Name() string {
	return gp.specP.Name
}

// Required returns true if the property is required, false otherwise.
func (gp *GeneratorProperty) Required() bool {
	return gp.specP.Required
}

// UpdateType returns the update type of the property ("Mutable", "Immutable", or "Conditional").
func (gp *GeneratorProperty) UpdateType() string {
	return gp.specP.UpdateType
}

func (gp *GeneratorProperty) goType(ic *importsCollector) string {
	return gp.parent.g.o.getGoSupportType(ic,
		fmt.Sprintf("%v[%v]", gp.goUnqualifiedOuterType(), gp.goGenericType(ic)))
}

func (gp *GeneratorProperty) goUnqualifiedOuterType() string {
	if gp.specP.Type == "List" {
		return "ExpressionSlice"
	}
	if gp.specP.Type == "Map" {
		return "ExpressionMap"
	}

	return "Expression"
}

func (gp *GeneratorProperty) goGenericType(ic *importsCollector) string {
	if gp.specP.PrimitiveType != "" {
		return gp.parent.g.getPrimitiveGoType(ic, gp.specP.PrimitiveType)
	}

	if gp.specP.Type == "List" || gp.specP.Type == "Map" {
		if gp.specP.PrimitiveItemType != "" {
			return gp.parent.g.getPrimitiveGoType(ic, gp.specP.PrimitiveItemType)
		}

		if gp.specP.ItemType == "Tag" {
			return gp.parent.g.o.getGoSupportType(ic, "Tag")
		}

		return gp.mustLookupType(gp.specP.ItemType).GoName()
	}

	if gp.specP.Type == "Tag" {
		return gp.parent.g.o.getGoSupportType(ic, "Tag")
	}

	return gp.mustLookupType(gp.specP.Type).GoName()
}

func (gp *GeneratorProperty) collectImports(ic *importsCollector) {
	gp.goType(ic)
}
