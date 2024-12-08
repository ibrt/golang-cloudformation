// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__Type is the CloudFormation type for AWS::Greengrass::DeviceDefinition.DeviceDefinitionVersion.
	AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__Type = "AWS::Greengrass::DeviceDefinition.DeviceDefinitionVersion"
)

var (
	// AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::DeviceDefinition.DeviceDefinitionVersion.
	AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__PropertiesMap = struct {
		Devices string
	}{
		Devices: "Devices",
	}

	// AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::DeviceDefinition.DeviceDefinitionVersion.
	AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__PropertiesSlice = []string{
		AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__PropertiesMap.Devices,
	}
)

// AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion is a binding for AWS::Greengrass::DeviceDefinition.DeviceDefinitionVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.html
type AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion struct {
	// Devices is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-devicedefinition-devicedefinitionversion.html#cfn-greengrass-devicedefinition-devicedefinitionversion-devices
	Devices cfz.ExpressionSlice[AWS_Greengrass_DeviceDefinition_Device] `json:"Devices,omitempty"`
}

// New__AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion initializes a new AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion.
func New__AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion() AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion {
	return AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion{}
}

// GetType returns the CloudFormation type.
func (AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion) GetType() string {
	return AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion__Type
}

// Set__Devices updates property "Devices".
func (t AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion) Set__Devices(v cfz.ExpressionSlice[AWS_Greengrass_DeviceDefinition_Device]) AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion {
	t.Devices = v
	return t
}

// SetS__Devices updates property "Devices".
func (t AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion) SetS__Devices(v ...cfz.Expression[AWS_Greengrass_DeviceDefinition_Device]) AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion {
	t.Devices = cfz.S(v...)
	return t
}

// SetSV__Devices updates property "Devices".
func (t AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion) SetSV__Devices(v ...AWS_Greengrass_DeviceDefinition_Device) AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion {
	t.Devices = cfz.SV(v...)
	return t
}
