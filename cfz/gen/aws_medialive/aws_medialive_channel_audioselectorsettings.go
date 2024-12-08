// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_AudioSelectorSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.AudioSelectorSettings.
	AWS_MediaLive_Channel_AudioSelectorSettings__Type = "AWS::MediaLive::Channel.AudioSelectorSettings"
)

var (
	// AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioSelectorSettings.
	AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesMap = struct {
		AudioHlsRenditionSelection string
		AudioLanguageSelection     string
		AudioPidSelection          string
		AudioTrackSelection        string
	}{
		AudioHlsRenditionSelection: "AudioHlsRenditionSelection",
		AudioLanguageSelection:     "AudioLanguageSelection",
		AudioPidSelection:          "AudioPidSelection",
		AudioTrackSelection:        "AudioTrackSelection",
	}

	// AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioSelectorSettings.
	AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesMap.AudioHlsRenditionSelection,
		AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesMap.AudioLanguageSelection,
		AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesMap.AudioPidSelection,
		AWS_MediaLive_Channel_AudioSelectorSettings__PropertiesMap.AudioTrackSelection,
	}
)

// AWS_MediaLive_Channel_AudioSelectorSettings is a binding for AWS::MediaLive::Channel.AudioSelectorSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html
type AWS_MediaLive_Channel_AudioSelectorSettings struct {
	// AudioHlsRenditionSelection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiohlsrenditionselection
	AudioHlsRenditionSelection cfz.Expression[AWS_MediaLive_Channel_AudioHlsRenditionSelection] `json:"AudioHlsRenditionSelection,omitempty"`

	// AudioLanguageSelection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiolanguageselection
	AudioLanguageSelection cfz.Expression[AWS_MediaLive_Channel_AudioLanguageSelection] `json:"AudioLanguageSelection,omitempty"`

	// AudioPidSelection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiopidselection
	AudioPidSelection cfz.Expression[AWS_MediaLive_Channel_AudioPidSelection] `json:"AudioPidSelection,omitempty"`

	// AudioTrackSelection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audioselectorsettings.html#cfn-medialive-channel-audioselectorsettings-audiotrackselection
	AudioTrackSelection cfz.Expression[AWS_MediaLive_Channel_AudioTrackSelection] `json:"AudioTrackSelection,omitempty"`
}

// New__AWS_MediaLive_Channel_AudioSelectorSettings initializes a new AWS_MediaLive_Channel_AudioSelectorSettings.
func New__AWS_MediaLive_Channel_AudioSelectorSettings() AWS_MediaLive_Channel_AudioSelectorSettings {
	return AWS_MediaLive_Channel_AudioSelectorSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_AudioSelectorSettings) GetType() string {
	return AWS_MediaLive_Channel_AudioSelectorSettings__Type
}

// Set__AudioHlsRenditionSelection updates property "AudioHlsRenditionSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) Set__AudioHlsRenditionSelection(v cfz.Expression[AWS_MediaLive_Channel_AudioHlsRenditionSelection]) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioHlsRenditionSelection = v
	return t
}

// SetV__AudioHlsRenditionSelection updates property "AudioHlsRenditionSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) SetV__AudioHlsRenditionSelection(v AWS_MediaLive_Channel_AudioHlsRenditionSelection) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioHlsRenditionSelection = cfz.V(v)
	return t
}

// Set__AudioLanguageSelection updates property "AudioLanguageSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) Set__AudioLanguageSelection(v cfz.Expression[AWS_MediaLive_Channel_AudioLanguageSelection]) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioLanguageSelection = v
	return t
}

// SetV__AudioLanguageSelection updates property "AudioLanguageSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) SetV__AudioLanguageSelection(v AWS_MediaLive_Channel_AudioLanguageSelection) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioLanguageSelection = cfz.V(v)
	return t
}

// Set__AudioPidSelection updates property "AudioPidSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) Set__AudioPidSelection(v cfz.Expression[AWS_MediaLive_Channel_AudioPidSelection]) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioPidSelection = v
	return t
}

// SetV__AudioPidSelection updates property "AudioPidSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) SetV__AudioPidSelection(v AWS_MediaLive_Channel_AudioPidSelection) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioPidSelection = cfz.V(v)
	return t
}

// Set__AudioTrackSelection updates property "AudioTrackSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) Set__AudioTrackSelection(v cfz.Expression[AWS_MediaLive_Channel_AudioTrackSelection]) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioTrackSelection = v
	return t
}

// SetV__AudioTrackSelection updates property "AudioTrackSelection".
func (t AWS_MediaLive_Channel_AudioSelectorSettings) SetV__AudioTrackSelection(v AWS_MediaLive_Channel_AudioTrackSelection) AWS_MediaLive_Channel_AudioSelectorSettings {
	t.AudioTrackSelection = cfz.V(v)
	return t
}
