// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_WavSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.WavSettings.
	AWS_MediaLive_Channel_WavSettings__Type = "AWS::MediaLive::Channel.WavSettings"
)

var (
	// AWS_MediaLive_Channel_WavSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.WavSettings.
	AWS_MediaLive_Channel_WavSettings__PropertiesMap = struct {
		BitDepth   string
		CodingMode string
		SampleRate string
	}{
		BitDepth:   "BitDepth",
		CodingMode: "CodingMode",
		SampleRate: "SampleRate",
	}

	// AWS_MediaLive_Channel_WavSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.WavSettings.
	AWS_MediaLive_Channel_WavSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_WavSettings__PropertiesMap.BitDepth,
		AWS_MediaLive_Channel_WavSettings__PropertiesMap.CodingMode,
		AWS_MediaLive_Channel_WavSettings__PropertiesMap.SampleRate,
	}
)

// AWS_MediaLive_Channel_WavSettings is a binding for AWS::MediaLive::Channel.WavSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html
type AWS_MediaLive_Channel_WavSettings struct {
	// BitDepth is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-bitdepth
	BitDepth cfz.Expression[float64] `json:"BitDepth,omitempty"`

	// CodingMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-codingmode
	CodingMode cfz.Expression[string] `json:"CodingMode,omitempty"`

	// SampleRate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-wavsettings.html#cfn-medialive-channel-wavsettings-samplerate
	SampleRate cfz.Expression[float64] `json:"SampleRate,omitempty"`
}

// New__AWS_MediaLive_Channel_WavSettings initializes a new AWS_MediaLive_Channel_WavSettings.
func New__AWS_MediaLive_Channel_WavSettings() AWS_MediaLive_Channel_WavSettings {
	return AWS_MediaLive_Channel_WavSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_WavSettings) GetType() string {
	return AWS_MediaLive_Channel_WavSettings__Type
}

// Set__BitDepth updates property "BitDepth".
func (t AWS_MediaLive_Channel_WavSettings) Set__BitDepth(v cfz.Expression[float64]) AWS_MediaLive_Channel_WavSettings {
	t.BitDepth = v
	return t
}

// SetV__BitDepth updates property "BitDepth".
func (t AWS_MediaLive_Channel_WavSettings) SetV__BitDepth(v float64) AWS_MediaLive_Channel_WavSettings {
	t.BitDepth = cfz.V(v)
	return t
}

// Set__CodingMode updates property "CodingMode".
func (t AWS_MediaLive_Channel_WavSettings) Set__CodingMode(v cfz.Expression[string]) AWS_MediaLive_Channel_WavSettings {
	t.CodingMode = v
	return t
}

// SetV__CodingMode updates property "CodingMode".
func (t AWS_MediaLive_Channel_WavSettings) SetV__CodingMode(v string) AWS_MediaLive_Channel_WavSettings {
	t.CodingMode = cfz.V(v)
	return t
}

// Set__SampleRate updates property "SampleRate".
func (t AWS_MediaLive_Channel_WavSettings) Set__SampleRate(v cfz.Expression[float64]) AWS_MediaLive_Channel_WavSettings {
	t.SampleRate = v
	return t
}

// SetV__SampleRate updates property "SampleRate".
func (t AWS_MediaLive_Channel_WavSettings) SetV__SampleRate(v float64) AWS_MediaLive_Channel_WavSettings {
	t.SampleRate = cfz.V(v)
	return t
}
