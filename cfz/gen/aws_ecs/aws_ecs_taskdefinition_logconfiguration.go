// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ecs

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ECS_TaskDefinition_LogConfiguration__Type is the CloudFormation type for AWS::ECS::TaskDefinition.LogConfiguration.
	AWS_ECS_TaskDefinition_LogConfiguration__Type = "AWS::ECS::TaskDefinition.LogConfiguration"
)

var (
	// AWS_ECS_TaskDefinition_LogConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ECS::TaskDefinition.LogConfiguration.
	AWS_ECS_TaskDefinition_LogConfiguration__PropertiesMap = struct {
		LogDriver     string
		Options       string
		SecretOptions string
	}{
		LogDriver:     "LogDriver",
		Options:       "Options",
		SecretOptions: "SecretOptions",
	}

	// AWS_ECS_TaskDefinition_LogConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ECS::TaskDefinition.LogConfiguration.
	AWS_ECS_TaskDefinition_LogConfiguration__PropertiesSlice = []string{
		AWS_ECS_TaskDefinition_LogConfiguration__PropertiesMap.LogDriver,
		AWS_ECS_TaskDefinition_LogConfiguration__PropertiesMap.Options,
		AWS_ECS_TaskDefinition_LogConfiguration__PropertiesMap.SecretOptions,
	}
)

// AWS_ECS_TaskDefinition_LogConfiguration is a binding for AWS::ECS::TaskDefinition.LogConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html
type AWS_ECS_TaskDefinition_LogConfiguration struct {
	// LogDriver is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html#cfn-ecs-taskdefinition-logconfiguration-logdriver
	LogDriver cfz.Expression[string] `json:"LogDriver,omitempty"`

	// Options is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html#cfn-ecs-taskdefinition-logconfiguration-options
	Options cfz.ExpressionMap[string] `json:"Options,omitempty"`

	// SecretOptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html#cfn-ecs-taskdefinition-logconfiguration-secretoptions
	SecretOptions cfz.ExpressionSlice[AWS_ECS_TaskDefinition_Secret] `json:"SecretOptions,omitempty"`
}

// New__AWS_ECS_TaskDefinition_LogConfiguration initializes a new AWS_ECS_TaskDefinition_LogConfiguration.
func New__AWS_ECS_TaskDefinition_LogConfiguration() AWS_ECS_TaskDefinition_LogConfiguration {
	return AWS_ECS_TaskDefinition_LogConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ECS_TaskDefinition_LogConfiguration) GetType() string {
	return AWS_ECS_TaskDefinition_LogConfiguration__Type
}

// Set__LogDriver updates property "LogDriver".
func (t AWS_ECS_TaskDefinition_LogConfiguration) Set__LogDriver(v cfz.Expression[string]) AWS_ECS_TaskDefinition_LogConfiguration {
	t.LogDriver = v
	return t
}

// SetV__LogDriver updates property "LogDriver".
func (t AWS_ECS_TaskDefinition_LogConfiguration) SetV__LogDriver(v string) AWS_ECS_TaskDefinition_LogConfiguration {
	t.LogDriver = cfz.V(v)
	return t
}

// Set__Options updates property "Options".
func (t AWS_ECS_TaskDefinition_LogConfiguration) Set__Options(v cfz.ExpressionMap[string]) AWS_ECS_TaskDefinition_LogConfiguration {
	t.Options = v
	return t
}

// SetM__Options updates property "Options".
func (t AWS_ECS_TaskDefinition_LogConfiguration) SetM__Options(v ...map[string]cfz.Expression[string]) AWS_ECS_TaskDefinition_LogConfiguration {
	t.Options = cfz.M(v...)
	return t
}

// SetMV__Options updates property "Options".
func (t AWS_ECS_TaskDefinition_LogConfiguration) SetMV__Options(v ...map[string]string) AWS_ECS_TaskDefinition_LogConfiguration {
	t.Options = cfz.MV(v...)
	return t
}

// Set__SecretOptions updates property "SecretOptions".
func (t AWS_ECS_TaskDefinition_LogConfiguration) Set__SecretOptions(v cfz.ExpressionSlice[AWS_ECS_TaskDefinition_Secret]) AWS_ECS_TaskDefinition_LogConfiguration {
	t.SecretOptions = v
	return t
}

// SetS__SecretOptions updates property "SecretOptions".
func (t AWS_ECS_TaskDefinition_LogConfiguration) SetS__SecretOptions(v ...cfz.Expression[AWS_ECS_TaskDefinition_Secret]) AWS_ECS_TaskDefinition_LogConfiguration {
	t.SecretOptions = cfz.S(v...)
	return t
}

// SetSV__SecretOptions updates property "SecretOptions".
func (t AWS_ECS_TaskDefinition_LogConfiguration) SetSV__SecretOptions(v ...AWS_ECS_TaskDefinition_Secret) AWS_ECS_TaskDefinition_LogConfiguration {
	t.SecretOptions = cfz.SV(v...)
	return t
}
