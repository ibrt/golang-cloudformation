// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_NetworkAclEntry_Icmp__Type is the CloudFormation type for AWS::EC2::NetworkAclEntry.Icmp.
	AWS_EC2_NetworkAclEntry_Icmp__Type = "AWS::EC2::NetworkAclEntry.Icmp"
)

var (
	// AWS_EC2_NetworkAclEntry_Icmp__PropertiesMap reports all the CloudFormation properties for AWS::EC2::NetworkAclEntry.Icmp.
	AWS_EC2_NetworkAclEntry_Icmp__PropertiesMap = struct {
		Code string
		Type string
	}{
		Code: "Code",
		Type: "Type",
	}

	// AWS_EC2_NetworkAclEntry_Icmp__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::NetworkAclEntry.Icmp.
	AWS_EC2_NetworkAclEntry_Icmp__PropertiesSlice = []string{
		AWS_EC2_NetworkAclEntry_Icmp__PropertiesMap.Code,
		AWS_EC2_NetworkAclEntry_Icmp__PropertiesMap.Type,
	}
)

// AWS_EC2_NetworkAclEntry_Icmp is a binding for AWS::EC2::NetworkAclEntry.Icmp.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html
type AWS_EC2_NetworkAclEntry_Icmp struct {
	// Code is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-code
	Code cfz.Expression[int32] `json:"Code,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-type
	Type cfz.Expression[int32] `json:"Type,omitempty"`
}

// New__AWS_EC2_NetworkAclEntry_Icmp initializes a new AWS_EC2_NetworkAclEntry_Icmp.
func New__AWS_EC2_NetworkAclEntry_Icmp() AWS_EC2_NetworkAclEntry_Icmp {
	return AWS_EC2_NetworkAclEntry_Icmp{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_NetworkAclEntry_Icmp) GetType() string {
	return AWS_EC2_NetworkAclEntry_Icmp__Type
}

// Set__Code updates property "Code".
func (t AWS_EC2_NetworkAclEntry_Icmp) Set__Code(v cfz.Expression[int32]) AWS_EC2_NetworkAclEntry_Icmp {
	t.Code = v
	return t
}

// SetV__Code updates property "Code".
func (t AWS_EC2_NetworkAclEntry_Icmp) SetV__Code(v int32) AWS_EC2_NetworkAclEntry_Icmp {
	t.Code = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t AWS_EC2_NetworkAclEntry_Icmp) Set__Type(v cfz.Expression[int32]) AWS_EC2_NetworkAclEntry_Icmp {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_EC2_NetworkAclEntry_Icmp) SetV__Type(v int32) AWS_EC2_NetworkAclEntry_Icmp {
	t.Type = cfz.V(v)
	return t
}
