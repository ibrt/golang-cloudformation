// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Lex_Bot_TextLogDestination__Type is the CloudFormation type for AWS::Lex::Bot.TextLogDestination.
	AWS_Lex_Bot_TextLogDestination__Type = "AWS::Lex::Bot.TextLogDestination"
)

var (
	// AWS_Lex_Bot_TextLogDestination__PropertiesMap reports all the CloudFormation properties for AWS::Lex::Bot.TextLogDestination.
	AWS_Lex_Bot_TextLogDestination__PropertiesMap = struct {
		CloudWatch string
	}{
		CloudWatch: "CloudWatch",
	}

	// AWS_Lex_Bot_TextLogDestination__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::Bot.TextLogDestination.
	AWS_Lex_Bot_TextLogDestination__PropertiesSlice = []string{
		AWS_Lex_Bot_TextLogDestination__PropertiesMap.CloudWatch,
	}
)

// AWS_Lex_Bot_TextLogDestination is a binding for AWS::Lex::Bot.TextLogDestination.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-textlogdestination.html
type AWS_Lex_Bot_TextLogDestination struct {
	// CloudWatch is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-textlogdestination.html#cfn-lex-bot-textlogdestination-cloudwatch
	CloudWatch cfz.Expression[AWS_Lex_Bot_CloudWatchLogGroupLogDestination] `json:"CloudWatch,omitempty"`
}

// New__AWS_Lex_Bot_TextLogDestination initializes a new AWS_Lex_Bot_TextLogDestination.
func New__AWS_Lex_Bot_TextLogDestination() AWS_Lex_Bot_TextLogDestination {
	return AWS_Lex_Bot_TextLogDestination{}
}

// GetType returns the CloudFormation type.
func (AWS_Lex_Bot_TextLogDestination) GetType() string {
	return AWS_Lex_Bot_TextLogDestination__Type
}

// Set__CloudWatch updates property "CloudWatch".
func (t AWS_Lex_Bot_TextLogDestination) Set__CloudWatch(v cfz.Expression[AWS_Lex_Bot_CloudWatchLogGroupLogDestination]) AWS_Lex_Bot_TextLogDestination {
	t.CloudWatch = v
	return t
}

// SetV__CloudWatch updates property "CloudWatch".
func (t AWS_Lex_Bot_TextLogDestination) SetV__CloudWatch(v AWS_Lex_Bot_CloudWatchLogGroupLogDestination) AWS_Lex_Bot_TextLogDestination {
	t.CloudWatch = cfz.V(v)
	return t
}
