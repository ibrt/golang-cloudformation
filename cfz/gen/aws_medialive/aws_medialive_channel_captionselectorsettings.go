// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_CaptionSelectorSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.CaptionSelectorSettings.
	AWS_MediaLive_Channel_CaptionSelectorSettings__Type = "AWS::MediaLive::Channel.CaptionSelectorSettings"
)

var (
	// AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.CaptionSelectorSettings.
	AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap = struct {
		AncillarySourceSettings string
		AribSourceSettings      string
		DvbSubSourceSettings    string
		EmbeddedSourceSettings  string
		Scte20SourceSettings    string
		Scte27SourceSettings    string
		TeletextSourceSettings  string
	}{
		AncillarySourceSettings: "AncillarySourceSettings",
		AribSourceSettings:      "AribSourceSettings",
		DvbSubSourceSettings:    "DvbSubSourceSettings",
		EmbeddedSourceSettings:  "EmbeddedSourceSettings",
		Scte20SourceSettings:    "Scte20SourceSettings",
		Scte27SourceSettings:    "Scte27SourceSettings",
		TeletextSourceSettings:  "TeletextSourceSettings",
	}

	// AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.CaptionSelectorSettings.
	AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.AncillarySourceSettings,
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.AribSourceSettings,
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.DvbSubSourceSettings,
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.EmbeddedSourceSettings,
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.Scte20SourceSettings,
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.Scte27SourceSettings,
		AWS_MediaLive_Channel_CaptionSelectorSettings__PropertiesMap.TeletextSourceSettings,
	}
)

// AWS_MediaLive_Channel_CaptionSelectorSettings is a binding for AWS::MediaLive::Channel.CaptionSelectorSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html
type AWS_MediaLive_Channel_CaptionSelectorSettings struct {
	// AncillarySourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-ancillarysourcesettings
	AncillarySourceSettings cfz.Expression[AWS_MediaLive_Channel_AncillarySourceSettings] `json:"AncillarySourceSettings,omitempty"`

	// AribSourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-aribsourcesettings
	AribSourceSettings cfz.Expression[AWS_MediaLive_Channel_AribSourceSettings] `json:"AribSourceSettings,omitempty"`

	// DvbSubSourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-dvbsubsourcesettings
	DvbSubSourceSettings cfz.Expression[AWS_MediaLive_Channel_DvbSubSourceSettings] `json:"DvbSubSourceSettings,omitempty"`

	// EmbeddedSourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-embeddedsourcesettings
	EmbeddedSourceSettings cfz.Expression[AWS_MediaLive_Channel_EmbeddedSourceSettings] `json:"EmbeddedSourceSettings,omitempty"`

	// Scte20SourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-scte20sourcesettings
	Scte20SourceSettings cfz.Expression[AWS_MediaLive_Channel_Scte20SourceSettings] `json:"Scte20SourceSettings,omitempty"`

	// Scte27SourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-scte27sourcesettings
	Scte27SourceSettings cfz.Expression[AWS_MediaLive_Channel_Scte27SourceSettings] `json:"Scte27SourceSettings,omitempty"`

	// TeletextSourceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-captionselectorsettings.html#cfn-medialive-channel-captionselectorsettings-teletextsourcesettings
	TeletextSourceSettings cfz.Expression[AWS_MediaLive_Channel_TeletextSourceSettings] `json:"TeletextSourceSettings,omitempty"`
}

// New__AWS_MediaLive_Channel_CaptionSelectorSettings initializes a new AWS_MediaLive_Channel_CaptionSelectorSettings.
func New__AWS_MediaLive_Channel_CaptionSelectorSettings() AWS_MediaLive_Channel_CaptionSelectorSettings {
	return AWS_MediaLive_Channel_CaptionSelectorSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_CaptionSelectorSettings) GetType() string {
	return AWS_MediaLive_Channel_CaptionSelectorSettings__Type
}

// Set__AncillarySourceSettings updates property "AncillarySourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__AncillarySourceSettings(v cfz.Expression[AWS_MediaLive_Channel_AncillarySourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.AncillarySourceSettings = v
	return t
}

// SetV__AncillarySourceSettings updates property "AncillarySourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__AncillarySourceSettings(v AWS_MediaLive_Channel_AncillarySourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.AncillarySourceSettings = cfz.V(v)
	return t
}

// Set__AribSourceSettings updates property "AribSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__AribSourceSettings(v cfz.Expression[AWS_MediaLive_Channel_AribSourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.AribSourceSettings = v
	return t
}

// SetV__AribSourceSettings updates property "AribSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__AribSourceSettings(v AWS_MediaLive_Channel_AribSourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.AribSourceSettings = cfz.V(v)
	return t
}

// Set__DvbSubSourceSettings updates property "DvbSubSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__DvbSubSourceSettings(v cfz.Expression[AWS_MediaLive_Channel_DvbSubSourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.DvbSubSourceSettings = v
	return t
}

// SetV__DvbSubSourceSettings updates property "DvbSubSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__DvbSubSourceSettings(v AWS_MediaLive_Channel_DvbSubSourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.DvbSubSourceSettings = cfz.V(v)
	return t
}

// Set__EmbeddedSourceSettings updates property "EmbeddedSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__EmbeddedSourceSettings(v cfz.Expression[AWS_MediaLive_Channel_EmbeddedSourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.EmbeddedSourceSettings = v
	return t
}

// SetV__EmbeddedSourceSettings updates property "EmbeddedSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__EmbeddedSourceSettings(v AWS_MediaLive_Channel_EmbeddedSourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.EmbeddedSourceSettings = cfz.V(v)
	return t
}

// Set__Scte20SourceSettings updates property "Scte20SourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__Scte20SourceSettings(v cfz.Expression[AWS_MediaLive_Channel_Scte20SourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.Scte20SourceSettings = v
	return t
}

// SetV__Scte20SourceSettings updates property "Scte20SourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__Scte20SourceSettings(v AWS_MediaLive_Channel_Scte20SourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.Scte20SourceSettings = cfz.V(v)
	return t
}

// Set__Scte27SourceSettings updates property "Scte27SourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__Scte27SourceSettings(v cfz.Expression[AWS_MediaLive_Channel_Scte27SourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.Scte27SourceSettings = v
	return t
}

// SetV__Scte27SourceSettings updates property "Scte27SourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__Scte27SourceSettings(v AWS_MediaLive_Channel_Scte27SourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.Scte27SourceSettings = cfz.V(v)
	return t
}

// Set__TeletextSourceSettings updates property "TeletextSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) Set__TeletextSourceSettings(v cfz.Expression[AWS_MediaLive_Channel_TeletextSourceSettings]) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.TeletextSourceSettings = v
	return t
}

// SetV__TeletextSourceSettings updates property "TeletextSourceSettings".
func (t AWS_MediaLive_Channel_CaptionSelectorSettings) SetV__TeletextSourceSettings(v AWS_MediaLive_Channel_TeletextSourceSettings) AWS_MediaLive_Channel_CaptionSelectorSettings {
	t.TeletextSourceSettings = cfz.V(v)
	return t
}
