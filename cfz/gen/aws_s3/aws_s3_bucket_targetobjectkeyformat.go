// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_TargetObjectKeyFormat__Type is the CloudFormation type for AWS::S3::Bucket.TargetObjectKeyFormat.
	AWS_S3_Bucket_TargetObjectKeyFormat__Type = "AWS::S3::Bucket.TargetObjectKeyFormat"
)

var (
	// AWS_S3_Bucket_TargetObjectKeyFormat__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.TargetObjectKeyFormat.
	AWS_S3_Bucket_TargetObjectKeyFormat__PropertiesMap = struct {
		PartitionedPrefix string
		SimplePrefix      string
	}{
		PartitionedPrefix: "PartitionedPrefix",
		SimplePrefix:      "SimplePrefix",
	}

	// AWS_S3_Bucket_TargetObjectKeyFormat__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.TargetObjectKeyFormat.
	AWS_S3_Bucket_TargetObjectKeyFormat__PropertiesSlice = []string{
		AWS_S3_Bucket_TargetObjectKeyFormat__PropertiesMap.PartitionedPrefix,
		AWS_S3_Bucket_TargetObjectKeyFormat__PropertiesMap.SimplePrefix,
	}
)

// AWS_S3_Bucket_TargetObjectKeyFormat is a binding for AWS::S3::Bucket.TargetObjectKeyFormat.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-targetobjectkeyformat.html
type AWS_S3_Bucket_TargetObjectKeyFormat struct {
	// PartitionedPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-targetobjectkeyformat.html#cfn-s3-bucket-targetobjectkeyformat-partitionedprefix
	PartitionedPrefix cfz.Expression[AWS_S3_Bucket_PartitionedPrefix] `json:"PartitionedPrefix,omitempty"`

	// SimplePrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-targetobjectkeyformat.html#cfn-s3-bucket-targetobjectkeyformat-simpleprefix
	SimplePrefix cfz.Expression[json.RawMessage] `json:"SimplePrefix,omitempty"`
}

// New__AWS_S3_Bucket_TargetObjectKeyFormat initializes a new AWS_S3_Bucket_TargetObjectKeyFormat.
func New__AWS_S3_Bucket_TargetObjectKeyFormat() AWS_S3_Bucket_TargetObjectKeyFormat {
	return AWS_S3_Bucket_TargetObjectKeyFormat{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_TargetObjectKeyFormat) GetType() string {
	return AWS_S3_Bucket_TargetObjectKeyFormat__Type
}

// Set__PartitionedPrefix updates property "PartitionedPrefix".
func (t AWS_S3_Bucket_TargetObjectKeyFormat) Set__PartitionedPrefix(v cfz.Expression[AWS_S3_Bucket_PartitionedPrefix]) AWS_S3_Bucket_TargetObjectKeyFormat {
	t.PartitionedPrefix = v
	return t
}

// SetV__PartitionedPrefix updates property "PartitionedPrefix".
func (t AWS_S3_Bucket_TargetObjectKeyFormat) SetV__PartitionedPrefix(v AWS_S3_Bucket_PartitionedPrefix) AWS_S3_Bucket_TargetObjectKeyFormat {
	t.PartitionedPrefix = cfz.V(v)
	return t
}

// Set__SimplePrefix updates property "SimplePrefix".
func (t AWS_S3_Bucket_TargetObjectKeyFormat) Set__SimplePrefix(v cfz.Expression[json.RawMessage]) AWS_S3_Bucket_TargetObjectKeyFormat {
	t.SimplePrefix = v
	return t
}

// SetV__SimplePrefix updates property "SimplePrefix".
func (t AWS_S3_Bucket_TargetObjectKeyFormat) SetV__SimplePrefix(v json.RawMessage) AWS_S3_Bucket_TargetObjectKeyFormat {
	t.SimplePrefix = cfz.V(v)
	return t
}
