package cfspecz

import (
	"github.com/ibrt/golang-utils/jsonz"
)

// SpecPatch describes a patch applicable to a spec.
type SpecPatch func(ic *SpecIssueCollector, s *Spec)

// TypePatch describes a patch applicable to a type.
type TypePatch func(ic *SpecIssueCollector, t *Type)

// SpecPatchManager manages a set of patches for a spec.
type SpecPatchManager struct {
	specPatches []SpecPatch
	typePatches map[string][]TypePatch
}

// NewSpecPatchManager initializes a new spec patch manager.
func NewSpecPatchManager() *SpecPatchManager {
	return &SpecPatchManager{
		specPatches: make([]SpecPatch, 0),
		typePatches: make(map[string][]TypePatch),
	}
}

// NewDefaultSpecPatchManager initializes a new spec patch manager and configures it with default, known patchers.
func NewDefaultSpecPatchManager() *SpecPatchManager {
	return NewSpecPatchManager().
		RegisterSpecPatch(
			func(ic *SpecIssueCollector, s *Spec) {
				if _, ok := s.PropertyTypes["Tag"]; ok {
					delete(s.PropertyTypes, "Tag")
					return
				}

				ic.CollectIssue(s, "outdated patch: missing property type: 'Tag'")
			}).
		RegisterTypePatchDeleteAttribute(
			"AWS::CleanRooms::IdNamespaceAssociation",
			"InputReferenceProperties.IdMappingWorkflowsSupported",
			func(a *Attribute) bool { return a.Type == "List" && a.PrimitiveItemType == "Json" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::CloudFormation::WaitCondition",
			"Data",
			func(a *Attribute) bool { return a.PrimitiveType == "Json" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::IoTTwinMaker::Scene",
			"GeneratedSceneMetadata",
			func(a *Attribute) bool { return a.Type == "Map" && a.PrimitiveItemType == "String" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::MediaLive::SignalMap",
			"MediaResourceMap",
			func(a *Attribute) bool { return a.Type == "Map" && a.ItemType == "MediaResource" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::MediaLive::SignalMap",
			"FailedMediaResourceMap",
			func(a *Attribute) bool { return a.Type == "Map" && a.ItemType == "MediaResource" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::OpenSearchService::Domain",
			"DomainEndpoints",
			func(a *Attribute) bool { return a.Type == "Map" && a.PrimitiveItemType == "String" }).
		RegisterTypePatchDeleteAttribute(
			"AWS::ServiceCatalog::CloudFormationProvisionedProduct",
			"Outputs",
			func(a *Attribute) bool { return a.Type == "Map" && a.PrimitiveItemType == "String" })
}

// RegisterSpecPatch registers a spec patch.
func (m *SpecPatchManager) RegisterSpecPatch(patch SpecPatch) *SpecPatchManager {
	m.specPatches = append(m.specPatches, patch)
	return m
}

// RegisterTypePatch registers a type patch.
func (m *SpecPatchManager) RegisterTypePatch(typeName string, patch TypePatch) *SpecPatchManager {
	typePatchesForType := m.typePatches[typeName]
	typePatchesForType = append(typePatchesForType, patch)
	m.typePatches[typeName] = typePatchesForType
	return m
}

// RegisterTypePatchDeleteAttribute registers a "delete attribute" type patch.
func (m *SpecPatchManager) RegisterTypePatchDeleteAttribute(typeName, attributeName string, cond func(a *Attribute) bool) *SpecPatchManager {
	return m.RegisterTypePatch(typeName, func(ic *SpecIssueCollector, t *Type) {
		a, ok := t.Attributes[attributeName]
		if !ok {
			ic.CollectIssue(t, "outdated patch: missing attribute: '%v'", attributeName)
			return
		}

		if !cond(a) {
			ic.CollectIssue(t,
				"outdated patch: attribute '%v' does not match condition (%v)",
				attributeName,
				jsonz.MustMarshalPrettyString(a))

			return
		}

		delete(t.Attributes, attributeName)
	})
}

func (m *SpecPatchManager) applySpecPatches(ic *SpecIssueCollector, s *Spec) {
	for _, patch := range m.specPatches {
		patch(ic, s)
	}
}

func (m *SpecPatchManager) applyTypePatches(ic *SpecIssueCollector, t *Type) {
	for _, patch := range m.typePatches[t.Name] {
		patch(ic, t)
	}
}
