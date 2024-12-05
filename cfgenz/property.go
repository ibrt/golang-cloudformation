package cfgenz

import (
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
