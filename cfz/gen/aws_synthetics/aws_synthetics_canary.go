// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_synthetics

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Synthetics_Canary)(nil)
	_ cfz.Resource                   = (*AWS_Synthetics_Canary)(nil)
)

const (
	// AWS_Synthetics_Canary__Type is the CloudFormation type for AWS::Synthetics::Canary.
	AWS_Synthetics_Canary__Type = "AWS::Synthetics::Canary"
)

var (
	// AWS_Synthetics_Canary__AttributesMap reports all the CloudFormation attributes for AWS::Synthetics::Canary.
	AWS_Synthetics_Canary__AttributesMap = struct {
		Code_SourceLocationArn string
		Id                     string
		State                  string
	}{
		Code_SourceLocationArn: "Code.SourceLocationArn",
		Id:                     "Id",
		State:                  "State",
	}

	// AWS_Synthetics_Canary__AttributesSlice reports all the CloudFormation attributes for AWS::Synthetics::Canary.
	AWS_Synthetics_Canary__AttributesSlice = []string{
		AWS_Synthetics_Canary__AttributesMap.Code_SourceLocationArn,
		AWS_Synthetics_Canary__AttributesMap.Id,
		AWS_Synthetics_Canary__AttributesMap.State,
	}
)

var (
	// AWS_Synthetics_Canary__PropertiesMap reports all the CloudFormation properties for AWS::Synthetics::Canary.
	AWS_Synthetics_Canary__PropertiesMap = struct {
		ArtifactConfig             string
		ArtifactS3Location         string
		Code                       string
		ExecutionRoleArn           string
		FailureRetentionPeriod     string
		Name                       string
		ProvisionedResourceCleanup string
		ResourcesToReplicateTags   string
		RunConfig                  string
		RuntimeVersion             string
		Schedule                   string
		StartCanaryAfterCreation   string
		SuccessRetentionPeriod     string
		Tags                       string
		VPCConfig                  string
		VisualReference            string
	}{
		ArtifactConfig:             "ArtifactConfig",
		ArtifactS3Location:         "ArtifactS3Location",
		Code:                       "Code",
		ExecutionRoleArn:           "ExecutionRoleArn",
		FailureRetentionPeriod:     "FailureRetentionPeriod",
		Name:                       "Name",
		ProvisionedResourceCleanup: "ProvisionedResourceCleanup",
		ResourcesToReplicateTags:   "ResourcesToReplicateTags",
		RunConfig:                  "RunConfig",
		RuntimeVersion:             "RuntimeVersion",
		Schedule:                   "Schedule",
		StartCanaryAfterCreation:   "StartCanaryAfterCreation",
		SuccessRetentionPeriod:     "SuccessRetentionPeriod",
		Tags:                       "Tags",
		VPCConfig:                  "VPCConfig",
		VisualReference:            "VisualReference",
	}

	// AWS_Synthetics_Canary__PropertiesSlice reports all the CloudFormation properties for AWS::Synthetics::Canary.
	AWS_Synthetics_Canary__PropertiesSlice = []string{
		AWS_Synthetics_Canary__PropertiesMap.ArtifactConfig,
		AWS_Synthetics_Canary__PropertiesMap.ArtifactS3Location,
		AWS_Synthetics_Canary__PropertiesMap.Code,
		AWS_Synthetics_Canary__PropertiesMap.ExecutionRoleArn,
		AWS_Synthetics_Canary__PropertiesMap.FailureRetentionPeriod,
		AWS_Synthetics_Canary__PropertiesMap.Name,
		AWS_Synthetics_Canary__PropertiesMap.ProvisionedResourceCleanup,
		AWS_Synthetics_Canary__PropertiesMap.ResourcesToReplicateTags,
		AWS_Synthetics_Canary__PropertiesMap.RunConfig,
		AWS_Synthetics_Canary__PropertiesMap.RuntimeVersion,
		AWS_Synthetics_Canary__PropertiesMap.Schedule,
		AWS_Synthetics_Canary__PropertiesMap.StartCanaryAfterCreation,
		AWS_Synthetics_Canary__PropertiesMap.SuccessRetentionPeriod,
		AWS_Synthetics_Canary__PropertiesMap.Tags,
		AWS_Synthetics_Canary__PropertiesMap.VPCConfig,
		AWS_Synthetics_Canary__PropertiesMap.VisualReference,
	}
)

// AWS_Synthetics_Canary is a binding for AWS::Synthetics::Canary.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html
type AWS_Synthetics_Canary struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ArtifactConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-artifactconfig
	ArtifactConfig cfz.Expression[AWS_Synthetics_Canary_ArtifactConfig] `json:"ArtifactConfig,omitempty"`

	// ArtifactS3Location is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-artifacts3location
	ArtifactS3Location cfz.Expression[string] `json:"ArtifactS3Location,omitempty"`

	// Code is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-code
	Code cfz.Expression[AWS_Synthetics_Canary_Code] `json:"Code,omitempty"`

	// ExecutionRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
	ExecutionRoleArn cfz.Expression[string] `json:"ExecutionRoleArn,omitempty"`

	// FailureRetentionPeriod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-failureretentionperiod
	FailureRetentionPeriod cfz.Expression[int32] `json:"FailureRetentionPeriod,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ProvisionedResourceCleanup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-provisionedresourcecleanup
	ProvisionedResourceCleanup cfz.Expression[string] `json:"ProvisionedResourceCleanup,omitempty"`

	// ResourcesToReplicateTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-resourcestoreplicatetags
	ResourcesToReplicateTags cfz.ExpressionSlice[string] `json:"ResourcesToReplicateTags,omitempty"`

	// RunConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runconfig
	RunConfig cfz.Expression[AWS_Synthetics_Canary_RunConfig] `json:"RunConfig,omitempty"`

	// RuntimeVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-runtimeversion
	RuntimeVersion cfz.Expression[string] `json:"RuntimeVersion,omitempty"`

	// Schedule is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-schedule
	Schedule cfz.Expression[AWS_Synthetics_Canary_Schedule] `json:"Schedule,omitempty"`

	// StartCanaryAfterCreation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-startcanaryaftercreation
	StartCanaryAfterCreation cfz.Expression[bool] `json:"StartCanaryAfterCreation,omitempty"`

	// SuccessRetentionPeriod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-successretentionperiod
	SuccessRetentionPeriod cfz.Expression[int32] `json:"SuccessRetentionPeriod,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// VPCConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-vpcconfig
	VPCConfig cfz.Expression[AWS_Synthetics_Canary_VPCConfig] `json:"VPCConfig,omitempty"`

	// VisualReference is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-visualreference
	VisualReference cfz.Expression[AWS_Synthetics_Canary_VisualReference] `json:"VisualReference,omitempty"`
}

// New__AWS_Synthetics_Canary initializes a new *AWS_Synthetics_Canary.
func New__AWS_Synthetics_Canary(logicalName string) *AWS_Synthetics_Canary {
	return &AWS_Synthetics_Canary{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Synthetics_Canary) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Synthetics_Canary) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Synthetics_Canary) GetType() string {
	return AWS_Synthetics_Canary__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Synthetics_Canary) Set__LogicalName(v string) *AWS_Synthetics_Canary {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Synthetics_Canary) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Synthetics_Canary {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Synthetics_Canary) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Synthetics_Canary {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Synthetics_Canary) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Synthetics_Canary {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Synthetics_Canary) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Synthetics_Canary {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Synthetics_Canary) Set__RequestedOutputs(v []cfz.Output) *AWS_Synthetics_Canary {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Synthetics_Canary) Add__RequestedOutputs(v ...cfz.Output) *AWS_Synthetics_Canary {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ArtifactConfig updates property "ArtifactConfig".
func (t *AWS_Synthetics_Canary) Set__ArtifactConfig(v cfz.Expression[AWS_Synthetics_Canary_ArtifactConfig]) *AWS_Synthetics_Canary {
	t.ArtifactConfig = v
	return t
}

// SetV__ArtifactConfig updates property "ArtifactConfig".
func (t *AWS_Synthetics_Canary) SetV__ArtifactConfig(v AWS_Synthetics_Canary_ArtifactConfig) *AWS_Synthetics_Canary {
	t.ArtifactConfig = cfz.V(v)
	return t
}

// Set__ArtifactS3Location updates property "ArtifactS3Location".
func (t *AWS_Synthetics_Canary) Set__ArtifactS3Location(v cfz.Expression[string]) *AWS_Synthetics_Canary {
	t.ArtifactS3Location = v
	return t
}

// SetV__ArtifactS3Location updates property "ArtifactS3Location".
func (t *AWS_Synthetics_Canary) SetV__ArtifactS3Location(v string) *AWS_Synthetics_Canary {
	t.ArtifactS3Location = cfz.V(v)
	return t
}

// Set__Code updates property "Code".
func (t *AWS_Synthetics_Canary) Set__Code(v cfz.Expression[AWS_Synthetics_Canary_Code]) *AWS_Synthetics_Canary {
	t.Code = v
	return t
}

// SetV__Code updates property "Code".
func (t *AWS_Synthetics_Canary) SetV__Code(v AWS_Synthetics_Canary_Code) *AWS_Synthetics_Canary {
	t.Code = cfz.V(v)
	return t
}

// Set__ExecutionRoleArn updates property "ExecutionRoleArn".
func (t *AWS_Synthetics_Canary) Set__ExecutionRoleArn(v cfz.Expression[string]) *AWS_Synthetics_Canary {
	t.ExecutionRoleArn = v
	return t
}

// SetV__ExecutionRoleArn updates property "ExecutionRoleArn".
func (t *AWS_Synthetics_Canary) SetV__ExecutionRoleArn(v string) *AWS_Synthetics_Canary {
	t.ExecutionRoleArn = cfz.V(v)
	return t
}

// Set__FailureRetentionPeriod updates property "FailureRetentionPeriod".
func (t *AWS_Synthetics_Canary) Set__FailureRetentionPeriod(v cfz.Expression[int32]) *AWS_Synthetics_Canary {
	t.FailureRetentionPeriod = v
	return t
}

// SetV__FailureRetentionPeriod updates property "FailureRetentionPeriod".
func (t *AWS_Synthetics_Canary) SetV__FailureRetentionPeriod(v int32) *AWS_Synthetics_Canary {
	t.FailureRetentionPeriod = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Synthetics_Canary) Set__Name(v cfz.Expression[string]) *AWS_Synthetics_Canary {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Synthetics_Canary) SetV__Name(v string) *AWS_Synthetics_Canary {
	t.Name = cfz.V(v)
	return t
}

// Set__ProvisionedResourceCleanup updates property "ProvisionedResourceCleanup".
func (t *AWS_Synthetics_Canary) Set__ProvisionedResourceCleanup(v cfz.Expression[string]) *AWS_Synthetics_Canary {
	t.ProvisionedResourceCleanup = v
	return t
}

// SetV__ProvisionedResourceCleanup updates property "ProvisionedResourceCleanup".
func (t *AWS_Synthetics_Canary) SetV__ProvisionedResourceCleanup(v string) *AWS_Synthetics_Canary {
	t.ProvisionedResourceCleanup = cfz.V(v)
	return t
}

// Set__ResourcesToReplicateTags updates property "ResourcesToReplicateTags".
func (t *AWS_Synthetics_Canary) Set__ResourcesToReplicateTags(v cfz.ExpressionSlice[string]) *AWS_Synthetics_Canary {
	t.ResourcesToReplicateTags = v
	return t
}

// SetS__ResourcesToReplicateTags updates property "ResourcesToReplicateTags".
func (t *AWS_Synthetics_Canary) SetS__ResourcesToReplicateTags(v ...cfz.Expression[string]) *AWS_Synthetics_Canary {
	t.ResourcesToReplicateTags = cfz.S(v...)
	return t
}

// SetSV__ResourcesToReplicateTags updates property "ResourcesToReplicateTags".
func (t *AWS_Synthetics_Canary) SetSV__ResourcesToReplicateTags(v ...string) *AWS_Synthetics_Canary {
	t.ResourcesToReplicateTags = cfz.SV(v...)
	return t
}

// Set__RunConfig updates property "RunConfig".
func (t *AWS_Synthetics_Canary) Set__RunConfig(v cfz.Expression[AWS_Synthetics_Canary_RunConfig]) *AWS_Synthetics_Canary {
	t.RunConfig = v
	return t
}

// SetV__RunConfig updates property "RunConfig".
func (t *AWS_Synthetics_Canary) SetV__RunConfig(v AWS_Synthetics_Canary_RunConfig) *AWS_Synthetics_Canary {
	t.RunConfig = cfz.V(v)
	return t
}

// Set__RuntimeVersion updates property "RuntimeVersion".
func (t *AWS_Synthetics_Canary) Set__RuntimeVersion(v cfz.Expression[string]) *AWS_Synthetics_Canary {
	t.RuntimeVersion = v
	return t
}

// SetV__RuntimeVersion updates property "RuntimeVersion".
func (t *AWS_Synthetics_Canary) SetV__RuntimeVersion(v string) *AWS_Synthetics_Canary {
	t.RuntimeVersion = cfz.V(v)
	return t
}

// Set__Schedule updates property "Schedule".
func (t *AWS_Synthetics_Canary) Set__Schedule(v cfz.Expression[AWS_Synthetics_Canary_Schedule]) *AWS_Synthetics_Canary {
	t.Schedule = v
	return t
}

// SetV__Schedule updates property "Schedule".
func (t *AWS_Synthetics_Canary) SetV__Schedule(v AWS_Synthetics_Canary_Schedule) *AWS_Synthetics_Canary {
	t.Schedule = cfz.V(v)
	return t
}

// Set__StartCanaryAfterCreation updates property "StartCanaryAfterCreation".
func (t *AWS_Synthetics_Canary) Set__StartCanaryAfterCreation(v cfz.Expression[bool]) *AWS_Synthetics_Canary {
	t.StartCanaryAfterCreation = v
	return t
}

// SetV__StartCanaryAfterCreation updates property "StartCanaryAfterCreation".
func (t *AWS_Synthetics_Canary) SetV__StartCanaryAfterCreation(v bool) *AWS_Synthetics_Canary {
	t.StartCanaryAfterCreation = cfz.V(v)
	return t
}

// Set__SuccessRetentionPeriod updates property "SuccessRetentionPeriod".
func (t *AWS_Synthetics_Canary) Set__SuccessRetentionPeriod(v cfz.Expression[int32]) *AWS_Synthetics_Canary {
	t.SuccessRetentionPeriod = v
	return t
}

// SetV__SuccessRetentionPeriod updates property "SuccessRetentionPeriod".
func (t *AWS_Synthetics_Canary) SetV__SuccessRetentionPeriod(v int32) *AWS_Synthetics_Canary {
	t.SuccessRetentionPeriod = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Synthetics_Canary) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Synthetics_Canary {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Synthetics_Canary) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Synthetics_Canary {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Synthetics_Canary) SetSV__Tags(v ...cfz.Tag) *AWS_Synthetics_Canary {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__VPCConfig updates property "VPCConfig".
func (t *AWS_Synthetics_Canary) Set__VPCConfig(v cfz.Expression[AWS_Synthetics_Canary_VPCConfig]) *AWS_Synthetics_Canary {
	t.VPCConfig = v
	return t
}

// SetV__VPCConfig updates property "VPCConfig".
func (t *AWS_Synthetics_Canary) SetV__VPCConfig(v AWS_Synthetics_Canary_VPCConfig) *AWS_Synthetics_Canary {
	t.VPCConfig = cfz.V(v)
	return t
}

// Set__VisualReference updates property "VisualReference".
func (t *AWS_Synthetics_Canary) Set__VisualReference(v cfz.Expression[AWS_Synthetics_Canary_VisualReference]) *AWS_Synthetics_Canary {
	t.VisualReference = v
	return t
}

// SetV__VisualReference updates property "VisualReference".
func (t *AWS_Synthetics_Canary) SetV__VisualReference(v AWS_Synthetics_Canary_VisualReference) *AWS_Synthetics_Canary {
	t.VisualReference = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Synthetics_Canary) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Code_SourceLocationArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Code.SourceLocationArn
func (t *AWS_Synthetics_Canary) GetAtt__Code_SourceLocationArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Synthetics_Canary__AttributesMap.Code_SourceLocationArn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Synthetics_Canary) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Synthetics_Canary__AttributesMap.Id))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_Synthetics_Canary) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Synthetics_Canary__AttributesMap.State))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Synthetics_Canary) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Code_SourceLocationArn returns a conventionally configured output for an attribute of this resource.
// Attribute: Code.SourceLocationArn
func (t *AWS_Synthetics_Canary) GetConventionalOutputAtt__Code_SourceLocationArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCodeSourceLocationArn", t.GetAtt__Code_SourceLocationArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Synthetics_Canary) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_Synthetics_Canary) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Synthetics_Canary) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Synthetics_Canary

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

func (t *AWS_Synthetics_Canary) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
