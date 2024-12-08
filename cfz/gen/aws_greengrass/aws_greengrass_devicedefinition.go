// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Greengrass_DeviceDefinition)(nil)
	_ cfz.Resource                   = (*AWS_Greengrass_DeviceDefinition)(nil)
)

const (
	// AWS_Greengrass_DeviceDefinition__Type is the CloudFormation type for AWS::Greengrass::DeviceDefinition.
	AWS_Greengrass_DeviceDefinition__Type = "AWS::Greengrass::DeviceDefinition"
)

var (
	// AWS_Greengrass_DeviceDefinition__AttributesMap reports all the CloudFormation attributes for AWS::Greengrass::DeviceDefinition.
	AWS_Greengrass_DeviceDefinition__AttributesMap = struct {
		Arn              string
		Id               string
		LatestVersionArn string
		Name             string
	}{
		Arn:              "Arn",
		Id:               "Id",
		LatestVersionArn: "LatestVersionArn",
		Name:             "Name",
	}

	// AWS_Greengrass_DeviceDefinition__AttributesSlice reports all the CloudFormation attributes for AWS::Greengrass::DeviceDefinition.
	AWS_Greengrass_DeviceDefinition__AttributesSlice = []string{
		AWS_Greengrass_DeviceDefinition__AttributesMap.Arn,
		AWS_Greengrass_DeviceDefinition__AttributesMap.Id,
		AWS_Greengrass_DeviceDefinition__AttributesMap.LatestVersionArn,
		AWS_Greengrass_DeviceDefinition__AttributesMap.Name,
	}
)

var (
	// AWS_Greengrass_DeviceDefinition__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::DeviceDefinition.
	AWS_Greengrass_DeviceDefinition__PropertiesMap = struct {
		InitialVersion string
		Name           string
		Tags           string
	}{
		InitialVersion: "InitialVersion",
		Name:           "Name",
		Tags:           "Tags",
	}

	// AWS_Greengrass_DeviceDefinition__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::DeviceDefinition.
	AWS_Greengrass_DeviceDefinition__PropertiesSlice = []string{
		AWS_Greengrass_DeviceDefinition__PropertiesMap.InitialVersion,
		AWS_Greengrass_DeviceDefinition__PropertiesMap.Name,
		AWS_Greengrass_DeviceDefinition__PropertiesMap.Tags,
	}
)

// AWS_Greengrass_DeviceDefinition is a binding for AWS::Greengrass::DeviceDefinition.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html
type AWS_Greengrass_DeviceDefinition struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// InitialVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html#cfn-greengrass-devicedefinition-initialversion
	InitialVersion cfz.Expression[AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion] `json:"InitialVersion,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html#cfn-greengrass-devicedefinition-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html#cfn-greengrass-devicedefinition-tags
	Tags cfz.Expression[json.RawMessage] `json:"Tags,omitempty"`
}

// New__AWS_Greengrass_DeviceDefinition initializes a new *AWS_Greengrass_DeviceDefinition.
func New__AWS_Greengrass_DeviceDefinition(logicalName string) *AWS_Greengrass_DeviceDefinition {
	return &AWS_Greengrass_DeviceDefinition{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Greengrass_DeviceDefinition) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Greengrass_DeviceDefinition) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Greengrass_DeviceDefinition) GetType() string {
	return AWS_Greengrass_DeviceDefinition__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Greengrass_DeviceDefinition) Set__LogicalName(v string) *AWS_Greengrass_DeviceDefinition {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Greengrass_DeviceDefinition) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Greengrass_DeviceDefinition {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Greengrass_DeviceDefinition) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Greengrass_DeviceDefinition {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Greengrass_DeviceDefinition) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Greengrass_DeviceDefinition {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Greengrass_DeviceDefinition) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Greengrass_DeviceDefinition {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Greengrass_DeviceDefinition) Set__RequestedOutputs(v []cfz.Output) *AWS_Greengrass_DeviceDefinition {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Greengrass_DeviceDefinition) Add__RequestedOutputs(v ...cfz.Output) *AWS_Greengrass_DeviceDefinition {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__InitialVersion updates property "InitialVersion".
func (t *AWS_Greengrass_DeviceDefinition) Set__InitialVersion(v cfz.Expression[AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion]) *AWS_Greengrass_DeviceDefinition {
	t.InitialVersion = v
	return t
}

// SetV__InitialVersion updates property "InitialVersion".
func (t *AWS_Greengrass_DeviceDefinition) SetV__InitialVersion(v AWS_Greengrass_DeviceDefinition_DeviceDefinitionVersion) *AWS_Greengrass_DeviceDefinition {
	t.InitialVersion = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Greengrass_DeviceDefinition) Set__Name(v cfz.Expression[string]) *AWS_Greengrass_DeviceDefinition {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Greengrass_DeviceDefinition) SetV__Name(v string) *AWS_Greengrass_DeviceDefinition {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Greengrass_DeviceDefinition) Set__Tags(v cfz.Expression[json.RawMessage]) *AWS_Greengrass_DeviceDefinition {
	t.Tags = v
	return t
}

// SetV__Tags updates property "Tags".
func (t *AWS_Greengrass_DeviceDefinition) SetV__Tags(v json.RawMessage) *AWS_Greengrass_DeviceDefinition {
	t.Tags = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Greengrass_DeviceDefinition) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Greengrass_DeviceDefinition) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Greengrass_DeviceDefinition__AttributesMap.Arn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Greengrass_DeviceDefinition) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Greengrass_DeviceDefinition__AttributesMap.Id))
}

// GetAtt__LatestVersionArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LatestVersionArn
func (t *AWS_Greengrass_DeviceDefinition) GetAtt__LatestVersionArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Greengrass_DeviceDefinition__AttributesMap.LatestVersionArn))
}

// GetAtt__Name returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Name
func (t *AWS_Greengrass_DeviceDefinition) GetAtt__Name() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Greengrass_DeviceDefinition__AttributesMap.Name))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Greengrass_DeviceDefinition) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Greengrass_DeviceDefinition) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Greengrass_DeviceDefinition) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LatestVersionArn returns a conventionally configured output for an attribute of this resource.
// Attribute: LatestVersionArn
func (t *AWS_Greengrass_DeviceDefinition) GetConventionalOutputAtt__LatestVersionArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLatestVersionArn", t.GetAtt__LatestVersionArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Name returns a conventionally configured output for an attribute of this resource.
// Attribute: Name
func (t *AWS_Greengrass_DeviceDefinition) GetConventionalOutputAtt__Name(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttName", t.GetAtt__Name())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Greengrass_DeviceDefinition) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Greengrass_DeviceDefinition

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

func (t *AWS_Greengrass_DeviceDefinition) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
