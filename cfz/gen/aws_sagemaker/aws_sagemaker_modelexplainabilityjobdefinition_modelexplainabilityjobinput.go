// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__Type is the CloudFormation type for AWS::SageMaker::ModelExplainabilityJobDefinition.ModelExplainabilityJobInput.
	AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__Type = "AWS::SageMaker::ModelExplainabilityJobDefinition.ModelExplainabilityJobInput"
)

var (
	// AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelExplainabilityJobDefinition.ModelExplainabilityJobInput.
	AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__PropertiesMap = struct {
		BatchTransformInput string
		EndpointInput       string
	}{
		BatchTransformInput: "BatchTransformInput",
		EndpointInput:       "EndpointInput",
	}

	// AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelExplainabilityJobDefinition.ModelExplainabilityJobInput.
	AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__PropertiesSlice = []string{
		AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__PropertiesMap.BatchTransformInput,
		AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__PropertiesMap.EndpointInput,
	}
)

// AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput is a binding for AWS::SageMaker::ModelExplainabilityJobDefinition.ModelExplainabilityJobInput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput.html
type AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput struct {
	// BatchTransformInput is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput-batchtransforminput
	BatchTransformInput cfz.Expression[AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput] `json:"BatchTransformInput,omitempty"`

	// EndpointInput is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput-endpointinput
	EndpointInput cfz.Expression[AWS_SageMaker_ModelExplainabilityJobDefinition_EndpointInput] `json:"EndpointInput,omitempty"`
}

// New__AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput initializes a new AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput.
func New__AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput() AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput {
	return AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput) GetType() string {
	return AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput__Type
}

// Set__BatchTransformInput updates property "BatchTransformInput".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput) Set__BatchTransformInput(v cfz.Expression[AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput]) AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput {
	t.BatchTransformInput = v
	return t
}

// SetV__BatchTransformInput updates property "BatchTransformInput".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput) SetV__BatchTransformInput(v AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput {
	t.BatchTransformInput = cfz.V(v)
	return t
}

// Set__EndpointInput updates property "EndpointInput".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput) Set__EndpointInput(v cfz.Expression[AWS_SageMaker_ModelExplainabilityJobDefinition_EndpointInput]) AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput {
	t.EndpointInput = v
	return t
}

// SetV__EndpointInput updates property "EndpointInput".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput) SetV__EndpointInput(v AWS_SageMaker_ModelExplainabilityJobDefinition_EndpointInput) AWS_SageMaker_ModelExplainabilityJobDefinition_ModelExplainabilityJobInput {
	t.EndpointInput = cfz.V(v)
	return t
}
