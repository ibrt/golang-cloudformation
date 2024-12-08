// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_acmpca

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ACMPCA_Certificate_KeyUsage__Type is the CloudFormation type for AWS::ACMPCA::Certificate.KeyUsage.
	AWS_ACMPCA_Certificate_KeyUsage__Type = "AWS::ACMPCA::Certificate.KeyUsage"
)

var (
	// AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap reports all the CloudFormation properties for AWS::ACMPCA::Certificate.KeyUsage.
	AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap = struct {
		CRLSign          string
		DataEncipherment string
		DecipherOnly     string
		DigitalSignature string
		EncipherOnly     string
		KeyAgreement     string
		KeyCertSign      string
		KeyEncipherment  string
		NonRepudiation   string
	}{
		CRLSign:          "CRLSign",
		DataEncipherment: "DataEncipherment",
		DecipherOnly:     "DecipherOnly",
		DigitalSignature: "DigitalSignature",
		EncipherOnly:     "EncipherOnly",
		KeyAgreement:     "KeyAgreement",
		KeyCertSign:      "KeyCertSign",
		KeyEncipherment:  "KeyEncipherment",
		NonRepudiation:   "NonRepudiation",
	}

	// AWS_ACMPCA_Certificate_KeyUsage__PropertiesSlice reports all the CloudFormation properties for AWS::ACMPCA::Certificate.KeyUsage.
	AWS_ACMPCA_Certificate_KeyUsage__PropertiesSlice = []string{
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.CRLSign,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.DataEncipherment,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.DecipherOnly,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.DigitalSignature,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.EncipherOnly,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.KeyAgreement,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.KeyCertSign,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.KeyEncipherment,
		AWS_ACMPCA_Certificate_KeyUsage__PropertiesMap.NonRepudiation,
	}
)

// AWS_ACMPCA_Certificate_KeyUsage is a binding for AWS::ACMPCA::Certificate.KeyUsage.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html
type AWS_ACMPCA_Certificate_KeyUsage struct {
	// CRLSign is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-crlsign
	CRLSign cfz.Expression[bool] `json:"CRLSign,omitempty"`

	// DataEncipherment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-dataencipherment
	DataEncipherment cfz.Expression[bool] `json:"DataEncipherment,omitempty"`

	// DecipherOnly is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-decipheronly
	DecipherOnly cfz.Expression[bool] `json:"DecipherOnly,omitempty"`

	// DigitalSignature is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-digitalsignature
	DigitalSignature cfz.Expression[bool] `json:"DigitalSignature,omitempty"`

	// EncipherOnly is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-encipheronly
	EncipherOnly cfz.Expression[bool] `json:"EncipherOnly,omitempty"`

	// KeyAgreement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-keyagreement
	KeyAgreement cfz.Expression[bool] `json:"KeyAgreement,omitempty"`

	// KeyCertSign is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-keycertsign
	KeyCertSign cfz.Expression[bool] `json:"KeyCertSign,omitempty"`

	// KeyEncipherment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-keyencipherment
	KeyEncipherment cfz.Expression[bool] `json:"KeyEncipherment,omitempty"`

	// NonRepudiation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-nonrepudiation
	NonRepudiation cfz.Expression[bool] `json:"NonRepudiation,omitempty"`
}

// New__AWS_ACMPCA_Certificate_KeyUsage initializes a new AWS_ACMPCA_Certificate_KeyUsage.
func New__AWS_ACMPCA_Certificate_KeyUsage() AWS_ACMPCA_Certificate_KeyUsage {
	return AWS_ACMPCA_Certificate_KeyUsage{}
}

// GetType returns the CloudFormation type.
func (AWS_ACMPCA_Certificate_KeyUsage) GetType() string {
	return AWS_ACMPCA_Certificate_KeyUsage__Type
}

// Set__CRLSign updates property "CRLSign".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__CRLSign(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.CRLSign = v
	return t
}

// SetV__CRLSign updates property "CRLSign".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__CRLSign(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.CRLSign = cfz.V(v)
	return t
}

// Set__DataEncipherment updates property "DataEncipherment".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__DataEncipherment(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.DataEncipherment = v
	return t
}

// SetV__DataEncipherment updates property "DataEncipherment".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__DataEncipherment(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.DataEncipherment = cfz.V(v)
	return t
}

// Set__DecipherOnly updates property "DecipherOnly".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__DecipherOnly(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.DecipherOnly = v
	return t
}

// SetV__DecipherOnly updates property "DecipherOnly".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__DecipherOnly(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.DecipherOnly = cfz.V(v)
	return t
}

// Set__DigitalSignature updates property "DigitalSignature".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__DigitalSignature(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.DigitalSignature = v
	return t
}

// SetV__DigitalSignature updates property "DigitalSignature".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__DigitalSignature(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.DigitalSignature = cfz.V(v)
	return t
}

// Set__EncipherOnly updates property "EncipherOnly".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__EncipherOnly(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.EncipherOnly = v
	return t
}

// SetV__EncipherOnly updates property "EncipherOnly".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__EncipherOnly(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.EncipherOnly = cfz.V(v)
	return t
}

// Set__KeyAgreement updates property "KeyAgreement".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__KeyAgreement(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.KeyAgreement = v
	return t
}

// SetV__KeyAgreement updates property "KeyAgreement".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__KeyAgreement(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.KeyAgreement = cfz.V(v)
	return t
}

// Set__KeyCertSign updates property "KeyCertSign".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__KeyCertSign(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.KeyCertSign = v
	return t
}

// SetV__KeyCertSign updates property "KeyCertSign".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__KeyCertSign(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.KeyCertSign = cfz.V(v)
	return t
}

// Set__KeyEncipherment updates property "KeyEncipherment".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__KeyEncipherment(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.KeyEncipherment = v
	return t
}

// SetV__KeyEncipherment updates property "KeyEncipherment".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__KeyEncipherment(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.KeyEncipherment = cfz.V(v)
	return t
}

// Set__NonRepudiation updates property "NonRepudiation".
func (t AWS_ACMPCA_Certificate_KeyUsage) Set__NonRepudiation(v cfz.Expression[bool]) AWS_ACMPCA_Certificate_KeyUsage {
	t.NonRepudiation = v
	return t
}

// SetV__NonRepudiation updates property "NonRepudiation".
func (t AWS_ACMPCA_Certificate_KeyUsage) SetV__NonRepudiation(v bool) AWS_ACMPCA_Certificate_KeyUsage {
	t.NonRepudiation = cfz.V(v)
	return t
}
