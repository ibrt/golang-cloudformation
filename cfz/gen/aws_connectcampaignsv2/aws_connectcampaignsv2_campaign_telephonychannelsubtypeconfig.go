// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_connectcampaignsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__Type is the CloudFormation type for AWS::ConnectCampaignsV2::Campaign.TelephonyChannelSubtypeConfig.
	AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__Type = "AWS::ConnectCampaignsV2::Campaign.TelephonyChannelSubtypeConfig"
)

var (
	// AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesMap reports all the CloudFormation properties for AWS::ConnectCampaignsV2::Campaign.TelephonyChannelSubtypeConfig.
	AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesMap = struct {
		Capacity              string
		ConnectQueueId        string
		DefaultOutboundConfig string
		OutboundMode          string
	}{
		Capacity:              "Capacity",
		ConnectQueueId:        "ConnectQueueId",
		DefaultOutboundConfig: "DefaultOutboundConfig",
		OutboundMode:          "OutboundMode",
	}

	// AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesSlice reports all the CloudFormation properties for AWS::ConnectCampaignsV2::Campaign.TelephonyChannelSubtypeConfig.
	AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesSlice = []string{
		AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesMap.Capacity,
		AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesMap.ConnectQueueId,
		AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesMap.DefaultOutboundConfig,
		AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__PropertiesMap.OutboundMode,
	}
)

// AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig is a binding for AWS::ConnectCampaignsV2::Campaign.TelephonyChannelSubtypeConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html
type AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig struct {
	// Capacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-capacity
	Capacity cfz.Expression[float64] `json:"Capacity,omitempty"`

	// ConnectQueueId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-connectqueueid
	ConnectQueueId cfz.Expression[string] `json:"ConnectQueueId,omitempty"`

	// DefaultOutboundConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-defaultoutboundconfig
	DefaultOutboundConfig cfz.Expression[AWS_ConnectCampaignsV2_Campaign_TelephonyOutboundConfig] `json:"DefaultOutboundConfig,omitempty"`

	// OutboundMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-outboundmode
	OutboundMode cfz.Expression[AWS_ConnectCampaignsV2_Campaign_TelephonyOutboundMode] `json:"OutboundMode,omitempty"`
}

// New__AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig initializes a new AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig.
func New__AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig() AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	return AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) GetType() string {
	return AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig__Type
}

// Set__Capacity updates property "Capacity".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) Set__Capacity(v cfz.Expression[float64]) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.Capacity = v
	return t
}

// SetV__Capacity updates property "Capacity".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) SetV__Capacity(v float64) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.Capacity = cfz.V(v)
	return t
}

// Set__ConnectQueueId updates property "ConnectQueueId".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) Set__ConnectQueueId(v cfz.Expression[string]) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.ConnectQueueId = v
	return t
}

// SetV__ConnectQueueId updates property "ConnectQueueId".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) SetV__ConnectQueueId(v string) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.ConnectQueueId = cfz.V(v)
	return t
}

// Set__DefaultOutboundConfig updates property "DefaultOutboundConfig".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) Set__DefaultOutboundConfig(v cfz.Expression[AWS_ConnectCampaignsV2_Campaign_TelephonyOutboundConfig]) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.DefaultOutboundConfig = v
	return t
}

// SetV__DefaultOutboundConfig updates property "DefaultOutboundConfig".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) SetV__DefaultOutboundConfig(v AWS_ConnectCampaignsV2_Campaign_TelephonyOutboundConfig) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.DefaultOutboundConfig = cfz.V(v)
	return t
}

// Set__OutboundMode updates property "OutboundMode".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) Set__OutboundMode(v cfz.Expression[AWS_ConnectCampaignsV2_Campaign_TelephonyOutboundMode]) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.OutboundMode = v
	return t
}

// SetV__OutboundMode updates property "OutboundMode".
func (t AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig) SetV__OutboundMode(v AWS_ConnectCampaignsV2_Campaign_TelephonyOutboundMode) AWS_ConnectCampaignsV2_Campaign_TelephonyChannelSubtypeConfig {
	t.OutboundMode = cfz.V(v)
	return t
}
