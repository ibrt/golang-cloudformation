package cfspecz

import (
	"bytes"
	"encoding/json"

	"github.com/Jeffail/gabs/v2"
	"github.com/ibrt/golang-utils/errorz"

	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	specDisplayPath = "spec"
)

var (
	_ cfz.ProblemLocation = (*Property)(nil)
)

// Spec describes the CloudFormation spec.
type Spec struct {
	ResourceSpecificationVersion string
	ResourceTypes                map[string]*Type
	PropertyTypes                map[string]*Type
}

// NewSpecFromBuffer parses, patches (using the given patch manager), and validates a CloudFormation spec from the
// given buffer (i.e. "CloudFormationResourceSpecification.json" file).
func NewSpecFromBuffer(buf []byte, pm *PatchManager) (*Spec, error) {
	d := json.NewDecoder(bytes.NewReader(buf))
	d.DisallowUnknownFields()
	d.UseNumber()

	rawSpec, err := gabs.ParseJSONDecoder(d)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	if err := pm.applyRawPatches(rawSpec); err != nil {
		return nil, errorz.Wrap(err)
	}

	d = json.NewDecoder(bytes.NewReader(rawSpec.EncodeJSON()))
	d.DisallowUnknownFields()

	var s *Spec

	if err := d.Decode(&s); err != nil {
		return nil, errorz.Wrap(err)
	}

	s.preProcess()

	if err := s.applyPatches(pm); err != nil {
		return nil, errorz.Wrap(err)
	}

	if err := s.collectProblems(); err != nil {
		return nil, errorz.Wrap(err)
	}

	return s, nil
}

// GetDisplayPath implements the ProblemLocation interface.
func (*Spec) GetDisplayPath() string {
	return specDisplayPath
}

func (s *Spec) preProcess() {
	for resourceTypeName, resourceType := range s.ResourceTypes {
		resourceType.preProcess(s, true, resourceTypeName)
	}

	for propertyTypeName, propertyType := range s.PropertyTypes {
		propertyType.preProcess(s, false, propertyTypeName)
	}
}

func (s *Spec) applyPatches(pm *PatchManager) error {
	pc := cfz.NewProblemsCollector()
	pm.applySpecPatches(pc, s)

	for _, resourceType := range s.ResourceTypes {
		resourceType.applyPatches(pc, pm)
	}

	for _, propertyType := range s.PropertyTypes {
		propertyType.applyPatches(pc, pm)
	}

	return errorz.MaybeWrap(pc.ToError())
}

func (s *Spec) collectProblems() error {
	pc := cfz.NewProblemsCollector()

	if s.ResourceSpecificationVersion == "" {
		pc.Collect(s, "missing ResourceSpecificationVersion")
		return pc.ToError()
	}

	for _, resourceType := range s.ResourceTypes {
		resourceType.collectProblems(pc)
	}

	for _, propertyType := range s.PropertyTypes {
		propertyType.collectProblems(pc)
	}

	for _, propertyType := range s.PropertyTypes {
		pc.MaybeCollect(propertyType, !propertyType.IsReferenced, "unreferenced structured type")
	}

	return errorz.MaybeWrap(pc.ToError())
}
