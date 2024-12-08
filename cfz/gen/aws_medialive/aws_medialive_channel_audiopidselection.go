// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_AudioPidSelection__Type is the CloudFormation type for AWS::MediaLive::Channel.AudioPidSelection.
	AWS_MediaLive_Channel_AudioPidSelection__Type = "AWS::MediaLive::Channel.AudioPidSelection"
)

var (
	// AWS_MediaLive_Channel_AudioPidSelection__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioPidSelection.
	AWS_MediaLive_Channel_AudioPidSelection__PropertiesMap = struct {
		Pid string
	}{
		Pid: "Pid",
	}

	// AWS_MediaLive_Channel_AudioPidSelection__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioPidSelection.
	AWS_MediaLive_Channel_AudioPidSelection__PropertiesSlice = []string{
		AWS_MediaLive_Channel_AudioPidSelection__PropertiesMap.Pid,
	}
)

// AWS_MediaLive_Channel_AudioPidSelection is a binding for AWS::MediaLive::Channel.AudioPidSelection.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html
type AWS_MediaLive_Channel_AudioPidSelection struct {
	// Pid is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiopidselection.html#cfn-medialive-channel-audiopidselection-pid
	Pid cfz.Expression[int32] `json:"Pid,omitempty"`
}

// New__AWS_MediaLive_Channel_AudioPidSelection initializes a new AWS_MediaLive_Channel_AudioPidSelection.
func New__AWS_MediaLive_Channel_AudioPidSelection() AWS_MediaLive_Channel_AudioPidSelection {
	return AWS_MediaLive_Channel_AudioPidSelection{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_AudioPidSelection) GetType() string {
	return AWS_MediaLive_Channel_AudioPidSelection__Type
}

// Set__Pid updates property "Pid".
func (t AWS_MediaLive_Channel_AudioPidSelection) Set__Pid(v cfz.Expression[int32]) AWS_MediaLive_Channel_AudioPidSelection {
	t.Pid = v
	return t
}

// SetV__Pid updates property "Pid".
func (t AWS_MediaLive_Channel_AudioPidSelection) SetV__Pid(v int32) AWS_MediaLive_Channel_AudioPidSelection {
	t.Pid = cfz.V(v)
	return t
}
