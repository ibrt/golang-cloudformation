// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_JsonMatchPattern__Type is the CloudFormation type for AWS::WAFv2::WebACL.JsonMatchPattern.
	AWS_WAFv2_WebACL_JsonMatchPattern__Type = "AWS::WAFv2::WebACL.JsonMatchPattern"
)

var (
	// AWS_WAFv2_WebACL_JsonMatchPattern__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.JsonMatchPattern.
	AWS_WAFv2_WebACL_JsonMatchPattern__PropertiesMap = struct {
		All           string
		IncludedPaths string
	}{
		All:           "All",
		IncludedPaths: "IncludedPaths",
	}

	// AWS_WAFv2_WebACL_JsonMatchPattern__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.JsonMatchPattern.
	AWS_WAFv2_WebACL_JsonMatchPattern__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_JsonMatchPattern__PropertiesMap.All,
		AWS_WAFv2_WebACL_JsonMatchPattern__PropertiesMap.IncludedPaths,
	}
)

// AWS_WAFv2_WebACL_JsonMatchPattern is a binding for AWS::WAFv2::WebACL.JsonMatchPattern.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonmatchpattern.html
type AWS_WAFv2_WebACL_JsonMatchPattern struct {
	// All is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonmatchpattern.html#cfn-wafv2-webacl-jsonmatchpattern-all
	All cfz.Expression[json.RawMessage] `json:"All,omitempty"`

	// IncludedPaths is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonmatchpattern.html#cfn-wafv2-webacl-jsonmatchpattern-includedpaths
	IncludedPaths cfz.ExpressionSlice[string] `json:"IncludedPaths,omitempty"`
}

// New__AWS_WAFv2_WebACL_JsonMatchPattern initializes a new AWS_WAFv2_WebACL_JsonMatchPattern.
func New__AWS_WAFv2_WebACL_JsonMatchPattern() AWS_WAFv2_WebACL_JsonMatchPattern {
	return AWS_WAFv2_WebACL_JsonMatchPattern{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_JsonMatchPattern) GetType() string {
	return AWS_WAFv2_WebACL_JsonMatchPattern__Type
}

// Set__All updates property "All".
func (t AWS_WAFv2_WebACL_JsonMatchPattern) Set__All(v cfz.Expression[json.RawMessage]) AWS_WAFv2_WebACL_JsonMatchPattern {
	t.All = v
	return t
}

// SetV__All updates property "All".
func (t AWS_WAFv2_WebACL_JsonMatchPattern) SetV__All(v json.RawMessage) AWS_WAFv2_WebACL_JsonMatchPattern {
	t.All = cfz.V(v)
	return t
}

// Set__IncludedPaths updates property "IncludedPaths".
func (t AWS_WAFv2_WebACL_JsonMatchPattern) Set__IncludedPaths(v cfz.ExpressionSlice[string]) AWS_WAFv2_WebACL_JsonMatchPattern {
	t.IncludedPaths = v
	return t
}

// SetS__IncludedPaths updates property "IncludedPaths".
func (t AWS_WAFv2_WebACL_JsonMatchPattern) SetS__IncludedPaths(v ...cfz.Expression[string]) AWS_WAFv2_WebACL_JsonMatchPattern {
	t.IncludedPaths = cfz.S(v...)
	return t
}

// SetSV__IncludedPaths updates property "IncludedPaths".
func (t AWS_WAFv2_WebACL_JsonMatchPattern) SetSV__IncludedPaths(v ...string) AWS_WAFv2_WebACL_JsonMatchPattern {
	t.IncludedPaths = cfz.SV(v...)
	return t
}
