package cfschemaz

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ibrt/golang-utils/errorz"

	"github.com/ibrt/golang-cloudformation/cfz"
)

// Definition describes part of the CloudFormation JSON schema.
type Definition struct {
	AdditionalProperties bool                       `json:"additionalProperties,omitempty"`
	AllOf                []*Definition              `json:"allOf,omitempty"`
	AnyOf                []*Definition              `json:"anyOf,omitempty"`
	ArrayType            string                     `json:"arrayType,omitempty"`
	Comment              string                     `json:"$comment,omitempty"`
	Const                string                     `json:"const,omitempty"`
	Default              any                        `json:"default,omitempty"` // If specified, it is a value of "type".
	Dependencies         map[string][]string        `json:"dependencies,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Enum                 *Enum                      `json:"enum,omitempty"`
	Examples             []json.RawMessage          `json:"examples,omitempty"`
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
	Type                 any                        `json:"type,omitempty"` // It can be string or []string.
	UniqueItems          *bool                      `json:"uniqueItems,omitempty"`
}

func (d *Definition) maybeGetValidation(pc *cfz.ProblemsCollector, plt *cfz.ProblemLocationTracker, r *Resource) *cfz.Validation {
	if d == nil {
		return nil
	}

	ref, err := d.maybeGetRef()
	if err != nil {
		pc.Collect(plt, "invalid $ref: %v", err.Error())
		return nil
	}

	if ref != "" {
		refD := r.Definitions[ref]

		if refD == nil {
			pc.Collect(plt, "unknown $ref: %v", ref)
			return nil
		}

		if refD.Type == "object" {
			return &cfz.Validation{
				Single: &cfz.ValidationSingle{
					Structured: &cfz.ValidationStructuredType{
						StructuredTypeName: fmt.Sprintf("%v.%v", r.TypeName, ref),
					},
				},
			}
		}

		return refD.maybeGetValidation(pc, plt.WithPathElements(fmt.Sprintf("reference[%v]", ref)), r)
	}

	if d.Type == "string" {
		minLength := d.MinLength
		maxLength := d.MaxLength
		pattern := d.Pattern

		if minLength != nil && *minLength < 0 {
			pc.Collect(plt, "invalid MinLength: %v", *minLength)
			minLength = nil
		}

		if maxLength != nil && *maxLength < 0 {
			pc.Collect(plt, "invalid MaxLength: %v", *maxLength)
			maxLength = nil
		}

		if pattern != nil && *pattern == "" {
			pc.Collect(plt, "invalid Pattern: '%v'", *pattern)
			pattern = nil
		}

		pc.MaybeCollect(plt, d.Enum.IsSet() && !d.Enum.CanBeString(), "invalid Enum")

		return &cfz.Validation{
			Single: &cfz.ValidationSingle{
				String: &cfz.ValidationString{
					MinLength: minLength,
					MaxLength: maxLength,
					Pattern:   pattern,
					Enum:      d.Enum.MaybeString(),
				},
			},
		}
	}

	if d.Type == "integer" {
		minimum := d.Minimum
		maximum := d.Maximum
		multipleOf := d.MultipleOf

		if minimum != nil && !canJSONNumberBeInt64(*minimum) {
			pc.Collect(plt, "invalid Minimum: %v", *minimum)
			minimum = nil
		}

		if maximum != nil && !canJSONNumberBeInt64(*maximum) {
			pc.Collect(plt, "invalid Maximum: %v", *maximum)
			maximum = nil
		}

		if multipleOf != nil && !canJSONNumberBeInt64(*multipleOf) {
			pc.Collect(plt, "invalid MultipleOf: %v", *multipleOf)
			multipleOf = nil
		}

		pc.MaybeCollect(plt, d.Enum.IsSet() && !d.Enum.CanBeInt64(), "invalid Enum")

		return &cfz.Validation{
			Single: &cfz.ValidationSingle{
				Int64: &cfz.ValidationNumber[int64]{
					Minimum:    maybeJSONNumberPtrInt64(minimum),
					Maximum:    maybeJSONNumberPtrInt64(maximum),
					MultipleOf: maybeJSONNumberPtrInt64(multipleOf),
					Enum:       d.Enum.MaybeInt64(),
				},
			},
		}
	}

	if d.Type == "number" {
		minimum := d.Minimum
		maximum := d.Maximum
		multipleOf := d.MultipleOf

		if minimum != nil && !canJSONNumberBeFloat64(*minimum) {
			pc.Collect(plt, "invalid Minimum: %v", *minimum)
			minimum = nil
		}

		if maximum != nil && !canJSONNumberBeFloat64(*maximum) {
			pc.Collect(plt, "invalid Maximum: %v", *maximum)
			maximum = nil
		}

		if multipleOf != nil && !canJSONNumberBeFloat64(*multipleOf) {
			pc.Collect(plt, "invalid MultipleOf: %v", *multipleOf)
			multipleOf = nil
		}
		pc.MaybeCollect(plt, d.Enum.IsSet() && !d.Enum.CanBeFloat64(), "invalid Enum")

		return &cfz.Validation{
			Single: &cfz.ValidationSingle{
				Float64: &cfz.ValidationNumber[float64]{
					Minimum:    maybeJSONNumberPtrFloat64(minimum),
					Maximum:    maybeJSONNumberPtrFloat64(maximum),
					MultipleOf: maybeJSONNumberPtrFloat64(multipleOf),
					Enum:       d.Enum.MaybeFloat64(),
				},
			},
		}
	}

	if d.Type == "array" {
		minItems := d.MinItems
		maxItems := d.MaxItems

		if minItems != nil && *minItems < 0 {
			pc.Collect(plt, "invalid MinItems: %v", *minItems)
			minItems = nil
		}

		if maxItems != nil && *maxItems < 0 {
			pc.Collect(plt, "invalid MaxItems: %v", *maxItems)
			maxItems = nil
		}

		return &cfz.Validation{
			Array: &cfz.ValidationArray{
				MinItems:       minItems,
				MaxItems:       maxItems,
				InsertionOrder: d.InsertionOrder,
				UniqueItems:    d.UniqueItems,
				Items:          d.Items.maybeGetValidation(pc, plt, r),
			},
		}
	}

	return nil
}

func (d *Definition) maybeGetRef() (string, error) {
	if d.Ref == nil {
		return "", nil
	}

	ref := *d.Ref

	if !strings.HasPrefix(ref, "#/definitions/") {
		return "", errorz.Errorf("unexpected $ref prefix: '%v'", ref)
	}

	ref = strings.TrimPrefix(ref, "#/definitions/")

	if ref == "" {
		return "", errorz.Errorf("empty $ref")
	}

	return ref, nil
}

// DefinitionRelationshipRef describes part of the CloudFormation JSON schema.
type DefinitionRelationshipRef struct {
	TypeName     string `json:"typeName,omitempty"`
	PropertyPath string `json:"propertyPath,omitempty"`
}
