// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_NetworkManager_CustomerGatewayAssociation)(nil)
	_ cfz.Resource                   = (*AWS_NetworkManager_CustomerGatewayAssociation)(nil)
)

const (
	// AWS_NetworkManager_CustomerGatewayAssociation__Type is the CloudFormation type for AWS::NetworkManager::CustomerGatewayAssociation.
	AWS_NetworkManager_CustomerGatewayAssociation__Type = "AWS::NetworkManager::CustomerGatewayAssociation"
)

var (
	// AWS_NetworkManager_CustomerGatewayAssociation__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::CustomerGatewayAssociation.
	AWS_NetworkManager_CustomerGatewayAssociation__PropertiesMap = struct {
		CustomerGatewayArn string
		DeviceId           string
		GlobalNetworkId    string
		LinkId             string
	}{
		CustomerGatewayArn: "CustomerGatewayArn",
		DeviceId:           "DeviceId",
		GlobalNetworkId:    "GlobalNetworkId",
		LinkId:             "LinkId",
	}

	// AWS_NetworkManager_CustomerGatewayAssociation__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::CustomerGatewayAssociation.
	AWS_NetworkManager_CustomerGatewayAssociation__PropertiesSlice = []string{
		AWS_NetworkManager_CustomerGatewayAssociation__PropertiesMap.CustomerGatewayArn,
		AWS_NetworkManager_CustomerGatewayAssociation__PropertiesMap.DeviceId,
		AWS_NetworkManager_CustomerGatewayAssociation__PropertiesMap.GlobalNetworkId,
		AWS_NetworkManager_CustomerGatewayAssociation__PropertiesMap.LinkId,
	}
)

// AWS_NetworkManager_CustomerGatewayAssociation is a binding for AWS::NetworkManager::CustomerGatewayAssociation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html
type AWS_NetworkManager_CustomerGatewayAssociation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CustomerGatewayArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-customergatewayarn
	CustomerGatewayArn cfz.Expression[string] `json:"CustomerGatewayArn,omitempty"`

	// DeviceId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-deviceid
	DeviceId cfz.Expression[string] `json:"DeviceId,omitempty"`

	// GlobalNetworkId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-globalnetworkid
	GlobalNetworkId cfz.Expression[string] `json:"GlobalNetworkId,omitempty"`

	// LinkId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html#cfn-networkmanager-customergatewayassociation-linkid
	LinkId cfz.Expression[string] `json:"LinkId,omitempty"`
}

// New__AWS_NetworkManager_CustomerGatewayAssociation initializes a new *AWS_NetworkManager_CustomerGatewayAssociation.
func New__AWS_NetworkManager_CustomerGatewayAssociation(logicalName string) *AWS_NetworkManager_CustomerGatewayAssociation {
	return &AWS_NetworkManager_CustomerGatewayAssociation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_NetworkManager_CustomerGatewayAssociation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_NetworkManager_CustomerGatewayAssociation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_NetworkManager_CustomerGatewayAssociation) GetType() string {
	return AWS_NetworkManager_CustomerGatewayAssociation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__LogicalName(v string) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__RequestedOutputs(v []cfz.Output) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Add__RequestedOutputs(v ...cfz.Output) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CustomerGatewayArn updates property "CustomerGatewayArn".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__CustomerGatewayArn(v cfz.Expression[string]) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.CustomerGatewayArn = v
	return t
}

// SetV__CustomerGatewayArn updates property "CustomerGatewayArn".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) SetV__CustomerGatewayArn(v string) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.CustomerGatewayArn = cfz.V(v)
	return t
}

// Set__DeviceId updates property "DeviceId".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__DeviceId(v cfz.Expression[string]) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.DeviceId = v
	return t
}

// SetV__DeviceId updates property "DeviceId".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) SetV__DeviceId(v string) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.DeviceId = cfz.V(v)
	return t
}

// Set__GlobalNetworkId updates property "GlobalNetworkId".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__GlobalNetworkId(v cfz.Expression[string]) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.GlobalNetworkId = v
	return t
}

// SetV__GlobalNetworkId updates property "GlobalNetworkId".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) SetV__GlobalNetworkId(v string) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.GlobalNetworkId = cfz.V(v)
	return t
}

// Set__LinkId updates property "LinkId".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Set__LinkId(v cfz.Expression[string]) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.LinkId = v
	return t
}

// SetV__LinkId updates property "LinkId".
func (t *AWS_NetworkManager_CustomerGatewayAssociation) SetV__LinkId(v string) *AWS_NetworkManager_CustomerGatewayAssociation {
	t.LinkId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_NetworkManager_CustomerGatewayAssociation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_NetworkManager_CustomerGatewayAssociation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_NetworkManager_CustomerGatewayAssociation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_NetworkManager_CustomerGatewayAssociation

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

func (t *AWS_NetworkManager_CustomerGatewayAssociation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
