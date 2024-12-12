package cfgenz

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ibrt/golang-cloudformation/cfspecz"
)

// Property describes a property of either a top-level resource or structured type.
type Property struct {
	parent         *Type
	mustLookupType func(unqualifiedStructuredTypeName string) *Type
	specP          *cfspecz.Property
}

func newProperty(t *Type, specP *cfspecz.Property) *Property {
	return &Property{
		parent:         t,
		mustLookupType: t.g.makeMustLookupType(t, specP),
		specP:          specP,
	}
}

// GoName returns the Go name for this property.
func (p *Property) GoName() string {
	s := []rune(strings.ReplaceAll(p.specP.Name, ".", "_"))
	s[0] = unicode.ToTitle(s[0])
	return string(s)
}

// GoType returns the Go type for this property.
func (p *Property) GoType() string {
	return p.goType(nil)
}

// GoGenericType returns the Go generic type for this property.
func (p *Property) GoGenericType() string {
	return p.goGenericType(nil)
}

// GoUnqualifiedOuterType returns the Go unqualified outer type for this property
// (i.e. Expression, ExpressionSlice, or ExpressionMap).
func (p *Property) GoUnqualifiedOuterType() string {
	return p.goUnqualifiedOuterType()
}

// DocumentationURL returns the documentation URL for this property.
func (p *Property) DocumentationURL() string {
	return p.specP.Documentation
}

// Name returns the CloudFormation name for this property.
func (p *Property) Name() string {
	return p.specP.Name
}

// Required returns true if the property is required, false otherwise.
func (p *Property) Required() bool {
	return p.specP.Required
}

// UpdateType returns the update type of the property ("Mutable", "Immutable", or "Conditional").
func (p *Property) UpdateType() string {
	return p.specP.UpdateType
}

func (p *Property) goType(ic *importsCollector) string {
	return p.parent.g.o.getGoSupportType(ic,
		fmt.Sprintf("%v[%v]", p.goUnqualifiedOuterType(), p.goGenericType(ic)))
}

func (p *Property) goUnqualifiedOuterType() string {
	if p.specP.Type == "List" {
		return "ExpressionSlice"
	}
	if p.specP.Type == "Map" {
		return "ExpressionMap"
	}

	return "Expression"
}

func (p *Property) goGenericType(ic *importsCollector) string {
	if p.specP.PrimitiveType != "" {
		return p.parent.g.getPrimitiveGoType(ic, p.specP.PrimitiveType)
	}

	if p.specP.Type == "List" || p.specP.Type == "Map" {
		if p.specP.PrimitiveItemType != "" {
			return p.parent.g.getPrimitiveGoType(ic, p.specP.PrimitiveItemType)
		}

		if p.specP.ItemType == "Tag" {
			return p.parent.g.o.getGoSupportType(ic, "Tag")
		}

		return p.mustLookupType(p.specP.ItemType).GoName()
	}

	if p.specP.Type == "Tag" {
		return p.parent.g.o.getGoSupportType(ic, "Tag")
	}

	return p.mustLookupType(p.specP.Type).GoName()
}

func (p *Property) collectImports(ic *importsCollector) {
	p.goType(ic)
}
