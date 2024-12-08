// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelCard_InferenceEnvironment__Type is the CloudFormation type for AWS::SageMaker::ModelCard.InferenceEnvironment.
	AWS_SageMaker_ModelCard_InferenceEnvironment__Type = "AWS::SageMaker::ModelCard.InferenceEnvironment"
)

var (
	// AWS_SageMaker_ModelCard_InferenceEnvironment__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelCard.InferenceEnvironment.
	AWS_SageMaker_ModelCard_InferenceEnvironment__PropertiesMap = struct {
		ContainerImage string
	}{
		ContainerImage: "ContainerImage",
	}

	// AWS_SageMaker_ModelCard_InferenceEnvironment__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelCard.InferenceEnvironment.
	AWS_SageMaker_ModelCard_InferenceEnvironment__PropertiesSlice = []string{
		AWS_SageMaker_ModelCard_InferenceEnvironment__PropertiesMap.ContainerImage,
	}
)

// AWS_SageMaker_ModelCard_InferenceEnvironment is a binding for AWS::SageMaker::ModelCard.InferenceEnvironment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-inferenceenvironment.html
type AWS_SageMaker_ModelCard_InferenceEnvironment struct {
	// ContainerImage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-inferenceenvironment.html#cfn-sagemaker-modelcard-inferenceenvironment-containerimage
	ContainerImage cfz.ExpressionSlice[string] `json:"ContainerImage,omitempty"`
}

// New__AWS_SageMaker_ModelCard_InferenceEnvironment initializes a new AWS_SageMaker_ModelCard_InferenceEnvironment.
func New__AWS_SageMaker_ModelCard_InferenceEnvironment() AWS_SageMaker_ModelCard_InferenceEnvironment {
	return AWS_SageMaker_ModelCard_InferenceEnvironment{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelCard_InferenceEnvironment) GetType() string {
	return AWS_SageMaker_ModelCard_InferenceEnvironment__Type
}

// Set__ContainerImage updates property "ContainerImage".
func (t AWS_SageMaker_ModelCard_InferenceEnvironment) Set__ContainerImage(v cfz.ExpressionSlice[string]) AWS_SageMaker_ModelCard_InferenceEnvironment {
	t.ContainerImage = v
	return t
}

// SetS__ContainerImage updates property "ContainerImage".
func (t AWS_SageMaker_ModelCard_InferenceEnvironment) SetS__ContainerImage(v ...cfz.Expression[string]) AWS_SageMaker_ModelCard_InferenceEnvironment {
	t.ContainerImage = cfz.S(v...)
	return t
}

// SetSV__ContainerImage updates property "ContainerImage".
func (t AWS_SageMaker_ModelCard_InferenceEnvironment) SetSV__ContainerImage(v ...string) AWS_SageMaker_ModelCard_InferenceEnvironment {
	t.ContainerImage = cfz.SV(v...)
	return t
}
