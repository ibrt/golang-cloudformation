// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_applicationinsights

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ApplicationInsights_Application_CustomComponent__Type is the CloudFormation type for AWS::ApplicationInsights::Application.CustomComponent.
	AWS_ApplicationInsights_Application_CustomComponent__Type = "AWS::ApplicationInsights::Application.CustomComponent"
)

var (
	// AWS_ApplicationInsights_Application_CustomComponent__PropertiesMap reports all the CloudFormation properties for AWS::ApplicationInsights::Application.CustomComponent.
	AWS_ApplicationInsights_Application_CustomComponent__PropertiesMap = struct {
		ComponentName string
		ResourceList  string
	}{
		ComponentName: "ComponentName",
		ResourceList:  "ResourceList",
	}

	// AWS_ApplicationInsights_Application_CustomComponent__PropertiesSlice reports all the CloudFormation properties for AWS::ApplicationInsights::Application.CustomComponent.
	AWS_ApplicationInsights_Application_CustomComponent__PropertiesSlice = []string{
		AWS_ApplicationInsights_Application_CustomComponent__PropertiesMap.ComponentName,
		AWS_ApplicationInsights_Application_CustomComponent__PropertiesMap.ResourceList,
	}
)

// AWS_ApplicationInsights_Application_CustomComponent is a binding for AWS::ApplicationInsights::Application.CustomComponent.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-customcomponent.html
type AWS_ApplicationInsights_Application_CustomComponent struct {
	// ComponentName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-customcomponent.html#cfn-applicationinsights-application-customcomponent-componentname
	ComponentName cfz.Expression[string] `json:"ComponentName,omitempty"`

	// ResourceList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationinsights-application-customcomponent.html#cfn-applicationinsights-application-customcomponent-resourcelist
	ResourceList cfz.ExpressionSlice[string] `json:"ResourceList,omitempty"`
}

// New__AWS_ApplicationInsights_Application_CustomComponent initializes a new AWS_ApplicationInsights_Application_CustomComponent.
func New__AWS_ApplicationInsights_Application_CustomComponent() AWS_ApplicationInsights_Application_CustomComponent {
	return AWS_ApplicationInsights_Application_CustomComponent{}
}

// GetType returns the CloudFormation type.
func (AWS_ApplicationInsights_Application_CustomComponent) GetType() string {
	return AWS_ApplicationInsights_Application_CustomComponent__Type
}

// Set__ComponentName updates property "ComponentName".
func (t AWS_ApplicationInsights_Application_CustomComponent) Set__ComponentName(v cfz.Expression[string]) AWS_ApplicationInsights_Application_CustomComponent {
	t.ComponentName = v
	return t
}

// SetV__ComponentName updates property "ComponentName".
func (t AWS_ApplicationInsights_Application_CustomComponent) SetV__ComponentName(v string) AWS_ApplicationInsights_Application_CustomComponent {
	t.ComponentName = cfz.V(v)
	return t
}

// Set__ResourceList updates property "ResourceList".
func (t AWS_ApplicationInsights_Application_CustomComponent) Set__ResourceList(v cfz.ExpressionSlice[string]) AWS_ApplicationInsights_Application_CustomComponent {
	t.ResourceList = v
	return t
}

// SetS__ResourceList updates property "ResourceList".
func (t AWS_ApplicationInsights_Application_CustomComponent) SetS__ResourceList(v ...cfz.Expression[string]) AWS_ApplicationInsights_Application_CustomComponent {
	t.ResourceList = cfz.S(v...)
	return t
}

// SetSV__ResourceList updates property "ResourceList".
func (t AWS_ApplicationInsights_Application_CustomComponent) SetSV__ResourceList(v ...string) AWS_ApplicationInsights_Application_CustomComponent {
	t.ResourceList = cfz.SV(v...)
	return t
}
