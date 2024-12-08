// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codepipeline

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_CodePipeline_Pipeline)(nil)
	_ cfz.Resource                   = (*AWS_CodePipeline_Pipeline)(nil)
)

const (
	// AWS_CodePipeline_Pipeline__Type is the CloudFormation type for AWS::CodePipeline::Pipeline.
	AWS_CodePipeline_Pipeline__Type = "AWS::CodePipeline::Pipeline"
)

var (
	// AWS_CodePipeline_Pipeline__AttributesMap reports all the CloudFormation attributes for AWS::CodePipeline::Pipeline.
	AWS_CodePipeline_Pipeline__AttributesMap = struct {
		Version string
	}{
		Version: "Version",
	}

	// AWS_CodePipeline_Pipeline__AttributesSlice reports all the CloudFormation attributes for AWS::CodePipeline::Pipeline.
	AWS_CodePipeline_Pipeline__AttributesSlice = []string{
		AWS_CodePipeline_Pipeline__AttributesMap.Version,
	}
)

var (
	// AWS_CodePipeline_Pipeline__PropertiesMap reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.
	AWS_CodePipeline_Pipeline__PropertiesMap = struct {
		ArtifactStore                  string
		ArtifactStores                 string
		DisableInboundStageTransitions string
		ExecutionMode                  string
		Name                           string
		PipelineType                   string
		RestartExecutionOnUpdate       string
		RoleArn                        string
		Stages                         string
		Tags                           string
		Triggers                       string
		Variables                      string
	}{
		ArtifactStore:                  "ArtifactStore",
		ArtifactStores:                 "ArtifactStores",
		DisableInboundStageTransitions: "DisableInboundStageTransitions",
		ExecutionMode:                  "ExecutionMode",
		Name:                           "Name",
		PipelineType:                   "PipelineType",
		RestartExecutionOnUpdate:       "RestartExecutionOnUpdate",
		RoleArn:                        "RoleArn",
		Stages:                         "Stages",
		Tags:                           "Tags",
		Triggers:                       "Triggers",
		Variables:                      "Variables",
	}

	// AWS_CodePipeline_Pipeline__PropertiesSlice reports all the CloudFormation properties for AWS::CodePipeline::Pipeline.
	AWS_CodePipeline_Pipeline__PropertiesSlice = []string{
		AWS_CodePipeline_Pipeline__PropertiesMap.ArtifactStore,
		AWS_CodePipeline_Pipeline__PropertiesMap.ArtifactStores,
		AWS_CodePipeline_Pipeline__PropertiesMap.DisableInboundStageTransitions,
		AWS_CodePipeline_Pipeline__PropertiesMap.ExecutionMode,
		AWS_CodePipeline_Pipeline__PropertiesMap.Name,
		AWS_CodePipeline_Pipeline__PropertiesMap.PipelineType,
		AWS_CodePipeline_Pipeline__PropertiesMap.RestartExecutionOnUpdate,
		AWS_CodePipeline_Pipeline__PropertiesMap.RoleArn,
		AWS_CodePipeline_Pipeline__PropertiesMap.Stages,
		AWS_CodePipeline_Pipeline__PropertiesMap.Tags,
		AWS_CodePipeline_Pipeline__PropertiesMap.Triggers,
		AWS_CodePipeline_Pipeline__PropertiesMap.Variables,
	}
)

// AWS_CodePipeline_Pipeline is a binding for AWS::CodePipeline::Pipeline.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
type AWS_CodePipeline_Pipeline struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ArtifactStore is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstore
	ArtifactStore cfz.Expression[AWS_CodePipeline_Pipeline_ArtifactStore] `json:"ArtifactStore,omitempty"`

	// ArtifactStores is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstores
	ArtifactStores cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_ArtifactStoreMap] `json:"ArtifactStores,omitempty"`

	// DisableInboundStageTransitions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-disableinboundstagetransitions
	DisableInboundStageTransitions cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_StageTransition] `json:"DisableInboundStageTransitions,omitempty"`

	// ExecutionMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-executionmode
	ExecutionMode cfz.Expression[string] `json:"ExecutionMode,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// PipelineType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-pipelinetype
	PipelineType cfz.Expression[string] `json:"PipelineType,omitempty"`

	// RestartExecutionOnUpdate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-restartexecutiononupdate
	RestartExecutionOnUpdate cfz.Expression[bool] `json:"RestartExecutionOnUpdate,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// Stages is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-stages
	Stages cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_StageDeclaration] `json:"Stages,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// Triggers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-triggers
	Triggers cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration] `json:"Triggers,omitempty"`

	// Variables is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-variables
	Variables cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_VariableDeclaration] `json:"Variables,omitempty"`
}

// New__AWS_CodePipeline_Pipeline initializes a new *AWS_CodePipeline_Pipeline.
func New__AWS_CodePipeline_Pipeline(logicalName string) *AWS_CodePipeline_Pipeline {
	return &AWS_CodePipeline_Pipeline{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_CodePipeline_Pipeline) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_CodePipeline_Pipeline) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_CodePipeline_Pipeline) GetType() string {
	return AWS_CodePipeline_Pipeline__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_CodePipeline_Pipeline) Set__LogicalName(v string) *AWS_CodePipeline_Pipeline {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_CodePipeline_Pipeline) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_CodePipeline_Pipeline {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_CodePipeline_Pipeline) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_CodePipeline_Pipeline {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_CodePipeline_Pipeline) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_CodePipeline_Pipeline {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_CodePipeline_Pipeline) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_CodePipeline_Pipeline {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_CodePipeline_Pipeline) Set__RequestedOutputs(v []cfz.Output) *AWS_CodePipeline_Pipeline {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_CodePipeline_Pipeline) Add__RequestedOutputs(v ...cfz.Output) *AWS_CodePipeline_Pipeline {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ArtifactStore updates property "ArtifactStore".
func (t *AWS_CodePipeline_Pipeline) Set__ArtifactStore(v cfz.Expression[AWS_CodePipeline_Pipeline_ArtifactStore]) *AWS_CodePipeline_Pipeline {
	t.ArtifactStore = v
	return t
}

// SetV__ArtifactStore updates property "ArtifactStore".
func (t *AWS_CodePipeline_Pipeline) SetV__ArtifactStore(v AWS_CodePipeline_Pipeline_ArtifactStore) *AWS_CodePipeline_Pipeline {
	t.ArtifactStore = cfz.V(v)
	return t
}

// Set__ArtifactStores updates property "ArtifactStores".
func (t *AWS_CodePipeline_Pipeline) Set__ArtifactStores(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_ArtifactStoreMap]) *AWS_CodePipeline_Pipeline {
	t.ArtifactStores = v
	return t
}

// SetS__ArtifactStores updates property "ArtifactStores".
func (t *AWS_CodePipeline_Pipeline) SetS__ArtifactStores(v ...cfz.Expression[AWS_CodePipeline_Pipeline_ArtifactStoreMap]) *AWS_CodePipeline_Pipeline {
	t.ArtifactStores = cfz.S(v...)
	return t
}

// SetSV__ArtifactStores updates property "ArtifactStores".
func (t *AWS_CodePipeline_Pipeline) SetSV__ArtifactStores(v ...AWS_CodePipeline_Pipeline_ArtifactStoreMap) *AWS_CodePipeline_Pipeline {
	t.ArtifactStores = cfz.SV(v...)
	return t
}

// Set__DisableInboundStageTransitions updates property "DisableInboundStageTransitions".
func (t *AWS_CodePipeline_Pipeline) Set__DisableInboundStageTransitions(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_StageTransition]) *AWS_CodePipeline_Pipeline {
	t.DisableInboundStageTransitions = v
	return t
}

// SetS__DisableInboundStageTransitions updates property "DisableInboundStageTransitions".
func (t *AWS_CodePipeline_Pipeline) SetS__DisableInboundStageTransitions(v ...cfz.Expression[AWS_CodePipeline_Pipeline_StageTransition]) *AWS_CodePipeline_Pipeline {
	t.DisableInboundStageTransitions = cfz.S(v...)
	return t
}

// SetSV__DisableInboundStageTransitions updates property "DisableInboundStageTransitions".
func (t *AWS_CodePipeline_Pipeline) SetSV__DisableInboundStageTransitions(v ...AWS_CodePipeline_Pipeline_StageTransition) *AWS_CodePipeline_Pipeline {
	t.DisableInboundStageTransitions = cfz.SV(v...)
	return t
}

// Set__ExecutionMode updates property "ExecutionMode".
func (t *AWS_CodePipeline_Pipeline) Set__ExecutionMode(v cfz.Expression[string]) *AWS_CodePipeline_Pipeline {
	t.ExecutionMode = v
	return t
}

// SetV__ExecutionMode updates property "ExecutionMode".
func (t *AWS_CodePipeline_Pipeline) SetV__ExecutionMode(v string) *AWS_CodePipeline_Pipeline {
	t.ExecutionMode = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_CodePipeline_Pipeline) Set__Name(v cfz.Expression[string]) *AWS_CodePipeline_Pipeline {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_CodePipeline_Pipeline) SetV__Name(v string) *AWS_CodePipeline_Pipeline {
	t.Name = cfz.V(v)
	return t
}

// Set__PipelineType updates property "PipelineType".
func (t *AWS_CodePipeline_Pipeline) Set__PipelineType(v cfz.Expression[string]) *AWS_CodePipeline_Pipeline {
	t.PipelineType = v
	return t
}

// SetV__PipelineType updates property "PipelineType".
func (t *AWS_CodePipeline_Pipeline) SetV__PipelineType(v string) *AWS_CodePipeline_Pipeline {
	t.PipelineType = cfz.V(v)
	return t
}

// Set__RestartExecutionOnUpdate updates property "RestartExecutionOnUpdate".
func (t *AWS_CodePipeline_Pipeline) Set__RestartExecutionOnUpdate(v cfz.Expression[bool]) *AWS_CodePipeline_Pipeline {
	t.RestartExecutionOnUpdate = v
	return t
}

// SetV__RestartExecutionOnUpdate updates property "RestartExecutionOnUpdate".
func (t *AWS_CodePipeline_Pipeline) SetV__RestartExecutionOnUpdate(v bool) *AWS_CodePipeline_Pipeline {
	t.RestartExecutionOnUpdate = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_CodePipeline_Pipeline) Set__RoleArn(v cfz.Expression[string]) *AWS_CodePipeline_Pipeline {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_CodePipeline_Pipeline) SetV__RoleArn(v string) *AWS_CodePipeline_Pipeline {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__Stages updates property "Stages".
func (t *AWS_CodePipeline_Pipeline) Set__Stages(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_StageDeclaration]) *AWS_CodePipeline_Pipeline {
	t.Stages = v
	return t
}

// SetS__Stages updates property "Stages".
func (t *AWS_CodePipeline_Pipeline) SetS__Stages(v ...cfz.Expression[AWS_CodePipeline_Pipeline_StageDeclaration]) *AWS_CodePipeline_Pipeline {
	t.Stages = cfz.S(v...)
	return t
}

// SetSV__Stages updates property "Stages".
func (t *AWS_CodePipeline_Pipeline) SetSV__Stages(v ...AWS_CodePipeline_Pipeline_StageDeclaration) *AWS_CodePipeline_Pipeline {
	t.Stages = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_CodePipeline_Pipeline) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_CodePipeline_Pipeline {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_CodePipeline_Pipeline) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_CodePipeline_Pipeline {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_CodePipeline_Pipeline) SetSV__Tags(v ...cfz.Tag) *AWS_CodePipeline_Pipeline {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__Triggers updates property "Triggers".
func (t *AWS_CodePipeline_Pipeline) Set__Triggers(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration]) *AWS_CodePipeline_Pipeline {
	t.Triggers = v
	return t
}

// SetS__Triggers updates property "Triggers".
func (t *AWS_CodePipeline_Pipeline) SetS__Triggers(v ...cfz.Expression[AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration]) *AWS_CodePipeline_Pipeline {
	t.Triggers = cfz.S(v...)
	return t
}

// SetSV__Triggers updates property "Triggers".
func (t *AWS_CodePipeline_Pipeline) SetSV__Triggers(v ...AWS_CodePipeline_Pipeline_PipelineTriggerDeclaration) *AWS_CodePipeline_Pipeline {
	t.Triggers = cfz.SV(v...)
	return t
}

// Set__Variables updates property "Variables".
func (t *AWS_CodePipeline_Pipeline) Set__Variables(v cfz.ExpressionSlice[AWS_CodePipeline_Pipeline_VariableDeclaration]) *AWS_CodePipeline_Pipeline {
	t.Variables = v
	return t
}

// SetS__Variables updates property "Variables".
func (t *AWS_CodePipeline_Pipeline) SetS__Variables(v ...cfz.Expression[AWS_CodePipeline_Pipeline_VariableDeclaration]) *AWS_CodePipeline_Pipeline {
	t.Variables = cfz.S(v...)
	return t
}

// SetSV__Variables updates property "Variables".
func (t *AWS_CodePipeline_Pipeline) SetSV__Variables(v ...AWS_CodePipeline_Pipeline_VariableDeclaration) *AWS_CodePipeline_Pipeline {
	t.Variables = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_CodePipeline_Pipeline) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Version returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Version
func (t *AWS_CodePipeline_Pipeline) GetAtt__Version() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CodePipeline_Pipeline__AttributesMap.Version))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_CodePipeline_Pipeline) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Version returns a conventionally configured output for an attribute of this resource.
// Attribute: Version
func (t *AWS_CodePipeline_Pipeline) GetConventionalOutputAtt__Version(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVersion", t.GetAtt__Version())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_CodePipeline_Pipeline) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_CodePipeline_Pipeline

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

func (t *AWS_CodePipeline_Pipeline) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
