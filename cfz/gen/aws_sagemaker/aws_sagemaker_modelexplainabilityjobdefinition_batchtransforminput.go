// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__Type is the CloudFormation type for AWS::SageMaker::ModelExplainabilityJobDefinition.BatchTransformInput.
	AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__Type = "AWS::SageMaker::ModelExplainabilityJobDefinition.BatchTransformInput"
)

var (
	// AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelExplainabilityJobDefinition.BatchTransformInput.
	AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap = struct {
		DataCapturedDestinationS3Uri string
		DatasetFormat                string
		FeaturesAttribute            string
		InferenceAttribute           string
		LocalPath                    string
		ProbabilityAttribute         string
		S3DataDistributionType       string
		S3InputMode                  string
	}{
		DataCapturedDestinationS3Uri: "DataCapturedDestinationS3Uri",
		DatasetFormat:                "DatasetFormat",
		FeaturesAttribute:            "FeaturesAttribute",
		InferenceAttribute:           "InferenceAttribute",
		LocalPath:                    "LocalPath",
		ProbabilityAttribute:         "ProbabilityAttribute",
		S3DataDistributionType:       "S3DataDistributionType",
		S3InputMode:                  "S3InputMode",
	}

	// AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelExplainabilityJobDefinition.BatchTransformInput.
	AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesSlice = []string{
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.DataCapturedDestinationS3Uri,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.DatasetFormat,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.FeaturesAttribute,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.InferenceAttribute,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.LocalPath,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.ProbabilityAttribute,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.S3DataDistributionType,
		AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__PropertiesMap.S3InputMode,
	}
)

// AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput is a binding for AWS::SageMaker::ModelExplainabilityJobDefinition.BatchTransformInput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html
type AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput struct {
	// DataCapturedDestinationS3Uri is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-datacaptureddestinations3uri
	DataCapturedDestinationS3Uri cfz.Expression[string] `json:"DataCapturedDestinationS3Uri,omitempty"`

	// DatasetFormat is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-datasetformat
	DatasetFormat cfz.Expression[AWS_SageMaker_ModelExplainabilityJobDefinition_DatasetFormat] `json:"DatasetFormat,omitempty"`

	// FeaturesAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-featuresattribute
	FeaturesAttribute cfz.Expression[string] `json:"FeaturesAttribute,omitempty"`

	// InferenceAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-inferenceattribute
	InferenceAttribute cfz.Expression[string] `json:"InferenceAttribute,omitempty"`

	// LocalPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-localpath
	LocalPath cfz.Expression[string] `json:"LocalPath,omitempty"`

	// ProbabilityAttribute is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-probabilityattribute
	ProbabilityAttribute cfz.Expression[string] `json:"ProbabilityAttribute,omitempty"`

	// S3DataDistributionType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-s3datadistributiontype
	S3DataDistributionType cfz.Expression[string] `json:"S3DataDistributionType,omitempty"`

	// S3InputMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-batchtransforminput.html#cfn-sagemaker-modelexplainabilityjobdefinition-batchtransforminput-s3inputmode
	S3InputMode cfz.Expression[string] `json:"S3InputMode,omitempty"`
}

// New__AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput initializes a new AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput.
func New__AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput() AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	return AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) GetType() string {
	return AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput__Type
}

// Set__DataCapturedDestinationS3Uri updates property "DataCapturedDestinationS3Uri".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__DataCapturedDestinationS3Uri(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.DataCapturedDestinationS3Uri = v
	return t
}

// SetV__DataCapturedDestinationS3Uri updates property "DataCapturedDestinationS3Uri".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__DataCapturedDestinationS3Uri(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.DataCapturedDestinationS3Uri = cfz.V(v)
	return t
}

// Set__DatasetFormat updates property "DatasetFormat".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__DatasetFormat(v cfz.Expression[AWS_SageMaker_ModelExplainabilityJobDefinition_DatasetFormat]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.DatasetFormat = v
	return t
}

// SetV__DatasetFormat updates property "DatasetFormat".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__DatasetFormat(v AWS_SageMaker_ModelExplainabilityJobDefinition_DatasetFormat) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.DatasetFormat = cfz.V(v)
	return t
}

// Set__FeaturesAttribute updates property "FeaturesAttribute".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__FeaturesAttribute(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.FeaturesAttribute = v
	return t
}

// SetV__FeaturesAttribute updates property "FeaturesAttribute".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__FeaturesAttribute(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.FeaturesAttribute = cfz.V(v)
	return t
}

// Set__InferenceAttribute updates property "InferenceAttribute".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__InferenceAttribute(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.InferenceAttribute = v
	return t
}

// SetV__InferenceAttribute updates property "InferenceAttribute".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__InferenceAttribute(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.InferenceAttribute = cfz.V(v)
	return t
}

// Set__LocalPath updates property "LocalPath".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__LocalPath(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.LocalPath = v
	return t
}

// SetV__LocalPath updates property "LocalPath".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__LocalPath(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.LocalPath = cfz.V(v)
	return t
}

// Set__ProbabilityAttribute updates property "ProbabilityAttribute".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__ProbabilityAttribute(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.ProbabilityAttribute = v
	return t
}

// SetV__ProbabilityAttribute updates property "ProbabilityAttribute".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__ProbabilityAttribute(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.ProbabilityAttribute = cfz.V(v)
	return t
}

// Set__S3DataDistributionType updates property "S3DataDistributionType".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__S3DataDistributionType(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.S3DataDistributionType = v
	return t
}

// SetV__S3DataDistributionType updates property "S3DataDistributionType".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__S3DataDistributionType(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.S3DataDistributionType = cfz.V(v)
	return t
}

// Set__S3InputMode updates property "S3InputMode".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) Set__S3InputMode(v cfz.Expression[string]) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.S3InputMode = v
	return t
}

// SetV__S3InputMode updates property "S3InputMode".
func (t AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput) SetV__S3InputMode(v string) AWS_SageMaker_ModelExplainabilityJobDefinition_BatchTransformInput {
	t.S3InputMode = cfz.V(v)
	return t
}
