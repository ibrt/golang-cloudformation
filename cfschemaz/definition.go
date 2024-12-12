package cfschemaz

import (
	"encoding/json"
)

// Definition describes part of the CloudFormation JSON schema.
type Definition struct {
	AdditionalProperties bool                       `json:"additionalProperties,omitempty"`
	AllOf                any                        `json:"allOf,omitempty"`
	AnyOf                any                        `json:"anyOf,omitempty"`
	ArrayType            string                     `json:"arrayType,omitempty"`
	Comment              string                     `json:"$comment,omitempty"`
	Const                string                     `json:"const,omitempty"`
	Default              any                        `json:"default,omitempty"` // If specified, it is a value of "type".
	Dependencies         map[string][]string        `json:"dependencies,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Enum                 *Enum                      `json:"enum,omitempty"`
	Examples             any                        `json:"examples,omitempty"`
	Format               string                     `json:"format,omitempty"`
	InsertionOrder       *bool                      `json:"insertionOrder,omitempty"`
	Items                *Definition                `json:"items,omitempty"`
	MaxItems             *int                       `json:"maxItems,omitempty"`
	MaxLength            *int                       `json:"maxLength,omitempty"`
	MaxProperties        *int                       `json:"maxProperties,omitempty"`
	Maximum              *json.Number               `json:"maximum,omitempty"`
	MinItems             *int                       `json:"minItems,omitempty"`
	MinLength            *int                       `json:"minLength,omitempty"`
	MinProperties        *int                       `json:"minProperties,omitempty"`
	Minimum              *json.Number               `json:"minimum,omitempty"`
	MultipleOf           *json.Number               `json:"multipleOf,omitempty"`
	OneOf                []*Definition              `json:"oneOf,omitempty"`
	Pattern              *string                    `json:"pattern,omitempty"`
	PatternProperties    map[string]*Definition     `json:"patternProperties,omitempty"`
	Properties           map[string]*Definition     `json:"properties,omitempty"`
	Ref                  *string                    `json:"$ref,omitempty"`
	RelationshipRef      *DefinitionRelationshipRef `json:"relationshipRef,omitempty"`
	Required             []string                   `json:"required,omitempty"`
	Title                string                     `json:"title,omitempty"`
	Type                 any                        `json:"type,omitempty"` // This can be string or []string.
	UniqueItems          *bool                      `json:"uniqueItems,omitempty"`
}

// DefinitionRelationshipRef describes part of the CloudFormation JSON schema.
type DefinitionRelationshipRef struct {
	TypeName     string `json:"typeName,omitempty"`
	PropertyPath string `json:"propertyPath,omitempty"`
}
