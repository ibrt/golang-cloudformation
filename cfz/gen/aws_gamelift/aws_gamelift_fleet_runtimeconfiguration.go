// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_GameLift_Fleet_RuntimeConfiguration__Type is the CloudFormation type for AWS::GameLift::Fleet.RuntimeConfiguration.
	AWS_GameLift_Fleet_RuntimeConfiguration__Type = "AWS::GameLift::Fleet.RuntimeConfiguration"
)

var (
	// AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::Fleet.RuntimeConfiguration.
	AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesMap = struct {
		GameSessionActivationTimeoutSeconds string
		MaxConcurrentGameSessionActivations string
		ServerProcesses                     string
	}{
		GameSessionActivationTimeoutSeconds: "GameSessionActivationTimeoutSeconds",
		MaxConcurrentGameSessionActivations: "MaxConcurrentGameSessionActivations",
		ServerProcesses:                     "ServerProcesses",
	}

	// AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::Fleet.RuntimeConfiguration.
	AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesSlice = []string{
		AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesMap.GameSessionActivationTimeoutSeconds,
		AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesMap.MaxConcurrentGameSessionActivations,
		AWS_GameLift_Fleet_RuntimeConfiguration__PropertiesMap.ServerProcesses,
	}
)

// AWS_GameLift_Fleet_RuntimeConfiguration is a binding for AWS::GameLift::Fleet.RuntimeConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html
type AWS_GameLift_Fleet_RuntimeConfiguration struct {
	// GameSessionActivationTimeoutSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-gamesessionactivationtimeoutseconds
	GameSessionActivationTimeoutSeconds cfz.Expression[int32] `json:"GameSessionActivationTimeoutSeconds,omitempty"`

	// MaxConcurrentGameSessionActivations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-maxconcurrentgamesessionactivations
	MaxConcurrentGameSessionActivations cfz.Expression[int32] `json:"MaxConcurrentGameSessionActivations,omitempty"`

	// ServerProcesses is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-serverprocesses
	ServerProcesses cfz.ExpressionSlice[AWS_GameLift_Fleet_ServerProcess] `json:"ServerProcesses,omitempty"`
}

// New__AWS_GameLift_Fleet_RuntimeConfiguration initializes a new AWS_GameLift_Fleet_RuntimeConfiguration.
func New__AWS_GameLift_Fleet_RuntimeConfiguration() AWS_GameLift_Fleet_RuntimeConfiguration {
	return AWS_GameLift_Fleet_RuntimeConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_GameLift_Fleet_RuntimeConfiguration) GetType() string {
	return AWS_GameLift_Fleet_RuntimeConfiguration__Type
}

// Set__GameSessionActivationTimeoutSeconds updates property "GameSessionActivationTimeoutSeconds".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) Set__GameSessionActivationTimeoutSeconds(v cfz.Expression[int32]) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.GameSessionActivationTimeoutSeconds = v
	return t
}

// SetV__GameSessionActivationTimeoutSeconds updates property "GameSessionActivationTimeoutSeconds".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) SetV__GameSessionActivationTimeoutSeconds(v int32) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.GameSessionActivationTimeoutSeconds = cfz.V(v)
	return t
}

// Set__MaxConcurrentGameSessionActivations updates property "MaxConcurrentGameSessionActivations".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) Set__MaxConcurrentGameSessionActivations(v cfz.Expression[int32]) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.MaxConcurrentGameSessionActivations = v
	return t
}

// SetV__MaxConcurrentGameSessionActivations updates property "MaxConcurrentGameSessionActivations".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) SetV__MaxConcurrentGameSessionActivations(v int32) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.MaxConcurrentGameSessionActivations = cfz.V(v)
	return t
}

// Set__ServerProcesses updates property "ServerProcesses".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) Set__ServerProcesses(v cfz.ExpressionSlice[AWS_GameLift_Fleet_ServerProcess]) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.ServerProcesses = v
	return t
}

// SetS__ServerProcesses updates property "ServerProcesses".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) SetS__ServerProcesses(v ...cfz.Expression[AWS_GameLift_Fleet_ServerProcess]) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.ServerProcesses = cfz.S(v...)
	return t
}

// SetSV__ServerProcesses updates property "ServerProcesses".
func (t AWS_GameLift_Fleet_RuntimeConfiguration) SetSV__ServerProcesses(v ...AWS_GameLift_Fleet_ServerProcess) AWS_GameLift_Fleet_RuntimeConfiguration {
	t.ServerProcesses = cfz.SV(v...)
	return t
}
