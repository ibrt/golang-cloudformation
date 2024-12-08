// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_AudioChannelMapping__Type is the CloudFormation type for AWS::MediaLive::Channel.AudioChannelMapping.
	AWS_MediaLive_Channel_AudioChannelMapping__Type = "AWS::MediaLive::Channel.AudioChannelMapping"
)

var (
	// AWS_MediaLive_Channel_AudioChannelMapping__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioChannelMapping.
	AWS_MediaLive_Channel_AudioChannelMapping__PropertiesMap = struct {
		InputChannelLevels string
		OutputChannel      string
	}{
		InputChannelLevels: "InputChannelLevels",
		OutputChannel:      "OutputChannel",
	}

	// AWS_MediaLive_Channel_AudioChannelMapping__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioChannelMapping.
	AWS_MediaLive_Channel_AudioChannelMapping__PropertiesSlice = []string{
		AWS_MediaLive_Channel_AudioChannelMapping__PropertiesMap.InputChannelLevels,
		AWS_MediaLive_Channel_AudioChannelMapping__PropertiesMap.OutputChannel,
	}
)

// AWS_MediaLive_Channel_AudioChannelMapping is a binding for AWS::MediaLive::Channel.AudioChannelMapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiochannelmapping.html
type AWS_MediaLive_Channel_AudioChannelMapping struct {
	// InputChannelLevels is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiochannelmapping.html#cfn-medialive-channel-audiochannelmapping-inputchannellevels
	InputChannelLevels cfz.ExpressionSlice[AWS_MediaLive_Channel_InputChannelLevel] `json:"InputChannelLevels,omitempty"`

	// OutputChannel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiochannelmapping.html#cfn-medialive-channel-audiochannelmapping-outputchannel
	OutputChannel cfz.Expression[int32] `json:"OutputChannel,omitempty"`
}

// New__AWS_MediaLive_Channel_AudioChannelMapping initializes a new AWS_MediaLive_Channel_AudioChannelMapping.
func New__AWS_MediaLive_Channel_AudioChannelMapping() AWS_MediaLive_Channel_AudioChannelMapping {
	return AWS_MediaLive_Channel_AudioChannelMapping{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_AudioChannelMapping) GetType() string {
	return AWS_MediaLive_Channel_AudioChannelMapping__Type
}

// Set__InputChannelLevels updates property "InputChannelLevels".
func (t AWS_MediaLive_Channel_AudioChannelMapping) Set__InputChannelLevels(v cfz.ExpressionSlice[AWS_MediaLive_Channel_InputChannelLevel]) AWS_MediaLive_Channel_AudioChannelMapping {
	t.InputChannelLevels = v
	return t
}

// SetS__InputChannelLevels updates property "InputChannelLevels".
func (t AWS_MediaLive_Channel_AudioChannelMapping) SetS__InputChannelLevels(v ...cfz.Expression[AWS_MediaLive_Channel_InputChannelLevel]) AWS_MediaLive_Channel_AudioChannelMapping {
	t.InputChannelLevels = cfz.S(v...)
	return t
}

// SetSV__InputChannelLevels updates property "InputChannelLevels".
func (t AWS_MediaLive_Channel_AudioChannelMapping) SetSV__InputChannelLevels(v ...AWS_MediaLive_Channel_InputChannelLevel) AWS_MediaLive_Channel_AudioChannelMapping {
	t.InputChannelLevels = cfz.SV(v...)
	return t
}

// Set__OutputChannel updates property "OutputChannel".
func (t AWS_MediaLive_Channel_AudioChannelMapping) Set__OutputChannel(v cfz.Expression[int32]) AWS_MediaLive_Channel_AudioChannelMapping {
	t.OutputChannel = v
	return t
}

// SetV__OutputChannel updates property "OutputChannel".
func (t AWS_MediaLive_Channel_AudioChannelMapping) SetV__OutputChannel(v int32) AWS_MediaLive_Channel_AudioChannelMapping {
	t.OutputChannel = cfz.V(v)
	return t
}
