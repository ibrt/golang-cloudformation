package cfschemaz

import (
	"encoding/json"
	"fmt"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/jsonz"

	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ json.Unmarshaler = (*UnprocessedEnum)(nil)
)

// UnprocessedEnum describes part of the CloudFormation JSON schema for a resource.
type UnprocessedEnum struct {
	ss []string
	nn []json.Number
}

// IsSet returns true if the enum was set in the schema.
func (ue *UnprocessedEnum) IsSet() bool {
	return ue != nil
}

// CanBeString returns true if the enum is set and can be parsed as []string.
func (ue *UnprocessedEnum) CanBeString() bool {
	return ue != nil && len(ue.ss) > 0
}

// MaybeString returns the enum as []string if possible.
func (ue *UnprocessedEnum) MaybeString() []string {
	if ue != nil && len(ue.ss) > 0 {
		return ue.ss
	}

	return nil
}

// CanBeFloat64 returns true if the enum is set and can be parsed as []float64.
func (ue *UnprocessedEnum) CanBeFloat64() bool {
	if ue == nil || len(ue.nn) == 0 {
		return false
	}

	for _, n := range ue.nn {
		if _, err := n.Float64(); err != nil {
			return false
		}
	}

	return true
}

// MaybeFloat64 returns the enum as []float64 if possible.
func (ue *UnprocessedEnum) MaybeFloat64() []float64 {
	if ue == nil || len(ue.nn) == 0 {
		return nil
	}

	ff := make([]float64, 0, len(ue.nn))

	for _, n := range ue.nn {
		f, err := n.Float64()
		if err != nil {
			return nil
		}

		ff = append(ff, f)
	}

	return ff
}

// CanBeInt64 returns true if the enum is set and can be parsed as []int64.
func (ue *UnprocessedEnum) CanBeInt64() bool {
	if ue == nil || len(ue.nn) == 0 {
		return false
	}

	for _, n := range ue.nn {
		if _, err := n.Int64(); err != nil {
			return false
		}
	}

	return true
}

// MaybeInt64 returns the enum as []int64 if possible.
func (ue *UnprocessedEnum) MaybeInt64() []int64 {
	if ue == nil || len(ue.nn) == 0 {
		return nil
	}

	ii := make([]int64, 0, len(ue.nn))

	for _, n := range ue.nn {
		i, err := n.Int64()
		if err != nil {
			return nil
		}

		ii = append(ii, i)
	}

	return ii
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ue *UnprocessedEnum) UnmarshalJSON(bytes []byte) error {
	if ss, err := jsonz.Unmarshal[[]string](bytes); err == nil {
		*ue = UnprocessedEnum{ss: ss}
		return nil
	}

	if nn, err := jsonz.Unmarshal[[]json.Number](bytes); err == nil {
		*ue = UnprocessedEnum{nn: nn}
		return nil
	}

	return errorz.Errorf("invalid enum value: '%s'", bytes)
}

// UnprocessedHandlers describes part of the CloudFormation JSON schema for a resource.
type UnprocessedHandlers struct {
	Create *UnprocessedHandlersOperation `json:"create,omitempty"`
	Delete *UnprocessedHandlersOperation `json:"delete,omitempty"`
	List   *UnprocessedHandlersOperation `json:"list,omitempty"`
	Read   *UnprocessedHandlersOperation `json:"read,omitempty"`
	Update *UnprocessedHandlersOperation `json:"update,omitempty"`
}

// UnprocessedHandlersOperation describes part of the CloudFormation JSON schema for a resource.
type UnprocessedHandlersOperation struct {
	HandlerSchema    *UnprocessedDefinition `json:"handlerSchema,omitempty"`
	Permissions      []string               `json:"permissions,omitempty"`
	TimeoutInMinutes *int                   `json:"timeoutInMinutes,omitempty"`
}

// UnprocessedRelationshipRef describes part of the CloudFormation JSON schema.
type UnprocessedRelationshipRef struct {
	TypeName     string `json:"typeName,omitempty"`
	PropertyPath string `json:"propertyPath,omitempty"`
}

// UnprocessedResourceLink describes part of the CloudFormation JSON schema.
type UnprocessedResourceLink struct {
	Mappings    map[string]string `json:"mappings,omitempty"`
	TemplateURI string            `json:"templateUri,omitempty"`
}

// UnprocessedDefinition describes part of the CloudFormation JSON schema.
type UnprocessedDefinition struct {
	AdditionalProperties bool                              `json:"additionalProperties,omitempty"`
	AllOf                any                               `json:"allOf,omitempty"`
	AnyOf                any                               `json:"anyOf,omitempty"`
	ArrayType            string                            `json:"arrayType,omitempty"`
	Comment              string                            `json:"$comment,omitempty"`
	Const                string                            `json:"const,omitempty"`
	Default              any                               `json:"default,omitempty"` // If specified, it is a value of "type".
	Dependencies         map[string][]string               `json:"dependencies,omitempty"`
	Description          string                            `json:"description,omitempty"`
	Enum                 *UnprocessedEnum                  `json:"enum,omitempty"`
	Examples             any                               `json:"examples,omitempty"`
	Format               string                            `json:"format,omitempty"`
	InsertionOrder       *bool                             `json:"insertionOrder,omitempty"`
	Items                *UnprocessedDefinition            `json:"items,omitempty"`
	MaxItems             *int                              `json:"maxItems,omitempty"`
	MaxLength            *int                              `json:"maxLength,omitempty"`
	MaxProperties        *int                              `json:"maxProperties,omitempty"`
	Maximum              *json.Number                      `json:"maximum,omitempty"`
	MinItems             *int                              `json:"minItems,omitempty"`
	MinLength            *int                              `json:"minLength,omitempty"`
	MinProperties        *int                              `json:"minProperties,omitempty"`
	Minimum              *json.Number                      `json:"minimum,omitempty"`
	MultipleOf           *json.Number                      `json:"multipleOf,omitempty"`
	OneOf                []*UnprocessedDefinition          `json:"oneOf,omitempty"`
	Pattern              *string                           `json:"pattern,omitempty"`
	PatternProperties    map[string]*UnprocessedDefinition `json:"patternProperties,omitempty"`
	Properties           map[string]*UnprocessedDefinition `json:"properties,omitempty"`
	Ref                  *string                           `json:"$ref,omitempty"`
	RelationshipRef      *UnprocessedRelationshipRef       `json:"relationshipRef,omitempty"`
	Required             []string                          `json:"required,omitempty"`
	Title                string                            `json:"title,omitempty"`
	Type                 any                               `json:"type,omitempty"` // This can be string or []string.
	UniqueItems          *bool                             `json:"uniqueItems,omitempty"`
}

/*
func (ust *UnprocessedDefinition) collectProblems(pc *cfz.ProblemsCollector, plt *cfz.ProblemLocationTracker) {

	if ust.Type == "integer" {
		if _, err := ust.Enum.Int64(); err != nil {
			pc.Collect(plt, "invalid enum: %v", err.Error())
		}
	}

	if ust.Type == "number" {
		if _, err := ust.Enum.Float64(); err != nil {
			pc.Collect(plt, "invalid enum: %v", err.Error())
		}
	}

	for name, pUST := range ust.Properties {
		pUST.collectProblems(pc, plt.WithPathElements(fmt.Sprintf("property[%v]", name)))
	}
}
*/

func (ust *UnprocessedDefinition) toType(
	pc *cfz.ProblemsCollector,
	plt *cfz.ProblemLocationTracker,
	parentULTR *UnprocessedTopLevelResource,
	name string,
) *Type {
	if ust.Type != "object" {
		return nil
	}

	t := &Type{
		IsTopLevelResourceType: false,
		Name:                   fmt.Sprintf("%v.%v", parentULTR.TypeName, name),
		Description:            ust.Description,
		Properties:             make(map[string]*Property),
	}

	for name, ud := range ust.Properties {
		t.Properties[name] = ud.toProperty(pc, plt.WithPathElements(fmt.Sprintf("property[%v]", name)), parentULTR, name)
	}

	return t
}

func (ust *UnprocessedDefinition) toProperty(
	pc *cfz.ProblemsCollector,
	plt *cfz.ProblemLocationTracker,
	parentULTR *UnprocessedTopLevelResource,
	name string,
) *Property {
	// TODO(ibrt): Find problems.

	return &Property{
		Name:       name,
		Validation: ust.toValidation(pc, plt, parentULTR),
	}
}

func (ust *UnprocessedDefinition) toValidation(
	pc *cfz.ProblemsCollector,
	plt *cfz.ProblemLocationTracker,
	parentULTR *UnprocessedTopLevelResource,
) *Validation {
	// TODO(ibrt): Find problems.

	if ust.Type == "string" {
		pc.MaybeCollect(plt, ust.MinLength == nil || *ust.MinLength >= 0, "invalid MinLength")
		pc.MaybeCollect(plt, ust.MaxLength == nil || *ust.MaxLength >= 0, "invalid MaxLength")
		pc.MaybeCollect(plt, ust.Pattern == nil || *ust.Pattern != "", "invalid Pattern")
		pc.MaybeCollect(plt, !ust.Enum.IsSet() || ust.Enum.CanBeString(), "invalid Enum")

		vs := &ValidationString{
			MinLength: ust.MinLength,
			MaxLength: ust.MaxLength,
			Pattern:   ust.Pattern,
			Enum:      ust.Enum.MaybeString(),
		}

		return &Validation{
			Single: &ValidationSingle{
				String: vs,
			},
		}
	}

	if ust.Type == "integer" {
		pc.MaybeCollect(plt, ust.Minimum == nil || canJSONNumberBeInt64(*ust.Minimum), "invalid Minimum")
		pc.MaybeCollect(plt, ust.Maximum == nil || canJSONNumberBeInt64(*ust.Maximum), "invalid Maximum")
		pc.MaybeCollect(plt, ust.MultipleOf == nil || canJSONNumberBeInt64(*ust.MultipleOf) || mustJSONNumberInt64(*ust.MultipleOf) <= 0, "invalid MultipleOf")
		pc.MaybeCollect(plt, !ust.Enum.IsSet() || ust.Enum.CanBeInt64(), "invalid Enum")

		vi := &ValidationNumber[int64]{
			Minimum:    maybeJSONNumberPtrInt64(ust.Minimum),
			Maximum:    maybeJSONNumberPtrInt64(ust.Maximum),
			MultipleOf: maybeJSONNumberPtrInt64(ust.MultipleOf),
			Enum:       ust.Enum.MaybeInt64(),
		}

		return &Validation{
			Single: &ValidationSingle{
				Int64: vi,
			},
		}
	}

	if ust.Type == "number" {
		pc.MaybeCollect(plt, ust.Minimum == nil || canJSONNumberBeFloat64(*ust.Minimum), "invalid Minimum")
		pc.MaybeCollect(plt, ust.Maximum == nil || canJSONNumberBeFloat64(*ust.Maximum), "invalid Maximum")
		pc.MaybeCollect(plt, ust.MultipleOf == nil || canJSONNumberBeFloat64(*ust.MultipleOf) || mustJSONNumberFloat64(*ust.MultipleOf) <= 0, "invalid MultipleOf")
		pc.MaybeCollect(plt, !ust.Enum.IsSet() || ust.Enum.CanBeFloat64(), "invalid Enum")

		vf := &ValidationNumber[float64]{
			Minimum:    maybeJSONNumberPtrFloat64(ust.Minimum),
			Maximum:    maybeJSONNumberPtrFloat64(ust.Maximum),
			MultipleOf: maybeJSONNumberPtrFloat64(ust.MultipleOf),
			Enum:       ust.Enum.MaybeFloat64(),
		}

		return &Validation{
			Single: &ValidationSingle{
				Float64: vf,
			},
		}
	}

	return nil
}

// UnprocessedTagging describes part of the CloudFormation JSON schema for a resource.
type UnprocessedTagging struct {
	CloudFormationSystemTags bool     `json:"cloudFormationSystemTags,omitempty"`
	Permissions              []string `json:"permissions,omitempty"`
	TagOnCreate              bool     `json:"tagOnCreate,omitempty"`
	TagProperty              string   `json:"tagProperty,omitempty"`
	TagUpdatable             bool     `json:"tagUpdatable,omitempty"`
	Taggable                 bool     `json:"taggable,omitempty"`
}

// UnprocessedTopLevelResource describes part of the CloudFormation JSON schema.
type UnprocessedTopLevelResource struct {
	AdditionalIdentifiers          [][]string                        `json:"additionalIdentifiers,omitempty"`
	AdditionalProperties           bool                              `json:"additionalProperties,omitempty"`
	AllOf                          any                               `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                          any                               `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	Comment                        string                            `json:"$comment,omitempty"`
	ConditionalCreatOnlyProperties []string                          `json:"conditionalCreateOnlyProperties,omitempty"`
	CreateOnlyProperties           []string                          `json:"createOnlyProperties,omitempty"`
	Definitions                    map[string]*UnprocessedDefinition `json:"definitions,omitempty"`
	DeprecatedProperties           []string                          `json:"deprecatedProperties,omitempty"`
	Description                    string                            `json:"description,omitempty"`
	DocumentationURL               string                            `json:"documentationUrl,omitempty"`
	Handlers                       *UnprocessedHandlers              `json:"handlers,omitempty"`
	OneOf                          []*UnprocessedDefinition          `json:"oneOf,omitempty"`
	PrimaryIdentifier              []string                          `json:"primaryIdentifier,omitempty"`
	Properties                     map[string]*UnprocessedDefinition `json:"properties,omitempty"`
	PropertyTransform              map[string]string                 `json:"propertyTransform,omitempty"`
	ReadOnlyProperties             []string                          `json:"readOnlyProperties,omitempty"`
	ReplacementStrategy            string                            `json:"replacementStrategy,omitempty"`
	Required                       []string                          `json:"required,omitempty"`
	ResourceLink                   *UnprocessedResourceLink          `json:"resourceLink,omitempty"`
	Schema                         string                            `json:"$schema,omitempty"`
	SourceURL                      string                            `json:"sourceUrl,omitempty"`
	Taggable                       bool                              `json:"taggable,omitempty"`
	Tagging                        *UnprocessedTagging               `json:"tagging,omitempty"`
	TypeName                       string                            `json:"typeName,omitempty"`
	WriteOnlyProperties            []string                          `json:"writeOnlyProperties,omitempty"`
}

/*
	func (utlr *UnprocessedTopLevelResource) collectProblems(pc *cfz.ProblemsCollector, plt *cfz.ProblemLocationTracker) {
		pc.MaybeCollect(plt, utlr.TypeName == "", "missing TypeName")
		pc.MaybeCollect(plt, utlr.Description == "", "missing Description")

		for name, dUST := range utlr.Definitions {
			dUST.collectProblems(pc, plt.WithPathElements(fmt.Sprintf("definition[%v]", name)))
		}

		for name, pUST := range utlr.Properties {
			pUST.collectProblems(pc, plt.WithPathElements(fmt.Sprintf("property[%v]", name)))
		}
	}
*/
func (utlr *UnprocessedTopLevelResource) toTypes(plt *cfz.ProblemLocationTracker, s *Schema) (*Type, []*Type, error) {
	pc := cfz.NewProblemsCollector()
	pc.MaybeCollect(plt, utlr.TypeName == "", "missing TypeName")
	pc.MaybeCollect(plt, utlr.Description == "", "missing Description")

	tlr := &Type{
		IsTopLevelResourceType: true,
		Name:                   utlr.TypeName,
		Description:            utlr.Description,
		Properties:             make(map[string]*Property),
	}

	sts := make([]*Type, 0, len(utlr.Definitions))

	for name, ust := range utlr.Definitions {
		if t := ust.toType(pc, plt, utlr, name); t != nil {
			sts = append(sts, t)
		}
	}

	return tlr, sts, pc.ToError()
}

func canJSONNumberBeFloat64(n json.Number) bool {
	_, err := n.Float64()
	return err == nil
}

func mustJSONNumberFloat64(n json.Number) float64 {
	f, err := n.Float64()
	errorz.MaybeMustWrap(err)
	return f
}

func maybeJSONNumberPtrFloat64(n *json.Number) *float64 {
	if n != nil {
		if f, err := n.Float64(); err == nil {
			return &f
		}
	}
	return nil
}

func canJSONNumberBeInt64(n json.Number) bool {
	_, err := n.Int64()
	return err == nil
}

func mustJSONNumberInt64(n json.Number) int64 {
	i, err := n.Int64()
	errorz.MaybeMustWrap(err)
	return i
}

func maybeJSONNumberPtrInt64(n *json.Number) *int64 {
	if n != nil {
		if i, err := n.Int64(); err == nil {
			return &i
		}
	}
	return nil
}
