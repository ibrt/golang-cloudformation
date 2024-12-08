// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Glue_Trigger)(nil)
	_ cfz.Resource                   = (*AWS_Glue_Trigger)(nil)
)

const (
	// AWS_Glue_Trigger__Type is the CloudFormation type for AWS::Glue::Trigger.
	AWS_Glue_Trigger__Type = "AWS::Glue::Trigger"
)

var (
	// AWS_Glue_Trigger__PropertiesMap reports all the CloudFormation properties for AWS::Glue::Trigger.
	AWS_Glue_Trigger__PropertiesMap = struct {
		Actions                string
		Description            string
		EventBatchingCondition string
		Name                   string
		Predicate              string
		Schedule               string
		StartOnCreation        string
		Tags                   string
		Type                   string
		WorkflowName           string
	}{
		Actions:                "Actions",
		Description:            "Description",
		EventBatchingCondition: "EventBatchingCondition",
		Name:                   "Name",
		Predicate:              "Predicate",
		Schedule:               "Schedule",
		StartOnCreation:        "StartOnCreation",
		Tags:                   "Tags",
		Type:                   "Type",
		WorkflowName:           "WorkflowName",
	}

	// AWS_Glue_Trigger__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::Trigger.
	AWS_Glue_Trigger__PropertiesSlice = []string{
		AWS_Glue_Trigger__PropertiesMap.Actions,
		AWS_Glue_Trigger__PropertiesMap.Description,
		AWS_Glue_Trigger__PropertiesMap.EventBatchingCondition,
		AWS_Glue_Trigger__PropertiesMap.Name,
		AWS_Glue_Trigger__PropertiesMap.Predicate,
		AWS_Glue_Trigger__PropertiesMap.Schedule,
		AWS_Glue_Trigger__PropertiesMap.StartOnCreation,
		AWS_Glue_Trigger__PropertiesMap.Tags,
		AWS_Glue_Trigger__PropertiesMap.Type,
		AWS_Glue_Trigger__PropertiesMap.WorkflowName,
	}
)

// AWS_Glue_Trigger is a binding for AWS::Glue::Trigger.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html
type AWS_Glue_Trigger struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Actions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-actions
	Actions cfz.ExpressionSlice[AWS_Glue_Trigger_Action] `json:"Actions,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EventBatchingCondition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-eventbatchingcondition
	EventBatchingCondition cfz.Expression[AWS_Glue_Trigger_EventBatchingCondition] `json:"EventBatchingCondition,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Predicate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-predicate
	Predicate cfz.Expression[AWS_Glue_Trigger_Predicate] `json:"Predicate,omitempty"`

	// Schedule is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-schedule
	Schedule cfz.Expression[string] `json:"Schedule,omitempty"`

	// StartOnCreation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-startoncreation
	StartOnCreation cfz.Expression[bool] `json:"StartOnCreation,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-tags
	Tags cfz.Expression[json.RawMessage] `json:"Tags,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-type
	Type cfz.Expression[string] `json:"Type,omitempty"`

	// WorkflowName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html#cfn-glue-trigger-workflowname
	WorkflowName cfz.Expression[string] `json:"WorkflowName,omitempty"`
}

// New__AWS_Glue_Trigger initializes a new *AWS_Glue_Trigger.
func New__AWS_Glue_Trigger(logicalName string) *AWS_Glue_Trigger {
	return &AWS_Glue_Trigger{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Glue_Trigger) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Glue_Trigger) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Glue_Trigger) GetType() string {
	return AWS_Glue_Trigger__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Glue_Trigger) Set__LogicalName(v string) *AWS_Glue_Trigger {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Glue_Trigger) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Glue_Trigger {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Glue_Trigger) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Glue_Trigger {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Glue_Trigger) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Glue_Trigger {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Glue_Trigger) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Glue_Trigger {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Glue_Trigger) Set__RequestedOutputs(v []cfz.Output) *AWS_Glue_Trigger {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Glue_Trigger) Add__RequestedOutputs(v ...cfz.Output) *AWS_Glue_Trigger {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Actions updates property "Actions".
func (t *AWS_Glue_Trigger) Set__Actions(v cfz.ExpressionSlice[AWS_Glue_Trigger_Action]) *AWS_Glue_Trigger {
	t.Actions = v
	return t
}

// SetS__Actions updates property "Actions".
func (t *AWS_Glue_Trigger) SetS__Actions(v ...cfz.Expression[AWS_Glue_Trigger_Action]) *AWS_Glue_Trigger {
	t.Actions = cfz.S(v...)
	return t
}

// SetSV__Actions updates property "Actions".
func (t *AWS_Glue_Trigger) SetSV__Actions(v ...AWS_Glue_Trigger_Action) *AWS_Glue_Trigger {
	t.Actions = cfz.SV(v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Glue_Trigger) Set__Description(v cfz.Expression[string]) *AWS_Glue_Trigger {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Glue_Trigger) SetV__Description(v string) *AWS_Glue_Trigger {
	t.Description = cfz.V(v)
	return t
}

// Set__EventBatchingCondition updates property "EventBatchingCondition".
func (t *AWS_Glue_Trigger) Set__EventBatchingCondition(v cfz.Expression[AWS_Glue_Trigger_EventBatchingCondition]) *AWS_Glue_Trigger {
	t.EventBatchingCondition = v
	return t
}

// SetV__EventBatchingCondition updates property "EventBatchingCondition".
func (t *AWS_Glue_Trigger) SetV__EventBatchingCondition(v AWS_Glue_Trigger_EventBatchingCondition) *AWS_Glue_Trigger {
	t.EventBatchingCondition = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Glue_Trigger) Set__Name(v cfz.Expression[string]) *AWS_Glue_Trigger {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Glue_Trigger) SetV__Name(v string) *AWS_Glue_Trigger {
	t.Name = cfz.V(v)
	return t
}

// Set__Predicate updates property "Predicate".
func (t *AWS_Glue_Trigger) Set__Predicate(v cfz.Expression[AWS_Glue_Trigger_Predicate]) *AWS_Glue_Trigger {
	t.Predicate = v
	return t
}

// SetV__Predicate updates property "Predicate".
func (t *AWS_Glue_Trigger) SetV__Predicate(v AWS_Glue_Trigger_Predicate) *AWS_Glue_Trigger {
	t.Predicate = cfz.V(v)
	return t
}

// Set__Schedule updates property "Schedule".
func (t *AWS_Glue_Trigger) Set__Schedule(v cfz.Expression[string]) *AWS_Glue_Trigger {
	t.Schedule = v
	return t
}

// SetV__Schedule updates property "Schedule".
func (t *AWS_Glue_Trigger) SetV__Schedule(v string) *AWS_Glue_Trigger {
	t.Schedule = cfz.V(v)
	return t
}

// Set__StartOnCreation updates property "StartOnCreation".
func (t *AWS_Glue_Trigger) Set__StartOnCreation(v cfz.Expression[bool]) *AWS_Glue_Trigger {
	t.StartOnCreation = v
	return t
}

// SetV__StartOnCreation updates property "StartOnCreation".
func (t *AWS_Glue_Trigger) SetV__StartOnCreation(v bool) *AWS_Glue_Trigger {
	t.StartOnCreation = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Glue_Trigger) Set__Tags(v cfz.Expression[json.RawMessage]) *AWS_Glue_Trigger {
	t.Tags = v
	return t
}

// SetV__Tags updates property "Tags".
func (t *AWS_Glue_Trigger) SetV__Tags(v json.RawMessage) *AWS_Glue_Trigger {
	t.Tags = cfz.V(v)
	return t
}

// Set__Type updates property "Type".
func (t *AWS_Glue_Trigger) Set__Type(v cfz.Expression[string]) *AWS_Glue_Trigger {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t *AWS_Glue_Trigger) SetV__Type(v string) *AWS_Glue_Trigger {
	t.Type = cfz.V(v)
	return t
}

// Set__WorkflowName updates property "WorkflowName".
func (t *AWS_Glue_Trigger) Set__WorkflowName(v cfz.Expression[string]) *AWS_Glue_Trigger {
	t.WorkflowName = v
	return t
}

// SetV__WorkflowName updates property "WorkflowName".
func (t *AWS_Glue_Trigger) SetV__WorkflowName(v string) *AWS_Glue_Trigger {
	t.WorkflowName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Glue_Trigger) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Glue_Trigger) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Glue_Trigger) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Glue_Trigger

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

func (t *AWS_Glue_Trigger) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
