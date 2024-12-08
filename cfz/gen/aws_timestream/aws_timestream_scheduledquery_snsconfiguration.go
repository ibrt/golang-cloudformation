// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_timestream

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Timestream_ScheduledQuery_SnsConfiguration__Type is the CloudFormation type for AWS::Timestream::ScheduledQuery.SnsConfiguration.
	AWS_Timestream_ScheduledQuery_SnsConfiguration__Type = "AWS::Timestream::ScheduledQuery.SnsConfiguration"
)

var (
	// AWS_Timestream_ScheduledQuery_SnsConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Timestream::ScheduledQuery.SnsConfiguration.
	AWS_Timestream_ScheduledQuery_SnsConfiguration__PropertiesMap = struct {
		TopicArn string
	}{
		TopicArn: "TopicArn",
	}

	// AWS_Timestream_ScheduledQuery_SnsConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Timestream::ScheduledQuery.SnsConfiguration.
	AWS_Timestream_ScheduledQuery_SnsConfiguration__PropertiesSlice = []string{
		AWS_Timestream_ScheduledQuery_SnsConfiguration__PropertiesMap.TopicArn,
	}
)

// AWS_Timestream_ScheduledQuery_SnsConfiguration is a binding for AWS::Timestream::ScheduledQuery.SnsConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-snsconfiguration.html
type AWS_Timestream_ScheduledQuery_SnsConfiguration struct {
	// TopicArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-snsconfiguration.html#cfn-timestream-scheduledquery-snsconfiguration-topicarn
	TopicArn cfz.Expression[string] `json:"TopicArn,omitempty"`
}

// New__AWS_Timestream_ScheduledQuery_SnsConfiguration initializes a new AWS_Timestream_ScheduledQuery_SnsConfiguration.
func New__AWS_Timestream_ScheduledQuery_SnsConfiguration() AWS_Timestream_ScheduledQuery_SnsConfiguration {
	return AWS_Timestream_ScheduledQuery_SnsConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Timestream_ScheduledQuery_SnsConfiguration) GetType() string {
	return AWS_Timestream_ScheduledQuery_SnsConfiguration__Type
}

// Set__TopicArn updates property "TopicArn".
func (t AWS_Timestream_ScheduledQuery_SnsConfiguration) Set__TopicArn(v cfz.Expression[string]) AWS_Timestream_ScheduledQuery_SnsConfiguration {
	t.TopicArn = v
	return t
}

// SetV__TopicArn updates property "TopicArn".
func (t AWS_Timestream_ScheduledQuery_SnsConfiguration) SetV__TopicArn(v string) AWS_Timestream_ScheduledQuery_SnsConfiguration {
	t.TopicArn = cfz.V(v)
	return t
}
