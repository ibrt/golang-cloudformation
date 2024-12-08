// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__Type is the CloudFormation type for AWS::Bedrock::FlowVersion.PromptInferenceConfiguration.
	AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__Type = "AWS::Bedrock::FlowVersion.PromptInferenceConfiguration"
)

var (
	// AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::FlowVersion.PromptInferenceConfiguration.
	AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__PropertiesMap = struct {
		Text string
	}{
		Text: "Text",
	}

	// AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::FlowVersion.PromptInferenceConfiguration.
	AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__PropertiesMap.Text,
	}
)

// AWS_Bedrock_FlowVersion_PromptInferenceConfiguration is a binding for AWS::Bedrock::FlowVersion.PromptInferenceConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptinferenceconfiguration.html
type AWS_Bedrock_FlowVersion_PromptInferenceConfiguration struct {
	// Text is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptinferenceconfiguration.html#cfn-bedrock-flowversion-promptinferenceconfiguration-text
	Text cfz.Expression[AWS_Bedrock_FlowVersion_PromptModelInferenceConfiguration] `json:"Text,omitempty"`
}

// New__AWS_Bedrock_FlowVersion_PromptInferenceConfiguration initializes a new AWS_Bedrock_FlowVersion_PromptInferenceConfiguration.
func New__AWS_Bedrock_FlowVersion_PromptInferenceConfiguration() AWS_Bedrock_FlowVersion_PromptInferenceConfiguration {
	return AWS_Bedrock_FlowVersion_PromptInferenceConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_FlowVersion_PromptInferenceConfiguration) GetType() string {
	return AWS_Bedrock_FlowVersion_PromptInferenceConfiguration__Type
}

// Set__Text updates property "Text".
func (t AWS_Bedrock_FlowVersion_PromptInferenceConfiguration) Set__Text(v cfz.Expression[AWS_Bedrock_FlowVersion_PromptModelInferenceConfiguration]) AWS_Bedrock_FlowVersion_PromptInferenceConfiguration {
	t.Text = v
	return t
}

// SetV__Text updates property "Text".
func (t AWS_Bedrock_FlowVersion_PromptInferenceConfiguration) SetV__Text(v AWS_Bedrock_FlowVersion_PromptModelInferenceConfiguration) AWS_Bedrock_FlowVersion_PromptInferenceConfiguration {
	t.Text = cfz.V(v)
	return t
}
