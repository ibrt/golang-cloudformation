// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudformation

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFormation_GuardHook_StackNames__Type is the CloudFormation type for AWS::CloudFormation::GuardHook.StackNames.
	AWS_CloudFormation_GuardHook_StackNames__Type = "AWS::CloudFormation::GuardHook.StackNames"
)

var (
	// AWS_CloudFormation_GuardHook_StackNames__PropertiesMap reports all the CloudFormation properties for AWS::CloudFormation::GuardHook.StackNames.
	AWS_CloudFormation_GuardHook_StackNames__PropertiesMap = struct {
		Exclude string
		Include string
	}{
		Exclude: "Exclude",
		Include: "Include",
	}

	// AWS_CloudFormation_GuardHook_StackNames__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFormation::GuardHook.StackNames.
	AWS_CloudFormation_GuardHook_StackNames__PropertiesSlice = []string{
		AWS_CloudFormation_GuardHook_StackNames__PropertiesMap.Exclude,
		AWS_CloudFormation_GuardHook_StackNames__PropertiesMap.Include,
	}
)

// AWS_CloudFormation_GuardHook_StackNames is a binding for AWS::CloudFormation::GuardHook.StackNames.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stacknames.html
type AWS_CloudFormation_GuardHook_StackNames struct {
	// Exclude is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stacknames.html#cfn-cloudformation-guardhook-stacknames-exclude
	Exclude cfz.ExpressionSlice[string] `json:"Exclude,omitempty"`

	// Include is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stacknames.html#cfn-cloudformation-guardhook-stacknames-include
	Include cfz.ExpressionSlice[string] `json:"Include,omitempty"`
}

// New__AWS_CloudFormation_GuardHook_StackNames initializes a new AWS_CloudFormation_GuardHook_StackNames.
func New__AWS_CloudFormation_GuardHook_StackNames() AWS_CloudFormation_GuardHook_StackNames {
	return AWS_CloudFormation_GuardHook_StackNames{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFormation_GuardHook_StackNames) GetType() string {
	return AWS_CloudFormation_GuardHook_StackNames__Type
}

// Set__Exclude updates property "Exclude".
func (t AWS_CloudFormation_GuardHook_StackNames) Set__Exclude(v cfz.ExpressionSlice[string]) AWS_CloudFormation_GuardHook_StackNames {
	t.Exclude = v
	return t
}

// SetS__Exclude updates property "Exclude".
func (t AWS_CloudFormation_GuardHook_StackNames) SetS__Exclude(v ...cfz.Expression[string]) AWS_CloudFormation_GuardHook_StackNames {
	t.Exclude = cfz.S(v...)
	return t
}

// SetSV__Exclude updates property "Exclude".
func (t AWS_CloudFormation_GuardHook_StackNames) SetSV__Exclude(v ...string) AWS_CloudFormation_GuardHook_StackNames {
	t.Exclude = cfz.SV(v...)
	return t
}

// Set__Include updates property "Include".
func (t AWS_CloudFormation_GuardHook_StackNames) Set__Include(v cfz.ExpressionSlice[string]) AWS_CloudFormation_GuardHook_StackNames {
	t.Include = v
	return t
}

// SetS__Include updates property "Include".
func (t AWS_CloudFormation_GuardHook_StackNames) SetS__Include(v ...cfz.Expression[string]) AWS_CloudFormation_GuardHook_StackNames {
	t.Include = cfz.S(v...)
	return t
}

// SetSV__Include updates property "Include".
func (t AWS_CloudFormation_GuardHook_StackNames) SetSV__Include(v ...string) AWS_CloudFormation_GuardHook_StackNames {
	t.Include = cfz.SV(v...)
	return t
}
