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

// GetDisplayPath implements the ProblemLocation interface.
func (a *Attribute) GetDisplayPath() string {
	return a.displayPath
}

func (a *Attribute) preProcess(spec *Spec, parent *Type, name string) {
	a.Name = name

	a.MaybeLookupType = func(unqualifiedStructuredTypeName string) *Type {
		if t := spec.PropertyTypes[parent.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]; t != nil {
			t.IsReferenced = true
			return t
		}

		return nil
	}

	a.displayPath = fmt.Sprintf("%v/attribute[%v]", parent.GetDisplayPath(), a.Name)
}

func (a *Attribute) collectProblems(pc *ProblemsCollector) {
	if a.PrimitiveType == "" && a.Type == "" {
		pc.Collect(a, "missing both PrimitiveType and Type")
		return
	}

	if a.PrimitiveType != "" {
		pc.MaybeCollect(a, a.Type != "", "unexpected Type")
		pc.MaybeCollect(a, a.PrimitiveItemType != "", "unexpected PrimitiveItemType")
		pc.MaybeCollect(a, a.ItemType != "", "unexpected ItemType")
		pc.MaybeCollect(a, !IsValidPrimitiveType(a.PrimitiveType), "invalid PrimitiveType: %v", a.PrimitiveType)
		pc.MaybeCollect(a, a.PrimitiveType == "Json", "unsupported PrimitiveType: 'Json'")

		return
	}

	if a.Type == "List" {
		if a.PrimitiveItemType == "" && a.ItemType == "" {
			pc.Collect(a, "missing both PrimitiveItemType and ItemType")
			return
		}

		if a.PrimitiveItemType != "" {
			pc.MaybeCollect(a, a.ItemType != "", "unexpected ItemType")
			pc.MaybeCollect(a, !IsValidPrimitiveType(a.PrimitiveItemType), "invalid PrimitiveItemType: %v", a.PrimitiveItemType)
			pc.MaybeCollect(a, a.PrimitiveItemType == "Json", "unsupported PrimitiveItemType: 'Json'")

			return
		}

		if a.ItemType == "List" || a.ItemType == "Map" {
			pc.Collect(a, "unsupported ItemType: '%v'", a.ItemType)
			return
		}

		pc.MaybeCollect(a,
			a.ItemType != "Tag" && a.MaybeLookupType(a.ItemType) == nil,
			"unknown ItemType: '%v'", a.ItemType)

		return
	}

	if a.Type == "Map" {
		pc.Collect(a, "unsupported Type: 'Map'")
		return
	}

	pc.MaybeCollect(a, a.PrimitiveItemType != "", "unexpected PrimitiveItemType")
	pc.MaybeCollect(a, a.ItemType != "", "unexpected ItemType")

	pc.MaybeCollect(a,
		a.Type != "Tag" && a.MaybeLookupType(a.Type) == nil,
		"unknown Type: '%v'", a.Type)
}
