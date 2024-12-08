// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_stepfunctions

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_StepFunctions_StateMachineAlias)(nil)
	_ cfz.Resource                   = (*AWS_StepFunctions_StateMachineAlias)(nil)
)

const (
	// AWS_StepFunctions_StateMachineAlias__Type is the CloudFormation type for AWS::StepFunctions::StateMachineAlias.
	AWS_StepFunctions_StateMachineAlias__Type = "AWS::StepFunctions::StateMachineAlias"
)

var (
	// AWS_StepFunctions_StateMachineAlias__AttributesMap reports all the CloudFormation attributes for AWS::StepFunctions::StateMachineAlias.
	AWS_StepFunctions_StateMachineAlias__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_StepFunctions_StateMachineAlias__AttributesSlice reports all the CloudFormation attributes for AWS::StepFunctions::StateMachineAlias.
	AWS_StepFunctions_StateMachineAlias__AttributesSlice = []string{
		AWS_StepFunctions_StateMachineAlias__AttributesMap.Arn,
	}
)

var (
	// AWS_StepFunctions_StateMachineAlias__PropertiesMap reports all the CloudFormation properties for AWS::StepFunctions::StateMachineAlias.
	AWS_StepFunctions_StateMachineAlias__PropertiesMap = struct {
		DeploymentPreference string
		Description          string
		Name                 string
		RoutingConfiguration string
	}{
		DeploymentPreference: "DeploymentPreference",
		Description:          "Description",
		Name:                 "Name",
		RoutingConfiguration: "RoutingConfiguration",
	}

	// AWS_StepFunctions_StateMachineAlias__PropertiesSlice reports all the CloudFormation properties for AWS::StepFunctions::StateMachineAlias.
	AWS_StepFunctions_StateMachineAlias__PropertiesSlice = []string{
		AWS_StepFunctions_StateMachineAlias__PropertiesMap.DeploymentPreference,
		AWS_StepFunctions_StateMachineAlias__PropertiesMap.Description,
		AWS_StepFunctions_StateMachineAlias__PropertiesMap.Name,
		AWS_StepFunctions_StateMachineAlias__PropertiesMap.RoutingConfiguration,
	}
)

// AWS_StepFunctions_StateMachineAlias is a binding for AWS::StepFunctions::StateMachineAlias.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html
type AWS_StepFunctions_StateMachineAlias struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DeploymentPreference is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-deploymentpreference
	DeploymentPreference cfz.Expression[AWS_StepFunctions_StateMachineAlias_DeploymentPreference] `json:"DeploymentPreference,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RoutingConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachinealias.html#cfn-stepfunctions-statemachinealias-routingconfiguration
	RoutingConfiguration cfz.ExpressionSlice[AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion] `json:"RoutingConfiguration,omitempty"`
}

// New__AWS_StepFunctions_StateMachineAlias initializes a new *AWS_StepFunctions_StateMachineAlias.
func New__AWS_StepFunctions_StateMachineAlias(logicalName string) *AWS_StepFunctions_StateMachineAlias {
	return &AWS_StepFunctions_StateMachineAlias{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_StepFunctions_StateMachineAlias) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_StepFunctions_StateMachineAlias) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_StepFunctions_StateMachineAlias) GetType() string {
	return AWS_StepFunctions_StateMachineAlias__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_StepFunctions_StateMachineAlias) Set__LogicalName(v string) *AWS_StepFunctions_StateMachineAlias {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_StepFunctions_StateMachineAlias) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_StepFunctions_StateMachineAlias {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_StepFunctions_StateMachineAlias) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_StepFunctions_StateMachineAlias {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_StepFunctions_StateMachineAlias) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_StepFunctions_StateMachineAlias {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_StepFunctions_StateMachineAlias) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_StepFunctions_StateMachineAlias {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_StepFunctions_StateMachineAlias) Set__RequestedOutputs(v []cfz.Output) *AWS_StepFunctions_StateMachineAlias {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_StepFunctions_StateMachineAlias) Add__RequestedOutputs(v ...cfz.Output) *AWS_StepFunctions_StateMachineAlias {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DeploymentPreference updates property "DeploymentPreference".
func (t *AWS_StepFunctions_StateMachineAlias) Set__DeploymentPreference(v cfz.Expression[AWS_StepFunctions_StateMachineAlias_DeploymentPreference]) *AWS_StepFunctions_StateMachineAlias {
	t.DeploymentPreference = v
	return t
}

// SetV__DeploymentPreference updates property "DeploymentPreference".
func (t *AWS_StepFunctions_StateMachineAlias) SetV__DeploymentPreference(v AWS_StepFunctions_StateMachineAlias_DeploymentPreference) *AWS_StepFunctions_StateMachineAlias {
	t.DeploymentPreference = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_StepFunctions_StateMachineAlias) Set__Description(v cfz.Expression[string]) *AWS_StepFunctions_StateMachineAlias {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_StepFunctions_StateMachineAlias) SetV__Description(v string) *AWS_StepFunctions_StateMachineAlias {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_StepFunctions_StateMachineAlias) Set__Name(v cfz.Expression[string]) *AWS_StepFunctions_StateMachineAlias {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_StepFunctions_StateMachineAlias) SetV__Name(v string) *AWS_StepFunctions_StateMachineAlias {
	t.Name = cfz.V(v)
	return t
}

// Set__RoutingConfiguration updates property "RoutingConfiguration".
func (t *AWS_StepFunctions_StateMachineAlias) Set__RoutingConfiguration(v cfz.ExpressionSlice[AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion]) *AWS_StepFunctions_StateMachineAlias {
	t.RoutingConfiguration = v
	return t
}

// SetS__RoutingConfiguration updates property "RoutingConfiguration".
func (t *AWS_StepFunctions_StateMachineAlias) SetS__RoutingConfiguration(v ...cfz.Expression[AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion]) *AWS_StepFunctions_StateMachineAlias {
	t.RoutingConfiguration = cfz.S(v...)
	return t
}

// SetSV__RoutingConfiguration updates property "RoutingConfiguration".
func (t *AWS_StepFunctions_StateMachineAlias) SetSV__RoutingConfiguration(v ...AWS_StepFunctions_StateMachineAlias_RoutingConfigurationVersion) *AWS_StepFunctions_StateMachineAlias {
	t.RoutingConfiguration = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_StepFunctions_StateMachineAlias) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_StepFunctions_StateMachineAlias) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_StepFunctions_StateMachineAlias__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_StepFunctions_StateMachineAlias) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_StepFunctions_StateMachineAlias) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_StepFunctions_StateMachineAlias) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_StepFunctions_StateMachineAlias

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

func (t *AWS_StepFunctions_StateMachineAlias) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
