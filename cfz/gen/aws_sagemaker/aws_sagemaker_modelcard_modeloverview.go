// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelCard_ModelOverview__Type is the CloudFormation type for AWS::SageMaker::ModelCard.ModelOverview.
	AWS_SageMaker_ModelCard_ModelOverview__Type = "AWS::SageMaker::ModelCard.ModelOverview"
)

var (
	// AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelCard.ModelOverview.
	AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap = struct {
		AlgorithmType        string
		InferenceEnvironment string
		ModelArtifact        string
		ModelCreator         string
		ModelDescription     string
		ModelId              string
		ModelName            string
		ModelOwner           string
		ModelVersion         string
		ProblemType          string
	}{
		AlgorithmType:        "AlgorithmType",
		InferenceEnvironment: "InferenceEnvironment",
		ModelArtifact:        "ModelArtifact",
		ModelCreator:         "ModelCreator",
		ModelDescription:     "ModelDescription",
		ModelId:              "ModelId",
		ModelName:            "ModelName",
		ModelOwner:           "ModelOwner",
		ModelVersion:         "ModelVersion",
		ProblemType:          "ProblemType",
	}

	// AWS_SageMaker_ModelCard_ModelOverview__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelCard.ModelOverview.
	AWS_SageMaker_ModelCard_ModelOverview__PropertiesSlice = []string{
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.AlgorithmType,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.InferenceEnvironment,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelArtifact,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelCreator,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelDescription,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelId,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelName,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelOwner,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ModelVersion,
		AWS_SageMaker_ModelCard_ModelOverview__PropertiesMap.ProblemType,
	}
)

// AWS_SageMaker_ModelCard_ModelOverview is a binding for AWS::SageMaker::ModelCard.ModelOverview.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html
type AWS_SageMaker_ModelCard_ModelOverview struct {
	// AlgorithmType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-algorithmtype
	AlgorithmType cfz.Expression[string] `json:"AlgorithmType,omitempty"`

	// InferenceEnvironment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-inferenceenvironment
	InferenceEnvironment cfz.Expression[AWS_SageMaker_ModelCard_InferenceEnvironment] `json:"InferenceEnvironment,omitempty"`

	// ModelArtifact is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modelartifact
	ModelArtifact cfz.ExpressionSlice[string] `json:"ModelArtifact,omitempty"`

	// ModelCreator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modelcreator
	ModelCreator cfz.Expression[string] `json:"ModelCreator,omitempty"`

	// ModelDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modeldescription
	ModelDescription cfz.Expression[string] `json:"ModelDescription,omitempty"`

	// ModelId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modelid
	ModelId cfz.Expression[string] `json:"ModelId,omitempty"`

	// ModelName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modelname
	ModelName cfz.Expression[string] `json:"ModelName,omitempty"`

	// ModelOwner is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modelowner
	ModelOwner cfz.Expression[string] `json:"ModelOwner,omitempty"`

	// ModelVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-modelversion
	ModelVersion cfz.Expression[float64] `json:"ModelVersion,omitempty"`

	// ProblemType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-modeloverview.html#cfn-sagemaker-modelcard-modeloverview-problemtype
	ProblemType cfz.Expression[string] `json:"ProblemType,omitempty"`
}

// New__AWS_SageMaker_ModelCard_ModelOverview initializes a new AWS_SageMaker_ModelCard_ModelOverview.
func New__AWS_SageMaker_ModelCard_ModelOverview() AWS_SageMaker_ModelCard_ModelOverview {
	return AWS_SageMaker_ModelCard_ModelOverview{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelCard_ModelOverview) GetType() string {
	return AWS_SageMaker_ModelCard_ModelOverview__Type
}

// Set__AlgorithmType updates property "AlgorithmType".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__AlgorithmType(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.AlgorithmType = v
	return t
}

// SetV__AlgorithmType updates property "AlgorithmType".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__AlgorithmType(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.AlgorithmType = cfz.V(v)
	return t
}

// Set__InferenceEnvironment updates property "InferenceEnvironment".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__InferenceEnvironment(v cfz.Expression[AWS_SageMaker_ModelCard_InferenceEnvironment]) AWS_SageMaker_ModelCard_ModelOverview {
	t.InferenceEnvironment = v
	return t
}

// SetV__InferenceEnvironment updates property "InferenceEnvironment".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__InferenceEnvironment(v AWS_SageMaker_ModelCard_InferenceEnvironment) AWS_SageMaker_ModelCard_ModelOverview {
	t.InferenceEnvironment = cfz.V(v)
	return t
}

// Set__ModelArtifact updates property "ModelArtifact".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelArtifact(v cfz.ExpressionSlice[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelArtifact = v
	return t
}

// SetS__ModelArtifact updates property "ModelArtifact".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetS__ModelArtifact(v ...cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelArtifact = cfz.S(v...)
	return t
}

// SetSV__ModelArtifact updates property "ModelArtifact".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetSV__ModelArtifact(v ...string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelArtifact = cfz.SV(v...)
	return t
}

// Set__ModelCreator updates property "ModelCreator".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelCreator(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelCreator = v
	return t
}

// SetV__ModelCreator updates property "ModelCreator".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ModelCreator(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelCreator = cfz.V(v)
	return t
}

// Set__ModelDescription updates property "ModelDescription".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelDescription(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelDescription = v
	return t
}

// SetV__ModelDescription updates property "ModelDescription".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ModelDescription(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelDescription = cfz.V(v)
	return t
}

// Set__ModelId updates property "ModelId".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelId(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelId = v
	return t
}

// SetV__ModelId updates property "ModelId".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ModelId(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelId = cfz.V(v)
	return t
}

// Set__ModelName updates property "ModelName".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelName(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelName = v
	return t
}

// SetV__ModelName updates property "ModelName".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ModelName(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelName = cfz.V(v)
	return t
}

// Set__ModelOwner updates property "ModelOwner".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelOwner(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelOwner = v
	return t
}

// SetV__ModelOwner updates property "ModelOwner".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ModelOwner(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelOwner = cfz.V(v)
	return t
}

// Set__ModelVersion updates property "ModelVersion".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ModelVersion(v cfz.Expression[float64]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelVersion = v
	return t
}

// SetV__ModelVersion updates property "ModelVersion".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ModelVersion(v float64) AWS_SageMaker_ModelCard_ModelOverview {
	t.ModelVersion = cfz.V(v)
	return t
}

// Set__ProblemType updates property "ProblemType".
func (t AWS_SageMaker_ModelCard_ModelOverview) Set__ProblemType(v cfz.Expression[string]) AWS_SageMaker_ModelCard_ModelOverview {
	t.ProblemType = v
	return t
}

// SetV__ProblemType updates property "ProblemType".
func (t AWS_SageMaker_ModelCard_ModelOverview) SetV__ProblemType(v string) AWS_SageMaker_ModelCard_ModelOverview {
	t.ProblemType = cfz.V(v)
	return t
}
