// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_b2bi

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_B2BI_Transformer_Mapping__Type is the CloudFormation type for AWS::B2BI::Transformer.Mapping.
	AWS_B2BI_Transformer_Mapping__Type = "AWS::B2BI::Transformer.Mapping"
)

var (
	// AWS_B2BI_Transformer_Mapping__PropertiesMap reports all the CloudFormation properties for AWS::B2BI::Transformer.Mapping.
	AWS_B2BI_Transformer_Mapping__PropertiesMap = struct {
		Template         string
		TemplateLanguage string
	}{
		Template:         "Template",
		TemplateLanguage: "TemplateLanguage",
	}

	// AWS_B2BI_Transformer_Mapping__PropertiesSlice reports all the CloudFormation properties for AWS::B2BI::Transformer.Mapping.
	AWS_B2BI_Transformer_Mapping__PropertiesSlice = []string{
		AWS_B2BI_Transformer_Mapping__PropertiesMap.Template,
		AWS_B2BI_Transformer_Mapping__PropertiesMap.TemplateLanguage,
	}
)

// AWS_B2BI_Transformer_Mapping is a binding for AWS::B2BI::Transformer.Mapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html
type AWS_B2BI_Transformer_Mapping struct {
	// Template is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html#cfn-b2bi-transformer-mapping-template
	Template cfz.Expression[string] `json:"Template,omitempty"`

	// TemplateLanguage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html#cfn-b2bi-transformer-mapping-templatelanguage
	TemplateLanguage cfz.Expression[string] `json:"TemplateLanguage,omitempty"`
}

// New__AWS_B2BI_Transformer_Mapping initializes a new AWS_B2BI_Transformer_Mapping.
func New__AWS_B2BI_Transformer_Mapping() AWS_B2BI_Transformer_Mapping {
	return AWS_B2BI_Transformer_Mapping{}
}

// GetType returns the CloudFormation type.
func (AWS_B2BI_Transformer_Mapping) GetType() string {
	return AWS_B2BI_Transformer_Mapping__Type
}

// Set__Template updates property "Template".
func (t AWS_B2BI_Transformer_Mapping) Set__Template(v cfz.Expression[string]) AWS_B2BI_Transformer_Mapping {
	t.Template = v
	return t
}

// SetV__Template updates property "Template".
func (t AWS_B2BI_Transformer_Mapping) SetV__Template(v string) AWS_B2BI_Transformer_Mapping {
	t.Template = cfz.V(v)
	return t
}

// Set__TemplateLanguage updates property "TemplateLanguage".
func (t AWS_B2BI_Transformer_Mapping) Set__TemplateLanguage(v cfz.Expression[string]) AWS_B2BI_Transformer_Mapping {
	t.TemplateLanguage = v
	return t
}

// SetV__TemplateLanguage updates property "TemplateLanguage".
func (t AWS_B2BI_Transformer_Mapping) SetV__TemplateLanguage(v string) AWS_B2BI_Transformer_Mapping {
	t.TemplateLanguage = cfz.V(v)
	return t
}
