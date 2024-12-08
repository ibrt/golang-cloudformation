// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediaconnect

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MediaConnect_Gateway)(nil)
	_ cfz.Resource                   = (*AWS_MediaConnect_Gateway)(nil)
)

const (
	// AWS_MediaConnect_Gateway__Type is the CloudFormation type for AWS::MediaConnect::Gateway.
	AWS_MediaConnect_Gateway__Type = "AWS::MediaConnect::Gateway"
)

var (
	// AWS_MediaConnect_Gateway__AttributesMap reports all the CloudFormation attributes for AWS::MediaConnect::Gateway.
	AWS_MediaConnect_Gateway__AttributesMap = struct {
		GatewayArn   string
		GatewayState string
	}{
		GatewayArn:   "GatewayArn",
		GatewayState: "GatewayState",
	}

	// AWS_MediaConnect_Gateway__AttributesSlice reports all the CloudFormation attributes for AWS::MediaConnect::Gateway.
	AWS_MediaConnect_Gateway__AttributesSlice = []string{
		AWS_MediaConnect_Gateway__AttributesMap.GatewayArn,
		AWS_MediaConnect_Gateway__AttributesMap.GatewayState,
	}
)

var (
	// AWS_MediaConnect_Gateway__PropertiesMap reports all the CloudFormation properties for AWS::MediaConnect::Gateway.
	AWS_MediaConnect_Gateway__PropertiesMap = struct {
		EgressCidrBlocks string
		Name             string
		Networks         string
	}{
		EgressCidrBlocks: "EgressCidrBlocks",
		Name:             "Name",
		Networks:         "Networks",
	}

	// AWS_MediaConnect_Gateway__PropertiesSlice reports all the CloudFormation properties for AWS::MediaConnect::Gateway.
	AWS_MediaConnect_Gateway__PropertiesSlice = []string{
		AWS_MediaConnect_Gateway__PropertiesMap.EgressCidrBlocks,
		AWS_MediaConnect_Gateway__PropertiesMap.Name,
		AWS_MediaConnect_Gateway__PropertiesMap.Networks,
	}
)

// AWS_MediaConnect_Gateway is a binding for AWS::MediaConnect::Gateway.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html
type AWS_MediaConnect_Gateway struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// EgressCidrBlocks is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html#cfn-mediaconnect-gateway-egresscidrblocks
	EgressCidrBlocks cfz.ExpressionSlice[string] `json:"EgressCidrBlocks,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html#cfn-mediaconnect-gateway-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Networks is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-gateway.html#cfn-mediaconnect-gateway-networks
	Networks cfz.ExpressionSlice[AWS_MediaConnect_Gateway_GatewayNetwork] `json:"Networks,omitempty"`
}

// New__AWS_MediaConnect_Gateway initializes a new *AWS_MediaConnect_Gateway.
func New__AWS_MediaConnect_Gateway(logicalName string) *AWS_MediaConnect_Gateway {
	return &AWS_MediaConnect_Gateway{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MediaConnect_Gateway) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MediaConnect_Gateway) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MediaConnect_Gateway) GetType() string {
	return AWS_MediaConnect_Gateway__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MediaConnect_Gateway) Set__LogicalName(v string) *AWS_MediaConnect_Gateway {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MediaConnect_Gateway) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MediaConnect_Gateway {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MediaConnect_Gateway) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MediaConnect_Gateway {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MediaConnect_Gateway) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MediaConnect_Gateway {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MediaConnect_Gateway) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MediaConnect_Gateway {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MediaConnect_Gateway) Set__RequestedOutputs(v []cfz.Output) *AWS_MediaConnect_Gateway {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MediaConnect_Gateway) Add__RequestedOutputs(v ...cfz.Output) *AWS_MediaConnect_Gateway {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__EgressCidrBlocks updates property "EgressCidrBlocks".
func (t *AWS_MediaConnect_Gateway) Set__EgressCidrBlocks(v cfz.ExpressionSlice[string]) *AWS_MediaConnect_Gateway {
	t.EgressCidrBlocks = v
	return t
}

// SetS__EgressCidrBlocks updates property "EgressCidrBlocks".
func (t *AWS_MediaConnect_Gateway) SetS__EgressCidrBlocks(v ...cfz.Expression[string]) *AWS_MediaConnect_Gateway {
	t.EgressCidrBlocks = cfz.S(v...)
	return t
}

// SetSV__EgressCidrBlocks updates property "EgressCidrBlocks".
func (t *AWS_MediaConnect_Gateway) SetSV__EgressCidrBlocks(v ...string) *AWS_MediaConnect_Gateway {
	t.EgressCidrBlocks = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_MediaConnect_Gateway) Set__Name(v cfz.Expression[string]) *AWS_MediaConnect_Gateway {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_MediaConnect_Gateway) SetV__Name(v string) *AWS_MediaConnect_Gateway {
	t.Name = cfz.V(v)
	return t
}

// Set__Networks updates property "Networks".
func (t *AWS_MediaConnect_Gateway) Set__Networks(v cfz.ExpressionSlice[AWS_MediaConnect_Gateway_GatewayNetwork]) *AWS_MediaConnect_Gateway {
	t.Networks = v
	return t
}

// SetS__Networks updates property "Networks".
func (t *AWS_MediaConnect_Gateway) SetS__Networks(v ...cfz.Expression[AWS_MediaConnect_Gateway_GatewayNetwork]) *AWS_MediaConnect_Gateway {
	t.Networks = cfz.S(v...)
	return t
}

// SetSV__Networks updates property "Networks".
func (t *AWS_MediaConnect_Gateway) SetSV__Networks(v ...AWS_MediaConnect_Gateway_GatewayNetwork) *AWS_MediaConnect_Gateway {
	t.Networks = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MediaConnect_Gateway) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__GatewayArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: GatewayArn
func (t *AWS_MediaConnect_Gateway) GetAtt__GatewayArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaConnect_Gateway__AttributesMap.GatewayArn))
}

// GetAtt__GatewayState returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: GatewayState
func (t *AWS_MediaConnect_Gateway) GetAtt__GatewayState() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaConnect_Gateway__AttributesMap.GatewayState))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MediaConnect_Gateway) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__GatewayArn returns a conventionally configured output for an attribute of this resource.
// Attribute: GatewayArn
func (t *AWS_MediaConnect_Gateway) GetConventionalOutputAtt__GatewayArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttGatewayArn", t.GetAtt__GatewayArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__GatewayState returns a conventionally configured output for an attribute of this resource.
// Attribute: GatewayState
func (t *AWS_MediaConnect_Gateway) GetConventionalOutputAtt__GatewayState(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttGatewayState", t.GetAtt__GatewayState())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MediaConnect_Gateway) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MediaConnect_Gateway

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

func (t *AWS_MediaConnect_Gateway) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
