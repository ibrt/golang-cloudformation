// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GameLift_GameServerGroup_InstanceDefinition__Type is the CloudFormation type for AWS::GameLift::GameServerGroup.InstanceDefinition.
	AWS_GameLift_GameServerGroup_InstanceDefinition__Type = "AWS::GameLift::GameServerGroup.InstanceDefinition"
)

var (
	// AWS_GameLift_GameServerGroup_InstanceDefinition__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::GameServerGroup.InstanceDefinition.
	AWS_GameLift_GameServerGroup_InstanceDefinition__PropertiesMap = struct {
		InstanceType     string
		WeightedCapacity string
	}{
		InstanceType:     "InstanceType",
		WeightedCapacity: "WeightedCapacity",
	}

	// AWS_GameLift_GameServerGroup_InstanceDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::GameServerGroup.InstanceDefinition.
	AWS_GameLift_GameServerGroup_InstanceDefinition__PropertiesSlice = []string{
		AWS_GameLift_GameServerGroup_InstanceDefinition__PropertiesMap.InstanceType,
		AWS_GameLift_GameServerGroup_InstanceDefinition__PropertiesMap.WeightedCapacity,
	}
)

// AWS_GameLift_GameServerGroup_InstanceDefinition is a binding for AWS::GameLift::GameServerGroup.InstanceDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html
type AWS_GameLift_GameServerGroup_InstanceDefinition struct {
	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html#cfn-gamelift-gameservergroup-instancedefinition-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`

	// WeightedCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-instancedefinition.html#cfn-gamelift-gameservergroup-instancedefinition-weightedcapacity
	WeightedCapacity cfz.Expression[string] `json:"WeightedCapacity,omitempty"`
}

// New__AWS_GameLift_GameServerGroup_InstanceDefinition initializes a new AWS_GameLift_GameServerGroup_InstanceDefinition.
func New__AWS_GameLift_GameServerGroup_InstanceDefinition() AWS_GameLift_GameServerGroup_InstanceDefinition {
	return AWS_GameLift_GameServerGroup_InstanceDefinition{}
}

// GetType returns the CloudFormation type.
func (AWS_GameLift_GameServerGroup_InstanceDefinition) GetType() string {
	return AWS_GameLift_GameServerGroup_InstanceDefinition__Type
}

// Set__InstanceType updates property "InstanceType".
func (t AWS_GameLift_GameServerGroup_InstanceDefinition) Set__InstanceType(v cfz.Expression[string]) AWS_GameLift_GameServerGroup_InstanceDefinition {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t AWS_GameLift_GameServerGroup_InstanceDefinition) SetV__InstanceType(v string) AWS_GameLift_GameServerGroup_InstanceDefinition {
	t.InstanceType = cfz.V(v)
	return t
}

// Set__WeightedCapacity updates property "WeightedCapacity".
func (t AWS_GameLift_GameServerGroup_InstanceDefinition) Set__WeightedCapacity(v cfz.Expression[string]) AWS_GameLift_GameServerGroup_InstanceDefinition {
	t.WeightedCapacity = v
	return t
}

// SetV__WeightedCapacity updates property "WeightedCapacity".
func (t AWS_GameLift_GameServerGroup_InstanceDefinition) SetV__WeightedCapacity(v string) AWS_GameLift_GameServerGroup_InstanceDefinition {
	t.WeightedCapacity = cfz.V(v)
	return t
}
