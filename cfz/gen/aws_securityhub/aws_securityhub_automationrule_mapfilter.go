// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_securityhub

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SecurityHub_AutomationRule_MapFilter__Type is the CloudFormation type for AWS::SecurityHub::AutomationRule.MapFilter.
	AWS_SecurityHub_AutomationRule_MapFilter__Type = "AWS::SecurityHub::AutomationRule.MapFilter"
)

var (
	// AWS_SecurityHub_AutomationRule_MapFilter__PropertiesMap reports all the CloudFormation properties for AWS::SecurityHub::AutomationRule.MapFilter.
	AWS_SecurityHub_AutomationRule_MapFilter__PropertiesMap = struct {
		Comparison string
		Key        string
		Value      string
	}{
		Comparison: "Comparison",
		Key:        "Key",
		Value:      "Value",
	}

	// AWS_SecurityHub_AutomationRule_MapFilter__PropertiesSlice reports all the CloudFormation properties for AWS::SecurityHub::AutomationRule.MapFilter.
	AWS_SecurityHub_AutomationRule_MapFilter__PropertiesSlice = []string{
		AWS_SecurityHub_AutomationRule_MapFilter__PropertiesMap.Comparison,
		AWS_SecurityHub_AutomationRule_MapFilter__PropertiesMap.Key,
		AWS_SecurityHub_AutomationRule_MapFilter__PropertiesMap.Value,
	}
)

// AWS_SecurityHub_AutomationRule_MapFilter is a binding for AWS::SecurityHub::AutomationRule.MapFilter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html
type AWS_SecurityHub_AutomationRule_MapFilter struct {
	// Comparison is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html#cfn-securityhub-automationrule-mapfilter-comparison
	Comparison cfz.Expression[string] `json:"Comparison,omitempty"`

	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html#cfn-securityhub-automationrule-mapfilter-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-mapfilter.html#cfn-securityhub-automationrule-mapfilter-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_SecurityHub_AutomationRule_MapFilter initializes a new AWS_SecurityHub_AutomationRule_MapFilter.
func New__AWS_SecurityHub_AutomationRule_MapFilter() AWS_SecurityHub_AutomationRule_MapFilter {
	return AWS_SecurityHub_AutomationRule_MapFilter{}
}

// GetType returns the CloudFormation type.
func (AWS_SecurityHub_AutomationRule_MapFilter) GetType() string {
	return AWS_SecurityHub_AutomationRule_MapFilter__Type
}

// Set__Comparison updates property "Comparison".
func (t AWS_SecurityHub_AutomationRule_MapFilter) Set__Comparison(v cfz.Expression[string]) AWS_SecurityHub_AutomationRule_MapFilter {
	t.Comparison = v
	return t
}

// SetV__Comparison updates property "Comparison".
func (t AWS_SecurityHub_AutomationRule_MapFilter) SetV__Comparison(v string) AWS_SecurityHub_AutomationRule_MapFilter {
	t.Comparison = cfz.V(v)
	return t
}

// Set__Key updates property "Key".
func (t AWS_SecurityHub_AutomationRule_MapFilter) Set__Key(v cfz.Expression[string]) AWS_SecurityHub_AutomationRule_MapFilter {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_SecurityHub_AutomationRule_MapFilter) SetV__Key(v string) AWS_SecurityHub_AutomationRule_MapFilter {
	t.Key = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_SecurityHub_AutomationRule_MapFilter) Set__Value(v cfz.Expression[string]) AWS_SecurityHub_AutomationRule_MapFilter {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_SecurityHub_AutomationRule_MapFilter) SetV__Value(v string) AWS_SecurityHub_AutomationRule_MapFilter {
	t.Value = cfz.V(v)
	return t
}
