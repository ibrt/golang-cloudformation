// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ecs

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ECS_Service_DeploymentAlarms__Type is the CloudFormation type for AWS::ECS::Service.DeploymentAlarms.
	AWS_ECS_Service_DeploymentAlarms__Type = "AWS::ECS::Service.DeploymentAlarms"
)

var (
	// AWS_ECS_Service_DeploymentAlarms__PropertiesMap reports all the CloudFormation properties for AWS::ECS::Service.DeploymentAlarms.
	AWS_ECS_Service_DeploymentAlarms__PropertiesMap = struct {
		AlarmNames string
		Enable     string
		Rollback   string
	}{
		AlarmNames: "AlarmNames",
		Enable:     "Enable",
		Rollback:   "Rollback",
	}

	// AWS_ECS_Service_DeploymentAlarms__PropertiesSlice reports all the CloudFormation properties for AWS::ECS::Service.DeploymentAlarms.
	AWS_ECS_Service_DeploymentAlarms__PropertiesSlice = []string{
		AWS_ECS_Service_DeploymentAlarms__PropertiesMap.AlarmNames,
		AWS_ECS_Service_DeploymentAlarms__PropertiesMap.Enable,
		AWS_ECS_Service_DeploymentAlarms__PropertiesMap.Rollback,
	}
)

// AWS_ECS_Service_DeploymentAlarms is a binding for AWS::ECS::Service.DeploymentAlarms.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html
type AWS_ECS_Service_DeploymentAlarms struct {
	// AlarmNames is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-alarmnames
	AlarmNames cfz.ExpressionSlice[string] `json:"AlarmNames,omitempty"`

	// Enable is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-enable
	Enable cfz.Expression[bool] `json:"Enable,omitempty"`

	// Rollback is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentalarms.html#cfn-ecs-service-deploymentalarms-rollback
	Rollback cfz.Expression[bool] `json:"Rollback,omitempty"`
}

// New__AWS_ECS_Service_DeploymentAlarms initializes a new AWS_ECS_Service_DeploymentAlarms.
func New__AWS_ECS_Service_DeploymentAlarms() AWS_ECS_Service_DeploymentAlarms {
	return AWS_ECS_Service_DeploymentAlarms{}
}

// GetType returns the CloudFormation type.
func (AWS_ECS_Service_DeploymentAlarms) GetType() string {
	return AWS_ECS_Service_DeploymentAlarms__Type
}

// Set__AlarmNames updates property "AlarmNames".
func (t AWS_ECS_Service_DeploymentAlarms) Set__AlarmNames(v cfz.ExpressionSlice[string]) AWS_ECS_Service_DeploymentAlarms {
	t.AlarmNames = v
	return t
}

// SetS__AlarmNames updates property "AlarmNames".
func (t AWS_ECS_Service_DeploymentAlarms) SetS__AlarmNames(v ...cfz.Expression[string]) AWS_ECS_Service_DeploymentAlarms {
	t.AlarmNames = cfz.S(v...)
	return t
}

// SetSV__AlarmNames updates property "AlarmNames".
func (t AWS_ECS_Service_DeploymentAlarms) SetSV__AlarmNames(v ...string) AWS_ECS_Service_DeploymentAlarms {
	t.AlarmNames = cfz.SV(v...)
	return t
}

// Set__Enable updates property "Enable".
func (t AWS_ECS_Service_DeploymentAlarms) Set__Enable(v cfz.Expression[bool]) AWS_ECS_Service_DeploymentAlarms {
	t.Enable = v
	return t
}

// SetV__Enable updates property "Enable".
func (t AWS_ECS_Service_DeploymentAlarms) SetV__Enable(v bool) AWS_ECS_Service_DeploymentAlarms {
	t.Enable = cfz.V(v)
	return t
}

// Set__Rollback updates property "Rollback".
func (t AWS_ECS_Service_DeploymentAlarms) Set__Rollback(v cfz.Expression[bool]) AWS_ECS_Service_DeploymentAlarms {
	t.Rollback = v
	return t
}

// SetV__Rollback updates property "Rollback".
func (t AWS_ECS_Service_DeploymentAlarms) SetV__Rollback(v bool) AWS_ECS_Service_DeploymentAlarms {
	t.Rollback = cfz.V(v)
	return t
}
