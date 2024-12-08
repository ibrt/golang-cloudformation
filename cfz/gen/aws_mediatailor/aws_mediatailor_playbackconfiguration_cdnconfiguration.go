// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediatailor

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__Type is the CloudFormation type for AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration.
	AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__Type = "AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration"
)

var (
	// AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration.
	AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__PropertiesMap = struct {
		AdSegmentUrlPrefix      string
		ContentSegmentUrlPrefix string
	}{
		AdSegmentUrlPrefix:      "AdSegmentUrlPrefix",
		ContentSegmentUrlPrefix: "ContentSegmentUrlPrefix",
	}

	// AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration.
	AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__PropertiesSlice = []string{
		AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__PropertiesMap.AdSegmentUrlPrefix,
		AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__PropertiesMap.ContentSegmentUrlPrefix,
	}
)

// AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration is a binding for AWS::MediaTailor::PlaybackConfiguration.CdnConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html
type AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration struct {
	// AdSegmentUrlPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html#cfn-mediatailor-playbackconfiguration-cdnconfiguration-adsegmenturlprefix
	AdSegmentUrlPrefix cfz.Expression[string] `json:"AdSegmentUrlPrefix,omitempty"`

	// ContentSegmentUrlPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html#cfn-mediatailor-playbackconfiguration-cdnconfiguration-contentsegmenturlprefix
	ContentSegmentUrlPrefix cfz.Expression[string] `json:"ContentSegmentUrlPrefix,omitempty"`
}

// New__AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration initializes a new AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration.
func New__AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration() AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration {
	return AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration) GetType() string {
	return AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration__Type
}

// Set__AdSegmentUrlPrefix updates property "AdSegmentUrlPrefix".
func (t AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration) Set__AdSegmentUrlPrefix(v cfz.Expression[string]) AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration {
	t.AdSegmentUrlPrefix = v
	return t
}

// SetV__AdSegmentUrlPrefix updates property "AdSegmentUrlPrefix".
func (t AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration) SetV__AdSegmentUrlPrefix(v string) AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration {
	t.AdSegmentUrlPrefix = cfz.V(v)
	return t
}

// Set__ContentSegmentUrlPrefix updates property "ContentSegmentUrlPrefix".
func (t AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration) Set__ContentSegmentUrlPrefix(v cfz.Expression[string]) AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration {
	t.ContentSegmentUrlPrefix = v
	return t
}

// SetV__ContentSegmentUrlPrefix updates property "ContentSegmentUrlPrefix".
func (t AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration) SetV__ContentSegmentUrlPrefix(v string) AWS_MediaTailor_PlaybackConfiguration_CdnConfiguration {
	t.ContentSegmentUrlPrefix = cfz.V(v)
	return t
}
