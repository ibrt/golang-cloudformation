package cfspecz

import (
	"bytes"
	"encoding/json"

	"github.com/Jeffail/gabs/v2"
	"github.com/ibrt/golang-utils/errorz"
)

const (
	specDisplayPath = "spec"
)

// SpecContext describes a place in the spec (i.e. spec, type, property, or attribute).
type SpecContext interface {
	GetDisplayPath() string
}

// Spec describes the CloudFormation spec.
type Spec struct {
	ResourceSpecificationVersion string
	ResourceTypes                map[string]*Type
	PropertyTypes                map[string]*Type
}

// NewSpecFromBuffer parses, patches (using the given spec patch manager), and validates a CloudFormation spec from the
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
	d.UseNumber()

	var spec *Spec

	if err := d.Decode(&spec); err != nil {
		return nil, errorz.Wrap(err)
	}

	spec.preProcess()

	if err := spec.applyPatches(pm); err != nil {
		return nil, errorz.Wrap(err)
	}

	if err := spec.collectIssues(); err != nil {
		return nil, errorz.Wrap(err)
	}

	return spec, nil
}

// GetDisplayPath implements the SpecContext interface.
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
	ic := NewSpecIssueCollector()
	pm.applySpecPatches(ic, s)

	for _, resourceType := range s.ResourceTypes {
		resourceType.applyPatches(ic, pm)
	}

	for _, propertyType := range s.PropertyTypes {
		propertyType.applyPatches(ic, pm)
	}

	return ic.MaybeToError()
}

func (s *Spec) collectIssues() error {
	ic := NewSpecIssueCollector()

	if s.ResourceSpecificationVersion == "" {
		ic.CollectIssue(s, "missing ResourceSpecificationVersion")
		return ic.MaybeToError()
	}

	for _, resourceType := range s.ResourceTypes {
		resourceType.collectIssues(ic)
	}

	for _, propertyType := range s.PropertyTypes {
		propertyType.collectIssues(ic)
	}

	return ic.MaybeToError()
}
