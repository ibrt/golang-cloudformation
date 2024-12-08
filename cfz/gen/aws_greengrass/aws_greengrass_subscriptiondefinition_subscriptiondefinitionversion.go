// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__Type is the CloudFormation type for AWS::Greengrass::SubscriptionDefinition.SubscriptionDefinitionVersion.
	AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__Type = "AWS::Greengrass::SubscriptionDefinition.SubscriptionDefinitionVersion"
)

var (
	// AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::SubscriptionDefinition.SubscriptionDefinitionVersion.
	AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__PropertiesMap = struct {
		Subscriptions string
	}{
		Subscriptions: "Subscriptions",
	}

	// AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::SubscriptionDefinition.SubscriptionDefinitionVersion.
	AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__PropertiesSlice = []string{
		AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__PropertiesMap.Subscriptions,
	}
)

// AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion is a binding for AWS::Greengrass::SubscriptionDefinition.SubscriptionDefinitionVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.html
type AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion struct {
	// Subscriptions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-subscriptiondefinition-subscriptiondefinitionversion.html#cfn-greengrass-subscriptiondefinition-subscriptiondefinitionversion-subscriptions
	Subscriptions cfz.ExpressionSlice[AWS_Greengrass_SubscriptionDefinition_Subscription] `json:"Subscriptions,omitempty"`
}

// New__AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion initializes a new AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion.
func New__AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion() AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion {
	return AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion{}
}

// GetType returns the CloudFormation type.
func (AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion) GetType() string {
	return AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion__Type
}

// Set__Subscriptions updates property "Subscriptions".
func (t AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion) Set__Subscriptions(v cfz.ExpressionSlice[AWS_Greengrass_SubscriptionDefinition_Subscription]) AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion {
	t.Subscriptions = v
	return t
}

// SetS__Subscriptions updates property "Subscriptions".
func (t AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion) SetS__Subscriptions(v ...cfz.Expression[AWS_Greengrass_SubscriptionDefinition_Subscription]) AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion {
	t.Subscriptions = cfz.S(v...)
	return t
}

// SetSV__Subscriptions updates property "Subscriptions".
func (t AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion) SetSV__Subscriptions(v ...AWS_Greengrass_SubscriptionDefinition_Subscription) AWS_Greengrass_SubscriptionDefinition_SubscriptionDefinitionVersion {
	t.Subscriptions = cfz.SV(v...)
	return t
}
