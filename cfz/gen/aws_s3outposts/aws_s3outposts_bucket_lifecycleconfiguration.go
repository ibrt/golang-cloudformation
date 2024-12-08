// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3outposts

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3Outposts_Bucket_LifecycleConfiguration__Type is the CloudFormation type for AWS::S3Outposts::Bucket.LifecycleConfiguration.
	AWS_S3Outposts_Bucket_LifecycleConfiguration__Type = "AWS::S3Outposts::Bucket.LifecycleConfiguration"
)

var (
	// AWS_S3Outposts_Bucket_LifecycleConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::S3Outposts::Bucket.LifecycleConfiguration.
	AWS_S3Outposts_Bucket_LifecycleConfiguration__PropertiesMap = struct {
		Rules string
	}{
		Rules: "Rules",
	}

	// AWS_S3Outposts_Bucket_LifecycleConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::S3Outposts::Bucket.LifecycleConfiguration.
	AWS_S3Outposts_Bucket_LifecycleConfiguration__PropertiesSlice = []string{
		AWS_S3Outposts_Bucket_LifecycleConfiguration__PropertiesMap.Rules,
	}
)

// AWS_S3Outposts_Bucket_LifecycleConfiguration is a binding for AWS::S3Outposts::Bucket.LifecycleConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html
type AWS_S3Outposts_Bucket_LifecycleConfiguration struct {
	// Rules is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
	Rules cfz.ExpressionSlice[AWS_S3Outposts_Bucket_Rule] `json:"Rules,omitempty"`
}

// New__AWS_S3Outposts_Bucket_LifecycleConfiguration initializes a new AWS_S3Outposts_Bucket_LifecycleConfiguration.
func New__AWS_S3Outposts_Bucket_LifecycleConfiguration() AWS_S3Outposts_Bucket_LifecycleConfiguration {
	return AWS_S3Outposts_Bucket_LifecycleConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_S3Outposts_Bucket_LifecycleConfiguration) GetType() string {
	return AWS_S3Outposts_Bucket_LifecycleConfiguration__Type
}

// Set__Rules updates property "Rules".
func (t AWS_S3Outposts_Bucket_LifecycleConfiguration) Set__Rules(v cfz.ExpressionSlice[AWS_S3Outposts_Bucket_Rule]) AWS_S3Outposts_Bucket_LifecycleConfiguration {
	t.Rules = v
	return t
}

// SetS__Rules updates property "Rules".
func (t AWS_S3Outposts_Bucket_LifecycleConfiguration) SetS__Rules(v ...cfz.Expression[AWS_S3Outposts_Bucket_Rule]) AWS_S3Outposts_Bucket_LifecycleConfiguration {
	t.Rules = cfz.S(v...)
	return t
}

// SetSV__Rules updates property "Rules".
func (t AWS_S3Outposts_Bucket_LifecycleConfiguration) SetSV__Rules(v ...AWS_S3Outposts_Bucket_Rule) AWS_S3Outposts_Bucket_LifecycleConfiguration {
	t.Rules = cfz.SV(v...)
	return t
}
