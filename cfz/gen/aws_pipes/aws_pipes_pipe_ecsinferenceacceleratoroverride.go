// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__Type is the CloudFormation type for AWS::Pipes::Pipe.EcsInferenceAcceleratorOverride.
	AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__Type = "AWS::Pipes::Pipe.EcsInferenceAcceleratorOverride"
)

var (
	// AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.EcsInferenceAcceleratorOverride.
	AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__PropertiesMap = struct {
		DeviceName string
		DeviceType string
	}{
		DeviceName: "DeviceName",
		DeviceType: "DeviceType",
	}

	// AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.EcsInferenceAcceleratorOverride.
	AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__PropertiesSlice = []string{
		AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__PropertiesMap.DeviceName,
		AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__PropertiesMap.DeviceType,
	}
)

// AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride is a binding for AWS::Pipes::Pipe.EcsInferenceAcceleratorOverride.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecsinferenceacceleratoroverride.html
type AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride struct {
	// DeviceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecsinferenceacceleratoroverride.html#cfn-pipes-pipe-ecsinferenceacceleratoroverride-devicename
	DeviceName cfz.Expression[string] `json:"DeviceName,omitempty"`

	// DeviceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecsinferenceacceleratoroverride.html#cfn-pipes-pipe-ecsinferenceacceleratoroverride-devicetype
	DeviceType cfz.Expression[string] `json:"DeviceType,omitempty"`
}

// New__AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride initializes a new AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride.
func New__AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride() AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride {
	return AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride) GetType() string {
	return AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride__Type
}

// Set__DeviceName updates property "DeviceName".
func (t AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride) Set__DeviceName(v cfz.Expression[string]) AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride {
	t.DeviceName = v
	return t
}

// SetV__DeviceName updates property "DeviceName".
func (t AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride) SetV__DeviceName(v string) AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride {
	t.DeviceName = cfz.V(v)
	return t
}

// Set__DeviceType updates property "DeviceType".
func (t AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride) Set__DeviceType(v cfz.Expression[string]) AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride {
	t.DeviceType = v
	return t
}

// SetV__DeviceType updates property "DeviceType".
func (t AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride) SetV__DeviceType(v string) AWS_Pipes_Pipe_EcsInferenceAcceleratorOverride {
	t.DeviceType = cfz.V(v)
	return t
}
