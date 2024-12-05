package cfspecz

import (
	"fmt"
	"strings"

	"github.com/ibrt/golang-utils/memz"
)

// Type describes either a top-level resource type or a structured type in the spec.
type Type struct {
	Documentation string
	Attributes    map[string]*Attribute
	Properties    map[string]*Property

	IsTopLevelResourceType bool   `json:"-"`
	Name                   string `json:"-"`
}

// GetRelatedTopLevelResourceTypeName returns the top-level resource type name related to this type, that is, Name if
// IsTopLevelResourceType = true, or the portion of Name before the first "." otherwise.
func (t *Type) GetRelatedTopLevelResourceTypeName() string {
	if t.IsTopLevelResourceType {
		return t.Name
	}
	return strings.Split(t.Name, ".")[0]
}

// GetRelatedStructuredTypeName converts the value of a Type or ItemType field of a property or attribute in this type
// to its fully qualified structured type name.
func (t *Type) GetRelatedStructuredTypeName(unqualifiedStructuredTypeName string) string {
	return fmt.Sprintf("%s.%s", t.GetRelatedTopLevelResourceTypeName(), unqualifiedStructuredTypeName)
}

// GetDisplayPath implements the SpecContext interface.
func (t *Type) GetDisplayPath() string {
	return fmt.Sprintf("%v/%v[%v]",
		specDisplayPath,
		memz.Ternary(t.IsTopLevelResourceType, "topLevelResourceType", "structuredType"),
		t.Name)
}

func (t *Type) preProcess(spec *Spec, isTopLevelResourceType bool, name string) {
	t.IsTopLevelResourceType = isTopLevelResourceType
	t.Name = name

	for attributeName, attribute := range t.Attributes {
		attribute.preProcess(spec, t, attributeName)
	}

	for propertyName, property := range t.Properties {
		property.preProcess(spec, t, propertyName)
	}
}

func (t *Type) applyPatches(ic *SpecIssueCollector, pm *SpecPatchManager) {
	pm.applyTypePatches(ic, t)
}

func (t *Type) collectIssues(ic *SpecIssueCollector) {
	ic.MaybeCollectIssue(t, t.Documentation == "", "missing Documentation")

	if t.Name == "" {
		ic.CollectIssue(t, "missing Name")
		return
	}

	if t.IsTopLevelResourceType && strings.Contains(t.Name, ".") {
		ic.CollectIssue(t, "unexpected '.' in Name")
		return
	}

	if !t.IsTopLevelResourceType && !strings.Contains(t.Name, ".") {
		ic.CollectIssue(t, "missing '.' in Name")
		return
	}

	for _, attribute := range t.Attributes {
		attribute.collectIssues(ic)
	}

	for _, property := range t.Properties {
		property.collectIssues(ic)
	}
}
