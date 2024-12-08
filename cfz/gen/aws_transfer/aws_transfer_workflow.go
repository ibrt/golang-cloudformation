// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_transfer

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Transfer_Workflow)(nil)
	_ cfz.Resource                   = (*AWS_Transfer_Workflow)(nil)
)

const (
	// AWS_Transfer_Workflow__Type is the CloudFormation type for AWS::Transfer::Workflow.
	AWS_Transfer_Workflow__Type = "AWS::Transfer::Workflow"
)

var (
	// AWS_Transfer_Workflow__AttributesMap reports all the CloudFormation attributes for AWS::Transfer::Workflow.
	AWS_Transfer_Workflow__AttributesMap = struct {
		Arn        string
		WorkflowId string
	}{
		Arn:        "Arn",
		WorkflowId: "WorkflowId",
	}

	// AWS_Transfer_Workflow__AttributesSlice reports all the CloudFormation attributes for AWS::Transfer::Workflow.
	AWS_Transfer_Workflow__AttributesSlice = []string{
		AWS_Transfer_Workflow__AttributesMap.Arn,
		AWS_Transfer_Workflow__AttributesMap.WorkflowId,
	}
)

var (
	// AWS_Transfer_Workflow__PropertiesMap reports all the CloudFormation properties for AWS::Transfer::Workflow.
	AWS_Transfer_Workflow__PropertiesMap = struct {
		Description      string
		OnExceptionSteps string
		Steps            string
		Tags             string
	}{
		Description:      "Description",
		OnExceptionSteps: "OnExceptionSteps",
		Steps:            "Steps",
		Tags:             "Tags",
	}

	// AWS_Transfer_Workflow__PropertiesSlice reports all the CloudFormation properties for AWS::Transfer::Workflow.
	AWS_Transfer_Workflow__PropertiesSlice = []string{
		AWS_Transfer_Workflow__PropertiesMap.Description,
		AWS_Transfer_Workflow__PropertiesMap.OnExceptionSteps,
		AWS_Transfer_Workflow__PropertiesMap.Steps,
		AWS_Transfer_Workflow__PropertiesMap.Tags,
	}
)

// AWS_Transfer_Workflow is a binding for AWS::Transfer::Workflow.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html
type AWS_Transfer_Workflow struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// OnExceptionSteps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-onexceptionsteps
	OnExceptionSteps cfz.ExpressionSlice[AWS_Transfer_Workflow_WorkflowStep] `json:"OnExceptionSteps,omitempty"`

	// Steps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-steps
	Steps cfz.ExpressionSlice[AWS_Transfer_Workflow_WorkflowStep] `json:"Steps,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html#cfn-transfer-workflow-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Transfer_Workflow initializes a new *AWS_Transfer_Workflow.
func New__AWS_Transfer_Workflow(logicalName string) *AWS_Transfer_Workflow {
	return &AWS_Transfer_Workflow{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Transfer_Workflow) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Transfer_Workflow) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Transfer_Workflow) GetType() string {
	return AWS_Transfer_Workflow__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Transfer_Workflow) Set__LogicalName(v string) *AWS_Transfer_Workflow {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Transfer_Workflow) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Transfer_Workflow {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Transfer_Workflow) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Transfer_Workflow {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Transfer_Workflow) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Transfer_Workflow {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Transfer_Workflow) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Transfer_Workflow {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Transfer_Workflow) Set__RequestedOutputs(v []cfz.Output) *AWS_Transfer_Workflow {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Transfer_Workflow) Add__RequestedOutputs(v ...cfz.Output) *AWS_Transfer_Workflow {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Transfer_Workflow) Set__Description(v cfz.Expression[string]) *AWS_Transfer_Workflow {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Transfer_Workflow) SetV__Description(v string) *AWS_Transfer_Workflow {
	t.Description = cfz.V(v)
	return t
}

// Set__OnExceptionSteps updates property "OnExceptionSteps".
func (t *AWS_Transfer_Workflow) Set__OnExceptionSteps(v cfz.ExpressionSlice[AWS_Transfer_Workflow_WorkflowStep]) *AWS_Transfer_Workflow {
	t.OnExceptionSteps = v
	return t
}

// SetS__OnExceptionSteps updates property "OnExceptionSteps".
func (t *AWS_Transfer_Workflow) SetS__OnExceptionSteps(v ...cfz.Expression[AWS_Transfer_Workflow_WorkflowStep]) *AWS_Transfer_Workflow {
	t.OnExceptionSteps = cfz.S(v...)
	return t
}

// SetSV__OnExceptionSteps updates property "OnExceptionSteps".
func (t *AWS_Transfer_Workflow) SetSV__OnExceptionSteps(v ...AWS_Transfer_Workflow_WorkflowStep) *AWS_Transfer_Workflow {
	t.OnExceptionSteps = cfz.SV(v...)
	return t
}

// Set__Steps updates property "Steps".
func (t *AWS_Transfer_Workflow) Set__Steps(v cfz.ExpressionSlice[AWS_Transfer_Workflow_WorkflowStep]) *AWS_Transfer_Workflow {
	t.Steps = v
	return t
}

// SetS__Steps updates property "Steps".
func (t *AWS_Transfer_Workflow) SetS__Steps(v ...cfz.Expression[AWS_Transfer_Workflow_WorkflowStep]) *AWS_Transfer_Workflow {
	t.Steps = cfz.S(v...)
	return t
}

// SetSV__Steps updates property "Steps".
func (t *AWS_Transfer_Workflow) SetSV__Steps(v ...AWS_Transfer_Workflow_WorkflowStep) *AWS_Transfer_Workflow {
	t.Steps = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Transfer_Workflow) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Transfer_Workflow {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Transfer_Workflow) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Transfer_Workflow {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Transfer_Workflow) SetSV__Tags(v ...cfz.Tag) *AWS_Transfer_Workflow {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Transfer_Workflow) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Transfer_Workflow) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Transfer_Workflow__AttributesMap.Arn))
}

// GetAtt__WorkflowId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: WorkflowId
func (t *AWS_Transfer_Workflow) GetAtt__WorkflowId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Transfer_Workflow__AttributesMap.WorkflowId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Transfer_Workflow) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Transfer_Workflow) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__WorkflowId returns a conventionally configured output for an attribute of this resource.
// Attribute: WorkflowId
func (t *AWS_Transfer_Workflow) GetConventionalOutputAtt__WorkflowId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttWorkflowId", t.GetAtt__WorkflowId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Transfer_Workflow) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Transfer_Workflow

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

func (t *AWS_Transfer_Workflow) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
