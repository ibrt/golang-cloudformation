// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_acmpca

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ACMPCA_CertificateAuthority_EdiPartyName__Type is the CloudFormation type for AWS::ACMPCA::CertificateAuthority.EdiPartyName.
	AWS_ACMPCA_CertificateAuthority_EdiPartyName__Type = "AWS::ACMPCA::CertificateAuthority.EdiPartyName"
)

var (
	// AWS_ACMPCA_CertificateAuthority_EdiPartyName__PropertiesMap reports all the CloudFormation properties for AWS::ACMPCA::CertificateAuthority.EdiPartyName.
	AWS_ACMPCA_CertificateAuthority_EdiPartyName__PropertiesMap = struct {
		NameAssigner string
		PartyName    string
	}{
		NameAssigner: "NameAssigner",
		PartyName:    "PartyName",
	}

	// AWS_ACMPCA_CertificateAuthority_EdiPartyName__PropertiesSlice reports all the CloudFormation properties for AWS::ACMPCA::CertificateAuthority.EdiPartyName.
	AWS_ACMPCA_CertificateAuthority_EdiPartyName__PropertiesSlice = []string{
		AWS_ACMPCA_CertificateAuthority_EdiPartyName__PropertiesMap.NameAssigner,
		AWS_ACMPCA_CertificateAuthority_EdiPartyName__PropertiesMap.PartyName,
	}
)

// AWS_ACMPCA_CertificateAuthority_EdiPartyName is a binding for AWS::ACMPCA::CertificateAuthority.EdiPartyName.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-edipartyname.html
type AWS_ACMPCA_CertificateAuthority_EdiPartyName struct {
	// NameAssigner is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-edipartyname.html#cfn-acmpca-certificateauthority-edipartyname-nameassigner
	NameAssigner cfz.Expression[string] `json:"NameAssigner,omitempty"`

	// PartyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-edipartyname.html#cfn-acmpca-certificateauthority-edipartyname-partyname
	PartyName cfz.Expression[string] `json:"PartyName,omitempty"`
}

// New__AWS_ACMPCA_CertificateAuthority_EdiPartyName initializes a new AWS_ACMPCA_CertificateAuthority_EdiPartyName.
func New__AWS_ACMPCA_CertificateAuthority_EdiPartyName() AWS_ACMPCA_CertificateAuthority_EdiPartyName {
	return AWS_ACMPCA_CertificateAuthority_EdiPartyName{}
}

// GetType returns the CloudFormation type.
func (AWS_ACMPCA_CertificateAuthority_EdiPartyName) GetType() string {
	return AWS_ACMPCA_CertificateAuthority_EdiPartyName__Type
}

// Set__NameAssigner updates property "NameAssigner".
func (t AWS_ACMPCA_CertificateAuthority_EdiPartyName) Set__NameAssigner(v cfz.Expression[string]) AWS_ACMPCA_CertificateAuthority_EdiPartyName {
	t.NameAssigner = v
	return t
}

// SetV__NameAssigner updates property "NameAssigner".
func (t AWS_ACMPCA_CertificateAuthority_EdiPartyName) SetV__NameAssigner(v string) AWS_ACMPCA_CertificateAuthority_EdiPartyName {
	t.NameAssigner = cfz.V(v)
	return t
}

// Set__PartyName updates property "PartyName".
func (t AWS_ACMPCA_CertificateAuthority_EdiPartyName) Set__PartyName(v cfz.Expression[string]) AWS_ACMPCA_CertificateAuthority_EdiPartyName {
	t.PartyName = v
	return t
}

// SetV__PartyName updates property "PartyName".
func (t AWS_ACMPCA_CertificateAuthority_EdiPartyName) SetV__PartyName(v string) AWS_ACMPCA_CertificateAuthority_EdiPartyName {
	t.PartyName = cfz.V(v)
	return t
}
