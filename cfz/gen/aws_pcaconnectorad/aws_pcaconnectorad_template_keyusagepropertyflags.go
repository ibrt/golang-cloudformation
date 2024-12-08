// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pcaconnectorad

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__Type is the CloudFormation type for AWS::PCAConnectorAD::Template.KeyUsagePropertyFlags.
	AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__Type = "AWS::PCAConnectorAD::Template.KeyUsagePropertyFlags"
)

var (
	// AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesMap reports all the CloudFormation properties for AWS::PCAConnectorAD::Template.KeyUsagePropertyFlags.
	AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesMap = struct {
		Decrypt      string
		KeyAgreement string
		Sign         string
	}{
		Decrypt:      "Decrypt",
		KeyAgreement: "KeyAgreement",
		Sign:         "Sign",
	}

	// AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesSlice reports all the CloudFormation properties for AWS::PCAConnectorAD::Template.KeyUsagePropertyFlags.
	AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesSlice = []string{
		AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesMap.Decrypt,
		AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesMap.KeyAgreement,
		AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__PropertiesMap.Sign,
	}
)

// AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags is a binding for AWS::PCAConnectorAD::Template.KeyUsagePropertyFlags.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html
type AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags struct {
	// Decrypt is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html#cfn-pcaconnectorad-template-keyusagepropertyflags-decrypt
	Decrypt cfz.Expression[bool] `json:"Decrypt,omitempty"`

	// KeyAgreement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html#cfn-pcaconnectorad-template-keyusagepropertyflags-keyagreement
	KeyAgreement cfz.Expression[bool] `json:"KeyAgreement,omitempty"`

	// Sign is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html#cfn-pcaconnectorad-template-keyusagepropertyflags-sign
	Sign cfz.Expression[bool] `json:"Sign,omitempty"`
}

// New__AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags initializes a new AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags.
func New__AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags() AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	return AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags{}
}

// GetType returns the CloudFormation type.
func (AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) GetType() string {
	return AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags__Type
}

// Set__Decrypt updates property "Decrypt".
func (t AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) Set__Decrypt(v cfz.Expression[bool]) AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	t.Decrypt = v
	return t
}

// SetV__Decrypt updates property "Decrypt".
func (t AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) SetV__Decrypt(v bool) AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	t.Decrypt = cfz.V(v)
	return t
}

// Set__KeyAgreement updates property "KeyAgreement".
func (t AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) Set__KeyAgreement(v cfz.Expression[bool]) AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	t.KeyAgreement = v
	return t
}

// SetV__KeyAgreement updates property "KeyAgreement".
func (t AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) SetV__KeyAgreement(v bool) AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	t.KeyAgreement = cfz.V(v)
	return t
}

// Set__Sign updates property "Sign".
func (t AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) Set__Sign(v cfz.Expression[bool]) AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	t.Sign = v
	return t
}

// SetV__Sign updates property "Sign".
func (t AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags) SetV__Sign(v bool) AWS_PCAConnectorAD_Template_KeyUsagePropertyFlags {
	t.Sign = cfz.V(v)
	return t
}
