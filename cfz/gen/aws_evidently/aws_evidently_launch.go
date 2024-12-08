// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_evidently

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Evidently_Launch)(nil)
	_ cfz.Resource                   = (*AWS_Evidently_Launch)(nil)
)

const (
	// AWS_Evidently_Launch__Type is the CloudFormation type for AWS::Evidently::Launch.
	AWS_Evidently_Launch__Type = "AWS::Evidently::Launch"
)

var (
	// AWS_Evidently_Launch__AttributesMap reports all the CloudFormation attributes for AWS::Evidently::Launch.
	AWS_Evidently_Launch__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Evidently_Launch__AttributesSlice reports all the CloudFormation attributes for AWS::Evidently::Launch.
	AWS_Evidently_Launch__AttributesSlice = []string{
		AWS_Evidently_Launch__AttributesMap.Arn,
	}
)

var (
	// AWS_Evidently_Launch__PropertiesMap reports all the CloudFormation properties for AWS::Evidently::Launch.
	AWS_Evidently_Launch__PropertiesMap = struct {
		Description           string
		ExecutionStatus       string
		Groups                string
		MetricMonitors        string
		Name                  string
		Project               string
		RandomizationSalt     string
		ScheduledSplitsConfig string
		Tags                  string
	}{
		Description:           "Description",
		ExecutionStatus:       "ExecutionStatus",
		Groups:                "Groups",
		MetricMonitors:        "MetricMonitors",
		Name:                  "Name",
		Project:               "Project",
		RandomizationSalt:     "RandomizationSalt",
		ScheduledSplitsConfig: "ScheduledSplitsConfig",
		Tags:                  "Tags",
	}

	// AWS_Evidently_Launch__PropertiesSlice reports all the CloudFormation properties for AWS::Evidently::Launch.
	AWS_Evidently_Launch__PropertiesSlice = []string{
		AWS_Evidently_Launch__PropertiesMap.Description,
		AWS_Evidently_Launch__PropertiesMap.ExecutionStatus,
		AWS_Evidently_Launch__PropertiesMap.Groups,
		AWS_Evidently_Launch__PropertiesMap.MetricMonitors,
		AWS_Evidently_Launch__PropertiesMap.Name,
		AWS_Evidently_Launch__PropertiesMap.Project,
		AWS_Evidently_Launch__PropertiesMap.RandomizationSalt,
		AWS_Evidently_Launch__PropertiesMap.ScheduledSplitsConfig,
		AWS_Evidently_Launch__PropertiesMap.Tags,
	}
)

// AWS_Evidently_Launch is a binding for AWS::Evidently::Launch.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html
type AWS_Evidently_Launch struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// ExecutionStatus is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-executionstatus
	ExecutionStatus cfz.Expression[AWS_Evidently_Launch_ExecutionStatusObject] `json:"ExecutionStatus,omitempty"`

	// Groups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-groups
	Groups cfz.ExpressionSlice[AWS_Evidently_Launch_LaunchGroupObject] `json:"Groups,omitempty"`

	// MetricMonitors is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-metricmonitors
	MetricMonitors cfz.ExpressionSlice[AWS_Evidently_Launch_MetricDefinitionObject] `json:"MetricMonitors,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Project is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-project
	Project cfz.Expression[string] `json:"Project,omitempty"`

	// RandomizationSalt is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-randomizationsalt
	RandomizationSalt cfz.Expression[string] `json:"RandomizationSalt,omitempty"`

	// ScheduledSplitsConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-scheduledsplitsconfig
	ScheduledSplitsConfig cfz.ExpressionSlice[AWS_Evidently_Launch_StepConfig] `json:"ScheduledSplitsConfig,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-evidently-launch.html#cfn-evidently-launch-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Evidently_Launch initializes a new *AWS_Evidently_Launch.
func New__AWS_Evidently_Launch(logicalName string) *AWS_Evidently_Launch {
	return &AWS_Evidently_Launch{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Evidently_Launch) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Evidently_Launch) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Evidently_Launch) GetType() string {
	return AWS_Evidently_Launch__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Evidently_Launch) Set__LogicalName(v string) *AWS_Evidently_Launch {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Evidently_Launch) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Evidently_Launch {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Evidently_Launch) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Evidently_Launch {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Evidently_Launch) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Evidently_Launch {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Evidently_Launch) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Evidently_Launch {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Evidently_Launch) Set__RequestedOutputs(v []cfz.Output) *AWS_Evidently_Launch {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Evidently_Launch) Add__RequestedOutputs(v ...cfz.Output) *AWS_Evidently_Launch {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Evidently_Launch) Set__Description(v cfz.Expression[string]) *AWS_Evidently_Launch {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Evidently_Launch) SetV__Description(v string) *AWS_Evidently_Launch {
	t.Description = cfz.V(v)
	return t
}

// Set__ExecutionStatus updates property "ExecutionStatus".
func (t *AWS_Evidently_Launch) Set__ExecutionStatus(v cfz.Expression[AWS_Evidently_Launch_ExecutionStatusObject]) *AWS_Evidently_Launch {
	t.ExecutionStatus = v
	return t
}

// SetV__ExecutionStatus updates property "ExecutionStatus".
func (t *AWS_Evidently_Launch) SetV__ExecutionStatus(v AWS_Evidently_Launch_ExecutionStatusObject) *AWS_Evidently_Launch {
	t.ExecutionStatus = cfz.V(v)
	return t
}

// Set__Groups updates property "Groups".
func (t *AWS_Evidently_Launch) Set__Groups(v cfz.ExpressionSlice[AWS_Evidently_Launch_LaunchGroupObject]) *AWS_Evidently_Launch {
	t.Groups = v
	return t
}

// SetS__Groups updates property "Groups".
func (t *AWS_Evidently_Launch) SetS__Groups(v ...cfz.Expression[AWS_Evidently_Launch_LaunchGroupObject]) *AWS_Evidently_Launch {
	t.Groups = cfz.S(v...)
	return t
}

// SetSV__Groups updates property "Groups".
func (t *AWS_Evidently_Launch) SetSV__Groups(v ...AWS_Evidently_Launch_LaunchGroupObject) *AWS_Evidently_Launch {
	t.Groups = cfz.SV(v...)
	return t
}

// Set__MetricMonitors updates property "MetricMonitors".
func (t *AWS_Evidently_Launch) Set__MetricMonitors(v cfz.ExpressionSlice[AWS_Evidently_Launch_MetricDefinitionObject]) *AWS_Evidently_Launch {
	t.MetricMonitors = v
	return t
}

// SetS__MetricMonitors updates property "MetricMonitors".
func (t *AWS_Evidently_Launch) SetS__MetricMonitors(v ...cfz.Expression[AWS_Evidently_Launch_MetricDefinitionObject]) *AWS_Evidently_Launch {
	t.MetricMonitors = cfz.S(v...)
	return t
}

// SetSV__MetricMonitors updates property "MetricMonitors".
func (t *AWS_Evidently_Launch) SetSV__MetricMonitors(v ...AWS_Evidently_Launch_MetricDefinitionObject) *AWS_Evidently_Launch {
	t.MetricMonitors = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Evidently_Launch) Set__Name(v cfz.Expression[string]) *AWS_Evidently_Launch {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Evidently_Launch) SetV__Name(v string) *AWS_Evidently_Launch {
	t.Name = cfz.V(v)
	return t
}

// Set__Project updates property "Project".
func (t *AWS_Evidently_Launch) Set__Project(v cfz.Expression[string]) *AWS_Evidently_Launch {
	t.Project = v
	return t
}

// SetV__Project updates property "Project".
func (t *AWS_Evidently_Launch) SetV__Project(v string) *AWS_Evidently_Launch {
	t.Project = cfz.V(v)
	return t
}

// Set__RandomizationSalt updates property "RandomizationSalt".
func (t *AWS_Evidently_Launch) Set__RandomizationSalt(v cfz.Expression[string]) *AWS_Evidently_Launch {
	t.RandomizationSalt = v
	return t
}

// SetV__RandomizationSalt updates property "RandomizationSalt".
func (t *AWS_Evidently_Launch) SetV__RandomizationSalt(v string) *AWS_Evidently_Launch {
	t.RandomizationSalt = cfz.V(v)
	return t
}

// Set__ScheduledSplitsConfig updates property "ScheduledSplitsConfig".
func (t *AWS_Evidently_Launch) Set__ScheduledSplitsConfig(v cfz.ExpressionSlice[AWS_Evidently_Launch_StepConfig]) *AWS_Evidently_Launch {
	t.ScheduledSplitsConfig = v
	return t
}

// SetS__ScheduledSplitsConfig updates property "ScheduledSplitsConfig".
func (t *AWS_Evidently_Launch) SetS__ScheduledSplitsConfig(v ...cfz.Expression[AWS_Evidently_Launch_StepConfig]) *AWS_Evidently_Launch {
	t.ScheduledSplitsConfig = cfz.S(v...)
	return t
}

// SetSV__ScheduledSplitsConfig updates property "ScheduledSplitsConfig".
func (t *AWS_Evidently_Launch) SetSV__ScheduledSplitsConfig(v ...AWS_Evidently_Launch_StepConfig) *AWS_Evidently_Launch {
	t.ScheduledSplitsConfig = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Evidently_Launch) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Evidently_Launch {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Evidently_Launch) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Evidently_Launch {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Evidently_Launch) SetSV__Tags(v ...cfz.Tag) *AWS_Evidently_Launch {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Evidently_Launch) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Evidently_Launch) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Evidently_Launch__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Evidently_Launch) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Evidently_Launch) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Evidently_Launch) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Evidently_Launch

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

func (t *AWS_Evidently_Launch) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
