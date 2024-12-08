// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_SingleHeader__Type is the CloudFormation type for AWS::WAFv2::WebACL.SingleHeader.
	AWS_WAFv2_WebACL_SingleHeader__Type = "AWS::WAFv2::WebACL.SingleHeader"
)

var (
	// AWS_WAFv2_WebACL_SingleHeader__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.SingleHeader.
	AWS_WAFv2_WebACL_SingleHeader__PropertiesMap = struct {
		Name string
	}{
		Name: "Name",
	}

	// AWS_WAFv2_WebACL_SingleHeader__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.SingleHeader.
	AWS_WAFv2_WebACL_SingleHeader__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_SingleHeader__PropertiesMap.Name,
	}
)

// AWS_WAFv2_WebACL_SingleHeader is a binding for AWS::WAFv2::WebACL.SingleHeader.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-singleheader.html
type AWS_WAFv2_WebACL_SingleHeader struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-singleheader.html#cfn-wafv2-webacl-singleheader-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_WAFv2_WebACL_SingleHeader initializes a new AWS_WAFv2_WebACL_SingleHeader.
func New__AWS_WAFv2_WebACL_SingleHeader() AWS_WAFv2_WebACL_SingleHeader {
	return AWS_WAFv2_WebACL_SingleHeader{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_SingleHeader) GetType() string {
	return AWS_WAFv2_WebACL_SingleHeader__Type
}

// Set__Name updates property "Name".
func (t AWS_WAFv2_WebACL_SingleHeader) Set__Name(v cfz.Expression[string]) AWS_WAFv2_WebACL_SingleHeader {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_WAFv2_WebACL_SingleHeader) SetV__Name(v string) AWS_WAFv2_WebACL_SingleHeader {
	t.Name = cfz.V(v)
	return t
}
