// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_OutputDestinationSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.OutputDestinationSettings.
	AWS_MediaLive_Channel_OutputDestinationSettings__Type = "AWS::MediaLive::Channel.OutputDestinationSettings"
)

var (
	// AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.OutputDestinationSettings.
	AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesMap = struct {
		PasswordParam string
		StreamName    string
		Url           string
		Username      string
	}{
		PasswordParam: "PasswordParam",
		StreamName:    "StreamName",
		Url:           "Url",
		Username:      "Username",
	}

	// AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.OutputDestinationSettings.
	AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesMap.PasswordParam,
		AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesMap.StreamName,
		AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesMap.Url,
		AWS_MediaLive_Channel_OutputDestinationSettings__PropertiesMap.Username,
	}
)

// AWS_MediaLive_Channel_OutputDestinationSettings is a binding for AWS::MediaLive::Channel.OutputDestinationSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html
type AWS_MediaLive_Channel_OutputDestinationSettings struct {
	// PasswordParam is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-passwordparam
	PasswordParam cfz.Expression[string] `json:"PasswordParam,omitempty"`

	// StreamName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-streamname
	StreamName cfz.Expression[string] `json:"StreamName,omitempty"`

	// Url is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-url
	Url cfz.Expression[string] `json:"Url,omitempty"`

	// Username is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-username
	Username cfz.Expression[string] `json:"Username,omitempty"`
}

// New__AWS_MediaLive_Channel_OutputDestinationSettings initializes a new AWS_MediaLive_Channel_OutputDestinationSettings.
func New__AWS_MediaLive_Channel_OutputDestinationSettings() AWS_MediaLive_Channel_OutputDestinationSettings {
	return AWS_MediaLive_Channel_OutputDestinationSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_OutputDestinationSettings) GetType() string {
	return AWS_MediaLive_Channel_OutputDestinationSettings__Type
}

// Set__PasswordParam updates property "PasswordParam".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) Set__PasswordParam(v cfz.Expression[string]) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.PasswordParam = v
	return t
}

// SetV__PasswordParam updates property "PasswordParam".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) SetV__PasswordParam(v string) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.PasswordParam = cfz.V(v)
	return t
}

// Set__StreamName updates property "StreamName".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) Set__StreamName(v cfz.Expression[string]) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.StreamName = v
	return t
}

// SetV__StreamName updates property "StreamName".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) SetV__StreamName(v string) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.StreamName = cfz.V(v)
	return t
}

// Set__Url updates property "Url".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) Set__Url(v cfz.Expression[string]) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.Url = v
	return t
}

// SetV__Url updates property "Url".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) SetV__Url(v string) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.Url = cfz.V(v)
	return t
}

// Set__Username updates property "Username".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) Set__Username(v cfz.Expression[string]) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.Username = v
	return t
}

// SetV__Username updates property "Username".
func (t AWS_MediaLive_Channel_OutputDestinationSettings) SetV__Username(v string) AWS_MediaLive_Channel_OutputDestinationSettings {
	t.Username = cfz.V(v)
	return t
}
