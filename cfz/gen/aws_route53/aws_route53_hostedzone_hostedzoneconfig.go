// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Route53_HostedZone_HostedZoneConfig__Type is the CloudFormation type for AWS::Route53::HostedZone.HostedZoneConfig.
	AWS_Route53_HostedZone_HostedZoneConfig__Type = "AWS::Route53::HostedZone.HostedZoneConfig"
)

var (
	// AWS_Route53_HostedZone_HostedZoneConfig__PropertiesMap reports all the CloudFormation properties for AWS::Route53::HostedZone.HostedZoneConfig.
	AWS_Route53_HostedZone_HostedZoneConfig__PropertiesMap = struct {
		Comment string
	}{
		Comment: "Comment",
	}

	// AWS_Route53_HostedZone_HostedZoneConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Route53::HostedZone.HostedZoneConfig.
	AWS_Route53_HostedZone_HostedZoneConfig__PropertiesSlice = []string{
		AWS_Route53_HostedZone_HostedZoneConfig__PropertiesMap.Comment,
	}
)

// AWS_Route53_HostedZone_HostedZoneConfig is a binding for AWS::Route53::HostedZone.HostedZoneConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
type AWS_Route53_HostedZone_HostedZoneConfig struct {
	// Comment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
	Comment cfz.Expression[string] `json:"Comment,omitempty"`
}

// New__AWS_Route53_HostedZone_HostedZoneConfig initializes a new AWS_Route53_HostedZone_HostedZoneConfig.
func New__AWS_Route53_HostedZone_HostedZoneConfig() AWS_Route53_HostedZone_HostedZoneConfig {
	return AWS_Route53_HostedZone_HostedZoneConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Route53_HostedZone_HostedZoneConfig) GetType() string {
	return AWS_Route53_HostedZone_HostedZoneConfig__Type
}

// Set__Comment updates property "Comment".
func (t AWS_Route53_HostedZone_HostedZoneConfig) Set__Comment(v cfz.Expression[string]) AWS_Route53_HostedZone_HostedZoneConfig {
	t.Comment = v
	return t
}

// SetV__Comment updates property "Comment".
func (t AWS_Route53_HostedZone_HostedZoneConfig) SetV__Comment(v string) AWS_Route53_HostedZone_HostedZoneConfig {
	t.Comment = cfz.V(v)
	return t
}
