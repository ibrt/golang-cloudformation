// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ssm

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SSM_MaintenanceWindowTask_NotificationConfig__Type is the CloudFormation type for AWS::SSM::MaintenanceWindowTask.NotificationConfig.
	AWS_SSM_MaintenanceWindowTask_NotificationConfig__Type = "AWS::SSM::MaintenanceWindowTask.NotificationConfig"
)

var (
	// AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesMap reports all the CloudFormation properties for AWS::SSM::MaintenanceWindowTask.NotificationConfig.
	AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesMap = struct {
		NotificationArn    string
		NotificationEvents string
		NotificationType   string
	}{
		NotificationArn:    "NotificationArn",
		NotificationEvents: "NotificationEvents",
		NotificationType:   "NotificationType",
	}

	// AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SSM::MaintenanceWindowTask.NotificationConfig.
	AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesSlice = []string{
		AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesMap.NotificationArn,
		AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesMap.NotificationEvents,
		AWS_SSM_MaintenanceWindowTask_NotificationConfig__PropertiesMap.NotificationType,
	}
)

// AWS_SSM_MaintenanceWindowTask_NotificationConfig is a binding for AWS::SSM::MaintenanceWindowTask.NotificationConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html
type AWS_SSM_MaintenanceWindowTask_NotificationConfig struct {
	// NotificationArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html#cfn-ssm-maintenancewindowtask-notificationconfig-notificationarn
	NotificationArn cfz.Expression[string] `json:"NotificationArn,omitempty"`

	// NotificationEvents is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html#cfn-ssm-maintenancewindowtask-notificationconfig-notificationevents
	NotificationEvents cfz.ExpressionSlice[string] `json:"NotificationEvents,omitempty"`

	// NotificationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html#cfn-ssm-maintenancewindowtask-notificationconfig-notificationtype
	NotificationType cfz.Expression[string] `json:"NotificationType,omitempty"`
}

// New__AWS_SSM_MaintenanceWindowTask_NotificationConfig initializes a new AWS_SSM_MaintenanceWindowTask_NotificationConfig.
func New__AWS_SSM_MaintenanceWindowTask_NotificationConfig() AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	return AWS_SSM_MaintenanceWindowTask_NotificationConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_SSM_MaintenanceWindowTask_NotificationConfig) GetType() string {
	return AWS_SSM_MaintenanceWindowTask_NotificationConfig__Type
}

// Set__NotificationArn updates property "NotificationArn".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) Set__NotificationArn(v cfz.Expression[string]) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationArn = v
	return t
}

// SetV__NotificationArn updates property "NotificationArn".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) SetV__NotificationArn(v string) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationArn = cfz.V(v)
	return t
}

// Set__NotificationEvents updates property "NotificationEvents".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) Set__NotificationEvents(v cfz.ExpressionSlice[string]) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationEvents = v
	return t
}

// SetS__NotificationEvents updates property "NotificationEvents".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) SetS__NotificationEvents(v ...cfz.Expression[string]) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationEvents = cfz.S(v...)
	return t
}

// SetSV__NotificationEvents updates property "NotificationEvents".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) SetSV__NotificationEvents(v ...string) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationEvents = cfz.SV(v...)
	return t
}

// Set__NotificationType updates property "NotificationType".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) Set__NotificationType(v cfz.Expression[string]) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationType = v
	return t
}

// SetV__NotificationType updates property "NotificationType".
func (t AWS_SSM_MaintenanceWindowTask_NotificationConfig) SetV__NotificationType(v string) AWS_SSM_MaintenanceWindowTask_NotificationConfig {
	t.NotificationType = cfz.V(v)
	return t
}
