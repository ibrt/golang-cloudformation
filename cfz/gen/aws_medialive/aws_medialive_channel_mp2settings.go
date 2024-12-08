// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_Mp2Settings__Type is the CloudFormation type for AWS::MediaLive::Channel.Mp2Settings.
	AWS_MediaLive_Channel_Mp2Settings__Type = "AWS::MediaLive::Channel.Mp2Settings"
)

var (
	// AWS_MediaLive_Channel_Mp2Settings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.Mp2Settings.
	AWS_MediaLive_Channel_Mp2Settings__PropertiesMap = struct {
		Bitrate    string
		CodingMode string
		SampleRate string
	}{
		Bitrate:    "Bitrate",
		CodingMode: "CodingMode",
		SampleRate: "SampleRate",
	}

	// AWS_MediaLive_Channel_Mp2Settings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.Mp2Settings.
	AWS_MediaLive_Channel_Mp2Settings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_Mp2Settings__PropertiesMap.Bitrate,
		AWS_MediaLive_Channel_Mp2Settings__PropertiesMap.CodingMode,
		AWS_MediaLive_Channel_Mp2Settings__PropertiesMap.SampleRate,
	}
)

// AWS_MediaLive_Channel_Mp2Settings is a binding for AWS::MediaLive::Channel.Mp2Settings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html
type AWS_MediaLive_Channel_Mp2Settings struct {
	// Bitrate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html#cfn-medialive-channel-mp2settings-bitrate
	Bitrate cfz.Expression[float64] `json:"Bitrate,omitempty"`

	// CodingMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html#cfn-medialive-channel-mp2settings-codingmode
	CodingMode cfz.Expression[string] `json:"CodingMode,omitempty"`

	// SampleRate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-mp2settings.html#cfn-medialive-channel-mp2settings-samplerate
	SampleRate cfz.Expression[float64] `json:"SampleRate,omitempty"`
}

// New__AWS_MediaLive_Channel_Mp2Settings initializes a new AWS_MediaLive_Channel_Mp2Settings.
func New__AWS_MediaLive_Channel_Mp2Settings() AWS_MediaLive_Channel_Mp2Settings {
	return AWS_MediaLive_Channel_Mp2Settings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_Mp2Settings) GetType() string {
	return AWS_MediaLive_Channel_Mp2Settings__Type
}

// Set__Bitrate updates property "Bitrate".
func (t AWS_MediaLive_Channel_Mp2Settings) Set__Bitrate(v cfz.Expression[float64]) AWS_MediaLive_Channel_Mp2Settings {
	t.Bitrate = v
	return t
}

// SetV__Bitrate updates property "Bitrate".
func (t AWS_MediaLive_Channel_Mp2Settings) SetV__Bitrate(v float64) AWS_MediaLive_Channel_Mp2Settings {
	t.Bitrate = cfz.V(v)
	return t
}

// Set__CodingMode updates property "CodingMode".
func (t AWS_MediaLive_Channel_Mp2Settings) Set__CodingMode(v cfz.Expression[string]) AWS_MediaLive_Channel_Mp2Settings {
	t.CodingMode = v
	return t
}

// SetV__CodingMode updates property "CodingMode".
func (t AWS_MediaLive_Channel_Mp2Settings) SetV__CodingMode(v string) AWS_MediaLive_Channel_Mp2Settings {
	t.CodingMode = cfz.V(v)
	return t
}

// Set__SampleRate updates property "SampleRate".
func (t AWS_MediaLive_Channel_Mp2Settings) Set__SampleRate(v cfz.Expression[float64]) AWS_MediaLive_Channel_Mp2Settings {
	t.SampleRate = v
	return t
}

// SetV__SampleRate updates property "SampleRate".
func (t AWS_MediaLive_Channel_Mp2Settings) SetV__SampleRate(v float64) AWS_MediaLive_Channel_Mp2Settings {
	t.SampleRate = cfz.V(v)
	return t
}
