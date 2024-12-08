// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pcaconnectorscep

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__Type is the CloudFormation type for AWS::PCAConnectorSCEP::Connector.OpenIdConfiguration.
	AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__Type = "AWS::PCAConnectorSCEP::Connector.OpenIdConfiguration"
)

var (
	// AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::PCAConnectorSCEP::Connector.OpenIdConfiguration.
	AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesMap = struct {
		Audience string
		Issuer   string
		Subject  string
	}{
		Audience: "Audience",
		Issuer:   "Issuer",
		Subject:  "Subject",
	}

	// AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::PCAConnectorSCEP::Connector.OpenIdConfiguration.
	AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesSlice = []string{
		AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesMap.Audience,
		AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesMap.Issuer,
		AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__PropertiesMap.Subject,
	}
)

// AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration is a binding for AWS::PCAConnectorSCEP::Connector.OpenIdConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html
type AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration struct {
	// Audience is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-audience
	Audience cfz.Expression[string] `json:"Audience,omitempty"`

	// Issuer is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-issuer
	Issuer cfz.Expression[string] `json:"Issuer,omitempty"`

	// Subject is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorscep-connector-openidconfiguration.html#cfn-pcaconnectorscep-connector-openidconfiguration-subject
	Subject cfz.Expression[string] `json:"Subject,omitempty"`
}

// New__AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration initializes a new AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration.
func New__AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration() AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	return AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) GetType() string {
	return AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration__Type
}

// Set__Audience updates property "Audience".
func (t AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) Set__Audience(v cfz.Expression[string]) AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	t.Audience = v
	return t
}

// SetV__Audience updates property "Audience".
func (t AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) SetV__Audience(v string) AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	t.Audience = cfz.V(v)
	return t
}

// Set__Issuer updates property "Issuer".
func (t AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) Set__Issuer(v cfz.Expression[string]) AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	t.Issuer = v
	return t
}

// SetV__Issuer updates property "Issuer".
func (t AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) SetV__Issuer(v string) AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	t.Issuer = cfz.V(v)
	return t
}

// Set__Subject updates property "Subject".
func (t AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) Set__Subject(v cfz.Expression[string]) AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	t.Subject = v
	return t
}

// SetV__Subject updates property "Subject".
func (t AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration) SetV__Subject(v string) AWS_PCAConnectorSCEP_Connector_OpenIdConfiguration {
	t.Subject = cfz.V(v)
	return t
}
