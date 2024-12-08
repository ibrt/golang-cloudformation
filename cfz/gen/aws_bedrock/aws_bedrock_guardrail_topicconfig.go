// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Guardrail_TopicConfig__Type is the CloudFormation type for AWS::Bedrock::Guardrail.TopicConfig.
	AWS_Bedrock_Guardrail_TopicConfig__Type = "AWS::Bedrock::Guardrail.TopicConfig"
)

var (
	// AWS_Bedrock_Guardrail_TopicConfig__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Guardrail.TopicConfig.
	AWS_Bedrock_Guardrail_TopicConfig__PropertiesMap = struct {
		Definition string
		Examples   string
		Name       string
		Type       string
	}{
		Definition: "Definition",
		Examples:   "Examples",
		Name:       "Name",
		Type:       "Type",
	}

	// AWS_Bedrock_Guardrail_TopicConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Guardrail.TopicConfig.
	AWS_Bedrock_Guardrail_TopicConfig__PropertiesSlice = []string{
		AWS_Bedrock_Guardrail_TopicConfig__PropertiesMap.Definition,
		AWS_Bedrock_Guardrail_TopicConfig__PropertiesMap.Examples,
		AWS_Bedrock_Guardrail_TopicConfig__PropertiesMap.Name,
		AWS_Bedrock_Guardrail_TopicConfig__PropertiesMap.Type,
	}
)

// AWS_Bedrock_Guardrail_TopicConfig is a binding for AWS::Bedrock::Guardrail.TopicConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html
type AWS_Bedrock_Guardrail_TopicConfig struct {
	// Definition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-definition
	Definition cfz.Expression[string] `json:"Definition,omitempty"`

	// Examples is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-examples
	Examples cfz.ExpressionSlice[string] `json:"Examples,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-topicconfig.html#cfn-bedrock-guardrail-topicconfig-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Bedrock_Guardrail_TopicConfig initializes a new AWS_Bedrock_Guardrail_TopicConfig.
func New__AWS_Bedrock_Guardrail_TopicConfig() AWS_Bedrock_Guardrail_TopicConfig {
	return AWS_Bedrock_Guardrail_TopicConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Guardrail_TopicConfig) GetType() string {
	return AWS_Bedrock_Guardrail_TopicConfig__Type
}

// Set__Definition updates property "Definition".
func (t AWS_Bedrock_Guardrail_TopicConfig) Set__Definition(v cfz.Expression[string]) AWS_Bedrock_Guardrail_TopicConfig {
	t.Definition = v
	return t
}

// SetV__Definition updates property "Definition".
func (t AWS_Bedrock_Guardrail_TopicConfig) SetV__Definition(v string) AWS_Bedrock_Guardrail_TopicConfig {
	t.Definition = cfz.V(v)
	return t
}

// Set__Examples updates property "Examples".
func (t AWS_Bedrock_Guardrail_TopicConfig) Set__Examples(v cfz.ExpressionSlice[string]) AWS_Bedrock_Guardrail_TopicConfig {
	t.Examples = v
	return t
}

// SetS__Examples updates property "Examples".
func (t AWS_Bedrock_Guardrail_TopicConfig) SetS__Examples(v ...cfz.Expression[string]) AWS_Bedrock_Guardrail_TopicConfig {
	t.Examples = cfz.S(v...)
	return t
}

// SetSV__Examples updates property "Examples".
func (t AWS_Bedrock_Guardrail_TopicConfig) SetSV__Examples(v ...string) AWS_Bedrock_Guardrail_TopicConfig {
	t.Examples = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t AWS_Bedrock_Guardrail_TopicConfig) Set__Name(v cfz.Expression[string]) AWS_Bedrock_Guardrail_TopicConfig {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Bedrock_Guardrail_TopicConfig) SetV__Name(v string) AWS_Bedrock_Guardrail_TopicConfig {
	t.Name = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Bedrock_Guardrail_TopicConfig) Set__Type(v cfz.Expression[string]) AWS_Bedrock_Guardrail_TopicConfig {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Bedrock_Guardrail_TopicConfig) SetV__Type(v string) AWS_Bedrock_Guardrail_TopicConfig {
	t.Type = cfz.V(v)
	return t
}
