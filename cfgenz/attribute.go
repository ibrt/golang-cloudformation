package cfgenz

import (
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
