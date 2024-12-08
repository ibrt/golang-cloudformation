// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__Type is the CloudFormation type for AWS::CloudFront::ResponseHeadersPolicy.XSSProtection.
	AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__Type = "AWS::CloudFront::ResponseHeadersPolicy.XSSProtection"
)

var (
	// AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::ResponseHeadersPolicy.XSSProtection.
	AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesMap = struct {
		ModeBlock  string
		Override   string
		Protection string
		ReportUri  string
	}{
		ModeBlock:  "ModeBlock",
		Override:   "Override",
		Protection: "Protection",
		ReportUri:  "ReportUri",
	}

	// AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::ResponseHeadersPolicy.XSSProtection.
	AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesSlice = []string{
		AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesMap.ModeBlock,
		AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesMap.Override,
		AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesMap.Protection,
		AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__PropertiesMap.ReportUri,
	}
)

// AWS_CloudFront_ResponseHeadersPolicy_XSSProtection is a binding for AWS::CloudFront::ResponseHeadersPolicy.XSSProtection.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html
type AWS_CloudFront_ResponseHeadersPolicy_XSSProtection struct {
	// ModeBlock is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-modeblock
	ModeBlock cfz.Expression[bool] `json:"ModeBlock,omitempty"`

	// Override is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-override
	Override cfz.Expression[bool] `json:"Override,omitempty"`

	// Protection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-protection
	Protection cfz.Expression[bool] `json:"Protection,omitempty"`

	// ReportUri is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-reporturi
	ReportUri cfz.Expression[string] `json:"ReportUri,omitempty"`
}

// New__AWS_CloudFront_ResponseHeadersPolicy_XSSProtection initializes a new AWS_CloudFront_ResponseHeadersPolicy_XSSProtection.
func New__AWS_CloudFront_ResponseHeadersPolicy_XSSProtection() AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	return AWS_CloudFront_ResponseHeadersPolicy_XSSProtection{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) GetType() string {
	return AWS_CloudFront_ResponseHeadersPolicy_XSSProtection__Type
}

// Set__ModeBlock updates property "ModeBlock".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) Set__ModeBlock(v cfz.Expression[bool]) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.ModeBlock = v
	return t
}

// SetV__ModeBlock updates property "ModeBlock".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) SetV__ModeBlock(v bool) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.ModeBlock = cfz.V(v)
	return t
}

// Set__Override updates property "Override".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) Set__Override(v cfz.Expression[bool]) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.Override = v
	return t
}

// SetV__Override updates property "Override".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) SetV__Override(v bool) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.Override = cfz.V(v)
	return t
}

// Set__Protection updates property "Protection".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) Set__Protection(v cfz.Expression[bool]) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.Protection = v
	return t
}

// SetV__Protection updates property "Protection".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) SetV__Protection(v bool) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.Protection = cfz.V(v)
	return t
}

// Set__ReportUri updates property "ReportUri".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) Set__ReportUri(v cfz.Expression[string]) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.ReportUri = v
	return t
}

// SetV__ReportUri updates property "ReportUri".
func (t AWS_CloudFront_ResponseHeadersPolicy_XSSProtection) SetV__ReportUri(v string) AWS_CloudFront_ResponseHeadersPolicy_XSSProtection {
	t.ReportUri = cfz.V(v)
	return t
}
