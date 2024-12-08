// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__Type is the CloudFormation type for AWS::SageMaker::ModelBiasJobDefinition.EndpointInput.
	AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__Type = "AWS::SageMaker::ModelBiasJobDefinition.EndpointInput"
)

var (
	// AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelBiasJobDefinition.EndpointInput.
	AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap = struct {
		EndTimeOffset                 string
		EndpointName                  string
		FeaturesAttribute             string
		InferenceAttribute            string
		LocalPath                     string
		ProbabilityAttribute          string
		ProbabilityThresholdAttribute string
		S3DataDistributionType        string
		S3InputMode                   string
		StartTimeOffset               string
	}{
		EndTimeOffset:                 "EndTimeOffset",
		EndpointName:                  "EndpointName",
		FeaturesAttribute:             "FeaturesAttribute",
		InferenceAttribute:            "InferenceAttribute",
		LocalPath:                     "LocalPath",
		ProbabilityAttribute:          "ProbabilityAttribute",
		ProbabilityThresholdAttribute: "ProbabilityThresholdAttribute",
		S3DataDistributionType:        "S3DataDistributionType",
		S3InputMode:                   "S3InputMode",
		StartTimeOffset:               "StartTimeOffset",
	}

	// AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelBiasJobDefinition.EndpointInput.
	AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesSlice = []string{
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.EndTimeOffset,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.EndpointName,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.FeaturesAttribute,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.InferenceAttribute,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.LocalPath,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.ProbabilityAttribute,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.ProbabilityThresholdAttribute,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.S3DataDistributionType,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.S3InputMode,
		AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__PropertiesMap.StartTimeOffset,
	}
)

// AWS_SageMaker_ModelBiasJobDefinition_EndpointInput is a binding for AWS::SageMaker::ModelBiasJobDefinition.EndpointInput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html
type AWS_SageMaker_ModelBiasJobDefinition_EndpointInput struct {
	// EndTimeOffset is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-endtimeoffset
	EndTimeOffset cfz.Expression[string] `json:"EndTimeOffset,omitempty"`

	// EndpointName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-endpointname
	EndpointName cfz.Expression[string] `json:"EndpointName,omitempty"`

	// FeaturesAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-featuresattribute
	FeaturesAttribute cfz.Expression[string] `json:"FeaturesAttribute,omitempty"`

	// InferenceAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-inferenceattribute
	InferenceAttribute cfz.Expression[string] `json:"InferenceAttribute,omitempty"`

	// LocalPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-localpath
	LocalPath cfz.Expression[string] `json:"LocalPath,omitempty"`

	// ProbabilityAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-probabilityattribute
	ProbabilityAttribute cfz.Expression[string] `json:"ProbabilityAttribute,omitempty"`

	// ProbabilityThresholdAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-probabilitythresholdattribute
	ProbabilityThresholdAttribute cfz.Expression[float64] `json:"ProbabilityThresholdAttribute,omitempty"`

	// S3DataDistributionType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-s3datadistributiontype
	S3DataDistributionType cfz.Expression[string] `json:"S3DataDistributionType,omitempty"`

	// S3InputMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-s3inputmode
	S3InputMode cfz.Expression[string] `json:"S3InputMode,omitempty"`

	// StartTimeOffset is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-endpointinput.html#cfn-sagemaker-modelbiasjobdefinition-endpointinput-starttimeoffset
	StartTimeOffset cfz.Expression[string] `json:"StartTimeOffset,omitempty"`
}

// New__AWS_SageMaker_ModelBiasJobDefinition_EndpointInput initializes a new AWS_SageMaker_ModelBiasJobDefinition_EndpointInput.
func New__AWS_SageMaker_ModelBiasJobDefinition_EndpointInput() AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	return AWS_SageMaker_ModelBiasJobDefinition_EndpointInput{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) GetType() string {
	return AWS_SageMaker_ModelBiasJobDefinition_EndpointInput__Type
}

// Set__EndTimeOffset updates property "EndTimeOffset".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__EndTimeOffset(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.EndTimeOffset = v
	return t
}

// SetV__EndTimeOffset updates property "EndTimeOffset".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__EndTimeOffset(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.EndTimeOffset = cfz.V(v)
	return t
}

// Set__EndpointName updates property "EndpointName".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__EndpointName(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.EndpointName = v
	return t
}

// SetV__EndpointName updates property "EndpointName".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__EndpointName(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.EndpointName = cfz.V(v)
	return t
}

// Set__FeaturesAttribute updates property "FeaturesAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__FeaturesAttribute(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.FeaturesAttribute = v
	return t
}

// SetV__FeaturesAttribute updates property "FeaturesAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__FeaturesAttribute(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.FeaturesAttribute = cfz.V(v)
	return t
}

// Set__InferenceAttribute updates property "InferenceAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__InferenceAttribute(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.InferenceAttribute = v
	return t
}

// SetV__InferenceAttribute updates property "InferenceAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__InferenceAttribute(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.InferenceAttribute = cfz.V(v)
	return t
}

// Set__LocalPath updates property "LocalPath".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__LocalPath(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.LocalPath = v
	return t
}

// SetV__LocalPath updates property "LocalPath".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__LocalPath(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.LocalPath = cfz.V(v)
	return t
}

// Set__ProbabilityAttribute updates property "ProbabilityAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__ProbabilityAttribute(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.ProbabilityAttribute = v
	return t
}

// SetV__ProbabilityAttribute updates property "ProbabilityAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__ProbabilityAttribute(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.ProbabilityAttribute = cfz.V(v)
	return t
}

// Set__ProbabilityThresholdAttribute updates property "ProbabilityThresholdAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__ProbabilityThresholdAttribute(v cfz.Expression[float64]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.ProbabilityThresholdAttribute = v
	return t
}

// SetV__ProbabilityThresholdAttribute updates property "ProbabilityThresholdAttribute".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__ProbabilityThresholdAttribute(v float64) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.ProbabilityThresholdAttribute = cfz.V(v)
	return t
}

// Set__S3DataDistributionType updates property "S3DataDistributionType".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__S3DataDistributionType(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.S3DataDistributionType = v
	return t
}

// SetV__S3DataDistributionType updates property "S3DataDistributionType".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__S3DataDistributionType(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.S3DataDistributionType = cfz.V(v)
	return t
}

// Set__S3InputMode updates property "S3InputMode".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__S3InputMode(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.S3InputMode = v
	return t
}

// SetV__S3InputMode updates property "S3InputMode".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__S3InputMode(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.S3InputMode = cfz.V(v)
	return t
}

// Set__StartTimeOffset updates property "StartTimeOffset".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) Set__StartTimeOffset(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.StartTimeOffset = v
	return t
}

// SetV__StartTimeOffset updates property "StartTimeOffset".
func (t AWS_SageMaker_ModelBiasJobDefinition_EndpointInput) SetV__StartTimeOffset(v string) AWS_SageMaker_ModelBiasJobDefinition_EndpointInput {
	t.StartTimeOffset = cfz.V(v)
	return t
}
