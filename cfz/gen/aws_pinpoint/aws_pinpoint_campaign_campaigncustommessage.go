// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpoint

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pinpoint_Campaign_CampaignCustomMessage__Type is the CloudFormation type for AWS::Pinpoint::Campaign.CampaignCustomMessage.
	AWS_Pinpoint_Campaign_CampaignCustomMessage__Type = "AWS::Pinpoint::Campaign.CampaignCustomMessage"
)

var (
	// AWS_Pinpoint_Campaign_CampaignCustomMessage__PropertiesMap reports all the CloudFormation properties for AWS::Pinpoint::Campaign.CampaignCustomMessage.
	AWS_Pinpoint_Campaign_CampaignCustomMessage__PropertiesMap = struct {
		Data string
	}{
		Data: "Data",
	}

	// AWS_Pinpoint_Campaign_CampaignCustomMessage__PropertiesSlice reports all the CloudFormation properties for AWS::Pinpoint::Campaign.CampaignCustomMessage.
	AWS_Pinpoint_Campaign_CampaignCustomMessage__PropertiesSlice = []string{
		AWS_Pinpoint_Campaign_CampaignCustomMessage__PropertiesMap.Data,
	}
)

// AWS_Pinpoint_Campaign_CampaignCustomMessage is a binding for AWS::Pinpoint::Campaign.CampaignCustomMessage.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigncustommessage.html
type AWS_Pinpoint_Campaign_CampaignCustomMessage struct {
	// Data is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigncustommessage.html#cfn-pinpoint-campaign-campaigncustommessage-data
	Data cfz.Expression[string] `json:"Data,omitempty"`
}

// New__AWS_Pinpoint_Campaign_CampaignCustomMessage initializes a new AWS_Pinpoint_Campaign_CampaignCustomMessage.
func New__AWS_Pinpoint_Campaign_CampaignCustomMessage() AWS_Pinpoint_Campaign_CampaignCustomMessage {
	return AWS_Pinpoint_Campaign_CampaignCustomMessage{}
}

// GetType returns the CloudFormation type.
func (AWS_Pinpoint_Campaign_CampaignCustomMessage) GetType() string {
	return AWS_Pinpoint_Campaign_CampaignCustomMessage__Type
}

// Set__Data updates property "Data".
func (t AWS_Pinpoint_Campaign_CampaignCustomMessage) Set__Data(v cfz.Expression[string]) AWS_Pinpoint_Campaign_CampaignCustomMessage {
	t.Data = v
	return t
}

// SetV__Data updates property "Data".
func (t AWS_Pinpoint_Campaign_CampaignCustomMessage) SetV__Data(v string) AWS_Pinpoint_Campaign_CampaignCustomMessage {
	t.Data = cfz.V(v)
	return t
}
