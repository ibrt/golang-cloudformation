// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_apigateway

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ApiGateway_UsagePlanKey)(nil)
	_ cfz.Resource                   = (*AWS_ApiGateway_UsagePlanKey)(nil)
)

const (
	// AWS_ApiGateway_UsagePlanKey__Type is the CloudFormation type for AWS::ApiGateway::UsagePlanKey.
	AWS_ApiGateway_UsagePlanKey__Type = "AWS::ApiGateway::UsagePlanKey"
)

var (
	// AWS_ApiGateway_UsagePlanKey__AttributesMap reports all the CloudFormation attributes for AWS::ApiGateway::UsagePlanKey.
	AWS_ApiGateway_UsagePlanKey__AttributesMap = struct {
		Id string
	}{
		Id: "Id",
	}

	// AWS_ApiGateway_UsagePlanKey__AttributesSlice reports all the CloudFormation attributes for AWS::ApiGateway::UsagePlanKey.
	AWS_ApiGateway_UsagePlanKey__AttributesSlice = []string{
		AWS_ApiGateway_UsagePlanKey__AttributesMap.Id,
	}
)

var (
	// AWS_ApiGateway_UsagePlanKey__PropertiesMap reports all the CloudFormation properties for AWS::ApiGateway::UsagePlanKey.
	AWS_ApiGateway_UsagePlanKey__PropertiesMap = struct {
		KeyId       string
		KeyType     string
		UsagePlanId string
	}{
		KeyId:       "KeyId",
		KeyType:     "KeyType",
		UsagePlanId: "UsagePlanId",
	}

	// AWS_ApiGateway_UsagePlanKey__PropertiesSlice reports all the CloudFormation properties for AWS::ApiGateway::UsagePlanKey.
	AWS_ApiGateway_UsagePlanKey__PropertiesSlice = []string{
		AWS_ApiGateway_UsagePlanKey__PropertiesMap.KeyId,
		AWS_ApiGateway_UsagePlanKey__PropertiesMap.KeyType,
		AWS_ApiGateway_UsagePlanKey__PropertiesMap.UsagePlanId,
	}
)

// AWS_ApiGateway_UsagePlanKey is a binding for AWS::ApiGateway::UsagePlanKey.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html
type AWS_ApiGateway_UsagePlanKey struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// KeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keyid
	KeyId cfz.Expression[string] `json:"KeyId,omitempty"`

	// KeyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keytype
	KeyType cfz.Expression[string] `json:"KeyType,omitempty"`

	// UsagePlanId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-usageplanid
	UsagePlanId cfz.Expression[string] `json:"UsagePlanId,omitempty"`
}

// New__AWS_ApiGateway_UsagePlanKey initializes a new *AWS_ApiGateway_UsagePlanKey.
func New__AWS_ApiGateway_UsagePlanKey(logicalName string) *AWS_ApiGateway_UsagePlanKey {
	return &AWS_ApiGateway_UsagePlanKey{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ApiGateway_UsagePlanKey) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ApiGateway_UsagePlanKey) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ApiGateway_UsagePlanKey) GetType() string {
	return AWS_ApiGateway_UsagePlanKey__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ApiGateway_UsagePlanKey) Set__LogicalName(v string) *AWS_ApiGateway_UsagePlanKey {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ApiGateway_UsagePlanKey) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ApiGateway_UsagePlanKey {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ApiGateway_UsagePlanKey) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ApiGateway_UsagePlanKey {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ApiGateway_UsagePlanKey) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ApiGateway_UsagePlanKey {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ApiGateway_UsagePlanKey) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ApiGateway_UsagePlanKey {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ApiGateway_UsagePlanKey) Set__RequestedOutputs(v []cfz.Output) *AWS_ApiGateway_UsagePlanKey {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ApiGateway_UsagePlanKey) Add__RequestedOutputs(v ...cfz.Output) *AWS_ApiGateway_UsagePlanKey {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__KeyId updates property "KeyId".
func (t *AWS_ApiGateway_UsagePlanKey) Set__KeyId(v cfz.Expression[string]) *AWS_ApiGateway_UsagePlanKey {
	t.KeyId = v
	return t
}

// SetV__KeyId updates property "KeyId".
func (t *AWS_ApiGateway_UsagePlanKey) SetV__KeyId(v string) *AWS_ApiGateway_UsagePlanKey {
	t.KeyId = cfz.V(v)
	return t
}

// Set__KeyType updates property "KeyType".
func (t *AWS_ApiGateway_UsagePlanKey) Set__KeyType(v cfz.Expression[string]) *AWS_ApiGateway_UsagePlanKey {
	t.KeyType = v
	return t
}

// SetV__KeyType updates property "KeyType".
func (t *AWS_ApiGateway_UsagePlanKey) SetV__KeyType(v string) *AWS_ApiGateway_UsagePlanKey {
	t.KeyType = cfz.V(v)
	return t
}

// Set__UsagePlanId updates property "UsagePlanId".
func (t *AWS_ApiGateway_UsagePlanKey) Set__UsagePlanId(v cfz.Expression[string]) *AWS_ApiGateway_UsagePlanKey {
	t.UsagePlanId = v
	return t
}

// SetV__UsagePlanId updates property "UsagePlanId".
func (t *AWS_ApiGateway_UsagePlanKey) SetV__UsagePlanId(v string) *AWS_ApiGateway_UsagePlanKey {
	t.UsagePlanId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ApiGateway_UsagePlanKey) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_ApiGateway_UsagePlanKey) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_ApiGateway_UsagePlanKey__AttributesMap.Id))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ApiGateway_UsagePlanKey) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_ApiGateway_UsagePlanKey) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ApiGateway_UsagePlanKey) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ApiGateway_UsagePlanKey

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

func (t *AWS_ApiGateway_UsagePlanKey) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
