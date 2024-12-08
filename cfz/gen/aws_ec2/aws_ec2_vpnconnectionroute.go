// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_VPNConnectionRoute)(nil)
	_ cfz.Resource                   = (*AWS_EC2_VPNConnectionRoute)(nil)
)

const (
	// AWS_EC2_VPNConnectionRoute__Type is the CloudFormation type for AWS::EC2::VPNConnectionRoute.
	AWS_EC2_VPNConnectionRoute__Type = "AWS::EC2::VPNConnectionRoute"
)

var (
	// AWS_EC2_VPNConnectionRoute__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VPNConnectionRoute.
	AWS_EC2_VPNConnectionRoute__PropertiesMap = struct {
		DestinationCidrBlock string
		VpnConnectionId      string
	}{
		DestinationCidrBlock: "DestinationCidrBlock",
		VpnConnectionId:      "VpnConnectionId",
	}

	// AWS_EC2_VPNConnectionRoute__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VPNConnectionRoute.
	AWS_EC2_VPNConnectionRoute__PropertiesSlice = []string{
		AWS_EC2_VPNConnectionRoute__PropertiesMap.DestinationCidrBlock,
		AWS_EC2_VPNConnectionRoute__PropertiesMap.VpnConnectionId,
	}
)

// AWS_EC2_VPNConnectionRoute is a binding for AWS::EC2::VPNConnectionRoute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html
type AWS_EC2_VPNConnectionRoute struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DestinationCidrBlock is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html#cfn-ec2-vpnconnectionroute-destinationcidrblock
	DestinationCidrBlock cfz.Expression[string] `json:"DestinationCidrBlock,omitempty"`

	// VpnConnectionId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpnconnectionroute.html#cfn-ec2-vpnconnectionroute-vpnconnectionid
	VpnConnectionId cfz.Expression[string] `json:"VpnConnectionId,omitempty"`
}

// New__AWS_EC2_VPNConnectionRoute initializes a new *AWS_EC2_VPNConnectionRoute.
func New__AWS_EC2_VPNConnectionRoute(logicalName string) *AWS_EC2_VPNConnectionRoute {
	return &AWS_EC2_VPNConnectionRoute{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_VPNConnectionRoute) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_VPNConnectionRoute) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_VPNConnectionRoute) GetType() string {
	return AWS_EC2_VPNConnectionRoute__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_VPNConnectionRoute) Set__LogicalName(v string) *AWS_EC2_VPNConnectionRoute {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_VPNConnectionRoute) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_VPNConnectionRoute {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_VPNConnectionRoute) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_VPNConnectionRoute {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_VPNConnectionRoute) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_VPNConnectionRoute {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_VPNConnectionRoute) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_VPNConnectionRoute {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_VPNConnectionRoute) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_VPNConnectionRoute {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_VPNConnectionRoute) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_VPNConnectionRoute {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DestinationCidrBlock updates property "DestinationCidrBlock".
func (t *AWS_EC2_VPNConnectionRoute) Set__DestinationCidrBlock(v cfz.Expression[string]) *AWS_EC2_VPNConnectionRoute {
	t.DestinationCidrBlock = v
	return t
}

// SetV__DestinationCidrBlock updates property "DestinationCidrBlock".
func (t *AWS_EC2_VPNConnectionRoute) SetV__DestinationCidrBlock(v string) *AWS_EC2_VPNConnectionRoute {
	t.DestinationCidrBlock = cfz.V(v)
	return t
}

// Set__VpnConnectionId updates property "VpnConnectionId".
func (t *AWS_EC2_VPNConnectionRoute) Set__VpnConnectionId(v cfz.Expression[string]) *AWS_EC2_VPNConnectionRoute {
	t.VpnConnectionId = v
	return t
}

// SetV__VpnConnectionId updates property "VpnConnectionId".
func (t *AWS_EC2_VPNConnectionRoute) SetV__VpnConnectionId(v string) *AWS_EC2_VPNConnectionRoute {
	t.VpnConnectionId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_VPNConnectionRoute) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_VPNConnectionRoute) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_VPNConnectionRoute) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_VPNConnectionRoute

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

func (t *AWS_EC2_VPNConnectionRoute) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
