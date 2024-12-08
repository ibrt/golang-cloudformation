// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_Instance_BlockDeviceMapping__Type is the CloudFormation type for AWS::EC2::Instance.BlockDeviceMapping.
	AWS_EC2_Instance_BlockDeviceMapping__Type = "AWS::EC2::Instance.BlockDeviceMapping"
)

var (
	// AWS_EC2_Instance_BlockDeviceMapping__PropertiesMap reports all the CloudFormation properties for AWS::EC2::Instance.BlockDeviceMapping.
	AWS_EC2_Instance_BlockDeviceMapping__PropertiesMap = struct {
		DeviceName  string
		Ebs         string
		NoDevice    string
		VirtualName string
	}{
		DeviceName:  "DeviceName",
		Ebs:         "Ebs",
		NoDevice:    "NoDevice",
		VirtualName: "VirtualName",
	}

	// AWS_EC2_Instance_BlockDeviceMapping__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::Instance.BlockDeviceMapping.
	AWS_EC2_Instance_BlockDeviceMapping__PropertiesSlice = []string{
		AWS_EC2_Instance_BlockDeviceMapping__PropertiesMap.DeviceName,
		AWS_EC2_Instance_BlockDeviceMapping__PropertiesMap.Ebs,
		AWS_EC2_Instance_BlockDeviceMapping__PropertiesMap.NoDevice,
		AWS_EC2_Instance_BlockDeviceMapping__PropertiesMap.VirtualName,
	}
)

// AWS_EC2_Instance_BlockDeviceMapping is a binding for AWS::EC2::Instance.BlockDeviceMapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html
type AWS_EC2_Instance_BlockDeviceMapping struct {
	// DeviceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-devicename
	DeviceName cfz.Expression[string] `json:"DeviceName,omitempty"`

	// Ebs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-ebs
	Ebs cfz.Expression[AWS_EC2_Instance_Ebs] `json:"Ebs,omitempty"`

	// NoDevice is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-nodevice
	NoDevice cfz.Expression[json.RawMessage] `json:"NoDevice,omitempty"`

	// VirtualName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html#cfn-ec2-instance-blockdevicemapping-virtualname
	VirtualName cfz.Expression[string] `json:"VirtualName,omitempty"`
}

// New__AWS_EC2_Instance_BlockDeviceMapping initializes a new AWS_EC2_Instance_BlockDeviceMapping.
func New__AWS_EC2_Instance_BlockDeviceMapping() AWS_EC2_Instance_BlockDeviceMapping {
	return AWS_EC2_Instance_BlockDeviceMapping{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_Instance_BlockDeviceMapping) GetType() string {
	return AWS_EC2_Instance_BlockDeviceMapping__Type
}

// Set__DeviceName updates property "DeviceName".
func (t AWS_EC2_Instance_BlockDeviceMapping) Set__DeviceName(v cfz.Expression[string]) AWS_EC2_Instance_BlockDeviceMapping {
	t.DeviceName = v
	return t
}

// SetV__DeviceName updates property "DeviceName".
func (t AWS_EC2_Instance_BlockDeviceMapping) SetV__DeviceName(v string) AWS_EC2_Instance_BlockDeviceMapping {
	t.DeviceName = cfz.V(v)
	return t
}

// Set__Ebs updates property "Ebs".
func (t AWS_EC2_Instance_BlockDeviceMapping) Set__Ebs(v cfz.Expression[AWS_EC2_Instance_Ebs]) AWS_EC2_Instance_BlockDeviceMapping {
	t.Ebs = v
	return t
}

// SetV__Ebs updates property "Ebs".
func (t AWS_EC2_Instance_BlockDeviceMapping) SetV__Ebs(v AWS_EC2_Instance_Ebs) AWS_EC2_Instance_BlockDeviceMapping {
	t.Ebs = cfz.V(v)
	return t
}

// Set__NoDevice updates property "NoDevice".
func (t AWS_EC2_Instance_BlockDeviceMapping) Set__NoDevice(v cfz.Expression[json.RawMessage]) AWS_EC2_Instance_BlockDeviceMapping {
	t.NoDevice = v
	return t
}

// SetV__NoDevice updates property "NoDevice".
func (t AWS_EC2_Instance_BlockDeviceMapping) SetV__NoDevice(v json.RawMessage) AWS_EC2_Instance_BlockDeviceMapping {
	t.NoDevice = cfz.V(v)
	return t
}

// Set__VirtualName updates property "VirtualName".
func (t AWS_EC2_Instance_BlockDeviceMapping) Set__VirtualName(v cfz.Expression[string]) AWS_EC2_Instance_BlockDeviceMapping {
	t.VirtualName = v
	return t
}

// SetV__VirtualName updates property "VirtualName".
func (t AWS_EC2_Instance_BlockDeviceMapping) SetV__VirtualName(v string) AWS_EC2_Instance_BlockDeviceMapping {
	t.VirtualName = cfz.V(v)
	return t
}
