// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appsync

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppSync_Api_EventLogConfig__Type is the CloudFormation type for AWS::AppSync::Api.EventLogConfig.
	AWS_AppSync_Api_EventLogConfig__Type = "AWS::AppSync::Api.EventLogConfig"
)

var (
	// AWS_AppSync_Api_EventLogConfig__PropertiesMap reports all the CloudFormation properties for AWS::AppSync::Api.EventLogConfig.
	AWS_AppSync_Api_EventLogConfig__PropertiesMap = struct {
		CloudWatchLogsRoleArn string
		LogLevel              string
	}{
		CloudWatchLogsRoleArn: "CloudWatchLogsRoleArn",
		LogLevel:              "LogLevel",
	}

	// AWS_AppSync_Api_EventLogConfig__PropertiesSlice reports all the CloudFormation properties for AWS::AppSync::Api.EventLogConfig.
	AWS_AppSync_Api_EventLogConfig__PropertiesSlice = []string{
		AWS_AppSync_Api_EventLogConfig__PropertiesMap.CloudWatchLogsRoleArn,
		AWS_AppSync_Api_EventLogConfig__PropertiesMap.LogLevel,
	}
)

// AWS_AppSync_Api_EventLogConfig is a binding for AWS::AppSync::Api.EventLogConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventlogconfig.html
type AWS_AppSync_Api_EventLogConfig struct {
	// CloudWatchLogsRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventlogconfig.html#cfn-appsync-api-eventlogconfig-cloudwatchlogsrolearn
	CloudWatchLogsRoleArn cfz.Expression[string] `json:"CloudWatchLogsRoleArn,omitempty"`

	// LogLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-eventlogconfig.html#cfn-appsync-api-eventlogconfig-loglevel
	LogLevel cfz.Expression[string] `json:"LogLevel,omitempty"`
}

// New__AWS_AppSync_Api_EventLogConfig initializes a new AWS_AppSync_Api_EventLogConfig.
func New__AWS_AppSync_Api_EventLogConfig() AWS_AppSync_Api_EventLogConfig {
	return AWS_AppSync_Api_EventLogConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_AppSync_Api_EventLogConfig) GetType() string {
	return AWS_AppSync_Api_EventLogConfig__Type
}

// Set__CloudWatchLogsRoleArn updates property "CloudWatchLogsRoleArn".
func (t AWS_AppSync_Api_EventLogConfig) Set__CloudWatchLogsRoleArn(v cfz.Expression[string]) AWS_AppSync_Api_EventLogConfig {
	t.CloudWatchLogsRoleArn = v
	return t
}

// SetV__CloudWatchLogsRoleArn updates property "CloudWatchLogsRoleArn".
func (t AWS_AppSync_Api_EventLogConfig) SetV__CloudWatchLogsRoleArn(v string) AWS_AppSync_Api_EventLogConfig {
	t.CloudWatchLogsRoleArn = cfz.V(v)
	return t
}

// Set__LogLevel updates property "LogLevel".
func (t AWS_AppSync_Api_EventLogConfig) Set__LogLevel(v cfz.Expression[string]) AWS_AppSync_Api_EventLogConfig {
	t.LogLevel = v
	return t
}

// SetV__LogLevel updates property "LogLevel".
func (t AWS_AppSync_Api_EventLogConfig) SetV__LogLevel(v string) AWS_AppSync_Api_EventLogConfig {
	t.LogLevel = cfz.V(v)
	return t
}
