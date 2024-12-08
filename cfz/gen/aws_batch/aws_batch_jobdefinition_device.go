// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_batch

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Batch_JobDefinition_Device__Type is the CloudFormation type for AWS::Batch::JobDefinition.Device.
	AWS_Batch_JobDefinition_Device__Type = "AWS::Batch::JobDefinition.Device"
)

var (
	// AWS_Batch_JobDefinition_Device__PropertiesMap reports all the CloudFormation properties for AWS::Batch::JobDefinition.Device.
	AWS_Batch_JobDefinition_Device__PropertiesMap = struct {
		ContainerPath string
		HostPath      string
		Permissions   string
	}{
		ContainerPath: "ContainerPath",
		HostPath:      "HostPath",
		Permissions:   "Permissions",
	}

	// AWS_Batch_JobDefinition_Device__PropertiesSlice reports all the CloudFormation properties for AWS::Batch::JobDefinition.Device.
	AWS_Batch_JobDefinition_Device__PropertiesSlice = []string{
		AWS_Batch_JobDefinition_Device__PropertiesMap.ContainerPath,
		AWS_Batch_JobDefinition_Device__PropertiesMap.HostPath,
		AWS_Batch_JobDefinition_Device__PropertiesMap.Permissions,
	}
)

// AWS_Batch_JobDefinition_Device is a binding for AWS::Batch::JobDefinition.Device.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html
type AWS_Batch_JobDefinition_Device struct {
	// ContainerPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html#cfn-batch-jobdefinition-device-containerpath
	ContainerPath cfz.Expression[string] `json:"ContainerPath,omitempty"`

	// HostPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html#cfn-batch-jobdefinition-device-hostpath
	HostPath cfz.Expression[string] `json:"HostPath,omitempty"`

	// Permissions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-device.html#cfn-batch-jobdefinition-device-permissions
	Permissions cfz.ExpressionSlice[string] `json:"Permissions,omitempty"`
}

// New__AWS_Batch_JobDefinition_Device initializes a new AWS_Batch_JobDefinition_Device.
func New__AWS_Batch_JobDefinition_Device() AWS_Batch_JobDefinition_Device {
	return AWS_Batch_JobDefinition_Device{}
}

// GetType returns the CloudFormation type.
func (AWS_Batch_JobDefinition_Device) GetType() string {
	return AWS_Batch_JobDefinition_Device__Type
}

// Set__ContainerPath updates property "ContainerPath".
func (t AWS_Batch_JobDefinition_Device) Set__ContainerPath(v cfz.Expression[string]) AWS_Batch_JobDefinition_Device {
	t.ContainerPath = v
	return t
}

// SetV__ContainerPath updates property "ContainerPath".
func (t AWS_Batch_JobDefinition_Device) SetV__ContainerPath(v string) AWS_Batch_JobDefinition_Device {
	t.ContainerPath = cfz.V(v)
	return t
}

// Set__HostPath updates property "HostPath".
func (t AWS_Batch_JobDefinition_Device) Set__HostPath(v cfz.Expression[string]) AWS_Batch_JobDefinition_Device {
	t.HostPath = v
	return t
}

// SetV__HostPath updates property "HostPath".
func (t AWS_Batch_JobDefinition_Device) SetV__HostPath(v string) AWS_Batch_JobDefinition_Device {
	t.HostPath = cfz.V(v)
	return t
}

// Set__Permissions updates property "Permissions".
func (t AWS_Batch_JobDefinition_Device) Set__Permissions(v cfz.ExpressionSlice[string]) AWS_Batch_JobDefinition_Device {
	t.Permissions = v
	return t
}

// SetS__Permissions updates property "Permissions".
func (t AWS_Batch_JobDefinition_Device) SetS__Permissions(v ...cfz.Expression[string]) AWS_Batch_JobDefinition_Device {
	t.Permissions = cfz.S(v...)
	return t
}

// SetSV__Permissions updates property "Permissions".
func (t AWS_Batch_JobDefinition_Device) SetSV__Permissions(v ...string) AWS_Batch_JobDefinition_Device {
	t.Permissions = cfz.SV(v...)
	return t
}
