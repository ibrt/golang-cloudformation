// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ivs

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IVS_PublicKey)(nil)
	_ cfz.Resource                   = (*AWS_IVS_PublicKey)(nil)
)

const (
	// AWS_IVS_PublicKey__Type is the CloudFormation type for AWS::IVS::PublicKey.
	AWS_IVS_PublicKey__Type = "AWS::IVS::PublicKey"
)

var (
	// AWS_IVS_PublicKey__AttributesMap reports all the CloudFormation attributes for AWS::IVS::PublicKey.
	AWS_IVS_PublicKey__AttributesMap = struct {
		Arn         string
		Fingerprint string
	}{
		Arn:         "Arn",
		Fingerprint: "Fingerprint",
	}

	// AWS_IVS_PublicKey__AttributesSlice reports all the CloudFormation attributes for AWS::IVS::PublicKey.
	AWS_IVS_PublicKey__AttributesSlice = []string{
		AWS_IVS_PublicKey__AttributesMap.Arn,
		AWS_IVS_PublicKey__AttributesMap.Fingerprint,
	}
)

var (
	// AWS_IVS_PublicKey__PropertiesMap reports all the CloudFormation properties for AWS::IVS::PublicKey.
	AWS_IVS_PublicKey__PropertiesMap = struct {
		Name              string
		PublicKeyMaterial string
		Tags              string
	}{
		Name:              "Name",
		PublicKeyMaterial: "PublicKeyMaterial",
		Tags:              "Tags",
	}

	// AWS_IVS_PublicKey__PropertiesSlice reports all the CloudFormation properties for AWS::IVS::PublicKey.
	AWS_IVS_PublicKey__PropertiesSlice = []string{
		AWS_IVS_PublicKey__PropertiesMap.Name,
		AWS_IVS_PublicKey__PropertiesMap.PublicKeyMaterial,
		AWS_IVS_PublicKey__PropertiesMap.Tags,
	}
)

// AWS_IVS_PublicKey is a binding for AWS::IVS::PublicKey.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-publickey.html
type AWS_IVS_PublicKey struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-publickey.html#cfn-ivs-publickey-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// PublicKeyMaterial is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-publickey.html#cfn-ivs-publickey-publickeymaterial
	PublicKeyMaterial cfz.Expression[string] `json:"PublicKeyMaterial,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-publickey.html#cfn-ivs-publickey-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_IVS_PublicKey initializes a new *AWS_IVS_PublicKey.
func New__AWS_IVS_PublicKey(logicalName string) *AWS_IVS_PublicKey {
	return &AWS_IVS_PublicKey{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IVS_PublicKey) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IVS_PublicKey) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IVS_PublicKey) GetType() string {
	return AWS_IVS_PublicKey__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IVS_PublicKey) Set__LogicalName(v string) *AWS_IVS_PublicKey {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IVS_PublicKey) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IVS_PublicKey {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IVS_PublicKey) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IVS_PublicKey {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IVS_PublicKey) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IVS_PublicKey {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IVS_PublicKey) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IVS_PublicKey {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IVS_PublicKey) Set__RequestedOutputs(v []cfz.Output) *AWS_IVS_PublicKey {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IVS_PublicKey) Add__RequestedOutputs(v ...cfz.Output) *AWS_IVS_PublicKey {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_IVS_PublicKey) Set__Name(v cfz.Expression[string]) *AWS_IVS_PublicKey {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_IVS_PublicKey) SetV__Name(v string) *AWS_IVS_PublicKey {
	t.Name = cfz.V(v)
	return t
}

// Set__PublicKeyMaterial updates property "PublicKeyMaterial".
func (t *AWS_IVS_PublicKey) Set__PublicKeyMaterial(v cfz.Expression[string]) *AWS_IVS_PublicKey {
	t.PublicKeyMaterial = v
	return t
}

// SetV__PublicKeyMaterial updates property "PublicKeyMaterial".
func (t *AWS_IVS_PublicKey) SetV__PublicKeyMaterial(v string) *AWS_IVS_PublicKey {
	t.PublicKeyMaterial = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IVS_PublicKey) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IVS_PublicKey {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IVS_PublicKey) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IVS_PublicKey {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IVS_PublicKey) SetSV__Tags(v ...cfz.Tag) *AWS_IVS_PublicKey {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IVS_PublicKey) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_IVS_PublicKey) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IVS_PublicKey__AttributesMap.Arn))
}

// GetAtt__Fingerprint returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Fingerprint
func (t *AWS_IVS_PublicKey) GetAtt__Fingerprint() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IVS_PublicKey__AttributesMap.Fingerprint))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IVS_PublicKey) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_IVS_PublicKey) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Fingerprint returns a conventionally configured output for an attribute of this resource.
// Attribute: Fingerprint
func (t *AWS_IVS_PublicKey) GetConventionalOutputAtt__Fingerprint(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFingerprint", t.GetAtt__Fingerprint())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IVS_PublicKey) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IVS_PublicKey

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

func (t *AWS_IVS_PublicKey) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
