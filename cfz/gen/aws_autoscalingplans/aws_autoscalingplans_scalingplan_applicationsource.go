// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_autoscalingplans

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__Type is the CloudFormation type for AWS::AutoScalingPlans::ScalingPlan.ApplicationSource.
	AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__Type = "AWS::AutoScalingPlans::ScalingPlan.ApplicationSource"
)

var (
	// AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__PropertiesMap reports all the CloudFormation properties for AWS::AutoScalingPlans::ScalingPlan.ApplicationSource.
	AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__PropertiesMap = struct {
		CloudFormationStackARN string
		TagFilters             string
	}{
		CloudFormationStackARN: "CloudFormationStackARN",
		TagFilters:             "TagFilters",
	}

	// AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__PropertiesSlice reports all the CloudFormation properties for AWS::AutoScalingPlans::ScalingPlan.ApplicationSource.
	AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__PropertiesSlice = []string{
		AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__PropertiesMap.CloudFormationStackARN,
		AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__PropertiesMap.TagFilters,
	}
)

// AWS_AutoScalingPlans_ScalingPlan_ApplicationSource is a binding for AWS::AutoScalingPlans::ScalingPlan.ApplicationSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html
type AWS_AutoScalingPlans_ScalingPlan_ApplicationSource struct {
	// CloudFormationStackARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html#cfn-autoscalingplans-scalingplan-applicationsource-cloudformationstackarn
	CloudFormationStackARN cfz.Expression[string] `json:"CloudFormationStackARN,omitempty"`

	// TagFilters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html#cfn-autoscalingplans-scalingplan-applicationsource-tagfilters
	TagFilters cfz.ExpressionSlice[AWS_AutoScalingPlans_ScalingPlan_TagFilter] `json:"TagFilters,omitempty"`
}

// New__AWS_AutoScalingPlans_ScalingPlan_ApplicationSource initializes a new AWS_AutoScalingPlans_ScalingPlan_ApplicationSource.
func New__AWS_AutoScalingPlans_ScalingPlan_ApplicationSource() AWS_AutoScalingPlans_ScalingPlan_ApplicationSource {
	return AWS_AutoScalingPlans_ScalingPlan_ApplicationSource{}
}

// GetType returns the CloudFormation type.
func (AWS_AutoScalingPlans_ScalingPlan_ApplicationSource) GetType() string {
	return AWS_AutoScalingPlans_ScalingPlan_ApplicationSource__Type
}

// Set__CloudFormationStackARN updates property "CloudFormationStackARN".
func (t AWS_AutoScalingPlans_ScalingPlan_ApplicationSource) Set__CloudFormationStackARN(v cfz.Expression[string]) AWS_AutoScalingPlans_ScalingPlan_ApplicationSource {
	t.CloudFormationStackARN = v
	return t
}

// SetV__CloudFormationStackARN updates property "CloudFormationStackARN".
func (t AWS_AutoScalingPlans_ScalingPlan_ApplicationSource) SetV__CloudFormationStackARN(v string) AWS_AutoScalingPlans_ScalingPlan_ApplicationSource {
	t.CloudFormationStackARN = cfz.V(v)
	return t
}

// Set__TagFilters updates property "TagFilters".
func (t AWS_AutoScalingPlans_ScalingPlan_ApplicationSource) Set__TagFilters(v cfz.ExpressionSlice[AWS_AutoScalingPlans_ScalingPlan_TagFilter]) AWS_AutoScalingPlans_ScalingPlan_ApplicationSource {
	t.TagFilters = v
	return t
}

// SetS__TagFilters updates property "TagFilters".
func (t AWS_AutoScalingPlans_ScalingPlan_ApplicationSource) SetS__TagFilters(v ...cfz.Expression[AWS_AutoScalingPlans_ScalingPlan_TagFilter]) AWS_AutoScalingPlans_ScalingPlan_ApplicationSource {
	t.TagFilters = cfz.S(v...)
	return t
}

// SetSV__TagFilters updates property "TagFilters".
func (t AWS_AutoScalingPlans_ScalingPlan_ApplicationSource) SetSV__TagFilters(v ...AWS_AutoScalingPlans_ScalingPlan_TagFilter) AWS_AutoScalingPlans_ScalingPlan_ApplicationSource {
	t.TagFilters = cfz.SV(v...)
	return t
}
