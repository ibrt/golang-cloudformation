// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_RuleGroup_BlockAction__Type is the CloudFormation type for AWS::WAFv2::RuleGroup.BlockAction.
	AWS_WAFv2_RuleGroup_BlockAction__Type = "AWS::WAFv2::RuleGroup.BlockAction"
)

var (
	// AWS_WAFv2_RuleGroup_BlockAction__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::RuleGroup.BlockAction.
	AWS_WAFv2_RuleGroup_BlockAction__PropertiesMap = struct {
		CustomResponse string
	}{
		CustomResponse: "CustomResponse",
	}

	// AWS_WAFv2_RuleGroup_BlockAction__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::RuleGroup.BlockAction.
	AWS_WAFv2_RuleGroup_BlockAction__PropertiesSlice = []string{
		AWS_WAFv2_RuleGroup_BlockAction__PropertiesMap.CustomResponse,
	}
)

// AWS_WAFv2_RuleGroup_BlockAction is a binding for AWS::WAFv2::RuleGroup.BlockAction.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-blockaction.html
type AWS_WAFv2_RuleGroup_BlockAction struct {
	// CustomResponse is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-blockaction.html#cfn-wafv2-rulegroup-blockaction-customresponse
	CustomResponse cfz.Expression[AWS_WAFv2_RuleGroup_CustomResponse] `json:"CustomResponse,omitempty"`
}

// New__AWS_WAFv2_RuleGroup_BlockAction initializes a new AWS_WAFv2_RuleGroup_BlockAction.
func New__AWS_WAFv2_RuleGroup_BlockAction() AWS_WAFv2_RuleGroup_BlockAction {
	return AWS_WAFv2_RuleGroup_BlockAction{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_RuleGroup_BlockAction) GetType() string {
	return AWS_WAFv2_RuleGroup_BlockAction__Type
}

// Set__CustomResponse updates property "CustomResponse".
func (t AWS_WAFv2_RuleGroup_BlockAction) Set__CustomResponse(v cfz.Expression[AWS_WAFv2_RuleGroup_CustomResponse]) AWS_WAFv2_RuleGroup_BlockAction {
	t.CustomResponse = v
	return t
}

// SetV__CustomResponse updates property "CustomResponse".
func (t AWS_WAFv2_RuleGroup_BlockAction) SetV__CustomResponse(v AWS_WAFv2_RuleGroup_CustomResponse) AWS_WAFv2_RuleGroup_BlockAction {
	t.CustomResponse = cfz.V(v)
	return t
}
