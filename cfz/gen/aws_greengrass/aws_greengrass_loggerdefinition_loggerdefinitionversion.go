// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__Type is the CloudFormation type for AWS::Greengrass::LoggerDefinition.LoggerDefinitionVersion.
	AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__Type = "AWS::Greengrass::LoggerDefinition.LoggerDefinitionVersion"
)

var (
	// AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::LoggerDefinition.LoggerDefinitionVersion.
	AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__PropertiesMap = struct {
		Loggers string
	}{
		Loggers: "Loggers",
	}

	// AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::LoggerDefinition.LoggerDefinitionVersion.
	AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__PropertiesSlice = []string{
		AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__PropertiesMap.Loggers,
	}
)

// AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion is a binding for AWS::Greengrass::LoggerDefinition.LoggerDefinitionVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-loggerdefinitionversion.html
type AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion struct {
	// Loggers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-loggerdefinitionversion.html#cfn-greengrass-loggerdefinition-loggerdefinitionversion-loggers
	Loggers cfz.ExpressionSlice[AWS_Greengrass_LoggerDefinition_Logger] `json:"Loggers,omitempty"`
}

// New__AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion initializes a new AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion.
func New__AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion() AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion {
	return AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion{}
}

// GetType returns the CloudFormation type.
func (AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion) GetType() string {
	return AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion__Type
}

// Set__Loggers updates property "Loggers".
func (t AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion) Set__Loggers(v cfz.ExpressionSlice[AWS_Greengrass_LoggerDefinition_Logger]) AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion {
	t.Loggers = v
	return t
}

// SetS__Loggers updates property "Loggers".
func (t AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion) SetS__Loggers(v ...cfz.Expression[AWS_Greengrass_LoggerDefinition_Logger]) AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion {
	t.Loggers = cfz.S(v...)
	return t
}

// SetSV__Loggers updates property "Loggers".
func (t AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion) SetSV__Loggers(v ...AWS_Greengrass_LoggerDefinition_Logger) AWS_Greengrass_LoggerDefinition_LoggerDefinitionVersion {
	t.Loggers = cfz.SV(v...)
	return t
}
