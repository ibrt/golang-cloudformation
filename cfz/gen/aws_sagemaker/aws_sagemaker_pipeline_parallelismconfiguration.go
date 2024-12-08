// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Pipeline_ParallelismConfiguration__Type is the CloudFormation type for AWS::SageMaker::Pipeline.ParallelismConfiguration.
	AWS_SageMaker_Pipeline_ParallelismConfiguration__Type = "AWS::SageMaker::Pipeline.ParallelismConfiguration"
)

var (
	// AWS_SageMaker_Pipeline_ParallelismConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Pipeline.ParallelismConfiguration.
	AWS_SageMaker_Pipeline_ParallelismConfiguration__PropertiesMap = struct {
		MaxParallelExecutionSteps string
	}{
		MaxParallelExecutionSteps: "MaxParallelExecutionSteps",
	}

	// AWS_SageMaker_Pipeline_ParallelismConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Pipeline.ParallelismConfiguration.
	AWS_SageMaker_Pipeline_ParallelismConfiguration__PropertiesSlice = []string{
		AWS_SageMaker_Pipeline_ParallelismConfiguration__PropertiesMap.MaxParallelExecutionSteps,
	}
)

// AWS_SageMaker_Pipeline_ParallelismConfiguration is a binding for AWS::SageMaker::Pipeline.ParallelismConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-parallelismconfiguration.html
type AWS_SageMaker_Pipeline_ParallelismConfiguration struct {
	// MaxParallelExecutionSteps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-parallelismconfiguration.html#cfn-sagemaker-pipeline-parallelismconfiguration-maxparallelexecutionsteps
	MaxParallelExecutionSteps cfz.Expression[int32] `json:"MaxParallelExecutionSteps,omitempty"`
}

// New__AWS_SageMaker_Pipeline_ParallelismConfiguration initializes a new AWS_SageMaker_Pipeline_ParallelismConfiguration.
func New__AWS_SageMaker_Pipeline_ParallelismConfiguration() AWS_SageMaker_Pipeline_ParallelismConfiguration {
	return AWS_SageMaker_Pipeline_ParallelismConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Pipeline_ParallelismConfiguration) GetType() string {
	return AWS_SageMaker_Pipeline_ParallelismConfiguration__Type
}

// Set__MaxParallelExecutionSteps updates property "MaxParallelExecutionSteps".
func (t AWS_SageMaker_Pipeline_ParallelismConfiguration) Set__MaxParallelExecutionSteps(v cfz.Expression[int32]) AWS_SageMaker_Pipeline_ParallelismConfiguration {
	t.MaxParallelExecutionSteps = v
	return t
}

// SetV__MaxParallelExecutionSteps updates property "MaxParallelExecutionSteps".
func (t AWS_SageMaker_Pipeline_ParallelismConfiguration) SetV__MaxParallelExecutionSteps(v int32) AWS_SageMaker_Pipeline_ParallelismConfiguration {
	t.MaxParallelExecutionSteps = cfz.V(v)
	return t
}
