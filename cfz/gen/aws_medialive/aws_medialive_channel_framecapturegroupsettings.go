// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_FrameCaptureGroupSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.FrameCaptureGroupSettings.
	AWS_MediaLive_Channel_FrameCaptureGroupSettings__Type = "AWS::MediaLive::Channel.FrameCaptureGroupSettings"
)

var (
	// AWS_MediaLive_Channel_FrameCaptureGroupSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.FrameCaptureGroupSettings.
	AWS_MediaLive_Channel_FrameCaptureGroupSettings__PropertiesMap = struct {
		Destination             string
		FrameCaptureCdnSettings string
	}{
		Destination:             "Destination",
		FrameCaptureCdnSettings: "FrameCaptureCdnSettings",
	}

	// AWS_MediaLive_Channel_FrameCaptureGroupSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.FrameCaptureGroupSettings.
	AWS_MediaLive_Channel_FrameCaptureGroupSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_FrameCaptureGroupSettings__PropertiesMap.Destination,
		AWS_MediaLive_Channel_FrameCaptureGroupSettings__PropertiesMap.FrameCaptureCdnSettings,
	}
)

// AWS_MediaLive_Channel_FrameCaptureGroupSettings is a binding for AWS::MediaLive::Channel.FrameCaptureGroupSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturegroupsettings.html
type AWS_MediaLive_Channel_FrameCaptureGroupSettings struct {
	// Destination is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturegroupsettings.html#cfn-medialive-channel-framecapturegroupsettings-destination
	Destination cfz.Expression[AWS_MediaLive_Channel_OutputLocationRef] `json:"Destination,omitempty"`

	// FrameCaptureCdnSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturegroupsettings.html#cfn-medialive-channel-framecapturegroupsettings-framecapturecdnsettings
	FrameCaptureCdnSettings cfz.Expression[AWS_MediaLive_Channel_FrameCaptureCdnSettings] `json:"FrameCaptureCdnSettings,omitempty"`
}

// New__AWS_MediaLive_Channel_FrameCaptureGroupSettings initializes a new AWS_MediaLive_Channel_FrameCaptureGroupSettings.
func New__AWS_MediaLive_Channel_FrameCaptureGroupSettings() AWS_MediaLive_Channel_FrameCaptureGroupSettings {
	return AWS_MediaLive_Channel_FrameCaptureGroupSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_FrameCaptureGroupSettings) GetType() string {
	return AWS_MediaLive_Channel_FrameCaptureGroupSettings__Type
}

// Set__Destination updates property "Destination".
func (t AWS_MediaLive_Channel_FrameCaptureGroupSettings) Set__Destination(v cfz.Expression[AWS_MediaLive_Channel_OutputLocationRef]) AWS_MediaLive_Channel_FrameCaptureGroupSettings {
	t.Destination = v
	return t
}

// SetV__Destination updates property "Destination".
func (t AWS_MediaLive_Channel_FrameCaptureGroupSettings) SetV__Destination(v AWS_MediaLive_Channel_OutputLocationRef) AWS_MediaLive_Channel_FrameCaptureGroupSettings {
	t.Destination = cfz.V(v)
	return t
}

// Set__FrameCaptureCdnSettings updates property "FrameCaptureCdnSettings".
func (t AWS_MediaLive_Channel_FrameCaptureGroupSettings) Set__FrameCaptureCdnSettings(v cfz.Expression[AWS_MediaLive_Channel_FrameCaptureCdnSettings]) AWS_MediaLive_Channel_FrameCaptureGroupSettings {
	t.FrameCaptureCdnSettings = v
	return t
}

// SetV__FrameCaptureCdnSettings updates property "FrameCaptureCdnSettings".
func (t AWS_MediaLive_Channel_FrameCaptureGroupSettings) SetV__FrameCaptureCdnSettings(v AWS_MediaLive_Channel_FrameCaptureCdnSettings) AWS_MediaLive_Channel_FrameCaptureGroupSettings {
	t.FrameCaptureCdnSettings = cfz.V(v)
	return t
}
