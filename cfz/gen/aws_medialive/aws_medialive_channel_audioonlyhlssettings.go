// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_AudioOnlyHlsSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.AudioOnlyHlsSettings.
	AWS_MediaLive_Channel_AudioOnlyHlsSettings__Type = "AWS::MediaLive::Channel.AudioOnlyHlsSettings"
)

var (
	// AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioOnlyHlsSettings.
	AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesMap = struct {
		AudioGroupId   string
		AudioOnlyImage string
		AudioTrackType string
		SegmentType    string
	}{
		AudioGroupId:   "AudioGroupId",
		AudioOnlyImage: "AudioOnlyImage",
		AudioTrackType: "AudioTrackType",
		SegmentType:    "SegmentType",
	}

	// AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioOnlyHlsSettings.
	AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesMap.AudioGroupId,
		AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesMap.AudioOnlyImage,
		AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesMap.AudioTrackType,
		AWS_MediaLive_Channel_AudioOnlyHlsSettings__PropertiesMap.SegmentType,
	}
)

// AWS_MediaLive_Channel_AudioOnlyHlsSettings is a binding for AWS::MediaLive::Channel.AudioOnlyHlsSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html
type AWS_MediaLive_Channel_AudioOnlyHlsSettings struct {
	// AudioGroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-audiogroupid
	AudioGroupId cfz.Expression[string] `json:"AudioGroupId,omitempty"`

	// AudioOnlyImage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-audioonlyimage
	AudioOnlyImage cfz.Expression[AWS_MediaLive_Channel_InputLocation] `json:"AudioOnlyImage,omitempty"`

	// AudioTrackType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-audiotracktype
	AudioTrackType cfz.Expression[string] `json:"AudioTrackType,omitempty"`

	// SegmentType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioonlyhlssettings.html#cfn-medialive-channel-audioonlyhlssettings-segmenttype
	SegmentType cfz.Expression[string] `json:"SegmentType,omitempty"`
}

// New__AWS_MediaLive_Channel_AudioOnlyHlsSettings initializes a new AWS_MediaLive_Channel_AudioOnlyHlsSettings.
func New__AWS_MediaLive_Channel_AudioOnlyHlsSettings() AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	return AWS_MediaLive_Channel_AudioOnlyHlsSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_AudioOnlyHlsSettings) GetType() string {
	return AWS_MediaLive_Channel_AudioOnlyHlsSettings__Type
}

// Set__AudioGroupId updates property "AudioGroupId".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) Set__AudioGroupId(v cfz.Expression[string]) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.AudioGroupId = v
	return t
}

// SetV__AudioGroupId updates property "AudioGroupId".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) SetV__AudioGroupId(v string) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.AudioGroupId = cfz.V(v)
	return t
}

// Set__AudioOnlyImage updates property "AudioOnlyImage".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) Set__AudioOnlyImage(v cfz.Expression[AWS_MediaLive_Channel_InputLocation]) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.AudioOnlyImage = v
	return t
}

// SetV__AudioOnlyImage updates property "AudioOnlyImage".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) SetV__AudioOnlyImage(v AWS_MediaLive_Channel_InputLocation) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.AudioOnlyImage = cfz.V(v)
	return t
}

// Set__AudioTrackType updates property "AudioTrackType".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) Set__AudioTrackType(v cfz.Expression[string]) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.AudioTrackType = v
	return t
}

// SetV__AudioTrackType updates property "AudioTrackType".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) SetV__AudioTrackType(v string) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.AudioTrackType = cfz.V(v)
	return t
}

// Set__SegmentType updates property "SegmentType".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) Set__SegmentType(v cfz.Expression[string]) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.SegmentType = v
	return t
}

// SetV__SegmentType updates property "SegmentType".
func (t AWS_MediaLive_Channel_AudioOnlyHlsSettings) SetV__SegmentType(v string) AWS_MediaLive_Channel_AudioOnlyHlsSettings {
	t.SegmentType = cfz.V(v)
	return t
}
