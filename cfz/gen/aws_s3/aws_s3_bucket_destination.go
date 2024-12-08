// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_Bucket_Destination__Type is the CloudFormation type for AWS::S3::Bucket.Destination.
	AWS_S3_Bucket_Destination__Type = "AWS::S3::Bucket.Destination"
)

var (
	// AWS_S3_Bucket_Destination__PropertiesMap reports all the CloudFormation properties for AWS::S3::Bucket.Destination.
	AWS_S3_Bucket_Destination__PropertiesMap = struct {
		BucketAccountId string
		BucketArn       string
		Format          string
		Prefix          string
	}{
		BucketAccountId: "BucketAccountId",
		BucketArn:       "BucketArn",
		Format:          "Format",
		Prefix:          "Prefix",
	}

	// AWS_S3_Bucket_Destination__PropertiesSlice reports all the CloudFormation properties for AWS::S3::Bucket.Destination.
	AWS_S3_Bucket_Destination__PropertiesSlice = []string{
		AWS_S3_Bucket_Destination__PropertiesMap.BucketAccountId,
		AWS_S3_Bucket_Destination__PropertiesMap.BucketArn,
		AWS_S3_Bucket_Destination__PropertiesMap.Format,
		AWS_S3_Bucket_Destination__PropertiesMap.Prefix,
	}
)

// AWS_S3_Bucket_Destination is a binding for AWS::S3::Bucket.Destination.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html
type AWS_S3_Bucket_Destination struct {
	// BucketAccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-bucketaccountid
	BucketAccountId cfz.Expression[string] `json:"BucketAccountId,omitempty"`

	// BucketArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-bucketarn
	BucketArn cfz.Expression[string] `json:"BucketArn,omitempty"`

	// Format is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-format
	Format cfz.Expression[string] `json:"Format,omitempty"`

	// Prefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-destination.html#cfn-s3-bucket-destination-prefix
	Prefix cfz.Expression[string] `json:"Prefix,omitempty"`
}

// New__AWS_S3_Bucket_Destination initializes a new AWS_S3_Bucket_Destination.
func New__AWS_S3_Bucket_Destination() AWS_S3_Bucket_Destination {
	return AWS_S3_Bucket_Destination{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_Bucket_Destination) GetType() string {
	return AWS_S3_Bucket_Destination__Type
}

// Set__BucketAccountId updates property "BucketAccountId".
func (t AWS_S3_Bucket_Destination) Set__BucketAccountId(v cfz.Expression[string]) AWS_S3_Bucket_Destination {
	t.BucketAccountId = v
	return t
}

// SetV__BucketAccountId updates property "BucketAccountId".
func (t AWS_S3_Bucket_Destination) SetV__BucketAccountId(v string) AWS_S3_Bucket_Destination {
	t.BucketAccountId = cfz.V(v)
	return t
}

// Set__BucketArn updates property "BucketArn".
func (t AWS_S3_Bucket_Destination) Set__BucketArn(v cfz.Expression[string]) AWS_S3_Bucket_Destination {
	t.BucketArn = v
	return t
}

// SetV__BucketArn updates property "BucketArn".
func (t AWS_S3_Bucket_Destination) SetV__BucketArn(v string) AWS_S3_Bucket_Destination {
	t.BucketArn = cfz.V(v)
	return t
}

// Set__Format updates property "Format".
func (t AWS_S3_Bucket_Destination) Set__Format(v cfz.Expression[string]) AWS_S3_Bucket_Destination {
	t.Format = v
	return t
}

// SetV__Format updates property "Format".
func (t AWS_S3_Bucket_Destination) SetV__Format(v string) AWS_S3_Bucket_Destination {
	t.Format = cfz.V(v)
	return t
}

// Set__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_Destination) Set__Prefix(v cfz.Expression[string]) AWS_S3_Bucket_Destination {
	t.Prefix = v
	return t
}

// SetV__Prefix updates property "Prefix".
func (t AWS_S3_Bucket_Destination) SetV__Prefix(v string) AWS_S3_Bucket_Destination {
	t.Prefix = cfz.V(v)
	return t
}
