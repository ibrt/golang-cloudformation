// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_budgets

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Budgets_Budget_Notification__Type is the CloudFormation type for AWS::Budgets::Budget.Notification.
	AWS_Budgets_Budget_Notification__Type = "AWS::Budgets::Budget.Notification"
)

var (
	// AWS_Budgets_Budget_Notification__PropertiesMap reports all the CloudFormation properties for AWS::Budgets::Budget.Notification.
	AWS_Budgets_Budget_Notification__PropertiesMap = struct {
		ComparisonOperator string
		NotificationType   string
		Threshold          string
		ThresholdType      string
	}{
		ComparisonOperator: "ComparisonOperator",
		NotificationType:   "NotificationType",
		Threshold:          "Threshold",
		ThresholdType:      "ThresholdType",
	}

	// AWS_Budgets_Budget_Notification__PropertiesSlice reports all the CloudFormation properties for AWS::Budgets::Budget.Notification.
	AWS_Budgets_Budget_Notification__PropertiesSlice = []string{
		AWS_Budgets_Budget_Notification__PropertiesMap.ComparisonOperator,
		AWS_Budgets_Budget_Notification__PropertiesMap.NotificationType,
		AWS_Budgets_Budget_Notification__PropertiesMap.Threshold,
		AWS_Budgets_Budget_Notification__PropertiesMap.ThresholdType,
	}
)

// AWS_Budgets_Budget_Notification is a binding for AWS::Budgets::Budget.Notification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html
type AWS_Budgets_Budget_Notification struct {
	// ComparisonOperator is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-comparisonoperator
	ComparisonOperator cfz.Expression[string] `json:"ComparisonOperator,omitempty"`

	// NotificationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-notificationtype
	NotificationType cfz.Expression[string] `json:"NotificationType,omitempty"`

	// Threshold is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-threshold
	Threshold cfz.Expression[float64] `json:"Threshold,omitempty"`

	// ThresholdType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notification.html#cfn-budgets-budget-notification-thresholdtype
	ThresholdType cfz.Expression[string] `json:"ThresholdType,omitempty"`
}

// New__AWS_Budgets_Budget_Notification initializes a new AWS_Budgets_Budget_Notification.
func New__AWS_Budgets_Budget_Notification() AWS_Budgets_Budget_Notification {
	return AWS_Budgets_Budget_Notification{}
}

// GetType returns the CloudFormation type.
func (AWS_Budgets_Budget_Notification) GetType() string {
	return AWS_Budgets_Budget_Notification__Type
}

// Set__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_Budgets_Budget_Notification) Set__ComparisonOperator(v cfz.Expression[string]) AWS_Budgets_Budget_Notification {
	t.ComparisonOperator = v
	return t
}

// SetV__ComparisonOperator updates property "ComparisonOperator".
func (t AWS_Budgets_Budget_Notification) SetV__ComparisonOperator(v string) AWS_Budgets_Budget_Notification {
	t.ComparisonOperator = cfz.V(v)
	return t
}

// Set__NotificationType updates property "NotificationType".
func (t AWS_Budgets_Budget_Notification) Set__NotificationType(v cfz.Expression[string]) AWS_Budgets_Budget_Notification {
	t.NotificationType = v
	return t
}

// SetV__NotificationType updates property "NotificationType".
func (t AWS_Budgets_Budget_Notification) SetV__NotificationType(v string) AWS_Budgets_Budget_Notification {
	t.NotificationType = cfz.V(v)
	return t
}

// Set__Threshold updates property "Threshold".
func (t AWS_Budgets_Budget_Notification) Set__Threshold(v cfz.Expression[float64]) AWS_Budgets_Budget_Notification {
	t.Threshold = v
	return t
}

// SetV__Threshold updates property "Threshold".
func (t AWS_Budgets_Budget_Notification) SetV__Threshold(v float64) AWS_Budgets_Budget_Notification {
	t.Threshold = cfz.V(v)
	return t
}

// Set__ThresholdType updates property "ThresholdType".
func (t AWS_Budgets_Budget_Notification) Set__ThresholdType(v cfz.Expression[string]) AWS_Budgets_Budget_Notification {
	t.ThresholdType = v
	return t
}

// SetV__ThresholdType updates property "ThresholdType".
func (t AWS_Budgets_Budget_Notification) SetV__ThresholdType(v string) AWS_Budgets_Budget_Notification {
	t.ThresholdType = cfz.V(v)
	return t
}
