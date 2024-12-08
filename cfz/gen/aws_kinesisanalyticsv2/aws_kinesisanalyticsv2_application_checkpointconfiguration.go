// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::Application.CheckpointConfiguration.
	AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__Type = "AWS::KinesisAnalyticsV2::Application.CheckpointConfiguration"
)

var (
	// AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.CheckpointConfiguration.
	AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesMap = struct {
		CheckpointInterval         string
		CheckpointingEnabled       string
		ConfigurationType          string
		MinPauseBetweenCheckpoints string
	}{
		CheckpointInterval:         "CheckpointInterval",
		CheckpointingEnabled:       "CheckpointingEnabled",
		ConfigurationType:          "ConfigurationType",
		MinPauseBetweenCheckpoints: "MinPauseBetweenCheckpoints",
	}

	// AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.CheckpointConfiguration.
	AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesMap.CheckpointInterval,
		AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesMap.CheckpointingEnabled,
		AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesMap.ConfigurationType,
		AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__PropertiesMap.MinPauseBetweenCheckpoints,
	}
)

// AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration is a binding for AWS::KinesisAnalyticsV2::Application.CheckpointConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html
type AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration struct {
	// CheckpointInterval is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-checkpointinterval
	CheckpointInterval cfz.Expression[int32] `json:"CheckpointInterval,omitempty"`

	// CheckpointingEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-checkpointingenabled
	CheckpointingEnabled cfz.Expression[bool] `json:"CheckpointingEnabled,omitempty"`

	// ConfigurationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-configurationtype
	ConfigurationType cfz.Expression[string] `json:"ConfigurationType,omitempty"`

	// MinPauseBetweenCheckpoints is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.html#cfn-kinesisanalyticsv2-application-checkpointconfiguration-minpausebetweencheckpoints
	MinPauseBetweenCheckpoints cfz.Expression[int32] `json:"MinPauseBetweenCheckpoints,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration initializes a new AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration.
func New__AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration() AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	return AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) GetType() string {
	return AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration__Type
}

// Set__CheckpointInterval updates property "CheckpointInterval".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) Set__CheckpointInterval(v cfz.Expression[int32]) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.CheckpointInterval = v
	return t
}

// SetV__CheckpointInterval updates property "CheckpointInterval".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) SetV__CheckpointInterval(v int32) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.CheckpointInterval = cfz.V(v)
	return t
}

// Set__CheckpointingEnabled updates property "CheckpointingEnabled".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) Set__CheckpointingEnabled(v cfz.Expression[bool]) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.CheckpointingEnabled = v
	return t
}

// SetV__CheckpointingEnabled updates property "CheckpointingEnabled".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) SetV__CheckpointingEnabled(v bool) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.CheckpointingEnabled = cfz.V(v)
	return t
}

// Set__ConfigurationType updates property "ConfigurationType".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) Set__ConfigurationType(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.ConfigurationType = v
	return t
}

// SetV__ConfigurationType updates property "ConfigurationType".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) SetV__ConfigurationType(v string) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.ConfigurationType = cfz.V(v)
	return t
}

// Set__MinPauseBetweenCheckpoints updates property "MinPauseBetweenCheckpoints".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) Set__MinPauseBetweenCheckpoints(v cfz.Expression[int32]) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.MinPauseBetweenCheckpoints = v
	return t
}

// SetV__MinPauseBetweenCheckpoints updates property "MinPauseBetweenCheckpoints".
func (t AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration) SetV__MinPauseBetweenCheckpoints(v int32) AWS_KinesisAnalyticsV2_Application_CheckpointConfiguration {
	t.MinPauseBetweenCheckpoints = cfz.V(v)
	return t
}
