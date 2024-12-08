// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__Type is the CloudFormation type for AWS::EC2::VerifiedAccessTrustProvider.OidcOptions.
	AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__Type = "AWS::EC2::VerifiedAccessTrustProvider.OidcOptions"
)

var (
	// AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VerifiedAccessTrustProvider.OidcOptions.
	AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap = struct {
		AuthorizationEndpoint string
		ClientId              string
		ClientSecret          string
		Issuer                string
		Scope                 string
		TokenEndpoint         string
		UserInfoEndpoint      string
	}{
		AuthorizationEndpoint: "AuthorizationEndpoint",
		ClientId:              "ClientId",
		ClientSecret:          "ClientSecret",
		Issuer:                "Issuer",
		Scope:                 "Scope",
		TokenEndpoint:         "TokenEndpoint",
		UserInfoEndpoint:      "UserInfoEndpoint",
	}

	// AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VerifiedAccessTrustProvider.OidcOptions.
	AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesSlice = []string{
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.AuthorizationEndpoint,
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.ClientId,
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.ClientSecret,
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.Issuer,
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.Scope,
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.TokenEndpoint,
		AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__PropertiesMap.UserInfoEndpoint,
	}
)

// AWS_EC2_VerifiedAccessTrustProvider_OidcOptions is a binding for AWS::EC2::VerifiedAccessTrustProvider.OidcOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html
type AWS_EC2_VerifiedAccessTrustProvider_OidcOptions struct {
	// AuthorizationEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-authorizationendpoint
	AuthorizationEndpoint cfz.Expression[string] `json:"AuthorizationEndpoint,omitempty"`

	// ClientId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-clientid
	ClientId cfz.Expression[string] `json:"ClientId,omitempty"`

	// ClientSecret is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-clientsecret
	ClientSecret cfz.Expression[string] `json:"ClientSecret,omitempty"`

	// Issuer is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-issuer
	Issuer cfz.Expression[string] `json:"Issuer,omitempty"`

	// Scope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-scope
	Scope cfz.Expression[string] `json:"Scope,omitempty"`

	// TokenEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-tokenendpoint
	TokenEndpoint cfz.Expression[string] `json:"TokenEndpoint,omitempty"`

	// UserInfoEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-oidcoptions-userinfoendpoint
	UserInfoEndpoint cfz.Expression[string] `json:"UserInfoEndpoint,omitempty"`
}

// New__AWS_EC2_VerifiedAccessTrustProvider_OidcOptions initializes a new AWS_EC2_VerifiedAccessTrustProvider_OidcOptions.
func New__AWS_EC2_VerifiedAccessTrustProvider_OidcOptions() AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	return AWS_EC2_VerifiedAccessTrustProvider_OidcOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) GetType() string {
	return AWS_EC2_VerifiedAccessTrustProvider_OidcOptions__Type
}

// Set__AuthorizationEndpoint updates property "AuthorizationEndpoint".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__AuthorizationEndpoint(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.AuthorizationEndpoint = v
	return t
}

// SetV__AuthorizationEndpoint updates property "AuthorizationEndpoint".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__AuthorizationEndpoint(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.AuthorizationEndpoint = cfz.V(v)
	return t
}

// Set__ClientId updates property "ClientId".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__ClientId(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.ClientId = v
	return t
}

// SetV__ClientId updates property "ClientId".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__ClientId(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.ClientId = cfz.V(v)
	return t
}

// Set__ClientSecret updates property "ClientSecret".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__ClientSecret(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.ClientSecret = v
	return t
}

// SetV__ClientSecret updates property "ClientSecret".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__ClientSecret(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.ClientSecret = cfz.V(v)
	return t
}

// Set__Issuer updates property "Issuer".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__Issuer(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.Issuer = v
	return t
}

// SetV__Issuer updates property "Issuer".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__Issuer(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.Issuer = cfz.V(v)
	return t
}

// Set__Scope updates property "Scope".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__Scope(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.Scope = v
	return t
}

// SetV__Scope updates property "Scope".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__Scope(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.Scope = cfz.V(v)
	return t
}

// Set__TokenEndpoint updates property "TokenEndpoint".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__TokenEndpoint(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.TokenEndpoint = v
	return t
}

// SetV__TokenEndpoint updates property "TokenEndpoint".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__TokenEndpoint(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.TokenEndpoint = cfz.V(v)
	return t
}

// Set__UserInfoEndpoint updates property "UserInfoEndpoint".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) Set__UserInfoEndpoint(v cfz.Expression[string]) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.UserInfoEndpoint = v
	return t
}

// SetV__UserInfoEndpoint updates property "UserInfoEndpoint".
func (t AWS_EC2_VerifiedAccessTrustProvider_OidcOptions) SetV__UserInfoEndpoint(v string) AWS_EC2_VerifiedAccessTrustProvider_OidcOptions {
	t.UserInfoEndpoint = cfz.V(v)
	return t
}
