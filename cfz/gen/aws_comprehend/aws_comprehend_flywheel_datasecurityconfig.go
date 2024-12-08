// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_comprehend

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Comprehend_Flywheel_DataSecurityConfig__Type is the CloudFormation type for AWS::Comprehend::Flywheel.DataSecurityConfig.
	AWS_Comprehend_Flywheel_DataSecurityConfig__Type = "AWS::Comprehend::Flywheel.DataSecurityConfig"
)

var (
	// AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesMap reports all the CloudFormation properties for AWS::Comprehend::Flywheel.DataSecurityConfig.
	AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesMap = struct {
		DataLakeKmsKeyId string
		ModelKmsKeyId    string
		VolumeKmsKeyId   string
		VpcConfig        string
	}{
		DataLakeKmsKeyId: "DataLakeKmsKeyId",
		ModelKmsKeyId:    "ModelKmsKeyId",
		VolumeKmsKeyId:   "VolumeKmsKeyId",
		VpcConfig:        "VpcConfig",
	}

	// AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Comprehend::Flywheel.DataSecurityConfig.
	AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesSlice = []string{
		AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesMap.DataLakeKmsKeyId,
		AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesMap.ModelKmsKeyId,
		AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesMap.VolumeKmsKeyId,
		AWS_Comprehend_Flywheel_DataSecurityConfig__PropertiesMap.VpcConfig,
	}
)

// AWS_Comprehend_Flywheel_DataSecurityConfig is a binding for AWS::Comprehend::Flywheel.DataSecurityConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html
type AWS_Comprehend_Flywheel_DataSecurityConfig struct {
	// DataLakeKmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-datalakekmskeyid
	DataLakeKmsKeyId cfz.Expression[string] `json:"DataLakeKmsKeyId,omitempty"`

	// ModelKmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-modelkmskeyid
	ModelKmsKeyId cfz.Expression[string] `json:"ModelKmsKeyId,omitempty"`

	// VolumeKmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-volumekmskeyid
	VolumeKmsKeyId cfz.Expression[string] `json:"VolumeKmsKeyId,omitempty"`

	// VpcConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-datasecurityconfig.html#cfn-comprehend-flywheel-datasecurityconfig-vpcconfig
	VpcConfig cfz.Expression[AWS_Comprehend_Flywheel_VpcConfig] `json:"VpcConfig,omitempty"`
}

// New__AWS_Comprehend_Flywheel_DataSecurityConfig initializes a new AWS_Comprehend_Flywheel_DataSecurityConfig.
func New__AWS_Comprehend_Flywheel_DataSecurityConfig() AWS_Comprehend_Flywheel_DataSecurityConfig {
	return AWS_Comprehend_Flywheel_DataSecurityConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Comprehend_Flywheel_DataSecurityConfig) GetType() string {
	return AWS_Comprehend_Flywheel_DataSecurityConfig__Type
}

// Set__DataLakeKmsKeyId updates property "DataLakeKmsKeyId".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) Set__DataLakeKmsKeyId(v cfz.Expression[string]) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.DataLakeKmsKeyId = v
	return t
}

// SetV__DataLakeKmsKeyId updates property "DataLakeKmsKeyId".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) SetV__DataLakeKmsKeyId(v string) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.DataLakeKmsKeyId = cfz.V(v)
	return t
}

// Set__ModelKmsKeyId updates property "ModelKmsKeyId".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) Set__ModelKmsKeyId(v cfz.Expression[string]) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.ModelKmsKeyId = v
	return t
}

// SetV__ModelKmsKeyId updates property "ModelKmsKeyId".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) SetV__ModelKmsKeyId(v string) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.ModelKmsKeyId = cfz.V(v)
	return t
}

// Set__VolumeKmsKeyId updates property "VolumeKmsKeyId".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) Set__VolumeKmsKeyId(v cfz.Expression[string]) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.VolumeKmsKeyId = v
	return t
}

// SetV__VolumeKmsKeyId updates property "VolumeKmsKeyId".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) SetV__VolumeKmsKeyId(v string) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.VolumeKmsKeyId = cfz.V(v)
	return t
}

// Set__VpcConfig updates property "VpcConfig".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) Set__VpcConfig(v cfz.Expression[AWS_Comprehend_Flywheel_VpcConfig]) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.VpcConfig = v
	return t
}

// SetV__VpcConfig updates property "VpcConfig".
func (t AWS_Comprehend_Flywheel_DataSecurityConfig) SetV__VpcConfig(v AWS_Comprehend_Flywheel_VpcConfig) AWS_Comprehend_Flywheel_DataSecurityConfig {
	t.VpcConfig = cfz.V(v)
	return t
}
