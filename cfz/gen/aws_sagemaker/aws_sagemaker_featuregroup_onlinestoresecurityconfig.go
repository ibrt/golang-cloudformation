// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__Type is the CloudFormation type for AWS::SageMaker::FeatureGroup.OnlineStoreSecurityConfig.
	AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__Type = "AWS::SageMaker::FeatureGroup.OnlineStoreSecurityConfig"
)

var (
	// AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::FeatureGroup.OnlineStoreSecurityConfig.
	AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__PropertiesMap = struct {
		KmsKeyId string
	}{
		KmsKeyId: "KmsKeyId",
	}

	// AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::FeatureGroup.OnlineStoreSecurityConfig.
	AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__PropertiesSlice = []string{
		AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__PropertiesMap.KmsKeyId,
	}
)

// AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig is a binding for AWS::SageMaker::FeatureGroup.OnlineStoreSecurityConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoresecurityconfig.html
type AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig struct {
	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-onlinestoresecurityconfig.html#cfn-sagemaker-featuregroup-onlinestoresecurityconfig-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`
}

// New__AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig initializes a new AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig.
func New__AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig() AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig {
	return AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig) GetType() string {
	return AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig__Type
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig) Set__KmsKeyId(v cfz.Expression[string]) AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig) SetV__KmsKeyId(v string) AWS_SageMaker_FeatureGroup_OnlineStoreSecurityConfig {
	t.KmsKeyId = cfz.V(v)
	return t
}
