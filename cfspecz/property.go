package cfspecz

import (
	"fmt"

	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ProblemLocation = (*Property)(nil)
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

// GetDisplayPath implements the ProblemLocation interface.
func (p *Property) GetDisplayPath() string {
	return p.displayPath
}

func (p *Property) preProcess(spec *Spec, parent *Type, name string) {
	p.Name = name

	p.MaybeLookupType = func(unqualifiedStructuredTypeName string) *Type {
		if t := spec.PropertyTypes[parent.GetRelatedStructuredTypeName(unqualifiedStructuredTypeName)]; t != nil {
			t.IsReferenced = true
			return t
		}

		return nil
	}

	p.displayPath = fmt.Sprintf("%v/property[%v]", parent.GetDisplayPath(), p.Name)
}

func (p *Property) collectProblems(pc *cfz.ProblemsCollector) {
	pc.MaybeCollect(p, p.Documentation == "", "missing Documentation")

	if p.PrimitiveType == "" && p.Type == "" {
		pc.Collect(p, "missing both PrimitiveType and Type")
		return
	}

	if p.PrimitiveType != "" {
		pc.MaybeCollect(p, p.Type != "", "unexpected Type")
		pc.MaybeCollect(p, p.PrimitiveItemType != "", "unexpected PrimitiveItemType")
		pc.MaybeCollect(p, p.ItemType != "", "unexpected ItemType")
		pc.MaybeCollect(p, p.DuplicatesAllowed, "unexpected DuplicatesAllowed")
		pc.MaybeCollect(p, !IsValidPrimitiveType(p.PrimitiveType), "invalid PrimitiveType: %v", p.PrimitiveType)

		return
	}

	if p.Type == "List" || p.Type == "Map" {
		pc.MaybeCollect(p, p.Type == "Map" && p.DuplicatesAllowed, "unexpected DuplicatesAllowed")

		if p.PrimitiveItemType == "" && p.ItemType == "" {
			pc.Collect(p, "missing both PrimitiveItemType and ItemType")
			return
		}

		if p.PrimitiveItemType != "" {
			pc.MaybeCollect(p, p.ItemType != "", "unexpected ItemType")
			pc.MaybeCollect(p, !IsValidPrimitiveType(p.PrimitiveItemType), "invalid PrimitiveItemType: %v", p.PrimitiveItemType)

			return
		}

		if p.ItemType == "List" || p.ItemType == "Map" {
			pc.Collect(p, "unsupported ItemType: '%v'", p.ItemType)
			return
		}

		pc.MaybeCollect(p,
			p.ItemType != "Tag" && p.MaybeLookupType(p.ItemType) == nil,
			"unknown ItemType: '%v'", p.ItemType)

		return
	}

	pc.MaybeCollect(p, p.PrimitiveItemType != "", "unexpected PrimitiveItemType")
	pc.MaybeCollect(p, p.ItemType != "", "unexpected ItemType")

	pc.MaybeCollect(p,
		p.Type != "Tag" && p.MaybeLookupType(p.Type) == nil,
		"unknown Type: '%v'", p.Type)
}
