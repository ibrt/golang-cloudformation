// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__Type is the CloudFormation type for AWS::SageMaker::ModelQualityJobDefinition.MonitoringOutput.
	AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__Type = "AWS::SageMaker::ModelQualityJobDefinition.MonitoringOutput"
)

var (
	// AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelQualityJobDefinition.MonitoringOutput.
	AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__PropertiesMap = struct {
		S3Output string
	}{
		S3Output: "S3Output",
	}

	// AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelQualityJobDefinition.MonitoringOutput.
	AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__PropertiesSlice = []string{
		AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__PropertiesMap.S3Output,
	}
)

// AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput is a binding for AWS::SageMaker::ModelQualityJobDefinition.MonitoringOutput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-monitoringoutput.html
type AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput struct {
	// S3Output is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelqualityjobdefinition-monitoringoutput.html#cfn-sagemaker-modelqualityjobdefinition-monitoringoutput-s3output
	S3Output cfz.Expression[AWS_SageMaker_ModelQualityJobDefinition_S3Output] `json:"S3Output,omitempty"`
}

// New__AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput initializes a new AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput.
func New__AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput() AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput {
	return AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput) GetType() string {
	return AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput__Type
}

// Set__S3Output updates property "S3Output".
func (t AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput) Set__S3Output(v cfz.Expression[AWS_SageMaker_ModelQualityJobDefinition_S3Output]) AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput {
	t.S3Output = v
	return t
}

// SetV__S3Output updates property "S3Output".
func (t AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput) SetV__S3Output(v AWS_SageMaker_ModelQualityJobDefinition_S3Output) AWS_SageMaker_ModelQualityJobDefinition_MonitoringOutput {
	t.S3Output = cfz.V(v)
	return t
}
