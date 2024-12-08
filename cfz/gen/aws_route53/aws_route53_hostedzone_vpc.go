// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Route53_HostedZone_VPC__Type is the CloudFormation type for AWS::Route53::HostedZone.VPC.
	AWS_Route53_HostedZone_VPC__Type = "AWS::Route53::HostedZone.VPC"
)

var (
	// AWS_Route53_HostedZone_VPC__PropertiesMap reports all the CloudFormation properties for AWS::Route53::HostedZone.VPC.
	AWS_Route53_HostedZone_VPC__PropertiesMap = struct {
		VPCId     string
		VPCRegion string
	}{
		VPCId:     "VPCId",
		VPCRegion: "VPCRegion",
	}

	// AWS_Route53_HostedZone_VPC__PropertiesSlice reports all the CloudFormation properties for AWS::Route53::HostedZone.VPC.
	AWS_Route53_HostedZone_VPC__PropertiesSlice = []string{
		AWS_Route53_HostedZone_VPC__PropertiesMap.VPCId,
		AWS_Route53_HostedZone_VPC__PropertiesMap.VPCRegion,
	}
)

// AWS_Route53_HostedZone_VPC is a binding for AWS::Route53::HostedZone.VPC.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html
type AWS_Route53_HostedZone_VPC struct {
	// VPCId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcid
	VPCId cfz.Expression[string] `json:"VPCId,omitempty"`

	// VPCRegion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcregion
	VPCRegion cfz.Expression[string] `json:"VPCRegion,omitempty"`
}

// New__AWS_Route53_HostedZone_VPC initializes a new AWS_Route53_HostedZone_VPC.
func New__AWS_Route53_HostedZone_VPC() AWS_Route53_HostedZone_VPC {
	return AWS_Route53_HostedZone_VPC{}
}

// GetType returns the CloudFormation type.
func (AWS_Route53_HostedZone_VPC) GetType() string {
	return AWS_Route53_HostedZone_VPC__Type
}

// Set__VPCId updates property "VPCId".
func (t AWS_Route53_HostedZone_VPC) Set__VPCId(v cfz.Expression[string]) AWS_Route53_HostedZone_VPC {
	t.VPCId = v
	return t
}

// SetV__VPCId updates property "VPCId".
func (t AWS_Route53_HostedZone_VPC) SetV__VPCId(v string) AWS_Route53_HostedZone_VPC {
	t.VPCId = cfz.V(v)
	return t
}

// Set__VPCRegion updates property "VPCRegion".
func (t AWS_Route53_HostedZone_VPC) Set__VPCRegion(v cfz.Expression[string]) AWS_Route53_HostedZone_VPC {
	t.VPCRegion = v
	return t
}

// SetV__VPCRegion updates property "VPCRegion".
func (t AWS_Route53_HostedZone_VPC) SetV__VPCRegion(v string) AWS_Route53_HostedZone_VPC {
	t.VPCRegion = cfz.V(v)
	return t
}
