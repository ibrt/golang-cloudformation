// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_glue

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Glue_SecurityConfiguration)(nil)
	_ cfz.Resource                   = (*AWS_Glue_SecurityConfiguration)(nil)
)

const (
	// AWS_Glue_SecurityConfiguration__Type is the CloudFormation type for AWS::Glue::SecurityConfiguration.
	AWS_Glue_SecurityConfiguration__Type = "AWS::Glue::SecurityConfiguration"
)

var (
	// AWS_Glue_SecurityConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Glue::SecurityConfiguration.
	AWS_Glue_SecurityConfiguration__PropertiesMap = struct {
		EncryptionConfiguration string
		Name                    string
	}{
		EncryptionConfiguration: "EncryptionConfiguration",
		Name:                    "Name",
	}

	// AWS_Glue_SecurityConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Glue::SecurityConfiguration.
	AWS_Glue_SecurityConfiguration__PropertiesSlice = []string{
		AWS_Glue_SecurityConfiguration__PropertiesMap.EncryptionConfiguration,
		AWS_Glue_SecurityConfiguration__PropertiesMap.Name,
	}
)

// AWS_Glue_SecurityConfiguration is a binding for AWS::Glue::SecurityConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html
type AWS_Glue_SecurityConfiguration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// EncryptionConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration
	EncryptionConfiguration cfz.Expression[AWS_Glue_SecurityConfiguration_EncryptionConfiguration] `json:"EncryptionConfiguration,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-securityconfiguration.html#cfn-glue-securityconfiguration-name
	Name cfz.Expression[string] `json:"Name,omitempty"`
}

// New__AWS_Glue_SecurityConfiguration initializes a new *AWS_Glue_SecurityConfiguration.
func New__AWS_Glue_SecurityConfiguration(logicalName string) *AWS_Glue_SecurityConfiguration {
	return &AWS_Glue_SecurityConfiguration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Glue_SecurityConfiguration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Glue_SecurityConfiguration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Glue_SecurityConfiguration) GetType() string {
	return AWS_Glue_SecurityConfiguration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Glue_SecurityConfiguration) Set__LogicalName(v string) *AWS_Glue_SecurityConfiguration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Glue_SecurityConfiguration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Glue_SecurityConfiguration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Glue_SecurityConfiguration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Glue_SecurityConfiguration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Glue_SecurityConfiguration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Glue_SecurityConfiguration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Glue_SecurityConfiguration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Glue_SecurityConfiguration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Glue_SecurityConfiguration) Set__RequestedOutputs(v []cfz.Output) *AWS_Glue_SecurityConfiguration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Glue_SecurityConfiguration) Add__RequestedOutputs(v ...cfz.Output) *AWS_Glue_SecurityConfiguration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__EncryptionConfiguration updates property "EncryptionConfiguration".
func (t *AWS_Glue_SecurityConfiguration) Set__EncryptionConfiguration(v cfz.Expression[AWS_Glue_SecurityConfiguration_EncryptionConfiguration]) *AWS_Glue_SecurityConfiguration {
	t.EncryptionConfiguration = v
	return t
}

// SetV__EncryptionConfiguration updates property "EncryptionConfiguration".
func (t *AWS_Glue_SecurityConfiguration) SetV__EncryptionConfiguration(v AWS_Glue_SecurityConfiguration_EncryptionConfiguration) *AWS_Glue_SecurityConfiguration {
	t.EncryptionConfiguration = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Glue_SecurityConfiguration) Set__Name(v cfz.Expression[string]) *AWS_Glue_SecurityConfiguration {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Glue_SecurityConfiguration) SetV__Name(v string) *AWS_Glue_SecurityConfiguration {
	t.Name = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Glue_SecurityConfiguration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Glue_SecurityConfiguration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Glue_SecurityConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Glue_SecurityConfiguration

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

func (t *AWS_Glue_SecurityConfiguration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
