package specz

import (
	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/jsonz"
)

// Spec describes the CloudFormation spec.
type Spec struct {
	ResourceSpecificationVersion string
	ResourceTypes                map[string]*Type
	PropertyTypes                map[string]*Type
}

func (s *Spec) preProcess() {
	delete(s.PropertyTypes, "Tag")

	for resourceTypeName, resourceType := range s.ResourceTypes {
		resourceType.preProcess(s, true, resourceTypeName)
	}

	for propertyTypeName, propertyType := range s.PropertyTypes {
		propertyType.preProcess(s, false, propertyTypeName)
	}
}

func (s *Spec) validate() error {
	if s.ResourceSpecificationVersion == "" {
		return errorz.Wrap(NewSpecValidationErrorByMessage("missing ResourceSpecificationVersion"))
	}

	ic := newIssuesCollector()

	for _, resourceType := range s.ResourceTypes {
		resourceType.validateAndCollectIssues(ic)
	}

	for _, propertyType := range s.PropertyTypes {
		propertyType.validateAndCollectIssues(ic)
	}

	return errorz.MaybeWrap(ic.maybeToError())
}

// ParseSpec parses a CloudFormation spec from the given buffer (i.e. "CloudFormationResourceSpecification.json" file).
func ParseSpec(buf []byte) (*Spec, error) {
	spec, err := jsonz.Unmarshal[*Spec](buf)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	spec.preProcess()

	if err := spec.validate(); err != nil {
		return nil, errorz.Wrap(err)
	}

	return spec, nil
}
