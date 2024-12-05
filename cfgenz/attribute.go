package cfgenz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ibrt/golang-utils/memz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// GeneratorAttribute describes an attribute of a top-level resource type in the generator spec.
type GeneratorAttribute struct {
	parent         *GeneratorType
	mustLookupType func(unqualifiedStructuredTypeName string) *GeneratorType
	a              *cfspecz.Attribute
}

func newGeneratorAttribute(gt *GeneratorType, a *cfspecz.Attribute) *GeneratorAttribute {
	return &GeneratorAttribute{
		parent:         gt,
		mustLookupType: gt.spec.makeMustLookupType(gt, a),
		a:              a,
	}
}

// GoName returns the Go name for this attribute.
func (ga *GeneratorAttribute) GoName() string {
	s := []rune(strings.ReplaceAll(ga.a.Name, ".", "_"))
	s[0] = unicode.ToTitle(s[0])
	return string(s)
}

// GoType returns the Go type for this attribute.
func (ga *GeneratorAttribute) GoType() string {
	return ga.goType(nil)
}

// GoGenericType returns the Go generic type for this attribute.
func (ga *GeneratorAttribute) GoGenericType() string {
	return ga.goGenericType(nil)
}

// SupportGetAttFunctionName returns the GetAtt* support function to use for this attribute.
func (ga *GeneratorAttribute) SupportGetAttFunctionName() string {
	return memz.Ternary(ga.a.Type == "List", "GetAttSlice", "GetAtt")
}

// Name returns the CloudFormation name for this attribute.
func (ga *GeneratorAttribute) Name() string {
	return ga.a.Name
}

func (ga *GeneratorAttribute) goType(ic *importsCollector) string {
	if ga.a.Type == "List" {
		return ga.parent.spec.o.getGoSupportType(ic,
			fmt.Sprintf("ExpressionList[%v]", ga.goGenericType(ic)))
	}

	return ga.parent.spec.o.getGoSupportType(ic,
		fmt.Sprintf("Expression[%v]", ga.goGenericType(ic)))
}

func (ga *GeneratorAttribute) goGenericType(ic *importsCollector) string {
	if ga.a.PrimitiveType != "" {
		return ga.parent.spec.getPrimitiveGoType(ic, ga.a.PrimitiveType)
	}

	if ga.a.Type == "List" {
		if ga.a.PrimitiveItemType != "" {
			return ga.parent.spec.getPrimitiveGoType(ic, ga.a.PrimitiveItemType)
		}

		if ga.a.ItemType == "Tag" {
			return ga.parent.spec.o.getGoSupportType(ic, "Tag")
		}

		return ga.mustLookupType(ga.a.ItemType).GoName()
	}

	if ga.a.Type == "Tag" {
		return ga.parent.spec.o.getGoSupportType(ic, "Tag")
	}

	return ga.mustLookupType(ga.a.Type).GoName()
}

func (ga *GeneratorAttribute) collectImports(ic *importsCollector) {
	ga.goType(ic)
}
