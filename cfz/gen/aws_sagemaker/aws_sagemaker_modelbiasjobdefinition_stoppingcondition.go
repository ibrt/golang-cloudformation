// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__Type is the CloudFormation type for AWS::SageMaker::ModelBiasJobDefinition.StoppingCondition.
	AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__Type = "AWS::SageMaker::ModelBiasJobDefinition.StoppingCondition"
)

var (
	// AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelBiasJobDefinition.StoppingCondition.
	AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__PropertiesMap = struct {
		MaxRuntimeInSeconds string
	}{
		MaxRuntimeInSeconds: "MaxRuntimeInSeconds",
	}

	// AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelBiasJobDefinition.StoppingCondition.
	AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__PropertiesSlice = []string{
		AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__PropertiesMap.MaxRuntimeInSeconds,
	}
)

// AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition is a binding for AWS::SageMaker::ModelBiasJobDefinition.StoppingCondition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-stoppingcondition.html
type AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition struct {
	// MaxRuntimeInSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-stoppingcondition.html#cfn-sagemaker-modelbiasjobdefinition-stoppingcondition-maxruntimeinseconds
	MaxRuntimeInSeconds cfz.Expression[int32] `json:"MaxRuntimeInSeconds,omitempty"`
}

// New__AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition initializes a new AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition.
func New__AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition() AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition {
	return AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition) GetType() string {
	return AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition__Type
}

// Set__MaxRuntimeInSeconds updates property "MaxRuntimeInSeconds".
func (t AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition) Set__MaxRuntimeInSeconds(v cfz.Expression[int32]) AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition {
	t.MaxRuntimeInSeconds = v
	return t
}

// SetV__MaxRuntimeInSeconds updates property "MaxRuntimeInSeconds".
func (t AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition) SetV__MaxRuntimeInSeconds(v int32) AWS_SageMaker_ModelBiasJobDefinition_StoppingCondition {
	t.MaxRuntimeInSeconds = cfz.V(v)
	return t
}
