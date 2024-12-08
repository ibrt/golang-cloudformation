// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_DialogAction__Type is the CloudFormation type for AWS::Lex::Bot.DialogAction.
	AWS_Lex_Bot_DialogAction__Type = "AWS::Lex::Bot.DialogAction"
)

var (
	// AWS_Lex_Bot_DialogAction__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.DialogAction.
	AWS_Lex_Bot_DialogAction__PropertiesMap = struct {
		SlotToElicit        string
		SuppressNextMessage string
		Type                string
	}{
		SlotToElicit:        "SlotToElicit",
		SuppressNextMessage: "SuppressNextMessage",
		Type:                "Type",
	}

	// AWS_Lex_Bot_DialogAction__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.DialogAction.
	AWS_Lex_Bot_DialogAction__PropertiesSlice = []string{
		AWS_Lex_Bot_DialogAction__PropertiesMap.SlotToElicit,
		AWS_Lex_Bot_DialogAction__PropertiesMap.SuppressNextMessage,
		AWS_Lex_Bot_DialogAction__PropertiesMap.Type,
	}
)

// AWS_Lex_Bot_DialogAction is a binding for AWS::Lex::Bot.DialogAction.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html
type AWS_Lex_Bot_DialogAction struct {
	// SlotToElicit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-slottoelicit
	SlotToElicit cfz.Expression[string] `json:"SlotToElicit,omitempty"`

	// SuppressNextMessage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-suppressnextmessage
	SuppressNextMessage cfz.Expression[bool] `json:"SuppressNextMessage,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-dialogaction.html#cfn-lex-bot-dialogaction-type
	Type cfz.Expression[string] `json:"Type,omitempty"`
}

// New__AWS_Lex_Bot_DialogAction initializes a new AWS_Lex_Bot_DialogAction.
func New__AWS_Lex_Bot_DialogAction() AWS_Lex_Bot_DialogAction {
	return AWS_Lex_Bot_DialogAction{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_DialogAction) GetType() string {
	return AWS_Lex_Bot_DialogAction__Type
}

// Set__SlotToElicit updates property "SlotToElicit".
func (t AWS_Lex_Bot_DialogAction) Set__SlotToElicit(v cfz.Expression[string]) AWS_Lex_Bot_DialogAction {
	t.SlotToElicit = v
	return t
}

// SetV__SlotToElicit updates property "SlotToElicit".
func (t AWS_Lex_Bot_DialogAction) SetV__SlotToElicit(v string) AWS_Lex_Bot_DialogAction {
	t.SlotToElicit = cfz.V(v)
	return t
}

// Set__SuppressNextMessage updates property "SuppressNextMessage".
func (t AWS_Lex_Bot_DialogAction) Set__SuppressNextMessage(v cfz.Expression[bool]) AWS_Lex_Bot_DialogAction {
	t.SuppressNextMessage = v
	return t
}

// SetV__SuppressNextMessage updates property "SuppressNextMessage".
func (t AWS_Lex_Bot_DialogAction) SetV__SuppressNextMessage(v bool) AWS_Lex_Bot_DialogAction {
	t.SuppressNextMessage = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_Lex_Bot_DialogAction) Set__Type(v cfz.Expression[string]) AWS_Lex_Bot_DialogAction {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_Lex_Bot_DialogAction) SetV__Type(v string) AWS_Lex_Bot_DialogAction {
	t.Type = cfz.V(v)
	return t
}
