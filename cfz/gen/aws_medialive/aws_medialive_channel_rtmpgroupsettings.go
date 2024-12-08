// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_RtmpGroupSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.RtmpGroupSettings.
	AWS_MediaLive_Channel_RtmpGroupSettings__Type = "AWS::MediaLive::Channel.RtmpGroupSettings"
)

var (
	// AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.RtmpGroupSettings.
	AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap = struct {
		AdMarkers             string
		AuthenticationScheme  string
		CacheFullBehavior     string
		CacheLength           string
		CaptionData           string
		IncludeFillerNalUnits string
		InputLossAction       string
		RestartDelay          string
	}{
		AdMarkers:             "AdMarkers",
		AuthenticationScheme:  "AuthenticationScheme",
		CacheFullBehavior:     "CacheFullBehavior",
		CacheLength:           "CacheLength",
		CaptionData:           "CaptionData",
		IncludeFillerNalUnits: "IncludeFillerNalUnits",
		InputLossAction:       "InputLossAction",
		RestartDelay:          "RestartDelay",
	}

	// AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.RtmpGroupSettings.
	AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.AdMarkers,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.AuthenticationScheme,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.CacheFullBehavior,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.CacheLength,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.CaptionData,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.IncludeFillerNalUnits,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.InputLossAction,
		AWS_MediaLive_Channel_RtmpGroupSettings__PropertiesMap.RestartDelay,
	}
)

// AWS_MediaLive_Channel_RtmpGroupSettings is a binding for AWS::MediaLive::Channel.RtmpGroupSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html
type AWS_MediaLive_Channel_RtmpGroupSettings struct {
	// AdMarkers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-admarkers
	AdMarkers cfz.ExpressionSlice[string] `json:"AdMarkers,omitempty"`

	// AuthenticationScheme is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-authenticationscheme
	AuthenticationScheme cfz.Expression[string] `json:"AuthenticationScheme,omitempty"`

	// CacheFullBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-cachefullbehavior
	CacheFullBehavior cfz.Expression[string] `json:"CacheFullBehavior,omitempty"`

	// CacheLength is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-cachelength
	CacheLength cfz.Expression[int32] `json:"CacheLength,omitempty"`

	// CaptionData is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-captiondata
	CaptionData cfz.Expression[string] `json:"CaptionData,omitempty"`

	// IncludeFillerNalUnits is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-includefillernalunits
	IncludeFillerNalUnits cfz.Expression[string] `json:"IncludeFillerNalUnits,omitempty"`

	// InputLossAction is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-inputlossaction
	InputLossAction cfz.Expression[string] `json:"InputLossAction,omitempty"`

	// RestartDelay is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-restartdelay
	RestartDelay cfz.Expression[int32] `json:"RestartDelay,omitempty"`
}

// New__AWS_MediaLive_Channel_RtmpGroupSettings initializes a new AWS_MediaLive_Channel_RtmpGroupSettings.
func New__AWS_MediaLive_Channel_RtmpGroupSettings() AWS_MediaLive_Channel_RtmpGroupSettings {
	return AWS_MediaLive_Channel_RtmpGroupSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_RtmpGroupSettings) GetType() string {
	return AWS_MediaLive_Channel_RtmpGroupSettings__Type
}

// Set__AdMarkers updates property "AdMarkers".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__AdMarkers(v cfz.ExpressionSlice[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.AdMarkers = v
	return t
}

// SetS__AdMarkers updates property "AdMarkers".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetS__AdMarkers(v ...cfz.Expression[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.AdMarkers = cfz.S(v...)
	return t
}

// SetSV__AdMarkers updates property "AdMarkers".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetSV__AdMarkers(v ...string) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.AdMarkers = cfz.SV(v...)
	return t
}

// Set__AuthenticationScheme updates property "AuthenticationScheme".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__AuthenticationScheme(v cfz.Expression[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.AuthenticationScheme = v
	return t
}

// SetV__AuthenticationScheme updates property "AuthenticationScheme".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__AuthenticationScheme(v string) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.AuthenticationScheme = cfz.V(v)
	return t
}

// Set__CacheFullBehavior updates property "CacheFullBehavior".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__CacheFullBehavior(v cfz.Expression[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.CacheFullBehavior = v
	return t
}

// SetV__CacheFullBehavior updates property "CacheFullBehavior".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__CacheFullBehavior(v string) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.CacheFullBehavior = cfz.V(v)
	return t
}

// Set__CacheLength updates property "CacheLength".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__CacheLength(v cfz.Expression[int32]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.CacheLength = v
	return t
}

// SetV__CacheLength updates property "CacheLength".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__CacheLength(v int32) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.CacheLength = cfz.V(v)
	return t
}

// Set__CaptionData updates property "CaptionData".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__CaptionData(v cfz.Expression[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.CaptionData = v
	return t
}

// SetV__CaptionData updates property "CaptionData".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__CaptionData(v string) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.CaptionData = cfz.V(v)
	return t
}

// Set__IncludeFillerNalUnits updates property "IncludeFillerNalUnits".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__IncludeFillerNalUnits(v cfz.Expression[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.IncludeFillerNalUnits = v
	return t
}

// SetV__IncludeFillerNalUnits updates property "IncludeFillerNalUnits".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__IncludeFillerNalUnits(v string) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.IncludeFillerNalUnits = cfz.V(v)
	return t
}

// Set__InputLossAction updates property "InputLossAction".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__InputLossAction(v cfz.Expression[string]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.InputLossAction = v
	return t
}

// SetV__InputLossAction updates property "InputLossAction".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__InputLossAction(v string) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.InputLossAction = cfz.V(v)
	return t
}

// Set__RestartDelay updates property "RestartDelay".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) Set__RestartDelay(v cfz.Expression[int32]) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.RestartDelay = v
	return t
}

// SetV__RestartDelay updates property "RestartDelay".
func (t AWS_MediaLive_Channel_RtmpGroupSettings) SetV__RestartDelay(v int32) AWS_MediaLive_Channel_RtmpGroupSettings {
	t.RestartDelay = cfz.V(v)
	return t
}
