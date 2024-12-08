// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opensearchservice

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpenSearchService_Domain_DomainEndpointOptions__Type is the CloudFormation type for AWS::OpenSearchService::Domain.DomainEndpointOptions.
	AWS_OpenSearchService_Domain_DomainEndpointOptions__Type = "AWS::OpenSearchService::Domain.DomainEndpointOptions"
)

var (
	// AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap reports all the CloudFormation properties for AWS::OpenSearchService::Domain.DomainEndpointOptions.
	AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap = struct {
		CustomEndpoint               string
		CustomEndpointCertificateArn string
		CustomEndpointEnabled        string
		EnforceHTTPS                 string
		TLSSecurityPolicy            string
	}{
		CustomEndpoint:               "CustomEndpoint",
		CustomEndpointCertificateArn: "CustomEndpointCertificateArn",
		CustomEndpointEnabled:        "CustomEndpointEnabled",
		EnforceHTTPS:                 "EnforceHTTPS",
		TLSSecurityPolicy:            "TLSSecurityPolicy",
	}

	// AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesSlice reports all the CloudFormation properties for AWS::OpenSearchService::Domain.DomainEndpointOptions.
	AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesSlice = []string{
		AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap.CustomEndpoint,
		AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap.CustomEndpointCertificateArn,
		AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap.CustomEndpointEnabled,
		AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap.EnforceHTTPS,
		AWS_OpenSearchService_Domain_DomainEndpointOptions__PropertiesMap.TLSSecurityPolicy,
	}
)

// AWS_OpenSearchService_Domain_DomainEndpointOptions is a binding for AWS::OpenSearchService::Domain.DomainEndpointOptions.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html
type AWS_OpenSearchService_Domain_DomainEndpointOptions struct {
	// CustomEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-customendpoint
	CustomEndpoint cfz.Expression[string] `json:"CustomEndpoint,omitempty"`

	// CustomEndpointCertificateArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-customendpointcertificatearn
	CustomEndpointCertificateArn cfz.Expression[string] `json:"CustomEndpointCertificateArn,omitempty"`

	// CustomEndpointEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-customendpointenabled
	CustomEndpointEnabled cfz.Expression[bool] `json:"CustomEndpointEnabled,omitempty"`

	// EnforceHTTPS is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-enforcehttps
	EnforceHTTPS cfz.Expression[bool] `json:"EnforceHTTPS,omitempty"`

	// TLSSecurityPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html#cfn-opensearchservice-domain-domainendpointoptions-tlssecuritypolicy
	TLSSecurityPolicy cfz.Expression[string] `json:"TLSSecurityPolicy,omitempty"`
}

// New__AWS_OpenSearchService_Domain_DomainEndpointOptions initializes a new AWS_OpenSearchService_Domain_DomainEndpointOptions.
func New__AWS_OpenSearchService_Domain_DomainEndpointOptions() AWS_OpenSearchService_Domain_DomainEndpointOptions {
	return AWS_OpenSearchService_Domain_DomainEndpointOptions{}
}

// GetType returns the CloudFormation type.
func (AWS_OpenSearchService_Domain_DomainEndpointOptions) GetType() string {
	return AWS_OpenSearchService_Domain_DomainEndpointOptions__Type
}

// Set__CustomEndpoint updates property "CustomEndpoint".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) Set__CustomEndpoint(v cfz.Expression[string]) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.CustomEndpoint = v
	return t
}

// SetV__CustomEndpoint updates property "CustomEndpoint".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) SetV__CustomEndpoint(v string) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.CustomEndpoint = cfz.V(v)
	return t
}

// Set__CustomEndpointCertificateArn updates property "CustomEndpointCertificateArn".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) Set__CustomEndpointCertificateArn(v cfz.Expression[string]) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.CustomEndpointCertificateArn = v
	return t
}

// SetV__CustomEndpointCertificateArn updates property "CustomEndpointCertificateArn".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) SetV__CustomEndpointCertificateArn(v string) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.CustomEndpointCertificateArn = cfz.V(v)
	return t
}

// Set__CustomEndpointEnabled updates property "CustomEndpointEnabled".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) Set__CustomEndpointEnabled(v cfz.Expression[bool]) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.CustomEndpointEnabled = v
	return t
}

// SetV__CustomEndpointEnabled updates property "CustomEndpointEnabled".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) SetV__CustomEndpointEnabled(v bool) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.CustomEndpointEnabled = cfz.V(v)
	return t
}

// Set__EnforceHTTPS updates property "EnforceHTTPS".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) Set__EnforceHTTPS(v cfz.Expression[bool]) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.EnforceHTTPS = v
	return t
}

// SetV__EnforceHTTPS updates property "EnforceHTTPS".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) SetV__EnforceHTTPS(v bool) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.EnforceHTTPS = cfz.V(v)
	return t
}

// Set__TLSSecurityPolicy updates property "TLSSecurityPolicy".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) Set__TLSSecurityPolicy(v cfz.Expression[string]) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.TLSSecurityPolicy = v
	return t
}

// SetV__TLSSecurityPolicy updates property "TLSSecurityPolicy".
func (t AWS_OpenSearchService_Domain_DomainEndpointOptions) SetV__TLSSecurityPolicy(v string) AWS_OpenSearchService_Domain_DomainEndpointOptions {
	t.TLSSecurityPolicy = cfz.V(v)
	return t
}
