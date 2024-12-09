package cfspecz

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/jsonz"
	"github.com/ibrt/golang-utils/memz"
)

// RawPatch describes a patch applicable to the raw spec (i.e. before parsing into structs).
type RawPatch interface {
	Apply(rawSpec *gabs.Container) error
}

// RawPatchFunc is a func shorthand for RawPatch.
type RawPatchFunc func(rawSpec *gabs.Container) error

// Apply implements the RawPatch interface.
func (f RawPatchFunc) Apply(rawSpec *gabs.Container) error {
	return errorz.MaybeWrap(f(rawSpec))
}

// SpecPatch describes a patch applicable to a spec.
type SpecPatch interface {
	Apply(ic *ProblemsCollector, s *Spec)
}

// SpecPatchFunc is a func shorthand for SpecPatch.
type SpecPatchFunc func(ic *ProblemsCollector, s *Spec)

// Apply implements the SpecPatch interface.
func (f SpecPatchFunc) Apply(ic *ProblemsCollector, s *Spec) {
	f(ic, s)
}

// TypePatch describes a patch applicable to a type.
type TypePatch interface {
	Apply(ic *ProblemsCollector, t *Type)
}

// TypePatchFunc is a func shorthand for TypePatch.
type TypePatchFunc func(ic *ProblemsCollector, t *Type)

// Apply implements the SpecPatch interface.
func (f TypePatchFunc) Apply(ic *ProblemsCollector, t *Type) {
	f(ic, t)
}

// PatchManager registers and applies patches for a spec.
type PatchManager struct {
	rawPatches  []RawPatch
	specPatches []SpecPatch
	typePatches map[string][]TypePatch
}

// NewPatchManager initializes a new patch manager.
func NewPatchManager() *PatchManager {
	return &PatchManager{
		rawPatches:  make([]RawPatch, 0),
		specPatches: make([]SpecPatch, 0),
		typePatches: make(map[string][]TypePatch),
	}
}

// RegisterRawPatch registers a raw patch.
func (m *PatchManager) RegisterRawPatch(patch RawPatch) *PatchManager {
	m.rawPatches = append(m.rawPatches, patch)
	return m
}

// RegisterSpecPatch registers a spec patch.
func (m *PatchManager) RegisterSpecPatch(patch SpecPatch) *PatchManager {
	m.specPatches = append(m.specPatches, patch)
	return m
}

// RegisterTypePatch registers a type patch.
func (m *PatchManager) RegisterTypePatch(typeName string, patch TypePatch) *PatchManager {
	typePatchesForType := m.typePatches[typeName]
	typePatchesForType = append(typePatchesForType, patch)
	m.typePatches[typeName] = typePatchesForType
	return m
}

func (m *PatchManager) applyRawPatches(raw *gabs.Container) error {
	for i, patch := range m.rawPatches {
		if err := patch.Apply(raw); err != nil {
			return errorz.Wrap(err, fmt.Errorf("raw patch #%v", i))
		}
	}

	return nil
}

func (m *PatchManager) applySpecPatches(ic *ProblemsCollector, s *Spec) {
	for _, patch := range m.specPatches {
		patch.Apply(ic, s)
	}
}

func (m *PatchManager) applyTypePatches(ic *ProblemsCollector, t *Type) {
	for _, patch := range m.typePatches[t.Name] {
		patch.Apply(ic, t)
	}
}

// PropertyOrAttributeTypeFields describes a set of type-related fields.
type PropertyOrAttributeTypeFields struct {
	PrimitiveType     string
	Type              string
	PrimitiveItemType string
	ItemType          string
}

// RawPatchDeleteType implements a category of raw patches.
type RawPatchDeleteType struct {
	TypeName string
}

// Apply implements the RawPatch interface.
func (p *RawPatchDeleteType) Apply(rawSpec *gabs.Container) error {
	path := []string{
		memz.Ternary(strings.Contains(p.TypeName, "."), "PropertyTypes", "ResourceTypes"),
		p.TypeName,
	}

	if !rawSpec.Exists(path...) {
		return errorz.Errorf("delete type raw patch: missing path: %v", jsonz.MustMarshalString(path))
	}

	if err := rawSpec.Delete(path...); err != nil {
		return errorz.Wrap(err)
	}

	return nil
}

// RawPatchFixPropertyType implements a category of raw patches.
type RawPatchFixPropertyType struct {
	TypeName       string
	PropertyName   string
	ExpectedFields *PropertyOrAttributeTypeFields
	FixedFields    *PropertyOrAttributeTypeFields
}

// Apply implements the RawPatch interface.
func (p *RawPatchFixPropertyType) Apply(rawSpec *gabs.Container) error {
	basePath := []string{
		memz.Ternary(strings.Contains(p.TypeName, "."), "PropertyTypes", "ResourceTypes"),
		p.TypeName,
		"Properties",
		p.PropertyName,
	}

	if !rawSpec.Exists(basePath...) {
		return errorz.Errorf("fix property type raw patch: missing path: %v", jsonz.MustMarshalString(basePath))
	}

	for _, m := range []struct {
		path          []string
		expectedValue string
		fixedValue    string
	}{
		{
			path:          append(memz.ShallowCopySlice(basePath), "PrimitiveType"),
			expectedValue: p.ExpectedFields.PrimitiveType,
			fixedValue:    p.FixedFields.PrimitiveType,
		},
		{
			path:          append(memz.ShallowCopySlice(basePath), "Type"),
			expectedValue: p.ExpectedFields.Type,
			fixedValue:    p.FixedFields.Type,
		},
		{
			path:          append(memz.ShallowCopySlice(basePath), "PrimitiveItemType"),
			expectedValue: p.ExpectedFields.PrimitiveItemType,
			fixedValue:    p.FixedFields.PrimitiveItemType,
		},
		{
			path:          append(memz.ShallowCopySlice(basePath), "ItemType"),
			expectedValue: p.ExpectedFields.ItemType,
			fixedValue:    p.FixedFields.ItemType,
		},
	} {
		if v := rawSpec.Search(m.path...).Data(); !reflect.DeepEqual(m.expectedValue, memz.Ternary(memz.IsAnyNil(v), "", v)) {
			return errorz.Errorf("malformed type raw patch: unexpected value for path: %v: expected '%v', got '%v'",
				jsonz.MustMarshalString(m.path),
				jsonz.MustMarshalString(p.ExpectedFields.PrimitiveType),
				jsonz.MustMarshalString(v))
		}

		if _, err := rawSpec.Set(m.fixedValue, m.path...); err != nil {
			return errorz.MaybeWrap(err)
		}
	}

	return nil
}

// SpecPatchDeleteType implements a category of spec patches.
type SpecPatchDeleteType struct {
	TypeName              string
	ForceIsStructuredType *bool
}

// Apply implements the SpecPatch interface.
func (p *SpecPatchDeleteType) Apply(ic *ProblemsCollector, s *Spec) {
	if strings.Contains(p.TypeName, ".") || memz.ValNilToZero(p.ForceIsStructuredType) {
		if _, ok := s.PropertyTypes[p.TypeName]; ok {
			delete(s.PropertyTypes, p.TypeName)
			return
		}
	} else {
		if _, ok := s.ResourceTypes[p.TypeName]; ok {
			delete(s.ResourceTypes, p.TypeName)
			return
		}
	}

	ic.Collect(s, "delete type spec patch: missing type: '%v'", p.TypeName)
}

// TypePatchDeleteAttribute implements a category of type patches.
type TypePatchDeleteAttribute struct {
	AttributeName  string
	ExpectedFields *PropertyOrAttributeTypeFields
}

// Apply implements the TypePatch interface.
func (p *TypePatchDeleteAttribute) Apply(ic *ProblemsCollector, t *Type) {
	a, ok := t.Attributes[p.AttributeName]
	if !ok {
		ic.Collect(t, "delete attribute type patch: missing attribute: '%v'", p.AttributeName)
		return
	}

	isMatching := true

	if p.ExpectedFields.PrimitiveType != a.PrimitiveType {
		isMatching = false
		ic.Collect(t, "delete attribute type patch: unexpected value for PrimitiveType: expected '%v', got '%v'",
			p.ExpectedFields.PrimitiveType, a.PrimitiveType)
	}

	if p.ExpectedFields.Type != a.Type {
		isMatching = false
		ic.Collect(t, "delete attribute type patch: unexpected value for Type: expected '%v', got '%v'",
			p.ExpectedFields.Type, a.Type)
	}

	if p.ExpectedFields.PrimitiveItemType != a.PrimitiveItemType {
		isMatching = false
		ic.Collect(t, "delete attribute type patch: unexpected value for PrimitiveItemType: expected '%v', got '%v'",
			p.ExpectedFields.PrimitiveItemType, a.PrimitiveItemType)
	}

	if p.ExpectedFields.ItemType != a.ItemType {
		isMatching = false
		ic.Collect(t, "delete attribute type patch: unexpected value for ItemType: expected '%v', got '%v'",
			p.ExpectedFields.ItemType, a.ItemType)
	}

	if isMatching {
		delete(t.Attributes, p.AttributeName)
	}
}
