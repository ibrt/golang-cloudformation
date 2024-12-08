// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::Application.SqlApplicationConfiguration.
	AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__Type = "AWS::KinesisAnalyticsV2::Application.SqlApplicationConfiguration"
)

var (
	// AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.SqlApplicationConfiguration.
	AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__PropertiesMap = struct {
		Inputs string
	}{
		Inputs: "Inputs",
	}

	// AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.SqlApplicationConfiguration.
	AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__PropertiesMap.Inputs,
	}
)

// AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration is a binding for AWS::KinesisAnalyticsV2::Application.SqlApplicationConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-sqlapplicationconfiguration.html
type AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration struct {
	// Inputs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-sqlapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-sqlapplicationconfiguration-inputs
	Inputs cfz.ExpressionSlice[AWS_KinesisAnalyticsV2_Application_Input] `json:"Inputs,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration initializes a new AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration.
func New__AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration() AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration {
	return AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration) GetType() string {
	return AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration__Type
}

// Set__Inputs updates property "Inputs".
func (t AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration) Set__Inputs(v cfz.ExpressionSlice[AWS_KinesisAnalyticsV2_Application_Input]) AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration {
	t.Inputs = v
	return t
}

// SetS__Inputs updates property "Inputs".
func (t AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration) SetS__Inputs(v ...cfz.Expression[AWS_KinesisAnalyticsV2_Application_Input]) AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration {
	t.Inputs = cfz.S(v...)
	return t
}

// SetSV__Inputs updates property "Inputs".
func (t AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration) SetSV__Inputs(v ...AWS_KinesisAnalyticsV2_Application_Input) AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration {
	t.Inputs = cfz.SV(v...)
	return t
}
