// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appstream

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppStream_AppBlock_S3Location__Type is the CloudFormation type for AWS::AppStream::AppBlock.S3Location.
	AWS_AppStream_AppBlock_S3Location__Type = "AWS::AppStream::AppBlock.S3Location"
)

var (
	// AWS_AppStream_AppBlock_S3Location__PropertiesMap reports all the CloudFormation properties for AWS::AppStream::AppBlock.S3Location.
	AWS_AppStream_AppBlock_S3Location__PropertiesMap = struct {
		S3Bucket string
		S3Key    string
	}{
		S3Bucket: "S3Bucket",
		S3Key:    "S3Key",
	}

	// AWS_AppStream_AppBlock_S3Location__PropertiesSlice reports all the CloudFormation properties for AWS::AppStream::AppBlock.S3Location.
	AWS_AppStream_AppBlock_S3Location__PropertiesSlice = []string{
		AWS_AppStream_AppBlock_S3Location__PropertiesMap.S3Bucket,
		AWS_AppStream_AppBlock_S3Location__PropertiesMap.S3Key,
	}
)

// AWS_AppStream_AppBlock_S3Location is a binding for AWS::AppStream::AppBlock.S3Location.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-s3location.html
type AWS_AppStream_AppBlock_S3Location struct {
	// S3Bucket is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-s3location.html#cfn-appstream-appblock-s3location-s3bucket
	S3Bucket cfz.Expression[string] `json:"S3Bucket,omitempty"`

	// S3Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-appblock-s3location.html#cfn-appstream-appblock-s3location-s3key
	S3Key cfz.Expression[string] `json:"S3Key,omitempty"`
}

// New__AWS_AppStream_AppBlock_S3Location initializes a new AWS_AppStream_AppBlock_S3Location.
func New__AWS_AppStream_AppBlock_S3Location() AWS_AppStream_AppBlock_S3Location {
	return AWS_AppStream_AppBlock_S3Location{}
}

// GetType returns the CloudFormation type.
func (AWS_AppStream_AppBlock_S3Location) GetType() string {
	return AWS_AppStream_AppBlock_S3Location__Type
}

// Set__S3Bucket updates property "S3Bucket".
func (t AWS_AppStream_AppBlock_S3Location) Set__S3Bucket(v cfz.Expression[string]) AWS_AppStream_AppBlock_S3Location {
	t.S3Bucket = v
	return t
}

// SetV__S3Bucket updates property "S3Bucket".
func (t AWS_AppStream_AppBlock_S3Location) SetV__S3Bucket(v string) AWS_AppStream_AppBlock_S3Location {
	t.S3Bucket = cfz.V(v)
	return t
}

// Set__S3Key updates property "S3Key".
func (t AWS_AppStream_AppBlock_S3Location) Set__S3Key(v cfz.Expression[string]) AWS_AppStream_AppBlock_S3Location {
	t.S3Key = v
	return t
}

// SetV__S3Key updates property "S3Key".
func (t AWS_AppStream_AppBlock_S3Location) SetV__S3Key(v string) AWS_AppStream_AppBlock_S3Location {
	t.S3Key = cfz.V(v)
	return t
}
