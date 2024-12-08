// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelCard_TrainingHyperParameter__Type is the CloudFormation type for AWS::SageMaker::ModelCard.TrainingHyperParameter.
	AWS_SageMaker_ModelCard_TrainingHyperParameter__Type = "AWS::SageMaker::ModelCard.TrainingHyperParameter"
)

var (
	// AWS_SageMaker_ModelCard_TrainingHyperParameter__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelCard.TrainingHyperParameter.
	AWS_SageMaker_ModelCard_TrainingHyperParameter__PropertiesMap = struct {
		Name  string
		Value string
	}{
		Name:  "Name",
		Value: "Value",
	}

	// AWS_SageMaker_ModelCard_TrainingHyperParameter__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelCard.TrainingHyperParameter.
	AWS_SageMaker_ModelCard_TrainingHyperParameter__PropertiesSlice = []string{
		AWS_SageMaker_ModelCard_TrainingHyperParameter__PropertiesMap.Name,
		AWS_SageMaker_ModelCard_TrainingHyperParameter__PropertiesMap.Value,
	}
)

// AWS_SageMaker_ModelCard_TrainingHyperParameter is a binding for AWS::SageMaker::ModelCard.TrainingHyperParameter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-traininghyperparameter.html
type AWS_SageMaker_ModelCard_TrainingHyperParameter struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-traininghyperparameter.html#cfn-sagemaker-modelcard-traininghyperparameter-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-traininghyperparameter.html#cfn-sagemaker-modelcard-traininghyperparameter-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_SageMaker_ModelCard_TrainingHyperParameter initializes a new AWS_SageMaker_ModelCard_TrainingHyperParameter.
func New__AWS_SageMaker_ModelCard_TrainingHyperParameter() AWS_SageMaker_ModelCard_TrainingHyperParameter {
	return AWS_SageMaker_ModelCard_TrainingHyperParameter{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelCard_TrainingHyperParameter) GetType() string {
	return AWS_SageMaker_ModelCard_TrainingHyperParameter__Type
}

// Set__Name updates property "Name".
func (t AWS_SageMaker_ModelCard_TrainingHyperParameter) Set__Name(v cfz.Expression[string]) AWS_SageMaker_ModelCard_TrainingHyperParameter {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_SageMaker_ModelCard_TrainingHyperParameter) SetV__Name(v string) AWS_SageMaker_ModelCard_TrainingHyperParameter {
	t.Name = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_SageMaker_ModelCard_TrainingHyperParameter) Set__Value(v cfz.Expression[string]) AWS_SageMaker_ModelCard_TrainingHyperParameter {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_SageMaker_ModelCard_TrainingHyperParameter) SetV__Value(v string) AWS_SageMaker_ModelCard_TrainingHyperParameter {
	t.Value = cfz.V(v)
	return t
}
