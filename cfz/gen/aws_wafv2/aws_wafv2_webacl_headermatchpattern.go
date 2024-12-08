// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_HeaderMatchPattern__Type is the CloudFormation type for AWS::WAFv2::WebACL.HeaderMatchPattern.
	AWS_WAFv2_WebACL_HeaderMatchPattern__Type = "AWS::WAFv2::WebACL.HeaderMatchPattern"
)

var (
	// AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.HeaderMatchPattern.
	AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesMap = struct {
		All             string
		ExcludedHeaders string
		IncludedHeaders string
	}{
		All:             "All",
		ExcludedHeaders: "ExcludedHeaders",
		IncludedHeaders: "IncludedHeaders",
	}

	// AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.HeaderMatchPattern.
	AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesMap.All,
		AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesMap.ExcludedHeaders,
		AWS_WAFv2_WebACL_HeaderMatchPattern__PropertiesMap.IncludedHeaders,
	}
)

// AWS_WAFv2_WebACL_HeaderMatchPattern is a binding for AWS::WAFv2::WebACL.HeaderMatchPattern.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html
type AWS_WAFv2_WebACL_HeaderMatchPattern struct {
	// All is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html#cfn-wafv2-webacl-headermatchpattern-all
	All cfz.Expression[json.RawMessage] `json:"All,omitempty"`

	// ExcludedHeaders is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html#cfn-wafv2-webacl-headermatchpattern-excludedheaders
	ExcludedHeaders cfz.ExpressionSlice[string] `json:"ExcludedHeaders,omitempty"`

	// IncludedHeaders is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-headermatchpattern.html#cfn-wafv2-webacl-headermatchpattern-includedheaders
	IncludedHeaders cfz.ExpressionSlice[string] `json:"IncludedHeaders,omitempty"`
}

// New__AWS_WAFv2_WebACL_HeaderMatchPattern initializes a new AWS_WAFv2_WebACL_HeaderMatchPattern.
func New__AWS_WAFv2_WebACL_HeaderMatchPattern() AWS_WAFv2_WebACL_HeaderMatchPattern {
	return AWS_WAFv2_WebACL_HeaderMatchPattern{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_HeaderMatchPattern) GetType() string {
	return AWS_WAFv2_WebACL_HeaderMatchPattern__Type
}

// Set__All updates property "All".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) Set__All(v cfz.Expression[json.RawMessage]) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.All = v
	return t
}

// SetV__All updates property "All".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) SetV__All(v json.RawMessage) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.All = cfz.V(v)
	return t
}

// Set__ExcludedHeaders updates property "ExcludedHeaders".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) Set__ExcludedHeaders(v cfz.ExpressionSlice[string]) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.ExcludedHeaders = v
	return t
}

// SetS__ExcludedHeaders updates property "ExcludedHeaders".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) SetS__ExcludedHeaders(v ...cfz.Expression[string]) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.ExcludedHeaders = cfz.S(v...)
	return t
}

// SetSV__ExcludedHeaders updates property "ExcludedHeaders".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) SetSV__ExcludedHeaders(v ...string) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.ExcludedHeaders = cfz.SV(v...)
	return t
}

// Set__IncludedHeaders updates property "IncludedHeaders".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) Set__IncludedHeaders(v cfz.ExpressionSlice[string]) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.IncludedHeaders = v
	return t
}

// SetS__IncludedHeaders updates property "IncludedHeaders".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) SetS__IncludedHeaders(v ...cfz.Expression[string]) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.IncludedHeaders = cfz.S(v...)
	return t
}

// SetSV__IncludedHeaders updates property "IncludedHeaders".
func (t AWS_WAFv2_WebACL_HeaderMatchPattern) SetSV__IncludedHeaders(v ...string) AWS_WAFv2_WebACL_HeaderMatchPattern {
	t.IncludedHeaders = cfz.SV(v...)
	return t
}
