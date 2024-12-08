// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_Scte27SourceSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.Scte27SourceSettings.
	AWS_MediaLive_Channel_Scte27SourceSettings__Type = "AWS::MediaLive::Channel.Scte27SourceSettings"
)

var (
	// AWS_MediaLive_Channel_Scte27SourceSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.Scte27SourceSettings.
	AWS_MediaLive_Channel_Scte27SourceSettings__PropertiesMap = struct {
		OcrLanguage string
		Pid         string
	}{
		OcrLanguage: "OcrLanguage",
		Pid:         "Pid",
	}

	// AWS_MediaLive_Channel_Scte27SourceSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.Scte27SourceSettings.
	AWS_MediaLive_Channel_Scte27SourceSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_Scte27SourceSettings__PropertiesMap.OcrLanguage,
		AWS_MediaLive_Channel_Scte27SourceSettings__PropertiesMap.Pid,
	}
)

// AWS_MediaLive_Channel_Scte27SourceSettings is a binding for AWS::MediaLive::Channel.Scte27SourceSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte27sourcesettings.html
type AWS_MediaLive_Channel_Scte27SourceSettings struct {
	// OcrLanguage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte27sourcesettings.html#cfn-medialive-channel-scte27sourcesettings-ocrlanguage
	OcrLanguage cfz.Expression[string] `json:"OcrLanguage,omitempty"`

	// Pid is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte27sourcesettings.html#cfn-medialive-channel-scte27sourcesettings-pid
	Pid cfz.Expression[int32] `json:"Pid,omitempty"`
}

// New__AWS_MediaLive_Channel_Scte27SourceSettings initializes a new AWS_MediaLive_Channel_Scte27SourceSettings.
func New__AWS_MediaLive_Channel_Scte27SourceSettings() AWS_MediaLive_Channel_Scte27SourceSettings {
	return AWS_MediaLive_Channel_Scte27SourceSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_Scte27SourceSettings) GetType() string {
	return AWS_MediaLive_Channel_Scte27SourceSettings__Type
}

// Set__OcrLanguage updates property "OcrLanguage".
func (t AWS_MediaLive_Channel_Scte27SourceSettings) Set__OcrLanguage(v cfz.Expression[string]) AWS_MediaLive_Channel_Scte27SourceSettings {
	t.OcrLanguage = v
	return t
}

// SetV__OcrLanguage updates property "OcrLanguage".
func (t AWS_MediaLive_Channel_Scte27SourceSettings) SetV__OcrLanguage(v string) AWS_MediaLive_Channel_Scte27SourceSettings {
	t.OcrLanguage = cfz.V(v)
	return t
}

// Set__Pid updates property "Pid".
func (t AWS_MediaLive_Channel_Scte27SourceSettings) Set__Pid(v cfz.Expression[int32]) AWS_MediaLive_Channel_Scte27SourceSettings {
	t.Pid = v
	return t
}

// SetV__Pid updates property "Pid".
func (t AWS_MediaLive_Channel_Scte27SourceSettings) SetV__Pid(v int32) AWS_MediaLive_Channel_Scte27SourceSettings {
	t.Pid = cfz.V(v)
	return t
}
