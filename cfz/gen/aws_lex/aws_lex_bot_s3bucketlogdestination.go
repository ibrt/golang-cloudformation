// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_S3BucketLogDestination__Type is the CloudFormation type for AWS::Lex::Bot.S3BucketLogDestination.
	AWS_Lex_Bot_S3BucketLogDestination__Type = "AWS::Lex::Bot.S3BucketLogDestination"
)

var (
	// AWS_Lex_Bot_S3BucketLogDestination__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.S3BucketLogDestination.
	AWS_Lex_Bot_S3BucketLogDestination__PropertiesMap = struct {
		KmsKeyArn   string
		LogPrefix   string
		S3BucketArn string
	}{
		KmsKeyArn:   "KmsKeyArn",
		LogPrefix:   "LogPrefix",
		S3BucketArn: "S3BucketArn",
	}

	// AWS_Lex_Bot_S3BucketLogDestination__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.S3BucketLogDestination.
	AWS_Lex_Bot_S3BucketLogDestination__PropertiesSlice = []string{
		AWS_Lex_Bot_S3BucketLogDestination__PropertiesMap.KmsKeyArn,
		AWS_Lex_Bot_S3BucketLogDestination__PropertiesMap.LogPrefix,
		AWS_Lex_Bot_S3BucketLogDestination__PropertiesMap.S3BucketArn,
	}
)

// AWS_Lex_Bot_S3BucketLogDestination is a binding for AWS::Lex::Bot.S3BucketLogDestination.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html
type AWS_Lex_Bot_S3BucketLogDestination struct {
	// KmsKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-kmskeyarn
	KmsKeyArn cfz.Expression[string] `json:"KmsKeyArn,omitempty"`

	// LogPrefix is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-logprefix
	LogPrefix cfz.Expression[string] `json:"LogPrefix,omitempty"`

	// S3BucketArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-s3bucketlogdestination.html#cfn-lex-bot-s3bucketlogdestination-s3bucketarn
	S3BucketArn cfz.Expression[string] `json:"S3BucketArn,omitempty"`
}

// New__AWS_Lex_Bot_S3BucketLogDestination initializes a new AWS_Lex_Bot_S3BucketLogDestination.
func New__AWS_Lex_Bot_S3BucketLogDestination() AWS_Lex_Bot_S3BucketLogDestination {
	return AWS_Lex_Bot_S3BucketLogDestination{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_S3BucketLogDestination) GetType() string {
	return AWS_Lex_Bot_S3BucketLogDestination__Type
}

// Set__KmsKeyArn updates property "KmsKeyArn".
func (t AWS_Lex_Bot_S3BucketLogDestination) Set__KmsKeyArn(v cfz.Expression[string]) AWS_Lex_Bot_S3BucketLogDestination {
	t.KmsKeyArn = v
	return t
}

// SetV__KmsKeyArn updates property "KmsKeyArn".
func (t AWS_Lex_Bot_S3BucketLogDestination) SetV__KmsKeyArn(v string) AWS_Lex_Bot_S3BucketLogDestination {
	t.KmsKeyArn = cfz.V(v)
	return t
}

// Set__LogPrefix updates property "LogPrefix".
func (t AWS_Lex_Bot_S3BucketLogDestination) Set__LogPrefix(v cfz.Expression[string]) AWS_Lex_Bot_S3BucketLogDestination {
	t.LogPrefix = v
	return t
}

// SetV__LogPrefix updates property "LogPrefix".
func (t AWS_Lex_Bot_S3BucketLogDestination) SetV__LogPrefix(v string) AWS_Lex_Bot_S3BucketLogDestination {
	t.LogPrefix = cfz.V(v)
	return t
}

// Set__S3BucketArn updates property "S3BucketArn".
func (t AWS_Lex_Bot_S3BucketLogDestination) Set__S3BucketArn(v cfz.Expression[string]) AWS_Lex_Bot_S3BucketLogDestination {
	t.S3BucketArn = v
	return t
}

// SetV__S3BucketArn updates property "S3BucketArn".
func (t AWS_Lex_Bot_S3BucketLogDestination) SetV__S3BucketArn(v string) AWS_Lex_Bot_S3BucketLogDestination {
	t.S3BucketArn = cfz.V(v)
	return t
}
