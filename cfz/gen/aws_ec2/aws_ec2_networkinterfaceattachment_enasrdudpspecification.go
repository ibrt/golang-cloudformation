// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__Type is the CloudFormation type for AWS::EC2::NetworkInterfaceAttachment.EnaSrdUdpSpecification.
	AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__Type = "AWS::EC2::NetworkInterfaceAttachment.EnaSrdUdpSpecification"
)

var (
	// AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkInterfaceAttachment.EnaSrdUdpSpecification.
	AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__PropertiesMap = struct {
		EnaSrdUdpEnabled string
	}{
		EnaSrdUdpEnabled: "EnaSrdUdpEnabled",
	}

	// AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkInterfaceAttachment.EnaSrdUdpSpecification.
	AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__PropertiesSlice = []string{
		AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__PropertiesMap.EnaSrdUdpEnabled,
	}
)

// AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification is a binding for AWS::EC2::NetworkInterfaceAttachment.EnaSrdUdpSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdudpspecification.html
type AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification struct {
	// EnaSrdUdpEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterfaceattachment-enasrdudpspecification.html#cfn-ec2-networkinterfaceattachment-enasrdudpspecification-enasrdudpenabled
	EnaSrdUdpEnabled cfz.Expression[bool] `json:"EnaSrdUdpEnabled,omitempty"`
}

// New__AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification initializes a new AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification.
func New__AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification() AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification {
	return AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification) GetType() string {
	return AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification__Type
}

// Set__EnaSrdUdpEnabled updates property "EnaSrdUdpEnabled".
func (t AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification) Set__EnaSrdUdpEnabled(v cfz.Expression[bool]) AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification {
	t.EnaSrdUdpEnabled = v
	return t
}

// SetV__EnaSrdUdpEnabled updates property "EnaSrdUdpEnabled".
func (t AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification) SetV__EnaSrdUdpEnabled(v bool) AWS_EC2_NetworkInterfaceAttachment_EnaSrdUdpSpecification {
	t.EnaSrdUdpEnabled = cfz.V(v)
	return t
}
