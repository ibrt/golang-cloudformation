package cfschemaz

// Property describes a property.
type Property struct {
	AdditionalProperties bool                 `json:"additionalProperties,omitempty"`
	AllOf                any                  `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                any                  `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	ArrayType            string               `json:"arrayType,omitempty"`
	Comment              string               `json:"$comment,omitempty"`
	Const                string               `json:"const,omitempty"`
	Default              any                  `json:"default,omitempty"` // TODO(ibrt): Type.
	Dependencies         map[string][]string  `json:"dependencies,omitempty"`
	Description          string               `json:"description,omitempty"`
	Enum                 any                  `json:"enum,omitempty"`     // TODO(ibrt): Type.
	Examples             any                  `json:"examples,omitempty"` // TODO(ibrt): Type.
	Format               string               `json:"format,omitempty"`
	InsertionOrder       *bool                `json:"insertionOrder,omitempty"`
	Items                *Property            `json:"items,omitempty"`
	MaxItems             *int                 `json:"maxItems,omitempty"`
	MaxLength            *int                 `json:"maxLength,omitempty"`
	MaxProperties        *int                 `json:"maxProperties,omitempty"`
	Maximum              *float64             `json:"maximum,omitempty"`
	MinItems             *int                 `json:"minItems,omitempty"`
	MinLength            *int                 `json:"minLength,omitempty"`
	MinProperties        *int                 `json:"minProperties,omitempty"`
	Minimum              *float64             `json:"minimum,omitempty"`
	MultipleOf           *float64             `json:"multipleOf,omitempty"`
	OneOf                []*Property          `json:"oneOf,omitempty"`
	Pattern              *string              `json:"pattern,omitempty"`
	PatternProperties    map[string]*Property `json:"patternProperties,omitempty"`
	Properties           map[string]*Property `json:"properties,omitempty"`
	Ref                  *string              `json:"$ref,omitempty"`
	RelationshipRef      *RelationshipRef     `json:"relationshipRef,omitempty"`
	Required             []string             `json:"required,omitempty"`
	Title                string               `json:"title,omitempty"`
	Type                 any                  `json:"type,omitempty"` // TODO(ibrt): Type, so far string or []string.
	UniqueItems          *bool                `json:"uniqueItems,omitempty"`
}

// RelationshipRef describes part of a property.
type RelationshipRef struct {
	TypeName     string `json:"typeName,omitempty"`
	PropertyPath string `json:"propertyPath,omitempty"`
}
