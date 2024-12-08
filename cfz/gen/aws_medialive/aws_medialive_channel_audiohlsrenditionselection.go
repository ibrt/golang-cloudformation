// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_AudioHlsRenditionSelection__Type is the CloudFormation type for AWS::MediaLive::Channel.AudioHlsRenditionSelection.
	AWS_MediaLive_Channel_AudioHlsRenditionSelection__Type = "AWS::MediaLive::Channel.AudioHlsRenditionSelection"
)

var (
	// AWS_MediaLive_Channel_AudioHlsRenditionSelection__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioHlsRenditionSelection.
	AWS_MediaLive_Channel_AudioHlsRenditionSelection__PropertiesMap = struct {
		GroupId string
		Name    string
	}{
		GroupId: "GroupId",
		Name:    "Name",
	}

	// AWS_MediaLive_Channel_AudioHlsRenditionSelection__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.AudioHlsRenditionSelection.
	AWS_MediaLive_Channel_AudioHlsRenditionSelection__PropertiesSlice = []string{
		AWS_MediaLive_Channel_AudioHlsRenditionSelection__PropertiesMap.GroupId,
		AWS_MediaLive_Channel_AudioHlsRenditionSelection__PropertiesMap.Name,
	}
)

// AWS_MediaLive_Channel_AudioHlsRenditionSelection is a binding for AWS::MediaLive::Channel.AudioHlsRenditionSelection.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiohlsrenditionselection.html
type AWS_MediaLive_Channel_AudioHlsRenditionSelection struct {
	// GroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiohlsrenditionselection.html#cfn-medialive-channel-audiohlsrenditionselection-groupid
	GroupId cfz.Expression[string] `json:"GroupId,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiohlsrenditionselection.html#cfn-medialive-channel-audiohlsrenditionselection-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_MediaLive_Channel_AudioHlsRenditionSelection initializes a new AWS_MediaLive_Channel_AudioHlsRenditionSelection.
func New__AWS_MediaLive_Channel_AudioHlsRenditionSelection() AWS_MediaLive_Channel_AudioHlsRenditionSelection {
	return AWS_MediaLive_Channel_AudioHlsRenditionSelection{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_AudioHlsRenditionSelection) GetType() string {
	return AWS_MediaLive_Channel_AudioHlsRenditionSelection__Type
}

// Set__GroupId updates property "GroupId".
func (t AWS_MediaLive_Channel_AudioHlsRenditionSelection) Set__GroupId(v cfz.Expression[string]) AWS_MediaLive_Channel_AudioHlsRenditionSelection {
	t.GroupId = v
	return t
}

// SetV__GroupId updates property "GroupId".
func (t AWS_MediaLive_Channel_AudioHlsRenditionSelection) SetV__GroupId(v string) AWS_MediaLive_Channel_AudioHlsRenditionSelection {
	t.GroupId = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_MediaLive_Channel_AudioHlsRenditionSelection) Set__Name(v cfz.Expression[string]) AWS_MediaLive_Channel_AudioHlsRenditionSelection {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_MediaLive_Channel_AudioHlsRenditionSelection) SetV__Name(v string) AWS_MediaLive_Channel_AudioHlsRenditionSelection {
	t.Name = cfz.V(v)
	return t
}
