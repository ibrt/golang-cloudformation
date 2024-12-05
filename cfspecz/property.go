package cfspecz

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

	Name            string                                           `json:"-"`
	MaybeLookupType func(unqualifiedStructuredTypeName string) *Type `json:"-"`
	displayPath     string
}

// GetDisplayPath implements the SpecContext interface.
func (p *Property) GetDisplayPath() string {
	return p.displayPath
}

func (p *Property) preProcess(spec *Spec, parent *Type, name string) {
	p.Name = name

	p.MaybeLookupType = func(unqualifiedStructuredTypeName string) *Type {
		return spec.PropertyTypes[parent.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]
	}

	p.displayPath = fmt.Sprintf("%v/property[%v]", parent.GetDisplayPath(), p.Name)
}

func (p *Property) collectIssues(ic *SpecIssueCollector) {
	ic.MaybeCollectIssue(p, p.Documentation == "", "missing Documentation")

	if p.PrimitiveType == "" && p.Type == "" {
		ic.CollectIssue(p, "missing both PrimitiveType and Type")
		return
	}

	if p.PrimitiveType != "" {
		ic.MaybeCollectIssue(p, p.Type != "", "unexpected Type")
		ic.MaybeCollectIssue(p, p.PrimitiveItemType != "", "unexpected PrimitiveItemType")
		ic.MaybeCollectIssue(p, p.ItemType != "", "unexpected ItemType")
		ic.MaybeCollectIssue(p, p.DuplicatesAllowed, "unexpected DuplicatesAllowed")
		ic.MaybeCollectIssue(p, !IsValidPrimitiveType(p.PrimitiveType), "invalid PrimitiveType: %v", p.PrimitiveType)

		return
	}

	if p.Type == "List" || p.Type == "Map" {
		ic.MaybeCollectIssue(p, p.Type == "Map" && p.DuplicatesAllowed, "unexpected DuplicatesAllowed")

		if p.PrimitiveItemType == "" && p.ItemType == "" {
			ic.CollectIssue(p, "missing both PrimitiveItemType and ItemType")
			return
		}

		if p.PrimitiveItemType != "" {
			ic.MaybeCollectIssue(p, p.ItemType != "", "unexpected ItemType")
			ic.MaybeCollectIssue(p, !IsValidPrimitiveType(p.PrimitiveItemType), "invalid PrimitiveItemType: %v", p.PrimitiveItemType)

			return
		}

		if p.ItemType == "List" || p.ItemType == "Map" {
			ic.CollectIssue(p, "unsupported ItemType: '%v'", p.ItemType)
			return
		}

		ic.MaybeCollectIssue(p,
			p.ItemType != "Tag" && p.MaybeLookupType(p.ItemType) == nil,
			"unknown ItemType: '%v'", p.ItemType)

		return
	}

	ic.MaybeCollectIssue(p, p.PrimitiveItemType != "", "unexpected PrimitiveItemType")
	ic.MaybeCollectIssue(p, p.ItemType != "", "unexpected ItemType")

	ic.MaybeCollectIssue(p,
		p.Type != "Tag" && p.MaybeLookupType(p.Type) == nil,
		"unknown Type: '%v'", p.Type)
}
