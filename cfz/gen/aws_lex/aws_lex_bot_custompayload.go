// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_CustomPayload__Type is the CloudFormation type for AWS::Lex::Bot.CustomPayload.
	AWS_Lex_Bot_CustomPayload__Type = "AWS::Lex::Bot.CustomPayload"
)

var (
	// AWS_Lex_Bot_CustomPayload__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.CustomPayload.
	AWS_Lex_Bot_CustomPayload__PropertiesMap = struct {
		Value string
	}{
		Value: "Value",
	}

	// AWS_Lex_Bot_CustomPayload__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.CustomPayload.
	AWS_Lex_Bot_CustomPayload__PropertiesSlice = []string{
		AWS_Lex_Bot_CustomPayload__PropertiesMap.Value,
	}
)

// AWS_Lex_Bot_CustomPayload is a binding for AWS::Lex::Bot.CustomPayload.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-custompayload.html
type AWS_Lex_Bot_CustomPayload struct {
	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-custompayload.html#cfn-lex-bot-custompayload-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_Lex_Bot_CustomPayload initializes a new AWS_Lex_Bot_CustomPayload.
func New__AWS_Lex_Bot_CustomPayload() AWS_Lex_Bot_CustomPayload {
	return AWS_Lex_Bot_CustomPayload{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_CustomPayload) GetType() string {
	return AWS_Lex_Bot_CustomPayload__Type
}

// Set__Value updates property "Value".
func (t AWS_Lex_Bot_CustomPayload) Set__Value(v cfz.Expression[string]) AWS_Lex_Bot_CustomPayload {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_Lex_Bot_CustomPayload) SetV__Value(v string) AWS_Lex_Bot_CustomPayload {
	t.Value = cfz.V(v)
	return t
}
