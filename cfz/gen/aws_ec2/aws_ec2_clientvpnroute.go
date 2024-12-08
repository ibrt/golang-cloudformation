// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_ClientVpnRoute)(nil)
	_ cfz.Resource                   = (*AWS_EC2_ClientVpnRoute)(nil)
)

const (
	// AWS_EC2_ClientVpnRoute__Type is the CloudFormation type for AWS::EC2::ClientVpnRoute.
	AWS_EC2_ClientVpnRoute__Type = "AWS::EC2::ClientVpnRoute"
)

var (
	// AWS_EC2_ClientVpnRoute__PropertiesMap reports all the CloudFormation properties for AWS::EC2::ClientVpnRoute.
	AWS_EC2_ClientVpnRoute__PropertiesMap = struct {
		ClientVpnEndpointId  string
		Description          string
		DestinationCidrBlock string
		TargetVpcSubnetId    string
	}{
		ClientVpnEndpointId:  "ClientVpnEndpointId",
		Description:          "Description",
		DestinationCidrBlock: "DestinationCidrBlock",
		TargetVpcSubnetId:    "TargetVpcSubnetId",
	}

	// AWS_EC2_ClientVpnRoute__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::ClientVpnRoute.
	AWS_EC2_ClientVpnRoute__PropertiesSlice = []string{
		AWS_EC2_ClientVpnRoute__PropertiesMap.ClientVpnEndpointId,
		AWS_EC2_ClientVpnRoute__PropertiesMap.Description,
		AWS_EC2_ClientVpnRoute__PropertiesMap.DestinationCidrBlock,
		AWS_EC2_ClientVpnRoute__PropertiesMap.TargetVpcSubnetId,
	}
)

// AWS_EC2_ClientVpnRoute is a binding for AWS::EC2::ClientVpnRoute.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html
type AWS_EC2_ClientVpnRoute struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ClientVpnEndpointId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-clientvpnendpointid
	ClientVpnEndpointId cfz.Expression[string] `json:"ClientVpnEndpointId,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DestinationCidrBlock is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-destinationcidrblock
	DestinationCidrBlock cfz.Expression[string] `json:"DestinationCidrBlock,omitempty"`

	// TargetVpcSubnetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnroute.html#cfn-ec2-clientvpnroute-targetvpcsubnetid
	TargetVpcSubnetId cfz.Expression[string] `json:"TargetVpcSubnetId,omitempty"`
}

// New__AWS_EC2_ClientVpnRoute initializes a new *AWS_EC2_ClientVpnRoute.
func New__AWS_EC2_ClientVpnRoute(logicalName string) *AWS_EC2_ClientVpnRoute {
	return &AWS_EC2_ClientVpnRoute{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_ClientVpnRoute) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_ClientVpnRoute) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_ClientVpnRoute) GetType() string {
	return AWS_EC2_ClientVpnRoute__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_ClientVpnRoute) Set__LogicalName(v string) *AWS_EC2_ClientVpnRoute {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_ClientVpnRoute) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_ClientVpnRoute {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_ClientVpnRoute) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_ClientVpnRoute {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_ClientVpnRoute) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_ClientVpnRoute {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_ClientVpnRoute) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_ClientVpnRoute {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_ClientVpnRoute) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_ClientVpnRoute {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_ClientVpnRoute) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_ClientVpnRoute {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ClientVpnEndpointId updates property "ClientVpnEndpointId".
func (t *AWS_EC2_ClientVpnRoute) Set__ClientVpnEndpointId(v cfz.Expression[string]) *AWS_EC2_ClientVpnRoute {
	t.ClientVpnEndpointId = v
	return t
}

// SetV__ClientVpnEndpointId updates property "ClientVpnEndpointId".
func (t *AWS_EC2_ClientVpnRoute) SetV__ClientVpnEndpointId(v string) *AWS_EC2_ClientVpnRoute {
	t.ClientVpnEndpointId = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_EC2_ClientVpnRoute) Set__Description(v cfz.Expression[string]) *AWS_EC2_ClientVpnRoute {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_EC2_ClientVpnRoute) SetV__Description(v string) *AWS_EC2_ClientVpnRoute {
	t.Description = cfz.V(v)
	return t
}

// Set__DestinationCidrBlock updates property "DestinationCidrBlock".
func (t *AWS_EC2_ClientVpnRoute) Set__DestinationCidrBlock(v cfz.Expression[string]) *AWS_EC2_ClientVpnRoute {
	t.DestinationCidrBlock = v
	return t
}

// SetV__DestinationCidrBlock updates property "DestinationCidrBlock".
func (t *AWS_EC2_ClientVpnRoute) SetV__DestinationCidrBlock(v string) *AWS_EC2_ClientVpnRoute {
	t.DestinationCidrBlock = cfz.V(v)
	return t
}

// Set__TargetVpcSubnetId updates property "TargetVpcSubnetId".
func (t *AWS_EC2_ClientVpnRoute) Set__TargetVpcSubnetId(v cfz.Expression[string]) *AWS_EC2_ClientVpnRoute {
	t.TargetVpcSubnetId = v
	return t
}

// SetV__TargetVpcSubnetId updates property "TargetVpcSubnetId".
func (t *AWS_EC2_ClientVpnRoute) SetV__TargetVpcSubnetId(v string) *AWS_EC2_ClientVpnRoute {
	t.TargetVpcSubnetId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_ClientVpnRoute) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_ClientVpnRoute) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_ClientVpnRoute) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_ClientVpnRoute

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

func (t *AWS_EC2_ClientVpnRoute) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
