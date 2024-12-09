package cfschemaz

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
	AllOf                any                                   `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                any                                   `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	ArrayType            string                                `json:"arrayType,omitempty"`
	Comment              string                                `json:"$comment,omitempty"`
	Const                string                                `json:"const,omitempty"`
	Default              any                                   `json:"default,omitempty"` // TODO(ibrt): Type.
	Dependencies         map[string][]string                   `json:"dependencies,omitempty"`
	Description          string                                `json:"description,omitempty"`
	Enum                 any                                   `json:"enum,omitempty"`     // TODO(ibrt): Type.
	Examples             any                                   `json:"examples,omitempty"` // TODO(ibrt): Type.
	Format               string                                `json:"format,omitempty"`
	InsertionOrder       *bool                                 `json:"insertionOrder,omitempty"`
	Items                *UnprocessedStructuredType            `json:"items,omitempty"`
	MaxItems             *int                                  `json:"maxItems,omitempty"`
	MaxLength            *int                                  `json:"maxLength,omitempty"`
	MaxProperties        *int                                  `json:"maxProperties,omitempty"`
	Maximum              *float64                              `json:"maximum,omitempty"`
	MinItems             *int                                  `json:"minItems,omitempty"`
	MinLength            *int                                  `json:"minLength,omitempty"`
	MinProperties        *int                                  `json:"minProperties,omitempty"`
	Minimum              *float64                              `json:"minimum,omitempty"`
	MultipleOf           *float64                              `json:"multipleOf,omitempty"`
	OneOf                []*UnprocessedStructuredType          `json:"oneOf,omitempty"`
	Pattern              *string                               `json:"pattern,omitempty"`
	PatternProperties    map[string]*UnprocessedStructuredType `json:"patternProperties,omitempty"`
	Properties           map[string]*UnprocessedStructuredType `json:"properties,omitempty"`
	Ref                  *string                               `json:"$ref,omitempty"`
	RelationshipRef      *UnprocessedRelationshipRef           `json:"relationshipRef,omitempty"`
	Required             []string                              `json:"required,omitempty"`
	Title                string                                `json:"title,omitempty"`
	Type                 any                                   `json:"type,omitempty"` // TODO(ibrt): Type, so far string or []string.
	UniqueItems          *bool                                 `json:"uniqueItems,omitempty"`
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
