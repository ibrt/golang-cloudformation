// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_IPAM_IpamOperatingRegion__Type is the CloudFormation type for AWS::EC2::IPAM.IpamOperatingRegion.
	AWS_EC2_IPAM_IpamOperatingRegion__Type = "AWS::EC2::IPAM.IpamOperatingRegion"
)

var (
	// AWS_EC2_IPAM_IpamOperatingRegion__PropertiesMap reports all the CloudFormation properties for AWS::EC2::IPAM.IpamOperatingRegion.
	AWS_EC2_IPAM_IpamOperatingRegion__PropertiesMap = struct {
		RegionName string
	}{
		RegionName: "RegionName",
	}

	// AWS_EC2_IPAM_IpamOperatingRegion__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::IPAM.IpamOperatingRegion.
	AWS_EC2_IPAM_IpamOperatingRegion__PropertiesSlice = []string{
		AWS_EC2_IPAM_IpamOperatingRegion__PropertiesMap.RegionName,
	}
)

// AWS_EC2_IPAM_IpamOperatingRegion is a binding for AWS::EC2::IPAM.IpamOperatingRegion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipam-ipamoperatingregion.html
type AWS_EC2_IPAM_IpamOperatingRegion struct {
	// RegionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ipam-ipamoperatingregion.html#cfn-ec2-ipam-ipamoperatingregion-regionname
	RegionName cfz.Expression[string] `json:"RegionName,omitempty"`
}

// New__AWS_EC2_IPAM_IpamOperatingRegion initializes a new AWS_EC2_IPAM_IpamOperatingRegion.
func New__AWS_EC2_IPAM_IpamOperatingRegion() AWS_EC2_IPAM_IpamOperatingRegion {
	return AWS_EC2_IPAM_IpamOperatingRegion{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_IPAM_IpamOperatingRegion) GetType() string {
	return AWS_EC2_IPAM_IpamOperatingRegion__Type
}

// Set__RegionName updates property "RegionName".
func (t AWS_EC2_IPAM_IpamOperatingRegion) Set__RegionName(v cfz.Expression[string]) AWS_EC2_IPAM_IpamOperatingRegion {
	t.RegionName = v
	return t
}

// SetV__RegionName updates property "RegionName".
func (t AWS_EC2_IPAM_IpamOperatingRegion) SetV__RegionName(v string) AWS_EC2_IPAM_IpamOperatingRegion {
	t.RegionName = cfz.V(v)
	return t
}
