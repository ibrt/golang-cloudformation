package cfschemaz

// Resource describes the CloudFormation JSON schema for a resource.
type Resource struct {
	AdditionalIdentifiers          [][]string           `json:"additionalIdentifiers,omitempty"`
	ConditionalCreatOnlyProperties []string             `json:"conditionalCreateOnlyProperties,omitempty"`
	CreateOnlyProperties           []string             `json:"createOnlyProperties,omitempty"`
	Definitions                    map[string]*Property `json:"definitions,omitempty"`
	DeprecatedProperties           []string             `json:"deprecatedProperties,omitempty"`
	Description                    string               `json:"description,omitempty"`
	DocumentationURL               string               `json:"documentationUrl,omitempty"`
	Handlers                       *ResourceHandlers    `json:"handlers,omitempty"`
	PrimaryIdentifier              []string             `json:"primaryIdentifier,omitempty"`
	PropertyTransform              map[string]string    `json:"propertyTransform,omitempty"`
	ReadOnlyProperties             []string             `json:"readOnlyProperties,omitempty"`
	ReplacementStrategy            string               `json:"replacementStrategy,omitempty"`
	ResourceLink                   *ResourceLink        `json:"resourceLink,omitempty"`
	Schema                         string               `json:"$schema,omitempty"`
	SourceURL                      string               `json:"sourceUrl,omitempty"`
	Taggable                       bool                 `json:"taggable,omitempty"`
	Tagging                        *ResourceTagging     `json:"tagging,omitempty"`
	TypeName                       string               `json:"typeName,omitempty"`
	WriteOnlyProperties            []string             `json:"writeOnlyProperties,omitempty"`

	// Shared with Property
	AdditionalProperties bool                 `json:"additionalProperties,omitempty"`
	AllOf                any                  `json:"allOf,omitempty"` // TODO(ibrt): Type.
	AnyOf                any                  `json:"anyOf,omitempty"` // TODO(ibrt): Type.
	Comment              string               `json:"$comment,omitempty"`
	OneOf                []*Property          `json:"oneOf,omitempty"`
	Properties           map[string]*Property `json:"properties,omitempty"`
	Required             []string             `json:"required,omitempty"`
}

// ResourceLink describes part of the CloudFormation JSON schema for a resource.
type ResourceLink struct {
	TemplateURI string            `json:"templateUri,omitempty"`
	Mappings    map[string]string `json:"mappings,omitempty"`
}

// ResourceHandlers describes part of the CloudFormation JSON schema for a resource.
type ResourceHandlers struct {
	Create *ResourceHandlersEntry `json:"create,omitempty"`
	Read   *ResourceHandlersEntry `json:"read,omitempty"`
	Update *ResourceHandlersEntry `json:"update,omitempty"`
	Delete *ResourceHandlersEntry `json:"delete,omitempty"`
	List   *ResourceHandlersEntry `json:"list,omitempty"`
}

// ResourceHandlersEntry describes part of the CloudFormation JSON schema for a resource.
type ResourceHandlersEntry struct {
	Permissions      []string  `json:"permissions,omitempty"`
	HandlerSchema    *Property `json:"handlerSchema,omitempty"`
	TimeoutInMinutes *int      `json:"timeoutInMinutes,omitempty"`
}

// ResourceTagging describes part of the CloudFormation JSON schema for a resource.
type ResourceTagging struct {
	Taggable                 bool     `json:"taggable,omitempty"`
	TagOnCreate              bool     `json:"tagOnCreate,omitempty"`
	TagUpdatable             bool     `json:"tagUpdatable,omitempty"`
	CloudFormationSystemTags bool     `json:"cloudFormationSystemTags,omitempty"`
	TagProperty              string   `json:"tagProperty,omitempty"`
	Permissions              []string `json:"permissions,omitempty"`
}
