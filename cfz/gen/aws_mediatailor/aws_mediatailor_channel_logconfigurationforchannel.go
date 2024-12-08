// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediatailor

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaTailor_Channel_LogConfigurationForChannel__Type is the CloudFormation type for AWS::MediaTailor::Channel.LogConfigurationForChannel.
	AWS_MediaTailor_Channel_LogConfigurationForChannel__Type = "AWS::MediaTailor::Channel.LogConfigurationForChannel"
)

var (
	// AWS_MediaTailor_Channel_LogConfigurationForChannel__PropertiesMap reports all the CloudFormation properties for AWS::MediaTailor::Channel.LogConfigurationForChannel.
	AWS_MediaTailor_Channel_LogConfigurationForChannel__PropertiesMap = struct {
		LogTypes string
	}{
		LogTypes: "LogTypes",
	}

	// AWS_MediaTailor_Channel_LogConfigurationForChannel__PropertiesSlice reports all the CloudFormation properties for AWS::MediaTailor::Channel.LogConfigurationForChannel.
	AWS_MediaTailor_Channel_LogConfigurationForChannel__PropertiesSlice = []string{
		AWS_MediaTailor_Channel_LogConfigurationForChannel__PropertiesMap.LogTypes,
	}
)

// AWS_MediaTailor_Channel_LogConfigurationForChannel is a binding for AWS::MediaTailor::Channel.LogConfigurationForChannel.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-logconfigurationforchannel.html
type AWS_MediaTailor_Channel_LogConfigurationForChannel struct {
	// LogTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-channel-logconfigurationforchannel.html#cfn-mediatailor-channel-logconfigurationforchannel-logtypes
	LogTypes cfz.ExpressionSlice[string] `json:"LogTypes,omitempty"`
}

// New__AWS_MediaTailor_Channel_LogConfigurationForChannel initializes a new AWS_MediaTailor_Channel_LogConfigurationForChannel.
func New__AWS_MediaTailor_Channel_LogConfigurationForChannel() AWS_MediaTailor_Channel_LogConfigurationForChannel {
	return AWS_MediaTailor_Channel_LogConfigurationForChannel{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaTailor_Channel_LogConfigurationForChannel) GetType() string {
	return AWS_MediaTailor_Channel_LogConfigurationForChannel__Type
}

// Set__LogTypes updates property "LogTypes".
func (t AWS_MediaTailor_Channel_LogConfigurationForChannel) Set__LogTypes(v cfz.ExpressionSlice[string]) AWS_MediaTailor_Channel_LogConfigurationForChannel {
	t.LogTypes = v
	return t
}

// SetS__LogTypes updates property "LogTypes".
func (t AWS_MediaTailor_Channel_LogConfigurationForChannel) SetS__LogTypes(v ...cfz.Expression[string]) AWS_MediaTailor_Channel_LogConfigurationForChannel {
	t.LogTypes = cfz.S(v...)
	return t
}

// SetSV__LogTypes updates property "LogTypes".
func (t AWS_MediaTailor_Channel_LogConfigurationForChannel) SetSV__LogTypes(v ...string) AWS_MediaTailor_Channel_LogConfigurationForChannel {
	t.LogTypes = cfz.SV(v...)
	return t
}
