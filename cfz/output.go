package cfz

// Output describes a CloudFormation output.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html
type Output struct {
	Description string             `json:"Description,omitempty"`
	Value       Expression[string] `json:"Value,omitempty"`
	Export      *OutputExport      `json:"Export,omitempty"`
}

// OutputExport describes part of a CloudFormation output.
type OutputExport struct {
	Name Expression[string] `json:"Name,omitempty"`
}
