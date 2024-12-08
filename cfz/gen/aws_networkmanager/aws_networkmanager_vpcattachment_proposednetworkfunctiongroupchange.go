// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__Type is the CloudFormation type for AWS::NetworkManager::VpcAttachment.ProposedNetworkFunctionGroupChange.
	AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__Type = "AWS::NetworkManager::VpcAttachment.ProposedNetworkFunctionGroupChange"
)

var (
	// AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::VpcAttachment.ProposedNetworkFunctionGroupChange.
	AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesMap = struct {
		AttachmentPolicyRuleNumber string
		NetworkFunctionGroupName   string
		Tags                       string
	}{
		AttachmentPolicyRuleNumber: "AttachmentPolicyRuleNumber",
		NetworkFunctionGroupName:   "NetworkFunctionGroupName",
		Tags:                       "Tags",
	}

	// AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::VpcAttachment.ProposedNetworkFunctionGroupChange.
	AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesSlice = []string{
		AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesMap.AttachmentPolicyRuleNumber,
		AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesMap.NetworkFunctionGroupName,
		AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__PropertiesMap.Tags,
	}
)

// AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange is a binding for AWS::NetworkManager::VpcAttachment.ProposedNetworkFunctionGroupChange.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposednetworkfunctiongroupchange.html
type AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange struct {
	// AttachmentPolicyRuleNumber is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposednetworkfunctiongroupchange.html#cfn-networkmanager-vpcattachment-proposednetworkfunctiongroupchange-attachmentpolicyrulenumber
	AttachmentPolicyRuleNumber cfz.Expression[int32] `json:"AttachmentPolicyRuleNumber,omitempty"`

	// NetworkFunctionGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposednetworkfunctiongroupchange.html#cfn-networkmanager-vpcattachment-proposednetworkfunctiongroupchange-networkfunctiongroupname
	NetworkFunctionGroupName cfz.Expression[string] `json:"NetworkFunctionGroupName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposednetworkfunctiongroupchange.html#cfn-networkmanager-vpcattachment-proposednetworkfunctiongroupchange-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange initializes a new AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange.
func New__AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange() AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	return AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange{}
}

// GetType returns the CloudFormation type.
func (AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) GetType() string {
	return AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange__Type
}

// Set__AttachmentPolicyRuleNumber updates property "AttachmentPolicyRuleNumber".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) Set__AttachmentPolicyRuleNumber(v cfz.Expression[int32]) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.AttachmentPolicyRuleNumber = v
	return t
}

// SetV__AttachmentPolicyRuleNumber updates property "AttachmentPolicyRuleNumber".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) SetV__AttachmentPolicyRuleNumber(v int32) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.AttachmentPolicyRuleNumber = cfz.V(v)
	return t
}

// Set__NetworkFunctionGroupName updates property "NetworkFunctionGroupName".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) Set__NetworkFunctionGroupName(v cfz.Expression[string]) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.NetworkFunctionGroupName = v
	return t
}

// SetV__NetworkFunctionGroupName updates property "NetworkFunctionGroupName".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) SetV__NetworkFunctionGroupName(v string) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.NetworkFunctionGroupName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) SetS__Tags(v ...cfz.Expression[cfz.Tag]) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange) SetSV__Tags(v ...cfz.Tag) AWS_NetworkManager_VpcAttachment_ProposedNetworkFunctionGroupChange {
	t.Tags = cfz.SV(v...)
	return t
}
