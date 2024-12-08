// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisanalyticsv2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisAnalyticsV2_Application_S3ContentLocation__Type is the CloudFormation type for AWS::KinesisAnalyticsV2::Application.S3ContentLocation.
	AWS_KinesisAnalyticsV2_Application_S3ContentLocation__Type = "AWS::KinesisAnalyticsV2::Application.S3ContentLocation"
)

var (
	// AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesMap reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.S3ContentLocation.
	AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesMap = struct {
		BucketARN     string
		FileKey       string
		ObjectVersion string
	}{
		BucketARN:     "BucketARN",
		FileKey:       "FileKey",
		ObjectVersion: "ObjectVersion",
	}

	// AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisAnalyticsV2::Application.S3ContentLocation.
	AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesSlice = []string{
		AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesMap.BucketARN,
		AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesMap.FileKey,
		AWS_KinesisAnalyticsV2_Application_S3ContentLocation__PropertiesMap.ObjectVersion,
	}
)

// AWS_KinesisAnalyticsV2_Application_S3ContentLocation is a binding for AWS::KinesisAnalyticsV2::Application.S3ContentLocation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html
type AWS_KinesisAnalyticsV2_Application_S3ContentLocation struct {
	// BucketARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-bucketarn
	BucketARN cfz.Expression[string] `json:"BucketARN,omitempty"`

	// FileKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-filekey
	FileKey cfz.Expression[string] `json:"FileKey,omitempty"`

	// ObjectVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-s3contentlocation.html#cfn-kinesisanalyticsv2-application-s3contentlocation-objectversion
	ObjectVersion cfz.Expression[string] `json:"ObjectVersion,omitempty"`
}

// New__AWS_KinesisAnalyticsV2_Application_S3ContentLocation initializes a new AWS_KinesisAnalyticsV2_Application_S3ContentLocation.
func New__AWS_KinesisAnalyticsV2_Application_S3ContentLocation() AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	return AWS_KinesisAnalyticsV2_Application_S3ContentLocation{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisAnalyticsV2_Application_S3ContentLocation) GetType() string {
	return AWS_KinesisAnalyticsV2_Application_S3ContentLocation__Type
}

// Set__BucketARN updates property "BucketARN".
func (t AWS_KinesisAnalyticsV2_Application_S3ContentLocation) Set__BucketARN(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	t.BucketARN = v
	return t
}

// SetV__BucketARN updates property "BucketARN".
func (t AWS_KinesisAnalyticsV2_Application_S3ContentLocation) SetV__BucketARN(v string) AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	t.BucketARN = cfz.V(v)
	return t
}

// Set__FileKey updates property "FileKey".
func (t AWS_KinesisAnalyticsV2_Application_S3ContentLocation) Set__FileKey(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	t.FileKey = v
	return t
}

// SetV__FileKey updates property "FileKey".
func (t AWS_KinesisAnalyticsV2_Application_S3ContentLocation) SetV__FileKey(v string) AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	t.FileKey = cfz.V(v)
	return t
}

// Set__ObjectVersion updates property "ObjectVersion".
func (t AWS_KinesisAnalyticsV2_Application_S3ContentLocation) Set__ObjectVersion(v cfz.Expression[string]) AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	t.ObjectVersion = v
	return t
}

// SetV__ObjectVersion updates property "ObjectVersion".
func (t AWS_KinesisAnalyticsV2_Application_S3ContentLocation) SetV__ObjectVersion(v string) AWS_KinesisAnalyticsV2_Application_S3ContentLocation {
	t.ObjectVersion = cfz.V(v)
	return t
}
