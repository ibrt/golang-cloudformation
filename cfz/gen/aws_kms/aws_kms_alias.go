// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kms

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_KMS_Alias)(nil)
	_ cfz.Resource                   = (*AWS_KMS_Alias)(nil)
)

const (
	// AWS_KMS_Alias__Type is the CloudFormation type for AWS::KMS::Alias.
	AWS_KMS_Alias__Type = "AWS::KMS::Alias"
)

var (
	// AWS_KMS_Alias__PropertiesMap reports all the CloudFormation properties for AWS::KMS::Alias.
	AWS_KMS_Alias__PropertiesMap = struct {
		AliasName   string
		TargetKeyId string
	}{
		AliasName:   "AliasName",
		TargetKeyId: "TargetKeyId",
	}

	// AWS_KMS_Alias__PropertiesSlice reports all the CloudFormation properties for AWS::KMS::Alias.
	AWS_KMS_Alias__PropertiesSlice = []string{
		AWS_KMS_Alias__PropertiesMap.AliasName,
		AWS_KMS_Alias__PropertiesMap.TargetKeyId,
	}
)

// AWS_KMS_Alias is a binding for AWS::KMS::Alias.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type AWS_KMS_Alias struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AliasName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName cfz.Expression[string] `json:"AliasName,omitempty"`

	// TargetKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyId cfz.Expression[string] `json:"TargetKeyId,omitempty"`
}

// New__AWS_KMS_Alias initializes a new *AWS_KMS_Alias.
func New__AWS_KMS_Alias(logicalName string) *AWS_KMS_Alias {
	return &AWS_KMS_Alias{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_KMS_Alias) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_KMS_Alias) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_KMS_Alias) GetType() string {
	return AWS_KMS_Alias__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_KMS_Alias) Set__LogicalName(v string) *AWS_KMS_Alias {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_KMS_Alias) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_KMS_Alias {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_KMS_Alias) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_KMS_Alias {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_KMS_Alias) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_KMS_Alias {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_KMS_Alias) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_KMS_Alias {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_KMS_Alias) Set__RequestedOutputs(v []cfz.Output) *AWS_KMS_Alias {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_KMS_Alias) Add__RequestedOutputs(v ...cfz.Output) *AWS_KMS_Alias {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AliasName updates property "AliasName".
func (t *AWS_KMS_Alias) Set__AliasName(v cfz.Expression[string]) *AWS_KMS_Alias {
	t.AliasName = v
	return t
}

// SetV__AliasName updates property "AliasName".
func (t *AWS_KMS_Alias) SetV__AliasName(v string) *AWS_KMS_Alias {
	t.AliasName = cfz.V(v)
	return t
}

// Set__TargetKeyId updates property "TargetKeyId".
func (t *AWS_KMS_Alias) Set__TargetKeyId(v cfz.Expression[string]) *AWS_KMS_Alias {
	t.TargetKeyId = v
	return t
}

// SetV__TargetKeyId updates property "TargetKeyId".
func (t *AWS_KMS_Alias) SetV__TargetKeyId(v string) *AWS_KMS_Alias {
	t.TargetKeyId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_KMS_Alias) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_KMS_Alias) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_KMS_Alias) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_KMS_Alias

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

func (t *AWS_KMS_Alias) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
