package cfschemaz

// ParsedSchema describes part of the CloudFormation JSON schema.
type ParsedSchema struct {
	AdditionalIdentifiers          [][]string                       `json:"additionalIdentifiers,omitempty"`
	ConditionalCreatOnlyProperties []string                         `json:"conditionalCreateOnlyProperties,omitempty"`
	CreateOnlyProperties           []string                         `json:"createOnlyProperties,omitempty"`
	Definitions                    map[string]*ParsedSchemaProperty `json:"definitions,omitempty"`
	DeprecatedProperties           []string                         `json:"deprecatedProperties,omitempty"`
	Description                    string                           `json:"description,omitempty"`
	DocumentationURL               string                           `json:"documentationUrl,omitempty"`
	Handlers                       *ParsedSchemaHandlers            `json:"handlers,omitempty"`
	PrimaryIdentifier              []string                         `json:"primaryIdentifier,omitempty"`
	PropertyTransform              map[string]string                `json:"propertyTransform,omitempty"`
	ReadOnlyProperties             []string                         `json:"readOnlyProperties,omitempty"`
	ReplacementStrategy            string                           `json:"replacementStrategy,omitempty"`
	ResourceLink                   *ParsedSchemaResourceLink        `json:"resourceLink,omitempty"`
	Schema                         string                           `json:"$schema,omitempty"`
	SourceURL                      string                           `json:"sourceUrl,omitempty"`
	Taggable                       bool                             `json:"taggable,omitempty"`
	Tagging                        *ParsedSchemaTagging             `json:"tagging,omitempty"`
	TypeName                       string                           `json:"typeName,omitempty"`
	WriteOnlyProperties            []string                         `json:"writeOnlyProperties,omitempty"`

	// Shared with ParsedSchemaProperty.
	AdditionalProperties bool                             `json:"additionalProperties,omitempty"`
	AllOf                any                              `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                any                              `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	Comment              string                           `json:"$comment,omitempty"`
	OneOf                []*ParsedSchemaProperty          `json:"oneOf,omitempty"`
	Properties           map[string]*ParsedSchemaProperty `json:"properties,omitempty"`
	Required             []string                         `json:"required,omitempty"`
}

// ParsedSchemaHandlers describes part of the CloudFormation JSON schema for a resource.
type ParsedSchemaHandlers struct {
	Create *ParsedSchemaHandlersOperation `json:"create,omitempty"`
	Delete *ParsedSchemaHandlersOperation `json:"delete,omitempty"`
	List   *ParsedSchemaHandlersOperation `json:"list,omitempty"`
	Read   *ParsedSchemaHandlersOperation `json:"read,omitempty"`
	Update *ParsedSchemaHandlersOperation `json:"update,omitempty"`
}

// ParsedSchemaHandlersOperation describes part of the CloudFormation JSON schema for a resource.
type ParsedSchemaHandlersOperation struct {
	HandlerSchema    *ParsedSchemaProperty `json:"handlerSchema,omitempty"`
	Permissions      []string              `json:"permissions,omitempty"`
	TimeoutInMinutes *int                  `json:"timeoutInMinutes,omitempty"`
}

// ParsedSchemaProperty describes part of the CloudFormation JSON schema.
type ParsedSchemaProperty struct {
	AdditionalProperties bool                             `json:"additionalProperties,omitempty"`
	AllOf                any                              `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                any                              `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	ArrayType            string                           `json:"arrayType,omitempty"`
	Comment              string                           `json:"$comment,omitempty"`
	Const                string                           `json:"const,omitempty"`
	Default              any                              `json:"default,omitempty"` // TODO(ibrt): Type.
	Dependencies         map[string][]string              `json:"dependencies,omitempty"`
	Description          string                           `json:"description,omitempty"`
	Enum                 any                              `json:"enum,omitempty"`     // TODO(ibrt): Type.
	Examples             any                              `json:"examples,omitempty"` // TODO(ibrt): Type.
	Format               string                           `json:"format,omitempty"`
	InsertionOrder       *bool                            `json:"insertionOrder,omitempty"`
	Items                *ParsedSchemaProperty            `json:"items,omitempty"`
	MaxItems             *int                             `json:"maxItems,omitempty"`
	MaxLength            *int                             `json:"maxLength,omitempty"`
	MaxProperties        *int                             `json:"maxProperties,omitempty"`
	Maximum              *float64                         `json:"maximum,omitempty"`
	MinItems             *int                             `json:"minItems,omitempty"`
	MinLength            *int                             `json:"minLength,omitempty"`
	MinProperties        *int                             `json:"minProperties,omitempty"`
	Minimum              *float64                         `json:"minimum,omitempty"`
	MultipleOf           *float64                         `json:"multipleOf,omitempty"`
	OneOf                []*ParsedSchemaProperty          `json:"oneOf,omitempty"`
	Pattern              *string                          `json:"pattern,omitempty"`
	PatternProperties    map[string]*ParsedSchemaProperty `json:"patternProperties,omitempty"`
	Properties           map[string]*ParsedSchemaProperty `json:"properties,omitempty"`
	Ref                  *string                          `json:"$ref,omitempty"`
	RelationshipRef      *ParsedPropertyRelationshipRef   `json:"relationshipRef,omitempty"`
	Required             []string                         `json:"required,omitempty"`
	Title                string                           `json:"title,omitempty"`
	Type                 any                              `json:"type,omitempty"` // TODO(ibrt): Type, so far string or []string.
	UniqueItems          *bool                            `json:"uniqueItems,omitempty"`
}

// ParsedPropertyRelationshipRef describes part of the CloudFormation JSON schema.
type ParsedPropertyRelationshipRef struct {
	TypeName     string `json:"typeName,omitempty"`
	PropertyPath string `json:"propertyPath,omitempty"`
}

// ParsedSchemaResourceLink describes part of the CloudFormation JSON schema.
type ParsedSchemaResourceLink struct {
	Mappings    map[string]string `json:"mappings,omitempty"`
	TemplateURI string            `json:"templateUri,omitempty"`
}

// ParsedSchemaTagging describes part of the CloudFormation JSON schema for a resource.
type ParsedSchemaTagging struct {
	CloudFormationSystemTags bool     `json:"cloudFormationSystemTags,omitempty"`
	Permissions              []string `json:"permissions,omitempty"`
	TagOnCreate              bool     `json:"tagOnCreate,omitempty"`
	TagProperty              string   `json:"tagProperty,omitempty"`
	TagUpdatable             bool     `json:"tagUpdatable,omitempty"`
	Taggable                 bool     `json:"taggable,omitempty"`
}
