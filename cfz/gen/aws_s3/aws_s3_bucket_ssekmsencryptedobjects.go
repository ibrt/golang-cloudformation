// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_SseKmsEncryptedObjects__Type is the CloudFormation type for AWS::S3::Bucket.SseKmsEncryptedObjects.
	AWS_S3_Bucket_SseKmsEncryptedObjects__Type = "AWS::S3::Bucket.SseKmsEncryptedObjects"
)

var (
	// AWS_S3_Bucket_SseKmsEncryptedObjects__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.SseKmsEncryptedObjects.
	AWS_S3_Bucket_SseKmsEncryptedObjects__PropertiesMap = struct {
		Status string
	}{
		Status: "Status",
	}

	// AWS_S3_Bucket_SseKmsEncryptedObjects__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.SseKmsEncryptedObjects.
	AWS_S3_Bucket_SseKmsEncryptedObjects__PropertiesSlice = []string{
		AWS_S3_Bucket_SseKmsEncryptedObjects__PropertiesMap.Status,
	}
)

// AWS_S3_Bucket_SseKmsEncryptedObjects is a binding for AWS::S3::Bucket.SseKmsEncryptedObjects.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html
type AWS_S3_Bucket_SseKmsEncryptedObjects struct {
	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html#cfn-s3-bucket-ssekmsencryptedobjects-status
	Status cfz.Expression[string] `json:"Status,omitempty"`
}

// New__AWS_S3_Bucket_SseKmsEncryptedObjects initializes a new AWS_S3_Bucket_SseKmsEncryptedObjects.
func New__AWS_S3_Bucket_SseKmsEncryptedObjects() AWS_S3_Bucket_SseKmsEncryptedObjects {
	return AWS_S3_Bucket_SseKmsEncryptedObjects{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_SseKmsEncryptedObjects) GetType() string {
	return AWS_S3_Bucket_SseKmsEncryptedObjects__Type
}

// Set__Status updates property "Status".
func (t AWS_S3_Bucket_SseKmsEncryptedObjects) Set__Status(v cfz.Expression[string]) AWS_S3_Bucket_SseKmsEncryptedObjects {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t AWS_S3_Bucket_SseKmsEncryptedObjects) SetV__Status(v string) AWS_S3_Bucket_SseKmsEncryptedObjects {
	t.Status = cfz.V(v)
	return t
}
