// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__Type is the CloudFormation type for AWS::SageMaker::InferenceExperiment.ModelInfrastructureConfig.
	AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__Type = "AWS::SageMaker::InferenceExperiment.ModelInfrastructureConfig"
)

var (
	// AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::InferenceExperiment.ModelInfrastructureConfig.
	AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__PropertiesMap = struct {
		InfrastructureType      string
		RealTimeInferenceConfig string
	}{
		InfrastructureType:      "InfrastructureType",
		RealTimeInferenceConfig: "RealTimeInferenceConfig",
	}

	// AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::InferenceExperiment.ModelInfrastructureConfig.
	AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__PropertiesSlice = []string{
		AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__PropertiesMap.InfrastructureType,
		AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__PropertiesMap.RealTimeInferenceConfig,
	}
)

// AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig is a binding for AWS::SageMaker::InferenceExperiment.ModelInfrastructureConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html
type AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig struct {
	// InfrastructureType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html#cfn-sagemaker-inferenceexperiment-modelinfrastructureconfig-infrastructuretype
	InfrastructureType cfz.Expression[string] `json:"InfrastructureType,omitempty"`

	// RealTimeInferenceConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-modelinfrastructureconfig.html#cfn-sagemaker-inferenceexperiment-modelinfrastructureconfig-realtimeinferenceconfig
	RealTimeInferenceConfig cfz.Expression[AWS_SageMaker_InferenceExperiment_RealTimeInferenceConfig] `json:"RealTimeInferenceConfig,omitempty"`
}

// New__AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig initializes a new AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig.
func New__AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig() AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig {
	return AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig) GetType() string {
	return AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig__Type
}

// Set__InfrastructureType updates property "InfrastructureType".
func (t AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig) Set__InfrastructureType(v cfz.Expression[string]) AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig {
	t.InfrastructureType = v
	return t
}

// SetV__InfrastructureType updates property "InfrastructureType".
func (t AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig) SetV__InfrastructureType(v string) AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig {
	t.InfrastructureType = cfz.V(v)
	return t
}

// Set__RealTimeInferenceConfig updates property "RealTimeInferenceConfig".
func (t AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig) Set__RealTimeInferenceConfig(v cfz.Expression[AWS_SageMaker_InferenceExperiment_RealTimeInferenceConfig]) AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig {
	t.RealTimeInferenceConfig = v
	return t
}

// SetV__RealTimeInferenceConfig updates property "RealTimeInferenceConfig".
func (t AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig) SetV__RealTimeInferenceConfig(v AWS_SageMaker_InferenceExperiment_RealTimeInferenceConfig) AWS_SageMaker_InferenceExperiment_ModelInfrastructureConfig {
	t.RealTimeInferenceConfig = cfz.V(v)
	return t
}
