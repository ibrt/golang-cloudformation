package specz

import (
	"fmt"
)

// Attribute describes an attribute of a top-level resource type.
type Attribute struct {
	ItemType          string
	PrimitiveItemType string
	PrimitiveType     string
	Type              string

	Parent     *Type                       `json:"-"`
	Name       string                      `json:"-"`
	LookupType func(typeName string) *Type `json:"-"`
}

func (a *Attribute) getDisplayPath() string {
	return fmt.Sprintf("%v/attribute[%v]", a.Parent.getDisplayPath(), a.Name)
}

func (a *Attribute) preProcess(spec *Spec, parent *Type, name string) {
	a.Parent = parent
	a.Name = name

	a.LookupType = func(unqualifiedStructuredTypeName string) *Type {
		return spec.PropertyTypes[parent.getRelatedStructuredTypeName(unqualifiedStructuredTypeName)]
	}
}

func (a *Attribute) validateAndCollectIssues(ic *issuesCollector) {
	if a.PrimitiveType == "" && a.Type == "" {
		ic.collectIssue(a, "missing both PrimitiveType and Type")
		return
	}

	if a.PrimitiveType != "" {
		if a.Type != "" || a.PrimitiveItemType != "" || a.ItemType != "" {
			ic.collectIssue(a, "unexpected Type, PrimitiveItemType, ItemType")
		}

		if !isValidPrimitiveType(a.PrimitiveType) {
			ic.collectIssue(a, "invalid PrimitiveType: %v", a.PrimitiveType)
		}

		if a.PrimitiveType == "Json" {
			ic.collectIssue(a, "unsupported PrimitiveType: 'Json'")
		}

		return
	}

	if a.Type == "List" {
		if a.PrimitiveItemType == "" && a.ItemType == "" {
			ic.collectIssue(a, "missing both PrimitiveItemType and ItemType")
			return
		}

		if a.PrimitiveItemType != "" {
			if a.ItemType != "" {
				ic.collectIssue(a, "unexpected ItemType")
			}

			if !isValidPrimitiveType(a.PrimitiveItemType) {
				ic.collectIssue(a, "invalid PrimitiveItemType: %v", a.PrimitiveItemType)
			}

			if a.PrimitiveItemType == "Json" {
				ic.collectIssue(a, "unsupported PrimitiveItemType: 'Json'")
			}

			return
		}

		if a.ItemType == "List" || a.ItemType == "Map" {
			ic.collectIssue(a, "unsupported ItemType: '%v'", a.ItemType)
			return
		}

		if a.ItemType != "Tag" && a.LookupType(a.ItemType) == nil {
			ic.collectIssue(a, "unknown ItemType: '%v'", a.ItemType)
		}

		return
	}

	if a.Type == "Map" {
		ic.collectIssue(a, "unsupported Type: 'Map'")
		return
	}

	if a.PrimitiveItemType != "" || a.ItemType != "" {
		ic.collectIssue(a, "unexpected PrimitiveItemType or ItemType")
	}

	if a.Type != "Tag" && a.LookupType(a.Type) == nil {
		ic.collectIssue(a, "unknown Type: '%v'", a.Type)
	}
}
