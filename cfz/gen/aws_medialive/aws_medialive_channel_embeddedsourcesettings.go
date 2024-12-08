// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_EmbeddedSourceSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.EmbeddedSourceSettings.
	AWS_MediaLive_Channel_EmbeddedSourceSettings__Type = "AWS::MediaLive::Channel.EmbeddedSourceSettings"
)

var (
	// AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.EmbeddedSourceSettings.
	AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesMap = struct {
		Convert608To708        string
		Scte20Detection        string
		Source608ChannelNumber string
		Source608TrackNumber   string
	}{
		Convert608To708:        "Convert608To708",
		Scte20Detection:        "Scte20Detection",
		Source608ChannelNumber: "Source608ChannelNumber",
		Source608TrackNumber:   "Source608TrackNumber",
	}

	// AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.EmbeddedSourceSettings.
	AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesMap.Convert608To708,
		AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesMap.Scte20Detection,
		AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesMap.Source608ChannelNumber,
		AWS_MediaLive_Channel_EmbeddedSourceSettings__PropertiesMap.Source608TrackNumber,
	}
)

// AWS_MediaLive_Channel_EmbeddedSourceSettings is a binding for AWS::MediaLive::Channel.EmbeddedSourceSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html
type AWS_MediaLive_Channel_EmbeddedSourceSettings struct {
	// Convert608To708 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-convert608to708
	Convert608To708 cfz.Expression[string] `json:"Convert608To708,omitempty"`

	// Scte20Detection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-scte20detection
	Scte20Detection cfz.Expression[string] `json:"Scte20Detection,omitempty"`

	// Source608ChannelNumber is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-source608channelnumber
	Source608ChannelNumber cfz.Expression[int32] `json:"Source608ChannelNumber,omitempty"`

	// Source608TrackNumber is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedsourcesettings.html#cfn-medialive-channel-embeddedsourcesettings-source608tracknumber
	Source608TrackNumber cfz.Expression[int32] `json:"Source608TrackNumber,omitempty"`
}

// New__AWS_MediaLive_Channel_EmbeddedSourceSettings initializes a new AWS_MediaLive_Channel_EmbeddedSourceSettings.
func New__AWS_MediaLive_Channel_EmbeddedSourceSettings() AWS_MediaLive_Channel_EmbeddedSourceSettings {
	return AWS_MediaLive_Channel_EmbeddedSourceSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_EmbeddedSourceSettings) GetType() string {
	return AWS_MediaLive_Channel_EmbeddedSourceSettings__Type
}

// Set__Convert608To708 updates property "Convert608To708".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) Set__Convert608To708(v cfz.Expression[string]) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Convert608To708 = v
	return t
}

// SetV__Convert608To708 updates property "Convert608To708".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) SetV__Convert608To708(v string) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Convert608To708 = cfz.V(v)
	return t
}

// Set__Scte20Detection updates property "Scte20Detection".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) Set__Scte20Detection(v cfz.Expression[string]) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Scte20Detection = v
	return t
}

// SetV__Scte20Detection updates property "Scte20Detection".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) SetV__Scte20Detection(v string) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Scte20Detection = cfz.V(v)
	return t
}

// Set__Source608ChannelNumber updates property "Source608ChannelNumber".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) Set__Source608ChannelNumber(v cfz.Expression[int32]) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Source608ChannelNumber = v
	return t
}

// SetV__Source608ChannelNumber updates property "Source608ChannelNumber".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) SetV__Source608ChannelNumber(v int32) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Source608ChannelNumber = cfz.V(v)
	return t
}

// Set__Source608TrackNumber updates property "Source608TrackNumber".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) Set__Source608TrackNumber(v cfz.Expression[int32]) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Source608TrackNumber = v
	return t
}

// SetV__Source608TrackNumber updates property "Source608TrackNumber".
func (t AWS_MediaLive_Channel_EmbeddedSourceSettings) SetV__Source608TrackNumber(v int32) AWS_MediaLive_Channel_EmbeddedSourceSettings {
	t.Source608TrackNumber = cfz.V(v)
	return t
}
