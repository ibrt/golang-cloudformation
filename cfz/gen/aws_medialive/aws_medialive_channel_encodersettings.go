// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_EncoderSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.EncoderSettings.
	AWS_MediaLive_Channel_EncoderSettings__Type = "AWS::MediaLive::Channel.EncoderSettings"
)

var (
	// AWS_MediaLive_Channel_EncoderSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.EncoderSettings.
	AWS_MediaLive_Channel_EncoderSettings__PropertiesMap = struct {
		AudioDescriptions           string
		AvailBlanking               string
		AvailConfiguration          string
		BlackoutSlate               string
		CaptionDescriptions         string
		ColorCorrectionSettings     string
		FeatureActivations          string
		GlobalConfiguration         string
		MotionGraphicsConfiguration string
		NielsenConfiguration        string
		OutputGroups                string
		ThumbnailConfiguration      string
		TimecodeConfig              string
		VideoDescriptions           string
	}{
		AudioDescriptions:           "AudioDescriptions",
		AvailBlanking:               "AvailBlanking",
		AvailConfiguration:          "AvailConfiguration",
		BlackoutSlate:               "BlackoutSlate",
		CaptionDescriptions:         "CaptionDescriptions",
		ColorCorrectionSettings:     "ColorCorrectionSettings",
		FeatureActivations:          "FeatureActivations",
		GlobalConfiguration:         "GlobalConfiguration",
		MotionGraphicsConfiguration: "MotionGraphicsConfiguration",
		NielsenConfiguration:        "NielsenConfiguration",
		OutputGroups:                "OutputGroups",
		ThumbnailConfiguration:      "ThumbnailConfiguration",
		TimecodeConfig:              "TimecodeConfig",
		VideoDescriptions:           "VideoDescriptions",
	}

	// AWS_MediaLive_Channel_EncoderSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.EncoderSettings.
	AWS_MediaLive_Channel_EncoderSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.AudioDescriptions,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.AvailBlanking,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.AvailConfiguration,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.BlackoutSlate,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.CaptionDescriptions,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.ColorCorrectionSettings,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.FeatureActivations,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.GlobalConfiguration,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.MotionGraphicsConfiguration,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.NielsenConfiguration,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.OutputGroups,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.ThumbnailConfiguration,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.TimecodeConfig,
		AWS_MediaLive_Channel_EncoderSettings__PropertiesMap.VideoDescriptions,
	}
)

// AWS_MediaLive_Channel_EncoderSettings is a binding for AWS::MediaLive::Channel.EncoderSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html
type AWS_MediaLive_Channel_EncoderSettings struct {
	// AudioDescriptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-audiodescriptions
	AudioDescriptions cfz.ExpressionSlice[AWS_MediaLive_Channel_AudioDescription] `json:"AudioDescriptions,omitempty"`

	// AvailBlanking is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availblanking
	AvailBlanking cfz.Expression[AWS_MediaLive_Channel_AvailBlanking] `json:"AvailBlanking,omitempty"`

	// AvailConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availconfiguration
	AvailConfiguration cfz.Expression[AWS_MediaLive_Channel_AvailConfiguration] `json:"AvailConfiguration,omitempty"`

	// BlackoutSlate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-blackoutslate
	BlackoutSlate cfz.Expression[AWS_MediaLive_Channel_BlackoutSlate] `json:"BlackoutSlate,omitempty"`

	// CaptionDescriptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-captiondescriptions
	CaptionDescriptions cfz.ExpressionSlice[AWS_MediaLive_Channel_CaptionDescription] `json:"CaptionDescriptions,omitempty"`

	// ColorCorrectionSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-colorcorrectionsettings
	ColorCorrectionSettings cfz.Expression[AWS_MediaLive_Channel_ColorCorrectionSettings] `json:"ColorCorrectionSettings,omitempty"`

	// FeatureActivations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-featureactivations
	FeatureActivations cfz.Expression[AWS_MediaLive_Channel_FeatureActivations] `json:"FeatureActivations,omitempty"`

	// GlobalConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-globalconfiguration
	GlobalConfiguration cfz.Expression[AWS_MediaLive_Channel_GlobalConfiguration] `json:"GlobalConfiguration,omitempty"`

	// MotionGraphicsConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-motiongraphicsconfiguration
	MotionGraphicsConfiguration cfz.Expression[AWS_MediaLive_Channel_MotionGraphicsConfiguration] `json:"MotionGraphicsConfiguration,omitempty"`

	// NielsenConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-nielsenconfiguration
	NielsenConfiguration cfz.Expression[AWS_MediaLive_Channel_NielsenConfiguration] `json:"NielsenConfiguration,omitempty"`

	// OutputGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-outputgroups
	OutputGroups cfz.ExpressionSlice[AWS_MediaLive_Channel_OutputGroup] `json:"OutputGroups,omitempty"`

	// ThumbnailConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-thumbnailconfiguration
	ThumbnailConfiguration cfz.Expression[AWS_MediaLive_Channel_ThumbnailConfiguration] `json:"ThumbnailConfiguration,omitempty"`

	// TimecodeConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-timecodeconfig
	TimecodeConfig cfz.Expression[AWS_MediaLive_Channel_TimecodeConfig] `json:"TimecodeConfig,omitempty"`

	// VideoDescriptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-videodescriptions
	VideoDescriptions cfz.ExpressionSlice[AWS_MediaLive_Channel_VideoDescription] `json:"VideoDescriptions,omitempty"`
}

// New__AWS_MediaLive_Channel_EncoderSettings initializes a new AWS_MediaLive_Channel_EncoderSettings.
func New__AWS_MediaLive_Channel_EncoderSettings() AWS_MediaLive_Channel_EncoderSettings {
	return AWS_MediaLive_Channel_EncoderSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_EncoderSettings) GetType() string {
	return AWS_MediaLive_Channel_EncoderSettings__Type
}

// Set__AudioDescriptions updates property "AudioDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__AudioDescriptions(v cfz.ExpressionSlice[AWS_MediaLive_Channel_AudioDescription]) AWS_MediaLive_Channel_EncoderSettings {
	t.AudioDescriptions = v
	return t
}

// SetS__AudioDescriptions updates property "AudioDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) SetS__AudioDescriptions(v ...cfz.Expression[AWS_MediaLive_Channel_AudioDescription]) AWS_MediaLive_Channel_EncoderSettings {
	t.AudioDescriptions = cfz.S(v...)
	return t
}

// SetSV__AudioDescriptions updates property "AudioDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) SetSV__AudioDescriptions(v ...AWS_MediaLive_Channel_AudioDescription) AWS_MediaLive_Channel_EncoderSettings {
	t.AudioDescriptions = cfz.SV(v...)
	return t
}

// Set__AvailBlanking updates property "AvailBlanking".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__AvailBlanking(v cfz.Expression[AWS_MediaLive_Channel_AvailBlanking]) AWS_MediaLive_Channel_EncoderSettings {
	t.AvailBlanking = v
	return t
}

// SetV__AvailBlanking updates property "AvailBlanking".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__AvailBlanking(v AWS_MediaLive_Channel_AvailBlanking) AWS_MediaLive_Channel_EncoderSettings {
	t.AvailBlanking = cfz.V(v)
	return t
}

// Set__AvailConfiguration updates property "AvailConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__AvailConfiguration(v cfz.Expression[AWS_MediaLive_Channel_AvailConfiguration]) AWS_MediaLive_Channel_EncoderSettings {
	t.AvailConfiguration = v
	return t
}

// SetV__AvailConfiguration updates property "AvailConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__AvailConfiguration(v AWS_MediaLive_Channel_AvailConfiguration) AWS_MediaLive_Channel_EncoderSettings {
	t.AvailConfiguration = cfz.V(v)
	return t
}

// Set__BlackoutSlate updates property "BlackoutSlate".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__BlackoutSlate(v cfz.Expression[AWS_MediaLive_Channel_BlackoutSlate]) AWS_MediaLive_Channel_EncoderSettings {
	t.BlackoutSlate = v
	return t
}

// SetV__BlackoutSlate updates property "BlackoutSlate".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__BlackoutSlate(v AWS_MediaLive_Channel_BlackoutSlate) AWS_MediaLive_Channel_EncoderSettings {
	t.BlackoutSlate = cfz.V(v)
	return t
}

// Set__CaptionDescriptions updates property "CaptionDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__CaptionDescriptions(v cfz.ExpressionSlice[AWS_MediaLive_Channel_CaptionDescription]) AWS_MediaLive_Channel_EncoderSettings {
	t.CaptionDescriptions = v
	return t
}

// SetS__CaptionDescriptions updates property "CaptionDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) SetS__CaptionDescriptions(v ...cfz.Expression[AWS_MediaLive_Channel_CaptionDescription]) AWS_MediaLive_Channel_EncoderSettings {
	t.CaptionDescriptions = cfz.S(v...)
	return t
}

// SetSV__CaptionDescriptions updates property "CaptionDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) SetSV__CaptionDescriptions(v ...AWS_MediaLive_Channel_CaptionDescription) AWS_MediaLive_Channel_EncoderSettings {
	t.CaptionDescriptions = cfz.SV(v...)
	return t
}

// Set__ColorCorrectionSettings updates property "ColorCorrectionSettings".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__ColorCorrectionSettings(v cfz.Expression[AWS_MediaLive_Channel_ColorCorrectionSettings]) AWS_MediaLive_Channel_EncoderSettings {
	t.ColorCorrectionSettings = v
	return t
}

// SetV__ColorCorrectionSettings updates property "ColorCorrectionSettings".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__ColorCorrectionSettings(v AWS_MediaLive_Channel_ColorCorrectionSettings) AWS_MediaLive_Channel_EncoderSettings {
	t.ColorCorrectionSettings = cfz.V(v)
	return t
}

// Set__FeatureActivations updates property "FeatureActivations".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__FeatureActivations(v cfz.Expression[AWS_MediaLive_Channel_FeatureActivations]) AWS_MediaLive_Channel_EncoderSettings {
	t.FeatureActivations = v
	return t
}

// SetV__FeatureActivations updates property "FeatureActivations".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__FeatureActivations(v AWS_MediaLive_Channel_FeatureActivations) AWS_MediaLive_Channel_EncoderSettings {
	t.FeatureActivations = cfz.V(v)
	return t
}

// Set__GlobalConfiguration updates property "GlobalConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__GlobalConfiguration(v cfz.Expression[AWS_MediaLive_Channel_GlobalConfiguration]) AWS_MediaLive_Channel_EncoderSettings {
	t.GlobalConfiguration = v
	return t
}

// SetV__GlobalConfiguration updates property "GlobalConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__GlobalConfiguration(v AWS_MediaLive_Channel_GlobalConfiguration) AWS_MediaLive_Channel_EncoderSettings {
	t.GlobalConfiguration = cfz.V(v)
	return t
}

// Set__MotionGraphicsConfiguration updates property "MotionGraphicsConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__MotionGraphicsConfiguration(v cfz.Expression[AWS_MediaLive_Channel_MotionGraphicsConfiguration]) AWS_MediaLive_Channel_EncoderSettings {
	t.MotionGraphicsConfiguration = v
	return t
}

// SetV__MotionGraphicsConfiguration updates property "MotionGraphicsConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__MotionGraphicsConfiguration(v AWS_MediaLive_Channel_MotionGraphicsConfiguration) AWS_MediaLive_Channel_EncoderSettings {
	t.MotionGraphicsConfiguration = cfz.V(v)
	return t
}

// Set__NielsenConfiguration updates property "NielsenConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__NielsenConfiguration(v cfz.Expression[AWS_MediaLive_Channel_NielsenConfiguration]) AWS_MediaLive_Channel_EncoderSettings {
	t.NielsenConfiguration = v
	return t
}

// SetV__NielsenConfiguration updates property "NielsenConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__NielsenConfiguration(v AWS_MediaLive_Channel_NielsenConfiguration) AWS_MediaLive_Channel_EncoderSettings {
	t.NielsenConfiguration = cfz.V(v)
	return t
}

// Set__OutputGroups updates property "OutputGroups".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__OutputGroups(v cfz.ExpressionSlice[AWS_MediaLive_Channel_OutputGroup]) AWS_MediaLive_Channel_EncoderSettings {
	t.OutputGroups = v
	return t
}

// SetS__OutputGroups updates property "OutputGroups".
func (t AWS_MediaLive_Channel_EncoderSettings) SetS__OutputGroups(v ...cfz.Expression[AWS_MediaLive_Channel_OutputGroup]) AWS_MediaLive_Channel_EncoderSettings {
	t.OutputGroups = cfz.S(v...)
	return t
}

// SetSV__OutputGroups updates property "OutputGroups".
func (t AWS_MediaLive_Channel_EncoderSettings) SetSV__OutputGroups(v ...AWS_MediaLive_Channel_OutputGroup) AWS_MediaLive_Channel_EncoderSettings {
	t.OutputGroups = cfz.SV(v...)
	return t
}

// Set__ThumbnailConfiguration updates property "ThumbnailConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__ThumbnailConfiguration(v cfz.Expression[AWS_MediaLive_Channel_ThumbnailConfiguration]) AWS_MediaLive_Channel_EncoderSettings {
	t.ThumbnailConfiguration = v
	return t
}

// SetV__ThumbnailConfiguration updates property "ThumbnailConfiguration".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__ThumbnailConfiguration(v AWS_MediaLive_Channel_ThumbnailConfiguration) AWS_MediaLive_Channel_EncoderSettings {
	t.ThumbnailConfiguration = cfz.V(v)
	return t
}

// Set__TimecodeConfig updates property "TimecodeConfig".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__TimecodeConfig(v cfz.Expression[AWS_MediaLive_Channel_TimecodeConfig]) AWS_MediaLive_Channel_EncoderSettings {
	t.TimecodeConfig = v
	return t
}

// SetV__TimecodeConfig updates property "TimecodeConfig".
func (t AWS_MediaLive_Channel_EncoderSettings) SetV__TimecodeConfig(v AWS_MediaLive_Channel_TimecodeConfig) AWS_MediaLive_Channel_EncoderSettings {
	t.TimecodeConfig = cfz.V(v)
	return t
}

// Set__VideoDescriptions updates property "VideoDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) Set__VideoDescriptions(v cfz.ExpressionSlice[AWS_MediaLive_Channel_VideoDescription]) AWS_MediaLive_Channel_EncoderSettings {
	t.VideoDescriptions = v
	return t
}

// SetS__VideoDescriptions updates property "VideoDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) SetS__VideoDescriptions(v ...cfz.Expression[AWS_MediaLive_Channel_VideoDescription]) AWS_MediaLive_Channel_EncoderSettings {
	t.VideoDescriptions = cfz.S(v...)
	return t
}

// SetSV__VideoDescriptions updates property "VideoDescriptions".
func (t AWS_MediaLive_Channel_EncoderSettings) SetSV__VideoDescriptions(v ...AWS_MediaLive_Channel_VideoDescription) AWS_MediaLive_Channel_EncoderSettings {
	t.VideoDescriptions = cfz.SV(v...)
	return t
}
