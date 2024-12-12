package cfgenz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ibrt/golang-utils/memz"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// Attribute describes an attribute of a top-level resource type.
type Attribute struct {
	parent         *Type
	mustLookupType func(unqualifiedStructuredTypeName string) *Type
	specA          *cfspecz.Attribute
}

func newAttribute(t *Type, specA *cfspecz.Attribute) *Attribute {
	return &Attribute{
		parent:         t,
		mustLookupType: t.g.makeMustLookupType(t, specA),
		specA:          specA,
	}
}

// GoName returns the Go name for this attribute.
func (a *Attribute) GoName() string {
	s := []rune(strings.ReplaceAll(a.specA.Name, ".", "_"))
	s[0] = unicode.ToTitle(s[0])
	return string(s)
}

// GoType returns the Go type for this attribute.
func (a *Attribute) GoType() string {
	return a.goType(nil)
}

// GoGenericType returns the Go generic type for this attribute.
func (a *Attribute) GoGenericType() string {
	return a.goGenericType(nil)
}

// SupportGetAttFunctionName returns the GetAtt* support function to use for this attribute.
func (a *Attribute) SupportGetAttFunctionName() string {
	return memz.Ternary(a.specA.Type == "List", "GetAttSlice", "GetAtt")
}

// SupportOutputFunctionName returns the *Output* support function to use for this attribute.
func (a *Attribute) SupportOutputFunctionName() string {
	return memz.Ternary(a.specA.Type == "List", "OutputSlice", "Output")
}

// NameForLogicalNames returns the name for this attribute modified so that it can be embedded in logical names.
func (a *Attribute) NameForLogicalNames() string {
	return strings.ReplaceAll(a.specA.Name, ".", "")
}

// Name returns the CloudFormation name for this attribute.
func (a *Attribute) Name() string {
	return a.specA.Name
}

func (a *Attribute) goType(ic *importsCollector) string {
	if a.specA.Type == "List" {
		return a.parent.g.o.getGoSupportType(ic,
			fmt.Sprintf("ExpressionSlice[%v]", a.goGenericType(ic)))
	}

	return a.parent.g.o.getGoSupportType(ic,
		fmt.Sprintf("Expression[%v]", a.goGenericType(ic)))
}

func (a *Attribute) goGenericType(ic *importsCollector) string {
	if a.specA.PrimitiveType != "" {
		return a.parent.g.getPrimitiveGoType(ic, a.specA.PrimitiveType)
	}

	if a.specA.Type == "List" {
		if a.specA.PrimitiveItemType != "" {
			return a.parent.g.getPrimitiveGoType(ic, a.specA.PrimitiveItemType)
		}

		if a.specA.ItemType == "Tag" {
			return a.parent.g.o.getGoSupportType(ic, "Tag")
		}

		return a.mustLookupType(a.specA.ItemType).GoName()
	}

	if a.specA.Type == "Tag" {
		return a.parent.g.o.getGoSupportType(ic, "Tag")
	}

	return a.mustLookupType(a.specA.Type).GoName()
}

func (a *Attribute) collectImports(ic *importsCollector) {
	a.goType(ic)
}
