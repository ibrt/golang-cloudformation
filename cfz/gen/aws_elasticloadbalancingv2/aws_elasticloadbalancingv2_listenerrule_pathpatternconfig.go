// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticloadbalancingv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__Type is the CloudFormation type for AWS::ElasticLoadBalancingV2::ListenerRule.PathPatternConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__Type = "AWS::ElasticLoadBalancingV2::ListenerRule.PathPatternConfig"
)

var (
	// AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__PropertiesMap reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::ListenerRule.PathPatternConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__PropertiesMap = struct {
		Values string
	}{
		Values: "Values",
	}

	// AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::ListenerRule.PathPatternConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__PropertiesSlice = []string{
		AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__PropertiesMap.Values,
	}
)

// AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig is a binding for AWS::ElasticLoadBalancingV2::ListenerRule.PathPatternConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-pathpatternconfig.html
type AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig struct {
	// Values is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-pathpatternconfig.html#cfn-elasticloadbalancingv2-listenerrule-pathpatternconfig-values
	Values cfz.ExpressionSlice[string] `json:"Values,omitempty"`
}

// New__AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig initializes a new AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig.
func New__AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig() AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig {
	return AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig) GetType() string {
	return AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig__Type
}

// Set__Values updates property "Values".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig) Set__Values(v cfz.ExpressionSlice[string]) AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig {
	t.Values = v
	return t
}

// SetS__Values updates property "Values".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig) SetS__Values(v ...cfz.Expression[string]) AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig {
	t.Values = cfz.S(v...)
	return t
}

// SetSV__Values updates property "Values".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig) SetSV__Values(v ...string) AWS_ElasticLoadBalancingV2_ListenerRule_PathPatternConfig {
	t.Values = cfz.SV(v...)
	return t
}
