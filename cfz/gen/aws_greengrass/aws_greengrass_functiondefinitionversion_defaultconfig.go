// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__Type is the CloudFormation type for AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig.
	AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__Type = "AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig"
)

var (
	// AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig.
	AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__PropertiesMap = struct {
		Execution string
	}{
		Execution: "Execution",
	}

	// AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig.
	AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__PropertiesSlice = []string{
		AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__PropertiesMap.Execution,
	}
)

// AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig is a binding for AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html
type AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig struct {
	// Execution is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html#cfn-greengrass-functiondefinitionversion-defaultconfig-execution
	Execution cfz.Expression[AWS_Greengrass_FunctionDefinitionVersion_Execution] `json:"Execution,omitempty"`
}

// New__AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig initializes a new AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig.
func New__AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig() AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig {
	return AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig) GetType() string {
	return AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig__Type
}

// Set__Execution updates property "Execution".
func (t AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig) Set__Execution(v cfz.Expression[AWS_Greengrass_FunctionDefinitionVersion_Execution]) AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig {
	t.Execution = v
	return t
}

// SetV__Execution updates property "Execution".
func (t AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig) SetV__Execution(v AWS_Greengrass_FunctionDefinitionVersion_Execution) AWS_Greengrass_FunctionDefinitionVersion_DefaultConfig {
	t.Execution = cfz.V(v)
	return t
}
