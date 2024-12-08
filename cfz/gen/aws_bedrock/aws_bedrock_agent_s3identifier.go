// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Agent_S3Identifier__Type is the CloudFormation type for AWS::Bedrock::Agent.S3Identifier.
	AWS_Bedrock_Agent_S3Identifier__Type = "AWS::Bedrock::Agent.S3Identifier"
)

var (
	// AWS_Bedrock_Agent_S3Identifier__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Agent.S3Identifier.
	AWS_Bedrock_Agent_S3Identifier__PropertiesMap = struct {
		S3BucketName string
		S3ObjectKey  string
	}{
		S3BucketName: "S3BucketName",
		S3ObjectKey:  "S3ObjectKey",
	}

	// AWS_Bedrock_Agent_S3Identifier__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Agent.S3Identifier.
	AWS_Bedrock_Agent_S3Identifier__PropertiesSlice = []string{
		AWS_Bedrock_Agent_S3Identifier__PropertiesMap.S3BucketName,
		AWS_Bedrock_Agent_S3Identifier__PropertiesMap.S3ObjectKey,
	}
)

// AWS_Bedrock_Agent_S3Identifier is a binding for AWS::Bedrock::Agent.S3Identifier.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-s3identifier.html
type AWS_Bedrock_Agent_S3Identifier struct {
	// S3BucketName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-s3identifier.html#cfn-bedrock-agent-s3identifier-s3bucketname
	S3BucketName cfz.Expression[string] `json:"S3BucketName,omitempty"`

	// S3ObjectKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-s3identifier.html#cfn-bedrock-agent-s3identifier-s3objectkey
	S3ObjectKey cfz.Expression[string] `json:"S3ObjectKey,omitempty"`
}

// New__AWS_Bedrock_Agent_S3Identifier initializes a new AWS_Bedrock_Agent_S3Identifier.
func New__AWS_Bedrock_Agent_S3Identifier() AWS_Bedrock_Agent_S3Identifier {
	return AWS_Bedrock_Agent_S3Identifier{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Agent_S3Identifier) GetType() string {
	return AWS_Bedrock_Agent_S3Identifier__Type
}

// Set__S3BucketName updates property "S3BucketName".
func (t AWS_Bedrock_Agent_S3Identifier) Set__S3BucketName(v cfz.Expression[string]) AWS_Bedrock_Agent_S3Identifier {
	t.S3BucketName = v
	return t
}

// SetV__S3BucketName updates property "S3BucketName".
func (t AWS_Bedrock_Agent_S3Identifier) SetV__S3BucketName(v string) AWS_Bedrock_Agent_S3Identifier {
	t.S3BucketName = cfz.V(v)
	return t
}

// Set__S3ObjectKey updates property "S3ObjectKey".
func (t AWS_Bedrock_Agent_S3Identifier) Set__S3ObjectKey(v cfz.Expression[string]) AWS_Bedrock_Agent_S3Identifier {
	t.S3ObjectKey = v
	return t
}

// SetV__S3ObjectKey updates property "S3ObjectKey".
func (t AWS_Bedrock_Agent_S3Identifier) SetV__S3ObjectKey(v string) AWS_Bedrock_Agent_S3Identifier {
	t.S3ObjectKey = cfz.V(v)
	return t
}
