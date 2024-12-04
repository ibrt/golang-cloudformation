package specz

import (
	"fmt"
	"strings"
)

// Type describes either a top-level resource type or a structured type in the spec.
type Type struct {
	Documentation string
	Attributes    map[string]*Attribute
	Properties    map[string]*Property

	IsTopLevelResourceType bool   `json:"-"`
	Name                   string `json:"-"`
}

func (t *Type) getRelatedTopLevelResourceTypeName() string {
	if t.IsTopLevelResourceType {
		return t.Name
	}
	return strings.Split(t.Name, ".")[0]
}

func (t *Type) getRelatedStructuredTypeName(unqualifiedStructuredTypeName string) string {
	return fmt.Sprintf("%s.%s", t.getRelatedTopLevelResourceTypeName(), unqualifiedStructuredTypeName)
}

func (t *Type) getDisplayPath() string {
	if t.IsTopLevelResourceType {
		return fmt.Sprintf("spec/topLevelResourceType[%v]", t.Name)
	}

	return fmt.Sprintf("spec/structuredType[%v]", t.Name)
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

func (t *Type) validateAndCollectIssues(ic *issuesCollector) {
	if t.Documentation == "" {
		ic.collectIssue(t, "missing Documentation")
	}

	if t.Name == "" {
		ic.collectIssue(t, "missing Name")
		return
	}

	if t.IsTopLevelResourceType && strings.Contains(t.Name, ".") {
		ic.collectIssue(t, "unexpected '.' in Name")
		return
	}

	if !t.IsTopLevelResourceType && !strings.Contains(t.Name, ".") {
		ic.collectIssue(t, "missing '.' in Name")
		return
	}

	for _, attribute := range t.Attributes {
		attribute.validateAndCollectIssues(ic)
	}

	for _, property := range t.Properties {
		property.validateAndCollectIssues(ic)
	}
}
