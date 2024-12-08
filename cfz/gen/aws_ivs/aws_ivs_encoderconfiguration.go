// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ivs

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IVS_EncoderConfiguration)(nil)
	_ cfz.Resource                   = (*AWS_IVS_EncoderConfiguration)(nil)
)

const (
	// AWS_IVS_EncoderConfiguration__Type is the CloudFormation type for AWS::IVS::EncoderConfiguration.
	AWS_IVS_EncoderConfiguration__Type = "AWS::IVS::EncoderConfiguration"
)

var (
	// AWS_IVS_EncoderConfiguration__AttributesMap reports all the CloudFormation attributes for AWS::IVS::EncoderConfiguration.
	AWS_IVS_EncoderConfiguration__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_IVS_EncoderConfiguration__AttributesSlice reports all the CloudFormation attributes for AWS::IVS::EncoderConfiguration.
	AWS_IVS_EncoderConfiguration__AttributesSlice = []string{
		AWS_IVS_EncoderConfiguration__AttributesMap.Arn,
	}
)

var (
	// AWS_IVS_EncoderConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::IVS::EncoderConfiguration.
	AWS_IVS_EncoderConfiguration__PropertiesMap = struct {
		Name  string
		Tags  string
		Video string
	}{
		Name:  "Name",
		Tags:  "Tags",
		Video: "Video",
	}

	// AWS_IVS_EncoderConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::IVS::EncoderConfiguration.
	AWS_IVS_EncoderConfiguration__PropertiesSlice = []string{
		AWS_IVS_EncoderConfiguration__PropertiesMap.Name,
		AWS_IVS_EncoderConfiguration__PropertiesMap.Tags,
		AWS_IVS_EncoderConfiguration__PropertiesMap.Video,
	}
)

// AWS_IVS_EncoderConfiguration is a binding for AWS::IVS::EncoderConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-encoderconfiguration.html
type AWS_IVS_EncoderConfiguration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-encoderconfiguration.html#cfn-ivs-encoderconfiguration-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-encoderconfiguration.html#cfn-ivs-encoderconfiguration-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// Video is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-encoderconfiguration.html#cfn-ivs-encoderconfiguration-video
	Video cfz.Expression[AWS_IVS_EncoderConfiguration_Video] `json:"Video,omitempty"`
}

// New__AWS_IVS_EncoderConfiguration initializes a new *AWS_IVS_EncoderConfiguration.
func New__AWS_IVS_EncoderConfiguration(logicalName string) *AWS_IVS_EncoderConfiguration {
	return &AWS_IVS_EncoderConfiguration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IVS_EncoderConfiguration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IVS_EncoderConfiguration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IVS_EncoderConfiguration) GetType() string {
	return AWS_IVS_EncoderConfiguration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IVS_EncoderConfiguration) Set__LogicalName(v string) *AWS_IVS_EncoderConfiguration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IVS_EncoderConfiguration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IVS_EncoderConfiguration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IVS_EncoderConfiguration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IVS_EncoderConfiguration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IVS_EncoderConfiguration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IVS_EncoderConfiguration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IVS_EncoderConfiguration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IVS_EncoderConfiguration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IVS_EncoderConfiguration) Set__RequestedOutputs(v []cfz.Output) *AWS_IVS_EncoderConfiguration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IVS_EncoderConfiguration) Add__RequestedOutputs(v ...cfz.Output) *AWS_IVS_EncoderConfiguration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_IVS_EncoderConfiguration) Set__Name(v cfz.Expression[string]) *AWS_IVS_EncoderConfiguration {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_IVS_EncoderConfiguration) SetV__Name(v string) *AWS_IVS_EncoderConfiguration {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IVS_EncoderConfiguration) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IVS_EncoderConfiguration {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IVS_EncoderConfiguration) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IVS_EncoderConfiguration {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IVS_EncoderConfiguration) SetSV__Tags(v ...cfz.Tag) *AWS_IVS_EncoderConfiguration {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__Video updates property "Video".
func (t *AWS_IVS_EncoderConfiguration) Set__Video(v cfz.Expression[AWS_IVS_EncoderConfiguration_Video]) *AWS_IVS_EncoderConfiguration {
	t.Video = v
	return t
}

// SetV__Video updates property "Video".
func (t *AWS_IVS_EncoderConfiguration) SetV__Video(v AWS_IVS_EncoderConfiguration_Video) *AWS_IVS_EncoderConfiguration {
	t.Video = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IVS_EncoderConfiguration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_IVS_EncoderConfiguration) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IVS_EncoderConfiguration__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IVS_EncoderConfiguration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_IVS_EncoderConfiguration) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IVS_EncoderConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IVS_EncoderConfiguration

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

func (t *AWS_IVS_EncoderConfiguration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
