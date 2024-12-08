// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_TransitGatewayRouteTable)(nil)
	_ cfz.Resource                   = (*AWS_EC2_TransitGatewayRouteTable)(nil)
)

const (
	// AWS_EC2_TransitGatewayRouteTable__Type is the CloudFormation type for AWS::EC2::TransitGatewayRouteTable.
	AWS_EC2_TransitGatewayRouteTable__Type = "AWS::EC2::TransitGatewayRouteTable"
)

var (
	// AWS_EC2_TransitGatewayRouteTable__AttributesMap reports all the CloudFormation attributes for AWS::EC2::TransitGatewayRouteTable.
	AWS_EC2_TransitGatewayRouteTable__AttributesMap = struct {
		TransitGatewayRouteTableId string
	}{
		TransitGatewayRouteTableId: "TransitGatewayRouteTableId",
	}

	// AWS_EC2_TransitGatewayRouteTable__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::TransitGatewayRouteTable.
	AWS_EC2_TransitGatewayRouteTable__AttributesSlice = []string{
		AWS_EC2_TransitGatewayRouteTable__AttributesMap.TransitGatewayRouteTableId,
	}
)

var (
	// AWS_EC2_TransitGatewayRouteTable__PropertiesMap reports all the CloudFormation properties for AWS::EC2::TransitGatewayRouteTable.
	AWS_EC2_TransitGatewayRouteTable__PropertiesMap = struct {
		Tags             string
		TransitGatewayId string
	}{
		Tags:             "Tags",
		TransitGatewayId: "TransitGatewayId",
	}

	// AWS_EC2_TransitGatewayRouteTable__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::TransitGatewayRouteTable.
	AWS_EC2_TransitGatewayRouteTable__PropertiesSlice = []string{
		AWS_EC2_TransitGatewayRouteTable__PropertiesMap.Tags,
		AWS_EC2_TransitGatewayRouteTable__PropertiesMap.TransitGatewayId,
	}
)

// AWS_EC2_TransitGatewayRouteTable is a binding for AWS::EC2::TransitGatewayRouteTable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html
type AWS_EC2_TransitGatewayRouteTable struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TransitGatewayId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetable.html#cfn-ec2-transitgatewayroutetable-transitgatewayid
	TransitGatewayId cfz.Expression[string] `json:"TransitGatewayId,omitempty"`
}

// New__AWS_EC2_TransitGatewayRouteTable initializes a new *AWS_EC2_TransitGatewayRouteTable.
func New__AWS_EC2_TransitGatewayRouteTable(logicalName string) *AWS_EC2_TransitGatewayRouteTable {
	return &AWS_EC2_TransitGatewayRouteTable{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_TransitGatewayRouteTable) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayRouteTable) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_TransitGatewayRouteTable) GetType() string {
	return AWS_EC2_TransitGatewayRouteTable__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__LogicalName(v string) *AWS_EC2_TransitGatewayRouteTable {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayRouteTable {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayRouteTable) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayRouteTable {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_TransitGatewayRouteTable {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_TransitGatewayRouteTable {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_TransitGatewayRouteTable {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayRouteTable) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_TransitGatewayRouteTable {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EC2_TransitGatewayRouteTable {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayRouteTable) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EC2_TransitGatewayRouteTable {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayRouteTable) SetSV__Tags(v ...cfz.Tag) *AWS_EC2_TransitGatewayRouteTable {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TransitGatewayId updates property "TransitGatewayId".
func (t *AWS_EC2_TransitGatewayRouteTable) Set__TransitGatewayId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayRouteTable {
	t.TransitGatewayId = v
	return t
}

// SetV__TransitGatewayId updates property "TransitGatewayId".
func (t *AWS_EC2_TransitGatewayRouteTable) SetV__TransitGatewayId(v string) *AWS_EC2_TransitGatewayRouteTable {
	t.TransitGatewayId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_TransitGatewayRouteTable) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__TransitGatewayRouteTableId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TransitGatewayRouteTableId
func (t *AWS_EC2_TransitGatewayRouteTable) GetAtt__TransitGatewayRouteTableId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayRouteTable__AttributesMap.TransitGatewayRouteTableId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_TransitGatewayRouteTable) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TransitGatewayRouteTableId returns a conventionally configured output for an attribute of this resource.
// Attribute: TransitGatewayRouteTableId
func (t *AWS_EC2_TransitGatewayRouteTable) GetConventionalOutputAtt__TransitGatewayRouteTableId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTransitGatewayRouteTableId", t.GetAtt__TransitGatewayRouteTableId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayRouteTable) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_TransitGatewayRouteTable

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

func (t *AWS_EC2_TransitGatewayRouteTable) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
