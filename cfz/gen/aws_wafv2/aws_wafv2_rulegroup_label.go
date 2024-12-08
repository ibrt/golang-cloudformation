// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_RuleGroup_Label__Type is the CloudFormation type for AWS::WAFv2::RuleGroup.Label.
	AWS_WAFv2_RuleGroup_Label__Type = "AWS::WAFv2::RuleGroup.Label"
)

var (
	// AWS_WAFv2_RuleGroup_Label__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::RuleGroup.Label.
	AWS_WAFv2_RuleGroup_Label__PropertiesMap = struct {
		Name string
	}{
		Name: "Name",
	}

	// AWS_WAFv2_RuleGroup_Label__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::RuleGroup.Label.
	AWS_WAFv2_RuleGroup_Label__PropertiesSlice = []string{
		AWS_WAFv2_RuleGroup_Label__PropertiesMap.Name,
	}
)

// AWS_WAFv2_RuleGroup_Label is a binding for AWS::WAFv2::RuleGroup.Label.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-label.html
type AWS_WAFv2_RuleGroup_Label struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-label.html#cfn-wafv2-rulegroup-label-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_WAFv2_RuleGroup_Label initializes a new AWS_WAFv2_RuleGroup_Label.
func New__AWS_WAFv2_RuleGroup_Label() AWS_WAFv2_RuleGroup_Label {
	return AWS_WAFv2_RuleGroup_Label{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_RuleGroup_Label) GetType() string {
	return AWS_WAFv2_RuleGroup_Label__Type
}

// Set__Name updates property "Name".
func (t AWS_WAFv2_RuleGroup_Label) Set__Name(v cfz.Expression[string]) AWS_WAFv2_RuleGroup_Label {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_WAFv2_RuleGroup_Label) SetV__Name(v string) AWS_WAFv2_RuleGroup_Label {
	t.Name = cfz.V(v)
	return t
}
