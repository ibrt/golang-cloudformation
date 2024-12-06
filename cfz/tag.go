package cfz

// Tag is a structured type used widely across many CloudFormation resources.
// It is actually defined in the spec, but for convenience we provide a manually crafted implementation here.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html
type Tag struct {
	// Key is a property.
	Key Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	Value Expression[string] `json:"Value,omitempty"`
}
