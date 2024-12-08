// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_connect

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Connect_SecurityKey)(nil)
	_ cfz.Resource                   = (*AWS_Connect_SecurityKey)(nil)
)

const (
	// AWS_Connect_SecurityKey__Type is the CloudFormation type for AWS::Connect::SecurityKey.
	AWS_Connect_SecurityKey__Type = "AWS::Connect::SecurityKey"
)

var (
	// AWS_Connect_SecurityKey__AttributesMap reports all the CloudFormation attributes for AWS::Connect::SecurityKey.
	AWS_Connect_SecurityKey__AttributesMap = struct {
		AssociationId string
	}{
		AssociationId: "AssociationId",
	}

	// AWS_Connect_SecurityKey__AttributesSlice reports all the CloudFormation attributes for AWS::Connect::SecurityKey.
	AWS_Connect_SecurityKey__AttributesSlice = []string{
		AWS_Connect_SecurityKey__AttributesMap.AssociationId,
	}
)

var (
	// AWS_Connect_SecurityKey__PropertiesMap reports all the CloudFormation properties for AWS::Connect::SecurityKey.
	AWS_Connect_SecurityKey__PropertiesMap = struct {
		InstanceId string
		Key        string
	}{
		InstanceId: "InstanceId",
		Key:        "Key",
	}

	// AWS_Connect_SecurityKey__PropertiesSlice reports all the CloudFormation properties for AWS::Connect::SecurityKey.
	AWS_Connect_SecurityKey__PropertiesSlice = []string{
		AWS_Connect_SecurityKey__PropertiesMap.InstanceId,
		AWS_Connect_SecurityKey__PropertiesMap.Key,
	}
)

// AWS_Connect_SecurityKey is a binding for AWS::Connect::SecurityKey.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securitykey.html
type AWS_Connect_SecurityKey struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// InstanceId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securitykey.html#cfn-connect-securitykey-instanceid
	InstanceId cfz.Expression[string] `json:"InstanceId,omitempty"`

	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-securitykey.html#cfn-connect-securitykey-key
	Key cfz.Expression[string] `json:"Key,omitempty"`
}

// New__AWS_Connect_SecurityKey initializes a new *AWS_Connect_SecurityKey.
func New__AWS_Connect_SecurityKey(logicalName string) *AWS_Connect_SecurityKey {
	return &AWS_Connect_SecurityKey{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Connect_SecurityKey) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Connect_SecurityKey) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Connect_SecurityKey) GetType() string {
	return AWS_Connect_SecurityKey__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Connect_SecurityKey) Set__LogicalName(v string) *AWS_Connect_SecurityKey {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Connect_SecurityKey) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Connect_SecurityKey {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Connect_SecurityKey) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Connect_SecurityKey {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Connect_SecurityKey) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Connect_SecurityKey {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Connect_SecurityKey) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Connect_SecurityKey {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Connect_SecurityKey) Set__RequestedOutputs(v []cfz.Output) *AWS_Connect_SecurityKey {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Connect_SecurityKey) Add__RequestedOutputs(v ...cfz.Output) *AWS_Connect_SecurityKey {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__InstanceId updates property "InstanceId".
func (t *AWS_Connect_SecurityKey) Set__InstanceId(v cfz.Expression[string]) *AWS_Connect_SecurityKey {
	t.InstanceId = v
	return t
}

// SetV__InstanceId updates property "InstanceId".
func (t *AWS_Connect_SecurityKey) SetV__InstanceId(v string) *AWS_Connect_SecurityKey {
	t.InstanceId = cfz.V(v)
	return t
}

// Set__Key updates property "Key".
func (t *AWS_Connect_SecurityKey) Set__Key(v cfz.Expression[string]) *AWS_Connect_SecurityKey {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t *AWS_Connect_SecurityKey) SetV__Key(v string) *AWS_Connect_SecurityKey {
	t.Key = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Connect_SecurityKey) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AssociationId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AssociationId
func (t *AWS_Connect_SecurityKey) GetAtt__AssociationId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Connect_SecurityKey__AttributesMap.AssociationId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Connect_SecurityKey) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AssociationId returns a conventionally configured output for an attribute of this resource.
// Attribute: AssociationId
func (t *AWS_Connect_SecurityKey) GetConventionalOutputAtt__AssociationId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAssociationId", t.GetAtt__AssociationId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Connect_SecurityKey) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Connect_SecurityKey

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

func (t *AWS_Connect_SecurityKey) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
