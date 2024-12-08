// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_TransitGatewayRoute)(nil)
	_ cfz.Resource                   = (*AWS_EC2_TransitGatewayRoute)(nil)
)

const (
	// AWS_EC2_TransitGatewayRoute__Type is the CloudFormation type for AWS::EC2::TransitGatewayRoute.
	AWS_EC2_TransitGatewayRoute__Type = "AWS::EC2::TransitGatewayRoute"
)

var (
	// AWS_EC2_TransitGatewayRoute__PropertiesMap reports all the CloudFormation properties for AWS::EC2::TransitGatewayRoute.
	AWS_EC2_TransitGatewayRoute__PropertiesMap = struct {
		Blackhole                  string
		DestinationCidrBlock       string
		TransitGatewayAttachmentId string
		TransitGatewayRouteTableId string
	}{
		Blackhole:                  "Blackhole",
		DestinationCidrBlock:       "DestinationCidrBlock",
		TransitGatewayAttachmentId: "TransitGatewayAttachmentId",
		TransitGatewayRouteTableId: "TransitGatewayRouteTableId",
	}

	// AWS_EC2_TransitGatewayRoute__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::TransitGatewayRoute.
	AWS_EC2_TransitGatewayRoute__PropertiesSlice = []string{
		AWS_EC2_TransitGatewayRoute__PropertiesMap.Blackhole,
		AWS_EC2_TransitGatewayRoute__PropertiesMap.DestinationCidrBlock,
		AWS_EC2_TransitGatewayRoute__PropertiesMap.TransitGatewayAttachmentId,
		AWS_EC2_TransitGatewayRoute__PropertiesMap.TransitGatewayRouteTableId,
	}
)

// AWS_EC2_TransitGatewayRoute is a binding for AWS::EC2::TransitGatewayRoute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html
type AWS_EC2_TransitGatewayRoute struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Blackhole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-blackhole
	Blackhole cfz.Expression[bool] `json:"Blackhole,omitempty"`

	// DestinationCidrBlock is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-destinationcidrblock
	DestinationCidrBlock cfz.Expression[string] `json:"DestinationCidrBlock,omitempty"`

	// TransitGatewayAttachmentId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-transitgatewayattachmentid
	TransitGatewayAttachmentId cfz.Expression[string] `json:"TransitGatewayAttachmentId,omitempty"`

	// TransitGatewayRouteTableId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroute.html#cfn-ec2-transitgatewayroute-transitgatewayroutetableid
	TransitGatewayRouteTableId cfz.Expression[string] `json:"TransitGatewayRouteTableId,omitempty"`
}

// New__AWS_EC2_TransitGatewayRoute initializes a new *AWS_EC2_TransitGatewayRoute.
func New__AWS_EC2_TransitGatewayRoute(logicalName string) *AWS_EC2_TransitGatewayRoute {
	return &AWS_EC2_TransitGatewayRoute{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_TransitGatewayRoute) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayRoute) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_TransitGatewayRoute) GetType() string {
	return AWS_EC2_TransitGatewayRoute__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_TransitGatewayRoute) Set__LogicalName(v string) *AWS_EC2_TransitGatewayRoute {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayRoute) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayRoute {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayRoute) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayRoute {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_TransitGatewayRoute) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_TransitGatewayRoute {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_TransitGatewayRoute) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_TransitGatewayRoute {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayRoute) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_TransitGatewayRoute {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayRoute) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_TransitGatewayRoute {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Blackhole updates property "Blackhole".
func (t *AWS_EC2_TransitGatewayRoute) Set__Blackhole(v cfz.Expression[bool]) *AWS_EC2_TransitGatewayRoute {
	t.Blackhole = v
	return t
}

// SetV__Blackhole updates property "Blackhole".
func (t *AWS_EC2_TransitGatewayRoute) SetV__Blackhole(v bool) *AWS_EC2_TransitGatewayRoute {
	t.Blackhole = cfz.V(v)
	return t
}

// Set__DestinationCidrBlock updates property "DestinationCidrBlock".
func (t *AWS_EC2_TransitGatewayRoute) Set__DestinationCidrBlock(v cfz.Expression[string]) *AWS_EC2_TransitGatewayRoute {
	t.DestinationCidrBlock = v
	return t
}

// SetV__DestinationCidrBlock updates property "DestinationCidrBlock".
func (t *AWS_EC2_TransitGatewayRoute) SetV__DestinationCidrBlock(v string) *AWS_EC2_TransitGatewayRoute {
	t.DestinationCidrBlock = cfz.V(v)
	return t
}

// Set__TransitGatewayAttachmentId updates property "TransitGatewayAttachmentId".
func (t *AWS_EC2_TransitGatewayRoute) Set__TransitGatewayAttachmentId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayRoute {
	t.TransitGatewayAttachmentId = v
	return t
}

// SetV__TransitGatewayAttachmentId updates property "TransitGatewayAttachmentId".
func (t *AWS_EC2_TransitGatewayRoute) SetV__TransitGatewayAttachmentId(v string) *AWS_EC2_TransitGatewayRoute {
	t.TransitGatewayAttachmentId = cfz.V(v)
	return t
}

// Set__TransitGatewayRouteTableId updates property "TransitGatewayRouteTableId".
func (t *AWS_EC2_TransitGatewayRoute) Set__TransitGatewayRouteTableId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayRoute {
	t.TransitGatewayRouteTableId = v
	return t
}

// SetV__TransitGatewayRouteTableId updates property "TransitGatewayRouteTableId".
func (t *AWS_EC2_TransitGatewayRoute) SetV__TransitGatewayRouteTableId(v string) *AWS_EC2_TransitGatewayRoute {
	t.TransitGatewayRouteTableId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_TransitGatewayRoute) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_TransitGatewayRoute) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayRoute) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_TransitGatewayRoute

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

func (t *AWS_EC2_TransitGatewayRoute) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
