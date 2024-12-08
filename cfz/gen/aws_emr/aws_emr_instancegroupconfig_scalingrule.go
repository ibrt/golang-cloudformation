// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_InstanceGroupConfig_ScalingRule__Type is the CloudFormation type for AWS::EMR::InstanceGroupConfig.ScalingRule.
	AWS_EMR_InstanceGroupConfig_ScalingRule__Type = "AWS::EMR::InstanceGroupConfig.ScalingRule"
)

var (
	// AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesMap reports all the CloudFormation properties for AWS::EMR::InstanceGroupConfig.ScalingRule.
	AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesMap = struct {
		Action      string
		Description string
		Name        string
		Trigger     string
	}{
		Action:      "Action",
		Description: "Description",
		Name:        "Name",
		Trigger:     "Trigger",
	}

	// AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::InstanceGroupConfig.ScalingRule.
	AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesSlice = []string{
		AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesMap.Action,
		AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesMap.Description,
		AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesMap.Name,
		AWS_EMR_InstanceGroupConfig_ScalingRule__PropertiesMap.Trigger,
	}
)

// AWS_EMR_InstanceGroupConfig_ScalingRule is a binding for AWS::EMR::InstanceGroupConfig.ScalingRule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html
type AWS_EMR_InstanceGroupConfig_ScalingRule struct {
	// Action is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-action
	Action cfz.Expression[AWS_EMR_InstanceGroupConfig_ScalingAction] `json:"Action,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Trigger is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-trigger
	Trigger cfz.Expression[AWS_EMR_InstanceGroupConfig_ScalingTrigger] `json:"Trigger,omitempty"`
}

// New__AWS_EMR_InstanceGroupConfig_ScalingRule initializes a new AWS_EMR_InstanceGroupConfig_ScalingRule.
func New__AWS_EMR_InstanceGroupConfig_ScalingRule() AWS_EMR_InstanceGroupConfig_ScalingRule {
	return AWS_EMR_InstanceGroupConfig_ScalingRule{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_InstanceGroupConfig_ScalingRule) GetType() string {
	return AWS_EMR_InstanceGroupConfig_ScalingRule__Type
}

// Set__Action updates property "Action".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) Set__Action(v cfz.Expression[AWS_EMR_InstanceGroupConfig_ScalingAction]) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Action = v
	return t
}

// SetV__Action updates property "Action".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) SetV__Action(v AWS_EMR_InstanceGroupConfig_ScalingAction) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Action = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) Set__Description(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) SetV__Description(v string) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) Set__Name(v cfz.Expression[string]) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) SetV__Name(v string) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Name = cfz.V(v)
	return t
}

// Set__Trigger updates property "Trigger".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) Set__Trigger(v cfz.Expression[AWS_EMR_InstanceGroupConfig_ScalingTrigger]) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Trigger = v
	return t
}

// SetV__Trigger updates property "Trigger".
func (t AWS_EMR_InstanceGroupConfig_ScalingRule) SetV__Trigger(v AWS_EMR_InstanceGroupConfig_ScalingTrigger) AWS_EMR_InstanceGroupConfig_ScalingRule {
	t.Trigger = cfz.V(v)
	return t
}
