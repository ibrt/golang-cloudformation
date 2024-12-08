// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_CmafIngestOutputSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.CmafIngestOutputSettings.
	AWS_MediaLive_Channel_CmafIngestOutputSettings__Type = "AWS::MediaLive::Channel.CmafIngestOutputSettings"
)

var (
	// AWS_MediaLive_Channel_CmafIngestOutputSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.CmafIngestOutputSettings.
	AWS_MediaLive_Channel_CmafIngestOutputSettings__PropertiesMap = struct {
		NameModifier string
	}{
		NameModifier: "NameModifier",
	}

	// AWS_MediaLive_Channel_CmafIngestOutputSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.CmafIngestOutputSettings.
	AWS_MediaLive_Channel_CmafIngestOutputSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_CmafIngestOutputSettings__PropertiesMap.NameModifier,
	}
)

// AWS_MediaLive_Channel_CmafIngestOutputSettings is a binding for AWS::MediaLive::Channel.CmafIngestOutputSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestoutputsettings.html
type AWS_MediaLive_Channel_CmafIngestOutputSettings struct {
	// NameModifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-cmafingestoutputsettings.html#cfn-medialive-channel-cmafingestoutputsettings-namemodifier
	NameModifier cfz.Expression[string] `json:"NameModifier,omitempty"`
}

// New__AWS_MediaLive_Channel_CmafIngestOutputSettings initializes a new AWS_MediaLive_Channel_CmafIngestOutputSettings.
func New__AWS_MediaLive_Channel_CmafIngestOutputSettings() AWS_MediaLive_Channel_CmafIngestOutputSettings {
	return AWS_MediaLive_Channel_CmafIngestOutputSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_CmafIngestOutputSettings) GetType() string {
	return AWS_MediaLive_Channel_CmafIngestOutputSettings__Type
}

// Set__NameModifier updates property "NameModifier".
func (t AWS_MediaLive_Channel_CmafIngestOutputSettings) Set__NameModifier(v cfz.Expression[string]) AWS_MediaLive_Channel_CmafIngestOutputSettings {
	t.NameModifier = v
	return t
}

// SetV__NameModifier updates property "NameModifier".
func (t AWS_MediaLive_Channel_CmafIngestOutputSettings) SetV__NameModifier(v string) AWS_MediaLive_Channel_CmafIngestOutputSettings {
	t.NameModifier = cfz.V(v)
	return t
}
