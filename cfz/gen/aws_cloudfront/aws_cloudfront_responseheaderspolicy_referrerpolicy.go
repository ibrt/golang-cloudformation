// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__Type is the CloudFormation type for AWS::CloudFront::ResponseHeadersPolicy.ReferrerPolicy.
	AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__Type = "AWS::CloudFront::ResponseHeadersPolicy.ReferrerPolicy"
)

var (
	// AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::ResponseHeadersPolicy.ReferrerPolicy.
	AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__PropertiesMap = struct {
		Override       string
		ReferrerPolicy string
	}{
		Override:       "Override",
		ReferrerPolicy: "ReferrerPolicy",
	}

	// AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::ResponseHeadersPolicy.ReferrerPolicy.
	AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__PropertiesSlice = []string{
		AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__PropertiesMap.Override,
		AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__PropertiesMap.ReferrerPolicy,
	}
)

// AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy is a binding for AWS::CloudFront::ResponseHeadersPolicy.ReferrerPolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.html
type AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy struct {
	// Override is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.html#cfn-cloudfront-responseheaderspolicy-referrerpolicy-override
	Override cfz.Expression[bool] `json:"Override,omitempty"`

	// ReferrerPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.html#cfn-cloudfront-responseheaderspolicy-referrerpolicy-referrerpolicy
	ReferrerPolicy cfz.Expression[string] `json:"ReferrerPolicy,omitempty"`
}

// New__AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy initializes a new AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy.
func New__AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy() AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy {
	return AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy) GetType() string {
	return AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy__Type
}

// Set__Override updates property "Override".
func (t AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy) Set__Override(v cfz.Expression[bool]) AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy {
	t.Override = v
	return t
}

// SetV__Override updates property "Override".
func (t AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy) SetV__Override(v bool) AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy {
	t.Override = cfz.V(v)
	return t
}

// Set__ReferrerPolicy updates property "ReferrerPolicy".
func (t AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy) Set__ReferrerPolicy(v cfz.Expression[string]) AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy {
	t.ReferrerPolicy = v
	return t
}

// SetV__ReferrerPolicy updates property "ReferrerPolicy".
func (t AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy) SetV__ReferrerPolicy(v string) AWS_CloudFront_ResponseHeadersPolicy_ReferrerPolicy {
	t.ReferrerPolicy = cfz.V(v)
	return t
}
