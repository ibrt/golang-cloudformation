package cfspecz

import (
	"fmt"
)

// Attribute describes an attribute of a top-level resource type.
type Attribute struct {
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Type              string

	Name            string                                           `json:"-"`
	MaybeLookupType func(unqualifiedStructuredTypeName string) *Type `json:"-"`

	displayPath string
}

// GetDisplayPath implements the SpecContext interface.
func (a *Attribute) GetDisplayPath() string {
	return a.displayPath
}

func (a *Attribute) preProcess(spec *Spec, parent *Type, name string) {
	a.Name = name

	a.MaybeLookupType = func(unqualifiedStructuredTypeName string) *Type {
		return spec.PropertyTypes[parent.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]
	}

	a.displayPath = fmt.Sprintf("%v/attribute[%v]", parent.GetDisplayPath(), a.Name)
}

func (a *Attribute) collectIssues(ic *SpecIssueCollector) {
	if a.PrimitiveType == "" && a.Type == "" {
		ic.CollectIssue(a, "missing both PrimitiveType and Type")
		return
	}

	if a.PrimitiveType != "" {
		ic.MaybeCollectIssue(a, a.Type != "", "unexpected Type")
		ic.MaybeCollectIssue(a, a.PrimitiveItemType != "", "unexpected PrimitiveItemType")
		ic.MaybeCollectIssue(a, a.ItemType != "", "unexpected ItemType")
		ic.MaybeCollectIssue(a, !IsValidPrimitiveType(a.PrimitiveType), "invalid PrimitiveType: %v", a.PrimitiveType)
		ic.MaybeCollectIssue(a, a.PrimitiveType == "Json", "unsupported PrimitiveType: 'Json'")

		return
	}

	if a.Type == "List" {
		if a.PrimitiveItemType == "" && a.ItemType == "" {
			ic.CollectIssue(a, "missing both PrimitiveItemType and ItemType")
			return
		}

		if a.PrimitiveItemType != "" {
			ic.MaybeCollectIssue(a, a.ItemType != "", "unexpected ItemType")
			ic.MaybeCollectIssue(a, !IsValidPrimitiveType(a.PrimitiveItemType), "invalid PrimitiveItemType: %v", a.PrimitiveItemType)
			ic.MaybeCollectIssue(a, a.PrimitiveItemType == "Json", "unsupported PrimitiveItemType: 'Json'")

			return
		}

		if a.ItemType == "List" || a.ItemType == "Map" {
			ic.CollectIssue(a, "unsupported ItemType: '%v'", a.ItemType)
			return
		}

		ic.MaybeCollectIssue(a,
			a.ItemType != "Tag" && a.MaybeLookupType(a.ItemType) == nil,
			"unknown ItemType: '%v'", a.ItemType)

		return
	}

	if a.Type == "Map" {
		ic.CollectIssue(a, "unsupported Type: 'Map'")
		return
	}

	ic.MaybeCollectIssue(a, a.PrimitiveItemType != "", "unexpected PrimitiveItemType")
	ic.MaybeCollectIssue(a, a.ItemType != "", "unexpected ItemType")

	ic.MaybeCollectIssue(a,
		a.Type != "Tag" && a.MaybeLookupType(a.Type) == nil,
		"unknown Type: '%v'", a.Type)
}
