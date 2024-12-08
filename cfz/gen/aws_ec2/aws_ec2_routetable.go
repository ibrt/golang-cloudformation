// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_RouteTable)(nil)
	_ cfz.Resource                   = (*AWS_EC2_RouteTable)(nil)
)

const (
	// AWS_EC2_RouteTable__Type is the CloudFormation type for AWS::EC2::RouteTable.
	AWS_EC2_RouteTable__Type = "AWS::EC2::RouteTable"
)

var (
	// AWS_EC2_RouteTable__AttributesMap reports all the CloudFormation attributes for AWS::EC2::RouteTable.
	AWS_EC2_RouteTable__AttributesMap = struct {
		RouteTableId string
	}{
		RouteTableId: "RouteTableId",
	}

	// AWS_EC2_RouteTable__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::RouteTable.
	AWS_EC2_RouteTable__AttributesSlice = []string{
		AWS_EC2_RouteTable__AttributesMap.RouteTableId,
	}
)

var (
	// AWS_EC2_RouteTable__PropertiesMap reports all the CloudFormation properties for AWS::EC2::RouteTable.
	AWS_EC2_RouteTable__PropertiesMap = struct {
		Tags  string
		VpcId string
	}{
		Tags:  "Tags",
		VpcId: "VpcId",
	}

	// AWS_EC2_RouteTable__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::RouteTable.
	AWS_EC2_RouteTable__PropertiesSlice = []string{
		AWS_EC2_RouteTable__PropertiesMap.Tags,
		AWS_EC2_RouteTable__PropertiesMap.VpcId,
	}
)

// AWS_EC2_RouteTable is a binding for AWS::EC2::RouteTable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html
type AWS_EC2_RouteTable struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html#cfn-ec2-routetable-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// VpcId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html#cfn-ec2-routetable-vpcid
	VpcId cfz.Expression[string] `json:"VpcId,omitempty"`
}

// New__AWS_EC2_RouteTable initializes a new *AWS_EC2_RouteTable.
func New__AWS_EC2_RouteTable(logicalName string) *AWS_EC2_RouteTable {
	return &AWS_EC2_RouteTable{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_RouteTable) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_RouteTable) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_RouteTable) GetType() string {
	return AWS_EC2_RouteTable__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_RouteTable) Set__LogicalName(v string) *AWS_EC2_RouteTable {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_RouteTable) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_RouteTable {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_RouteTable) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_RouteTable {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_RouteTable) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_RouteTable {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_RouteTable) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_RouteTable {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_RouteTable) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_RouteTable {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_RouteTable) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_RouteTable {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EC2_RouteTable) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EC2_RouteTable {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EC2_RouteTable) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EC2_RouteTable {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EC2_RouteTable) SetSV__Tags(v ...cfz.Tag) *AWS_EC2_RouteTable {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__VpcId updates property "VpcId".
func (t *AWS_EC2_RouteTable) Set__VpcId(v cfz.Expression[string]) *AWS_EC2_RouteTable {
	t.VpcId = v
	return t
}

// SetV__VpcId updates property "VpcId".
func (t *AWS_EC2_RouteTable) SetV__VpcId(v string) *AWS_EC2_RouteTable {
	t.VpcId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_RouteTable) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__RouteTableId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RouteTableId
func (t *AWS_EC2_RouteTable) GetAtt__RouteTableId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_RouteTable__AttributesMap.RouteTableId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_RouteTable) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RouteTableId returns a conventionally configured output for an attribute of this resource.
// Attribute: RouteTableId
func (t *AWS_EC2_RouteTable) GetConventionalOutputAtt__RouteTableId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRouteTableId", t.GetAtt__RouteTableId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_RouteTable) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_RouteTable

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

func (t *AWS_EC2_RouteTable) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
