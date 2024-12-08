// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_scheduler

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Scheduler_ScheduleGroup)(nil)
	_ cfz.Resource                   = (*AWS_Scheduler_ScheduleGroup)(nil)
)

const (
	// AWS_Scheduler_ScheduleGroup__Type is the CloudFormation type for AWS::Scheduler::ScheduleGroup.
	AWS_Scheduler_ScheduleGroup__Type = "AWS::Scheduler::ScheduleGroup"
)

var (
	// AWS_Scheduler_ScheduleGroup__AttributesMap reports all the CloudFormation attributes for AWS::Scheduler::ScheduleGroup.
	AWS_Scheduler_ScheduleGroup__AttributesMap = struct {
		Arn                  string
		CreationDate         string
		LastModificationDate string
		State                string
	}{
		Arn:                  "Arn",
		CreationDate:         "CreationDate",
		LastModificationDate: "LastModificationDate",
		State:                "State",
	}

	// AWS_Scheduler_ScheduleGroup__AttributesSlice reports all the CloudFormation attributes for AWS::Scheduler::ScheduleGroup.
	AWS_Scheduler_ScheduleGroup__AttributesSlice = []string{
		AWS_Scheduler_ScheduleGroup__AttributesMap.Arn,
		AWS_Scheduler_ScheduleGroup__AttributesMap.CreationDate,
		AWS_Scheduler_ScheduleGroup__AttributesMap.LastModificationDate,
		AWS_Scheduler_ScheduleGroup__AttributesMap.State,
	}
)

var (
	// AWS_Scheduler_ScheduleGroup__PropertiesMap reports all the CloudFormation properties for AWS::Scheduler::ScheduleGroup.
	AWS_Scheduler_ScheduleGroup__PropertiesMap = struct {
		Name string
		Tags string
	}{
		Name: "Name",
		Tags: "Tags",
	}

	// AWS_Scheduler_ScheduleGroup__PropertiesSlice reports all the CloudFormation properties for AWS::Scheduler::ScheduleGroup.
	AWS_Scheduler_ScheduleGroup__PropertiesSlice = []string{
		AWS_Scheduler_ScheduleGroup__PropertiesMap.Name,
		AWS_Scheduler_ScheduleGroup__PropertiesMap.Tags,
	}
)

// AWS_Scheduler_ScheduleGroup is a binding for AWS::Scheduler::ScheduleGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedulegroup.html
type AWS_Scheduler_ScheduleGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedulegroup.html#cfn-scheduler-schedulegroup-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedulegroup.html#cfn-scheduler-schedulegroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Scheduler_ScheduleGroup initializes a new *AWS_Scheduler_ScheduleGroup.
func New__AWS_Scheduler_ScheduleGroup(logicalName string) *AWS_Scheduler_ScheduleGroup {
	return &AWS_Scheduler_ScheduleGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Scheduler_ScheduleGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Scheduler_ScheduleGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Scheduler_ScheduleGroup) GetType() string {
	return AWS_Scheduler_ScheduleGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Scheduler_ScheduleGroup) Set__LogicalName(v string) *AWS_Scheduler_ScheduleGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Scheduler_ScheduleGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Scheduler_ScheduleGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Scheduler_ScheduleGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Scheduler_ScheduleGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Scheduler_ScheduleGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Scheduler_ScheduleGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Scheduler_ScheduleGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Scheduler_ScheduleGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Scheduler_ScheduleGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_Scheduler_ScheduleGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Scheduler_ScheduleGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_Scheduler_ScheduleGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Scheduler_ScheduleGroup) Set__Name(v cfz.Expression[string]) *AWS_Scheduler_ScheduleGroup {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Scheduler_ScheduleGroup) SetV__Name(v string) *AWS_Scheduler_ScheduleGroup {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Scheduler_ScheduleGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Scheduler_ScheduleGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Scheduler_ScheduleGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Scheduler_ScheduleGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Scheduler_ScheduleGroup) SetSV__Tags(v ...cfz.Tag) *AWS_Scheduler_ScheduleGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Scheduler_ScheduleGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Scheduler_ScheduleGroup) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Scheduler_ScheduleGroup__AttributesMap.Arn))
}

// GetAtt__CreationDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationDate
func (t *AWS_Scheduler_ScheduleGroup) GetAtt__CreationDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Scheduler_ScheduleGroup__AttributesMap.CreationDate))
}

// GetAtt__LastModificationDate returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastModificationDate
func (t *AWS_Scheduler_ScheduleGroup) GetAtt__LastModificationDate() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Scheduler_ScheduleGroup__AttributesMap.LastModificationDate))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_Scheduler_ScheduleGroup) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Scheduler_ScheduleGroup__AttributesMap.State))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Scheduler_ScheduleGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Scheduler_ScheduleGroup) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationDate returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationDate
func (t *AWS_Scheduler_ScheduleGroup) GetConventionalOutputAtt__CreationDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationDate", t.GetAtt__CreationDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastModificationDate returns a conventionally configured output for an attribute of this resource.
// Attribute: LastModificationDate
func (t *AWS_Scheduler_ScheduleGroup) GetConventionalOutputAtt__LastModificationDate(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastModificationDate", t.GetAtt__LastModificationDate())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_Scheduler_ScheduleGroup) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Scheduler_ScheduleGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Scheduler_ScheduleGroup

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

func (t *AWS_Scheduler_ScheduleGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
