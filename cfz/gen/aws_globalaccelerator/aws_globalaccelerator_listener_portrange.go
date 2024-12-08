// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_globalaccelerator

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GlobalAccelerator_Listener_PortRange__Type is the CloudFormation type for AWS::GlobalAccelerator::Listener.PortRange.
	AWS_GlobalAccelerator_Listener_PortRange__Type = "AWS::GlobalAccelerator::Listener.PortRange"
)

var (
	// AWS_GlobalAccelerator_Listener_PortRange__PropertiesMap reports all the CloudFormation properties for AWS::GlobalAccelerator::Listener.PortRange.
	AWS_GlobalAccelerator_Listener_PortRange__PropertiesMap = struct {
		FromPort string
		ToPort   string
	}{
		FromPort: "FromPort",
		ToPort:   "ToPort",
	}

	// AWS_GlobalAccelerator_Listener_PortRange__PropertiesSlice reports all the CloudFormation properties for AWS::GlobalAccelerator::Listener.PortRange.
	AWS_GlobalAccelerator_Listener_PortRange__PropertiesSlice = []string{
		AWS_GlobalAccelerator_Listener_PortRange__PropertiesMap.FromPort,
		AWS_GlobalAccelerator_Listener_PortRange__PropertiesMap.ToPort,
	}
)

// AWS_GlobalAccelerator_Listener_PortRange is a binding for AWS::GlobalAccelerator::Listener.PortRange.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html
type AWS_GlobalAccelerator_Listener_PortRange struct {
	// FromPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-fromport
	FromPort cfz.Expression[int32] `json:"FromPort,omitempty"`

	// ToPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-globalaccelerator-listener-portrange.html#cfn-globalaccelerator-listener-portrange-toport
	ToPort cfz.Expression[int32] `json:"ToPort,omitempty"`
}

// New__AWS_GlobalAccelerator_Listener_PortRange initializes a new AWS_GlobalAccelerator_Listener_PortRange.
func New__AWS_GlobalAccelerator_Listener_PortRange() AWS_GlobalAccelerator_Listener_PortRange {
	return AWS_GlobalAccelerator_Listener_PortRange{}
}

// GetType returns the CloudFormation type.
func (AWS_GlobalAccelerator_Listener_PortRange) GetType() string {
	return AWS_GlobalAccelerator_Listener_PortRange__Type
}

// Set__FromPort updates property "FromPort".
func (t AWS_GlobalAccelerator_Listener_PortRange) Set__FromPort(v cfz.Expression[int32]) AWS_GlobalAccelerator_Listener_PortRange {
	t.FromPort = v
	return t
}

// SetV__FromPort updates property "FromPort".
func (t AWS_GlobalAccelerator_Listener_PortRange) SetV__FromPort(v int32) AWS_GlobalAccelerator_Listener_PortRange {
	t.FromPort = cfz.V(v)
	return t
}

// Set__ToPort updates property "ToPort".
func (t AWS_GlobalAccelerator_Listener_PortRange) Set__ToPort(v cfz.Expression[int32]) AWS_GlobalAccelerator_Listener_PortRange {
	t.ToPort = v
	return t
}

// SetV__ToPort updates property "ToPort".
func (t AWS_GlobalAccelerator_Listener_PortRange) SetV__ToPort(v int32) AWS_GlobalAccelerator_Listener_PortRange {
	t.ToPort = cfz.V(v)
	return t
}
