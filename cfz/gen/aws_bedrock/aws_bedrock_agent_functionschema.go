// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Agent_FunctionSchema__Type is the CloudFormation type for AWS::Bedrock::Agent.FunctionSchema.
	AWS_Bedrock_Agent_FunctionSchema__Type = "AWS::Bedrock::Agent.FunctionSchema"
)

var (
	// AWS_Bedrock_Agent_FunctionSchema__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Agent.FunctionSchema.
	AWS_Bedrock_Agent_FunctionSchema__PropertiesMap = struct {
		Functions string
	}{
		Functions: "Functions",
	}

	// AWS_Bedrock_Agent_FunctionSchema__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Agent.FunctionSchema.
	AWS_Bedrock_Agent_FunctionSchema__PropertiesSlice = []string{
		AWS_Bedrock_Agent_FunctionSchema__PropertiesMap.Functions,
	}
)

// AWS_Bedrock_Agent_FunctionSchema is a binding for AWS::Bedrock::Agent.FunctionSchema.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-functionschema.html
type AWS_Bedrock_Agent_FunctionSchema struct {
	// Functions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-functionschema.html#cfn-bedrock-agent-functionschema-functions
	Functions cfz.ExpressionSlice[AWS_Bedrock_Agent_Function] `json:"Functions,omitempty"`
}

// New__AWS_Bedrock_Agent_FunctionSchema initializes a new AWS_Bedrock_Agent_FunctionSchema.
func New__AWS_Bedrock_Agent_FunctionSchema() AWS_Bedrock_Agent_FunctionSchema {
	return AWS_Bedrock_Agent_FunctionSchema{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Agent_FunctionSchema) GetType() string {
	return AWS_Bedrock_Agent_FunctionSchema__Type
}

// Set__Functions updates property "Functions".
func (t AWS_Bedrock_Agent_FunctionSchema) Set__Functions(v cfz.ExpressionSlice[AWS_Bedrock_Agent_Function]) AWS_Bedrock_Agent_FunctionSchema {
	t.Functions = v
	return t
}

// SetS__Functions updates property "Functions".
func (t AWS_Bedrock_Agent_FunctionSchema) SetS__Functions(v ...cfz.Expression[AWS_Bedrock_Agent_Function]) AWS_Bedrock_Agent_FunctionSchema {
	t.Functions = cfz.S(v...)
	return t
}

// SetSV__Functions updates property "Functions".
func (t AWS_Bedrock_Agent_FunctionSchema) SetSV__Functions(v ...AWS_Bedrock_Agent_Function) AWS_Bedrock_Agent_FunctionSchema {
	t.Functions = cfz.SV(v...)
	return t
}
