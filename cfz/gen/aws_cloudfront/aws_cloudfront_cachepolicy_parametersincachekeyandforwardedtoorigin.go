// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__Type is the CloudFormation type for AWS::CloudFront::CachePolicy.ParametersInCacheKeyAndForwardedToOrigin.
	AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__Type = "AWS::CloudFront::CachePolicy.ParametersInCacheKeyAndForwardedToOrigin"
)

var (
	// AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::CachePolicy.ParametersInCacheKeyAndForwardedToOrigin.
	AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap = struct {
		CookiesConfig              string
		EnableAcceptEncodingBrotli string
		EnableAcceptEncodingGzip   string
		HeadersConfig              string
		QueryStringsConfig         string
	}{
		CookiesConfig:              "CookiesConfig",
		EnableAcceptEncodingBrotli: "EnableAcceptEncodingBrotli",
		EnableAcceptEncodingGzip:   "EnableAcceptEncodingGzip",
		HeadersConfig:              "HeadersConfig",
		QueryStringsConfig:         "QueryStringsConfig",
	}

	// AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::CachePolicy.ParametersInCacheKeyAndForwardedToOrigin.
	AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesSlice = []string{
		AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap.CookiesConfig,
		AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap.EnableAcceptEncodingBrotli,
		AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap.EnableAcceptEncodingGzip,
		AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap.HeadersConfig,
		AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__PropertiesMap.QueryStringsConfig,
	}
)

// AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin is a binding for AWS::CloudFront::CachePolicy.ParametersInCacheKeyAndForwardedToOrigin.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.html
type AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin struct {
	// CookiesConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.html#cfn-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin-cookiesconfig
	CookiesConfig cfz.Expression[AWS_CloudFront_CachePolicy_CookiesConfig] `json:"CookiesConfig,omitempty"`

	// EnableAcceptEncodingBrotli is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.html#cfn-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin-enableacceptencodingbrotli
	EnableAcceptEncodingBrotli cfz.Expression[bool] `json:"EnableAcceptEncodingBrotli,omitempty"`

	// EnableAcceptEncodingGzip is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.html#cfn-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin-enableacceptencodinggzip
	EnableAcceptEncodingGzip cfz.Expression[bool] `json:"EnableAcceptEncodingGzip,omitempty"`

	// HeadersConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.html#cfn-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin-headersconfig
	HeadersConfig cfz.Expression[AWS_CloudFront_CachePolicy_HeadersConfig] `json:"HeadersConfig,omitempty"`

	// QueryStringsConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin.html#cfn-cloudfront-cachepolicy-parametersincachekeyandforwardedtoorigin-querystringsconfig
	QueryStringsConfig cfz.Expression[AWS_CloudFront_CachePolicy_QueryStringsConfig] `json:"QueryStringsConfig,omitempty"`
}

// New__AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin initializes a new AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin.
func New__AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin() AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	return AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) GetType() string {
	return AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin__Type
}

// Set__CookiesConfig updates property "CookiesConfig".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) Set__CookiesConfig(v cfz.Expression[AWS_CloudFront_CachePolicy_CookiesConfig]) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.CookiesConfig = v
	return t
}

// SetV__CookiesConfig updates property "CookiesConfig".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) SetV__CookiesConfig(v AWS_CloudFront_CachePolicy_CookiesConfig) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.CookiesConfig = cfz.V(v)
	return t
}

// Set__EnableAcceptEncodingBrotli updates property "EnableAcceptEncodingBrotli".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) Set__EnableAcceptEncodingBrotli(v cfz.Expression[bool]) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.EnableAcceptEncodingBrotli = v
	return t
}

// SetV__EnableAcceptEncodingBrotli updates property "EnableAcceptEncodingBrotli".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) SetV__EnableAcceptEncodingBrotli(v bool) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.EnableAcceptEncodingBrotli = cfz.V(v)
	return t
}

// Set__EnableAcceptEncodingGzip updates property "EnableAcceptEncodingGzip".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) Set__EnableAcceptEncodingGzip(v cfz.Expression[bool]) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.EnableAcceptEncodingGzip = v
	return t
}

// SetV__EnableAcceptEncodingGzip updates property "EnableAcceptEncodingGzip".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) SetV__EnableAcceptEncodingGzip(v bool) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.EnableAcceptEncodingGzip = cfz.V(v)
	return t
}

// Set__HeadersConfig updates property "HeadersConfig".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) Set__HeadersConfig(v cfz.Expression[AWS_CloudFront_CachePolicy_HeadersConfig]) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.HeadersConfig = v
	return t
}

// SetV__HeadersConfig updates property "HeadersConfig".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) SetV__HeadersConfig(v AWS_CloudFront_CachePolicy_HeadersConfig) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.HeadersConfig = cfz.V(v)
	return t
}

// Set__QueryStringsConfig updates property "QueryStringsConfig".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) Set__QueryStringsConfig(v cfz.Expression[AWS_CloudFront_CachePolicy_QueryStringsConfig]) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.QueryStringsConfig = v
	return t
}

// SetV__QueryStringsConfig updates property "QueryStringsConfig".
func (t AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin) SetV__QueryStringsConfig(v AWS_CloudFront_CachePolicy_QueryStringsConfig) AWS_CloudFront_CachePolicy_ParametersInCacheKeyAndForwardedToOrigin {
	t.QueryStringsConfig = cfz.V(v)
	return t
}
