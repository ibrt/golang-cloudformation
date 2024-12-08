// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Model_ModelDataSource__Type is the CloudFormation type for AWS::SageMaker::Model.ModelDataSource.
	AWS_SageMaker_Model_ModelDataSource__Type = "AWS::SageMaker::Model.ModelDataSource"
)

var (
	// AWS_SageMaker_Model_ModelDataSource__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Model.ModelDataSource.
	AWS_SageMaker_Model_ModelDataSource__PropertiesMap = struct {
		S3DataSource string
	}{
		S3DataSource: "S3DataSource",
	}

	// AWS_SageMaker_Model_ModelDataSource__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Model.ModelDataSource.
	AWS_SageMaker_Model_ModelDataSource__PropertiesSlice = []string{
		AWS_SageMaker_Model_ModelDataSource__PropertiesMap.S3DataSource,
	}
)

// AWS_SageMaker_Model_ModelDataSource is a binding for AWS::SageMaker::Model.ModelDataSource.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition-modeldatasource.html
type AWS_SageMaker_Model_ModelDataSource struct {
	// S3DataSource is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition-modeldatasource.html#cfn-sagemaker-model-containerdefinition-modeldatasource-s3datasource
	S3DataSource cfz.Expression[AWS_SageMaker_Model_S3DataSource] `json:"S3DataSource,omitempty"`
}

// New__AWS_SageMaker_Model_ModelDataSource initializes a new AWS_SageMaker_Model_ModelDataSource.
func New__AWS_SageMaker_Model_ModelDataSource() AWS_SageMaker_Model_ModelDataSource {
	return AWS_SageMaker_Model_ModelDataSource{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Model_ModelDataSource) GetType() string {
	return AWS_SageMaker_Model_ModelDataSource__Type
}

// Set__S3DataSource updates property "S3DataSource".
func (t AWS_SageMaker_Model_ModelDataSource) Set__S3DataSource(v cfz.Expression[AWS_SageMaker_Model_S3DataSource]) AWS_SageMaker_Model_ModelDataSource {
	t.S3DataSource = v
	return t
}

// SetV__S3DataSource updates property "S3DataSource".
func (t AWS_SageMaker_Model_ModelDataSource) SetV__S3DataSource(v AWS_SageMaker_Model_S3DataSource) AWS_SageMaker_Model_ModelDataSource {
	t.S3DataSource = cfz.V(v)
	return t
}
