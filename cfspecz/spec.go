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
	TopLevelResourceTypes        map[string]*Type `json:"ResourceTypes"`
	StructuredTypes              map[string]*Type `json:"PropertyTypes"`
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
	for name, t := range s.TopLevelResourceTypes {
		t.preProcess(s, true, name)
	}

	for name, t := range s.StructuredTypes {
		t.preProcess(s, false, name)
	}
}

func (s *Spec) applyPatches(pm *PatchManager) error {
	pc := cfz.NewProblemsCollector()
	pm.applySpecPatches(pc, s)

	for _, t := range s.TopLevelResourceTypes {
		t.applyPatches(pc, pm)
	}

	for _, t := range s.StructuredTypes {
		t.applyPatches(pc, pm)
	}

	return errorz.MaybeWrap(pc.ToError())
}

func (s *Spec) collectProblems() error {
	pc := cfz.NewProblemsCollector()

	if s.ResourceSpecificationVersion == "" {
		pc.Collect(s, "missing ResourceSpecificationVersion")
		return pc.ToError()
	}

	for _, t := range s.TopLevelResourceTypes {
		t.collectProblems(pc)
	}

	for _, t := range s.StructuredTypes {
		t.collectProblems(pc)
	}

	for _, t := range s.StructuredTypes {
		pc.MaybeCollect(t, !t.IsReferenced, "unreferenced structured type")
	}

	return errorz.MaybeWrap(pc.ToError())
}
