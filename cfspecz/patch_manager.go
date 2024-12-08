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
	Apply(ic *SpecIssueCollector, s *Spec)
}

// SpecPatchFunc is a func shorthand for SpecPatch.
type SpecPatchFunc func(ic *SpecIssueCollector, s *Spec)

// Apply implements the SpecPatch interface.
func (f SpecPatchFunc) Apply(ic *SpecIssueCollector, s *Spec) {
	f(ic, s)
}

// TypePatch describes a patch applicable to a type.
type TypePatch interface {
	Apply(ic *SpecIssueCollector, t *Type)
}

// TypePatchFunc is a func shorthand for TypePatch.
type TypePatchFunc func(ic *SpecIssueCollector, t *Type)

// Apply implements the SpecPatch interface.
func (f TypePatchFunc) Apply(ic *SpecIssueCollector, t *Type) {
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

func (m *PatchManager) applySpecPatches(ic *SpecIssueCollector, s *Spec) {
	for _, patch := range m.specPatches {
		patch.Apply(ic, s)
	}
}

func (m *PatchManager) applyTypePatches(ic *SpecIssueCollector, t *Type) {
	for _, patch := range m.typePatches[t.Name] {
		patch.Apply(ic, t)
	}
}

// TypeRelatedFields describes a set of type-related fields.
type TypeRelatedFields struct {
	PrimitiveType     string
	Type              string
	PrimitiveItemType string
	ItemType          string
}

// RawPatchMalformedType implements a category of raw patches.
type RawPatchMalformedType struct {
	TypeName      string
	PropertyFixes []*RawPatchMalformedTypePropertyFix
}

// RawPatchMalformedTypePropertyFix describes how to fix a property.
type RawPatchMalformedTypePropertyFix struct {
	TypeName       string
	PropertyName   string
	ExpectedFields *TypeRelatedFields
	FixedFields    *TypeRelatedFields
}

// Apply implements the RawPatch interface.
func (p *RawPatchMalformedType) Apply(rawSpec *gabs.Container) error {
	if err := p.removeMalformedType(rawSpec); err != nil {
		return errorz.Wrap(err)
	}

	for _, propertyFix := range p.PropertyFixes {
		if err := p.fixProperty(rawSpec, propertyFix); err != nil {
			return errorz.Wrap(err)
		}
	}

	return nil
}

func (p *RawPatchMalformedType) removeMalformedType(rawSpec *gabs.Container) error {
	path := []string{
		memz.Ternary(strings.Contains(p.TypeName, "."), "PropertyTypes", "ResourceTypes"),
		p.TypeName,
	}

	if !rawSpec.Exists(path...) {
		return errorz.Errorf("malformed type raw patch: missing path: %v", jsonz.MustMarshalString(path))
	}

	if err := rawSpec.Delete(path...); err != nil {
		return errorz.Wrap(err)
	}

	return nil
}

func (p *RawPatchMalformedType) fixProperty(rawSpec *gabs.Container, propertyFix *RawPatchMalformedTypePropertyFix) error {
	basePath := []string{
		memz.Ternary(strings.Contains(propertyFix.TypeName, "."), "PropertyTypes", "ResourceTypes"),
		propertyFix.TypeName,
		"Properties",
		propertyFix.PropertyName,
	}

	if !rawSpec.Exists(basePath...) {
		return errorz.Errorf("malformed type raw patch: missing path: %v", jsonz.MustMarshalString(basePath))
	}

	for expectedValue, path := range map[string][]string{
		propertyFix.ExpectedFields.PrimitiveType:     append(memz.ShallowCopySlice(basePath), "PrimitiveType"),
		propertyFix.ExpectedFields.Type:              append(memz.ShallowCopySlice(basePath), "Type"),
		propertyFix.ExpectedFields.PrimitiveItemType: append(memz.ShallowCopySlice(basePath), "PrimitiveItemType"),
		propertyFix.ExpectedFields.ItemType:          append(memz.ShallowCopySlice(basePath), "ItemType"),
	} {
		if v := rawSpec.Search(path...).Data(); reflect.DeepEqual(expectedValue, memz.Ternary(memz.IsAnyNil(v), "", v)) {
			return errorz.Errorf("malformed type raw patch: unexpected value for path: %v: expected '%v', got '%v'",
				jsonz.MustMarshalString(path),
				jsonz.MustMarshalString(propertyFix.ExpectedFields.PrimitiveType),
				jsonz.MustMarshalString(v))
		}
	}

	for fixedValue, path := range map[string][]string{
		propertyFix.FixedFields.PrimitiveType:     append(memz.ShallowCopySlice(basePath), "PrimitiveType"),
		propertyFix.FixedFields.Type:              append(memz.ShallowCopySlice(basePath), "Type"),
		propertyFix.FixedFields.PrimitiveItemType: append(memz.ShallowCopySlice(basePath), "PrimitiveItemType"),
		propertyFix.FixedFields.ItemType:          append(memz.ShallowCopySlice(basePath), "ItemType"),
	} {
		_, err := rawSpec.Set(fixedValue, path...)
		return errorz.MaybeWrap(err)
	}

	return nil
}

// SpecPatchDeleteType implements a category of spec patches.
type SpecPatchDeleteType struct {
	TypeName              string
	ForceIsStructuredType *bool
}

// Apply implements the SpecPatch interface.
func (p *SpecPatchDeleteType) Apply(ic *SpecIssueCollector, s *Spec) {
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

	ic.CollectIssue(s, "delete type spec patch: missing type: '%v'", p.TypeName)
}

// TypePatchDeleteAttribute implements a category of type patches.
type TypePatchDeleteAttribute struct {
	AttributeName  string
	ExpectedFields *TypeRelatedFields
}

// Apply implements the TypePatch interface.
func (p *TypePatchDeleteAttribute) Apply(ic *SpecIssueCollector, t *Type) {
	a, ok := t.Attributes[p.AttributeName]
	if !ok {
		ic.CollectIssue(t, "delete attribute type patch: missing attribute: '%v'", p.AttributeName)
		return
	}

	isMatching := true

	if p.ExpectedFields.PrimitiveType != a.PrimitiveType {
		isMatching = false
		ic.CollectIssue(t, "delete attribute type patch: unexpected value for PrimitiveType: expected '%v', got '%v'",
			p.ExpectedFields.PrimitiveType, a.PrimitiveType)
	}

	if p.ExpectedFields.Type != a.Type {
		isMatching = false
		ic.CollectIssue(t, "delete attribute type patch: unexpected value for Type: expected '%v', got '%v'",
			p.ExpectedFields.Type, a.Type)
	}

	if p.ExpectedFields.PrimitiveItemType != a.PrimitiveItemType {
		isMatching = false
		ic.CollectIssue(t, "delete attribute type patch: unexpected value for PrimitiveItemType: expected '%v', got '%v'",
			p.ExpectedFields.PrimitiveItemType, a.PrimitiveItemType)
	}

	if p.ExpectedFields.ItemType != a.ItemType {
		isMatching = false
		ic.CollectIssue(t, "delete attribute type patch: unexpected value for ItemType: expected '%v', got '%v'",
			p.ExpectedFields.ItemType, a.ItemType)
	}

	if isMatching {
		delete(t.Attributes, p.AttributeName)
	}
}
