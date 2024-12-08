// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_AnywhereSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.AnywhereSettings.
	AWS_MediaLive_Channel_AnywhereSettings__Type = "AWS::MediaLive::Channel.AnywhereSettings"
)

var (
	// AWS_MediaLive_Channel_AnywhereSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.AnywhereSettings.
	AWS_MediaLive_Channel_AnywhereSettings__PropertiesMap = struct {
		ChannelPlacementGroupId string
		ClusterId               string
	}{
		ChannelPlacementGroupId: "ChannelPlacementGroupId",
		ClusterId:               "ClusterId",
	}

	// AWS_MediaLive_Channel_AnywhereSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.AnywhereSettings.
	AWS_MediaLive_Channel_AnywhereSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_AnywhereSettings__PropertiesMap.ChannelPlacementGroupId,
		AWS_MediaLive_Channel_AnywhereSettings__PropertiesMap.ClusterId,
	}
)

// AWS_MediaLive_Channel_AnywhereSettings is a binding for AWS::MediaLive::Channel.AnywhereSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-anywheresettings.html
type AWS_MediaLive_Channel_AnywhereSettings struct {
	// ChannelPlacementGroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-anywheresettings.html#cfn-medialive-channel-anywheresettings-channelplacementgroupid
	ChannelPlacementGroupId cfz.Expression[string] `json:"ChannelPlacementGroupId,omitempty"`

	// ClusterId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-anywheresettings.html#cfn-medialive-channel-anywheresettings-clusterid
	ClusterId cfz.Expression[string] `json:"ClusterId,omitempty"`
}

// New__AWS_MediaLive_Channel_AnywhereSettings initializes a new AWS_MediaLive_Channel_AnywhereSettings.
func New__AWS_MediaLive_Channel_AnywhereSettings() AWS_MediaLive_Channel_AnywhereSettings {
	return AWS_MediaLive_Channel_AnywhereSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_AnywhereSettings) GetType() string {
	return AWS_MediaLive_Channel_AnywhereSettings__Type
}

// Set__ChannelPlacementGroupId updates property "ChannelPlacementGroupId".
func (t AWS_MediaLive_Channel_AnywhereSettings) Set__ChannelPlacementGroupId(v cfz.Expression[string]) AWS_MediaLive_Channel_AnywhereSettings {
	t.ChannelPlacementGroupId = v
	return t
}

// SetV__ChannelPlacementGroupId updates property "ChannelPlacementGroupId".
func (t AWS_MediaLive_Channel_AnywhereSettings) SetV__ChannelPlacementGroupId(v string) AWS_MediaLive_Channel_AnywhereSettings {
	t.ChannelPlacementGroupId = cfz.V(v)
	return t
}

// Set__ClusterId updates property "ClusterId".
func (t AWS_MediaLive_Channel_AnywhereSettings) Set__ClusterId(v cfz.Expression[string]) AWS_MediaLive_Channel_AnywhereSettings {
	t.ClusterId = v
	return t
}

// SetV__ClusterId updates property "ClusterId".
func (t AWS_MediaLive_Channel_AnywhereSettings) SetV__ClusterId(v string) AWS_MediaLive_Channel_AnywhereSettings {
	t.ClusterId = cfz.V(v)
	return t
}
