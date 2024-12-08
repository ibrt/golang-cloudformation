// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediaconvert

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MediaConvert_Preset)(nil)
	_ cfz.Resource                   = (*AWS_MediaConvert_Preset)(nil)
)

const (
	// AWS_MediaConvert_Preset__Type is the CloudFormation type for AWS::MediaConvert::Preset.
	AWS_MediaConvert_Preset__Type = "AWS::MediaConvert::Preset"
)

var (
	// AWS_MediaConvert_Preset__AttributesMap reports all the CloudFormation attributes for AWS::MediaConvert::Preset.
	AWS_MediaConvert_Preset__AttributesMap = struct {
		Arn  string
		Name string
	}{
		Arn:  "Arn",
		Name: "Name",
	}

	// AWS_MediaConvert_Preset__AttributesSlice reports all the CloudFormation attributes for AWS::MediaConvert::Preset.
	AWS_MediaConvert_Preset__AttributesSlice = []string{
		AWS_MediaConvert_Preset__AttributesMap.Arn,
		AWS_MediaConvert_Preset__AttributesMap.Name,
	}
)

var (
	// AWS_MediaConvert_Preset__PropertiesMap reports all the CloudFormation properties for AWS::MediaConvert::Preset.
	AWS_MediaConvert_Preset__PropertiesMap = struct {
		Category     string
		Description  string
		Name         string
		SettingsJson string
		Tags         string
	}{
		Category:     "Category",
		Description:  "Description",
		Name:         "Name",
		SettingsJson: "SettingsJson",
		Tags:         "Tags",
	}

	// AWS_MediaConvert_Preset__PropertiesSlice reports all the CloudFormation properties for AWS::MediaConvert::Preset.
	AWS_MediaConvert_Preset__PropertiesSlice = []string{
		AWS_MediaConvert_Preset__PropertiesMap.Category,
		AWS_MediaConvert_Preset__PropertiesMap.Description,
		AWS_MediaConvert_Preset__PropertiesMap.Name,
		AWS_MediaConvert_Preset__PropertiesMap.SettingsJson,
		AWS_MediaConvert_Preset__PropertiesMap.Tags,
	}
)

// AWS_MediaConvert_Preset is a binding for AWS::MediaConvert::Preset.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html
type AWS_MediaConvert_Preset struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Category is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-category
	Category cfz.Expression[string] `json:"Category,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// SettingsJson is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-settingsjson
	SettingsJson cfz.Expression[json.RawMessage] `json:"SettingsJson,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-preset.html#cfn-mediaconvert-preset-tags
	Tags cfz.Expression[json.RawMessage] `json:"Tags,omitempty"`
}

// New__AWS_MediaConvert_Preset initializes a new *AWS_MediaConvert_Preset.
func New__AWS_MediaConvert_Preset(logicalName string) *AWS_MediaConvert_Preset {
	return &AWS_MediaConvert_Preset{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MediaConvert_Preset) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MediaConvert_Preset) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MediaConvert_Preset) GetType() string {
	return AWS_MediaConvert_Preset__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MediaConvert_Preset) Set__LogicalName(v string) *AWS_MediaConvert_Preset {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MediaConvert_Preset) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MediaConvert_Preset {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MediaConvert_Preset) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MediaConvert_Preset {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MediaConvert_Preset) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MediaConvert_Preset {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MediaConvert_Preset) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MediaConvert_Preset {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MediaConvert_Preset) Set__RequestedOutputs(v []cfz.Output) *AWS_MediaConvert_Preset {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MediaConvert_Preset) Add__RequestedOutputs(v ...cfz.Output) *AWS_MediaConvert_Preset {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Category updates property "Category".
func (t *AWS_MediaConvert_Preset) Set__Category(v cfz.Expression[string]) *AWS_MediaConvert_Preset {
	t.Category = v
	return t
}

// SetV__Category updates property "Category".
func (t *AWS_MediaConvert_Preset) SetV__Category(v string) *AWS_MediaConvert_Preset {
	t.Category = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_MediaConvert_Preset) Set__Description(v cfz.Expression[string]) *AWS_MediaConvert_Preset {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_MediaConvert_Preset) SetV__Description(v string) *AWS_MediaConvert_Preset {
	t.Description = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_MediaConvert_Preset) Set__Name(v cfz.Expression[string]) *AWS_MediaConvert_Preset {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_MediaConvert_Preset) SetV__Name(v string) *AWS_MediaConvert_Preset {
	t.Name = cfz.V(v)
	return t
}

// Set__SettingsJson updates property "SettingsJson".
func (t *AWS_MediaConvert_Preset) Set__SettingsJson(v cfz.Expression[json.RawMessage]) *AWS_MediaConvert_Preset {
	t.SettingsJson = v
	return t
}

// SetV__SettingsJson updates property "SettingsJson".
func (t *AWS_MediaConvert_Preset) SetV__SettingsJson(v json.RawMessage) *AWS_MediaConvert_Preset {
	t.SettingsJson = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MediaConvert_Preset) Set__Tags(v cfz.Expression[json.RawMessage]) *AWS_MediaConvert_Preset {
	t.Tags = v
	return t
}

// SetV__Tags updates property "Tags".
func (t *AWS_MediaConvert_Preset) SetV__Tags(v json.RawMessage) *AWS_MediaConvert_Preset {
	t.Tags = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MediaConvert_Preset) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_MediaConvert_Preset) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaConvert_Preset__AttributesMap.Arn))
}

// GetAtt__Name returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Name
func (t *AWS_MediaConvert_Preset) GetAtt__Name() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaConvert_Preset__AttributesMap.Name))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MediaConvert_Preset) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_MediaConvert_Preset) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Name returns a conventionally configured output for an attribute of this resource.
// Attribute: Name
func (t *AWS_MediaConvert_Preset) GetConventionalOutputAtt__Name(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttName", t.GetAtt__Name())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MediaConvert_Preset) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MediaConvert_Preset

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

func (t *AWS_MediaConvert_Preset) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
