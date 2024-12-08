// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__Type is the CloudFormation type for AWS::EC2::LaunchTemplate.PrivateDnsNameOptions.
	AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__Type = "AWS::EC2::LaunchTemplate.PrivateDnsNameOptions"
)

var (
	// AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesMap reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.PrivateDnsNameOptions.
	AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesMap = struct {
		EnableResourceNameDnsAAAARecord string
		EnableResourceNameDnsARecord    string
		HostnameType                    string
	}{
		EnableResourceNameDnsAAAARecord: "EnableResourceNameDnsAAAARecord",
		EnableResourceNameDnsARecord:    "EnableResourceNameDnsARecord",
		HostnameType:                    "HostnameType",
	}

	// AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::LaunchTemplate.PrivateDnsNameOptions.
	AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesSlice = []string{
		AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesMap.EnableResourceNameDnsAAAARecord,
		AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesMap.EnableResourceNameDnsARecord,
		AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__PropertiesMap.HostnameType,
	}
)

// AWS_EC2_LaunchTemplate_PrivateDnsNameOptions is a binding for AWS::EC2::LaunchTemplate.PrivateDnsNameOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-privatednsnameoptions.html
type AWS_EC2_LaunchTemplate_PrivateDnsNameOptions struct {
	// EnableResourceNameDnsAAAARecord is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-privatednsnameoptions.html#cfn-ec2-launchtemplate-privatednsnameoptions-enableresourcenamednsaaaarecord
	EnableResourceNameDnsAAAARecord cfz.Expression[bool] `json:"EnableResourceNameDnsAAAARecord,omitempty"`

	// EnableResourceNameDnsARecord is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-privatednsnameoptions.html#cfn-ec2-launchtemplate-privatednsnameoptions-enableresourcenamednsarecord
	EnableResourceNameDnsARecord cfz.Expression[bool] `json:"EnableResourceNameDnsARecord,omitempty"`

	// HostnameType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-privatednsnameoptions.html#cfn-ec2-launchtemplate-privatednsnameoptions-hostnametype
	HostnameType cfz.Expression[string] `json:"HostnameType,omitempty"`
}

// New__AWS_EC2_LaunchTemplate_PrivateDnsNameOptions initializes a new AWS_EC2_LaunchTemplate_PrivateDnsNameOptions.
func New__AWS_EC2_LaunchTemplate_PrivateDnsNameOptions() AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	return AWS_EC2_LaunchTemplate_PrivateDnsNameOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) GetType() string {
	return AWS_EC2_LaunchTemplate_PrivateDnsNameOptions__Type
}

// Set__EnableResourceNameDnsAAAARecord updates property "EnableResourceNameDnsAAAARecord".
func (t AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) Set__EnableResourceNameDnsAAAARecord(v cfz.Expression[bool]) AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	t.EnableResourceNameDnsAAAARecord = v
	return t
}

// SetV__EnableResourceNameDnsAAAARecord updates property "EnableResourceNameDnsAAAARecord".
func (t AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) SetV__EnableResourceNameDnsAAAARecord(v bool) AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	t.EnableResourceNameDnsAAAARecord = cfz.V(v)
	return t
}

// Set__EnableResourceNameDnsARecord updates property "EnableResourceNameDnsARecord".
func (t AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) Set__EnableResourceNameDnsARecord(v cfz.Expression[bool]) AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	t.EnableResourceNameDnsARecord = v
	return t
}

// SetV__EnableResourceNameDnsARecord updates property "EnableResourceNameDnsARecord".
func (t AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) SetV__EnableResourceNameDnsARecord(v bool) AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	t.EnableResourceNameDnsARecord = cfz.V(v)
	return t
}

// Set__HostnameType updates property "HostnameType".
func (t AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) Set__HostnameType(v cfz.Expression[string]) AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	t.HostnameType = v
	return t
}

// SetV__HostnameType updates property "HostnameType".
func (t AWS_EC2_LaunchTemplate_PrivateDnsNameOptions) SetV__HostnameType(v string) AWS_EC2_LaunchTemplate_PrivateDnsNameOptions {
	t.HostnameType = cfz.V(v)
	return t
}
