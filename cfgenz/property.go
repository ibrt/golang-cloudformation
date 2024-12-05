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

func (gp *GeneratorProperty) goType(ic *importsCollector) string {
	if gp.p.Type == "List" {
		return gp.parent.spec.o.getGoSupportType(ic,
			fmt.Sprintf("ExpressionList[%v]", gp.goGenericType(ic)))
	}

	if gp.p.Type == "Map" {
		return gp.parent.spec.o.getGoSupportType(ic,
			fmt.Sprintf("ExpressionMap[%v]", gp.goGenericType(ic)))
	}

	return gp.parent.spec.o.getGoSupportType(ic,
		fmt.Sprintf("Expression[%v]", gp.goGenericType(ic)))
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
