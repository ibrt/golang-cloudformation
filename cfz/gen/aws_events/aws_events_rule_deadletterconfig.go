// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_events

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Events_Rule_DeadLetterConfig__Type is the CloudFormation type for AWS::Events::Rule.DeadLetterConfig.
	AWS_Events_Rule_DeadLetterConfig__Type = "AWS::Events::Rule.DeadLetterConfig"
)

var (
	// AWS_Events_Rule_DeadLetterConfig__PropertiesMap reports all the CloudFormation properties for AWS::Events::Rule.DeadLetterConfig.
	AWS_Events_Rule_DeadLetterConfig__PropertiesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Events_Rule_DeadLetterConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Events::Rule.DeadLetterConfig.
	AWS_Events_Rule_DeadLetterConfig__PropertiesSlice = []string{
		AWS_Events_Rule_DeadLetterConfig__PropertiesMap.Arn,
	}
)

// AWS_Events_Rule_DeadLetterConfig is a binding for AWS::Events::Rule.DeadLetterConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-deadletterconfig.html
type AWS_Events_Rule_DeadLetterConfig struct {
	// Arn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-deadletterconfig.html#cfn-events-rule-deadletterconfig-arn
	Arn cfz.Expression[string] `json:"Arn,omitempty"`
}

// New__AWS_Events_Rule_DeadLetterConfig initializes a new AWS_Events_Rule_DeadLetterConfig.
func New__AWS_Events_Rule_DeadLetterConfig() AWS_Events_Rule_DeadLetterConfig {
	return AWS_Events_Rule_DeadLetterConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Events_Rule_DeadLetterConfig) GetType() string {
	return AWS_Events_Rule_DeadLetterConfig__Type
}

// Set__Arn updates property "Arn".
func (t AWS_Events_Rule_DeadLetterConfig) Set__Arn(v cfz.Expression[string]) AWS_Events_Rule_DeadLetterConfig {
	t.Arn = v
	return t
}

// SetV__Arn updates property "Arn".
func (t AWS_Events_Rule_DeadLetterConfig) SetV__Arn(v string) AWS_Events_Rule_DeadLetterConfig {
	t.Arn = cfz.V(v)
	return t
}
