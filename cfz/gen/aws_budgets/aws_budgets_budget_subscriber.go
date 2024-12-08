// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_budgets

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Budgets_Budget_Subscriber__Type is the CloudFormation type for AWS::Budgets::Budget.Subscriber.
	AWS_Budgets_Budget_Subscriber__Type = "AWS::Budgets::Budget.Subscriber"
)

var (
	// AWS_Budgets_Budget_Subscriber__PropertiesMap reports all the CloudFormation properties for AWS::Budgets::Budget.Subscriber.
	AWS_Budgets_Budget_Subscriber__PropertiesMap = struct {
		Address          string
		SubscriptionType string
	}{
		Address:          "Address",
		SubscriptionType: "SubscriptionType",
	}

	// AWS_Budgets_Budget_Subscriber__PropertiesSlice reports all the CloudFormation properties for AWS::Budgets::Budget.Subscriber.
	AWS_Budgets_Budget_Subscriber__PropertiesSlice = []string{
		AWS_Budgets_Budget_Subscriber__PropertiesMap.Address,
		AWS_Budgets_Budget_Subscriber__PropertiesMap.SubscriptionType,
	}
)

// AWS_Budgets_Budget_Subscriber is a binding for AWS::Budgets::Budget.Subscriber.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html
type AWS_Budgets_Budget_Subscriber struct {
	// Address is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-address
	Address cfz.Expression[string] `json:"Address,omitempty"`

	// SubscriptionType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-subscriber.html#cfn-budgets-budget-subscriber-subscriptiontype
	SubscriptionType cfz.Expression[string] `json:"SubscriptionType,omitempty"`
}

// New__AWS_Budgets_Budget_Subscriber initializes a new AWS_Budgets_Budget_Subscriber.
func New__AWS_Budgets_Budget_Subscriber() AWS_Budgets_Budget_Subscriber {
	return AWS_Budgets_Budget_Subscriber{}
}

// GetType returns the CloudFormation type.
func (AWS_Budgets_Budget_Subscriber) GetType() string {
	return AWS_Budgets_Budget_Subscriber__Type
}

// Set__Address updates property "Address".
func (t AWS_Budgets_Budget_Subscriber) Set__Address(v cfz.Expression[string]) AWS_Budgets_Budget_Subscriber {
	t.Address = v
	return t
}

// SetV__Address updates property "Address".
func (t AWS_Budgets_Budget_Subscriber) SetV__Address(v string) AWS_Budgets_Budget_Subscriber {
	t.Address = cfz.V(v)
	return t
}

// Set__SubscriptionType updates property "SubscriptionType".
func (t AWS_Budgets_Budget_Subscriber) Set__SubscriptionType(v cfz.Expression[string]) AWS_Budgets_Budget_Subscriber {
	t.SubscriptionType = v
	return t
}

// SetV__SubscriptionType updates property "SubscriptionType".
func (t AWS_Budgets_Budget_Subscriber) SetV__SubscriptionType(v string) AWS_Budgets_Budget_Subscriber {
	t.SubscriptionType = cfz.V(v)
	return t
}
