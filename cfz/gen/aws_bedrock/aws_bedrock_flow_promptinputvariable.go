// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Flow_PromptInputVariable__Type is the CloudFormation type for AWS::Bedrock::Flow.PromptInputVariable.
	AWS_Bedrock_Flow_PromptInputVariable__Type = "AWS::Bedrock::Flow.PromptInputVariable"
)

var (
	// AWS_Bedrock_Flow_PromptInputVariable__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Flow.PromptInputVariable.
	AWS_Bedrock_Flow_PromptInputVariable__PropertiesMap = struct {
		Name string
	}{
		Name: "Name",
	}

	// AWS_Bedrock_Flow_PromptInputVariable__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Flow.PromptInputVariable.
	AWS_Bedrock_Flow_PromptInputVariable__PropertiesSlice = []string{
		AWS_Bedrock_Flow_PromptInputVariable__PropertiesMap.Name,
	}
)

// AWS_Bedrock_Flow_PromptInputVariable is a binding for AWS::Bedrock::Flow.PromptInputVariable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinputvariable.html
type AWS_Bedrock_Flow_PromptInputVariable struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptinputvariable.html#cfn-bedrock-flow-promptinputvariable-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_Bedrock_Flow_PromptInputVariable initializes a new AWS_Bedrock_Flow_PromptInputVariable.
func New__AWS_Bedrock_Flow_PromptInputVariable() AWS_Bedrock_Flow_PromptInputVariable {
	return AWS_Bedrock_Flow_PromptInputVariable{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Flow_PromptInputVariable) GetType() string {
	return AWS_Bedrock_Flow_PromptInputVariable__Type
}

// Set__Name updates property "Name".
func (t AWS_Bedrock_Flow_PromptInputVariable) Set__Name(v cfz.Expression[string]) AWS_Bedrock_Flow_PromptInputVariable {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Bedrock_Flow_PromptInputVariable) SetV__Name(v string) AWS_Bedrock_Flow_PromptInputVariable {
	t.Name = cfz.V(v)
	return t
}
