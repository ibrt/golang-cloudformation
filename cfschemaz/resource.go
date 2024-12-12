package cfschemaz

// Resource describes part of the CloudFormation JSON schema.
type Resource struct {
	AdditionalIdentifiers          [][]string               `json:"additionalIdentifiers,omitempty"`
	AdditionalProperties           bool                     `json:"additionalProperties,omitempty"`
	AllOf                          any                      `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                          any                      `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	Comment                        string                   `json:"$comment,omitempty"`
	ConditionalCreatOnlyProperties []string                 `json:"conditionalCreateOnlyProperties,omitempty"`
	CreateOnlyProperties           []string                 `json:"createOnlyProperties,omitempty"`
	Definitions                    map[string]*Definition   `json:"definitions,omitempty"`
	DeprecatedProperties           []string                 `json:"deprecatedProperties,omitempty"`
	Description                    string                   `json:"description,omitempty"`
	DocumentationURL               string                   `json:"documentationUrl,omitempty"`
	Handlers                       *ResourceHandlers        `json:"handlers,omitempty"`
	OneOf                          []*Definition            `json:"oneOf,omitempty"`
	PrimaryIdentifier              []string                 `json:"primaryIdentifier,omitempty"`
	Properties                     map[string]*Definition   `json:"properties,omitempty"`
	PropertyTransform              map[string]string        `json:"propertyTransform,omitempty"`
	ReadOnlyProperties             []string                 `json:"readOnlyProperties,omitempty"`
	ReplacementStrategy            string                   `json:"replacementStrategy,omitempty"`
	Required                       []string                 `json:"required,omitempty"`
	ResourceLink                   *UnprocessedResourceLink `json:"resourceLink,omitempty"`
	Schema                         string                   `json:"$schema,omitempty"`
	SourceURL                      string                   `json:"sourceUrl,omitempty"`
	Taggable                       bool                     `json:"taggable,omitempty"`
	Tagging                        *ResourceTagging         `json:"tagging,omitempty"`
	TypeName                       string                   `json:"typeName,omitempty"`
	WriteOnlyProperties            []string                 `json:"writeOnlyProperties,omitempty"`
}

// ResourceHandlers describes part of the CloudFormation JSON schema for a resource.
type ResourceHandlers struct {
	Create *ResourceHandlersOperation `json:"create,omitempty"`
	Delete *ResourceHandlersOperation `json:"delete,omitempty"`
	List   *ResourceHandlersOperation `json:"list,omitempty"`
	Read   *ResourceHandlersOperation `json:"read,omitempty"`
	Update *ResourceHandlersOperation `json:"update,omitempty"`
}

// ResourceHandlersOperation describes part of the CloudFormation JSON schema for a resource.
type ResourceHandlersOperation struct {
	HandlerSchema    *Definition `json:"handlerSchema,omitempty"`
	Permissions      []string    `json:"permissions,omitempty"`
	TimeoutInMinutes *int        `json:"timeoutInMinutes,omitempty"`
}

// UnprocessedResourceLink describes part of the CloudFormation JSON schema.
type UnprocessedResourceLink struct {
	Mappings    map[string]string `json:"mappings,omitempty"`
	TemplateURI string            `json:"templateUri,omitempty"`
}

// ResourceTagging describes part of the CloudFormation JSON schema for a resource.
type ResourceTagging struct {
	CloudFormationSystemTags bool     `json:"cloudFormationSystemTags,omitempty"`
	Permissions              []string `json:"permissions,omitempty"`
	TagOnCreate              bool     `json:"tagOnCreate,omitempty"`
	TagProperty              string   `json:"tagProperty,omitempty"`
	TagUpdatable             bool     `json:"tagUpdatable,omitempty"`
	Taggable                 bool     `json:"taggable,omitempty"`
}
