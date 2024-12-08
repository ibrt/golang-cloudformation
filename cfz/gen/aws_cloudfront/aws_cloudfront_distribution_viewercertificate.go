// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cloudfront

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_CloudFront_Distribution_ViewerCertificate__Type is the CloudFormation type for AWS::CloudFront::Distribution.ViewerCertificate.
	AWS_CloudFront_Distribution_ViewerCertificate__Type = "AWS::CloudFront::Distribution.ViewerCertificate"
)

var (
	// AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap reports all the CloudFormation properties for AWS::CloudFront::Distribution.ViewerCertificate.
	AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap = struct {
		AcmCertificateArn            string
		CloudFrontDefaultCertificate string
		IamCertificateId             string
		MinimumProtocolVersion       string
		SslSupportMethod             string
	}{
		AcmCertificateArn:            "AcmCertificateArn",
		CloudFrontDefaultCertificate: "CloudFrontDefaultCertificate",
		IamCertificateId:             "IamCertificateId",
		MinimumProtocolVersion:       "MinimumProtocolVersion",
		SslSupportMethod:             "SslSupportMethod",
	}

	// AWS_CloudFront_Distribution_ViewerCertificate__PropertiesSlice reports all the CloudFormation properties for AWS::CloudFront::Distribution.ViewerCertificate.
	AWS_CloudFront_Distribution_ViewerCertificate__PropertiesSlice = []string{
		AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap.AcmCertificateArn,
		AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap.CloudFrontDefaultCertificate,
		AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap.IamCertificateId,
		AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap.MinimumProtocolVersion,
		AWS_CloudFront_Distribution_ViewerCertificate__PropertiesMap.SslSupportMethod,
	}
)

// AWS_CloudFront_Distribution_ViewerCertificate is a binding for AWS::CloudFront::Distribution.ViewerCertificate.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html
type AWS_CloudFront_Distribution_ViewerCertificate struct {
	// AcmCertificateArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html#cfn-cloudfront-distribution-viewercertificate-acmcertificatearn
	AcmCertificateArn cfz.Expression[string] `json:"AcmCertificateArn,omitempty"`

	// CloudFrontDefaultCertificate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html#cfn-cloudfront-distribution-viewercertificate-cloudfrontdefaultcertificate
	CloudFrontDefaultCertificate cfz.Expression[bool] `json:"CloudFrontDefaultCertificate,omitempty"`

	// IamCertificateId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html#cfn-cloudfront-distribution-viewercertificate-iamcertificateid
	IamCertificateId cfz.Expression[string] `json:"IamCertificateId,omitempty"`

	// MinimumProtocolVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html#cfn-cloudfront-distribution-viewercertificate-minimumprotocolversion
	MinimumProtocolVersion cfz.Expression[string] `json:"MinimumProtocolVersion,omitempty"`

	// SslSupportMethod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-viewercertificate.html#cfn-cloudfront-distribution-viewercertificate-sslsupportmethod
	SslSupportMethod cfz.Expression[string] `json:"SslSupportMethod,omitempty"`
}

// New__AWS_CloudFront_Distribution_ViewerCertificate initializes a new AWS_CloudFront_Distribution_ViewerCertificate.
func New__AWS_CloudFront_Distribution_ViewerCertificate() AWS_CloudFront_Distribution_ViewerCertificate {
	return AWS_CloudFront_Distribution_ViewerCertificate{}
}

// GetType returns the CloudFormation type.
func (AWS_CloudFront_Distribution_ViewerCertificate) GetType() string {
	return AWS_CloudFront_Distribution_ViewerCertificate__Type
}

// Set__AcmCertificateArn updates property "AcmCertificateArn".
func (t AWS_CloudFront_Distribution_ViewerCertificate) Set__AcmCertificateArn(v cfz.Expression[string]) AWS_CloudFront_Distribution_ViewerCertificate {
	t.AcmCertificateArn = v
	return t
}

// SetV__AcmCertificateArn updates property "AcmCertificateArn".
func (t AWS_CloudFront_Distribution_ViewerCertificate) SetV__AcmCertificateArn(v string) AWS_CloudFront_Distribution_ViewerCertificate {
	t.AcmCertificateArn = cfz.V(v)
	return t
}

// Set__CloudFrontDefaultCertificate updates property "CloudFrontDefaultCertificate".
func (t AWS_CloudFront_Distribution_ViewerCertificate) Set__CloudFrontDefaultCertificate(v cfz.Expression[bool]) AWS_CloudFront_Distribution_ViewerCertificate {
	t.CloudFrontDefaultCertificate = v
	return t
}

// SetV__CloudFrontDefaultCertificate updates property "CloudFrontDefaultCertificate".
func (t AWS_CloudFront_Distribution_ViewerCertificate) SetV__CloudFrontDefaultCertificate(v bool) AWS_CloudFront_Distribution_ViewerCertificate {
	t.CloudFrontDefaultCertificate = cfz.V(v)
	return t
}

// Set__IamCertificateId updates property "IamCertificateId".
func (t AWS_CloudFront_Distribution_ViewerCertificate) Set__IamCertificateId(v cfz.Expression[string]) AWS_CloudFront_Distribution_ViewerCertificate {
	t.IamCertificateId = v
	return t
}

// SetV__IamCertificateId updates property "IamCertificateId".
func (t AWS_CloudFront_Distribution_ViewerCertificate) SetV__IamCertificateId(v string) AWS_CloudFront_Distribution_ViewerCertificate {
	t.IamCertificateId = cfz.V(v)
	return t
}

// Set__MinimumProtocolVersion updates property "MinimumProtocolVersion".
func (t AWS_CloudFront_Distribution_ViewerCertificate) Set__MinimumProtocolVersion(v cfz.Expression[string]) AWS_CloudFront_Distribution_ViewerCertificate {
	t.MinimumProtocolVersion = v
	return t
}

// SetV__MinimumProtocolVersion updates property "MinimumProtocolVersion".
func (t AWS_CloudFront_Distribution_ViewerCertificate) SetV__MinimumProtocolVersion(v string) AWS_CloudFront_Distribution_ViewerCertificate {
	t.MinimumProtocolVersion = cfz.V(v)
	return t
}

// Set__SslSupportMethod updates property "SslSupportMethod".
func (t AWS_CloudFront_Distribution_ViewerCertificate) Set__SslSupportMethod(v cfz.Expression[string]) AWS_CloudFront_Distribution_ViewerCertificate {
	t.SslSupportMethod = v
	return t
}

// SetV__SslSupportMethod updates property "SslSupportMethod".
func (t AWS_CloudFront_Distribution_ViewerCertificate) SetV__SslSupportMethod(v string) AWS_CloudFront_Distribution_ViewerCertificate {
	t.SslSupportMethod = cfz.V(v)
	return t
}
