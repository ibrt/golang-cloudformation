// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_verifiedpermissions

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__Type is the CloudFormation type for AWS::VerifiedPermissions::IdentitySource.OpenIdConnectConfiguration.
	AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__Type = "AWS::VerifiedPermissions::IdentitySource.OpenIdConnectConfiguration"
)

var (
	// AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::VerifiedPermissions::IdentitySource.OpenIdConnectConfiguration.
	AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesMap = struct {
		EntityIdPrefix     string
		GroupConfiguration string
		Issuer             string
		TokenSelection     string
	}{
		EntityIdPrefix:     "EntityIdPrefix",
		GroupConfiguration: "GroupConfiguration",
		Issuer:             "Issuer",
		TokenSelection:     "TokenSelection",
	}

	// AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::VerifiedPermissions::IdentitySource.OpenIdConnectConfiguration.
	AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesSlice = []string{
		AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesMap.EntityIdPrefix,
		AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesMap.GroupConfiguration,
		AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesMap.Issuer,
		AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__PropertiesMap.TokenSelection,
	}
)

// AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration is a binding for AWS::VerifiedPermissions::IdentitySource.OpenIdConnectConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html
type AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration struct {
	// EntityIdPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-entityidprefix
	EntityIdPrefix cfz.Expression[string] `json:"EntityIdPrefix,omitempty"`

	// GroupConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-groupconfiguration
	GroupConfiguration cfz.Expression[AWS_VerifiedPermissions_IdentitySource_OpenIdConnectGroupConfiguration] `json:"GroupConfiguration,omitempty"`

	// Issuer is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-issuer
	Issuer cfz.Expression[string] `json:"Issuer,omitempty"`

	// TokenSelection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-tokenselection
	TokenSelection cfz.Expression[AWS_VerifiedPermissions_IdentitySource_OpenIdConnectTokenSelection] `json:"TokenSelection,omitempty"`
}

// New__AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration initializes a new AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration.
func New__AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration() AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	return AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) GetType() string {
	return AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration__Type
}

// Set__EntityIdPrefix updates property "EntityIdPrefix".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) Set__EntityIdPrefix(v cfz.Expression[string]) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.EntityIdPrefix = v
	return t
}

// SetV__EntityIdPrefix updates property "EntityIdPrefix".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) SetV__EntityIdPrefix(v string) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.EntityIdPrefix = cfz.V(v)
	return t
}

// Set__GroupConfiguration updates property "GroupConfiguration".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) Set__GroupConfiguration(v cfz.Expression[AWS_VerifiedPermissions_IdentitySource_OpenIdConnectGroupConfiguration]) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.GroupConfiguration = v
	return t
}

// SetV__GroupConfiguration updates property "GroupConfiguration".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) SetV__GroupConfiguration(v AWS_VerifiedPermissions_IdentitySource_OpenIdConnectGroupConfiguration) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.GroupConfiguration = cfz.V(v)
	return t
}

// Set__Issuer updates property "Issuer".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) Set__Issuer(v cfz.Expression[string]) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.Issuer = v
	return t
}

// SetV__Issuer updates property "Issuer".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) SetV__Issuer(v string) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.Issuer = cfz.V(v)
	return t
}

// Set__TokenSelection updates property "TokenSelection".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) Set__TokenSelection(v cfz.Expression[AWS_VerifiedPermissions_IdentitySource_OpenIdConnectTokenSelection]) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.TokenSelection = v
	return t
}

// SetV__TokenSelection updates property "TokenSelection".
func (t AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration) SetV__TokenSelection(v AWS_VerifiedPermissions_IdentitySource_OpenIdConnectTokenSelection) AWS_VerifiedPermissions_IdentitySource_OpenIdConnectConfiguration {
	t.TokenSelection = cfz.V(v)
	return t
}
