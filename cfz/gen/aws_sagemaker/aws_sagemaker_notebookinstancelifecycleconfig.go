// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_SageMaker_NotebookInstanceLifecycleConfig)(nil)
	_ cfz.Resource                   = (*AWS_SageMaker_NotebookInstanceLifecycleConfig)(nil)
)

const (
	// AWS_SageMaker_NotebookInstanceLifecycleConfig__Type is the CloudFormation type for AWS::SageMaker::NotebookInstanceLifecycleConfig.
	AWS_SageMaker_NotebookInstanceLifecycleConfig__Type = "AWS::SageMaker::NotebookInstanceLifecycleConfig"
)

var (
	// AWS_SageMaker_NotebookInstanceLifecycleConfig__AttributesMap reports all the CloudFormation attributes for AWS::SageMaker::NotebookInstanceLifecycleConfig.
	AWS_SageMaker_NotebookInstanceLifecycleConfig__AttributesMap = struct {
		NotebookInstanceLifecycleConfigName string
	}{
		NotebookInstanceLifecycleConfigName: "NotebookInstanceLifecycleConfigName",
	}

	// AWS_SageMaker_NotebookInstanceLifecycleConfig__AttributesSlice reports all the CloudFormation attributes for AWS::SageMaker::NotebookInstanceLifecycleConfig.
	AWS_SageMaker_NotebookInstanceLifecycleConfig__AttributesSlice = []string{
		AWS_SageMaker_NotebookInstanceLifecycleConfig__AttributesMap.NotebookInstanceLifecycleConfigName,
	}
)

var (
	// AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::NotebookInstanceLifecycleConfig.
	AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesMap = struct {
		NotebookInstanceLifecycleConfigName string
		OnCreate                            string
		OnStart                             string
	}{
		NotebookInstanceLifecycleConfigName: "NotebookInstanceLifecycleConfigName",
		OnCreate:                            "OnCreate",
		OnStart:                             "OnStart",
	}

	// AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::NotebookInstanceLifecycleConfig.
	AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesSlice = []string{
		AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesMap.NotebookInstanceLifecycleConfigName,
		AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesMap.OnCreate,
		AWS_SageMaker_NotebookInstanceLifecycleConfig__PropertiesMap.OnStart,
	}
)

// AWS_SageMaker_NotebookInstanceLifecycleConfig is a binding for AWS::SageMaker::NotebookInstanceLifecycleConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html
type AWS_SageMaker_NotebookInstanceLifecycleConfig struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// NotebookInstanceLifecycleConfigName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecycleconfigname
	NotebookInstanceLifecycleConfigName cfz.Expression[string] `json:"NotebookInstanceLifecycleConfigName,omitempty"`

	// OnCreate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-oncreate
	OnCreate cfz.ExpressionSlice[AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook] `json:"OnCreate,omitempty"`

	// OnStart is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-onstart
	OnStart cfz.ExpressionSlice[AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook] `json:"OnStart,omitempty"`
}

// New__AWS_SageMaker_NotebookInstanceLifecycleConfig initializes a new *AWS_SageMaker_NotebookInstanceLifecycleConfig.
func New__AWS_SageMaker_NotebookInstanceLifecycleConfig(logicalName string) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	return &AWS_SageMaker_NotebookInstanceLifecycleConfig{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_SageMaker_NotebookInstanceLifecycleConfig) GetType() string {
	return AWS_SageMaker_NotebookInstanceLifecycleConfig__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__LogicalName(v string) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__RequestedOutputs(v []cfz.Output) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Add__RequestedOutputs(v ...cfz.Output) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__NotebookInstanceLifecycleConfigName updates property "NotebookInstanceLifecycleConfigName".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__NotebookInstanceLifecycleConfigName(v cfz.Expression[string]) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.NotebookInstanceLifecycleConfigName = v
	return t
}

// SetV__NotebookInstanceLifecycleConfigName updates property "NotebookInstanceLifecycleConfigName".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) SetV__NotebookInstanceLifecycleConfigName(v string) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.NotebookInstanceLifecycleConfigName = cfz.V(v)
	return t
}

// Set__OnCreate updates property "OnCreate".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__OnCreate(v cfz.ExpressionSlice[AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook]) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.OnCreate = v
	return t
}

// SetS__OnCreate updates property "OnCreate".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) SetS__OnCreate(v ...cfz.Expression[AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook]) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.OnCreate = cfz.S(v...)
	return t
}

// SetSV__OnCreate updates property "OnCreate".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) SetSV__OnCreate(v ...AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.OnCreate = cfz.SV(v...)
	return t
}

// Set__OnStart updates property "OnStart".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Set__OnStart(v cfz.ExpressionSlice[AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook]) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.OnStart = v
	return t
}

// SetS__OnStart updates property "OnStart".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) SetS__OnStart(v ...cfz.Expression[AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook]) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.OnStart = cfz.S(v...)
	return t
}

// SetSV__OnStart updates property "OnStart".
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) SetSV__OnStart(v ...AWS_SageMaker_NotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) *AWS_SageMaker_NotebookInstanceLifecycleConfig {
	t.OnStart = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__NotebookInstanceLifecycleConfigName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NotebookInstanceLifecycleConfigName
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) GetAtt__NotebookInstanceLifecycleConfigName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_NotebookInstanceLifecycleConfig__AttributesMap.NotebookInstanceLifecycleConfigName))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NotebookInstanceLifecycleConfigName returns a conventionally configured output for an attribute of this resource.
// Attribute: NotebookInstanceLifecycleConfigName
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) GetConventionalOutputAtt__NotebookInstanceLifecycleConfigName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNotebookInstanceLifecycleConfigName", t.GetAtt__NotebookInstanceLifecycleConfigName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_SageMaker_NotebookInstanceLifecycleConfig

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

func (t *AWS_SageMaker_NotebookInstanceLifecycleConfig) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
