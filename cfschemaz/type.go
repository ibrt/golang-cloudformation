package cfschemaz

type Type struct {
	TypeName                        string                     `json:"typeName,omitempty"`
	Description                     string                     `json:"description,omitempty"`
	SourceURL                       string                     `json:"sourceUrl,omitempty"`
	ResourceLink                    *TypeResourceLink          `json:"resourceLink,omitempty"`
	Definitions                     map[string]*TypeDefinition `json:"definitions,omitempty"`
	Properties                      map[string]*TypeDefinition `json:"properties,omitempty"`
	Required                        []string                   `json:"required,omitempty"`
	ReadOnlyProperties              []string                   `json:"readOnlyProperties,omitempty"`
	CreateOnlyProperties            []string                   `json:"createOnlyProperties,omitempty"`
	ConditionalCreateOnlyProperties []string                   `json:"conditionalCreateOnlyProperties,omitempty"`
	PrimaryIdentifier               []string                   `json:"primaryIdentifier,omitempty"`
	WriteOnlyProperties             []string                   `json:"writeOnlyProperties,omitempty"`
	AdditionalProperties            bool                       `json:"additionalProperties,omitempty"`
}

type TypeResourceLink struct {
	TemplateURI string            `json:"templateUri,omitempty"`
	Mappings    map[string]string `json:"mappings,omitempty"`
}

type TypeDefinition struct {
	Description          string                             `json:"description,omitempty"`
	Type                 string                             `json:"type,omitempty"`
	Enum                 []string                           `json:"enum,omitempty"`
	Properties           map[string]*TypeDefinitionProperty `json:"properties,omitempty"`
	Required             []string                           `json:"required,omitempty"`
	AdditionalProperties bool                               `json:"additionalProperties,omitempty"`
	Handlers             *TypeHandlers                      `json:"handlers,omitempty"`
	Tagging              *TypeTagging                       `json:"tagging,omitempty"`
}

type TypeDefinitionProperty struct {
	Type           string                             `json:"type,omitempty"`
	Description    *string                            `json:"description,omitempty"`
	InsertionOrder *bool                              `json:"insertionOrder,omitempty"`
	UniqueItems    *bool                              `json:"uniqueItems,omitempty"`
	MinItems       *int                               `json:"minItems,omitempty"`
	MaxItems       *int                               `json:"maxItems,omitempty"`
	MinLength      *int                               `json:"minLength,omitempty"`
	MaxLength      *int                               `json:"maxLength,omitempty"`
	Minimum        *int                               `json:"minimum,omitempty"`
	Maximum        *int                               `json:"maximum,omitempty"`
	Items          *TypeDefinitionPropertyItems       `json:"items,omitempty"`
	Properties     map[string]*TypeDefinitionProperty `json:"properties,omitempty"`
}

type TypeDefinitionPropertyItems struct {
	Ref  *string `json:"$ref,omitempty"`
	Type *string `json:"type,omitempty"`
}

type TypeHandlers struct {
	Create *TypeHandlersOperation `json:"create,omitempty"`
	Read   *TypeHandlersOperation `json:"read,omitempty"`
	Update *TypeHandlersOperation `json:"update,omitempty"`
	Delete *TypeHandlersOperation `json:"delete,omitempty"`
	List   *TypeHandlersOperation `json:"list,omitempty"`
}

type TypeHandlersOperation struct {
	Permissions      []string `json:"permissions,omitempty"`
	TimeoutInMinutes *int     `json:"timeoutInMinutes,omitempty"`
}

type TypeTagging struct {
	Taggable                 bool     `json:"taggable,omitempty"`
	TagOnCreate              bool     `json:"tagOnCreate,omitempty"`
	TagUpdatable             bool     `json:"tagUpdatable,omitempty"`
	CloudFormationSystemTags bool     `json:"cloudFormationSystemTags,omitempty"`
	TagProperty              string   `json:"tagProperty,omitempty"`
	Permissions              []string `json:"permissions,omitempty"`
}
