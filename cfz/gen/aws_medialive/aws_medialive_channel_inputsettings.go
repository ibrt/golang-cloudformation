// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_InputSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.InputSettings.
	AWS_MediaLive_Channel_InputSettings__Type = "AWS::MediaLive::Channel.InputSettings"
)

var (
	// AWS_MediaLive_Channel_InputSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.InputSettings.
	AWS_MediaLive_Channel_InputSettings__PropertiesMap = struct {
		AudioSelectors          string
		CaptionSelectors        string
		DeblockFilter           string
		DenoiseFilter           string
		FilterStrength          string
		InputFilter             string
		NetworkInputSettings    string
		Scte35Pid               string
		Smpte2038DataPreference string
		SourceEndBehavior       string
		VideoSelector           string
	}{
		AudioSelectors:          "AudioSelectors",
		CaptionSelectors:        "CaptionSelectors",
		DeblockFilter:           "DeblockFilter",
		DenoiseFilter:           "DenoiseFilter",
		FilterStrength:          "FilterStrength",
		InputFilter:             "InputFilter",
		NetworkInputSettings:    "NetworkInputSettings",
		Scte35Pid:               "Scte35Pid",
		Smpte2038DataPreference: "Smpte2038DataPreference",
		SourceEndBehavior:       "SourceEndBehavior",
		VideoSelector:           "VideoSelector",
	}

	// AWS_MediaLive_Channel_InputSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.InputSettings.
	AWS_MediaLive_Channel_InputSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.AudioSelectors,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.CaptionSelectors,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.DeblockFilter,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.DenoiseFilter,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.FilterStrength,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.InputFilter,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.NetworkInputSettings,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.Scte35Pid,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.Smpte2038DataPreference,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.SourceEndBehavior,
		AWS_MediaLive_Channel_InputSettings__PropertiesMap.VideoSelector,
	}
)

// AWS_MediaLive_Channel_InputSettings is a binding for AWS::MediaLive::Channel.InputSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html
type AWS_MediaLive_Channel_InputSettings struct {
	// AudioSelectors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-audioselectors
	AudioSelectors cfz.ExpressionSlice[AWS_MediaLive_Channel_AudioSelector] `json:"AudioSelectors,omitempty"`

	// CaptionSelectors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-captionselectors
	CaptionSelectors cfz.ExpressionSlice[AWS_MediaLive_Channel_CaptionSelector] `json:"CaptionSelectors,omitempty"`

	// DeblockFilter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-deblockfilter
	DeblockFilter cfz.Expression[string] `json:"DeblockFilter,omitempty"`

	// DenoiseFilter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-denoisefilter
	DenoiseFilter cfz.Expression[string] `json:"DenoiseFilter,omitempty"`

	// FilterStrength is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-filterstrength
	FilterStrength cfz.Expression[int32] `json:"FilterStrength,omitempty"`

	// InputFilter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-inputfilter
	InputFilter cfz.Expression[string] `json:"InputFilter,omitempty"`

	// NetworkInputSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-networkinputsettings
	NetworkInputSettings cfz.Expression[AWS_MediaLive_Channel_NetworkInputSettings] `json:"NetworkInputSettings,omitempty"`

	// Scte35Pid is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-scte35pid
	Scte35Pid cfz.Expression[int32] `json:"Scte35Pid,omitempty"`

	// Smpte2038DataPreference is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-smpte2038datapreference
	Smpte2038DataPreference cfz.Expression[string] `json:"Smpte2038DataPreference,omitempty"`

	// SourceEndBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-sourceendbehavior
	SourceEndBehavior cfz.Expression[string] `json:"SourceEndBehavior,omitempty"`

	// VideoSelector is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputsettings.html#cfn-medialive-channel-inputsettings-videoselector
	VideoSelector cfz.Expression[AWS_MediaLive_Channel_VideoSelector] `json:"VideoSelector,omitempty"`
}

// New__AWS_MediaLive_Channel_InputSettings initializes a new AWS_MediaLive_Channel_InputSettings.
func New__AWS_MediaLive_Channel_InputSettings() AWS_MediaLive_Channel_InputSettings {
	return AWS_MediaLive_Channel_InputSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_InputSettings) GetType() string {
	return AWS_MediaLive_Channel_InputSettings__Type
}

// Set__AudioSelectors updates property "AudioSelectors".
func (t AWS_MediaLive_Channel_InputSettings) Set__AudioSelectors(v cfz.ExpressionSlice[AWS_MediaLive_Channel_AudioSelector]) AWS_MediaLive_Channel_InputSettings {
	t.AudioSelectors = v
	return t
}

// SetS__AudioSelectors updates property "AudioSelectors".
func (t AWS_MediaLive_Channel_InputSettings) SetS__AudioSelectors(v ...cfz.Expression[AWS_MediaLive_Channel_AudioSelector]) AWS_MediaLive_Channel_InputSettings {
	t.AudioSelectors = cfz.S(v...)
	return t
}

// SetSV__AudioSelectors updates property "AudioSelectors".
func (t AWS_MediaLive_Channel_InputSettings) SetSV__AudioSelectors(v ...AWS_MediaLive_Channel_AudioSelector) AWS_MediaLive_Channel_InputSettings {
	t.AudioSelectors = cfz.SV(v...)
	return t
}

// Set__CaptionSelectors updates property "CaptionSelectors".
func (t AWS_MediaLive_Channel_InputSettings) Set__CaptionSelectors(v cfz.ExpressionSlice[AWS_MediaLive_Channel_CaptionSelector]) AWS_MediaLive_Channel_InputSettings {
	t.CaptionSelectors = v
	return t
}

// SetS__CaptionSelectors updates property "CaptionSelectors".
func (t AWS_MediaLive_Channel_InputSettings) SetS__CaptionSelectors(v ...cfz.Expression[AWS_MediaLive_Channel_CaptionSelector]) AWS_MediaLive_Channel_InputSettings {
	t.CaptionSelectors = cfz.S(v...)
	return t
}

// SetSV__CaptionSelectors updates property "CaptionSelectors".
func (t AWS_MediaLive_Channel_InputSettings) SetSV__CaptionSelectors(v ...AWS_MediaLive_Channel_CaptionSelector) AWS_MediaLive_Channel_InputSettings {
	t.CaptionSelectors = cfz.SV(v...)
	return t
}

// Set__DeblockFilter updates property "DeblockFilter".
func (t AWS_MediaLive_Channel_InputSettings) Set__DeblockFilter(v cfz.Expression[string]) AWS_MediaLive_Channel_InputSettings {
	t.DeblockFilter = v
	return t
}

// SetV__DeblockFilter updates property "DeblockFilter".
func (t AWS_MediaLive_Channel_InputSettings) SetV__DeblockFilter(v string) AWS_MediaLive_Channel_InputSettings {
	t.DeblockFilter = cfz.V(v)
	return t
}

// Set__DenoiseFilter updates property "DenoiseFilter".
func (t AWS_MediaLive_Channel_InputSettings) Set__DenoiseFilter(v cfz.Expression[string]) AWS_MediaLive_Channel_InputSettings {
	t.DenoiseFilter = v
	return t
}

// SetV__DenoiseFilter updates property "DenoiseFilter".
func (t AWS_MediaLive_Channel_InputSettings) SetV__DenoiseFilter(v string) AWS_MediaLive_Channel_InputSettings {
	t.DenoiseFilter = cfz.V(v)
	return t
}

// Set__FilterStrength updates property "FilterStrength".
func (t AWS_MediaLive_Channel_InputSettings) Set__FilterStrength(v cfz.Expression[int32]) AWS_MediaLive_Channel_InputSettings {
	t.FilterStrength = v
	return t
}

// SetV__FilterStrength updates property "FilterStrength".
func (t AWS_MediaLive_Channel_InputSettings) SetV__FilterStrength(v int32) AWS_MediaLive_Channel_InputSettings {
	t.FilterStrength = cfz.V(v)
	return t
}

// Set__InputFilter updates property "InputFilter".
func (t AWS_MediaLive_Channel_InputSettings) Set__InputFilter(v cfz.Expression[string]) AWS_MediaLive_Channel_InputSettings {
	t.InputFilter = v
	return t
}

// SetV__InputFilter updates property "InputFilter".
func (t AWS_MediaLive_Channel_InputSettings) SetV__InputFilter(v string) AWS_MediaLive_Channel_InputSettings {
	t.InputFilter = cfz.V(v)
	return t
}

// Set__NetworkInputSettings updates property "NetworkInputSettings".
func (t AWS_MediaLive_Channel_InputSettings) Set__NetworkInputSettings(v cfz.Expression[AWS_MediaLive_Channel_NetworkInputSettings]) AWS_MediaLive_Channel_InputSettings {
	t.NetworkInputSettings = v
	return t
}

// SetV__NetworkInputSettings updates property "NetworkInputSettings".
func (t AWS_MediaLive_Channel_InputSettings) SetV__NetworkInputSettings(v AWS_MediaLive_Channel_NetworkInputSettings) AWS_MediaLive_Channel_InputSettings {
	t.NetworkInputSettings = cfz.V(v)
	return t
}

// Set__Scte35Pid updates property "Scte35Pid".
func (t AWS_MediaLive_Channel_InputSettings) Set__Scte35Pid(v cfz.Expression[int32]) AWS_MediaLive_Channel_InputSettings {
	t.Scte35Pid = v
	return t
}

// SetV__Scte35Pid updates property "Scte35Pid".
func (t AWS_MediaLive_Channel_InputSettings) SetV__Scte35Pid(v int32) AWS_MediaLive_Channel_InputSettings {
	t.Scte35Pid = cfz.V(v)
	return t
}

// Set__Smpte2038DataPreference updates property "Smpte2038DataPreference".
func (t AWS_MediaLive_Channel_InputSettings) Set__Smpte2038DataPreference(v cfz.Expression[string]) AWS_MediaLive_Channel_InputSettings {
	t.Smpte2038DataPreference = v
	return t
}

// SetV__Smpte2038DataPreference updates property "Smpte2038DataPreference".
func (t AWS_MediaLive_Channel_InputSettings) SetV__Smpte2038DataPreference(v string) AWS_MediaLive_Channel_InputSettings {
	t.Smpte2038DataPreference = cfz.V(v)
	return t
}

// Set__SourceEndBehavior updates property "SourceEndBehavior".
func (t AWS_MediaLive_Channel_InputSettings) Set__SourceEndBehavior(v cfz.Expression[string]) AWS_MediaLive_Channel_InputSettings {
	t.SourceEndBehavior = v
	return t
}

// SetV__SourceEndBehavior updates property "SourceEndBehavior".
func (t AWS_MediaLive_Channel_InputSettings) SetV__SourceEndBehavior(v string) AWS_MediaLive_Channel_InputSettings {
	t.SourceEndBehavior = cfz.V(v)
	return t
}

// Set__VideoSelector updates property "VideoSelector".
func (t AWS_MediaLive_Channel_InputSettings) Set__VideoSelector(v cfz.Expression[AWS_MediaLive_Channel_VideoSelector]) AWS_MediaLive_Channel_InputSettings {
	t.VideoSelector = v
	return t
}

// SetV__VideoSelector updates property "VideoSelector".
func (t AWS_MediaLive_Channel_InputSettings) SetV__VideoSelector(v AWS_MediaLive_Channel_VideoSelector) AWS_MediaLive_Channel_InputSettings {
	t.VideoSelector = cfz.V(v)
	return t
}
