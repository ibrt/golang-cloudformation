// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pinpoint

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pinpoint_Campaign_CampaignEventFilter__Type is the CloudFormation type for AWS::Pinpoint::Campaign.CampaignEventFilter.
	AWS_Pinpoint_Campaign_CampaignEventFilter__Type = "AWS::Pinpoint::Campaign.CampaignEventFilter"
)

var (
	// AWS_Pinpoint_Campaign_CampaignEventFilter__PropertiesMap reports all the CloudFormation properties for AWS::Pinpoint::Campaign.CampaignEventFilter.
	AWS_Pinpoint_Campaign_CampaignEventFilter__PropertiesMap = struct {
		Dimensions string
		FilterType string
	}{
		Dimensions: "Dimensions",
		FilterType: "FilterType",
	}

	// AWS_Pinpoint_Campaign_CampaignEventFilter__PropertiesSlice reports all the CloudFormation properties for AWS::Pinpoint::Campaign.CampaignEventFilter.
	AWS_Pinpoint_Campaign_CampaignEventFilter__PropertiesSlice = []string{
		AWS_Pinpoint_Campaign_CampaignEventFilter__PropertiesMap.Dimensions,
		AWS_Pinpoint_Campaign_CampaignEventFilter__PropertiesMap.FilterType,
	}
)

// AWS_Pinpoint_Campaign_CampaignEventFilter is a binding for AWS::Pinpoint::Campaign.CampaignEventFilter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigneventfilter.html
type AWS_Pinpoint_Campaign_CampaignEventFilter struct {
	// Dimensions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigneventfilter.html#cfn-pinpoint-campaign-campaigneventfilter-dimensions
	Dimensions cfz.Expression[AWS_Pinpoint_Campaign_EventDimensions] `json:"Dimensions,omitempty"`

	// FilterType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-campaigneventfilter.html#cfn-pinpoint-campaign-campaigneventfilter-filtertype
	FilterType cfz.Expression[string] `json:"FilterType,omitempty"`
}

// New__AWS_Pinpoint_Campaign_CampaignEventFilter initializes a new AWS_Pinpoint_Campaign_CampaignEventFilter.
func New__AWS_Pinpoint_Campaign_CampaignEventFilter() AWS_Pinpoint_Campaign_CampaignEventFilter {
	return AWS_Pinpoint_Campaign_CampaignEventFilter{}
}

// GetType returns the CloudFormation type.
func (AWS_Pinpoint_Campaign_CampaignEventFilter) GetType() string {
	return AWS_Pinpoint_Campaign_CampaignEventFilter__Type
}

// Set__Dimensions updates property "Dimensions".
func (t AWS_Pinpoint_Campaign_CampaignEventFilter) Set__Dimensions(v cfz.Expression[AWS_Pinpoint_Campaign_EventDimensions]) AWS_Pinpoint_Campaign_CampaignEventFilter {
	t.Dimensions = v
	return t
}

// SetV__Dimensions updates property "Dimensions".
func (t AWS_Pinpoint_Campaign_CampaignEventFilter) SetV__Dimensions(v AWS_Pinpoint_Campaign_EventDimensions) AWS_Pinpoint_Campaign_CampaignEventFilter {
	t.Dimensions = cfz.V(v)
	return t
}

// Set__FilterType updates property "FilterType".
func (t AWS_Pinpoint_Campaign_CampaignEventFilter) Set__FilterType(v cfz.Expression[string]) AWS_Pinpoint_Campaign_CampaignEventFilter {
	t.FilterType = v
	return t
}

// SetV__FilterType updates property "FilterType".
func (t AWS_Pinpoint_Campaign_CampaignEventFilter) SetV__FilterType(v string) AWS_Pinpoint_Campaign_CampaignEventFilter {
	t.FilterType = cfz.V(v)
	return t
}
