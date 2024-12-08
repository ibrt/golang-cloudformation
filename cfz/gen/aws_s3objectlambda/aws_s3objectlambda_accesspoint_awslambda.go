// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_s3objectlambda

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_S3ObjectLambda_AccessPoint_AwsLambda__Type is the CloudFormation type for AWS::S3ObjectLambda::AccessPoint.AwsLambda.
	AWS_S3ObjectLambda_AccessPoint_AwsLambda__Type = "AWS::S3ObjectLambda::AccessPoint.AwsLambda"
)

var (
	// AWS_S3ObjectLambda_AccessPoint_AwsLambda__PropertiesMap reports all the CloudFormation properties for AWS::S3ObjectLambda::AccessPoint.AwsLambda.
	AWS_S3ObjectLambda_AccessPoint_AwsLambda__PropertiesMap = struct {
		FunctionArn     string
		FunctionPayload string
	}{
		FunctionArn:     "FunctionArn",
		FunctionPayload: "FunctionPayload",
	}

	// AWS_S3ObjectLambda_AccessPoint_AwsLambda__PropertiesSlice reports all the CloudFormation properties for AWS::S3ObjectLambda::AccessPoint.AwsLambda.
	AWS_S3ObjectLambda_AccessPoint_AwsLambda__PropertiesSlice = []string{
		AWS_S3ObjectLambda_AccessPoint_AwsLambda__PropertiesMap.FunctionArn,
		AWS_S3ObjectLambda_AccessPoint_AwsLambda__PropertiesMap.FunctionPayload,
	}
)

// AWS_S3ObjectLambda_AccessPoint_AwsLambda is a binding for AWS::S3ObjectLambda::AccessPoint.AwsLambda.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-awslambda.html
type AWS_S3ObjectLambda_AccessPoint_AwsLambda struct {
	// FunctionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-awslambda.html#cfn-s3objectlambda-accesspoint-awslambda-functionarn
	FunctionArn cfz.Expression[string] `json:"FunctionArn,omitempty"`

	// FunctionPayload is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-awslambda.html#cfn-s3objectlambda-accesspoint-awslambda-functionpayload
	FunctionPayload cfz.Expression[string] `json:"FunctionPayload,omitempty"`
}

// New__AWS_S3ObjectLambda_AccessPoint_AwsLambda initializes a new AWS_S3ObjectLambda_AccessPoint_AwsLambda.
func New__AWS_S3ObjectLambda_AccessPoint_AwsLambda() AWS_S3ObjectLambda_AccessPoint_AwsLambda {
	return AWS_S3ObjectLambda_AccessPoint_AwsLambda{}
}

// GetType returns the CloudFormation type.
func (AWS_S3ObjectLambda_AccessPoint_AwsLambda) GetType() string {
	return AWS_S3ObjectLambda_AccessPoint_AwsLambda__Type
}

// Set__FunctionArn updates property "FunctionArn".
func (t AWS_S3ObjectLambda_AccessPoint_AwsLambda) Set__FunctionArn(v cfz.Expression[string]) AWS_S3ObjectLambda_AccessPoint_AwsLambda {
	t.FunctionArn = v
	return t
}

// SetV__FunctionArn updates property "FunctionArn".
func (t AWS_S3ObjectLambda_AccessPoint_AwsLambda) SetV__FunctionArn(v string) AWS_S3ObjectLambda_AccessPoint_AwsLambda {
	t.FunctionArn = cfz.V(v)
	return t
}

// Set__FunctionPayload updates property "FunctionPayload".
func (t AWS_S3ObjectLambda_AccessPoint_AwsLambda) Set__FunctionPayload(v cfz.Expression[string]) AWS_S3ObjectLambda_AccessPoint_AwsLambda {
	t.FunctionPayload = v
	return t
}

// SetV__FunctionPayload updates property "FunctionPayload".
func (t AWS_S3ObjectLambda_AccessPoint_AwsLambda) SetV__FunctionPayload(v string) AWS_S3ObjectLambda_AccessPoint_AwsLambda {
	t.FunctionPayload = cfz.V(v)
	return t
}
