// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_datazone

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_DataZone_DataSource_ScheduleConfiguration__Type is the CloudFormation type for AWS::DataZone::DataSource.ScheduleConfiguration.
	AWS_DataZone_DataSource_ScheduleConfiguration__Type = "AWS::DataZone::DataSource.ScheduleConfiguration"
)

var (
	// AWS_DataZone_DataSource_ScheduleConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::DataZone::DataSource.ScheduleConfiguration.
	AWS_DataZone_DataSource_ScheduleConfiguration__PropertiesMap = struct {
		Schedule string
		Timezone string
	}{
		Schedule: "Schedule",
		Timezone: "Timezone",
	}

	// AWS_DataZone_DataSource_ScheduleConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::DataZone::DataSource.ScheduleConfiguration.
	AWS_DataZone_DataSource_ScheduleConfiguration__PropertiesSlice = []string{
		AWS_DataZone_DataSource_ScheduleConfiguration__PropertiesMap.Schedule,
		AWS_DataZone_DataSource_ScheduleConfiguration__PropertiesMap.Timezone,
	}
)

// AWS_DataZone_DataSource_ScheduleConfiguration is a binding for AWS::DataZone::DataSource.ScheduleConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html
type AWS_DataZone_DataSource_ScheduleConfiguration struct {
	// Schedule is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html#cfn-datazone-datasource-scheduleconfiguration-schedule
	Schedule cfz.Expression[string] `json:"Schedule,omitempty"`

	// Timezone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-scheduleconfiguration.html#cfn-datazone-datasource-scheduleconfiguration-timezone
	Timezone cfz.Expression[string] `json:"Timezone,omitempty"`
}

// New__AWS_DataZone_DataSource_ScheduleConfiguration initializes a new AWS_DataZone_DataSource_ScheduleConfiguration.
func New__AWS_DataZone_DataSource_ScheduleConfiguration() AWS_DataZone_DataSource_ScheduleConfiguration {
	return AWS_DataZone_DataSource_ScheduleConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_DataZone_DataSource_ScheduleConfiguration) GetType() string {
	return AWS_DataZone_DataSource_ScheduleConfiguration__Type
}

// Set__Schedule updates property "Schedule".
func (t AWS_DataZone_DataSource_ScheduleConfiguration) Set__Schedule(v cfz.Expression[string]) AWS_DataZone_DataSource_ScheduleConfiguration {
	t.Schedule = v
	return t
}

// SetV__Schedule updates property "Schedule".
func (t AWS_DataZone_DataSource_ScheduleConfiguration) SetV__Schedule(v string) AWS_DataZone_DataSource_ScheduleConfiguration {
	t.Schedule = cfz.V(v)
	return t
}

// Set__Timezone updates property "Timezone".
func (t AWS_DataZone_DataSource_ScheduleConfiguration) Set__Timezone(v cfz.Expression[string]) AWS_DataZone_DataSource_ScheduleConfiguration {
	t.Timezone = v
	return t
}

// SetV__Timezone updates property "Timezone".
func (t AWS_DataZone_DataSource_ScheduleConfiguration) SetV__Timezone(v string) AWS_DataZone_DataSource_ScheduleConfiguration {
	t.Timezone = cfz.V(v)
	return t
}
