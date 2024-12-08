// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_WAFv2_WebACL_JsonBody__Type is the CloudFormation type for AWS::WAFv2::WebACL.JsonBody.
	AWS_WAFv2_WebACL_JsonBody__Type = "AWS::WAFv2::WebACL.JsonBody"
)

var (
	// AWS_WAFv2_WebACL_JsonBody__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::WebACL.JsonBody.
	AWS_WAFv2_WebACL_JsonBody__PropertiesMap = struct {
		InvalidFallbackBehavior string
		MatchPattern            string
		MatchScope              string
		OversizeHandling        string
	}{
		InvalidFallbackBehavior: "InvalidFallbackBehavior",
		MatchPattern:            "MatchPattern",
		MatchScope:              "MatchScope",
		OversizeHandling:        "OversizeHandling",
	}

	// AWS_WAFv2_WebACL_JsonBody__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::WebACL.JsonBody.
	AWS_WAFv2_WebACL_JsonBody__PropertiesSlice = []string{
		AWS_WAFv2_WebACL_JsonBody__PropertiesMap.InvalidFallbackBehavior,
		AWS_WAFv2_WebACL_JsonBody__PropertiesMap.MatchPattern,
		AWS_WAFv2_WebACL_JsonBody__PropertiesMap.MatchScope,
		AWS_WAFv2_WebACL_JsonBody__PropertiesMap.OversizeHandling,
	}
)

// AWS_WAFv2_WebACL_JsonBody is a binding for AWS::WAFv2::WebACL.JsonBody.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html
type AWS_WAFv2_WebACL_JsonBody struct {
	// InvalidFallbackBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-invalidfallbackbehavior
	InvalidFallbackBehavior cfz.Expression[string] `json:"InvalidFallbackBehavior,omitempty"`

	// MatchPattern is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-matchpattern
	MatchPattern cfz.Expression[AWS_WAFv2_WebACL_JsonMatchPattern] `json:"MatchPattern,omitempty"`

	// MatchScope is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-matchscope
	MatchScope cfz.Expression[string] `json:"MatchScope,omitempty"`

	// OversizeHandling is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-jsonbody.html#cfn-wafv2-webacl-jsonbody-oversizehandling
	OversizeHandling cfz.Expression[string] `json:"OversizeHandling,omitempty"`
}

// New__AWS_WAFv2_WebACL_JsonBody initializes a new AWS_WAFv2_WebACL_JsonBody.
func New__AWS_WAFv2_WebACL_JsonBody() AWS_WAFv2_WebACL_JsonBody {
	return AWS_WAFv2_WebACL_JsonBody{}
}

// GetType returns the CloudFormation type.
func (AWS_WAFv2_WebACL_JsonBody) GetType() string {
	return AWS_WAFv2_WebACL_JsonBody__Type
}

// Set__InvalidFallbackBehavior updates property "InvalidFallbackBehavior".
func (t AWS_WAFv2_WebACL_JsonBody) Set__InvalidFallbackBehavior(v cfz.Expression[string]) AWS_WAFv2_WebACL_JsonBody {
	t.InvalidFallbackBehavior = v
	return t
}

// SetV__InvalidFallbackBehavior updates property "InvalidFallbackBehavior".
func (t AWS_WAFv2_WebACL_JsonBody) SetV__InvalidFallbackBehavior(v string) AWS_WAFv2_WebACL_JsonBody {
	t.InvalidFallbackBehavior = cfz.V(v)
	return t
}

// Set__MatchPattern updates property "MatchPattern".
func (t AWS_WAFv2_WebACL_JsonBody) Set__MatchPattern(v cfz.Expression[AWS_WAFv2_WebACL_JsonMatchPattern]) AWS_WAFv2_WebACL_JsonBody {
	t.MatchPattern = v
	return t
}

// SetV__MatchPattern updates property "MatchPattern".
func (t AWS_WAFv2_WebACL_JsonBody) SetV__MatchPattern(v AWS_WAFv2_WebACL_JsonMatchPattern) AWS_WAFv2_WebACL_JsonBody {
	t.MatchPattern = cfz.V(v)
	return t
}

// Set__MatchScope updates property "MatchScope".
func (t AWS_WAFv2_WebACL_JsonBody) Set__MatchScope(v cfz.Expression[string]) AWS_WAFv2_WebACL_JsonBody {
	t.MatchScope = v
	return t
}

// SetV__MatchScope updates property "MatchScope".
func (t AWS_WAFv2_WebACL_JsonBody) SetV__MatchScope(v string) AWS_WAFv2_WebACL_JsonBody {
	t.MatchScope = cfz.V(v)
	return t
}

// Set__OversizeHandling updates property "OversizeHandling".
func (t AWS_WAFv2_WebACL_JsonBody) Set__OversizeHandling(v cfz.Expression[string]) AWS_WAFv2_WebACL_JsonBody {
	t.OversizeHandling = v
	return t
}

// SetV__OversizeHandling updates property "OversizeHandling".
func (t AWS_WAFv2_WebACL_JsonBody) SetV__OversizeHandling(v string) AWS_WAFv2_WebACL_JsonBody {
	t.OversizeHandling = cfz.V(v)
	return t
}
