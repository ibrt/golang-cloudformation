// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_TopicRule_TopicRulePayload__Type is the CloudFormation type for AWS::IoT::TopicRule.TopicRulePayload.
	AWS_IoT_TopicRule_TopicRulePayload__Type = "AWS::IoT::TopicRule.TopicRulePayload"
)

var (
	// AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap reports all the CloudFormation properties for AWS::IoT::TopicRule.TopicRulePayload.
	AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap = struct {
		Actions          string
		AwsIotSqlVersion string
		Description      string
		ErrorAction      string
		RuleDisabled     string
		Sql              string
	}{
		Actions:          "Actions",
		AwsIotSqlVersion: "AwsIotSqlVersion",
		Description:      "Description",
		ErrorAction:      "ErrorAction",
		RuleDisabled:     "RuleDisabled",
		Sql:              "Sql",
	}

	// AWS_IoT_TopicRule_TopicRulePayload__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::TopicRule.TopicRulePayload.
	AWS_IoT_TopicRule_TopicRulePayload__PropertiesSlice = []string{
		AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap.Actions,
		AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap.AwsIotSqlVersion,
		AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap.Description,
		AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap.ErrorAction,
		AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap.RuleDisabled,
		AWS_IoT_TopicRule_TopicRulePayload__PropertiesMap.Sql,
	}
)

// AWS_IoT_TopicRule_TopicRulePayload is a binding for AWS::IoT::TopicRule.TopicRulePayload.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html
type AWS_IoT_TopicRule_TopicRulePayload struct {
	// Actions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-actions
	Actions cfz.ExpressionSlice[AWS_IoT_TopicRule_Action] `json:"Actions,omitempty"`

	// AwsIotSqlVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-awsiotsqlversion
	AwsIotSqlVersion cfz.Expression[string] `json:"AwsIotSqlVersion,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// ErrorAction is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-erroraction
	ErrorAction cfz.Expression[AWS_IoT_TopicRule_Action] `json:"ErrorAction,omitempty"`

	// RuleDisabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-ruledisabled
	RuleDisabled cfz.Expression[bool] `json:"RuleDisabled,omitempty"`

	// Sql is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-sql
	Sql cfz.Expression[string] `json:"Sql,omitempty"`
}

// New__AWS_IoT_TopicRule_TopicRulePayload initializes a new AWS_IoT_TopicRule_TopicRulePayload.
func New__AWS_IoT_TopicRule_TopicRulePayload() AWS_IoT_TopicRule_TopicRulePayload {
	return AWS_IoT_TopicRule_TopicRulePayload{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_TopicRule_TopicRulePayload) GetType() string {
	return AWS_IoT_TopicRule_TopicRulePayload__Type
}

// Set__Actions updates property "Actions".
func (t AWS_IoT_TopicRule_TopicRulePayload) Set__Actions(v cfz.ExpressionSlice[AWS_IoT_TopicRule_Action]) AWS_IoT_TopicRule_TopicRulePayload {
	t.Actions = v
	return t
}

// SetS__Actions updates property "Actions".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetS__Actions(v ...cfz.Expression[AWS_IoT_TopicRule_Action]) AWS_IoT_TopicRule_TopicRulePayload {
	t.Actions = cfz.S(v...)
	return t
}

// SetSV__Actions updates property "Actions".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetSV__Actions(v ...AWS_IoT_TopicRule_Action) AWS_IoT_TopicRule_TopicRulePayload {
	t.Actions = cfz.SV(v...)
	return t
}

// Set__AwsIotSqlVersion updates property "AwsIotSqlVersion".
func (t AWS_IoT_TopicRule_TopicRulePayload) Set__AwsIotSqlVersion(v cfz.Expression[string]) AWS_IoT_TopicRule_TopicRulePayload {
	t.AwsIotSqlVersion = v
	return t
}

// SetV__AwsIotSqlVersion updates property "AwsIotSqlVersion".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetV__AwsIotSqlVersion(v string) AWS_IoT_TopicRule_TopicRulePayload {
	t.AwsIotSqlVersion = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_IoT_TopicRule_TopicRulePayload) Set__Description(v cfz.Expression[string]) AWS_IoT_TopicRule_TopicRulePayload {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetV__Description(v string) AWS_IoT_TopicRule_TopicRulePayload {
	t.Description = cfz.V(v)
	return t
}

// Set__ErrorAction updates property "ErrorAction".
func (t AWS_IoT_TopicRule_TopicRulePayload) Set__ErrorAction(v cfz.Expression[AWS_IoT_TopicRule_Action]) AWS_IoT_TopicRule_TopicRulePayload {
	t.ErrorAction = v
	return t
}

// SetV__ErrorAction updates property "ErrorAction".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetV__ErrorAction(v AWS_IoT_TopicRule_Action) AWS_IoT_TopicRule_TopicRulePayload {
	t.ErrorAction = cfz.V(v)
	return t
}

// Set__RuleDisabled updates property "RuleDisabled".
func (t AWS_IoT_TopicRule_TopicRulePayload) Set__RuleDisabled(v cfz.Expression[bool]) AWS_IoT_TopicRule_TopicRulePayload {
	t.RuleDisabled = v
	return t
}

// SetV__RuleDisabled updates property "RuleDisabled".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetV__RuleDisabled(v bool) AWS_IoT_TopicRule_TopicRulePayload {
	t.RuleDisabled = cfz.V(v)
	return t
}

// Set__Sql updates property "Sql".
func (t AWS_IoT_TopicRule_TopicRulePayload) Set__Sql(v cfz.Expression[string]) AWS_IoT_TopicRule_TopicRulePayload {
	t.Sql = v
	return t
}

// SetV__Sql updates property "Sql".
func (t AWS_IoT_TopicRule_TopicRulePayload) SetV__Sql(v string) AWS_IoT_TopicRule_TopicRulePayload {
	t.Sql = cfz.V(v)
	return t
}
