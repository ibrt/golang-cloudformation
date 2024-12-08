// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_memorydb

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MemoryDB_SubnetGroup)(nil)
	_ cfz.Resource                   = (*AWS_MemoryDB_SubnetGroup)(nil)
)

const (
	// AWS_MemoryDB_SubnetGroup__Type is the CloudFormation type for AWS::MemoryDB::SubnetGroup.
	AWS_MemoryDB_SubnetGroup__Type = "AWS::MemoryDB::SubnetGroup"
)

var (
	// AWS_MemoryDB_SubnetGroup__AttributesMap reports all the CloudFormation attributes for AWS::MemoryDB::SubnetGroup.
	AWS_MemoryDB_SubnetGroup__AttributesMap = struct {
		ARN string
	}{
		ARN: "ARN",
	}

	// AWS_MemoryDB_SubnetGroup__AttributesSlice reports all the CloudFormation attributes for AWS::MemoryDB::SubnetGroup.
	AWS_MemoryDB_SubnetGroup__AttributesSlice = []string{
		AWS_MemoryDB_SubnetGroup__AttributesMap.ARN,
	}
)

var (
	// AWS_MemoryDB_SubnetGroup__PropertiesMap reports all the CloudFormation properties for AWS::MemoryDB::SubnetGroup.
	AWS_MemoryDB_SubnetGroup__PropertiesMap = struct {
		Description     string
		SubnetGroupName string
		SubnetIds       string
		Tags            string
	}{
		Description:     "Description",
		SubnetGroupName: "SubnetGroupName",
		SubnetIds:       "SubnetIds",
		Tags:            "Tags",
	}

	// AWS_MemoryDB_SubnetGroup__PropertiesSlice reports all the CloudFormation properties for AWS::MemoryDB::SubnetGroup.
	AWS_MemoryDB_SubnetGroup__PropertiesSlice = []string{
		AWS_MemoryDB_SubnetGroup__PropertiesMap.Description,
		AWS_MemoryDB_SubnetGroup__PropertiesMap.SubnetGroupName,
		AWS_MemoryDB_SubnetGroup__PropertiesMap.SubnetIds,
		AWS_MemoryDB_SubnetGroup__PropertiesMap.Tags,
	}
)

// AWS_MemoryDB_SubnetGroup is a binding for AWS::MemoryDB::SubnetGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-subnetgroup.html
type AWS_MemoryDB_SubnetGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-subnetgroup.html#cfn-memorydb-subnetgroup-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// SubnetGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-subnetgroup.html#cfn-memorydb-subnetgroup-subnetgroupname
	SubnetGroupName cfz.Expression[string] `json:"SubnetGroupName,omitempty"`

	// SubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-subnetgroup.html#cfn-memorydb-subnetgroup-subnetids
	SubnetIds cfz.ExpressionSlice[string] `json:"SubnetIds,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-subnetgroup.html#cfn-memorydb-subnetgroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_MemoryDB_SubnetGroup initializes a new *AWS_MemoryDB_SubnetGroup.
func New__AWS_MemoryDB_SubnetGroup(logicalName string) *AWS_MemoryDB_SubnetGroup {
	return &AWS_MemoryDB_SubnetGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MemoryDB_SubnetGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MemoryDB_SubnetGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MemoryDB_SubnetGroup) GetType() string {
	return AWS_MemoryDB_SubnetGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MemoryDB_SubnetGroup) Set__LogicalName(v string) *AWS_MemoryDB_SubnetGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MemoryDB_SubnetGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MemoryDB_SubnetGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MemoryDB_SubnetGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MemoryDB_SubnetGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MemoryDB_SubnetGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MemoryDB_SubnetGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MemoryDB_SubnetGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MemoryDB_SubnetGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MemoryDB_SubnetGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_MemoryDB_SubnetGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MemoryDB_SubnetGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_MemoryDB_SubnetGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_MemoryDB_SubnetGroup) Set__Description(v cfz.Expression[string]) *AWS_MemoryDB_SubnetGroup {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_MemoryDB_SubnetGroup) SetV__Description(v string) *AWS_MemoryDB_SubnetGroup {
	t.Description = cfz.V(v)
	return t
}

// Set__SubnetGroupName updates property "SubnetGroupName".
func (t *AWS_MemoryDB_SubnetGroup) Set__SubnetGroupName(v cfz.Expression[string]) *AWS_MemoryDB_SubnetGroup {
	t.SubnetGroupName = v
	return t
}

// SetV__SubnetGroupName updates property "SubnetGroupName".
func (t *AWS_MemoryDB_SubnetGroup) SetV__SubnetGroupName(v string) *AWS_MemoryDB_SubnetGroup {
	t.SubnetGroupName = cfz.V(v)
	return t
}

// Set__SubnetIds updates property "SubnetIds".
func (t *AWS_MemoryDB_SubnetGroup) Set__SubnetIds(v cfz.ExpressionSlice[string]) *AWS_MemoryDB_SubnetGroup {
	t.SubnetIds = v
	return t
}

// SetS__SubnetIds updates property "SubnetIds".
func (t *AWS_MemoryDB_SubnetGroup) SetS__SubnetIds(v ...cfz.Expression[string]) *AWS_MemoryDB_SubnetGroup {
	t.SubnetIds = cfz.S(v...)
	return t
}

// SetSV__SubnetIds updates property "SubnetIds".
func (t *AWS_MemoryDB_SubnetGroup) SetSV__SubnetIds(v ...string) *AWS_MemoryDB_SubnetGroup {
	t.SubnetIds = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MemoryDB_SubnetGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_MemoryDB_SubnetGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_MemoryDB_SubnetGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_MemoryDB_SubnetGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_MemoryDB_SubnetGroup) SetSV__Tags(v ...cfz.Tag) *AWS_MemoryDB_SubnetGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MemoryDB_SubnetGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__ARN returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ARN
func (t *AWS_MemoryDB_SubnetGroup) GetAtt__ARN() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MemoryDB_SubnetGroup__AttributesMap.ARN))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MemoryDB_SubnetGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ARN returns a conventionally configured output for an attribute of this resource.
// Attribute: ARN
func (t *AWS_MemoryDB_SubnetGroup) GetConventionalOutputAtt__ARN(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttARN", t.GetAtt__ARN())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MemoryDB_SubnetGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MemoryDB_SubnetGroup

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

func (t *AWS_MemoryDB_SubnetGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
