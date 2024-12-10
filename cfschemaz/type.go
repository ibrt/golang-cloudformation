package cfschemaz

// Type describes either a top-level resource type or a structured type.
type Type struct {
	IsTopLevelResourceType bool
	Name                   string
	Description            string
}
