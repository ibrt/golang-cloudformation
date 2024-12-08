// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_stepfunctions

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__Type is the CloudFormation type for AWS::StepFunctions::StateMachineAlias.RoutingConfigurationVersion.
	AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__Type = "AWS::StepFunctions::StateMachineAlias.RoutingConfigurationVersion"
)

var (
	// AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__PropertiesMap reports all the CloudFormation properties for AWS::StepFunctions::StateMachineAlias.RoutingConfigurationVersion.
	AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__PropertiesMap = struct {
		StateMachineVersionArn string
		Weight                 string
	}{
		StateMachineVersionArn: "StateMachineVersionArn",
		Weight:                 "Weight",
	}

	// AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__PropertiesSlice reports all the CloudFormation properties for AWS::StepFunctions::StateMachineAlias.RoutingConfigurationVersion.
	AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__PropertiesSlice = []string{
		AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__PropertiesMap.StateMachineVersionArn,
		AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__PropertiesMap.Weight,
	}
)

// AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion is a binding for AWS::StepFunctions::StateMachineAlias.RoutingConfigurationVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html
type AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion struct {
	// StateMachineVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html#cfn-stepfunctions-statemachinealias-routingconfigurationversion-statemachineversionarn
	StateMachineVersionArn cfz.Expression[string] `json:"StateMachineVersionArn,omitempty"`

	// Weight is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stepfunctions-statemachinealias-routingconfigurationversion.html#cfn-stepfunctions-statemachinealias-routingconfigurationversion-weight
	Weight cfz.Expression[int32] `json:"Weight,omitempty"`
}

// New__AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion initializes a new AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion.
func New__AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion() AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion {
	return AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion{}
}

// GetType returns the CloudFormation type.
func (AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion) GetType() string {
	return AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion__Type
}

// Set__StateMachineVersionArn updates property "StateMachineVersionArn".
func (t AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion) Set__StateMachineVersionArn(v cfz.Expression[string]) AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion {
	t.StateMachineVersionArn = v
	return t
}

// SetV__StateMachineVersionArn updates property "StateMachineVersionArn".
func (t AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion) SetV__StateMachineVersionArn(v string) AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion {
	t.StateMachineVersionArn = cfz.V(v)
	return t
}

// Set__Weight updates property "Weight".
func (t AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion) Set__Weight(v cfz.Expression[int32]) AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion {
	t.Weight = v
	return t
}

// SetV__Weight updates property "Weight".
func (t AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion) SetV__Weight(v int32) AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion {
	t.Weight = cfz.V(v)
	return t
}
