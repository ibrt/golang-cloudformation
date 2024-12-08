// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Flow_PromptInferenceConfiguration__Type is the CloudFormation type for AWS::Bedrock::Flow.PromptInferenceConfiguration.
	AWS_Bedrock_Flow_PromptInferenceConfiguration__Type = "AWS::Bedrock::Flow.PromptInferenceConfiguration"
)

var (
	// AWS_Bedrock_Flow_PromptInferenceConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Flow.PromptInferenceConfiguration.
	AWS_Bedrock_Flow_PromptInferenceConfiguration__PropertiesMap = struct {
		Text string
	}{
		Text: "Text",
	}

	// AWS_Bedrock_Flow_PromptInferenceConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Flow.PromptInferenceConfiguration.
	AWS_Bedrock_Flow_PromptInferenceConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_Flow_PromptInferenceConfiguration__PropertiesMap.Text,
	}
)

// AWS_Bedrock_Flow_PromptInferenceConfiguration is a binding for AWS::Bedrock::Flow.PromptInferenceConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinferenceconfiguration.html
type AWS_Bedrock_Flow_PromptInferenceConfiguration struct {
	// Text is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinferenceconfiguration.html#cfn-bedrock-flow-promptinferenceconfiguration-text
	Text cfz.Expression[AWS_Bedrock_Flow_PromptModelInferenceConfiguration] `json:"Text,omitempty"`
}

// New__AWS_Bedrock_Flow_PromptInferenceConfiguration initializes a new AWS_Bedrock_Flow_PromptInferenceConfiguration.
func New__AWS_Bedrock_Flow_PromptInferenceConfiguration() AWS_Bedrock_Flow_PromptInferenceConfiguration {
	return AWS_Bedrock_Flow_PromptInferenceConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Flow_PromptInferenceConfiguration) GetType() string {
	return AWS_Bedrock_Flow_PromptInferenceConfiguration__Type
}

// Set__Text updates property "Text".
func (t AWS_Bedrock_Flow_PromptInferenceConfiguration) Set__Text(v cfz.Expression[AWS_Bedrock_Flow_PromptModelInferenceConfiguration]) AWS_Bedrock_Flow_PromptInferenceConfiguration {
	t.Text = v
	return t
}

// SetV__Text updates property "Text".
func (t AWS_Bedrock_Flow_PromptInferenceConfiguration) SetV__Text(v AWS_Bedrock_Flow_PromptModelInferenceConfiguration) AWS_Bedrock_Flow_PromptInferenceConfiguration {
	t.Text = cfz.V(v)
	return t
}
