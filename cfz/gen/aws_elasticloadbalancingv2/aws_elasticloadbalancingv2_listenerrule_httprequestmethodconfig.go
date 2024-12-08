// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticloadbalancingv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__Type is the CloudFormation type for AWS::ElasticLoadBalancingV2::ListenerRule.HttpRequestMethodConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__Type = "AWS::ElasticLoadBalancingV2::ListenerRule.HttpRequestMethodConfig"
)

var (
	// AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__PropertiesMap reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::ListenerRule.HttpRequestMethodConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__PropertiesMap = struct {
		Values string
	}{
		Values: "Values",
	}

	// AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__PropertiesSlice reports all the CloudFormation properties for AWS::ElasticLoadBalancingV2::ListenerRule.HttpRequestMethodConfig.
	AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__PropertiesSlice = []string{
		AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__PropertiesMap.Values,
	}
)

// AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig is a binding for AWS::ElasticLoadBalancingV2::ListenerRule.HttpRequestMethodConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-httprequestmethodconfig.html
type AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig struct {
	// Values is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-httprequestmethodconfig.html#cfn-elasticloadbalancingv2-listenerrule-httprequestmethodconfig-values
	Values cfz.ExpressionSlice[string] `json:"Values,omitempty"`
}

// New__AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig initializes a new AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig.
func New__AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig() AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig {
	return AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig) GetType() string {
	return AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig__Type
}

// Set__Values updates property "Values".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig) Set__Values(v cfz.ExpressionSlice[string]) AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig {
	t.Values = v
	return t
}

// SetS__Values updates property "Values".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig) SetS__Values(v ...cfz.Expression[string]) AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig {
	t.Values = cfz.S(v...)
	return t
}

// SetSV__Values updates property "Values".
func (t AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig) SetSV__Values(v ...string) AWS_ElasticLoadBalancingV2_ListenerRule_HttpRequestMethodConfig {
	t.Values = cfz.SV(v...)
	return t
}
