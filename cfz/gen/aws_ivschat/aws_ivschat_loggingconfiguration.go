// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ivschat

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IVSChat_LoggingConfiguration)(nil)
	_ cfz.Resource                   = (*AWS_IVSChat_LoggingConfiguration)(nil)
)

const (
	// AWS_IVSChat_LoggingConfiguration__Type is the CloudFormation type for AWS::IVSChat::LoggingConfiguration.
	AWS_IVSChat_LoggingConfiguration__Type = "AWS::IVSChat::LoggingConfiguration"
)

var (
	// AWS_IVSChat_LoggingConfiguration__AttributesMap reports all the CloudFormation attributes for AWS::IVSChat::LoggingConfiguration.
	AWS_IVSChat_LoggingConfiguration__AttributesMap = struct {
		Arn   string
		Id    string
		State string
	}{
		Arn:   "Arn",
		Id:    "Id",
		State: "State",
	}

	// AWS_IVSChat_LoggingConfiguration__AttributesSlice reports all the CloudFormation attributes for AWS::IVSChat::LoggingConfiguration.
	AWS_IVSChat_LoggingConfiguration__AttributesSlice = []string{
		AWS_IVSChat_LoggingConfiguration__AttributesMap.Arn,
		AWS_IVSChat_LoggingConfiguration__AttributesMap.Id,
		AWS_IVSChat_LoggingConfiguration__AttributesMap.State,
	}
)

var (
	// AWS_IVSChat_LoggingConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::IVSChat::LoggingConfiguration.
	AWS_IVSChat_LoggingConfiguration__PropertiesMap = struct {
		DestinationConfiguration string
		Name                     string
		Tags                     string
	}{
		DestinationConfiguration: "DestinationConfiguration",
		Name:                     "Name",
		Tags:                     "Tags",
	}

	// AWS_IVSChat_LoggingConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::IVSChat::LoggingConfiguration.
	AWS_IVSChat_LoggingConfiguration__PropertiesSlice = []string{
		AWS_IVSChat_LoggingConfiguration__PropertiesMap.DestinationConfiguration,
		AWS_IVSChat_LoggingConfiguration__PropertiesMap.Name,
		AWS_IVSChat_LoggingConfiguration__PropertiesMap.Tags,
	}
)

// AWS_IVSChat_LoggingConfiguration is a binding for AWS::IVSChat::LoggingConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-loggingconfiguration.html
type AWS_IVSChat_LoggingConfiguration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DestinationConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-loggingconfiguration.html#cfn-ivschat-loggingconfiguration-destinationconfiguration
	DestinationConfiguration cfz.Expression[AWS_IVSChat_LoggingConfiguration_DestinationConfiguration] `json:"DestinationConfiguration,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-loggingconfiguration.html#cfn-ivschat-loggingconfiguration-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivschat-loggingconfiguration.html#cfn-ivschat-loggingconfiguration-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_IVSChat_LoggingConfiguration initializes a new *AWS_IVSChat_LoggingConfiguration.
func New__AWS_IVSChat_LoggingConfiguration(logicalName string) *AWS_IVSChat_LoggingConfiguration {
	return &AWS_IVSChat_LoggingConfiguration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IVSChat_LoggingConfiguration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IVSChat_LoggingConfiguration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IVSChat_LoggingConfiguration) GetType() string {
	return AWS_IVSChat_LoggingConfiguration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IVSChat_LoggingConfiguration) Set__LogicalName(v string) *AWS_IVSChat_LoggingConfiguration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IVSChat_LoggingConfiguration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IVSChat_LoggingConfiguration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IVSChat_LoggingConfiguration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IVSChat_LoggingConfiguration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IVSChat_LoggingConfiguration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IVSChat_LoggingConfiguration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IVSChat_LoggingConfiguration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IVSChat_LoggingConfiguration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IVSChat_LoggingConfiguration) Set__RequestedOutputs(v []cfz.Output) *AWS_IVSChat_LoggingConfiguration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IVSChat_LoggingConfiguration) Add__RequestedOutputs(v ...cfz.Output) *AWS_IVSChat_LoggingConfiguration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DestinationConfiguration updates property "DestinationConfiguration".
func (t *AWS_IVSChat_LoggingConfiguration) Set__DestinationConfiguration(v cfz.Expression[AWS_IVSChat_LoggingConfiguration_DestinationConfiguration]) *AWS_IVSChat_LoggingConfiguration {
	t.DestinationConfiguration = v
	return t
}

// SetV__DestinationConfiguration updates property "DestinationConfiguration".
func (t *AWS_IVSChat_LoggingConfiguration) SetV__DestinationConfiguration(v AWS_IVSChat_LoggingConfiguration_DestinationConfiguration) *AWS_IVSChat_LoggingConfiguration {
	t.DestinationConfiguration = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_IVSChat_LoggingConfiguration) Set__Name(v cfz.Expression[string]) *AWS_IVSChat_LoggingConfiguration {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_IVSChat_LoggingConfiguration) SetV__Name(v string) *AWS_IVSChat_LoggingConfiguration {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IVSChat_LoggingConfiguration) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IVSChat_LoggingConfiguration {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IVSChat_LoggingConfiguration) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IVSChat_LoggingConfiguration {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IVSChat_LoggingConfiguration) SetSV__Tags(v ...cfz.Tag) *AWS_IVSChat_LoggingConfiguration {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IVSChat_LoggingConfiguration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_IVSChat_LoggingConfiguration) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IVSChat_LoggingConfiguration__AttributesMap.Arn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_IVSChat_LoggingConfiguration) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IVSChat_LoggingConfiguration__AttributesMap.Id))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_IVSChat_LoggingConfiguration) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IVSChat_LoggingConfiguration__AttributesMap.State))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IVSChat_LoggingConfiguration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_IVSChat_LoggingConfiguration) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_IVSChat_LoggingConfiguration) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_IVSChat_LoggingConfiguration) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IVSChat_LoggingConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IVSChat_LoggingConfiguration

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

func (t *AWS_IVSChat_LoggingConfiguration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
