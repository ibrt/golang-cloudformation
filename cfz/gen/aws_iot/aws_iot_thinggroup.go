// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IoT_ThingGroup)(nil)
	_ cfz.Resource                   = (*AWS_IoT_ThingGroup)(nil)
)

const (
	// AWS_IoT_ThingGroup__Type is the CloudFormation type for AWS::IoT::ThingGroup.
	AWS_IoT_ThingGroup__Type = "AWS::IoT::ThingGroup"
)

var (
	// AWS_IoT_ThingGroup__AttributesMap reports all the CloudFormation attributes for AWS::IoT::ThingGroup.
	AWS_IoT_ThingGroup__AttributesMap = struct {
		Arn string
		Id  string
	}{
		Arn: "Arn",
		Id:  "Id",
	}

	// AWS_IoT_ThingGroup__AttributesSlice reports all the CloudFormation attributes for AWS::IoT::ThingGroup.
	AWS_IoT_ThingGroup__AttributesSlice = []string{
		AWS_IoT_ThingGroup__AttributesMap.Arn,
		AWS_IoT_ThingGroup__AttributesMap.Id,
	}
)

var (
	// AWS_IoT_ThingGroup__PropertiesMap reports all the CloudFormation properties for AWS::IoT::ThingGroup.
	AWS_IoT_ThingGroup__PropertiesMap = struct {
		ParentGroupName      string
		QueryString          string
		Tags                 string
		ThingGroupName       string
		ThingGroupProperties string
	}{
		ParentGroupName:      "ParentGroupName",
		QueryString:          "QueryString",
		Tags:                 "Tags",
		ThingGroupName:       "ThingGroupName",
		ThingGroupProperties: "ThingGroupProperties",
	}

	// AWS_IoT_ThingGroup__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::ThingGroup.
	AWS_IoT_ThingGroup__PropertiesSlice = []string{
		AWS_IoT_ThingGroup__PropertiesMap.ParentGroupName,
		AWS_IoT_ThingGroup__PropertiesMap.QueryString,
		AWS_IoT_ThingGroup__PropertiesMap.Tags,
		AWS_IoT_ThingGroup__PropertiesMap.ThingGroupName,
		AWS_IoT_ThingGroup__PropertiesMap.ThingGroupProperties,
	}
)

// AWS_IoT_ThingGroup is a binding for AWS::IoT::ThingGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html
type AWS_IoT_ThingGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ParentGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-parentgroupname
	ParentGroupName cfz.Expression[string] `json:"ParentGroupName,omitempty"`

	// QueryString is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-querystring
	QueryString cfz.Expression[string] `json:"QueryString,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// ThingGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-thinggroupname
	ThingGroupName cfz.Expression[string] `json:"ThingGroupName,omitempty"`

	// ThingGroupProperties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-thinggroupproperties
	ThingGroupProperties cfz.Expression[AWS_IoT_ThingGroup_ThingGroupProperties] `json:"ThingGroupProperties,omitempty"`
}

// New__AWS_IoT_ThingGroup initializes a new *AWS_IoT_ThingGroup.
func New__AWS_IoT_ThingGroup(logicalName string) *AWS_IoT_ThingGroup {
	return &AWS_IoT_ThingGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IoT_ThingGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IoT_ThingGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IoT_ThingGroup) GetType() string {
	return AWS_IoT_ThingGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IoT_ThingGroup) Set__LogicalName(v string) *AWS_IoT_ThingGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IoT_ThingGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IoT_ThingGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IoT_ThingGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IoT_ThingGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IoT_ThingGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IoT_ThingGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IoT_ThingGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IoT_ThingGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IoT_ThingGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_IoT_ThingGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IoT_ThingGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_IoT_ThingGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ParentGroupName updates property "ParentGroupName".
func (t *AWS_IoT_ThingGroup) Set__ParentGroupName(v cfz.Expression[string]) *AWS_IoT_ThingGroup {
	t.ParentGroupName = v
	return t
}

// SetV__ParentGroupName updates property "ParentGroupName".
func (t *AWS_IoT_ThingGroup) SetV__ParentGroupName(v string) *AWS_IoT_ThingGroup {
	t.ParentGroupName = cfz.V(v)
	return t
}

// Set__QueryString updates property "QueryString".
func (t *AWS_IoT_ThingGroup) Set__QueryString(v cfz.Expression[string]) *AWS_IoT_ThingGroup {
	t.QueryString = v
	return t
}

// SetV__QueryString updates property "QueryString".
func (t *AWS_IoT_ThingGroup) SetV__QueryString(v string) *AWS_IoT_ThingGroup {
	t.QueryString = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IoT_ThingGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IoT_ThingGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IoT_ThingGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IoT_ThingGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IoT_ThingGroup) SetSV__Tags(v ...cfz.Tag) *AWS_IoT_ThingGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__ThingGroupName updates property "ThingGroupName".
func (t *AWS_IoT_ThingGroup) Set__ThingGroupName(v cfz.Expression[string]) *AWS_IoT_ThingGroup {
	t.ThingGroupName = v
	return t
}

// SetV__ThingGroupName updates property "ThingGroupName".
func (t *AWS_IoT_ThingGroup) SetV__ThingGroupName(v string) *AWS_IoT_ThingGroup {
	t.ThingGroupName = cfz.V(v)
	return t
}

// Set__ThingGroupProperties updates property "ThingGroupProperties".
func (t *AWS_IoT_ThingGroup) Set__ThingGroupProperties(v cfz.Expression[AWS_IoT_ThingGroup_ThingGroupProperties]) *AWS_IoT_ThingGroup {
	t.ThingGroupProperties = v
	return t
}

// SetV__ThingGroupProperties updates property "ThingGroupProperties".
func (t *AWS_IoT_ThingGroup) SetV__ThingGroupProperties(v AWS_IoT_ThingGroup_ThingGroupProperties) *AWS_IoT_ThingGroup {
	t.ThingGroupProperties = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IoT_ThingGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_IoT_ThingGroup) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoT_ThingGroup__AttributesMap.Arn))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_IoT_ThingGroup) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoT_ThingGroup__AttributesMap.Id))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IoT_ThingGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_IoT_ThingGroup) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_IoT_ThingGroup) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IoT_ThingGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IoT_ThingGroup

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

func (t *AWS_IoT_ThingGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
