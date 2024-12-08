// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_autoscaling

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__Type is the CloudFormation type for AWS::AutoScaling::AutoScalingGroup.LaunchTemplate.
	AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__Type = "AWS::AutoScaling::AutoScalingGroup.LaunchTemplate"
)

var (
	// AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__PropertiesMap reports all the CloudFormation properties for AWS::AutoScaling::AutoScalingGroup.LaunchTemplate.
	AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__PropertiesMap = struct {
		LaunchTemplateSpecification string
		Overrides                   string
	}{
		LaunchTemplateSpecification: "LaunchTemplateSpecification",
		Overrides:                   "Overrides",
	}

	// AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__PropertiesSlice reports all the CloudFormation properties for AWS::AutoScaling::AutoScalingGroup.LaunchTemplate.
	AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__PropertiesSlice = []string{
		AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__PropertiesMap.LaunchTemplateSpecification,
		AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__PropertiesMap.Overrides,
	}
)

// AWS_AutoScaling_AutoScalingGroup_LaunchTemplate is a binding for AWS::AutoScaling::AutoScalingGroup.LaunchTemplate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html
type AWS_AutoScaling_AutoScalingGroup_LaunchTemplate struct {
	// LaunchTemplateSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html#cfn-autoscaling-autoscalinggroup-launchtemplate-launchtemplatespecification
	LaunchTemplateSpecification cfz.Expression[AWS_AutoScaling_AutoScalingGroup_LaunchTemplateSpecification] `json:"LaunchTemplateSpecification,omitempty"`

	// Overrides is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.html#cfn-autoscaling-autoscalinggroup-launchtemplate-overrides
	Overrides cfz.ExpressionSlice[AWS_AutoScaling_AutoScalingGroup_LaunchTemplateOverrides] `json:"Overrides,omitempty"`
}

// New__AWS_AutoScaling_AutoScalingGroup_LaunchTemplate initializes a new AWS_AutoScaling_AutoScalingGroup_LaunchTemplate.
func New__AWS_AutoScaling_AutoScalingGroup_LaunchTemplate() AWS_AutoScaling_AutoScalingGroup_LaunchTemplate {
	return AWS_AutoScaling_AutoScalingGroup_LaunchTemplate{}
}

// GetType returns the CloudFormation type.
func (AWS_AutoScaling_AutoScalingGroup_LaunchTemplate) GetType() string {
	return AWS_AutoScaling_AutoScalingGroup_LaunchTemplate__Type
}

// Set__LaunchTemplateSpecification updates property "LaunchTemplateSpecification".
func (t AWS_AutoScaling_AutoScalingGroup_LaunchTemplate) Set__LaunchTemplateSpecification(v cfz.Expression[AWS_AutoScaling_AutoScalingGroup_LaunchTemplateSpecification]) AWS_AutoScaling_AutoScalingGroup_LaunchTemplate {
	t.LaunchTemplateSpecification = v
	return t
}

// SetV__LaunchTemplateSpecification updates property "LaunchTemplateSpecification".
func (t AWS_AutoScaling_AutoScalingGroup_LaunchTemplate) SetV__LaunchTemplateSpecification(v AWS_AutoScaling_AutoScalingGroup_LaunchTemplateSpecification) AWS_AutoScaling_AutoScalingGroup_LaunchTemplate {
	t.LaunchTemplateSpecification = cfz.V(v)
	return t
}

// Set__Overrides updates property "Overrides".
func (t AWS_AutoScaling_AutoScalingGroup_LaunchTemplate) Set__Overrides(v cfz.ExpressionSlice[AWS_AutoScaling_AutoScalingGroup_LaunchTemplateOverrides]) AWS_AutoScaling_AutoScalingGroup_LaunchTemplate {
	t.Overrides = v
	return t
}

// SetS__Overrides updates property "Overrides".
func (t AWS_AutoScaling_AutoScalingGroup_LaunchTemplate) SetS__Overrides(v ...cfz.Expression[AWS_AutoScaling_AutoScalingGroup_LaunchTemplateOverrides]) AWS_AutoScaling_AutoScalingGroup_LaunchTemplate {
	t.Overrides = cfz.S(v...)
	return t
}

// SetSV__Overrides updates property "Overrides".
func (t AWS_AutoScaling_AutoScalingGroup_LaunchTemplate) SetSV__Overrides(v ...AWS_AutoScaling_AutoScalingGroup_LaunchTemplateOverrides) AWS_AutoScaling_AutoScalingGroup_LaunchTemplate {
	t.Overrides = cfz.SV(v...)
	return t
}
