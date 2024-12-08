// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::Application.ApplicationConfiguration.
	AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__Type = "AWS::KinesisAnalyticsV2::Application.ApplicationConfiguration"
)

var (
	// AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.ApplicationConfiguration.
	AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap = struct {
		ApplicationCodeConfiguration           string
		ApplicationSnapshotConfiguration       string
		ApplicationSystemRollbackConfiguration string
		EnvironmentProperties                  string
		FlinkApplicationConfiguration          string
		SqlApplicationConfiguration            string
		VpcConfigurations                      string
		ZeppelinApplicationConfiguration       string
	}{
		ApplicationCodeConfiguration:           "ApplicationCodeConfiguration",
		ApplicationSnapshotConfiguration:       "ApplicationSnapshotConfiguration",
		ApplicationSystemRollbackConfiguration: "ApplicationSystemRollbackConfiguration",
		EnvironmentProperties:                  "EnvironmentProperties",
		FlinkApplicationConfiguration:          "FlinkApplicationConfiguration",
		SqlApplicationConfiguration:            "SqlApplicationConfiguration",
		VpcConfigurations:                      "VpcConfigurations",
		ZeppelinApplicationConfiguration:       "ZeppelinApplicationConfiguration",
	}

	// AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.ApplicationConfiguration.
	AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.ApplicationCodeConfiguration,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.ApplicationSnapshotConfiguration,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.ApplicationSystemRollbackConfiguration,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.EnvironmentProperties,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.FlinkApplicationConfiguration,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.SqlApplicationConfiguration,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.VpcConfigurations,
		AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__PropertiesMap.ZeppelinApplicationConfiguration,
	}
)

// AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration is a binding for AWS::KinesisAnalyticsV2::Application.ApplicationConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html
type AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration struct {
	// ApplicationCodeConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationcodeconfiguration
	ApplicationCodeConfiguration cfz.Expression[AWS_KinesisAnalyticsV2_Application_ApplicationCodeConfiguration] `json:"ApplicationCodeConfiguration,omitempty"`

	// ApplicationSnapshotConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationsnapshotconfiguration
	ApplicationSnapshotConfiguration cfz.Expression[AWS_KinesisAnalyticsV2_Application_ApplicationSnapshotConfiguration] `json:"ApplicationSnapshotConfiguration,omitempty"`

	// ApplicationSystemRollbackConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-applicationsystemrollbackconfiguration
	ApplicationSystemRollbackConfiguration cfz.Expression[AWS_KinesisAnalyticsV2_Application_ApplicationSystemRollbackConfiguration] `json:"ApplicationSystemRollbackConfiguration,omitempty"`

	// EnvironmentProperties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-environmentproperties
	EnvironmentProperties cfz.Expression[AWS_KinesisAnalyticsV2_Application_EnvironmentProperties] `json:"EnvironmentProperties,omitempty"`

	// FlinkApplicationConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-flinkapplicationconfiguration
	FlinkApplicationConfiguration cfz.Expression[AWS_KinesisAnalyticsV2_Application_FlinkApplicationConfiguration] `json:"FlinkApplicationConfiguration,omitempty"`

	// SqlApplicationConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-sqlapplicationconfiguration
	SqlApplicationConfiguration cfz.Expression[AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration] `json:"SqlApplicationConfiguration,omitempty"`

	// VpcConfigurations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-vpcconfigurations
	VpcConfigurations cfz.ExpressionSlice[AWS_KinesisAnalyticsV2_Application_VpcConfiguration] `json:"VpcConfigurations,omitempty"`

	// ZeppelinApplicationConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationconfiguration.html#cfn-kinesisanalyticsv2-application-applicationconfiguration-zeppelinapplicationconfiguration
	ZeppelinApplicationConfiguration cfz.Expression[AWS_KinesisAnalyticsV2_Application_ZeppelinApplicationConfiguration] `json:"ZeppelinApplicationConfiguration,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration initializes a new AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration.
func New__AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration() AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	return AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) GetType() string {
	return AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration__Type
}

// Set__ApplicationCodeConfiguration updates property "ApplicationCodeConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__ApplicationCodeConfiguration(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_ApplicationCodeConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ApplicationCodeConfiguration = v
	return t
}

// SetV__ApplicationCodeConfiguration updates property "ApplicationCodeConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__ApplicationCodeConfiguration(v AWS_KinesisAnalyticsV2_Application_ApplicationCodeConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ApplicationCodeConfiguration = cfz.V(v)
	return t
}

// Set__ApplicationSnapshotConfiguration updates property "ApplicationSnapshotConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__ApplicationSnapshotConfiguration(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_ApplicationSnapshotConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ApplicationSnapshotConfiguration = v
	return t
}

// SetV__ApplicationSnapshotConfiguration updates property "ApplicationSnapshotConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__ApplicationSnapshotConfiguration(v AWS_KinesisAnalyticsV2_Application_ApplicationSnapshotConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ApplicationSnapshotConfiguration = cfz.V(v)
	return t
}

// Set__ApplicationSystemRollbackConfiguration updates property "ApplicationSystemRollbackConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__ApplicationSystemRollbackConfiguration(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_ApplicationSystemRollbackConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ApplicationSystemRollbackConfiguration = v
	return t
}

// SetV__ApplicationSystemRollbackConfiguration updates property "ApplicationSystemRollbackConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__ApplicationSystemRollbackConfiguration(v AWS_KinesisAnalyticsV2_Application_ApplicationSystemRollbackConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ApplicationSystemRollbackConfiguration = cfz.V(v)
	return t
}

// Set__EnvironmentProperties updates property "EnvironmentProperties".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__EnvironmentProperties(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_EnvironmentProperties]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.EnvironmentProperties = v
	return t
}

// SetV__EnvironmentProperties updates property "EnvironmentProperties".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__EnvironmentProperties(v AWS_KinesisAnalyticsV2_Application_EnvironmentProperties) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.EnvironmentProperties = cfz.V(v)
	return t
}

// Set__FlinkApplicationConfiguration updates property "FlinkApplicationConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__FlinkApplicationConfiguration(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_FlinkApplicationConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.FlinkApplicationConfiguration = v
	return t
}

// SetV__FlinkApplicationConfiguration updates property "FlinkApplicationConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__FlinkApplicationConfiguration(v AWS_KinesisAnalyticsV2_Application_FlinkApplicationConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.FlinkApplicationConfiguration = cfz.V(v)
	return t
}

// Set__SqlApplicationConfiguration updates property "SqlApplicationConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__SqlApplicationConfiguration(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.SqlApplicationConfiguration = v
	return t
}

// SetV__SqlApplicationConfiguration updates property "SqlApplicationConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__SqlApplicationConfiguration(v AWS_KinesisAnalyticsV2_Application_SqlApplicationConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.SqlApplicationConfiguration = cfz.V(v)
	return t
}

// Set__VpcConfigurations updates property "VpcConfigurations".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__VpcConfigurations(v cfz.ExpressionSlice[AWS_KinesisAnalyticsV2_Application_VpcConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.VpcConfigurations = v
	return t
}

// SetS__VpcConfigurations updates property "VpcConfigurations".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetS__VpcConfigurations(v ...cfz.Expression[AWS_KinesisAnalyticsV2_Application_VpcConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.VpcConfigurations = cfz.S(v...)
	return t
}

// SetSV__VpcConfigurations updates property "VpcConfigurations".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetSV__VpcConfigurations(v ...AWS_KinesisAnalyticsV2_Application_VpcConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.VpcConfigurations = cfz.SV(v...)
	return t
}

// Set__ZeppelinApplicationConfiguration updates property "ZeppelinApplicationConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) Set__ZeppelinApplicationConfiguration(v cfz.Expression[AWS_KinesisAnalyticsV2_Application_ZeppelinApplicationConfiguration]) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ZeppelinApplicationConfiguration = v
	return t
}

// SetV__ZeppelinApplicationConfiguration updates property "ZeppelinApplicationConfiguration".
func (t AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration) SetV__ZeppelinApplicationConfiguration(v AWS_KinesisAnalyticsV2_Application_ZeppelinApplicationConfiguration) AWS_KinesisAnalyticsV2_Application_ApplicationConfiguration {
	t.ZeppelinApplicationConfiguration = cfz.V(v)
	return t
}
