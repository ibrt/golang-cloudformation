// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__Type is the CloudFormation type for AWS::CloudFront::OriginAccessControl.OriginAccessControlConfig.
	AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__Type = "AWS::CloudFront::OriginAccessControl.OriginAccessControlConfig"
)

var (
	// AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::OriginAccessControl.OriginAccessControlConfig.
	AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap = struct {
		Description                   string
		Name                          string
		OriginAccessControlOriginType string
		SigningBehavior               string
		SigningProtocol               string
	}{
		Description:                   "Description",
		Name:                          "Name",
		OriginAccessControlOriginType: "OriginAccessControlOriginType",
		SigningBehavior:               "SigningBehavior",
		SigningProtocol:               "SigningProtocol",
	}

	// AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::OriginAccessControl.OriginAccessControlConfig.
	AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesSlice = []string{
		AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap.Description,
		AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap.Name,
		AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap.OriginAccessControlOriginType,
		AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap.SigningBehavior,
		AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__PropertiesMap.SigningProtocol,
	}
)

// AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig is a binding for AWS::CloudFront::OriginAccessControl.OriginAccessControlConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html
type AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig struct {
	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// OriginAccessControlOriginType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-originaccesscontrolorigintype
	OriginAccessControlOriginType cfz.Expression[string] `json:"OriginAccessControlOriginType,omitempty"`

	// SigningBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-signingbehavior
	SigningBehavior cfz.Expression[string] `json:"SigningBehavior,omitempty"`

	// SigningProtocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-signingprotocol
	SigningProtocol cfz.Expression[string] `json:"SigningProtocol,omitempty"`
}

// New__AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig initializes a new AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig.
func New__AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig() AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	return AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) GetType() string {
	return AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig__Type
}

// Set__Description updates property "Description".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) Set__Description(v cfz.Expression[string]) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) SetV__Description(v string) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) Set__Name(v cfz.Expression[string]) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) SetV__Name(v string) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.Name = cfz.V(v)
	return t
}

// Set__OriginAccessControlOriginType updates property "OriginAccessControlOriginType".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) Set__OriginAccessControlOriginType(v cfz.Expression[string]) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.OriginAccessControlOriginType = v
	return t
}

// SetV__OriginAccessControlOriginType updates property "OriginAccessControlOriginType".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) SetV__OriginAccessControlOriginType(v string) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.OriginAccessControlOriginType = cfz.V(v)
	return t
}

// Set__SigningBehavior updates property "SigningBehavior".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) Set__SigningBehavior(v cfz.Expression[string]) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.SigningBehavior = v
	return t
}

// SetV__SigningBehavior updates property "SigningBehavior".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) SetV__SigningBehavior(v string) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.SigningBehavior = cfz.V(v)
	return t
}

// Set__SigningProtocol updates property "SigningProtocol".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) Set__SigningProtocol(v cfz.Expression[string]) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.SigningProtocol = v
	return t
}

// SetV__SigningProtocol updates property "SigningProtocol".
func (t AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig) SetV__SigningProtocol(v string) AWS_CloudFront_OriginAccessControl_OriginAccessControlConfig {
	t.SigningProtocol = cfz.V(v)
	return t
}
