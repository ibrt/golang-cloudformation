// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_SecurityGroup_Egress__Type is the CloudFormation type for AWS::EC2::SecurityGroup.Egress.
	AWS_EC2_SecurityGroup_Egress__Type = "AWS::EC2::SecurityGroup.Egress"
)

var (
	// AWS_EC2_SecurityGroup_Egress__PropertiesMap reports all the CloudFormation properties for AWS::EC2::SecurityGroup.Egress.
	AWS_EC2_SecurityGroup_Egress__PropertiesMap = struct {
		CidrIp                     string
		CidrIpv6                   string
		Description                string
		DestinationPrefixListId    string
		DestinationSecurityGroupId string
		FromPort                   string
		IpProtocol                 string
		ToPort                     string
	}{
		CidrIp:                     "CidrIp",
		CidrIpv6:                   "CidrIpv6",
		Description:                "Description",
		DestinationPrefixListId:    "DestinationPrefixListId",
		DestinationSecurityGroupId: "DestinationSecurityGroupId",
		FromPort:                   "FromPort",
		IpProtocol:                 "IpProtocol",
		ToPort:                     "ToPort",
	}

	// AWS_EC2_SecurityGroup_Egress__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::SecurityGroup.Egress.
	AWS_EC2_SecurityGroup_Egress__PropertiesSlice = []string{
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.CidrIp,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.CidrIpv6,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.Description,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.DestinationPrefixListId,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.DestinationSecurityGroupId,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.FromPort,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.IpProtocol,
		AWS_EC2_SecurityGroup_Egress__PropertiesMap.ToPort,
	}
)

// AWS_EC2_SecurityGroup_Egress is a binding for AWS::EC2::SecurityGroup.Egress.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html
type AWS_EC2_SecurityGroup_Egress struct {
	// CidrIp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-cidrip
	CidrIp cfz.Expression[string] `json:"CidrIp,omitempty"`

	// CidrIpv6 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-cidripv6
	CidrIpv6 cfz.Expression[string] `json:"CidrIpv6,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DestinationPrefixListId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-destinationprefixlistid
	DestinationPrefixListId cfz.Expression[string] `json:"DestinationPrefixListId,omitempty"`

	// DestinationSecurityGroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-destinationsecuritygroupid
	DestinationSecurityGroupId cfz.Expression[string] `json:"DestinationSecurityGroupId,omitempty"`

	// FromPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-fromport
	FromPort cfz.Expression[int32] `json:"FromPort,omitempty"`

	// IpProtocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-ipprotocol
	IpProtocol cfz.Expression[string] `json:"IpProtocol,omitempty"`

	// ToPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-toport
	ToPort cfz.Expression[int32] `json:"ToPort,omitempty"`
}

// New__AWS_EC2_SecurityGroup_Egress initializes a new AWS_EC2_SecurityGroup_Egress.
func New__AWS_EC2_SecurityGroup_Egress() AWS_EC2_SecurityGroup_Egress {
	return AWS_EC2_SecurityGroup_Egress{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_SecurityGroup_Egress) GetType() string {
	return AWS_EC2_SecurityGroup_Egress__Type
}

// Set__CidrIp updates property "CidrIp".
func (t AWS_EC2_SecurityGroup_Egress) Set__CidrIp(v cfz.Expression[string]) AWS_EC2_SecurityGroup_Egress {
	t.CidrIp = v
	return t
}

// SetV__CidrIp updates property "CidrIp".
func (t AWS_EC2_SecurityGroup_Egress) SetV__CidrIp(v string) AWS_EC2_SecurityGroup_Egress {
	t.CidrIp = cfz.V(v)
	return t
}

// Set__CidrIpv6 updates property "CidrIpv6".
func (t AWS_EC2_SecurityGroup_Egress) Set__CidrIpv6(v cfz.Expression[string]) AWS_EC2_SecurityGroup_Egress {
	t.CidrIpv6 = v
	return t
}

// SetV__CidrIpv6 updates property "CidrIpv6".
func (t AWS_EC2_SecurityGroup_Egress) SetV__CidrIpv6(v string) AWS_EC2_SecurityGroup_Egress {
	t.CidrIpv6 = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_EC2_SecurityGroup_Egress) Set__Description(v cfz.Expression[string]) AWS_EC2_SecurityGroup_Egress {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_EC2_SecurityGroup_Egress) SetV__Description(v string) AWS_EC2_SecurityGroup_Egress {
	t.Description = cfz.V(v)
	return t
}

// Set__DestinationPrefixListId updates property "DestinationPrefixListId".
func (t AWS_EC2_SecurityGroup_Egress) Set__DestinationPrefixListId(v cfz.Expression[string]) AWS_EC2_SecurityGroup_Egress {
	t.DestinationPrefixListId = v
	return t
}

// SetV__DestinationPrefixListId updates property "DestinationPrefixListId".
func (t AWS_EC2_SecurityGroup_Egress) SetV__DestinationPrefixListId(v string) AWS_EC2_SecurityGroup_Egress {
	t.DestinationPrefixListId = cfz.V(v)
	return t
}

// Set__DestinationSecurityGroupId updates property "DestinationSecurityGroupId".
func (t AWS_EC2_SecurityGroup_Egress) Set__DestinationSecurityGroupId(v cfz.Expression[string]) AWS_EC2_SecurityGroup_Egress {
	t.DestinationSecurityGroupId = v
	return t
}

// SetV__DestinationSecurityGroupId updates property "DestinationSecurityGroupId".
func (t AWS_EC2_SecurityGroup_Egress) SetV__DestinationSecurityGroupId(v string) AWS_EC2_SecurityGroup_Egress {
	t.DestinationSecurityGroupId = cfz.V(v)
	return t
}

// Set__FromPort updates property "FromPort".
func (t AWS_EC2_SecurityGroup_Egress) Set__FromPort(v cfz.Expression[int32]) AWS_EC2_SecurityGroup_Egress {
	t.FromPort = v
	return t
}

// SetV__FromPort updates property "FromPort".
func (t AWS_EC2_SecurityGroup_Egress) SetV__FromPort(v int32) AWS_EC2_SecurityGroup_Egress {
	t.FromPort = cfz.V(v)
	return t
}

// Set__IpProtocol updates property "IpProtocol".
func (t AWS_EC2_SecurityGroup_Egress) Set__IpProtocol(v cfz.Expression[string]) AWS_EC2_SecurityGroup_Egress {
	t.IpProtocol = v
	return t
}

// SetV__IpProtocol updates property "IpProtocol".
func (t AWS_EC2_SecurityGroup_Egress) SetV__IpProtocol(v string) AWS_EC2_SecurityGroup_Egress {
	t.IpProtocol = cfz.V(v)
	return t
}

// Set__ToPort updates property "ToPort".
func (t AWS_EC2_SecurityGroup_Egress) Set__ToPort(v cfz.Expression[int32]) AWS_EC2_SecurityGroup_Egress {
	t.ToPort = v
	return t
}

// SetV__ToPort updates property "ToPort".
func (t AWS_EC2_SecurityGroup_Egress) SetV__ToPort(v int32) AWS_EC2_SecurityGroup_Egress {
	t.ToPort = cfz.V(v)
	return t
}
