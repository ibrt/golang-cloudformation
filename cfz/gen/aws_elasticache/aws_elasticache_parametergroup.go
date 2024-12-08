// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_elasticache

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ElastiCache_ParameterGroup)(nil)
	_ cfz.Resource                   = (*AWS_ElastiCache_ParameterGroup)(nil)
)

const (
	// AWS_ElastiCache_ParameterGroup__Type is the CloudFormation type for AWS::ElastiCache::ParameterGroup.
	AWS_ElastiCache_ParameterGroup__Type = "AWS::ElastiCache::ParameterGroup"
)

var (
	// AWS_ElastiCache_ParameterGroup__AttributesMap reports all the CloudFormation attributes for AWS::ElastiCache::ParameterGroup.
	AWS_ElastiCache_ParameterGroup__AttributesMap = struct {
		CacheParameterGroupName string
	}{
		CacheParameterGroupName: "CacheParameterGroupName",
	}

	// AWS_ElastiCache_ParameterGroup__AttributesSlice reports all the CloudFormation attributes for AWS::ElastiCache::ParameterGroup.
	AWS_ElastiCache_ParameterGroup__AttributesSlice = []string{
		AWS_ElastiCache_ParameterGroup__AttributesMap.CacheParameterGroupName,
	}
)

var (
	// AWS_ElastiCache_ParameterGroup__PropertiesMap reports all the CloudFormation properties for AWS::ElastiCache::ParameterGroup.
	AWS_ElastiCache_ParameterGroup__PropertiesMap = struct {
		CacheParameterGroupFamily string
		Description               string
		Properties                string
		Tags                      string
	}{
		CacheParameterGroupFamily: "CacheParameterGroupFamily",
		Description:               "Description",
		Properties:                "Properties",
		Tags:                      "Tags",
	}

	// AWS_ElastiCache_ParameterGroup__PropertiesSlice reports all the CloudFormation properties for AWS::ElastiCache::ParameterGroup.
	AWS_ElastiCache_ParameterGroup__PropertiesSlice = []string{
		AWS_ElastiCache_ParameterGroup__PropertiesMap.CacheParameterGroupFamily,
		AWS_ElastiCache_ParameterGroup__PropertiesMap.Description,
		AWS_ElastiCache_ParameterGroup__PropertiesMap.Properties,
		AWS_ElastiCache_ParameterGroup__PropertiesMap.Tags,
	}
)

// AWS_ElastiCache_ParameterGroup is a binding for AWS::ElastiCache::ParameterGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html
type AWS_ElastiCache_ParameterGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CacheParameterGroupFamily is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-cacheparametergroupfamily
	CacheParameterGroupFamily cfz.Expression[string] `json:"CacheParameterGroupFamily,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// Properties is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-properties
	Properties cfz.ExpressionMap[string] `json:"Properties,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_ElastiCache_ParameterGroup initializes a new *AWS_ElastiCache_ParameterGroup.
func New__AWS_ElastiCache_ParameterGroup(logicalName string) *AWS_ElastiCache_ParameterGroup {
	return &AWS_ElastiCache_ParameterGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ElastiCache_ParameterGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ElastiCache_ParameterGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ElastiCache_ParameterGroup) GetType() string {
	return AWS_ElastiCache_ParameterGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ElastiCache_ParameterGroup) Set__LogicalName(v string) *AWS_ElastiCache_ParameterGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ElastiCache_ParameterGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ElastiCache_ParameterGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ElastiCache_ParameterGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ElastiCache_ParameterGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ElastiCache_ParameterGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ElastiCache_ParameterGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ElastiCache_ParameterGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ElastiCache_ParameterGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ElastiCache_ParameterGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_ElastiCache_ParameterGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ElastiCache_ParameterGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_ElastiCache_ParameterGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CacheParameterGroupFamily updates property "CacheParameterGroupFamily".
func (t *AWS_ElastiCache_ParameterGroup) Set__CacheParameterGroupFamily(v cfz.Expression[string]) *AWS_ElastiCache_ParameterGroup {
	t.CacheParameterGroupFamily = v
	return t
}

// SetV__CacheParameterGroupFamily updates property "CacheParameterGroupFamily".
func (t *AWS_ElastiCache_ParameterGroup) SetV__CacheParameterGroupFamily(v string) *AWS_ElastiCache_ParameterGroup {
	t.CacheParameterGroupFamily = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_ElastiCache_ParameterGroup) Set__Description(v cfz.Expression[string]) *AWS_ElastiCache_ParameterGroup {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_ElastiCache_ParameterGroup) SetV__Description(v string) *AWS_ElastiCache_ParameterGroup {
	t.Description = cfz.V(v)
	return t
}

// Set__Properties updates property "Properties".
func (t *AWS_ElastiCache_ParameterGroup) Set__Properties(v cfz.ExpressionMap[string]) *AWS_ElastiCache_ParameterGroup {
	t.Properties = v
	return t
}

// SetM__Properties updates property "Properties".
func (t *AWS_ElastiCache_ParameterGroup) SetM__Properties(v ...map[string]cfz.Expression[string]) *AWS_ElastiCache_ParameterGroup {
	t.Properties = cfz.M(v...)
	return t
}

// SetMV__Properties updates property "Properties".
func (t *AWS_ElastiCache_ParameterGroup) SetMV__Properties(v ...map[string]string) *AWS_ElastiCache_ParameterGroup {
	t.Properties = cfz.MV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_ElastiCache_ParameterGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_ElastiCache_ParameterGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_ElastiCache_ParameterGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_ElastiCache_ParameterGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_ElastiCache_ParameterGroup) SetSV__Tags(v ...cfz.Tag) *AWS_ElastiCache_ParameterGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ElastiCache_ParameterGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CacheParameterGroupName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CacheParameterGroupName
func (t *AWS_ElastiCache_ParameterGroup) GetAtt__CacheParameterGroupName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ElastiCache_ParameterGroup__AttributesMap.CacheParameterGroupName))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ElastiCache_ParameterGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CacheParameterGroupName returns a conventionally configured output for an attribute of this resource.
// Attribute: CacheParameterGroupName
func (t *AWS_ElastiCache_ParameterGroup) GetConventionalOutputAtt__CacheParameterGroupName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCacheParameterGroupName", t.GetAtt__CacheParameterGroupName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ElastiCache_ParameterGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ElastiCache_ParameterGroup

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

func (t *AWS_ElastiCache_ParameterGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
