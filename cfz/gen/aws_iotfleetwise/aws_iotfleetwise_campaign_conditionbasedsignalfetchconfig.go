// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotfleetwise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__Type is the CloudFormation type for AWS::IoTFleetWise::Campaign.ConditionBasedSignalFetchConfig.
	AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__Type = "AWS::IoTFleetWise::Campaign.ConditionBasedSignalFetchConfig"
)

var (
	// AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__PropertiesMap reports all the CloudFormation properties for AWS::IoTFleetWise::Campaign.ConditionBasedSignalFetchConfig.
	AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__PropertiesMap = struct {
		ConditionExpression string
		TriggerMode         string
	}{
		ConditionExpression: "ConditionExpression",
		TriggerMode:         "TriggerMode",
	}

	// AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__PropertiesSlice reports all the CloudFormation properties for AWS::IoTFleetWise::Campaign.ConditionBasedSignalFetchConfig.
	AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__PropertiesSlice = []string{
		AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__PropertiesMap.ConditionExpression,
		AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__PropertiesMap.TriggerMode,
	}
)

// AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig is a binding for AWS::IoTFleetWise::Campaign.ConditionBasedSignalFetchConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html
type AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig struct {
	// ConditionExpression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-conditionbasedsignalfetchconfig-conditionexpression
	ConditionExpression cfz.Expression[string] `json:"ConditionExpression,omitempty"`

	// TriggerMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-conditionbasedsignalfetchconfig-triggermode
	TriggerMode cfz.Expression[string] `json:"TriggerMode,omitempty"`
}

// New__AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig initializes a new AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig.
func New__AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig() AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig {
	return AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig) GetType() string {
	return AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig__Type
}

// Set__ConditionExpression updates property "ConditionExpression".
func (t AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig) Set__ConditionExpression(v cfz.Expression[string]) AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig {
	t.ConditionExpression = v
	return t
}

// SetV__ConditionExpression updates property "ConditionExpression".
func (t AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig) SetV__ConditionExpression(v string) AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig {
	t.ConditionExpression = cfz.V(v)
	return t
}

// Set__TriggerMode updates property "TriggerMode".
func (t AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig) Set__TriggerMode(v cfz.Expression[string]) AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig {
	t.TriggerMode = v
	return t
}

// SetV__TriggerMode updates property "TriggerMode".
func (t AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig) SetV__TriggerMode(v string) AWS_IoTFleetWise_Campaign_ConditionBasedSignalFetchConfig {
	t.TriggerMode = cfz.V(v)
	return t
}
