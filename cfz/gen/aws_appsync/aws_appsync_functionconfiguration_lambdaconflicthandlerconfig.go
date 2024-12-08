// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appsync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__Type is the CloudFormation type for AWS::AppSync::FunctionConfiguration.LambdaConflictHandlerConfig.
	AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__Type = "AWS::AppSync::FunctionConfiguration.LambdaConflictHandlerConfig"
)

var (
	// AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__PropertiesMap reports all the CloudFormation properties for AWS::AppSync::FunctionConfiguration.LambdaConflictHandlerConfig.
	AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__PropertiesMap = struct {
		LambdaConflictHandlerArn string
	}{
		LambdaConflictHandlerArn: "LambdaConflictHandlerArn",
	}

	// AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__PropertiesSlice reports all the CloudFormation properties for AWS::AppSync::FunctionConfiguration.LambdaConflictHandlerConfig.
	AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__PropertiesSlice = []string{
		AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__PropertiesMap.LambdaConflictHandlerArn,
	}
)

// AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig is a binding for AWS::AppSync::FunctionConfiguration.LambdaConflictHandlerConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-lambdaconflicthandlerconfig.html
type AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig struct {
	// LambdaConflictHandlerArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-lambdaconflicthandlerconfig.html#cfn-appsync-functionconfiguration-lambdaconflicthandlerconfig-lambdaconflicthandlerarn
	LambdaConflictHandlerArn cfz.Expression[string] `json:"LambdaConflictHandlerArn,omitempty"`
}

// New__AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig initializes a new AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig.
func New__AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig() AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig {
	return AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig) GetType() string {
	return AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig__Type
}

// Set__LambdaConflictHandlerArn updates property "LambdaConflictHandlerArn".
func (t AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig) Set__LambdaConflictHandlerArn(v cfz.Expression[string]) AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig {
	t.LambdaConflictHandlerArn = v
	return t
}

// SetV__LambdaConflictHandlerArn updates property "LambdaConflictHandlerArn".
func (t AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig) SetV__LambdaConflictHandlerArn(v string) AWS_AppSync_FunctionConfiguration_LambdaConflictHandlerConfig {
	t.LambdaConflictHandlerArn = cfz.V(v)
	return t
}
