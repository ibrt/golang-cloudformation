// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_AppImageConfig_FileSystemConfig__Type is the CloudFormation type for AWS::SageMaker::AppImageConfig.FileSystemConfig.
	AWS_SageMaker_AppImageConfig_FileSystemConfig__Type = "AWS::SageMaker::AppImageConfig.FileSystemConfig"
)

var (
	// AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::AppImageConfig.FileSystemConfig.
	AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesMap = struct {
		DefaultGid string
		DefaultUid string
		MountPath  string
	}{
		DefaultGid: "DefaultGid",
		DefaultUid: "DefaultUid",
		MountPath:  "MountPath",
	}

	// AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::AppImageConfig.FileSystemConfig.
	AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesSlice = []string{
		AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesMap.DefaultGid,
		AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesMap.DefaultUid,
		AWS_SageMaker_AppImageConfig_FileSystemConfig__PropertiesMap.MountPath,
	}
)

// AWS_SageMaker_AppImageConfig_FileSystemConfig is a binding for AWS::SageMaker::AppImageConfig.FileSystemConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html
type AWS_SageMaker_AppImageConfig_FileSystemConfig struct {
	// DefaultGid is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-defaultgid
	DefaultGid cfz.Expression[int32] `json:"DefaultGid,omitempty"`

	// DefaultUid is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-defaultuid
	DefaultUid cfz.Expression[int32] `json:"DefaultUid,omitempty"`

	// MountPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-mountpath
	MountPath cfz.Expression[string] `json:"MountPath,omitempty"`
}

// New__AWS_SageMaker_AppImageConfig_FileSystemConfig initializes a new AWS_SageMaker_AppImageConfig_FileSystemConfig.
func New__AWS_SageMaker_AppImageConfig_FileSystemConfig() AWS_SageMaker_AppImageConfig_FileSystemConfig {
	return AWS_SageMaker_AppImageConfig_FileSystemConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_AppImageConfig_FileSystemConfig) GetType() string {
	return AWS_SageMaker_AppImageConfig_FileSystemConfig__Type
}

// Set__DefaultGid updates property "DefaultGid".
func (t AWS_SageMaker_AppImageConfig_FileSystemConfig) Set__DefaultGid(v cfz.Expression[int32]) AWS_SageMaker_AppImageConfig_FileSystemConfig {
	t.DefaultGid = v
	return t
}

// SetV__DefaultGid updates property "DefaultGid".
func (t AWS_SageMaker_AppImageConfig_FileSystemConfig) SetV__DefaultGid(v int32) AWS_SageMaker_AppImageConfig_FileSystemConfig {
	t.DefaultGid = cfz.V(v)
	return t
}

// Set__DefaultUid updates property "DefaultUid".
func (t AWS_SageMaker_AppImageConfig_FileSystemConfig) Set__DefaultUid(v cfz.Expression[int32]) AWS_SageMaker_AppImageConfig_FileSystemConfig {
	t.DefaultUid = v
	return t
}

// SetV__DefaultUid updates property "DefaultUid".
func (t AWS_SageMaker_AppImageConfig_FileSystemConfig) SetV__DefaultUid(v int32) AWS_SageMaker_AppImageConfig_FileSystemConfig {
	t.DefaultUid = cfz.V(v)
	return t
}

// Set__MountPath updates property "MountPath".
func (t AWS_SageMaker_AppImageConfig_FileSystemConfig) Set__MountPath(v cfz.Expression[string]) AWS_SageMaker_AppImageConfig_FileSystemConfig {
	t.MountPath = v
	return t
}

// SetV__MountPath updates property "MountPath".
func (t AWS_SageMaker_AppImageConfig_FileSystemConfig) SetV__MountPath(v string) AWS_SageMaker_AppImageConfig_FileSystemConfig {
	t.MountPath = cfz.V(v)
	return t
}
