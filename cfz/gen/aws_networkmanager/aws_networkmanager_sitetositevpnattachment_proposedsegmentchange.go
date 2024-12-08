// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__Type is the CloudFormation type for AWS::NetworkManager::SiteToSiteVpnAttachment.ProposedSegmentChange.
	AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__Type = "AWS::NetworkManager::SiteToSiteVpnAttachment.ProposedSegmentChange"
)

var (
	// AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::SiteToSiteVpnAttachment.ProposedSegmentChange.
	AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesMap = struct {
		AttachmentPolicyRuleNumber string
		SegmentName                string
		Tags                       string
	}{
		AttachmentPolicyRuleNumber: "AttachmentPolicyRuleNumber",
		SegmentName:                "SegmentName",
		Tags:                       "Tags",
	}

	// AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::SiteToSiteVpnAttachment.ProposedSegmentChange.
	AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesSlice = []string{
		AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesMap.AttachmentPolicyRuleNumber,
		AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesMap.SegmentName,
		AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__PropertiesMap.Tags,
	}
)

// AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange is a binding for AWS::NetworkManager::SiteToSiteVpnAttachment.ProposedSegmentChange.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-sitetositevpnattachment-proposedsegmentchange.html
type AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange struct {
	// AttachmentPolicyRuleNumber is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-sitetositevpnattachment-proposedsegmentchange.html#cfn-networkmanager-sitetositevpnattachment-proposedsegmentchange-attachmentpolicyrulenumber
	AttachmentPolicyRuleNumber cfz.Expression[int32] `json:"AttachmentPolicyRuleNumber,omitempty"`

	// SegmentName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-sitetositevpnattachment-proposedsegmentchange.html#cfn-networkmanager-sitetositevpnattachment-proposedsegmentchange-segmentname
	SegmentName cfz.Expression[string] `json:"SegmentName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-sitetositevpnattachment-proposedsegmentchange.html#cfn-networkmanager-sitetositevpnattachment-proposedsegmentchange-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange initializes a new AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange.
func New__AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange() AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	return AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange{}
}

// GetType returns the CloudFormation type.
func (AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) GetType() string {
	return AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange__Type
}

// Set__AttachmentPolicyRuleNumber updates property "AttachmentPolicyRuleNumber".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) Set__AttachmentPolicyRuleNumber(v cfz.Expression[int32]) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.AttachmentPolicyRuleNumber = v
	return t
}

// SetV__AttachmentPolicyRuleNumber updates property "AttachmentPolicyRuleNumber".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) SetV__AttachmentPolicyRuleNumber(v int32) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.AttachmentPolicyRuleNumber = cfz.V(v)
	return t
}

// Set__SegmentName updates property "SegmentName".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) Set__SegmentName(v cfz.Expression[string]) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.SegmentName = v
	return t
}

// SetV__SegmentName updates property "SegmentName".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) SetV__SegmentName(v string) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.SegmentName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) SetS__Tags(v ...cfz.Expression[cfz.Tag]) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange) SetSV__Tags(v ...cfz.Tag) AWS_NetworkManager_SiteToSiteVpnAttachment_ProposedSegmentChange {
	t.Tags = cfz.SV(v...)
	return t
}
