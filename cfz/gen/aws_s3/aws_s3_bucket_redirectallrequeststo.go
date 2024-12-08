// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_RedirectAllRequestsTo__Type is the CloudFormation type for AWS::S3::Bucket.RedirectAllRequestsTo.
	AWS_S3_Bucket_RedirectAllRequestsTo__Type = "AWS::S3::Bucket.RedirectAllRequestsTo"
)

var (
	// AWS_S3_Bucket_RedirectAllRequestsTo__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.RedirectAllRequestsTo.
	AWS_S3_Bucket_RedirectAllRequestsTo__PropertiesMap = struct {
		HostName string
		Protocol string
	}{
		HostName: "HostName",
		Protocol: "Protocol",
	}

	// AWS_S3_Bucket_RedirectAllRequestsTo__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.RedirectAllRequestsTo.
	AWS_S3_Bucket_RedirectAllRequestsTo__PropertiesSlice = []string{
		AWS_S3_Bucket_RedirectAllRequestsTo__PropertiesMap.HostName,
		AWS_S3_Bucket_RedirectAllRequestsTo__PropertiesMap.Protocol,
	}
)

// AWS_S3_Bucket_RedirectAllRequestsTo is a binding for AWS::S3::Bucket.RedirectAllRequestsTo.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-redirectallrequeststo.html
type AWS_S3_Bucket_RedirectAllRequestsTo struct {
	// HostName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-redirectallrequeststo.html#cfn-s3-bucket-redirectallrequeststo-hostname
	HostName cfz.Expression[string] `json:"HostName,omitempty"`

	// Protocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-redirectallrequeststo.html#cfn-s3-bucket-redirectallrequeststo-protocol
	Protocol cfz.Expression[string] `json:"Protocol,omitempty"`
}

// New__AWS_S3_Bucket_RedirectAllRequestsTo initializes a new AWS_S3_Bucket_RedirectAllRequestsTo.
func New__AWS_S3_Bucket_RedirectAllRequestsTo() AWS_S3_Bucket_RedirectAllRequestsTo {
	return AWS_S3_Bucket_RedirectAllRequestsTo{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_RedirectAllRequestsTo) GetType() string {
	return AWS_S3_Bucket_RedirectAllRequestsTo__Type
}

// Set__HostName updates property "HostName".
func (t AWS_S3_Bucket_RedirectAllRequestsTo) Set__HostName(v cfz.Expression[string]) AWS_S3_Bucket_RedirectAllRequestsTo {
	t.HostName = v
	return t
}

// SetV__HostName updates property "HostName".
func (t AWS_S3_Bucket_RedirectAllRequestsTo) SetV__HostName(v string) AWS_S3_Bucket_RedirectAllRequestsTo {
	t.HostName = cfz.V(v)
	return t
}

// Set__Protocol updates property "Protocol".
func (t AWS_S3_Bucket_RedirectAllRequestsTo) Set__Protocol(v cfz.Expression[string]) AWS_S3_Bucket_RedirectAllRequestsTo {
	t.Protocol = v
	return t
}

// SetV__Protocol updates property "Protocol".
func (t AWS_S3_Bucket_RedirectAllRequestsTo) SetV__Protocol(v string) AWS_S3_Bucket_RedirectAllRequestsTo {
	t.Protocol = cfz.V(v)
	return t
}
