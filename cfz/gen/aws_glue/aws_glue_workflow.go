// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Glue_Workflow)(nil)
	_ cfz.Resource                   = (*AWS_Glue_Workflow)(nil)
)

const (
	// AWS_Glue_Workflow__Type is the CloudFormation type for AWS::Glue::Workflow.
	AWS_Glue_Workflow__Type = "AWS::Glue::Workflow"
)

var (
	// AWS_Glue_Workflow__PropertiesMap reports all the CloudFormation properties for AWS::Glue::Workflow.
	AWS_Glue_Workflow__PropertiesMap = struct {
		DefaultRunProperties string
		Description          string
		MaxConcurrentRuns    string
		Name                 string
		Tags                 string
	}{
		DefaultRunProperties: "DefaultRunProperties",
		Description:          "Description",
		MaxConcurrentRuns:    "MaxConcurrentRuns",
		Name:                 "Name",
		Tags:                 "Tags",
	}

	// AWS_Glue_Workflow__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::Workflow.
	AWS_Glue_Workflow__PropertiesSlice = []string{
		AWS_Glue_Workflow__PropertiesMap.DefaultRunProperties,
		AWS_Glue_Workflow__PropertiesMap.Description,
		AWS_Glue_Workflow__PropertiesMap.MaxConcurrentRuns,
		AWS_Glue_Workflow__PropertiesMap.Name,
		AWS_Glue_Workflow__PropertiesMap.Tags,
	}
)

// AWS_Glue_Workflow is a binding for AWS::Glue::Workflow.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html
type AWS_Glue_Workflow struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DefaultRunProperties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-defaultrunproperties
	DefaultRunProperties cfz.Expression[json.RawMessage] `json:"DefaultRunProperties,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// MaxConcurrentRuns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-maxconcurrentruns
	MaxConcurrentRuns cfz.Expression[int32] `json:"MaxConcurrentRuns,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-workflow.html#cfn-glue-workflow-tags
	Tags cfz.Expression[json.RawMessage] `json:"Tags,omitempty"`
}

// New__AWS_Glue_Workflow initializes a new *AWS_Glue_Workflow.
func New__AWS_Glue_Workflow(logicalName string) *AWS_Glue_Workflow {
	return &AWS_Glue_Workflow{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Glue_Workflow) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Glue_Workflow) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Glue_Workflow) GetType() string {
	return AWS_Glue_Workflow__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Glue_Workflow) Set__LogicalName(v string) *AWS_Glue_Workflow {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Glue_Workflow) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Glue_Workflow {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Glue_Workflow) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Glue_Workflow {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Glue_Workflow) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Glue_Workflow {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Glue_Workflow) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Glue_Workflow {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Glue_Workflow) Set__RequestedOutputs(v []cfz.Output) *AWS_Glue_Workflow {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Glue_Workflow) Add__RequestedOutputs(v ...cfz.Output) *AWS_Glue_Workflow {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DefaultRunProperties updates property "DefaultRunProperties".
func (t *AWS_Glue_Workflow) Set__DefaultRunProperties(v cfz.Expression[json.RawMessage]) *AWS_Glue_Workflow {
	t.DefaultRunProperties = v
	return t
}

// SetV__DefaultRunProperties updates property "DefaultRunProperties".
func (t *AWS_Glue_Workflow) SetV__DefaultRunProperties(v json.RawMessage) *AWS_Glue_Workflow {
	t.DefaultRunProperties = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Glue_Workflow) Set__Description(v cfz.Expression[string]) *AWS_Glue_Workflow {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Glue_Workflow) SetV__Description(v string) *AWS_Glue_Workflow {
	t.Description = cfz.V(v)
	return t
}

// Set__MaxConcurrentRuns updates property "MaxConcurrentRuns".
func (t *AWS_Glue_Workflow) Set__MaxConcurrentRuns(v cfz.Expression[int32]) *AWS_Glue_Workflow {
	t.MaxConcurrentRuns = v
	return t
}

// SetV__MaxConcurrentRuns updates property "MaxConcurrentRuns".
func (t *AWS_Glue_Workflow) SetV__MaxConcurrentRuns(v int32) *AWS_Glue_Workflow {
	t.MaxConcurrentRuns = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Glue_Workflow) Set__Name(v cfz.Expression[string]) *AWS_Glue_Workflow {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Glue_Workflow) SetV__Name(v string) *AWS_Glue_Workflow {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Glue_Workflow) Set__Tags(v cfz.Expression[json.RawMessage]) *AWS_Glue_Workflow {
	t.Tags = v
	return t
}

// SetV__Tags updates property "Tags".
func (t *AWS_Glue_Workflow) SetV__Tags(v json.RawMessage) *AWS_Glue_Workflow {
	t.Tags = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Glue_Workflow) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Glue_Workflow) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Glue_Workflow) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Glue_Workflow

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

func (t *AWS_Glue_Workflow) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
