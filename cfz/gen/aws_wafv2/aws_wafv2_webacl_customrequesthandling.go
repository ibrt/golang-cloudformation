// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_CustomRequestHandling__Type is the CloudFormation type for AWS::WAFv2::WebACL.CustomRequestHandling.
	AWS_WAFv2_WebACL_CustomRequestHandling__Type = "AWS::WAFv2::WebACL.CustomRequestHandling"
)

var (
	// AWS_WAFv2_WebACL_CustomRequestHandling__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.CustomRequestHandling.
	AWS_WAFv2_WebACL_CustomRequestHandling__PropertiesMap = struct {
		InsertHeaders string
	}{
		InsertHeaders: "InsertHeaders",
	}

	// AWS_WAFv2_WebACL_CustomRequestHandling__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.CustomRequestHandling.
	AWS_WAFv2_WebACL_CustomRequestHandling__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_CustomRequestHandling__PropertiesMap.InsertHeaders,
	}
)

// AWS_WAFv2_WebACL_CustomRequestHandling is a binding for AWS::WAFv2::WebACL.CustomRequestHandling.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customrequesthandling.html
type AWS_WAFv2_WebACL_CustomRequestHandling struct {
	// InsertHeaders is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-customrequesthandling.html#cfn-wafv2-webacl-customrequesthandling-insertheaders
	InsertHeaders cfz.ExpressionSlice[AWS_WAFv2_WebACL_CustomHTTPHeader] `json:"InsertHeaders,omitempty"`
}

// New__AWS_WAFv2_WebACL_CustomRequestHandling initializes a new AWS_WAFv2_WebACL_CustomRequestHandling.
func New__AWS_WAFv2_WebACL_CustomRequestHandling() AWS_WAFv2_WebACL_CustomRequestHandling {
	return AWS_WAFv2_WebACL_CustomRequestHandling{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_CustomRequestHandling) GetType() string {
	return AWS_WAFv2_WebACL_CustomRequestHandling__Type
}

// Set__InsertHeaders updates property "InsertHeaders".
func (t AWS_WAFv2_WebACL_CustomRequestHandling) Set__InsertHeaders(v cfz.ExpressionSlice[AWS_WAFv2_WebACL_CustomHTTPHeader]) AWS_WAFv2_WebACL_CustomRequestHandling {
	t.InsertHeaders = v
	return t
}

// SetS__InsertHeaders updates property "InsertHeaders".
func (t AWS_WAFv2_WebACL_CustomRequestHandling) SetS__InsertHeaders(v ...cfz.Expression[AWS_WAFv2_WebACL_CustomHTTPHeader]) AWS_WAFv2_WebACL_CustomRequestHandling {
	t.InsertHeaders = cfz.S(v...)
	return t
}

// SetSV__InsertHeaders updates property "InsertHeaders".
func (t AWS_WAFv2_WebACL_CustomRequestHandling) SetSV__InsertHeaders(v ...AWS_WAFv2_WebACL_CustomHTTPHeader) AWS_WAFv2_WebACL_CustomRequestHandling {
	t.InsertHeaders = cfz.SV(v...)
	return t
}
