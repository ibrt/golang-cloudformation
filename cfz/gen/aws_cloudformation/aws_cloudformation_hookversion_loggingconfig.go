// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudformation

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFormation_HookVersion_LoggingConfig__Type is the CloudFormation type for AWS::CloudFormation::HookVersion.LoggingConfig.
	AWS_CloudFormation_HookVersion_LoggingConfig__Type = "AWS::CloudFormation::HookVersion.LoggingConfig"
)

var (
	// AWS_CloudFormation_HookVersion_LoggingConfig__PropertiesMap reports all the CloudFormation properties for AWS::CloudFormation::HookVersion.LoggingConfig.
	AWS_CloudFormation_HookVersion_LoggingConfig__PropertiesMap = struct {
		LogGroupName string
		LogRoleArn   string
	}{
		LogGroupName: "LogGroupName",
		LogRoleArn:   "LogRoleArn",
	}

	// AWS_CloudFormation_HookVersion_LoggingConfig__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFormation::HookVersion.LoggingConfig.
	AWS_CloudFormation_HookVersion_LoggingConfig__PropertiesSlice = []string{
		AWS_CloudFormation_HookVersion_LoggingConfig__PropertiesMap.LogGroupName,
		AWS_CloudFormation_HookVersion_LoggingConfig__PropertiesMap.LogRoleArn,
	}
)

// AWS_CloudFormation_HookVersion_LoggingConfig is a binding for AWS::CloudFormation::HookVersion.LoggingConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-hookversion-loggingconfig.html
type AWS_CloudFormation_HookVersion_LoggingConfig struct {
	// LogGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-hookversion-loggingconfig.html#cfn-cloudformation-hookversion-loggingconfig-loggroupname
	LogGroupName cfz.Expression[string] `json:"LogGroupName,omitempty"`

	// LogRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-hookversion-loggingconfig.html#cfn-cloudformation-hookversion-loggingconfig-logrolearn
	LogRoleArn cfz.Expression[string] `json:"LogRoleArn,omitempty"`
}

// New__AWS_CloudFormation_HookVersion_LoggingConfig initializes a new AWS_CloudFormation_HookVersion_LoggingConfig.
func New__AWS_CloudFormation_HookVersion_LoggingConfig() AWS_CloudFormation_HookVersion_LoggingConfig {
	return AWS_CloudFormation_HookVersion_LoggingConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFormation_HookVersion_LoggingConfig) GetType() string {
	return AWS_CloudFormation_HookVersion_LoggingConfig__Type
}

// Set__LogGroupName updates property "LogGroupName".
func (t AWS_CloudFormation_HookVersion_LoggingConfig) Set__LogGroupName(v cfz.Expression[string]) AWS_CloudFormation_HookVersion_LoggingConfig {
	t.LogGroupName = v
	return t
}

// SetV__LogGroupName updates property "LogGroupName".
func (t AWS_CloudFormation_HookVersion_LoggingConfig) SetV__LogGroupName(v string) AWS_CloudFormation_HookVersion_LoggingConfig {
	t.LogGroupName = cfz.V(v)
	return t
}

// Set__LogRoleArn updates property "LogRoleArn".
func (t AWS_CloudFormation_HookVersion_LoggingConfig) Set__LogRoleArn(v cfz.Expression[string]) AWS_CloudFormation_HookVersion_LoggingConfig {
	t.LogRoleArn = v
	return t
}

// SetV__LogRoleArn updates property "LogRoleArn".
func (t AWS_CloudFormation_HookVersion_LoggingConfig) SetV__LogRoleArn(v string) AWS_CloudFormation_HookVersion_LoggingConfig {
	t.LogRoleArn = cfz.V(v)
	return t
}
