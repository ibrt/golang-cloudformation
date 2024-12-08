// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_chatbot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Chatbot_CustomAction_CustomActionDefinition__Type is the CloudFormation type for AWS::Chatbot::CustomAction.CustomActionDefinition.
	AWS_Chatbot_CustomAction_CustomActionDefinition__Type = "AWS::Chatbot::CustomAction.CustomActionDefinition"
)

var (
	// AWS_Chatbot_CustomAction_CustomActionDefinition__PropertiesMap reports all the CloudFormation properties for AWS::Chatbot::CustomAction.CustomActionDefinition.
	AWS_Chatbot_CustomAction_CustomActionDefinition__PropertiesMap = struct {
		CommandText string
	}{
		CommandText: "CommandText",
	}

	// AWS_Chatbot_CustomAction_CustomActionDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::Chatbot::CustomAction.CustomActionDefinition.
	AWS_Chatbot_CustomAction_CustomActionDefinition__PropertiesSlice = []string{
		AWS_Chatbot_CustomAction_CustomActionDefinition__PropertiesMap.CommandText,
	}
)

// AWS_Chatbot_CustomAction_CustomActionDefinition is a binding for AWS::Chatbot::CustomAction.CustomActionDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactiondefinition.html
type AWS_Chatbot_CustomAction_CustomActionDefinition struct {
	// CommandText is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chatbot-customaction-customactiondefinition.html#cfn-chatbot-customaction-customactiondefinition-commandtext
	CommandText cfz.Expression[string] `json:"CommandText,omitempty"`
}

// New__AWS_Chatbot_CustomAction_CustomActionDefinition initializes a new AWS_Chatbot_CustomAction_CustomActionDefinition.
func New__AWS_Chatbot_CustomAction_CustomActionDefinition() AWS_Chatbot_CustomAction_CustomActionDefinition {
	return AWS_Chatbot_CustomAction_CustomActionDefinition{}
}

// GetType returns the CloudFormation type.
func (AWS_Chatbot_CustomAction_CustomActionDefinition) GetType() string {
	return AWS_Chatbot_CustomAction_CustomActionDefinition__Type
}

// Set__CommandText updates property "CommandText".
func (t AWS_Chatbot_CustomAction_CustomActionDefinition) Set__CommandText(v cfz.Expression[string]) AWS_Chatbot_CustomAction_CustomActionDefinition {
	t.CommandText = v
	return t
}

// SetV__CommandText updates property "CommandText".
func (t AWS_Chatbot_CustomAction_CustomActionDefinition) SetV__CommandText(v string) AWS_Chatbot_CustomAction_CustomActionDefinition {
	t.CommandText = cfz.V(v)
	return t
}
