// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pcaconnectorad

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_PCAConnectorAD_Template_TemplateV3__Type is the CloudFormation type for AWS::PCAConnectorAD::Template.TemplateV3.
	AWS_PCAConnectorAD_Template_TemplateV3__Type = "AWS::PCAConnectorAD::Template.TemplateV3"
)

var (
	// AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap reports all the CloudFormation properties for AWS::PCAConnectorAD::Template.TemplateV3.
	AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap = struct {
		CertificateValidity  string
		EnrollmentFlags      string
		Extensions           string
		GeneralFlags         string
		HashAlgorithm        string
		PrivateKeyAttributes string
		PrivateKeyFlags      string
		SubjectNameFlags     string
		SupersededTemplates  string
	}{
		CertificateValidity:  "CertificateValidity",
		EnrollmentFlags:      "EnrollmentFlags",
		Extensions:           "Extensions",
		GeneralFlags:         "GeneralFlags",
		HashAlgorithm:        "HashAlgorithm",
		PrivateKeyAttributes: "PrivateKeyAttributes",
		PrivateKeyFlags:      "PrivateKeyFlags",
		SubjectNameFlags:     "SubjectNameFlags",
		SupersededTemplates:  "SupersededTemplates",
	}

	// AWS_PCAConnectorAD_Template_TemplateV3__PropertiesSlice reports all the CloudFormation properties for AWS::PCAConnectorAD::Template.TemplateV3.
	AWS_PCAConnectorAD_Template_TemplateV3__PropertiesSlice = []string{
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.CertificateValidity,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.EnrollmentFlags,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.Extensions,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.GeneralFlags,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.HashAlgorithm,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.PrivateKeyAttributes,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.PrivateKeyFlags,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.SubjectNameFlags,
		AWS_PCAConnectorAD_Template_TemplateV3__PropertiesMap.SupersededTemplates,
	}
)

// AWS_PCAConnectorAD_Template_TemplateV3 is a binding for AWS::PCAConnectorAD::Template.TemplateV3.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html
type AWS_PCAConnectorAD_Template_TemplateV3 struct {
	// CertificateValidity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-certificatevalidity
	CertificateValidity cfz.Expression[AWS_PCAConnectorAD_Template_CertificateValidity] `json:"CertificateValidity,omitempty"`

	// EnrollmentFlags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-enrollmentflags
	EnrollmentFlags cfz.Expression[AWS_PCAConnectorAD_Template_EnrollmentFlagsV3] `json:"EnrollmentFlags,omitempty"`

	// Extensions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-extensions
	Extensions cfz.Expression[AWS_PCAConnectorAD_Template_ExtensionsV3] `json:"Extensions,omitempty"`

	// GeneralFlags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-generalflags
	GeneralFlags cfz.Expression[AWS_PCAConnectorAD_Template_GeneralFlagsV3] `json:"GeneralFlags,omitempty"`

	// HashAlgorithm is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-hashalgorithm
	HashAlgorithm cfz.Expression[string] `json:"HashAlgorithm,omitempty"`

	// PrivateKeyAttributes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-privatekeyattributes
	PrivateKeyAttributes cfz.Expression[AWS_PCAConnectorAD_Template_PrivateKeyAttributesV3] `json:"PrivateKeyAttributes,omitempty"`

	// PrivateKeyFlags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-privatekeyflags
	PrivateKeyFlags cfz.Expression[AWS_PCAConnectorAD_Template_PrivateKeyFlagsV3] `json:"PrivateKeyFlags,omitempty"`

	// SubjectNameFlags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-subjectnameflags
	SubjectNameFlags cfz.Expression[AWS_PCAConnectorAD_Template_SubjectNameFlagsV3] `json:"SubjectNameFlags,omitempty"`

	// SupersededTemplates is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-templatev3.html#cfn-pcaconnectorad-template-templatev3-supersededtemplates
	SupersededTemplates cfz.ExpressionSlice[string] `json:"SupersededTemplates,omitempty"`
}

// New__AWS_PCAConnectorAD_Template_TemplateV3 initializes a new AWS_PCAConnectorAD_Template_TemplateV3.
func New__AWS_PCAConnectorAD_Template_TemplateV3() AWS_PCAConnectorAD_Template_TemplateV3 {
	return AWS_PCAConnectorAD_Template_TemplateV3{}
}

// GetType returns the CloudFormation type.
func (AWS_PCAConnectorAD_Template_TemplateV3) GetType() string {
	return AWS_PCAConnectorAD_Template_TemplateV3__Type
}

// Set__CertificateValidity updates property "CertificateValidity".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__CertificateValidity(v cfz.Expression[AWS_PCAConnectorAD_Template_CertificateValidity]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.CertificateValidity = v
	return t
}

// SetV__CertificateValidity updates property "CertificateValidity".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__CertificateValidity(v AWS_PCAConnectorAD_Template_CertificateValidity) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.CertificateValidity = cfz.V(v)
	return t
}

// Set__EnrollmentFlags updates property "EnrollmentFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__EnrollmentFlags(v cfz.Expression[AWS_PCAConnectorAD_Template_EnrollmentFlagsV3]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.EnrollmentFlags = v
	return t
}

// SetV__EnrollmentFlags updates property "EnrollmentFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__EnrollmentFlags(v AWS_PCAConnectorAD_Template_EnrollmentFlagsV3) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.EnrollmentFlags = cfz.V(v)
	return t
}

// Set__Extensions updates property "Extensions".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__Extensions(v cfz.Expression[AWS_PCAConnectorAD_Template_ExtensionsV3]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.Extensions = v
	return t
}

// SetV__Extensions updates property "Extensions".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__Extensions(v AWS_PCAConnectorAD_Template_ExtensionsV3) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.Extensions = cfz.V(v)
	return t
}

// Set__GeneralFlags updates property "GeneralFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__GeneralFlags(v cfz.Expression[AWS_PCAConnectorAD_Template_GeneralFlagsV3]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.GeneralFlags = v
	return t
}

// SetV__GeneralFlags updates property "GeneralFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__GeneralFlags(v AWS_PCAConnectorAD_Template_GeneralFlagsV3) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.GeneralFlags = cfz.V(v)
	return t
}

// Set__HashAlgorithm updates property "HashAlgorithm".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__HashAlgorithm(v cfz.Expression[string]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.HashAlgorithm = v
	return t
}

// SetV__HashAlgorithm updates property "HashAlgorithm".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__HashAlgorithm(v string) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.HashAlgorithm = cfz.V(v)
	return t
}

// Set__PrivateKeyAttributes updates property "PrivateKeyAttributes".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__PrivateKeyAttributes(v cfz.Expression[AWS_PCAConnectorAD_Template_PrivateKeyAttributesV3]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.PrivateKeyAttributes = v
	return t
}

// SetV__PrivateKeyAttributes updates property "PrivateKeyAttributes".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__PrivateKeyAttributes(v AWS_PCAConnectorAD_Template_PrivateKeyAttributesV3) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.PrivateKeyAttributes = cfz.V(v)
	return t
}

// Set__PrivateKeyFlags updates property "PrivateKeyFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__PrivateKeyFlags(v cfz.Expression[AWS_PCAConnectorAD_Template_PrivateKeyFlagsV3]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.PrivateKeyFlags = v
	return t
}

// SetV__PrivateKeyFlags updates property "PrivateKeyFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__PrivateKeyFlags(v AWS_PCAConnectorAD_Template_PrivateKeyFlagsV3) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.PrivateKeyFlags = cfz.V(v)
	return t
}

// Set__SubjectNameFlags updates property "SubjectNameFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__SubjectNameFlags(v cfz.Expression[AWS_PCAConnectorAD_Template_SubjectNameFlagsV3]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.SubjectNameFlags = v
	return t
}

// SetV__SubjectNameFlags updates property "SubjectNameFlags".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetV__SubjectNameFlags(v AWS_PCAConnectorAD_Template_SubjectNameFlagsV3) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.SubjectNameFlags = cfz.V(v)
	return t
}

// Set__SupersededTemplates updates property "SupersededTemplates".
func (t AWS_PCAConnectorAD_Template_TemplateV3) Set__SupersededTemplates(v cfz.ExpressionSlice[string]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.SupersededTemplates = v
	return t
}

// SetS__SupersededTemplates updates property "SupersededTemplates".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetS__SupersededTemplates(v ...cfz.Expression[string]) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.SupersededTemplates = cfz.S(v...)
	return t
}

// SetSV__SupersededTemplates updates property "SupersededTemplates".
func (t AWS_PCAConnectorAD_Template_TemplateV3) SetSV__SupersededTemplates(v ...string) AWS_PCAConnectorAD_Template_TemplateV3 {
	t.SupersededTemplates = cfz.SV(v...)
	return t
}
