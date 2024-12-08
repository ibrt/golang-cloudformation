// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ecs

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ECS_TaskDefinition_Device__Type is the CloudFormation type for AWS::ECS::TaskDefinition.Device.
	AWS_ECS_TaskDefinition_Device__Type = "AWS::ECS::TaskDefinition.Device"
)

var (
	// AWS_ECS_TaskDefinition_Device__PropertiesMap reports all the CloudFormation properties for AWS::ECS::TaskDefinition.Device.
	AWS_ECS_TaskDefinition_Device__PropertiesMap = struct {
		ContainerPath string
		HostPath      string
		Permissions   string
	}{
		ContainerPath: "ContainerPath",
		HostPath:      "HostPath",
		Permissions:   "Permissions",
	}

	// AWS_ECS_TaskDefinition_Device__PropertiesSlice reports all the CloudFormation properties for AWS::ECS::TaskDefinition.Device.
	AWS_ECS_TaskDefinition_Device__PropertiesSlice = []string{
		AWS_ECS_TaskDefinition_Device__PropertiesMap.ContainerPath,
		AWS_ECS_TaskDefinition_Device__PropertiesMap.HostPath,
		AWS_ECS_TaskDefinition_Device__PropertiesMap.Permissions,
	}
)

// AWS_ECS_TaskDefinition_Device is a binding for AWS::ECS::TaskDefinition.Device.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html
type AWS_ECS_TaskDefinition_Device struct {
	// ContainerPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-containerpath
	ContainerPath cfz.Expression[string] `json:"ContainerPath,omitempty"`

	// HostPath is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-hostpath
	HostPath cfz.Expression[string] `json:"HostPath,omitempty"`

	// Permissions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-device.html#cfn-ecs-taskdefinition-device-permissions
	Permissions cfz.ExpressionSlice[string] `json:"Permissions,omitempty"`
}

// New__AWS_ECS_TaskDefinition_Device initializes a new AWS_ECS_TaskDefinition_Device.
func New__AWS_ECS_TaskDefinition_Device() AWS_ECS_TaskDefinition_Device {
	return AWS_ECS_TaskDefinition_Device{}
}

// GetType returns the CloudFormation type.
func (AWS_ECS_TaskDefinition_Device) GetType() string {
	return AWS_ECS_TaskDefinition_Device__Type
}

// Set__ContainerPath updates property "ContainerPath".
func (t AWS_ECS_TaskDefinition_Device) Set__ContainerPath(v cfz.Expression[string]) AWS_ECS_TaskDefinition_Device {
	t.ContainerPath = v
	return t
}

// SetV__ContainerPath updates property "ContainerPath".
func (t AWS_ECS_TaskDefinition_Device) SetV__ContainerPath(v string) AWS_ECS_TaskDefinition_Device {
	t.ContainerPath = cfz.V(v)
	return t
}

// Set__HostPath updates property "HostPath".
func (t AWS_ECS_TaskDefinition_Device) Set__HostPath(v cfz.Expression[string]) AWS_ECS_TaskDefinition_Device {
	t.HostPath = v
	return t
}

// SetV__HostPath updates property "HostPath".
func (t AWS_ECS_TaskDefinition_Device) SetV__HostPath(v string) AWS_ECS_TaskDefinition_Device {
	t.HostPath = cfz.V(v)
	return t
}

// Set__Permissions updates property "Permissions".
func (t AWS_ECS_TaskDefinition_Device) Set__Permissions(v cfz.ExpressionSlice[string]) AWS_ECS_TaskDefinition_Device {
	t.Permissions = v
	return t
}

// SetS__Permissions updates property "Permissions".
func (t AWS_ECS_TaskDefinition_Device) SetS__Permissions(v ...cfz.Expression[string]) AWS_ECS_TaskDefinition_Device {
	t.Permissions = cfz.S(v...)
	return t
}

// SetSV__Permissions updates property "Permissions".
func (t AWS_ECS_TaskDefinition_Device) SetSV__Permissions(v ...string) AWS_ECS_TaskDefinition_Device {
	t.Permissions = cfz.SV(v...)
	return t
}
