// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_RequestInspectionACFP__Type is the CloudFormation type for AWS::WAFv2::WebACL.RequestInspectionACFP.
	AWS_WAFv2_WebACL_RequestInspectionACFP__Type = "AWS::WAFv2::WebACL.RequestInspectionACFP"
)

var (
	// AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.RequestInspectionACFP.
	AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap = struct {
		AddressFields     string
		EmailField        string
		PasswordField     string
		PayloadType       string
		PhoneNumberFields string
		UsernameField     string
	}{
		AddressFields:     "AddressFields",
		EmailField:        "EmailField",
		PasswordField:     "PasswordField",
		PayloadType:       "PayloadType",
		PhoneNumberFields: "PhoneNumberFields",
		UsernameField:     "UsernameField",
	}

	// AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.RequestInspectionACFP.
	AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap.AddressFields,
		AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap.EmailField,
		AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap.PasswordField,
		AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap.PayloadType,
		AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap.PhoneNumberFields,
		AWS_WAFv2_WebACL_RequestInspectionACFP__PropertiesMap.UsernameField,
	}
)

// AWS_WAFv2_WebACL_RequestInspectionACFP is a binding for AWS::WAFv2::WebACL.RequestInspectionACFP.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html
type AWS_WAFv2_WebACL_RequestInspectionACFP struct {
	// AddressFields is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-addressfields
	AddressFields cfz.ExpressionSlice[AWS_WAFv2_WebACL_FieldIdentifier] `json:"AddressFields,omitempty"`

	// EmailField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-emailfield
	EmailField cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier] `json:"EmailField,omitempty"`

	// PasswordField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-passwordfield
	PasswordField cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier] `json:"PasswordField,omitempty"`

	// PayloadType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-payloadtype
	PayloadType cfz.Expression[string] `json:"PayloadType,omitempty"`

	// PhoneNumberFields is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-phonenumberfields
	PhoneNumberFields cfz.ExpressionSlice[AWS_WAFv2_WebACL_FieldIdentifier] `json:"PhoneNumberFields,omitempty"`

	// UsernameField is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-requestinspectionacfp.html#cfn-wafv2-webacl-requestinspectionacfp-usernamefield
	UsernameField cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier] `json:"UsernameField,omitempty"`
}

// New__AWS_WAFv2_WebACL_RequestInspectionACFP initializes a new AWS_WAFv2_WebACL_RequestInspectionACFP.
func New__AWS_WAFv2_WebACL_RequestInspectionACFP() AWS_WAFv2_WebACL_RequestInspectionACFP {
	return AWS_WAFv2_WebACL_RequestInspectionACFP{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_RequestInspectionACFP) GetType() string {
	return AWS_WAFv2_WebACL_RequestInspectionACFP__Type
}

// Set__AddressFields updates property "AddressFields".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) Set__AddressFields(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.AddressFields = v
	return t
}

// SetS__AddressFields updates property "AddressFields".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetS__AddressFields(v ...cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.AddressFields = cfz.S(v...)
	return t
}

// SetSV__AddressFields updates property "AddressFields".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetSV__AddressFields(v ...AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.AddressFields = cfz.SV(v...)
	return t
}

// Set__EmailField updates property "EmailField".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) Set__EmailField(v cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.EmailField = v
	return t
}

// SetV__EmailField updates property "EmailField".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetV__EmailField(v AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.EmailField = cfz.V(v)
	return t
}

// Set__PasswordField updates property "PasswordField".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) Set__PasswordField(v cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PasswordField = v
	return t
}

// SetV__PasswordField updates property "PasswordField".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetV__PasswordField(v AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PasswordField = cfz.V(v)
	return t
}

// Set__PayloadType updates property "PayloadType".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) Set__PayloadType(v cfz.Expression[string]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PayloadType = v
	return t
}

// SetV__PayloadType updates property "PayloadType".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetV__PayloadType(v string) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PayloadType = cfz.V(v)
	return t
}

// Set__PhoneNumberFields updates property "PhoneNumberFields".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) Set__PhoneNumberFields(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PhoneNumberFields = v
	return t
}

// SetS__PhoneNumberFields updates property "PhoneNumberFields".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetS__PhoneNumberFields(v ...cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PhoneNumberFields = cfz.S(v...)
	return t
}

// SetSV__PhoneNumberFields updates property "PhoneNumberFields".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetSV__PhoneNumberFields(v ...AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.PhoneNumberFields = cfz.SV(v...)
	return t
}

// Set__UsernameField updates property "UsernameField".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) Set__UsernameField(v cfz.Expression[AWS_WAFv2_WebACL_FieldIdentifier]) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.UsernameField = v
	return t
}

// SetV__UsernameField updates property "UsernameField".
func (t AWS_WAFv2_WebACL_RequestInspectionACFP) SetV__UsernameField(v AWS_WAFv2_WebACL_FieldIdentifier) AWS_WAFv2_WebACL_RequestInspectionACFP {
	t.UsernameField = cfz.V(v)
	return t
}
