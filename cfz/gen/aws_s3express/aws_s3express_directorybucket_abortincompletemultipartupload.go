// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3express

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__Type is the CloudFormation type for AWS::S3Express::DirectoryBucket.AbortIncompleteMultipartUpload.
	AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__Type = "AWS::S3Express::DirectoryBucket.AbortIncompleteMultipartUpload"
)

var (
	// AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__PropertiesMap reports all the CloudFormation properties for AWS::S3Express::DirectoryBucket.AbortIncompleteMultipartUpload.
	AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__PropertiesMap = struct {
		DaysAfterInitiation string
	}{
		DaysAfterInitiation: "DaysAfterInitiation",
	}

	// AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__PropertiesSlice reports all the CloudFormation properties for AWS::S3Express::DirectoryBucket.AbortIncompleteMultipartUpload.
	AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__PropertiesSlice = []string{
		AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__PropertiesMap.DaysAfterInitiation,
	}
)

// AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload is a binding for AWS::S3Express::DirectoryBucket.AbortIncompleteMultipartUpload.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-abortincompletemultipartupload.html
type AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload struct {
	// DaysAfterInitiation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-abortincompletemultipartupload.html#cfn-s3express-directorybucket-abortincompletemultipartupload-daysafterinitiation
	DaysAfterInitiation cfz.Expression[int32] `json:"DaysAfterInitiation,omitempty"`
}

// New__AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload initializes a new AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload.
func New__AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload() AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload {
	return AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload{}
}

// GetType returns the CloudFormation type.
func (AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload) GetType() string {
	return AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload__Type
}

// Set__DaysAfterInitiation updates property "DaysAfterInitiation".
func (t AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload) Set__DaysAfterInitiation(v cfz.Expression[int32]) AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload {
	t.DaysAfterInitiation = v
	return t
}

// SetV__DaysAfterInitiation updates property "DaysAfterInitiation".
func (t AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload) SetV__DaysAfterInitiation(v int32) AWS_S3Express_DirectoryBucket_AbortIncompleteMultipartUpload {
	t.DaysAfterInitiation = cfz.V(v)
	return t
}
