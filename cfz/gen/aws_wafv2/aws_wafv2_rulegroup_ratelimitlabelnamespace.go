// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__Type is the CloudFormation type for AWS::WAFv2::RuleGroup.RateLimitLabelNamespace.
	AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__Type = "AWS::WAFv2::RuleGroup.RateLimitLabelNamespace"
)

var (
	// AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::RuleGroup.RateLimitLabelNamespace.
	AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__PropertiesMap = struct {
		Namespace string
	}{
		Namespace: "Namespace",
	}

	// AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::RuleGroup.RateLimitLabelNamespace.
	AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__PropertiesSlice = []string{
		AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__PropertiesMap.Namespace,
	}
)

// AWS_WAFv2_RuleGroup_RateLimitLabelNamespace is a binding for AWS::WAFv2::RuleGroup.RateLimitLabelNamespace.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitlabelnamespace.html
type AWS_WAFv2_RuleGroup_RateLimitLabelNamespace struct {
	// Namespace is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-ratelimitlabelnamespace.html#cfn-wafv2-rulegroup-ratelimitlabelnamespace-namespace
	Namespace cfz.Expression[string] `json:"Namespace,omitempty"`
}

// New__AWS_WAFv2_RuleGroup_RateLimitLabelNamespace initializes a new AWS_WAFv2_RuleGroup_RateLimitLabelNamespace.
func New__AWS_WAFv2_RuleGroup_RateLimitLabelNamespace() AWS_WAFv2_RuleGroup_RateLimitLabelNamespace {
	return AWS_WAFv2_RuleGroup_RateLimitLabelNamespace{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_RuleGroup_RateLimitLabelNamespace) GetType() string {
	return AWS_WAFv2_RuleGroup_RateLimitLabelNamespace__Type
}

// Set__Namespace updates property "Namespace".
func (t AWS_WAFv2_RuleGroup_RateLimitLabelNamespace) Set__Namespace(v cfz.Expression[string]) AWS_WAFv2_RuleGroup_RateLimitLabelNamespace {
	t.Namespace = v
	return t
}

// SetV__Namespace updates property "Namespace".
func (t AWS_WAFv2_RuleGroup_RateLimitLabelNamespace) SetV__Namespace(v string) AWS_WAFv2_RuleGroup_RateLimitLabelNamespace {
	t.Namespace = cfz.V(v)
	return t
}
