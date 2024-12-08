// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_events

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Events_Rule)(nil)
	_ cfz.Resource                   = (*AWS_Events_Rule)(nil)
)

const (
	// AWS_Events_Rule__Type is the CloudFormation type for AWS::Events::Rule.
	AWS_Events_Rule__Type = "AWS::Events::Rule"
)

var (
	// AWS_Events_Rule__AttributesMap reports all the CloudFormation attributes for AWS::Events::Rule.
	AWS_Events_Rule__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Events_Rule__AttributesSlice reports all the CloudFormation attributes for AWS::Events::Rule.
	AWS_Events_Rule__AttributesSlice = []string{
		AWS_Events_Rule__AttributesMap.Arn,
	}
)

var (
	// AWS_Events_Rule__PropertiesMap reports all the CloudFormation properties for AWS::Events::Rule.
	AWS_Events_Rule__PropertiesMap = struct {
		Description        string
		EventBusName       string
		EventPattern       string
		Name               string
		RoleArn            string
		ScheduleExpression string
		State              string
		Targets            string
	}{
		Description:        "Description",
		EventBusName:       "EventBusName",
		EventPattern:       "EventPattern",
		Name:               "Name",
		RoleArn:            "RoleArn",
		ScheduleExpression: "ScheduleExpression",
		State:              "State",
		Targets:            "Targets",
	}

	// AWS_Events_Rule__PropertiesSlice reports all the CloudFormation properties for AWS::Events::Rule.
	AWS_Events_Rule__PropertiesSlice = []string{
		AWS_Events_Rule__PropertiesMap.Description,
		AWS_Events_Rule__PropertiesMap.EventBusName,
		AWS_Events_Rule__PropertiesMap.EventPattern,
		AWS_Events_Rule__PropertiesMap.Name,
		AWS_Events_Rule__PropertiesMap.RoleArn,
		AWS_Events_Rule__PropertiesMap.ScheduleExpression,
		AWS_Events_Rule__PropertiesMap.State,
		AWS_Events_Rule__PropertiesMap.Targets,
	}
)

// AWS_Events_Rule is a binding for AWS::Events::Rule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html
type AWS_Events_Rule struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// EventBusName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventbusname
	EventBusName cfz.Expression[string] `json:"EventBusName,omitempty"`

	// EventPattern is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventpattern
	EventPattern cfz.Expression[json.RawMessage] `json:"EventPattern,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// ScheduleExpression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-scheduleexpression
	ScheduleExpression cfz.Expression[string] `json:"ScheduleExpression,omitempty"`

	// State is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-state
	State cfz.Expression[string] `json:"State,omitempty"`

	// Targets is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-targets
	Targets cfz.ExpressionSlice[AWS_Events_Rule_Target] `json:"Targets,omitempty"`
}

// New__AWS_Events_Rule initializes a new *AWS_Events_Rule.
func New__AWS_Events_Rule(logicalName string) *AWS_Events_Rule {
	return &AWS_Events_Rule{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Events_Rule) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Events_Rule) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Events_Rule) GetType() string {
	return AWS_Events_Rule__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Events_Rule) Set__LogicalName(v string) *AWS_Events_Rule {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Events_Rule) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Events_Rule {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Events_Rule) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Events_Rule {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Events_Rule) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Events_Rule {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Events_Rule) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Events_Rule {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Events_Rule) Set__RequestedOutputs(v []cfz.Output) *AWS_Events_Rule {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Events_Rule) Add__RequestedOutputs(v ...cfz.Output) *AWS_Events_Rule {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Events_Rule) Set__Description(v cfz.Expression[string]) *AWS_Events_Rule {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Events_Rule) SetV__Description(v string) *AWS_Events_Rule {
	t.Description = cfz.V(v)
	return t
}

// Set__EventBusName updates property "EventBusName".
func (t *AWS_Events_Rule) Set__EventBusName(v cfz.Expression[string]) *AWS_Events_Rule {
	t.EventBusName = v
	return t
}

// SetV__EventBusName updates property "EventBusName".
func (t *AWS_Events_Rule) SetV__EventBusName(v string) *AWS_Events_Rule {
	t.EventBusName = cfz.V(v)
	return t
}

// Set__EventPattern updates property "EventPattern".
func (t *AWS_Events_Rule) Set__EventPattern(v cfz.Expression[json.RawMessage]) *AWS_Events_Rule {
	t.EventPattern = v
	return t
}

// SetV__EventPattern updates property "EventPattern".
func (t *AWS_Events_Rule) SetV__EventPattern(v json.RawMessage) *AWS_Events_Rule {
	t.EventPattern = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Events_Rule) Set__Name(v cfz.Expression[string]) *AWS_Events_Rule {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Events_Rule) SetV__Name(v string) *AWS_Events_Rule {
	t.Name = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_Events_Rule) Set__RoleArn(v cfz.Expression[string]) *AWS_Events_Rule {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_Events_Rule) SetV__RoleArn(v string) *AWS_Events_Rule {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__ScheduleExpression updates property "ScheduleExpression".
func (t *AWS_Events_Rule) Set__ScheduleExpression(v cfz.Expression[string]) *AWS_Events_Rule {
	t.ScheduleExpression = v
	return t
}

// SetV__ScheduleExpression updates property "ScheduleExpression".
func (t *AWS_Events_Rule) SetV__ScheduleExpression(v string) *AWS_Events_Rule {
	t.ScheduleExpression = cfz.V(v)
	return t
}

// Set__State updates property "State".
func (t *AWS_Events_Rule) Set__State(v cfz.Expression[string]) *AWS_Events_Rule {
	t.State = v
	return t
}

// SetV__State updates property "State".
func (t *AWS_Events_Rule) SetV__State(v string) *AWS_Events_Rule {
	t.State = cfz.V(v)
	return t
}

// Set__Targets updates property "Targets".
func (t *AWS_Events_Rule) Set__Targets(v cfz.ExpressionSlice[AWS_Events_Rule_Target]) *AWS_Events_Rule {
	t.Targets = v
	return t
}

// SetS__Targets updates property "Targets".
func (t *AWS_Events_Rule) SetS__Targets(v ...cfz.Expression[AWS_Events_Rule_Target]) *AWS_Events_Rule {
	t.Targets = cfz.S(v...)
	return t
}

// SetSV__Targets updates property "Targets".
func (t *AWS_Events_Rule) SetSV__Targets(v ...AWS_Events_Rule_Target) *AWS_Events_Rule {
	t.Targets = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Events_Rule) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Events_Rule) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Events_Rule__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Events_Rule) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Events_Rule) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Events_Rule) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Events_Rule

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

func (t *AWS_Events_Rule) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
