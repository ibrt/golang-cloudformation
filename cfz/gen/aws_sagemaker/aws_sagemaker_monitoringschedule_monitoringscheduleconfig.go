// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__Type is the CloudFormation type for AWS::SageMaker::MonitoringSchedule.MonitoringScheduleConfig.
	AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__Type = "AWS::SageMaker::MonitoringSchedule.MonitoringScheduleConfig"
)

var (
	// AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.MonitoringScheduleConfig.
	AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesMap = struct {
		MonitoringJobDefinition     string
		MonitoringJobDefinitionName string
		MonitoringType              string
		ScheduleConfig              string
	}{
		MonitoringJobDefinition:     "MonitoringJobDefinition",
		MonitoringJobDefinitionName: "MonitoringJobDefinitionName",
		MonitoringType:              "MonitoringType",
		ScheduleConfig:              "ScheduleConfig",
	}

	// AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.MonitoringScheduleConfig.
	AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesSlice = []string{
		AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesMap.MonitoringJobDefinition,
		AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesMap.MonitoringJobDefinitionName,
		AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesMap.MonitoringType,
		AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__PropertiesMap.ScheduleConfig,
	}
)

// AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig is a binding for AWS::SageMaker::MonitoringSchedule.MonitoringScheduleConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html
type AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig struct {
	// MonitoringJobDefinition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-monitoringjobdefinition
	MonitoringJobDefinition cfz.Expression[AWS_SageMaker_MonitoringSchedule_MonitoringJobDefinition] `json:"MonitoringJobDefinition,omitempty"`

	// MonitoringJobDefinitionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-monitoringjobdefinitionname
	MonitoringJobDefinitionName cfz.Expression[string] `json:"MonitoringJobDefinitionName,omitempty"`

	// MonitoringType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-monitoringtype
	MonitoringType cfz.Expression[string] `json:"MonitoringType,omitempty"`

	// ScheduleConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringscheduleconfig.html#cfn-sagemaker-monitoringschedule-monitoringscheduleconfig-scheduleconfig
	ScheduleConfig cfz.Expression[AWS_SageMaker_MonitoringSchedule_ScheduleConfig] `json:"ScheduleConfig,omitempty"`
}

// New__AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig initializes a new AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig.
func New__AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig() AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	return AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) GetType() string {
	return AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig__Type
}

// Set__MonitoringJobDefinition updates property "MonitoringJobDefinition".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) Set__MonitoringJobDefinition(v cfz.Expression[AWS_SageMaker_MonitoringSchedule_MonitoringJobDefinition]) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.MonitoringJobDefinition = v
	return t
}

// SetV__MonitoringJobDefinition updates property "MonitoringJobDefinition".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) SetV__MonitoringJobDefinition(v AWS_SageMaker_MonitoringSchedule_MonitoringJobDefinition) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.MonitoringJobDefinition = cfz.V(v)
	return t
}

// Set__MonitoringJobDefinitionName updates property "MonitoringJobDefinitionName".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) Set__MonitoringJobDefinitionName(v cfz.Expression[string]) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.MonitoringJobDefinitionName = v
	return t
}

// SetV__MonitoringJobDefinitionName updates property "MonitoringJobDefinitionName".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) SetV__MonitoringJobDefinitionName(v string) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.MonitoringJobDefinitionName = cfz.V(v)
	return t
}

// Set__MonitoringType updates property "MonitoringType".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) Set__MonitoringType(v cfz.Expression[string]) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.MonitoringType = v
	return t
}

// SetV__MonitoringType updates property "MonitoringType".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) SetV__MonitoringType(v string) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.MonitoringType = cfz.V(v)
	return t
}

// Set__ScheduleConfig updates property "ScheduleConfig".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) Set__ScheduleConfig(v cfz.Expression[AWS_SageMaker_MonitoringSchedule_ScheduleConfig]) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.ScheduleConfig = v
	return t
}

// SetV__ScheduleConfig updates property "ScheduleConfig".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig) SetV__ScheduleConfig(v AWS_SageMaker_MonitoringSchedule_ScheduleConfig) AWS_SageMaker_MonitoringSchedule_MonitoringScheduleConfig {
	t.ScheduleConfig = cfz.V(v)
	return t
}
