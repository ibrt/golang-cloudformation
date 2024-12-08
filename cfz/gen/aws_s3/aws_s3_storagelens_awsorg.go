// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3_StorageLens_AwsOrg__Type is the CloudFormation type for AWS::S3::StorageLens.AwsOrg.
	AWS_S3_StorageLens_AwsOrg__Type = "AWS::S3::StorageLens.AwsOrg"
)

var (
	// AWS_S3_StorageLens_AwsOrg__PropertiesMap reports all the CloudFormation properties for AWS::S3::StorageLens.AwsOrg.
	AWS_S3_StorageLens_AwsOrg__PropertiesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_S3_StorageLens_AwsOrg__PropertiesSlice reports all the CloudFormation properties for AWS::S3::StorageLens.AwsOrg.
	AWS_S3_StorageLens_AwsOrg__PropertiesSlice = []string{
		AWS_S3_StorageLens_AwsOrg__PropertiesMap.Arn,
	}
)

// AWS_S3_StorageLens_AwsOrg is a binding for AWS::S3::StorageLens.AwsOrg.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-awsorg.html
type AWS_S3_StorageLens_AwsOrg struct {
	// Arn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-awsorg.html#cfn-s3-storagelens-awsorg-arn
	Arn cfz.Expression[string] `json:"Arn,omitempty"`
}

// New__AWS_S3_StorageLens_AwsOrg initializes a new AWS_S3_StorageLens_AwsOrg.
func New__AWS_S3_StorageLens_AwsOrg() AWS_S3_StorageLens_AwsOrg {
	return AWS_S3_StorageLens_AwsOrg{}
}

// GetType returns the CloudFormation type.
func (AWS_S3_StorageLens_AwsOrg) GetType() string {
	return AWS_S3_StorageLens_AwsOrg__Type
}

// Set__Arn updates property "Arn".
func (t AWS_S3_StorageLens_AwsOrg) Set__Arn(v cfz.Expression[string]) AWS_S3_StorageLens_AwsOrg {
	t.Arn = v
	return t
}

// SetV__Arn updates property "Arn".
func (t AWS_S3_StorageLens_AwsOrg) SetV__Arn(v string) AWS_S3_StorageLens_AwsOrg {
	t.Arn = cfz.V(v)
	return t
}
