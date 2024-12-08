// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_TopicRule_SnsAction__Type is the CloudFormation type for AWS::IoT::TopicRule.SnsAction.
	AWS_IoT_TopicRule_SnsAction__Type = "AWS::IoT::TopicRule.SnsAction"
)

var (
	// AWS_IoT_TopicRule_SnsAction__PropertiesMap reports all the CloudFormation properties for AWS::IoT::TopicRule.SnsAction.
	AWS_IoT_TopicRule_SnsAction__PropertiesMap = struct {
		MessageFormat string
		RoleArn       string
		TargetArn     string
	}{
		MessageFormat: "MessageFormat",
		RoleArn:       "RoleArn",
		TargetArn:     "TargetArn",
	}

	// AWS_IoT_TopicRule_SnsAction__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::TopicRule.SnsAction.
	AWS_IoT_TopicRule_SnsAction__PropertiesSlice = []string{
		AWS_IoT_TopicRule_SnsAction__PropertiesMap.MessageFormat,
		AWS_IoT_TopicRule_SnsAction__PropertiesMap.RoleArn,
		AWS_IoT_TopicRule_SnsAction__PropertiesMap.TargetArn,
	}
)

// AWS_IoT_TopicRule_SnsAction is a binding for AWS::IoT::TopicRule.SnsAction.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html
type AWS_IoT_TopicRule_SnsAction struct {
	// MessageFormat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-messageformat
	MessageFormat cfz.Expression[string] `json:"MessageFormat,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// TargetArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-snsaction.html#cfn-iot-topicrule-snsaction-targetarn
	TargetArn cfz.Expression[string] `json:"TargetArn,omitempty"`
}

// New__AWS_IoT_TopicRule_SnsAction initializes a new AWS_IoT_TopicRule_SnsAction.
func New__AWS_IoT_TopicRule_SnsAction() AWS_IoT_TopicRule_SnsAction {
	return AWS_IoT_TopicRule_SnsAction{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_TopicRule_SnsAction) GetType() string {
	return AWS_IoT_TopicRule_SnsAction__Type
}

// Set__MessageFormat updates property "MessageFormat".
func (t AWS_IoT_TopicRule_SnsAction) Set__MessageFormat(v cfz.Expression[string]) AWS_IoT_TopicRule_SnsAction {
	t.MessageFormat = v
	return t
}

// SetV__MessageFormat updates property "MessageFormat".
func (t AWS_IoT_TopicRule_SnsAction) SetV__MessageFormat(v string) AWS_IoT_TopicRule_SnsAction {
	t.MessageFormat = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t AWS_IoT_TopicRule_SnsAction) Set__RoleArn(v cfz.Expression[string]) AWS_IoT_TopicRule_SnsAction {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t AWS_IoT_TopicRule_SnsAction) SetV__RoleArn(v string) AWS_IoT_TopicRule_SnsAction {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__TargetArn updates property "TargetArn".
func (t AWS_IoT_TopicRule_SnsAction) Set__TargetArn(v cfz.Expression[string]) AWS_IoT_TopicRule_SnsAction {
	t.TargetArn = v
	return t
}

// SetV__TargetArn updates property "TargetArn".
func (t AWS_IoT_TopicRule_SnsAction) SetV__TargetArn(v string) AWS_IoT_TopicRule_SnsAction {
	t.TargetArn = cfz.V(v)
	return t
}
