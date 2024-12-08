// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_SampleValue__Type is the CloudFormation type for AWS::Lex::Bot.SampleValue.
	AWS_Lex_Bot_SampleValue__Type = "AWS::Lex::Bot.SampleValue"
)

var (
	// AWS_Lex_Bot_SampleValue__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.SampleValue.
	AWS_Lex_Bot_SampleValue__PropertiesMap = struct {
		Value string
	}{
		Value: "Value",
	}

	// AWS_Lex_Bot_SampleValue__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.SampleValue.
	AWS_Lex_Bot_SampleValue__PropertiesSlice = []string{
		AWS_Lex_Bot_SampleValue__PropertiesMap.Value,
	}
)

// AWS_Lex_Bot_SampleValue is a binding for AWS::Lex::Bot.SampleValue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-samplevalue.html
type AWS_Lex_Bot_SampleValue struct {
	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-samplevalue.html#cfn-lex-bot-samplevalue-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_Lex_Bot_SampleValue initializes a new AWS_Lex_Bot_SampleValue.
func New__AWS_Lex_Bot_SampleValue() AWS_Lex_Bot_SampleValue {
	return AWS_Lex_Bot_SampleValue{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_SampleValue) GetType() string {
	return AWS_Lex_Bot_SampleValue__Type
}

// Set__Value updates property "Value".
func (t AWS_Lex_Bot_SampleValue) Set__Value(v cfz.Expression[string]) AWS_Lex_Bot_SampleValue {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_Lex_Bot_SampleValue) SetV__Value(v string) AWS_Lex_Bot_SampleValue {
	t.Value = cfz.V(v)
	return t
}
