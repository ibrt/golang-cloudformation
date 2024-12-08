// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emrserverless

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMRServerless_Application_AutoStartConfiguration__Type is the CloudFormation type for AWS::EMRServerless::Application.AutoStartConfiguration.
	AWS_EMRServerless_Application_AutoStartConfiguration__Type = "AWS::EMRServerless::Application.AutoStartConfiguration"
)

var (
	// AWS_EMRServerless_Application_AutoStartConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::EMRServerless::Application.AutoStartConfiguration.
	AWS_EMRServerless_Application_AutoStartConfiguration__PropertiesMap = struct {
		Enabled string
	}{
		Enabled: "Enabled",
	}

	// AWS_EMRServerless_Application_AutoStartConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::EMRServerless::Application.AutoStartConfiguration.
	AWS_EMRServerless_Application_AutoStartConfiguration__PropertiesSlice = []string{
		AWS_EMRServerless_Application_AutoStartConfiguration__PropertiesMap.Enabled,
	}
)

// AWS_EMRServerless_Application_AutoStartConfiguration is a binding for AWS::EMRServerless::Application.AutoStartConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostartconfiguration.html
type AWS_EMRServerless_Application_AutoStartConfiguration struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-autostartconfiguration.html#cfn-emrserverless-application-autostartconfiguration-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_EMRServerless_Application_AutoStartConfiguration initializes a new AWS_EMRServerless_Application_AutoStartConfiguration.
func New__AWS_EMRServerless_Application_AutoStartConfiguration() AWS_EMRServerless_Application_AutoStartConfiguration {
	return AWS_EMRServerless_Application_AutoStartConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_EMRServerless_Application_AutoStartConfiguration) GetType() string {
	return AWS_EMRServerless_Application_AutoStartConfiguration__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_EMRServerless_Application_AutoStartConfiguration) Set__Enabled(v cfz.Expression[bool]) AWS_EMRServerless_Application_AutoStartConfiguration {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_EMRServerless_Application_AutoStartConfiguration) SetV__Enabled(v bool) AWS_EMRServerless_Application_AutoStartConfiguration {
	t.Enabled = cfz.V(v)
	return t
}
