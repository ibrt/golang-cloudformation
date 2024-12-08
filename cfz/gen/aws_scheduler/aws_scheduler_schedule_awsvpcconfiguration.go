// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_scheduler

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Scheduler_Schedule_AwsVpcConfiguration__Type is the CloudFormation type for AWS::Scheduler::Schedule.AwsVpcConfiguration.
	AWS_Scheduler_Schedule_AwsVpcConfiguration__Type = "AWS::Scheduler::Schedule.AwsVpcConfiguration"
)

var (
	// AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Scheduler::Schedule.AwsVpcConfiguration.
	AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesMap = struct {
		AssignPublicIp string
		SecurityGroups string
		Subnets        string
	}{
		AssignPublicIp: "AssignPublicIp",
		SecurityGroups: "SecurityGroups",
		Subnets:        "Subnets",
	}

	// AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Scheduler::Schedule.AwsVpcConfiguration.
	AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesSlice = []string{
		AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesMap.AssignPublicIp,
		AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesMap.SecurityGroups,
		AWS_Scheduler_Schedule_AwsVpcConfiguration__PropertiesMap.Subnets,
	}
)

// AWS_Scheduler_Schedule_AwsVpcConfiguration is a binding for AWS::Scheduler::Schedule.AwsVpcConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-awsvpcconfiguration.html
type AWS_Scheduler_Schedule_AwsVpcConfiguration struct {
	// AssignPublicIp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-awsvpcconfiguration.html#cfn-scheduler-schedule-awsvpcconfiguration-assignpublicip
	AssignPublicIp cfz.Expression[string] `json:"AssignPublicIp,omitempty"`

	// SecurityGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-awsvpcconfiguration.html#cfn-scheduler-schedule-awsvpcconfiguration-securitygroups
	SecurityGroups cfz.ExpressionSlice[string] `json:"SecurityGroups,omitempty"`

	// Subnets is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-awsvpcconfiguration.html#cfn-scheduler-schedule-awsvpcconfiguration-subnets
	Subnets cfz.ExpressionSlice[string] `json:"Subnets,omitempty"`
}

// New__AWS_Scheduler_Schedule_AwsVpcConfiguration initializes a new AWS_Scheduler_Schedule_AwsVpcConfiguration.
func New__AWS_Scheduler_Schedule_AwsVpcConfiguration() AWS_Scheduler_Schedule_AwsVpcConfiguration {
	return AWS_Scheduler_Schedule_AwsVpcConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Scheduler_Schedule_AwsVpcConfiguration) GetType() string {
	return AWS_Scheduler_Schedule_AwsVpcConfiguration__Type
}

// Set__AssignPublicIp updates property "AssignPublicIp".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) Set__AssignPublicIp(v cfz.Expression[string]) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.AssignPublicIp = v
	return t
}

// SetV__AssignPublicIp updates property "AssignPublicIp".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) SetV__AssignPublicIp(v string) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.AssignPublicIp = cfz.V(v)
	return t
}

// Set__SecurityGroups updates property "SecurityGroups".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) Set__SecurityGroups(v cfz.ExpressionSlice[string]) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.SecurityGroups = v
	return t
}

// SetS__SecurityGroups updates property "SecurityGroups".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) SetS__SecurityGroups(v ...cfz.Expression[string]) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.SecurityGroups = cfz.S(v...)
	return t
}

// SetSV__SecurityGroups updates property "SecurityGroups".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) SetSV__SecurityGroups(v ...string) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.SecurityGroups = cfz.SV(v...)
	return t
}

// Set__Subnets updates property "Subnets".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) Set__Subnets(v cfz.ExpressionSlice[string]) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.Subnets = v
	return t
}

// SetS__Subnets updates property "Subnets".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) SetS__Subnets(v ...cfz.Expression[string]) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.Subnets = cfz.S(v...)
	return t
}

// SetSV__Subnets updates property "Subnets".
func (t AWS_Scheduler_Schedule_AwsVpcConfiguration) SetSV__Subnets(v ...string) AWS_Scheduler_Schedule_AwsVpcConfiguration {
	t.Subnets = cfz.SV(v...)
	return t
}
