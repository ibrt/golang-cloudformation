// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpoint

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__Type is the CloudFormation type for AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration.
	AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__Type = "AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration"
)

var (
	// AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration.
	AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__PropertiesMap = struct {
		ButtonAction string
		Link         string
	}{
		ButtonAction: "ButtonAction",
		Link:         "Link",
	}

	// AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration.
	AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__PropertiesSlice = []string{
		AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__PropertiesMap.ButtonAction,
		AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__PropertiesMap.Link,
	}
)

// AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration is a binding for AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.html
type AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration struct {
	// ButtonAction is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.html#cfn-pinpoint-inapptemplate-overridebuttonconfiguration-buttonaction
	ButtonAction cfz.Expression[string] `json:"ButtonAction,omitempty"`

	// Link is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.html#cfn-pinpoint-inapptemplate-overridebuttonconfiguration-link
	Link cfz.Expression[string] `json:"Link,omitempty"`
}

// New__AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration initializes a new AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration.
func New__AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration() AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration {
	return AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration) GetType() string {
	return AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration__Type
}

// Set__ButtonAction updates property "ButtonAction".
func (t AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration) Set__ButtonAction(v cfz.Expression[string]) AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration {
	t.ButtonAction = v
	return t
}

// SetV__ButtonAction updates property "ButtonAction".
func (t AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration) SetV__ButtonAction(v string) AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration {
	t.ButtonAction = cfz.V(v)
	return t
}

// Set__Link updates property "Link".
func (t AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration) Set__Link(v cfz.Expression[string]) AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration {
	t.Link = v
	return t
}

// SetV__Link updates property "Link".
func (t AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration) SetV__Link(v string) AWS_Pinpoint_InAppTemplate_OverrideButtonConfiguration {
	t.Link = cfz.V(v)
	return t
}
