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

// AsString returns the enum as []string if possible.
func (ue *UnprocessedEnum) AsString() []string {
	if ue != nil {
		return ue.ss
	}
	return nil
}

// AsFloat64 returns the enum as []float64 if possible.
func (ue *UnprocessedEnum) AsFloat64() ([]float64, error) {
	if ue == nil || ue.nn == nil {
		return nil, nil
	}

	ff := make([]float64, len(ue.nn))

	for _, n := range ue.nn {
		f, err := n.Float64()
		if err != nil {
			return nil, errorz.Wrap(err)
		}
		ff = append(ff, f)
	}

	return ff, nil
}

// MustAsFloat64 returns the enum as []float64 if possible, panics on error.
func (ue *UnprocessedEnum) MustAsFloat64() []float64 {
	ff, err := ue.AsFloat64()
	errorz.MaybeMustWrap(err)
	return ff
}

// AsInt64 returns the enum as []int64 if possible.
func (ue *UnprocessedEnum) AsInt64() ([]int64, error) {
	if ue == nil || ue.nn == nil {
		return nil, nil
	}

	ii := make([]int64, len(ue.nn))

	for _, n := range ue.nn {
		i, err := n.Int64()
		if err != nil {
			return nil, errorz.Wrap(err)
		}
		ii = append(ii, i)
	}

	return ii, nil
}

// MustAsInt64 returns the enum as []int64 if possible, panics on error.
func (ue *UnprocessedEnum) MustAsInt64() []int64 {
	ii, err := ue.AsInt64()
	errorz.MaybeMustWrap(err)
	return ii
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
	HandlerSchema    *UnprocessedStructuredType `json:"handlerSchema,omitempty"`
	Permissions      []string                   `json:"permissions,omitempty"`
	TimeoutInMinutes *int                       `json:"timeoutInMinutes,omitempty"`
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

// UnprocessedStructuredType describes part of the CloudFormation JSON schema.
type UnprocessedStructuredType struct {
	AdditionalProperties bool                                  `json:"additionalProperties,omitempty"`
	AllOf                any                                   `json:"allOf,omitempty"`
	AnyOf                any                                   `json:"anyOf,omitempty"`
	ArrayType            string                                `json:"arrayType,omitempty"`
	Comment              string                                `json:"$comment,omitempty"`
	Const                string                                `json:"const,omitempty"`
	Default              any                                   `json:"default,omitempty"` // If specified, it is a value of "type".
	Dependencies         map[string][]string                   `json:"dependencies,omitempty"`
	Description          string                                `json:"description,omitempty"`
	Enum                 *UnprocessedEnum                      `json:"enum,omitempty"`
	Examples             any                                   `json:"examples,omitempty"`
	Format               string                                `json:"format,omitempty"`
	InsertionOrder       *bool                                 `json:"insertionOrder,omitempty"`
	Items                *UnprocessedStructuredType            `json:"items,omitempty"`
	MaxItems             *int                                  `json:"maxItems,omitempty"`
	MaxLength            *int                                  `json:"maxLength,omitempty"`
	MaxProperties        *int                                  `json:"maxProperties,omitempty"`
	Maximum              *json.Number                          `json:"maximum,omitempty"`
	MinItems             *int                                  `json:"minItems,omitempty"`
	MinLength            *int                                  `json:"minLength,omitempty"`
	MinProperties        *int                                  `json:"minProperties,omitempty"`
	Minimum              *json.Number                          `json:"minimum,omitempty"`
	MultipleOf           *json.Number                          `json:"multipleOf,omitempty"`
	OneOf                []*UnprocessedStructuredType          `json:"oneOf,omitempty"`
	Pattern              *string                               `json:"pattern,omitempty"`
	PatternProperties    map[string]*UnprocessedStructuredType `json:"patternProperties,omitempty"`
	Properties           map[string]*UnprocessedStructuredType `json:"properties,omitempty"`
	Ref                  *string                               `json:"$ref,omitempty"`
	RelationshipRef      *UnprocessedRelationshipRef           `json:"relationshipRef,omitempty"`
	Required             []string                              `json:"required,omitempty"`
	Title                string                                `json:"title,omitempty"`
	Type                 any                                   `json:"type,omitempty"` // This can be string or []string.
	UniqueItems          *bool                                 `json:"uniqueItems,omitempty"`
}

func (ust *UnprocessedStructuredType) collectProblems(pc *cfz.ProblemsCollector, plt *cfz.ProblemLocationTracker) {
	if ust.Type == "integer" {
		if _, err := ust.Enum.AsInt64(); err != nil {
			pc.Collect(plt, "invalid enum: %v", err.Error())
		}
	}

	if ust.Type == "number" {
		if _, err := ust.Enum.AsFloat64(); err != nil {
			pc.Collect(plt, "invalid enum: %v", err.Error())
		}
	}

	for name, pUST := range ust.Properties {
		pUST.collectProblems(pc, plt.WithPathElements(fmt.Sprintf("property[%v]", name)))
	}
}

func (ust *UnprocessedStructuredType) maybeToPreprocessedType(s *Schema, utlr *UnprocessedTopLevelResource, name string) *Type {
	if ust.Type != "object" {
		return nil
	}

	t := &Type{
		IsTopLevelResourceType: false,
		Name:                   fmt.Sprintf("%v.%v", utlr.TypeName, name),
		Description:            ust.Description,
		Properties:             make(map[string]*Property),
		s:                      s,
		utlr:                   utlr,
		ust:                    ust,
	}

	return t
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
	AdditionalIdentifiers          [][]string                            `json:"additionalIdentifiers,omitempty"`
	AdditionalProperties           bool                                  `json:"additionalProperties,omitempty"`
	AllOf                          any                                   `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                          any                                   `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	Comment                        string                                `json:"$comment,omitempty"`
	ConditionalCreatOnlyProperties []string                              `json:"conditionalCreateOnlyProperties,omitempty"`
	CreateOnlyProperties           []string                              `json:"createOnlyProperties,omitempty"`
	Definitions                    map[string]*UnprocessedStructuredType `json:"definitions,omitempty"`
	DeprecatedProperties           []string                              `json:"deprecatedProperties,omitempty"`
	Description                    string                                `json:"description,omitempty"`
	DocumentationURL               string                                `json:"documentationUrl,omitempty"`
	Handlers                       *UnprocessedHandlers                  `json:"handlers,omitempty"`
	OneOf                          []*UnprocessedStructuredType          `json:"oneOf,omitempty"`
	PrimaryIdentifier              []string                              `json:"primaryIdentifier,omitempty"`
	Properties                     map[string]*UnprocessedStructuredType `json:"properties,omitempty"`
	PropertyTransform              map[string]string                     `json:"propertyTransform,omitempty"`
	ReadOnlyProperties             []string                              `json:"readOnlyProperties,omitempty"`
	ReplacementStrategy            string                                `json:"replacementStrategy,omitempty"`
	Required                       []string                              `json:"required,omitempty"`
	ResourceLink                   *UnprocessedResourceLink              `json:"resourceLink,omitempty"`
	Schema                         string                                `json:"$schema,omitempty"`
	SourceURL                      string                                `json:"sourceUrl,omitempty"`
	Taggable                       bool                                  `json:"taggable,omitempty"`
	Tagging                        *UnprocessedTagging                   `json:"tagging,omitempty"`
	TypeName                       string                                `json:"typeName,omitempty"`
	WriteOnlyProperties            []string                              `json:"writeOnlyProperties,omitempty"`
}

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

func (utlr *UnprocessedTopLevelResource) toPreprocessedTypes(s *Schema) (*Type, []*Type) {
	tlr := &Type{
		IsTopLevelResourceType: true,
		Name:                   utlr.TypeName,
		Description:            utlr.Description,
		Properties:             make(map[string]*Property),
		s:                      s,
		utlr:                   utlr,
		ust:                    nil,
	}

	sts := make([]*Type, 0, len(utlr.Definitions))

	for name, ust := range utlr.Definitions {
		if t := ust.maybeToPreprocessedType(s, utlr, name); t != nil {
			sts = append(sts, t)
		}
	}

	return tlr, sts
}
