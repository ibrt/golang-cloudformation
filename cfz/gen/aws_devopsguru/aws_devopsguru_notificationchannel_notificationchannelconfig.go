// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_devopsguru

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__Type is the CloudFormation type for AWS::DevOpsGuru::NotificationChannel.NotificationChannelConfig.
	AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__Type = "AWS::DevOpsGuru::NotificationChannel.NotificationChannelConfig"
)

var (
	// AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__PropertiesMap reports all the CloudFormation properties for AWS::DevOpsGuru::NotificationChannel.NotificationChannelConfig.
	AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__PropertiesMap = struct {
		Filters string
		Sns     string
	}{
		Filters: "Filters",
		Sns:     "Sns",
	}

	// AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__PropertiesSlice reports all the CloudFormation properties for AWS::DevOpsGuru::NotificationChannel.NotificationChannelConfig.
	AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__PropertiesSlice = []string{
		AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__PropertiesMap.Filters,
		AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__PropertiesMap.Sns,
	}
)

// AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig is a binding for AWS::DevOpsGuru::NotificationChannel.NotificationChannelConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationchannelconfig.html
type AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig struct {
	// Filters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationchannelconfig.html#cfn-devopsguru-notificationchannel-notificationchannelconfig-filters
	Filters cfz.Expression[AWS_DevOpsGuru_NotificationChannel_NotificationFilterConfig] `json:"Filters,omitempty"`

	// Sns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-notificationchannelconfig.html#cfn-devopsguru-notificationchannel-notificationchannelconfig-sns
	Sns cfz.Expression[AWS_DevOpsGuru_NotificationChannel_SnsChannelConfig] `json:"Sns,omitempty"`
}

// New__AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig initializes a new AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig.
func New__AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig() AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig {
	return AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig) GetType() string {
	return AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig__Type
}

// Set__Filters updates property "Filters".
func (t AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig) Set__Filters(v cfz.Expression[AWS_DevOpsGuru_NotificationChannel_NotificationFilterConfig]) AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig {
	t.Filters = v
	return t
}

// SetV__Filters updates property "Filters".
func (t AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig) SetV__Filters(v AWS_DevOpsGuru_NotificationChannel_NotificationFilterConfig) AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig {
	t.Filters = cfz.V(v)
	return t
}

// Set__Sns updates property "Sns".
func (t AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig) Set__Sns(v cfz.Expression[AWS_DevOpsGuru_NotificationChannel_SnsChannelConfig]) AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig {
	t.Sns = v
	return t
}

// SetV__Sns updates property "Sns".
func (t AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig) SetV__Sns(v AWS_DevOpsGuru_NotificationChannel_SnsChannelConfig) AWS_DevOpsGuru_NotificationChannel_NotificationChannelConfig {
	t.Sns = cfz.V(v)
	return t
}
