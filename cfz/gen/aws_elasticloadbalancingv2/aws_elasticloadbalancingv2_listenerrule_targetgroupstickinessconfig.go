// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticloadbalancingv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__Type is the CloudFormation type for AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupStickinessConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__Type = "AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupStickinessConfig"
)

var (
	// AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__PropertiesMap reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupStickinessConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__PropertiesMap = struct {
		DurationSeconds string
		Enabled         string
	}{
		DurationSeconds: "DurationSeconds",
		Enabled:         "Enabled",
	}

	// AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupStickinessConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__PropertiesSlice = []string{
		AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__PropertiesMap.DurationSeconds,
		AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__PropertiesMap.Enabled,
	}
)

// AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig is a binding for AWS::ElasticLoadBalancingV2::ListenerRule.TargetGroupStickinessConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-targetgroupstickinessconfig.html
type AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig struct {
	// DurationSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-targetgroupstickinessconfig.html#cfn-elasticloadbalancingv2-listenerrule-targetgroupstickinessconfig-durationseconds
	DurationSeconds cfz.Expression[int32] `json:"DurationSeconds,omitempty"`

	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-targetgroupstickinessconfig.html#cfn-elasticloadbalancingv2-listenerrule-targetgroupstickinessconfig-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig initializes a new AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig.
func New__AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig() AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig {
	return AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig) GetType() string {
	return AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig__Type
}

// Set__DurationSeconds updates property "DurationSeconds".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig) Set__DurationSeconds(v cfz.Expression[int32]) AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig {
	t.DurationSeconds = v
	return t
}

// SetV__DurationSeconds updates property "DurationSeconds".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig) SetV__DurationSeconds(v int32) AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig {
	t.DurationSeconds = cfz.V(v)
	return t
}

// Set__Enabled updates property "Enabled".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig) Set__Enabled(v cfz.Expression[bool]) AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig) SetV__Enabled(v bool) AWS_ElasticLoadBalancingV2_ListenerRule_TargetGroupStickinessConfig {
	t.Enabled = cfz.V(v)
	return t
}
