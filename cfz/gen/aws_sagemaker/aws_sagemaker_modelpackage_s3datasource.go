// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelPackage_S3DataSource__Type is the CloudFormation type for AWS::SageMaker::ModelPackage.S3DataSource.
	AWS_SageMaker_ModelPackage_S3DataSource__Type = "AWS::SageMaker::ModelPackage.S3DataSource"
)

var (
	// AWS_SageMaker_ModelPackage_S3DataSource__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelPackage.S3DataSource.
	AWS_SageMaker_ModelPackage_S3DataSource__PropertiesMap = struct {
		S3DataType string
		S3Uri      string
	}{
		S3DataType: "S3DataType",
		S3Uri:      "S3Uri",
	}

	// AWS_SageMaker_ModelPackage_S3DataSource__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelPackage.S3DataSource.
	AWS_SageMaker_ModelPackage_S3DataSource__PropertiesSlice = []string{
		AWS_SageMaker_ModelPackage_S3DataSource__PropertiesMap.S3DataType,
		AWS_SageMaker_ModelPackage_S3DataSource__PropertiesMap.S3Uri,
	}
)

// AWS_SageMaker_ModelPackage_S3DataSource is a binding for AWS::SageMaker::ModelPackage.S3DataSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-s3datasource.html
type AWS_SageMaker_ModelPackage_S3DataSource struct {
	// S3DataType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-s3datasource.html#cfn-sagemaker-modelpackage-s3datasource-s3datatype
	S3DataType cfz.Expression[string] `json:"S3DataType,omitempty"`

	// S3Uri is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-s3datasource.html#cfn-sagemaker-modelpackage-s3datasource-s3uri
	S3Uri cfz.Expression[string] `json:"S3Uri,omitempty"`
}

// New__AWS_SageMaker_ModelPackage_S3DataSource initializes a new AWS_SageMaker_ModelPackage_S3DataSource.
func New__AWS_SageMaker_ModelPackage_S3DataSource() AWS_SageMaker_ModelPackage_S3DataSource {
	return AWS_SageMaker_ModelPackage_S3DataSource{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelPackage_S3DataSource) GetType() string {
	return AWS_SageMaker_ModelPackage_S3DataSource__Type
}

// Set__S3DataType updates property "S3DataType".
func (t AWS_SageMaker_ModelPackage_S3DataSource) Set__S3DataType(v cfz.Expression[string]) AWS_SageMaker_ModelPackage_S3DataSource {
	t.S3DataType = v
	return t
}

// SetV__S3DataType updates property "S3DataType".
func (t AWS_SageMaker_ModelPackage_S3DataSource) SetV__S3DataType(v string) AWS_SageMaker_ModelPackage_S3DataSource {
	t.S3DataType = cfz.V(v)
	return t
}

// Set__S3Uri updates property "S3Uri".
func (t AWS_SageMaker_ModelPackage_S3DataSource) Set__S3Uri(v cfz.Expression[string]) AWS_SageMaker_ModelPackage_S3DataSource {
	t.S3Uri = v
	return t
}

// SetV__S3Uri updates property "S3Uri".
func (t AWS_SageMaker_ModelPackage_S3DataSource) SetV__S3Uri(v string) AWS_SageMaker_ModelPackage_S3DataSource {
	t.S3Uri = cfz.V(v)
	return t
}
