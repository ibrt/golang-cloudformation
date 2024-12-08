// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Pipes_Pipe)(nil)
	_ cfz.Resource                   = (*AWS_Pipes_Pipe)(nil)
)

const (
	// AWS_Pipes_Pipe__Type is the CloudFormation type for AWS::Pipes::Pipe.
	AWS_Pipes_Pipe__Type = "AWS::Pipes::Pipe"
)

var (
	// AWS_Pipes_Pipe__AttributesMap reports all the CloudFormation attributes for AWS::Pipes::Pipe.
	AWS_Pipes_Pipe__AttributesMap = struct {
		Arn              string
		CreationTime     string
		CurrentState     string
		LastModifiedTime string
		StateReason      string
	}{
		Arn:              "Arn",
		CreationTime:     "CreationTime",
		CurrentState:     "CurrentState",
		LastModifiedTime: "LastModifiedTime",
		StateReason:      "StateReason",
	}

	// AWS_Pipes_Pipe__AttributesSlice reports all the CloudFormation attributes for AWS::Pipes::Pipe.
	AWS_Pipes_Pipe__AttributesSlice = []string{
		AWS_Pipes_Pipe__AttributesMap.Arn,
		AWS_Pipes_Pipe__AttributesMap.CreationTime,
		AWS_Pipes_Pipe__AttributesMap.CurrentState,
		AWS_Pipes_Pipe__AttributesMap.LastModifiedTime,
		AWS_Pipes_Pipe__AttributesMap.StateReason,
	}
)

var (
	// AWS_Pipes_Pipe__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.
	AWS_Pipes_Pipe__PropertiesMap = struct {
		Description          string
		DesiredState         string
		Enrichment           string
		EnrichmentParameters string
		KmsKeyIdentifier     string
		LogConfiguration     string
		Name                 string
		RoleArn              string
		Source               string
		SourceParameters     string
		Tags                 string
		Target               string
		TargetParameters     string
	}{
		Description:          "Description",
		DesiredState:         "DesiredState",
		Enrichment:           "Enrichment",
		EnrichmentParameters: "EnrichmentParameters",
		KmsKeyIdentifier:     "KmsKeyIdentifier",
		LogConfiguration:     "LogConfiguration",
		Name:                 "Name",
		RoleArn:              "RoleArn",
		Source:               "Source",
		SourceParameters:     "SourceParameters",
		Tags:                 "Tags",
		Target:               "Target",
		TargetParameters:     "TargetParameters",
	}

	// AWS_Pipes_Pipe__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.
	AWS_Pipes_Pipe__PropertiesSlice = []string{
		AWS_Pipes_Pipe__PropertiesMap.Description,
		AWS_Pipes_Pipe__PropertiesMap.DesiredState,
		AWS_Pipes_Pipe__PropertiesMap.Enrichment,
		AWS_Pipes_Pipe__PropertiesMap.EnrichmentParameters,
		AWS_Pipes_Pipe__PropertiesMap.KmsKeyIdentifier,
		AWS_Pipes_Pipe__PropertiesMap.LogConfiguration,
		AWS_Pipes_Pipe__PropertiesMap.Name,
		AWS_Pipes_Pipe__PropertiesMap.RoleArn,
		AWS_Pipes_Pipe__PropertiesMap.Source,
		AWS_Pipes_Pipe__PropertiesMap.SourceParameters,
		AWS_Pipes_Pipe__PropertiesMap.Tags,
		AWS_Pipes_Pipe__PropertiesMap.Target,
		AWS_Pipes_Pipe__PropertiesMap.TargetParameters,
	}
)

// AWS_Pipes_Pipe is a binding for AWS::Pipes::Pipe.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html
type AWS_Pipes_Pipe struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DesiredState is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-desiredstate
	DesiredState cfz.Expression[string] `json:"DesiredState,omitempty"`

	// Enrichment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-enrichment
	Enrichment cfz.Expression[string] `json:"Enrichment,omitempty"`

	// EnrichmentParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-enrichmentparameters
	EnrichmentParameters cfz.Expression[AWS_Pipes_Pipe_PipeEnrichmentParameters] `json:"EnrichmentParameters,omitempty"`

	// KmsKeyIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-kmskeyidentifier
	KmsKeyIdentifier cfz.Expression[string] `json:"KmsKeyIdentifier,omitempty"`

	// LogConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-logconfiguration
	LogConfiguration cfz.Expression[AWS_Pipes_Pipe_PipeLogConfiguration] `json:"LogConfiguration,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// Source is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-source
	Source cfz.Expression[string] `json:"Source,omitempty"`

	// SourceParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-sourceparameters
	SourceParameters cfz.Expression[AWS_Pipes_Pipe_PipeSourceParameters] `json:"SourceParameters,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`

	// Target is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-target
	Target cfz.Expression[string] `json:"Target,omitempty"`

	// TargetParameters is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-targetparameters
	TargetParameters cfz.Expression[AWS_Pipes_Pipe_PipeTargetParameters] `json:"TargetParameters,omitempty"`
}

// New__AWS_Pipes_Pipe initializes a new *AWS_Pipes_Pipe.
func New__AWS_Pipes_Pipe(logicalName string) *AWS_Pipes_Pipe {
	return &AWS_Pipes_Pipe{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Pipes_Pipe) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Pipes_Pipe) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Pipes_Pipe) GetType() string {
	return AWS_Pipes_Pipe__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Pipes_Pipe) Set__LogicalName(v string) *AWS_Pipes_Pipe {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Pipes_Pipe) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Pipes_Pipe {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Pipes_Pipe) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Pipes_Pipe {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Pipes_Pipe) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Pipes_Pipe {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Pipes_Pipe) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Pipes_Pipe {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Pipes_Pipe) Set__RequestedOutputs(v []cfz.Output) *AWS_Pipes_Pipe {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Pipes_Pipe) Add__RequestedOutputs(v ...cfz.Output) *AWS_Pipes_Pipe {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Pipes_Pipe) Set__Description(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Pipes_Pipe) SetV__Description(v string) *AWS_Pipes_Pipe {
	t.Description = cfz.V(v)
	return t
}

// Set__DesiredState updates property "DesiredState".
func (t *AWS_Pipes_Pipe) Set__DesiredState(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.DesiredState = v
	return t
}

// SetV__DesiredState updates property "DesiredState".
func (t *AWS_Pipes_Pipe) SetV__DesiredState(v string) *AWS_Pipes_Pipe {
	t.DesiredState = cfz.V(v)
	return t
}

// Set__Enrichment updates property "Enrichment".
func (t *AWS_Pipes_Pipe) Set__Enrichment(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.Enrichment = v
	return t
}

// SetV__Enrichment updates property "Enrichment".
func (t *AWS_Pipes_Pipe) SetV__Enrichment(v string) *AWS_Pipes_Pipe {
	t.Enrichment = cfz.V(v)
	return t
}

// Set__EnrichmentParameters updates property "EnrichmentParameters".
func (t *AWS_Pipes_Pipe) Set__EnrichmentParameters(v cfz.Expression[AWS_Pipes_Pipe_PipeEnrichmentParameters]) *AWS_Pipes_Pipe {
	t.EnrichmentParameters = v
	return t
}

// SetV__EnrichmentParameters updates property "EnrichmentParameters".
func (t *AWS_Pipes_Pipe) SetV__EnrichmentParameters(v AWS_Pipes_Pipe_PipeEnrichmentParameters) *AWS_Pipes_Pipe {
	t.EnrichmentParameters = cfz.V(v)
	return t
}

// Set__KmsKeyIdentifier updates property "KmsKeyIdentifier".
func (t *AWS_Pipes_Pipe) Set__KmsKeyIdentifier(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.KmsKeyIdentifier = v
	return t
}

// SetV__KmsKeyIdentifier updates property "KmsKeyIdentifier".
func (t *AWS_Pipes_Pipe) SetV__KmsKeyIdentifier(v string) *AWS_Pipes_Pipe {
	t.KmsKeyIdentifier = cfz.V(v)
	return t
}

// Set__LogConfiguration updates property "LogConfiguration".
func (t *AWS_Pipes_Pipe) Set__LogConfiguration(v cfz.Expression[AWS_Pipes_Pipe_PipeLogConfiguration]) *AWS_Pipes_Pipe {
	t.LogConfiguration = v
	return t
}

// SetV__LogConfiguration updates property "LogConfiguration".
func (t *AWS_Pipes_Pipe) SetV__LogConfiguration(v AWS_Pipes_Pipe_PipeLogConfiguration) *AWS_Pipes_Pipe {
	t.LogConfiguration = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Pipes_Pipe) Set__Name(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Pipes_Pipe) SetV__Name(v string) *AWS_Pipes_Pipe {
	t.Name = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_Pipes_Pipe) Set__RoleArn(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_Pipes_Pipe) SetV__RoleArn(v string) *AWS_Pipes_Pipe {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__Source updates property "Source".
func (t *AWS_Pipes_Pipe) Set__Source(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.Source = v
	return t
}

// SetV__Source updates property "Source".
func (t *AWS_Pipes_Pipe) SetV__Source(v string) *AWS_Pipes_Pipe {
	t.Source = cfz.V(v)
	return t
}

// Set__SourceParameters updates property "SourceParameters".
func (t *AWS_Pipes_Pipe) Set__SourceParameters(v cfz.Expression[AWS_Pipes_Pipe_PipeSourceParameters]) *AWS_Pipes_Pipe {
	t.SourceParameters = v
	return t
}

// SetV__SourceParameters updates property "SourceParameters".
func (t *AWS_Pipes_Pipe) SetV__SourceParameters(v AWS_Pipes_Pipe_PipeSourceParameters) *AWS_Pipes_Pipe {
	t.SourceParameters = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Pipes_Pipe) Set__Tags(v cfz.ExpressionMap[string]) *AWS_Pipes_Pipe {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_Pipes_Pipe) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_Pipes_Pipe) SetMV__Tags(v ...map[string]string) *AWS_Pipes_Pipe {
	t.Tags = cfz.MV(v...)
	return t
}

// Set__Target updates property "Target".
func (t *AWS_Pipes_Pipe) Set__Target(v cfz.Expression[string]) *AWS_Pipes_Pipe {
	t.Target = v
	return t
}

// SetV__Target updates property "Target".
func (t *AWS_Pipes_Pipe) SetV__Target(v string) *AWS_Pipes_Pipe {
	t.Target = cfz.V(v)
	return t
}

// Set__TargetParameters updates property "TargetParameters".
func (t *AWS_Pipes_Pipe) Set__TargetParameters(v cfz.Expression[AWS_Pipes_Pipe_PipeTargetParameters]) *AWS_Pipes_Pipe {
	t.TargetParameters = v
	return t
}

// SetV__TargetParameters updates property "TargetParameters".
func (t *AWS_Pipes_Pipe) SetV__TargetParameters(v AWS_Pipes_Pipe_PipeTargetParameters) *AWS_Pipes_Pipe {
	t.TargetParameters = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Pipes_Pipe) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Pipes_Pipe) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Pipes_Pipe__AttributesMap.Arn))
}

// GetAtt__CreationTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTime
func (t *AWS_Pipes_Pipe) GetAtt__CreationTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Pipes_Pipe__AttributesMap.CreationTime))
}

// GetAtt__CurrentState returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CurrentState
func (t *AWS_Pipes_Pipe) GetAtt__CurrentState() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Pipes_Pipe__AttributesMap.CurrentState))
}

// GetAtt__LastModifiedTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LastModifiedTime
func (t *AWS_Pipes_Pipe) GetAtt__LastModifiedTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Pipes_Pipe__AttributesMap.LastModifiedTime))
}

// GetAtt__StateReason returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: StateReason
func (t *AWS_Pipes_Pipe) GetAtt__StateReason() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Pipes_Pipe__AttributesMap.StateReason))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Pipes_Pipe) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Pipes_Pipe) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTime
func (t *AWS_Pipes_Pipe) GetConventionalOutputAtt__CreationTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTime", t.GetAtt__CreationTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CurrentState returns a conventionally configured output for an attribute of this resource.
// Attribute: CurrentState
func (t *AWS_Pipes_Pipe) GetConventionalOutputAtt__CurrentState(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCurrentState", t.GetAtt__CurrentState())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LastModifiedTime returns a conventionally configured output for an attribute of this resource.
// Attribute: LastModifiedTime
func (t *AWS_Pipes_Pipe) GetConventionalOutputAtt__LastModifiedTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLastModifiedTime", t.GetAtt__LastModifiedTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__StateReason returns a conventionally configured output for an attribute of this resource.
// Attribute: StateReason
func (t *AWS_Pipes_Pipe) GetConventionalOutputAtt__StateReason(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStateReason", t.GetAtt__StateReason())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Pipes_Pipe) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Pipes_Pipe

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

func (t *AWS_Pipes_Pipe) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
