// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_EbuTtDDestinationSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.EbuTtDDestinationSettings.
	AWS_MediaLive_Channel_EbuTtDDestinationSettings__Type = "AWS::MediaLive::Channel.EbuTtDDestinationSettings"
)

var (
	// AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.EbuTtDDestinationSettings.
	AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesMap = struct {
		CopyrightHolder string
		FillLineGap     string
		FontFamily      string
		StyleControl    string
	}{
		CopyrightHolder: "CopyrightHolder",
		FillLineGap:     "FillLineGap",
		FontFamily:      "FontFamily",
		StyleControl:    "StyleControl",
	}

	// AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.EbuTtDDestinationSettings.
	AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesMap.CopyrightHolder,
		AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesMap.FillLineGap,
		AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesMap.FontFamily,
		AWS_MediaLive_Channel_EbuTtDDestinationSettings__PropertiesMap.StyleControl,
	}
)

// AWS_MediaLive_Channel_EbuTtDDestinationSettings is a binding for AWS::MediaLive::Channel.EbuTtDDestinationSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html
type AWS_MediaLive_Channel_EbuTtDDestinationSettings struct {
	// CopyrightHolder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-copyrightholder
	CopyrightHolder cfz.Expression[string] `json:"CopyrightHolder,omitempty"`

	// FillLineGap is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-filllinegap
	FillLineGap cfz.Expression[string] `json:"FillLineGap,omitempty"`

	// FontFamily is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-fontfamily
	FontFamily cfz.Expression[string] `json:"FontFamily,omitempty"`

	// StyleControl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-ebuttddestinationsettings.html#cfn-medialive-channel-ebuttddestinationsettings-stylecontrol
	StyleControl cfz.Expression[string] `json:"StyleControl,omitempty"`
}

// New__AWS_MediaLive_Channel_EbuTtDDestinationSettings initializes a new AWS_MediaLive_Channel_EbuTtDDestinationSettings.
func New__AWS_MediaLive_Channel_EbuTtDDestinationSettings() AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	return AWS_MediaLive_Channel_EbuTtDDestinationSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_EbuTtDDestinationSettings) GetType() string {
	return AWS_MediaLive_Channel_EbuTtDDestinationSettings__Type
}

// Set__CopyrightHolder updates property "CopyrightHolder".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) Set__CopyrightHolder(v cfz.Expression[string]) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.CopyrightHolder = v
	return t
}

// SetV__CopyrightHolder updates property "CopyrightHolder".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) SetV__CopyrightHolder(v string) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.CopyrightHolder = cfz.V(v)
	return t
}

// Set__FillLineGap updates property "FillLineGap".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) Set__FillLineGap(v cfz.Expression[string]) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.FillLineGap = v
	return t
}

// SetV__FillLineGap updates property "FillLineGap".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) SetV__FillLineGap(v string) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.FillLineGap = cfz.V(v)
	return t
}

// Set__FontFamily updates property "FontFamily".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) Set__FontFamily(v cfz.Expression[string]) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.FontFamily = v
	return t
}

// SetV__FontFamily updates property "FontFamily".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) SetV__FontFamily(v string) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.FontFamily = cfz.V(v)
	return t
}

// Set__StyleControl updates property "StyleControl".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) Set__StyleControl(v cfz.Expression[string]) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.StyleControl = v
	return t
}

// SetV__StyleControl updates property "StyleControl".
func (t AWS_MediaLive_Channel_EbuTtDDestinationSettings) SetV__StyleControl(v string) AWS_MediaLive_Channel_EbuTtDDestinationSettings {
	t.StyleControl = cfz.V(v)
	return t
}
