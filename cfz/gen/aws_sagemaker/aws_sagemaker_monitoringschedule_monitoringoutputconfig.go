// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__Type is the CloudFormation type for AWS::SageMaker::MonitoringSchedule.MonitoringOutputConfig.
	AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__Type = "AWS::SageMaker::MonitoringSchedule.MonitoringOutputConfig"
)

var (
	// AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.MonitoringOutputConfig.
	AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__PropertiesMap = struct {
		KmsKeyId          string
		MonitoringOutputs string
	}{
		KmsKeyId:          "KmsKeyId",
		MonitoringOutputs: "MonitoringOutputs",
	}

	// AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::MonitoringSchedule.MonitoringOutputConfig.
	AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__PropertiesSlice = []string{
		AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__PropertiesMap.KmsKeyId,
		AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__PropertiesMap.MonitoringOutputs,
	}
)

// AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig is a binding for AWS::SageMaker::MonitoringSchedule.MonitoringOutputConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.html
type AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig struct {
	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.html#cfn-sagemaker-monitoringschedule-monitoringoutputconfig-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`

	// MonitoringOutputs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedule-monitoringoutputconfig.html#cfn-sagemaker-monitoringschedule-monitoringoutputconfig-monitoringoutputs
	MonitoringOutputs cfz.ExpressionSlice[AWS_SageMaker_MonitoringSchedule_MonitoringOutput] `json:"MonitoringOutputs,omitempty"`
}

// New__AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig initializes a new AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig.
func New__AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig() AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig {
	return AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig) GetType() string {
	return AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig__Type
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig) Set__KmsKeyId(v cfz.Expression[string]) AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig) SetV__KmsKeyId(v string) AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig {
	t.KmsKeyId = cfz.V(v)
	return t
}

// Set__MonitoringOutputs updates property "MonitoringOutputs".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig) Set__MonitoringOutputs(v cfz.ExpressionSlice[AWS_SageMaker_MonitoringSchedule_MonitoringOutput]) AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig {
	t.MonitoringOutputs = v
	return t
}

// SetS__MonitoringOutputs updates property "MonitoringOutputs".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig) SetS__MonitoringOutputs(v ...cfz.Expression[AWS_SageMaker_MonitoringSchedule_MonitoringOutput]) AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig {
	t.MonitoringOutputs = cfz.S(v...)
	return t
}

// SetSV__MonitoringOutputs updates property "MonitoringOutputs".
func (t AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig) SetSV__MonitoringOutputs(v ...AWS_SageMaker_MonitoringSchedule_MonitoringOutput) AWS_SageMaker_MonitoringSchedule_MonitoringOutputConfig {
	t.MonitoringOutputs = cfz.SV(v...)
	return t
}
