// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_ClientVpnTargetNetworkAssociation)(nil)
	_ cfz.Resource                   = (*AWS_EC2_ClientVpnTargetNetworkAssociation)(nil)
)

const (
	// AWS_EC2_ClientVpnTargetNetworkAssociation__Type is the CloudFormation type for AWS::EC2::ClientVpnTargetNetworkAssociation.
	AWS_EC2_ClientVpnTargetNetworkAssociation__Type = "AWS::EC2::ClientVpnTargetNetworkAssociation"
)

var (
	// AWS_EC2_ClientVpnTargetNetworkAssociation__PropertiesMap reports all the CloudFormation properties for AWS::EC2::ClientVpnTargetNetworkAssociation.
	AWS_EC2_ClientVpnTargetNetworkAssociation__PropertiesMap = struct {
		ClientVpnEndpointId string
		SubnetId            string
	}{
		ClientVpnEndpointId: "ClientVpnEndpointId",
		SubnetId:            "SubnetId",
	}

	// AWS_EC2_ClientVpnTargetNetworkAssociation__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::ClientVpnTargetNetworkAssociation.
	AWS_EC2_ClientVpnTargetNetworkAssociation__PropertiesSlice = []string{
		AWS_EC2_ClientVpnTargetNetworkAssociation__PropertiesMap.ClientVpnEndpointId,
		AWS_EC2_ClientVpnTargetNetworkAssociation__PropertiesMap.SubnetId,
	}
)

// AWS_EC2_ClientVpnTargetNetworkAssociation is a binding for AWS::EC2::ClientVpnTargetNetworkAssociation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html
type AWS_EC2_ClientVpnTargetNetworkAssociation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html#cfn-ec2-clientvpntargetnetworkassociation-clientvpnendpointid
	ClientVpnEndpointId cfz.Expression[string] `json:"ClientVpnEndpointId,omitempty"`

	// SubnetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html#cfn-ec2-clientvpntargetnetworkassociation-subnetid
	SubnetId cfz.Expression[string] `json:"SubnetId,omitempty"`
}

// New__AWS_EC2_ClientVpnTargetNetworkAssociation initializes a new *AWS_EC2_ClientVpnTargetNetworkAssociation.
func New__AWS_EC2_ClientVpnTargetNetworkAssociation(logicalName string) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	return &AWS_EC2_ClientVpnTargetNetworkAssociation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_ClientVpnTargetNetworkAssociation) GetType() string {
	return AWS_EC2_ClientVpnTargetNetworkAssociation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__LogicalName(v string) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ClientVpnEndpointId updates property "ClientVpnEndpointId".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__ClientVpnEndpointId(v cfz.Expression[string]) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.ClientVpnEndpointId = v
	return t
}

// SetV__ClientVpnEndpointId updates property "ClientVpnEndpointId".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) SetV__ClientVpnEndpointId(v string) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.ClientVpnEndpointId = cfz.V(v)
	return t
}

// Set__SubnetId updates property "SubnetId".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Set__SubnetId(v cfz.Expression[string]) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.SubnetId = v
	return t
}

// SetV__SubnetId updates property "SubnetId".
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) SetV__SubnetId(v string) *AWS_EC2_ClientVpnTargetNetworkAssociation {
	t.SubnetId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_ClientVpnTargetNetworkAssociation

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

func (t *AWS_EC2_ClientVpnTargetNetworkAssociation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
