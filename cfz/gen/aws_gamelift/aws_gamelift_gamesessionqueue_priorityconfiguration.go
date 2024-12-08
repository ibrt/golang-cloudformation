// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GameLift_GameSessionQueue_PriorityConfiguration__Type is the CloudFormation type for AWS::GameLift::GameSessionQueue.PriorityConfiguration.
	AWS_GameLift_GameSessionQueue_PriorityConfiguration__Type = "AWS::GameLift::GameSessionQueue.PriorityConfiguration"
)

var (
	// AWS_GameLift_GameSessionQueue_PriorityConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::GameSessionQueue.PriorityConfiguration.
	AWS_GameLift_GameSessionQueue_PriorityConfiguration__PropertiesMap = struct {
		LocationOrder string
		PriorityOrder string
	}{
		LocationOrder: "LocationOrder",
		PriorityOrder: "PriorityOrder",
	}

	// AWS_GameLift_GameSessionQueue_PriorityConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::GameSessionQueue.PriorityConfiguration.
	AWS_GameLift_GameSessionQueue_PriorityConfiguration__PropertiesSlice = []string{
		AWS_GameLift_GameSessionQueue_PriorityConfiguration__PropertiesMap.LocationOrder,
		AWS_GameLift_GameSessionQueue_PriorityConfiguration__PropertiesMap.PriorityOrder,
	}
)

// AWS_GameLift_GameSessionQueue_PriorityConfiguration is a binding for AWS::GameLift::GameSessionQueue.PriorityConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html
type AWS_GameLift_GameSessionQueue_PriorityConfiguration struct {
	// LocationOrder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html#cfn-gamelift-gamesessionqueue-priorityconfiguration-locationorder
	LocationOrder cfz.ExpressionSlice[string] `json:"LocationOrder,omitempty"`

	// PriorityOrder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gamesessionqueue-priorityconfiguration.html#cfn-gamelift-gamesessionqueue-priorityconfiguration-priorityorder
	PriorityOrder cfz.ExpressionSlice[string] `json:"PriorityOrder,omitempty"`
}

// New__AWS_GameLift_GameSessionQueue_PriorityConfiguration initializes a new AWS_GameLift_GameSessionQueue_PriorityConfiguration.
func New__AWS_GameLift_GameSessionQueue_PriorityConfiguration() AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	return AWS_GameLift_GameSessionQueue_PriorityConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_GameLift_GameSessionQueue_PriorityConfiguration) GetType() string {
	return AWS_GameLift_GameSessionQueue_PriorityConfiguration__Type
}

// Set__LocationOrder updates property "LocationOrder".
func (t AWS_GameLift_GameSessionQueue_PriorityConfiguration) Set__LocationOrder(v cfz.ExpressionSlice[string]) AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	t.LocationOrder = v
	return t
}

// SetS__LocationOrder updates property "LocationOrder".
func (t AWS_GameLift_GameSessionQueue_PriorityConfiguration) SetS__LocationOrder(v ...cfz.Expression[string]) AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	t.LocationOrder = cfz.S(v...)
	return t
}

// SetSV__LocationOrder updates property "LocationOrder".
func (t AWS_GameLift_GameSessionQueue_PriorityConfiguration) SetSV__LocationOrder(v ...string) AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	t.LocationOrder = cfz.SV(v...)
	return t
}

// Set__PriorityOrder updates property "PriorityOrder".
func (t AWS_GameLift_GameSessionQueue_PriorityConfiguration) Set__PriorityOrder(v cfz.ExpressionSlice[string]) AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	t.PriorityOrder = v
	return t
}

// SetS__PriorityOrder updates property "PriorityOrder".
func (t AWS_GameLift_GameSessionQueue_PriorityConfiguration) SetS__PriorityOrder(v ...cfz.Expression[string]) AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	t.PriorityOrder = cfz.S(v...)
	return t
}

// SetSV__PriorityOrder updates property "PriorityOrder".
func (t AWS_GameLift_GameSessionQueue_PriorityConfiguration) SetSV__PriorityOrder(v ...string) AWS_GameLift_GameSessionQueue_PriorityConfiguration {
	t.PriorityOrder = cfz.SV(v...)
	return t
}
