// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaLive_Channel_ArchiveContainerSettings__Type is the CloudFormation type for AWS::MediaLive::Channel.ArchiveContainerSettings.
	AWS_MediaLive_Channel_ArchiveContainerSettings__Type = "AWS::MediaLive::Channel.ArchiveContainerSettings"
)

var (
	// AWS_MediaLive_Channel_ArchiveContainerSettings__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.ArchiveContainerSettings.
	AWS_MediaLive_Channel_ArchiveContainerSettings__PropertiesMap = struct {
		M2tsSettings string
		RawSettings  string
	}{
		M2tsSettings: "M2tsSettings",
		RawSettings:  "RawSettings",
	}

	// AWS_MediaLive_Channel_ArchiveContainerSettings__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.ArchiveContainerSettings.
	AWS_MediaLive_Channel_ArchiveContainerSettings__PropertiesSlice = []string{
		AWS_MediaLive_Channel_ArchiveContainerSettings__PropertiesMap.M2tsSettings,
		AWS_MediaLive_Channel_ArchiveContainerSettings__PropertiesMap.RawSettings,
	}
)

// AWS_MediaLive_Channel_ArchiveContainerSettings is a binding for AWS::MediaLive::Channel.ArchiveContainerSettings.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html
type AWS_MediaLive_Channel_ArchiveContainerSettings struct {
	// M2tsSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html#cfn-medialive-channel-archivecontainersettings-m2tssettings
	M2tsSettings cfz.Expression[AWS_MediaLive_Channel_M2tsSettings] `json:"M2tsSettings,omitempty"`

	// RawSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html#cfn-medialive-channel-archivecontainersettings-rawsettings
	RawSettings cfz.Expression[AWS_MediaLive_Channel_RawSettings] `json:"RawSettings,omitempty"`
}

// New__AWS_MediaLive_Channel_ArchiveContainerSettings initializes a new AWS_MediaLive_Channel_ArchiveContainerSettings.
func New__AWS_MediaLive_Channel_ArchiveContainerSettings() AWS_MediaLive_Channel_ArchiveContainerSettings {
	return AWS_MediaLive_Channel_ArchiveContainerSettings{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaLive_Channel_ArchiveContainerSettings) GetType() string {
	return AWS_MediaLive_Channel_ArchiveContainerSettings__Type
}

// Set__M2tsSettings updates property "M2tsSettings".
func (t AWS_MediaLive_Channel_ArchiveContainerSettings) Set__M2tsSettings(v cfz.Expression[AWS_MediaLive_Channel_M2tsSettings]) AWS_MediaLive_Channel_ArchiveContainerSettings {
	t.M2tsSettings = v
	return t
}

// SetV__M2tsSettings updates property "M2tsSettings".
func (t AWS_MediaLive_Channel_ArchiveContainerSettings) SetV__M2tsSettings(v AWS_MediaLive_Channel_M2tsSettings) AWS_MediaLive_Channel_ArchiveContainerSettings {
	t.M2tsSettings = cfz.V(v)
	return t
}

// Set__RawSettings updates property "RawSettings".
func (t AWS_MediaLive_Channel_ArchiveContainerSettings) Set__RawSettings(v cfz.Expression[AWS_MediaLive_Channel_RawSettings]) AWS_MediaLive_Channel_ArchiveContainerSettings {
	t.RawSettings = v
	return t
}

// SetV__RawSettings updates property "RawSettings".
func (t AWS_MediaLive_Channel_ArchiveContainerSettings) SetV__RawSettings(v AWS_MediaLive_Channel_RawSettings) AWS_MediaLive_Channel_ArchiveContainerSettings {
	t.RawSettings = cfz.V(v)
	return t
}
