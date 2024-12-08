// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_Av1ColorSpaceSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.Av1ColorSpaceSettings.
	AWS_MediaLive_Channel_Av1ColorSpaceSettings__Type = "AWS::MediaLive::Channel.Av1ColorSpaceSettings"
)

var (
	// AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.Av1ColorSpaceSettings.
	AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesMap = struct {
		ColorSpacePassthroughSettings string
		Hdr10Settings                 string
		Rec601Settings                string
		Rec709Settings                string
	}{
		ColorSpacePassthroughSettings: "ColorSpacePassthroughSettings",
		Hdr10Settings:                 "Hdr10Settings",
		Rec601Settings:                "Rec601Settings",
		Rec709Settings:                "Rec709Settings",
	}

	// AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.Av1ColorSpaceSettings.
	AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesMap.ColorSpacePassthroughSettings,
		AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesMap.Hdr10Settings,
		AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesMap.Rec601Settings,
		AWS_MediaLive_Channel_Av1ColorSpaceSettings__PropertiesMap.Rec709Settings,
	}
)

// AWS_MediaLive_Channel_Av1ColorSpaceSettings is a binding for AWS::MediaLive::Channel.Av1ColorSpaceSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html
type AWS_MediaLive_Channel_Av1ColorSpaceSettings struct {
	// ColorSpacePassthroughSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-colorspacepassthroughsettings
	ColorSpacePassthroughSettings cfz.Expression[AWS_MediaLive_Channel_ColorSpacePassthroughSettings] `json:"ColorSpacePassthroughSettings,omitempty"`

	// Hdr10Settings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-hdr10settings
	Hdr10Settings cfz.Expression[AWS_MediaLive_Channel_Hdr10Settings] `json:"Hdr10Settings,omitempty"`

	// Rec601Settings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-rec601settings
	Rec601Settings cfz.Expression[AWS_MediaLive_Channel_Rec601Settings] `json:"Rec601Settings,omitempty"`

	// Rec709Settings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-av1colorspacesettings.html#cfn-medialive-channel-av1colorspacesettings-rec709settings
	Rec709Settings cfz.Expression[AWS_MediaLive_Channel_Rec709Settings] `json:"Rec709Settings,omitempty"`
}

// New__AWS_MediaLive_Channel_Av1ColorSpaceSettings initializes a new AWS_MediaLive_Channel_Av1ColorSpaceSettings.
func New__AWS_MediaLive_Channel_Av1ColorSpaceSettings() AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	return AWS_MediaLive_Channel_Av1ColorSpaceSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_Av1ColorSpaceSettings) GetType() string {
	return AWS_MediaLive_Channel_Av1ColorSpaceSettings__Type
}

// Set__ColorSpacePassthroughSettings updates property "ColorSpacePassthroughSettings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) Set__ColorSpacePassthroughSettings(v cfz.Expression[AWS_MediaLive_Channel_ColorSpacePassthroughSettings]) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.ColorSpacePassthroughSettings = v
	return t
}

// SetV__ColorSpacePassthroughSettings updates property "ColorSpacePassthroughSettings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) SetV__ColorSpacePassthroughSettings(v AWS_MediaLive_Channel_ColorSpacePassthroughSettings) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.ColorSpacePassthroughSettings = cfz.V(v)
	return t
}

// Set__Hdr10Settings updates property "Hdr10Settings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) Set__Hdr10Settings(v cfz.Expression[AWS_MediaLive_Channel_Hdr10Settings]) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.Hdr10Settings = v
	return t
}

// SetV__Hdr10Settings updates property "Hdr10Settings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) SetV__Hdr10Settings(v AWS_MediaLive_Channel_Hdr10Settings) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.Hdr10Settings = cfz.V(v)
	return t
}

// Set__Rec601Settings updates property "Rec601Settings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) Set__Rec601Settings(v cfz.Expression[AWS_MediaLive_Channel_Rec601Settings]) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.Rec601Settings = v
	return t
}

// SetV__Rec601Settings updates property "Rec601Settings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) SetV__Rec601Settings(v AWS_MediaLive_Channel_Rec601Settings) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.Rec601Settings = cfz.V(v)
	return t
}

// Set__Rec709Settings updates property "Rec709Settings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) Set__Rec709Settings(v cfz.Expression[AWS_MediaLive_Channel_Rec709Settings]) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.Rec709Settings = v
	return t
}

// SetV__Rec709Settings updates property "Rec709Settings".
func (t AWS_MediaLive_Channel_Av1ColorSpaceSettings) SetV__Rec709Settings(v AWS_MediaLive_Channel_Rec709Settings) AWS_MediaLive_Channel_Av1ColorSpaceSettings {
	t.Rec709Settings = cfz.V(v)
	return t
}
