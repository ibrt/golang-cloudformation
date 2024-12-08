// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_acmpca

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ACMPCA_CertificateAuthority_OcspConfiguration__Type is the CloudFormation type for AWS::ACMPCA::CertificateAuthority.OcspConfiguration.
	AWS_ACMPCA_CertificateAuthority_OcspConfiguration__Type = "AWS::ACMPCA::CertificateAuthority.OcspConfiguration"
)

var (
	// AWS_ACMPCA_CertificateAuthority_OcspConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ACMPCA::CertificateAuthority.OcspConfiguration.
	AWS_ACMPCA_CertificateAuthority_OcspConfiguration__PropertiesMap = struct {
		Enabled         string
		OcspCustomCname string
	}{
		Enabled:         "Enabled",
		OcspCustomCname: "OcspCustomCname",
	}

	// AWS_ACMPCA_CertificateAuthority_OcspConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ACMPCA::CertificateAuthority.OcspConfiguration.
	AWS_ACMPCA_CertificateAuthority_OcspConfiguration__PropertiesSlice = []string{
		AWS_ACMPCA_CertificateAuthority_OcspConfiguration__PropertiesMap.Enabled,
		AWS_ACMPCA_CertificateAuthority_OcspConfiguration__PropertiesMap.OcspCustomCname,
	}
)

// AWS_ACMPCA_CertificateAuthority_OcspConfiguration is a binding for AWS::ACMPCA::CertificateAuthority.OcspConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-ocspconfiguration.html
type AWS_ACMPCA_CertificateAuthority_OcspConfiguration struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-ocspconfiguration.html#cfn-acmpca-certificateauthority-ocspconfiguration-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`

	// OcspCustomCname is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-ocspconfiguration.html#cfn-acmpca-certificateauthority-ocspconfiguration-ocspcustomcname
	OcspCustomCname cfz.Expression[string] `json:"OcspCustomCname,omitempty"`
}

// New__AWS_ACMPCA_CertificateAuthority_OcspConfiguration initializes a new AWS_ACMPCA_CertificateAuthority_OcspConfiguration.
func New__AWS_ACMPCA_CertificateAuthority_OcspConfiguration() AWS_ACMPCA_CertificateAuthority_OcspConfiguration {
	return AWS_ACMPCA_CertificateAuthority_OcspConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ACMPCA_CertificateAuthority_OcspConfiguration) GetType() string {
	return AWS_ACMPCA_CertificateAuthority_OcspConfiguration__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_ACMPCA_CertificateAuthority_OcspConfiguration) Set__Enabled(v cfz.Expression[bool]) AWS_ACMPCA_CertificateAuthority_OcspConfiguration {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_ACMPCA_CertificateAuthority_OcspConfiguration) SetV__Enabled(v bool) AWS_ACMPCA_CertificateAuthority_OcspConfiguration {
	t.Enabled = cfz.V(v)
	return t
}

// Set__OcspCustomCname updates property "OcspCustomCname".
func (t AWS_ACMPCA_CertificateAuthority_OcspConfiguration) Set__OcspCustomCname(v cfz.Expression[string]) AWS_ACMPCA_CertificateAuthority_OcspConfiguration {
	t.OcspCustomCname = v
	return t
}

// SetV__OcspCustomCname updates property "OcspCustomCname".
func (t AWS_ACMPCA_CertificateAuthority_OcspConfiguration) SetV__OcspCustomCname(v string) AWS_ACMPCA_CertificateAuthority_OcspConfiguration {
	t.OcspCustomCname = cfz.V(v)
	return t
}
