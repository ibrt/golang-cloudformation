// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_rds

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_RDS_OptionGroup)(nil)
	_ cfz.Resource                   = (*AWS_RDS_OptionGroup)(nil)
)

const (
	// AWS_RDS_OptionGroup__Type is the CloudFormation type for AWS::RDS::OptionGroup.
	AWS_RDS_OptionGroup__Type = "AWS::RDS::OptionGroup"
)

var (
	// AWS_RDS_OptionGroup__PropertiesMap reports all the CloudFormation properties for AWS::RDS::OptionGroup.
	AWS_RDS_OptionGroup__PropertiesMap = struct {
		EngineName             string
		MajorEngineVersion     string
		OptionConfigurations   string
		OptionGroupDescription string
		OptionGroupName        string
		Tags                   string
	}{
		EngineName:             "EngineName",
		MajorEngineVersion:     "MajorEngineVersion",
		OptionConfigurations:   "OptionConfigurations",
		OptionGroupDescription: "OptionGroupDescription",
		OptionGroupName:        "OptionGroupName",
		Tags:                   "Tags",
	}

	// AWS_RDS_OptionGroup__PropertiesSlice reports all the CloudFormation properties for AWS::RDS::OptionGroup.
	AWS_RDS_OptionGroup__PropertiesSlice = []string{
		AWS_RDS_OptionGroup__PropertiesMap.EngineName,
		AWS_RDS_OptionGroup__PropertiesMap.MajorEngineVersion,
		AWS_RDS_OptionGroup__PropertiesMap.OptionConfigurations,
		AWS_RDS_OptionGroup__PropertiesMap.OptionGroupDescription,
		AWS_RDS_OptionGroup__PropertiesMap.OptionGroupName,
		AWS_RDS_OptionGroup__PropertiesMap.Tags,
	}
)

// AWS_RDS_OptionGroup is a binding for AWS::RDS::OptionGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
type AWS_RDS_OptionGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// EngineName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-enginename
	EngineName cfz.Expression[string] `json:"EngineName,omitempty"`

	// MajorEngineVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-majorengineversion
	MajorEngineVersion cfz.Expression[string] `json:"MajorEngineVersion,omitempty"`

	// OptionConfigurations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optionconfigurations
	OptionConfigurations cfz.ExpressionSlice[AWS_RDS_OptionGroup_OptionConfiguration] `json:"OptionConfigurations,omitempty"`

	// OptionGroupDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optiongroupdescription
	OptionGroupDescription cfz.Expression[string] `json:"OptionGroupDescription,omitempty"`

	// OptionGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optiongroupname
	OptionGroupName cfz.Expression[string] `json:"OptionGroupName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_RDS_OptionGroup initializes a new *AWS_RDS_OptionGroup.
func New__AWS_RDS_OptionGroup(logicalName string) *AWS_RDS_OptionGroup {
	return &AWS_RDS_OptionGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_RDS_OptionGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_RDS_OptionGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_RDS_OptionGroup) GetType() string {
	return AWS_RDS_OptionGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_RDS_OptionGroup) Set__LogicalName(v string) *AWS_RDS_OptionGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_RDS_OptionGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_RDS_OptionGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_RDS_OptionGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_RDS_OptionGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_RDS_OptionGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_RDS_OptionGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_RDS_OptionGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_RDS_OptionGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_RDS_OptionGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_RDS_OptionGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_RDS_OptionGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_RDS_OptionGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__EngineName updates property "EngineName".
func (t *AWS_RDS_OptionGroup) Set__EngineName(v cfz.Expression[string]) *AWS_RDS_OptionGroup {
	t.EngineName = v
	return t
}

// SetV__EngineName updates property "EngineName".
func (t *AWS_RDS_OptionGroup) SetV__EngineName(v string) *AWS_RDS_OptionGroup {
	t.EngineName = cfz.V(v)
	return t
}

// Set__MajorEngineVersion updates property "MajorEngineVersion".
func (t *AWS_RDS_OptionGroup) Set__MajorEngineVersion(v cfz.Expression[string]) *AWS_RDS_OptionGroup {
	t.MajorEngineVersion = v
	return t
}

// SetV__MajorEngineVersion updates property "MajorEngineVersion".
func (t *AWS_RDS_OptionGroup) SetV__MajorEngineVersion(v string) *AWS_RDS_OptionGroup {
	t.MajorEngineVersion = cfz.V(v)
	return t
}

// Set__OptionConfigurations updates property "OptionConfigurations".
func (t *AWS_RDS_OptionGroup) Set__OptionConfigurations(v cfz.ExpressionSlice[AWS_RDS_OptionGroup_OptionConfiguration]) *AWS_RDS_OptionGroup {
	t.OptionConfigurations = v
	return t
}

// SetS__OptionConfigurations updates property "OptionConfigurations".
func (t *AWS_RDS_OptionGroup) SetS__OptionConfigurations(v ...cfz.Expression[AWS_RDS_OptionGroup_OptionConfiguration]) *AWS_RDS_OptionGroup {
	t.OptionConfigurations = cfz.S(v...)
	return t
}

// SetSV__OptionConfigurations updates property "OptionConfigurations".
func (t *AWS_RDS_OptionGroup) SetSV__OptionConfigurations(v ...AWS_RDS_OptionGroup_OptionConfiguration) *AWS_RDS_OptionGroup {
	t.OptionConfigurations = cfz.SV(v...)
	return t
}

// Set__OptionGroupDescription updates property "OptionGroupDescription".
func (t *AWS_RDS_OptionGroup) Set__OptionGroupDescription(v cfz.Expression[string]) *AWS_RDS_OptionGroup {
	t.OptionGroupDescription = v
	return t
}

// SetV__OptionGroupDescription updates property "OptionGroupDescription".
func (t *AWS_RDS_OptionGroup) SetV__OptionGroupDescription(v string) *AWS_RDS_OptionGroup {
	t.OptionGroupDescription = cfz.V(v)
	return t
}

// Set__OptionGroupName updates property "OptionGroupName".
func (t *AWS_RDS_OptionGroup) Set__OptionGroupName(v cfz.Expression[string]) *AWS_RDS_OptionGroup {
	t.OptionGroupName = v
	return t
}

// SetV__OptionGroupName updates property "OptionGroupName".
func (t *AWS_RDS_OptionGroup) SetV__OptionGroupName(v string) *AWS_RDS_OptionGroup {
	t.OptionGroupName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_RDS_OptionGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_RDS_OptionGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_RDS_OptionGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_RDS_OptionGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_RDS_OptionGroup) SetSV__Tags(v ...cfz.Tag) *AWS_RDS_OptionGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_RDS_OptionGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_RDS_OptionGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_RDS_OptionGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_RDS_OptionGroup

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

func (t *AWS_RDS_OptionGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
