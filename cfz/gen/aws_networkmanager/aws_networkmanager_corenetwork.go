// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkmanager

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_NetworkManager_CoreNetwork)(nil)
	_ cfz.Resource                   = (*AWS_NetworkManager_CoreNetwork)(nil)
)

const (
	// AWS_NetworkManager_CoreNetwork__Type is the CloudFormation type for AWS::NetworkManager::CoreNetwork.
	AWS_NetworkManager_CoreNetwork__Type = "AWS::NetworkManager::CoreNetwork"
)

var (
	// AWS_NetworkManager_CoreNetwork__AttributesMap reports all the CloudFormation attributes for AWS::NetworkManager::CoreNetwork.
	AWS_NetworkManager_CoreNetwork__AttributesMap = struct {
		CoreNetworkArn        string
		CoreNetworkId         string
		CreatedAt             string
		Edges                 string
		NetworkFunctionGroups string
		OwnerAccount          string
		Segments              string
		State                 string
	}{
		CoreNetworkArn:        "CoreNetworkArn",
		CoreNetworkId:         "CoreNetworkId",
		CreatedAt:             "CreatedAt",
		Edges:                 "Edges",
		NetworkFunctionGroups: "NetworkFunctionGroups",
		OwnerAccount:          "OwnerAccount",
		Segments:              "Segments",
		State:                 "State",
	}

	// AWS_NetworkManager_CoreNetwork__AttributesSlice reports all the CloudFormation attributes for AWS::NetworkManager::CoreNetwork.
	AWS_NetworkManager_CoreNetwork__AttributesSlice = []string{
		AWS_NetworkManager_CoreNetwork__AttributesMap.CoreNetworkArn,
		AWS_NetworkManager_CoreNetwork__AttributesMap.CoreNetworkId,
		AWS_NetworkManager_CoreNetwork__AttributesMap.CreatedAt,
		AWS_NetworkManager_CoreNetwork__AttributesMap.Edges,
		AWS_NetworkManager_CoreNetwork__AttributesMap.NetworkFunctionGroups,
		AWS_NetworkManager_CoreNetwork__AttributesMap.OwnerAccount,
		AWS_NetworkManager_CoreNetwork__AttributesMap.Segments,
		AWS_NetworkManager_CoreNetwork__AttributesMap.State,
	}
)

var (
	// AWS_NetworkManager_CoreNetwork__PropertiesMap reports all the CloudFormation properties for AWS::NetworkManager::CoreNetwork.
	AWS_NetworkManager_CoreNetwork__PropertiesMap = struct {
		Description     string
		GlobalNetworkId string
		PolicyDocument  string
		Tags            string
	}{
		Description:     "Description",
		GlobalNetworkId: "GlobalNetworkId",
		PolicyDocument:  "PolicyDocument",
		Tags:            "Tags",
	}

	// AWS_NetworkManager_CoreNetwork__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkManager::CoreNetwork.
	AWS_NetworkManager_CoreNetwork__PropertiesSlice = []string{
		AWS_NetworkManager_CoreNetwork__PropertiesMap.Description,
		AWS_NetworkManager_CoreNetwork__PropertiesMap.GlobalNetworkId,
		AWS_NetworkManager_CoreNetwork__PropertiesMap.PolicyDocument,
		AWS_NetworkManager_CoreNetwork__PropertiesMap.Tags,
	}
)

// AWS_NetworkManager_CoreNetwork is a binding for AWS::NetworkManager::CoreNetwork.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetwork.html
type AWS_NetworkManager_CoreNetwork struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetwork.html#cfn-networkmanager-corenetwork-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// GlobalNetworkId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetwork.html#cfn-networkmanager-corenetwork-globalnetworkid
	GlobalNetworkId cfz.Expression[string] `json:"GlobalNetworkId,omitempty"`

	// PolicyDocument is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetwork.html#cfn-networkmanager-corenetwork-policydocument
	PolicyDocument cfz.Expression[json.RawMessage] `json:"PolicyDocument,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-corenetwork.html#cfn-networkmanager-corenetwork-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_NetworkManager_CoreNetwork initializes a new *AWS_NetworkManager_CoreNetwork.
func New__AWS_NetworkManager_CoreNetwork(logicalName string) *AWS_NetworkManager_CoreNetwork {
	return &AWS_NetworkManager_CoreNetwork{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_NetworkManager_CoreNetwork) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_NetworkManager_CoreNetwork) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_NetworkManager_CoreNetwork) GetType() string {
	return AWS_NetworkManager_CoreNetwork__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_NetworkManager_CoreNetwork) Set__LogicalName(v string) *AWS_NetworkManager_CoreNetwork {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_NetworkManager_CoreNetwork) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_NetworkManager_CoreNetwork {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_NetworkManager_CoreNetwork) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_NetworkManager_CoreNetwork {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_NetworkManager_CoreNetwork) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_NetworkManager_CoreNetwork {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_NetworkManager_CoreNetwork) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_NetworkManager_CoreNetwork {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_NetworkManager_CoreNetwork) Set__RequestedOutputs(v []cfz.Output) *AWS_NetworkManager_CoreNetwork {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_NetworkManager_CoreNetwork) Add__RequestedOutputs(v ...cfz.Output) *AWS_NetworkManager_CoreNetwork {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_NetworkManager_CoreNetwork) Set__Description(v cfz.Expression[string]) *AWS_NetworkManager_CoreNetwork {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_NetworkManager_CoreNetwork) SetV__Description(v string) *AWS_NetworkManager_CoreNetwork {
	t.Description = cfz.V(v)
	return t
}

// Set__GlobalNetworkId updates property "GlobalNetworkId".
func (t *AWS_NetworkManager_CoreNetwork) Set__GlobalNetworkId(v cfz.Expression[string]) *AWS_NetworkManager_CoreNetwork {
	t.GlobalNetworkId = v
	return t
}

// SetV__GlobalNetworkId updates property "GlobalNetworkId".
func (t *AWS_NetworkManager_CoreNetwork) SetV__GlobalNetworkId(v string) *AWS_NetworkManager_CoreNetwork {
	t.GlobalNetworkId = cfz.V(v)
	return t
}

// Set__PolicyDocument updates property "PolicyDocument".
func (t *AWS_NetworkManager_CoreNetwork) Set__PolicyDocument(v cfz.Expression[json.RawMessage]) *AWS_NetworkManager_CoreNetwork {
	t.PolicyDocument = v
	return t
}

// SetV__PolicyDocument updates property "PolicyDocument".
func (t *AWS_NetworkManager_CoreNetwork) SetV__PolicyDocument(v json.RawMessage) *AWS_NetworkManager_CoreNetwork {
	t.PolicyDocument = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_NetworkManager_CoreNetwork) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_NetworkManager_CoreNetwork {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_NetworkManager_CoreNetwork) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_NetworkManager_CoreNetwork {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_NetworkManager_CoreNetwork) SetSV__Tags(v ...cfz.Tag) *AWS_NetworkManager_CoreNetwork {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_NetworkManager_CoreNetwork) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CoreNetworkArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CoreNetworkArn
func (t *AWS_NetworkManager_CoreNetwork) GetAtt__CoreNetworkArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.CoreNetworkArn))
}

// GetAtt__CoreNetworkId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CoreNetworkId
func (t *AWS_NetworkManager_CoreNetwork) GetAtt__CoreNetworkId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.CoreNetworkId))
}

// GetAtt__CreatedAt returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedAt
func (t *AWS_NetworkManager_CoreNetwork) GetAtt__CreatedAt() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.CreatedAt))
}

// GetAttSlice__Edges returns a $cfz.ExpressionSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkEdge] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Edges
func (t *AWS_NetworkManager_CoreNetwork) GetAttSlice__Edges() cfz.ExpressionSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkEdge] {
	return cfz.GetAttSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkEdge](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.Edges))
}

// GetAttSlice__NetworkFunctionGroups returns a $cfz.ExpressionSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkNetworkFunctionGroup] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NetworkFunctionGroups
func (t *AWS_NetworkManager_CoreNetwork) GetAttSlice__NetworkFunctionGroups() cfz.ExpressionSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkNetworkFunctionGroup] {
	return cfz.GetAttSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkNetworkFunctionGroup](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.NetworkFunctionGroups))
}

// GetAtt__OwnerAccount returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: OwnerAccount
func (t *AWS_NetworkManager_CoreNetwork) GetAtt__OwnerAccount() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.OwnerAccount))
}

// GetAttSlice__Segments returns a $cfz.ExpressionSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkSegment] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Segments
func (t *AWS_NetworkManager_CoreNetwork) GetAttSlice__Segments() cfz.ExpressionSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkSegment] {
	return cfz.GetAttSlice[AWS_NetworkManager_CoreNetwork_CoreNetworkSegment](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.Segments))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_NetworkManager_CoreNetwork) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_NetworkManager_CoreNetwork__AttributesMap.State))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CoreNetworkArn returns a conventionally configured output for an attribute of this resource.
// Attribute: CoreNetworkArn
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__CoreNetworkArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCoreNetworkArn", t.GetAtt__CoreNetworkArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CoreNetworkId returns a conventionally configured output for an attribute of this resource.
// Attribute: CoreNetworkId
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__CoreNetworkId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCoreNetworkId", t.GetAtt__CoreNetworkId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedAt returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedAt
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__CreatedAt(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedAt", t.GetAtt__CreatedAt())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Edges returns a conventionally configured output for an attribute of this resource.
// Attribute: Edges
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__Edges(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttEdges", t.GetAttSlice__Edges())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NetworkFunctionGroups returns a conventionally configured output for an attribute of this resource.
// Attribute: NetworkFunctionGroups
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__NetworkFunctionGroups(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNetworkFunctionGroups", t.GetAttSlice__NetworkFunctionGroups())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__OwnerAccount returns a conventionally configured output for an attribute of this resource.
// Attribute: OwnerAccount
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__OwnerAccount(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttOwnerAccount", t.GetAtt__OwnerAccount())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Segments returns a conventionally configured output for an attribute of this resource.
// Attribute: Segments
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__Segments(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSegments", t.GetAttSlice__Segments())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_NetworkManager_CoreNetwork) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_NetworkManager_CoreNetwork) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_NetworkManager_CoreNetwork

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

func (t *AWS_NetworkManager_CoreNetwork) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
