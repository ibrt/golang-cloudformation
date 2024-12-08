// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_scheduler

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Scheduler_Schedule)(nil)
	_ cfz.Resource                   = (*AWS_Scheduler_Schedule)(nil)
)

const (
	// AWS_Scheduler_Schedule__Type is the CloudFormation type for AWS::Scheduler::Schedule.
	AWS_Scheduler_Schedule__Type = "AWS::Scheduler::Schedule"
)

var (
	// AWS_Scheduler_Schedule__AttributesMap reports all the CloudFormation attributes for AWS::Scheduler::Schedule.
	AWS_Scheduler_Schedule__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Scheduler_Schedule__AttributesSlice reports all the CloudFormation attributes for AWS::Scheduler::Schedule.
	AWS_Scheduler_Schedule__AttributesSlice = []string{
		AWS_Scheduler_Schedule__AttributesMap.Arn,
	}
)

var (
	// AWS_Scheduler_Schedule__PropertiesMap reports all the CloudFormation properties for AWS::Scheduler::Schedule.
	AWS_Scheduler_Schedule__PropertiesMap = struct {
		Description                string
		EndDate                    string
		FlexibleTimeWindow         string
		GroupName                  string
		KmsKeyArn                  string
		Name                       string
		ScheduleExpression         string
		ScheduleExpressionTimezone string
		StartDate                  string
		State                      string
		Target                     string
	}{
		Description:                "Description",
		EndDate:                    "EndDate",
		FlexibleTimeWindow:         "FlexibleTimeWindow",
		GroupName:                  "GroupName",
		KmsKeyArn:                  "KmsKeyArn",
		Name:                       "Name",
		ScheduleExpression:         "ScheduleExpression",
		ScheduleExpressionTimezone: "ScheduleExpressionTimezone",
		StartDate:                  "StartDate",
		State:                      "State",
		Target:                     "Target",
	}

	// AWS_Scheduler_Schedule__PropertiesSlice reports all the CloudFormation properties for AWS::Scheduler::Schedule.
	AWS_Scheduler_Schedule__PropertiesSlice = []string{
		AWS_Scheduler_Schedule__PropertiesMap.Description,
		AWS_Scheduler_Schedule__PropertiesMap.EndDate,
		AWS_Scheduler_Schedule__PropertiesMap.FlexibleTimeWindow,
		AWS_Scheduler_Schedule__PropertiesMap.GroupName,
		AWS_Scheduler_Schedule__PropertiesMap.KmsKeyArn,
		AWS_Scheduler_Schedule__PropertiesMap.Name,
		AWS_Scheduler_Schedule__PropertiesMap.ScheduleExpression,
		AWS_Scheduler_Schedule__PropertiesMap.ScheduleExpressionTimezone,
		AWS_Scheduler_Schedule__PropertiesMap.StartDate,
		AWS_Scheduler_Schedule__PropertiesMap.State,
		AWS_Scheduler_Schedule__PropertiesMap.Target,
	}
)

// AWS_Scheduler_Schedule is a binding for AWS::Scheduler::Schedule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html
type AWS_Scheduler_Schedule struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EndDate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-enddate
	EndDate cfz.Expression[string] `json:"EndDate,omitempty"`

	// FlexibleTimeWindow is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-flexibletimewindow
	FlexibleTimeWindow cfz.Expression[AWS_Scheduler_Schedule_FlexibleTimeWindow] `json:"FlexibleTimeWindow,omitempty"`

	// GroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-groupname
	GroupName cfz.Expression[string] `json:"GroupName,omitempty"`

	// KmsKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-kmskeyarn
	KmsKeyArn cfz.Expression[string] `json:"KmsKeyArn,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ScheduleExpression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-scheduleexpression
	ScheduleExpression cfz.Expression[string] `json:"ScheduleExpression,omitempty"`

	// ScheduleExpressionTimezone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-scheduleexpressiontimezone
	ScheduleExpressionTimezone cfz.Expression[string] `json:"ScheduleExpressionTimezone,omitempty"`

	// StartDate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-startdate
	StartDate cfz.Expression[string] `json:"StartDate,omitempty"`

	// State is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-state
	State cfz.Expression[string] `json:"State,omitempty"`

	// Target is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html#cfn-scheduler-schedule-target
	Target cfz.Expression[AWS_Scheduler_Schedule_Target] `json:"Target,omitempty"`
}

// New__AWS_Scheduler_Schedule initializes a new *AWS_Scheduler_Schedule.
func New__AWS_Scheduler_Schedule(logicalName string) *AWS_Scheduler_Schedule {
	return &AWS_Scheduler_Schedule{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Scheduler_Schedule) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Scheduler_Schedule) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Scheduler_Schedule) GetType() string {
	return AWS_Scheduler_Schedule__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Scheduler_Schedule) Set__LogicalName(v string) *AWS_Scheduler_Schedule {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Scheduler_Schedule) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Scheduler_Schedule {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Scheduler_Schedule) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Scheduler_Schedule {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Scheduler_Schedule) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Scheduler_Schedule {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Scheduler_Schedule) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Scheduler_Schedule {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Scheduler_Schedule) Set__RequestedOutputs(v []cfz.Output) *AWS_Scheduler_Schedule {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Scheduler_Schedule) Add__RequestedOutputs(v ...cfz.Output) *AWS_Scheduler_Schedule {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Scheduler_Schedule) Set__Description(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Scheduler_Schedule) SetV__Description(v string) *AWS_Scheduler_Schedule {
	t.Description = cfz.V(v)
	return t
}

// Set__EndDate updates property "EndDate".
func (t *AWS_Scheduler_Schedule) Set__EndDate(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.EndDate = v
	return t
}

// SetV__EndDate updates property "EndDate".
func (t *AWS_Scheduler_Schedule) SetV__EndDate(v string) *AWS_Scheduler_Schedule {
	t.EndDate = cfz.V(v)
	return t
}

// Set__FlexibleTimeWindow updates property "FlexibleTimeWindow".
func (t *AWS_Scheduler_Schedule) Set__FlexibleTimeWindow(v cfz.Expression[AWS_Scheduler_Schedule_FlexibleTimeWindow]) *AWS_Scheduler_Schedule {
	t.FlexibleTimeWindow = v
	return t
}

// SetV__FlexibleTimeWindow updates property "FlexibleTimeWindow".
func (t *AWS_Scheduler_Schedule) SetV__FlexibleTimeWindow(v AWS_Scheduler_Schedule_FlexibleTimeWindow) *AWS_Scheduler_Schedule {
	t.FlexibleTimeWindow = cfz.V(v)
	return t
}

// Set__GroupName updates property "GroupName".
func (t *AWS_Scheduler_Schedule) Set__GroupName(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.GroupName = v
	return t
}

// SetV__GroupName updates property "GroupName".
func (t *AWS_Scheduler_Schedule) SetV__GroupName(v string) *AWS_Scheduler_Schedule {
	t.GroupName = cfz.V(v)
	return t
}

// Set__KmsKeyArn updates property "KmsKeyArn".
func (t *AWS_Scheduler_Schedule) Set__KmsKeyArn(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.KmsKeyArn = v
	return t
}

// SetV__KmsKeyArn updates property "KmsKeyArn".
func (t *AWS_Scheduler_Schedule) SetV__KmsKeyArn(v string) *AWS_Scheduler_Schedule {
	t.KmsKeyArn = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Scheduler_Schedule) Set__Name(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Scheduler_Schedule) SetV__Name(v string) *AWS_Scheduler_Schedule {
	t.Name = cfz.V(v)
	return t
}

// Set__ScheduleExpression updates property "ScheduleExpression".
func (t *AWS_Scheduler_Schedule) Set__ScheduleExpression(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.ScheduleExpression = v
	return t
}

// SetV__ScheduleExpression updates property "ScheduleExpression".
func (t *AWS_Scheduler_Schedule) SetV__ScheduleExpression(v string) *AWS_Scheduler_Schedule {
	t.ScheduleExpression = cfz.V(v)
	return t
}

// Set__ScheduleExpressionTimezone updates property "ScheduleExpressionTimezone".
func (t *AWS_Scheduler_Schedule) Set__ScheduleExpressionTimezone(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.ScheduleExpressionTimezone = v
	return t
}

// SetV__ScheduleExpressionTimezone updates property "ScheduleExpressionTimezone".
func (t *AWS_Scheduler_Schedule) SetV__ScheduleExpressionTimezone(v string) *AWS_Scheduler_Schedule {
	t.ScheduleExpressionTimezone = cfz.V(v)
	return t
}

// Set__StartDate updates property "StartDate".
func (t *AWS_Scheduler_Schedule) Set__StartDate(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.StartDate = v
	return t
}

// SetV__StartDate updates property "StartDate".
func (t *AWS_Scheduler_Schedule) SetV__StartDate(v string) *AWS_Scheduler_Schedule {
	t.StartDate = cfz.V(v)
	return t
}

// Set__State updates property "State".
func (t *AWS_Scheduler_Schedule) Set__State(v cfz.Expression[string]) *AWS_Scheduler_Schedule {
	t.State = v
	return t
}

// SetV__State updates property "State".
func (t *AWS_Scheduler_Schedule) SetV__State(v string) *AWS_Scheduler_Schedule {
	t.State = cfz.V(v)
	return t
}

// Set__Target updates property "Target".
func (t *AWS_Scheduler_Schedule) Set__Target(v cfz.Expression[AWS_Scheduler_Schedule_Target]) *AWS_Scheduler_Schedule {
	t.Target = v
	return t
}

// SetV__Target updates property "Target".
func (t *AWS_Scheduler_Schedule) SetV__Target(v AWS_Scheduler_Schedule_Target) *AWS_Scheduler_Schedule {
	t.Target = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Scheduler_Schedule) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Scheduler_Schedule) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Scheduler_Schedule__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Scheduler_Schedule) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Scheduler_Schedule) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Scheduler_Schedule) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Scheduler_Schedule

	return json.Marshal(struct {
		Type                string                          `json:"Type"`
		DependsOn           []string                        `json:"DependsOn,omitempty"`
		DeletionPolicy      cfz.ResourceDeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Properties          *Properties                     `json:"Properties,omitempty"`
	}{
		Type:       t.GetType(),
		DependsOn:  t.getDependsOn(),
		Properties: (*Properties)(t),
	})
}

func (t *AWS_Scheduler_Schedule) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
