// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::Application.MonitoringConfiguration.
	AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__Type = "AWS::KinesisAnalyticsV2::Application.MonitoringConfiguration"
)

var (
	// AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.MonitoringConfiguration.
	AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesMap = struct {
		ConfigurationType string
		LogLevel          string
		MetricsLevel      string
	}{
		ConfigurationType: "ConfigurationType",
		LogLevel:          "LogLevel",
		MetricsLevel:      "MetricsLevel",
	}

	// AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.MonitoringConfiguration.
	AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesMap.ConfigurationType,
		AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesMap.LogLevel,
		AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__PropertiesMap.MetricsLevel,
	}
)

// AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration is a binding for AWS::KinesisAnalyticsV2::Application.MonitoringConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-monitoringconfiguration.html
type AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration struct {
	// ConfigurationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-monitoringconfiguration.html#cfn-kinesisanalyticsv2-application-monitoringconfiguration-configurationtype
	ConfigurationType cfz.Expression[string] `json:"ConfigurationType,omitempty"`

	// LogLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-monitoringconfiguration.html#cfn-kinesisanalyticsv2-application-monitoringconfiguration-loglevel
	LogLevel cfz.Expression[string] `json:"LogLevel,omitempty"`

	// MetricsLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-monitoringconfiguration.html#cfn-kinesisanalyticsv2-application-monitoringconfiguration-metricslevel
	MetricsLevel cfz.Expression[string] `json:"MetricsLevel,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration initializes a new AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration.
func New__AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration() AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	return AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) GetType() string {
	return AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration__Type
}

// Set__ConfigurationType updates property "ConfigurationType".
func (t AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) Set__ConfigurationType(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	t.ConfigurationType = v
	return t
}

// SetV__ConfigurationType updates property "ConfigurationType".
func (t AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) SetV__ConfigurationType(v string) AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	t.ConfigurationType = cfz.V(v)
	return t
}

// Set__LogLevel updates property "LogLevel".
func (t AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) Set__LogLevel(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	t.LogLevel = v
	return t
}

// SetV__LogLevel updates property "LogLevel".
func (t AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) SetV__LogLevel(v string) AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	t.LogLevel = cfz.V(v)
	return t
}

// Set__MetricsLevel updates property "MetricsLevel".
func (t AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) Set__MetricsLevel(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	t.MetricsLevel = v
	return t
}

// SetV__MetricsLevel updates property "MetricsLevel".
func (t AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration) SetV__MetricsLevel(v string) AWS_KinesisAnalyticsV2_Application_MonitoringConfiguration {
	t.MetricsLevel = cfz.V(v)
	return t
}
