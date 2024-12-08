// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotfleetwise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTFleetWise_Campaign_CollectionScheme__Type is the CloudFormation type for AWS::IoTFleetWise::Campaign.CollectionScheme.
	AWS_IoTFleetWise_Campaign_CollectionScheme__Type = "AWS::IoTFleetWise::Campaign.CollectionScheme"
)

var (
	// AWS_IoTFleetWise_Campaign_CollectionScheme__PropertiesMap reports all the CloudFormation properties for AWS::IoTFleetWise::Campaign.CollectionScheme.
	AWS_IoTFleetWise_Campaign_CollectionScheme__PropertiesMap = struct {
		ConditionBasedCollectionScheme string
		TimeBasedCollectionScheme      string
	}{
		ConditionBasedCollectionScheme: "ConditionBasedCollectionScheme",
		TimeBasedCollectionScheme:      "TimeBasedCollectionScheme",
	}

	// AWS_IoTFleetWise_Campaign_CollectionScheme__PropertiesSlice reports all the CloudFormation properties for AWS::IoTFleetWise::Campaign.CollectionScheme.
	AWS_IoTFleetWise_Campaign_CollectionScheme__PropertiesSlice = []string{
		AWS_IoTFleetWise_Campaign_CollectionScheme__PropertiesMap.ConditionBasedCollectionScheme,
		AWS_IoTFleetWise_Campaign_CollectionScheme__PropertiesMap.TimeBasedCollectionScheme,
	}
)

// AWS_IoTFleetWise_Campaign_CollectionScheme is a binding for AWS::IoTFleetWise::Campaign.CollectionScheme.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-collectionscheme.html
type AWS_IoTFleetWise_Campaign_CollectionScheme struct {
	// ConditionBasedCollectionScheme is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-collectionscheme.html#cfn-iotfleetwise-campaign-collectionscheme-conditionbasedcollectionscheme
	ConditionBasedCollectionScheme cfz.Expression[AWS_IoTFleetWise_Campaign_ConditionBasedCollectionScheme] `json:"ConditionBasedCollectionScheme,omitempty"`

	// TimeBasedCollectionScheme is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-collectionscheme.html#cfn-iotfleetwise-campaign-collectionscheme-timebasedcollectionscheme
	TimeBasedCollectionScheme cfz.Expression[AWS_IoTFleetWise_Campaign_TimeBasedCollectionScheme] `json:"TimeBasedCollectionScheme,omitempty"`
}

// New__AWS_IoTFleetWise_Campaign_CollectionScheme initializes a new AWS_IoTFleetWise_Campaign_CollectionScheme.
func New__AWS_IoTFleetWise_Campaign_CollectionScheme() AWS_IoTFleetWise_Campaign_CollectionScheme {
	return AWS_IoTFleetWise_Campaign_CollectionScheme{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTFleetWise_Campaign_CollectionScheme) GetType() string {
	return AWS_IoTFleetWise_Campaign_CollectionScheme__Type
}

// Set__ConditionBasedCollectionScheme updates property "ConditionBasedCollectionScheme".
func (t AWS_IoTFleetWise_Campaign_CollectionScheme) Set__ConditionBasedCollectionScheme(v cfz.Expression[AWS_IoTFleetWise_Campaign_ConditionBasedCollectionScheme]) AWS_IoTFleetWise_Campaign_CollectionScheme {
	t.ConditionBasedCollectionScheme = v
	return t
}

// SetV__ConditionBasedCollectionScheme updates property "ConditionBasedCollectionScheme".
func (t AWS_IoTFleetWise_Campaign_CollectionScheme) SetV__ConditionBasedCollectionScheme(v AWS_IoTFleetWise_Campaign_ConditionBasedCollectionScheme) AWS_IoTFleetWise_Campaign_CollectionScheme {
	t.ConditionBasedCollectionScheme = cfz.V(v)
	return t
}

// Set__TimeBasedCollectionScheme updates property "TimeBasedCollectionScheme".
func (t AWS_IoTFleetWise_Campaign_CollectionScheme) Set__TimeBasedCollectionScheme(v cfz.Expression[AWS_IoTFleetWise_Campaign_TimeBasedCollectionScheme]) AWS_IoTFleetWise_Campaign_CollectionScheme {
	t.TimeBasedCollectionScheme = v
	return t
}

// SetV__TimeBasedCollectionScheme updates property "TimeBasedCollectionScheme".
func (t AWS_IoTFleetWise_Campaign_CollectionScheme) SetV__TimeBasedCollectionScheme(v AWS_IoTFleetWise_Campaign_TimeBasedCollectionScheme) AWS_IoTFleetWise_Campaign_CollectionScheme {
	t.TimeBasedCollectionScheme = cfz.V(v)
	return t
}
