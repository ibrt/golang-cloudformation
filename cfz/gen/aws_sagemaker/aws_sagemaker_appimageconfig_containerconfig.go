// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SageMaker_AppImageConfig_ContainerConfig__Type is the CloudFormation type for AWS::SageMaker::AppImageConfig.ContainerConfig.
	AWS_SageMaker_AppImageConfig_ContainerConfig__Type = "AWS::SageMaker::AppImageConfig.ContainerConfig"
)

var (
	// AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::AppImageConfig.ContainerConfig.
	AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesMap = struct {
		ContainerArguments            string
		ContainerEntrypoint           string
		ContainerEnvironmentVariables string
	}{
		ContainerArguments:            "ContainerArguments",
		ContainerEntrypoint:           "ContainerEntrypoint",
		ContainerEnvironmentVariables: "ContainerEnvironmentVariables",
	}

	// AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::AppImageConfig.ContainerConfig.
	AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesSlice = []string{
		AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesMap.ContainerArguments,
		AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesMap.ContainerEntrypoint,
		AWS_SageMaker_AppImageConfig_ContainerConfig__PropertiesMap.ContainerEnvironmentVariables,
	}
)

// AWS_SageMaker_AppImageConfig_ContainerConfig is a binding for AWS::SageMaker::AppImageConfig.ContainerConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html
type AWS_SageMaker_AppImageConfig_ContainerConfig struct {
	// ContainerArguments is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html#cfn-sagemaker-appimageconfig-containerconfig-containerarguments
	ContainerArguments cfz.ExpressionSlice[string] `json:"ContainerArguments,omitempty"`

	// ContainerEntrypoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html#cfn-sagemaker-appimageconfig-containerconfig-containerentrypoint
	ContainerEntrypoint cfz.ExpressionSlice[string] `json:"ContainerEntrypoint,omitempty"`

	// ContainerEnvironmentVariables is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-containerconfig.html#cfn-sagemaker-appimageconfig-containerconfig-containerenvironmentvariables
	ContainerEnvironmentVariables cfz.ExpressionSlice[AWS_SageMaker_AppImageConfig_CustomImageContainerEnvironmentVariable] `json:"ContainerEnvironmentVariables,omitempty"`
}

// New__AWS_SageMaker_AppImageConfig_ContainerConfig initializes a new AWS_SageMaker_AppImageConfig_ContainerConfig.
func New__AWS_SageMaker_AppImageConfig_ContainerConfig() AWS_SageMaker_AppImageConfig_ContainerConfig {
	return AWS_SageMaker_AppImageConfig_ContainerConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SageMaker_AppImageConfig_ContainerConfig) GetType() string {
	return AWS_SageMaker_AppImageConfig_ContainerConfig__Type
}

// Set__ContainerArguments updates property "ContainerArguments".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) Set__ContainerArguments(v cfz.ExpressionSlice[string]) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerArguments = v
	return t
}

// SetS__ContainerArguments updates property "ContainerArguments".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) SetS__ContainerArguments(v ...cfz.Expression[string]) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerArguments = cfz.S(v...)
	return t
}

// SetSV__ContainerArguments updates property "ContainerArguments".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) SetSV__ContainerArguments(v ...string) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerArguments = cfz.SV(v...)
	return t
}

// Set__ContainerEntrypoint updates property "ContainerEntrypoint".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) Set__ContainerEntrypoint(v cfz.ExpressionSlice[string]) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerEntrypoint = v
	return t
}

// SetS__ContainerEntrypoint updates property "ContainerEntrypoint".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) SetS__ContainerEntrypoint(v ...cfz.Expression[string]) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerEntrypoint = cfz.S(v...)
	return t
}

// SetSV__ContainerEntrypoint updates property "ContainerEntrypoint".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) SetSV__ContainerEntrypoint(v ...string) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerEntrypoint = cfz.SV(v...)
	return t
}

// Set__ContainerEnvironmentVariables updates property "ContainerEnvironmentVariables".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) Set__ContainerEnvironmentVariables(v cfz.ExpressionSlice[AWS_SageMaker_AppImageConfig_CustomImageContainerEnvironmentVariable]) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerEnvironmentVariables = v
	return t
}

// SetS__ContainerEnvironmentVariables updates property "ContainerEnvironmentVariables".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) SetS__ContainerEnvironmentVariables(v ...cfz.Expression[AWS_SageMaker_AppImageConfig_CustomImageContainerEnvironmentVariable]) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerEnvironmentVariables = cfz.S(v...)
	return t
}

// SetSV__ContainerEnvironmentVariables updates property "ContainerEnvironmentVariables".
func (t AWS_SageMaker_AppImageConfig_ContainerConfig) SetSV__ContainerEnvironmentVariables(v ...AWS_SageMaker_AppImageConfig_CustomImageContainerEnvironmentVariable) AWS_SageMaker_AppImageConfig_ContainerConfig {
	t.ContainerEnvironmentVariables = cfz.SV(v...)
	return t
}
