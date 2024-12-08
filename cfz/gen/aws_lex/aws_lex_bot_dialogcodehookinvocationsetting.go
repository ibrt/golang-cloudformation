// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_DialogCodeHookInvocationSetting__Type is the CloudFormation type for AWS::Lex::Bot.DialogCodeHookInvocationSetting.
	AWS_Lex_Bot_DialogCodeHookInvocationSetting__Type = "AWS::Lex::Bot.DialogCodeHookInvocationSetting"
)

var (
	// AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.DialogCodeHookInvocationSetting.
	AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesMap = struct {
		EnableCodeHookInvocation  string
		InvocationLabel           string
		IsActive                  string
		PostCodeHookSpecification string
	}{
		EnableCodeHookInvocation:  "EnableCodeHookInvocation",
		InvocationLabel:           "InvocationLabel",
		IsActive:                  "IsActive",
		PostCodeHookSpecification: "PostCodeHookSpecification",
	}

	// AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.DialogCodeHookInvocationSetting.
	AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesSlice = []string{
		AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesMap.EnableCodeHookInvocation,
		AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesMap.InvocationLabel,
		AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesMap.IsActive,
		AWS_Lex_Bot_DialogCodeHookInvocationSetting__PropertiesMap.PostCodeHookSpecification,
	}
)

// AWS_Lex_Bot_DialogCodeHookInvocationSetting is a binding for AWS::Lex::Bot.DialogCodeHookInvocationSetting.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html
type AWS_Lex_Bot_DialogCodeHookInvocationSetting struct {
	// EnableCodeHookInvocation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-enablecodehookinvocation
	EnableCodeHookInvocation cfz.Expression[bool] `json:"EnableCodeHookInvocation,omitempty"`

	// InvocationLabel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-invocationlabel
	InvocationLabel cfz.Expression[string] `json:"InvocationLabel,omitempty"`

	// IsActive is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-isactive
	IsActive cfz.Expression[bool] `json:"IsActive,omitempty"`

	// PostCodeHookSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogcodehookinvocationsetting.html#cfn-lex-bot-dialogcodehookinvocationsetting-postcodehookspecification
	PostCodeHookSpecification cfz.Expression[AWS_Lex_Bot_PostDialogCodeHookInvocationSpecification] `json:"PostCodeHookSpecification,omitempty"`
}

// New__AWS_Lex_Bot_DialogCodeHookInvocationSetting initializes a new AWS_Lex_Bot_DialogCodeHookInvocationSetting.
func New__AWS_Lex_Bot_DialogCodeHookInvocationSetting() AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	return AWS_Lex_Bot_DialogCodeHookInvocationSetting{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_DialogCodeHookInvocationSetting) GetType() string {
	return AWS_Lex_Bot_DialogCodeHookInvocationSetting__Type
}

// Set__EnableCodeHookInvocation updates property "EnableCodeHookInvocation".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) Set__EnableCodeHookInvocation(v cfz.Expression[bool]) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.EnableCodeHookInvocation = v
	return t
}

// SetV__EnableCodeHookInvocation updates property "EnableCodeHookInvocation".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) SetV__EnableCodeHookInvocation(v bool) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.EnableCodeHookInvocation = cfz.V(v)
	return t
}

// Set__InvocationLabel updates property "InvocationLabel".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) Set__InvocationLabel(v cfz.Expression[string]) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.InvocationLabel = v
	return t
}

// SetV__InvocationLabel updates property "InvocationLabel".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) SetV__InvocationLabel(v string) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.InvocationLabel = cfz.V(v)
	return t
}

// Set__IsActive updates property "IsActive".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) Set__IsActive(v cfz.Expression[bool]) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.IsActive = v
	return t
}

// SetV__IsActive updates property "IsActive".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) SetV__IsActive(v bool) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.IsActive = cfz.V(v)
	return t
}

// Set__PostCodeHookSpecification updates property "PostCodeHookSpecification".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) Set__PostCodeHookSpecification(v cfz.Expression[AWS_Lex_Bot_PostDialogCodeHookInvocationSpecification]) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.PostCodeHookSpecification = v
	return t
}

// SetV__PostCodeHookSpecification updates property "PostCodeHookSpecification".
func (t AWS_Lex_Bot_DialogCodeHookInvocationSetting) SetV__PostCodeHookSpecification(v AWS_Lex_Bot_PostDialogCodeHookInvocationSpecification) AWS_Lex_Bot_DialogCodeHookInvocationSetting {
	t.PostCodeHookSpecification = cfz.V(v)
	return t
}
