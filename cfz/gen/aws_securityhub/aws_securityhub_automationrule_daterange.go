// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_securityhub

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SecurityHub_AutomationRule_DateRange__Type is the CloudFormation type for AWS::SecurityHub::AutomationRule.DateRange.
	AWS_SecurityHub_AutomationRule_DateRange__Type = "AWS::SecurityHub::AutomationRule.DateRange"
)

var (
	// AWS_SecurityHub_AutomationRule_DateRange__PropertiesMap reports all the CloudFormation properties for AWS::SecurityHub::AutomationRule.DateRange.
	AWS_SecurityHub_AutomationRule_DateRange__PropertiesMap = struct {
		Unit  string
		Value string
	}{
		Unit:  "Unit",
		Value: "Value",
	}

	// AWS_SecurityHub_AutomationRule_DateRange__PropertiesSlice reports all the CloudFormation properties for AWS::SecurityHub::AutomationRule.DateRange.
	AWS_SecurityHub_AutomationRule_DateRange__PropertiesSlice = []string{
		AWS_SecurityHub_AutomationRule_DateRange__PropertiesMap.Unit,
		AWS_SecurityHub_AutomationRule_DateRange__PropertiesMap.Value,
	}
)

// AWS_SecurityHub_AutomationRule_DateRange is a binding for AWS::SecurityHub::AutomationRule.DateRange.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-daterange.html
type AWS_SecurityHub_AutomationRule_DateRange struct {
	// Unit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-daterange.html#cfn-securityhub-automationrule-daterange-unit
	Unit cfz.Expression[string] `json:"Unit,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-daterange.html#cfn-securityhub-automationrule-daterange-value
	Value cfz.Expression[float64] `json:"Value,omitempty"`
}

// New__AWS_SecurityHub_AutomationRule_DateRange initializes a new AWS_SecurityHub_AutomationRule_DateRange.
func New__AWS_SecurityHub_AutomationRule_DateRange() AWS_SecurityHub_AutomationRule_DateRange {
	return AWS_SecurityHub_AutomationRule_DateRange{}
}

// GetType returns the CloudFormation type.
func (AWS_SecurityHub_AutomationRule_DateRange) GetType() string {
	return AWS_SecurityHub_AutomationRule_DateRange__Type
}

// Set__Unit updates property "Unit".
func (t AWS_SecurityHub_AutomationRule_DateRange) Set__Unit(v cfz.Expression[string]) AWS_SecurityHub_AutomationRule_DateRange {
	t.Unit = v
	return t
}

// SetV__Unit updates property "Unit".
func (t AWS_SecurityHub_AutomationRule_DateRange) SetV__Unit(v string) AWS_SecurityHub_AutomationRule_DateRange {
	t.Unit = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_SecurityHub_AutomationRule_DateRange) Set__Value(v cfz.Expression[float64]) AWS_SecurityHub_AutomationRule_DateRange {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_SecurityHub_AutomationRule_DateRange) SetV__Value(v float64) AWS_SecurityHub_AutomationRule_DateRange {
	t.Value = cfz.V(v)
	return t
}
