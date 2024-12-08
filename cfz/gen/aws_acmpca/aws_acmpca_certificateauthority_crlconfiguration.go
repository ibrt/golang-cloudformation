// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_acmpca

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ACMPCA_CertificateAuthority_CrlConfiguration__Type is the CloudFormation type for AWS::ACMPCA::CertificateAuthority.CrlConfiguration.
	AWS_ACMPCA_CertificateAuthority_CrlConfiguration__Type = "AWS::ACMPCA::CertificateAuthority.CrlConfiguration"
)

var (
	// AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::ACMPCA::CertificateAuthority.CrlConfiguration.
	AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap = struct {
		CrlDistributionPointExtensionConfiguration string
		CustomCname                                string
		Enabled                                    string
		ExpirationInDays                           string
		S3BucketName                               string
		S3ObjectAcl                                string
	}{
		CrlDistributionPointExtensionConfiguration: "CrlDistributionPointExtensionConfiguration",
		CustomCname:      "CustomCname",
		Enabled:          "Enabled",
		ExpirationInDays: "ExpirationInDays",
		S3BucketName:     "S3BucketName",
		S3ObjectAcl:      "S3ObjectAcl",
	}

	// AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::ACMPCA::CertificateAuthority.CrlConfiguration.
	AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesSlice = []string{
		AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap.CrlDistributionPointExtensionConfiguration,
		AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap.CustomCname,
		AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap.Enabled,
		AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap.ExpirationInDays,
		AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap.S3BucketName,
		AWS_ACMPCA_CertificateAuthority_CrlConfiguration__PropertiesMap.S3ObjectAcl,
	}
)

// AWS_ACMPCA_CertificateAuthority_CrlConfiguration is a binding for AWS::ACMPCA::CertificateAuthority.CrlConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html
type AWS_ACMPCA_CertificateAuthority_CrlConfiguration struct {
	// CrlDistributionPointExtensionConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-crldistributionpointextensionconfiguration
	CrlDistributionPointExtensionConfiguration cfz.Expression[AWS_ACMPCA_CertificateAuthority_CrlDistributionPointExtensionConfiguration] `json:"CrlDistributionPointExtensionConfiguration,omitempty"`

	// CustomCname is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-customcname
	CustomCname cfz.Expression[string] `json:"CustomCname,omitempty"`

	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`

	// ExpirationInDays is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-expirationindays
	ExpirationInDays cfz.Expression[int32] `json:"ExpirationInDays,omitempty"`

	// S3BucketName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-s3bucketname
	S3BucketName cfz.Expression[string] `json:"S3BucketName,omitempty"`

	// S3ObjectAcl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-crlconfiguration.html#cfn-acmpca-certificateauthority-crlconfiguration-s3objectacl
	S3ObjectAcl cfz.Expression[string] `json:"S3ObjectAcl,omitempty"`
}

// New__AWS_ACMPCA_CertificateAuthority_CrlConfiguration initializes a new AWS_ACMPCA_CertificateAuthority_CrlConfiguration.
func New__AWS_ACMPCA_CertificateAuthority_CrlConfiguration() AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	return AWS_ACMPCA_CertificateAuthority_CrlConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_ACMPCA_CertificateAuthority_CrlConfiguration) GetType() string {
	return AWS_ACMPCA_CertificateAuthority_CrlConfiguration__Type
}

// Set__CrlDistributionPointExtensionConfiguration updates property "CrlDistributionPointExtensionConfiguration".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) Set__CrlDistributionPointExtensionConfiguration(v cfz.Expression[AWS_ACMPCA_CertificateAuthority_CrlDistributionPointExtensionConfiguration]) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.CrlDistributionPointExtensionConfiguration = v
	return t
}

// SetV__CrlDistributionPointExtensionConfiguration updates property "CrlDistributionPointExtensionConfiguration".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) SetV__CrlDistributionPointExtensionConfiguration(v AWS_ACMPCA_CertificateAuthority_CrlDistributionPointExtensionConfiguration) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.CrlDistributionPointExtensionConfiguration = cfz.V(v)
	return t
}

// Set__CustomCname updates property "CustomCname".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) Set__CustomCname(v cfz.Expression[string]) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.CustomCname = v
	return t
}

// SetV__CustomCname updates property "CustomCname".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) SetV__CustomCname(v string) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.CustomCname = cfz.V(v)
	return t
}

// Set__Enabled updates property "Enabled".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) Set__Enabled(v cfz.Expression[bool]) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) SetV__Enabled(v bool) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.Enabled = cfz.V(v)
	return t
}

// Set__ExpirationInDays updates property "ExpirationInDays".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) Set__ExpirationInDays(v cfz.Expression[int32]) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.ExpirationInDays = v
	return t
}

// SetV__ExpirationInDays updates property "ExpirationInDays".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) SetV__ExpirationInDays(v int32) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.ExpirationInDays = cfz.V(v)
	return t
}

// Set__S3BucketName updates property "S3BucketName".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) Set__S3BucketName(v cfz.Expression[string]) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.S3BucketName = v
	return t
}

// SetV__S3BucketName updates property "S3BucketName".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) SetV__S3BucketName(v string) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.S3BucketName = cfz.V(v)
	return t
}

// Set__S3ObjectAcl updates property "S3ObjectAcl".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) Set__S3ObjectAcl(v cfz.Expression[string]) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.S3ObjectAcl = v
	return t
}

// SetV__S3ObjectAcl updates property "S3ObjectAcl".
func (t AWS_ACMPCA_CertificateAuthority_CrlConfiguration) SetV__S3ObjectAcl(v string) AWS_ACMPCA_CertificateAuthority_CrlConfiguration {
	t.S3ObjectAcl = cfz.V(v)
	return t
}
