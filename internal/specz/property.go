package specz

import (
	"fmt"
)

// Property describes a property of either a top-level resource or structured type.
type Property struct {
	Documentation     string
	DuplicatesAllowed bool
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Required          bool
	Type              string
	UpdateType        string

	Parent     *Type                       `json:"-"`
	Name       string                      `json:"-"`
	LookupType func(typeName string) *Type `json:"-"`
}

func (p *Property) getDisplayPath() string {
	return fmt.Sprintf("%v/property[%v]", p.Parent.getDisplayPath(), p.Name)
}

func (p *Property) preProcess(spec *Spec, parent *Type, name string) {
	p.Parent = parent
	p.Name = name

	p.LookupType = func(unqualifiedStructuredTypeName string) *Type {
		return spec.PropertyTypes[parent.getRelatedStructuredTypeName(unqualifiedStructuredTypeName)]
	}
}

func (p *Property) validateAndCollectIssues(ic *issuesCollector) {
	if p.Documentation == "" {
		ic.collectIssue(p, "missing Documentation")
	}

	if p.PrimitiveType == "" && p.Type == "" {
		ic.collectIssue(p, "missing both PrimitiveType and Type")
		return
	}

	if p.PrimitiveType != "" {
		if p.Type != "" || p.PrimitiveItemType != "" || p.ItemType != "" {
			ic.collectIssue(p, "unexpected Type, PrimitiveItemType, ItemType")
		}

		if p.DuplicatesAllowed {
			ic.collectIssue(p, "unexpected DuplicatesAllowed")
		}

		if !isValidPrimitiveType(p.PrimitiveType) {
			ic.collectIssue(p, "invalid PrimitiveType: %v", p.PrimitiveType)
		}

		return
	}

	if p.Type == "List" || p.Type == "Map" {
		if p.Type == "Map" && p.DuplicatesAllowed {
			ic.collectIssue(p, "unexpected DuplicatesAllowed")
		}

		if p.PrimitiveItemType == "" && p.ItemType == "" {
			ic.collectIssue(p, "missing both PrimitiveItemType and ItemType")
			return
		}

		if p.PrimitiveItemType != "" {
			if p.ItemType != "" {
				ic.collectIssue(p, "unexpected ItemType")
			}

			if !isValidPrimitiveType(p.PrimitiveItemType) {
				ic.collectIssue(p, "invalid PrimitiveItemType: %v", p.PrimitiveItemType)
			}

			return
		}

		if p.ItemType == "List" || p.ItemType == "Map" {
			ic.collectIssue(p, "unsupported ItemType: '%v'", p.ItemType)
			return
		}

		if p.ItemType != "Tag" && p.LookupType(p.ItemType) == nil {
			ic.collectIssue(p, "unknown ItemType: '%v'", p.ItemType)
		}

		return
	}

	if p.PrimitiveItemType != "" || p.ItemType != "" {
		ic.collectIssue(p, "unexpected PrimitiveItemType or ItemType")
	}

	if p.Type != "Tag" && p.LookupType(p.Type) == nil {
		ic.collectIssue(p, "unknown Type: '%v'", p.Type)
	}
}
