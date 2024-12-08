// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_ByteMatchStatement__Type is the CloudFormation type for AWS::WAFv2::WebACL.ByteMatchStatement.
	AWS_WAFv2_WebACL_ByteMatchStatement__Type = "AWS::WAFv2::WebACL.ByteMatchStatement"
)

var (
	// AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.ByteMatchStatement.
	AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap = struct {
		FieldToMatch         string
		PositionalConstraint string
		SearchString         string
		SearchStringBase64   string
		TextTransformations  string
	}{
		FieldToMatch:         "FieldToMatch",
		PositionalConstraint: "PositionalConstraint",
		SearchString:         "SearchString",
		SearchStringBase64:   "SearchStringBase64",
		TextTransformations:  "TextTransformations",
	}

	// AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.ByteMatchStatement.
	AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap.FieldToMatch,
		AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap.PositionalConstraint,
		AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap.SearchString,
		AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap.SearchStringBase64,
		AWS_WAFv2_WebACL_ByteMatchStatement__PropertiesMap.TextTransformations,
	}
)

// AWS_WAFv2_WebACL_ByteMatchStatement is a binding for AWS::WAFv2::WebACL.ByteMatchStatement.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-bytematchstatement.html
type AWS_WAFv2_WebACL_ByteMatchStatement struct {
	// FieldToMatch is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-bytematchstatement.html#cfn-wafv2-webacl-bytematchstatement-fieldtomatch
	FieldToMatch cfz.Expression[AWS_WAFv2_WebACL_FieldToMatch] `json:"FieldToMatch,omitempty"`

	// PositionalConstraint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-bytematchstatement.html#cfn-wafv2-webacl-bytematchstatement-positionalconstraint
	PositionalConstraint cfz.Expression[string] `json:"PositionalConstraint,omitempty"`

	// SearchString is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-bytematchstatement.html#cfn-wafv2-webacl-bytematchstatement-searchstring
	SearchString cfz.Expression[string] `json:"SearchString,omitempty"`

	// SearchStringBase64 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-bytematchstatement.html#cfn-wafv2-webacl-bytematchstatement-searchstringbase64
	SearchStringBase64 cfz.Expression[string] `json:"SearchStringBase64,omitempty"`

	// TextTransformations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-bytematchstatement.html#cfn-wafv2-webacl-bytematchstatement-texttransformations
	TextTransformations cfz.ExpressionSlice[AWS_WAFv2_WebACL_TextTransformation] `json:"TextTransformations,omitempty"`
}

// New__AWS_WAFv2_WebACL_ByteMatchStatement initializes a new AWS_WAFv2_WebACL_ByteMatchStatement.
func New__AWS_WAFv2_WebACL_ByteMatchStatement() AWS_WAFv2_WebACL_ByteMatchStatement {
	return AWS_WAFv2_WebACL_ByteMatchStatement{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_ByteMatchStatement) GetType() string {
	return AWS_WAFv2_WebACL_ByteMatchStatement__Type
}

// Set__FieldToMatch updates property "FieldToMatch".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) Set__FieldToMatch(v cfz.Expression[AWS_WAFv2_WebACL_FieldToMatch]) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.FieldToMatch = v
	return t
}

// SetV__FieldToMatch updates property "FieldToMatch".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) SetV__FieldToMatch(v AWS_WAFv2_WebACL_FieldToMatch) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.FieldToMatch = cfz.V(v)
	return t
}

// Set__PositionalConstraint updates property "PositionalConstraint".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) Set__PositionalConstraint(v cfz.Expression[string]) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.PositionalConstraint = v
	return t
}

// SetV__PositionalConstraint updates property "PositionalConstraint".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) SetV__PositionalConstraint(v string) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.PositionalConstraint = cfz.V(v)
	return t
}

// Set__SearchString updates property "SearchString".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) Set__SearchString(v cfz.Expression[string]) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.SearchString = v
	return t
}

// SetV__SearchString updates property "SearchString".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) SetV__SearchString(v string) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.SearchString = cfz.V(v)
	return t
}

// Set__SearchStringBase64 updates property "SearchStringBase64".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) Set__SearchStringBase64(v cfz.Expression[string]) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.SearchStringBase64 = v
	return t
}

// SetV__SearchStringBase64 updates property "SearchStringBase64".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) SetV__SearchStringBase64(v string) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.SearchStringBase64 = cfz.V(v)
	return t
}

// Set__TextTransformations updates property "TextTransformations".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) Set__TextTransformations(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_TextTransformation]) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.TextTransformations = v
	return t
}

// SetS__TextTransformations updates property "TextTransformations".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) SetS__TextTransformations(v ...cfz.Expression[AWS_WAFv2_WebACL_TextTransformation]) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.TextTransformations = cfz.S(v...)
	return t
}

// SetSV__TextTransformations updates property "TextTransformations".
func (t AWS_WAFv2_WebACL_ByteMatchStatement) SetSV__TextTransformations(v ...AWS_WAFv2_WebACL_TextTransformation) AWS_WAFv2_WebACL_ByteMatchStatement {
	t.TextTransformations = cfz.SV(v...)
	return t
}
