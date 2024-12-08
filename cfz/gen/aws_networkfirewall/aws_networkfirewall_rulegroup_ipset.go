// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkfirewall

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_NetworkFirewall_RuleGroup_IPSet__Type is the CloudFormation type for AWS::NetworkFirewall::RuleGroup.IPSet.
	AWS_NetworkFirewall_RuleGroup_IPSet__Type = "AWS::NetworkFirewall::RuleGroup.IPSet"
)

var (
	// AWS_NetworkFirewall_RuleGroup_IPSet__PropertiesMap reports all the CloudFormation properties for AWS::NetworkFirewall::RuleGroup.IPSet.
	AWS_NetworkFirewall_RuleGroup_IPSet__PropertiesMap = struct {
		Definition string
	}{
		Definition: "Definition",
	}

	// AWS_NetworkFirewall_RuleGroup_IPSet__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkFirewall::RuleGroup.IPSet.
	AWS_NetworkFirewall_RuleGroup_IPSet__PropertiesSlice = []string{
		AWS_NetworkFirewall_RuleGroup_IPSet__PropertiesMap.Definition,
	}
)

// AWS_NetworkFirewall_RuleGroup_IPSet is a binding for AWS::NetworkFirewall::RuleGroup.IPSet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ipset.html
type AWS_NetworkFirewall_RuleGroup_IPSet struct {
	// Definition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-ipset.html#cfn-networkfirewall-rulegroup-ipset-definition
	Definition cfz.ExpressionSlice[string] `json:"Definition,omitempty"`
}

// New__AWS_NetworkFirewall_RuleGroup_IPSet initializes a new AWS_NetworkFirewall_RuleGroup_IPSet.
func New__AWS_NetworkFirewall_RuleGroup_IPSet() AWS_NetworkFirewall_RuleGroup_IPSet {
	return AWS_NetworkFirewall_RuleGroup_IPSet{}
}

// GetType returns the CloudFormation type.
func (AWS_NetworkFirewall_RuleGroup_IPSet) GetType() string {
	return AWS_NetworkFirewall_RuleGroup_IPSet__Type
}

// Set__Definition updates property "Definition".
func (t AWS_NetworkFirewall_RuleGroup_IPSet) Set__Definition(v cfz.ExpressionSlice[string]) AWS_NetworkFirewall_RuleGroup_IPSet {
	t.Definition = v
	return t
}

// SetS__Definition updates property "Definition".
func (t AWS_NetworkFirewall_RuleGroup_IPSet) SetS__Definition(v ...cfz.Expression[string]) AWS_NetworkFirewall_RuleGroup_IPSet {
	t.Definition = cfz.S(v...)
	return t
}

// SetSV__Definition updates property "Definition".
func (t AWS_NetworkFirewall_RuleGroup_IPSet) SetSV__Definition(v ...string) AWS_NetworkFirewall_RuleGroup_IPSet {
	t.Definition = cfz.SV(v...)
	return t
}
