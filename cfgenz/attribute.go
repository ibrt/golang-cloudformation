package cfgenz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ibrt/golang-utils/memz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// GeneratorAttribute describes an attribute of a top-level resource type in the generator g.
type GeneratorAttribute struct {
	parent         *GeneratorType
	mustLookupType func(unqualifiedStructuredTypeName string) *GeneratorType
	specA          *cfspecz.Attribute
}

func newGeneratorAttribute(gt *GeneratorType, specA *cfspecz.Attribute) *GeneratorAttribute {
	return &GeneratorAttribute{
		parent:         gt,
		mustLookupType: gt.g.makeMustLookupType(gt, specA),
		specA:          specA,
	}
}

// GoName returns the Go name for this attribute.
func (ga *GeneratorAttribute) GoName() string {
	s := []rune(strings.ReplaceAll(ga.specA.Name, ".", "_"))
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
	return memz.Ternary(ga.specA.Type == "List", "GetAttSlice", "GetAtt")
}

// SupportOutputFunctionName returns the *Output* support function to use for this attribute.
func (ga *GeneratorAttribute) SupportOutputFunctionName() string {
	return memz.Ternary(ga.specA.Type == "List", "OutputSlice", "Output")
}

// NameForLogicalNames returns the name for this attribute modified so that it can be embedded in logical names.
func (ga *GeneratorAttribute) NameForLogicalNames() string {
	return strings.ReplaceAll(ga.specA.Name, ".", "")
}

// Name returns the CloudFormation name for this attribute.
func (ga *GeneratorAttribute) Name() string {
	return ga.specA.Name
}

func (ga *GeneratorAttribute) goType(ic *importsCollector) string {
	if ga.specA.Type == "List" {
		return ga.parent.g.o.getGoSupportType(ic,
			fmt.Sprintf("ExpressionSlice[%v]", ga.goGenericType(ic)))
	}

	return ga.parent.g.o.getGoSupportType(ic,
		fmt.Sprintf("Expression[%v]", ga.goGenericType(ic)))
}

func (ga *GeneratorAttribute) goGenericType(ic *importsCollector) string {
	if ga.specA.PrimitiveType != "" {
		return ga.parent.g.getPrimitiveGoType(ic, ga.specA.PrimitiveType)
	}

	if ga.specA.Type == "List" {
		if ga.specA.PrimitiveItemType != "" {
			return ga.parent.g.getPrimitiveGoType(ic, ga.specA.PrimitiveItemType)
		}

		if ga.specA.ItemType == "Tag" {
			return ga.parent.g.o.getGoSupportType(ic, "Tag")
		}

		return ga.mustLookupType(ga.specA.ItemType).GoName()
	}

	if ga.specA.Type == "Tag" {
		return ga.parent.g.o.getGoSupportType(ic, "Tag")
	}

	return ga.mustLookupType(ga.specA.Type).GoName()
}

func (ga *GeneratorAttribute) collectImports(ic *importsCollector) {
	ga.goType(ic)
}
