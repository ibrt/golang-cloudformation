package cfschemaz

// Type describes either a top-level resource type or a structured type in the schema.
type Type struct {
	IsTopLevelResourceType bool `json:"-"`
	Name                   string
	Description            string
}