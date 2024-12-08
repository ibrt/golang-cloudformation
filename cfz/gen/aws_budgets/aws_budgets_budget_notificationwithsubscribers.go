// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_budgets

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Budgets_Budget_NotificationWithSubscribers__Type is the CloudFormation type for AWS::Budgets::Budget.NotificationWithSubscribers.
	AWS_Budgets_Budget_NotificationWithSubscribers__Type = "AWS::Budgets::Budget.NotificationWithSubscribers"
)

var (
	// AWS_Budgets_Budget_NotificationWithSubscribers__PropertiesMap reports all the CloudFormation properties for AWS::Budgets::Budget.NotificationWithSubscribers.
	AWS_Budgets_Budget_NotificationWithSubscribers__PropertiesMap = struct {
		Notification string
		Subscribers  string
	}{
		Notification: "Notification",
		Subscribers:  "Subscribers",
	}

	// AWS_Budgets_Budget_NotificationWithSubscribers__PropertiesSlice reports all the CloudFormation properties for AWS::Budgets::Budget.NotificationWithSubscribers.
	AWS_Budgets_Budget_NotificationWithSubscribers__PropertiesSlice = []string{
		AWS_Budgets_Budget_NotificationWithSubscribers__PropertiesMap.Notification,
		AWS_Budgets_Budget_NotificationWithSubscribers__PropertiesMap.Subscribers,
	}
)

// AWS_Budgets_Budget_NotificationWithSubscribers is a binding for AWS::Budgets::Budget.NotificationWithSubscribers.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html
type AWS_Budgets_Budget_NotificationWithSubscribers struct {
	// Notification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-notification
	Notification cfz.Expression[AWS_Budgets_Budget_Notification] `json:"Notification,omitempty"`

	// Subscribers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-notificationwithsubscribers.html#cfn-budgets-budget-notificationwithsubscribers-subscribers
	Subscribers cfz.ExpressionSlice[AWS_Budgets_Budget_Subscriber] `json:"Subscribers,omitempty"`
}

// New__AWS_Budgets_Budget_NotificationWithSubscribers initializes a new AWS_Budgets_Budget_NotificationWithSubscribers.
func New__AWS_Budgets_Budget_NotificationWithSubscribers() AWS_Budgets_Budget_NotificationWithSubscribers {
	return AWS_Budgets_Budget_NotificationWithSubscribers{}
}

// GetType returns the CloudFormation type.
func (AWS_Budgets_Budget_NotificationWithSubscribers) GetType() string {
	return AWS_Budgets_Budget_NotificationWithSubscribers__Type
}

// Set__Notification updates property "Notification".
func (t AWS_Budgets_Budget_NotificationWithSubscribers) Set__Notification(v cfz.Expression[AWS_Budgets_Budget_Notification]) AWS_Budgets_Budget_NotificationWithSubscribers {
	t.Notification = v
	return t
}

// SetV__Notification updates property "Notification".
func (t AWS_Budgets_Budget_NotificationWithSubscribers) SetV__Notification(v AWS_Budgets_Budget_Notification) AWS_Budgets_Budget_NotificationWithSubscribers {
	t.Notification = cfz.V(v)
	return t
}

// Set__Subscribers updates property "Subscribers".
func (t AWS_Budgets_Budget_NotificationWithSubscribers) Set__Subscribers(v cfz.ExpressionSlice[AWS_Budgets_Budget_Subscriber]) AWS_Budgets_Budget_NotificationWithSubscribers {
	t.Subscribers = v
	return t
}

// SetS__Subscribers updates property "Subscribers".
func (t AWS_Budgets_Budget_NotificationWithSubscribers) SetS__Subscribers(v ...cfz.Expression[AWS_Budgets_Budget_Subscriber]) AWS_Budgets_Budget_NotificationWithSubscribers {
	t.Subscribers = cfz.S(v...)
	return t
}

// SetSV__Subscribers updates property "Subscribers".
func (t AWS_Budgets_Budget_NotificationWithSubscribers) SetSV__Subscribers(v ...AWS_Budgets_Budget_Subscriber) AWS_Budgets_Budget_NotificationWithSubscribers {
	t.Subscribers = cfz.SV(v...)
	return t
}
