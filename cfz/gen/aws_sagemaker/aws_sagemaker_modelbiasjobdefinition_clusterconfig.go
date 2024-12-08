// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__Type is the CloudFormation type for AWS::SageMaker::ModelBiasJobDefinition.ClusterConfig.
	AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__Type = "AWS::SageMaker::ModelBiasJobDefinition.ClusterConfig"
)

var (
	// AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::ModelBiasJobDefinition.ClusterConfig.
	AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesMap = struct {
		InstanceCount  string
		InstanceType   string
		VolumeKmsKeyId string
		VolumeSizeInGB string
	}{
		InstanceCount:  "InstanceCount",
		InstanceType:   "InstanceType",
		VolumeKmsKeyId: "VolumeKmsKeyId",
		VolumeSizeInGB: "VolumeSizeInGB",
	}

	// AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::ModelBiasJobDefinition.ClusterConfig.
	AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesSlice = []string{
		AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesMap.InstanceCount,
		AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesMap.InstanceType,
		AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesMap.VolumeKmsKeyId,
		AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__PropertiesMap.VolumeSizeInGB,
	}
)

// AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig is a binding for AWS::SageMaker::ModelBiasJobDefinition.ClusterConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-clusterconfig.html
type AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig struct {
	// InstanceCount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-clusterconfig.html#cfn-sagemaker-modelbiasjobdefinition-clusterconfig-instancecount
	InstanceCount cfz.Expression[int32] `json:"InstanceCount,omitempty"`

	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-clusterconfig.html#cfn-sagemaker-modelbiasjobdefinition-clusterconfig-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`

	// VolumeKmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-clusterconfig.html#cfn-sagemaker-modelbiasjobdefinition-clusterconfig-volumekmskeyid
	VolumeKmsKeyId cfz.Expression[string] `json:"VolumeKmsKeyId,omitempty"`

	// VolumeSizeInGB is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-clusterconfig.html#cfn-sagemaker-modelbiasjobdefinition-clusterconfig-volumesizeingb
	VolumeSizeInGB cfz.Expression[int32] `json:"VolumeSizeInGB,omitempty"`
}

// New__AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig initializes a new AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig.
func New__AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig() AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	return AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) GetType() string {
	return AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig__Type
}

// Set__InstanceCount updates property "InstanceCount".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) Set__InstanceCount(v cfz.Expression[int32]) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.InstanceCount = v
	return t
}

// SetV__InstanceCount updates property "InstanceCount".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) SetV__InstanceCount(v int32) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.InstanceCount = cfz.V(v)
	return t
}

// Set__InstanceType updates property "InstanceType".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) Set__InstanceType(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) SetV__InstanceType(v string) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.InstanceType = cfz.V(v)
	return t
}

// Set__VolumeKmsKeyId updates property "VolumeKmsKeyId".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) Set__VolumeKmsKeyId(v cfz.Expression[string]) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.VolumeKmsKeyId = v
	return t
}

// SetV__VolumeKmsKeyId updates property "VolumeKmsKeyId".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) SetV__VolumeKmsKeyId(v string) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.VolumeKmsKeyId = cfz.V(v)
	return t
}

// Set__VolumeSizeInGB updates property "VolumeSizeInGB".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) Set__VolumeSizeInGB(v cfz.Expression[int32]) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.VolumeSizeInGB = v
	return t
}

// SetV__VolumeSizeInGB updates property "VolumeSizeInGB".
func (t AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig) SetV__VolumeSizeInGB(v int32) AWS_SageMaker_ModelBiasJobDefinition_ClusterConfig {
	t.VolumeSizeInGB = cfz.V(v)
	return t
}
