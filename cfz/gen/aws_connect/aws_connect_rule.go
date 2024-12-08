// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_connect

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Connect_Rule)(nil)
	_ cfz.Resource                   = (*AWS_Connect_Rule)(nil)
)

const (
	// AWS_Connect_Rule__Type is the CloudFormation type for AWS::Connect::Rule.
	AWS_Connect_Rule__Type = "AWS::Connect::Rule"
)

var (
	// AWS_Connect_Rule__AttributesMap reports all the CloudFormation attributes for AWS::Connect::Rule.
	AWS_Connect_Rule__AttributesMap = struct {
		RuleArn string
	}{
		RuleArn: "RuleArn",
	}

	// AWS_Connect_Rule__AttributesSlice reports all the CloudFormation attributes for AWS::Connect::Rule.
	AWS_Connect_Rule__AttributesSlice = []string{
		AWS_Connect_Rule__AttributesMap.RuleArn,
	}
)

var (
	// AWS_Connect_Rule__PropertiesMap reports all the CloudFormation properties for AWS::Connect::Rule.
	AWS_Connect_Rule__PropertiesMap = struct {
		Actions            string
		Function           string
		InstanceArn        string
		Name               string
		PublishStatus      string
		Tags               string
		TriggerEventSource string
	}{
		Actions:            "Actions",
		Function:           "Function",
		InstanceArn:        "InstanceArn",
		Name:               "Name",
		PublishStatus:      "PublishStatus",
		Tags:               "Tags",
		TriggerEventSource: "TriggerEventSource",
	}

	// AWS_Connect_Rule__PropertiesSlice reports all the CloudFormation properties for AWS::Connect::Rule.
	AWS_Connect_Rule__PropertiesSlice = []string{
		AWS_Connect_Rule__PropertiesMap.Actions,
		AWS_Connect_Rule__PropertiesMap.Function,
		AWS_Connect_Rule__PropertiesMap.InstanceArn,
		AWS_Connect_Rule__PropertiesMap.Name,
		AWS_Connect_Rule__PropertiesMap.PublishStatus,
		AWS_Connect_Rule__PropertiesMap.Tags,
		AWS_Connect_Rule__PropertiesMap.TriggerEventSource,
	}
)

// AWS_Connect_Rule is a binding for AWS::Connect::Rule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html
type AWS_Connect_Rule struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-actions
	Actions cfz.Expression[AWS_Connect_Rule_Actions] `json:"Actions,omitempty"`

	// Function is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-function
	Function cfz.Expression[string] `json:"Function,omitempty"`

	// InstanceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-instancearn
	InstanceArn cfz.Expression[string] `json:"InstanceArn,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// PublishStatus is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-publishstatus
	PublishStatus cfz.Expression[string] `json:"PublishStatus,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TriggerEventSource is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html#cfn-connect-rule-triggereventsource
	TriggerEventSource cfz.Expression[AWS_Connect_Rule_RuleTriggerEventSource] `json:"TriggerEventSource,omitempty"`
}

// New__AWS_Connect_Rule initializes a new *AWS_Connect_Rule.
func New__AWS_Connect_Rule(logicalName string) *AWS_Connect_Rule {
	return &AWS_Connect_Rule{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Connect_Rule) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Connect_Rule) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Connect_Rule) GetType() string {
	return AWS_Connect_Rule__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Connect_Rule) Set__LogicalName(v string) *AWS_Connect_Rule {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Connect_Rule) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Connect_Rule {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Connect_Rule) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Connect_Rule {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Connect_Rule) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Connect_Rule {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Connect_Rule) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Connect_Rule {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Connect_Rule) Set__RequestedOutputs(v []cfz.Output) *AWS_Connect_Rule {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Connect_Rule) Add__RequestedOutputs(v ...cfz.Output) *AWS_Connect_Rule {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Actions updates property "Actions".
func (t *AWS_Connect_Rule) Set__Actions(v cfz.Expression[AWS_Connect_Rule_Actions]) *AWS_Connect_Rule {
	t.Actions = v
	return t
}

// SetV__Actions updates property "Actions".
func (t *AWS_Connect_Rule) SetV__Actions(v AWS_Connect_Rule_Actions) *AWS_Connect_Rule {
	t.Actions = cfz.V(v)
	return t
}

// Set__Function updates property "Function".
func (t *AWS_Connect_Rule) Set__Function(v cfz.Expression[string]) *AWS_Connect_Rule {
	t.Function = v
	return t
}

// SetV__Function updates property "Function".
func (t *AWS_Connect_Rule) SetV__Function(v string) *AWS_Connect_Rule {
	t.Function = cfz.V(v)
	return t
}

// Set__InstanceArn updates property "InstanceArn".
func (t *AWS_Connect_Rule) Set__InstanceArn(v cfz.Expression[string]) *AWS_Connect_Rule {
	t.InstanceArn = v
	return t
}

// SetV__InstanceArn updates property "InstanceArn".
func (t *AWS_Connect_Rule) SetV__InstanceArn(v string) *AWS_Connect_Rule {
	t.InstanceArn = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Connect_Rule) Set__Name(v cfz.Expression[string]) *AWS_Connect_Rule {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Connect_Rule) SetV__Name(v string) *AWS_Connect_Rule {
	t.Name = cfz.V(v)
	return t
}

// Set__PublishStatus updates property "PublishStatus".
func (t *AWS_Connect_Rule) Set__PublishStatus(v cfz.Expression[string]) *AWS_Connect_Rule {
	t.PublishStatus = v
	return t
}

// SetV__PublishStatus updates property "PublishStatus".
func (t *AWS_Connect_Rule) SetV__PublishStatus(v string) *AWS_Connect_Rule {
	t.PublishStatus = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Connect_Rule) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Connect_Rule {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Connect_Rule) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Connect_Rule {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Connect_Rule) SetSV__Tags(v ...cfz.Tag) *AWS_Connect_Rule {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TriggerEventSource updates property "TriggerEventSource".
func (t *AWS_Connect_Rule) Set__TriggerEventSource(v cfz.Expression[AWS_Connect_Rule_RuleTriggerEventSource]) *AWS_Connect_Rule {
	t.TriggerEventSource = v
	return t
}

// SetV__TriggerEventSource updates property "TriggerEventSource".
func (t *AWS_Connect_Rule) SetV__TriggerEventSource(v AWS_Connect_Rule_RuleTriggerEventSource) *AWS_Connect_Rule {
	t.TriggerEventSource = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Connect_Rule) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__RuleArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RuleArn
func (t *AWS_Connect_Rule) GetAtt__RuleArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Connect_Rule__AttributesMap.RuleArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Connect_Rule) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RuleArn returns a conventionally configured output for an attribute of this resource.
// Attribute: RuleArn
func (t *AWS_Connect_Rule) GetConventionalOutputAtt__RuleArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRuleArn", t.GetAtt__RuleArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Connect_Rule) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Connect_Rule

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

func (t *AWS_Connect_Rule) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
