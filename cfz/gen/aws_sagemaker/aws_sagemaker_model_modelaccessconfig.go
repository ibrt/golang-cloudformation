// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_Model_ModelAccessConfig__Type is the CloudFormation type for AWS::SageMaker::Model.ModelAccessConfig.
	AWS_SageMaker_Model_ModelAccessConfig__Type = "AWS::SageMaker::Model.ModelAccessConfig"
)

var (
	// AWS_SageMaker_Model_ModelAccessConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Model.ModelAccessConfig.
	AWS_SageMaker_Model_ModelAccessConfig__PropertiesMap = struct {
		AcceptEula string
	}{
		AcceptEula: "AcceptEula",
	}

	// AWS_SageMaker_Model_ModelAccessConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Model.ModelAccessConfig.
	AWS_SageMaker_Model_ModelAccessConfig__PropertiesSlice = []string{
		AWS_SageMaker_Model_ModelAccessConfig__PropertiesMap.AcceptEula,
	}
)

// AWS_SageMaker_Model_ModelAccessConfig is a binding for AWS::SageMaker::Model.ModelAccessConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-s3datasource-modelaccessconfig.html
type AWS_SageMaker_Model_ModelAccessConfig struct {
	// AcceptEula is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-s3datasource-modelaccessconfig.html#cfn-sagemaker-model-s3datasource-modelaccessconfig-accepteula
	AcceptEula cfz.Expression[bool] `json:"AcceptEula,omitempty"`
}

// New__AWS_SageMaker_Model_ModelAccessConfig initializes a new AWS_SageMaker_Model_ModelAccessConfig.
func New__AWS_SageMaker_Model_ModelAccessConfig() AWS_SageMaker_Model_ModelAccessConfig {
	return AWS_SageMaker_Model_ModelAccessConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_Model_ModelAccessConfig) GetType() string {
	return AWS_SageMaker_Model_ModelAccessConfig__Type
}

// Set__AcceptEula updates property "AcceptEula".
func (t AWS_SageMaker_Model_ModelAccessConfig) Set__AcceptEula(v cfz.Expression[bool]) AWS_SageMaker_Model_ModelAccessConfig {
	t.AcceptEula = v
	return t
}

// SetV__AcceptEula updates property "AcceptEula".
func (t AWS_SageMaker_Model_ModelAccessConfig) SetV__AcceptEula(v bool) AWS_SageMaker_Model_ModelAccessConfig {
	t.AcceptEula = cfz.V(v)
	return t
}
