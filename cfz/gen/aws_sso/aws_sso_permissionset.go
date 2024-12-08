// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sso

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_SSO_PermissionSet)(nil)
	_ cfz.Resource                   = (*AWS_SSO_PermissionSet)(nil)
)

const (
	// AWS_SSO_PermissionSet__Type is the CloudFormation type for AWS::SSO::PermissionSet.
	AWS_SSO_PermissionSet__Type = "AWS::SSO::PermissionSet"
)

var (
	// AWS_SSO_PermissionSet__AttributesMap reports all the CloudFormation attributes for AWS::SSO::PermissionSet.
	AWS_SSO_PermissionSet__AttributesMap = struct {
		PermissionSetArn string
	}{
		PermissionSetArn: "PermissionSetArn",
	}

	// AWS_SSO_PermissionSet__AttributesSlice reports all the CloudFormation attributes for AWS::SSO::PermissionSet.
	AWS_SSO_PermissionSet__AttributesSlice = []string{
		AWS_SSO_PermissionSet__AttributesMap.PermissionSetArn,
	}
)

var (
	// AWS_SSO_PermissionSet__PropertiesMap reports all the CloudFormation properties for AWS::SSO::PermissionSet.
	AWS_SSO_PermissionSet__PropertiesMap = struct {
		CustomerManagedPolicyReferences string
		Description                     string
		InlinePolicy                    string
		InstanceArn                     string
		ManagedPolicies                 string
		Name                            string
		PermissionsBoundary             string
		RelayStateType                  string
		SessionDuration                 string
		Tags                            string
	}{
		CustomerManagedPolicyReferences: "CustomerManagedPolicyReferences",
		Description:                     "Description",
		InlinePolicy:                    "InlinePolicy",
		InstanceArn:                     "InstanceArn",
		ManagedPolicies:                 "ManagedPolicies",
		Name:                            "Name",
		PermissionsBoundary:             "PermissionsBoundary",
		RelayStateType:                  "RelayStateType",
		SessionDuration:                 "SessionDuration",
		Tags:                            "Tags",
	}

	// AWS_SSO_PermissionSet__PropertiesSlice reports all the CloudFormation properties for AWS::SSO::PermissionSet.
	AWS_SSO_PermissionSet__PropertiesSlice = []string{
		AWS_SSO_PermissionSet__PropertiesMap.CustomerManagedPolicyReferences,
		AWS_SSO_PermissionSet__PropertiesMap.Description,
		AWS_SSO_PermissionSet__PropertiesMap.InlinePolicy,
		AWS_SSO_PermissionSet__PropertiesMap.InstanceArn,
		AWS_SSO_PermissionSet__PropertiesMap.ManagedPolicies,
		AWS_SSO_PermissionSet__PropertiesMap.Name,
		AWS_SSO_PermissionSet__PropertiesMap.PermissionsBoundary,
		AWS_SSO_PermissionSet__PropertiesMap.RelayStateType,
		AWS_SSO_PermissionSet__PropertiesMap.SessionDuration,
		AWS_SSO_PermissionSet__PropertiesMap.Tags,
	}
)

// AWS_SSO_PermissionSet is a binding for AWS::SSO::PermissionSet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html
type AWS_SSO_PermissionSet struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CustomerManagedPolicyReferences is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-customermanagedpolicyreferences
	CustomerManagedPolicyReferences cfz.ExpressionSlice[AWS_SSO_PermissionSet_CustomerManagedPolicyReference] `json:"CustomerManagedPolicyReferences,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// InlinePolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-inlinepolicy
	InlinePolicy cfz.Expression[json.RawMessage] `json:"InlinePolicy,omitempty"`

	// InstanceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-instancearn
	InstanceArn cfz.Expression[string] `json:"InstanceArn,omitempty"`

	// ManagedPolicies is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-managedpolicies
	ManagedPolicies cfz.ExpressionSlice[string] `json:"ManagedPolicies,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// PermissionsBoundary is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-permissionsboundary
	PermissionsBoundary cfz.Expression[AWS_SSO_PermissionSet_PermissionsBoundary] `json:"PermissionsBoundary,omitempty"`

	// RelayStateType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-relaystatetype
	RelayStateType cfz.Expression[string] `json:"RelayStateType,omitempty"`

	// SessionDuration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-sessionduration
	SessionDuration cfz.Expression[string] `json:"SessionDuration,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_SSO_PermissionSet initializes a new *AWS_SSO_PermissionSet.
func New__AWS_SSO_PermissionSet(logicalName string) *AWS_SSO_PermissionSet {
	return &AWS_SSO_PermissionSet{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_SSO_PermissionSet) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_SSO_PermissionSet) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_SSO_PermissionSet) GetType() string {
	return AWS_SSO_PermissionSet__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_SSO_PermissionSet) Set__LogicalName(v string) *AWS_SSO_PermissionSet {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_SSO_PermissionSet) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_SSO_PermissionSet {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_SSO_PermissionSet) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_SSO_PermissionSet {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_SSO_PermissionSet) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_SSO_PermissionSet {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_SSO_PermissionSet) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_SSO_PermissionSet {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_SSO_PermissionSet) Set__RequestedOutputs(v []cfz.Output) *AWS_SSO_PermissionSet {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_SSO_PermissionSet) Add__RequestedOutputs(v ...cfz.Output) *AWS_SSO_PermissionSet {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CustomerManagedPolicyReferences updates property "CustomerManagedPolicyReferences".
func (t *AWS_SSO_PermissionSet) Set__CustomerManagedPolicyReferences(v cfz.ExpressionSlice[AWS_SSO_PermissionSet_CustomerManagedPolicyReference]) *AWS_SSO_PermissionSet {
	t.CustomerManagedPolicyReferences = v
	return t
}

// SetS__CustomerManagedPolicyReferences updates property "CustomerManagedPolicyReferences".
func (t *AWS_SSO_PermissionSet) SetS__CustomerManagedPolicyReferences(v ...cfz.Expression[AWS_SSO_PermissionSet_CustomerManagedPolicyReference]) *AWS_SSO_PermissionSet {
	t.CustomerManagedPolicyReferences = cfz.S(v...)
	return t
}

// SetSV__CustomerManagedPolicyReferences updates property "CustomerManagedPolicyReferences".
func (t *AWS_SSO_PermissionSet) SetSV__CustomerManagedPolicyReferences(v ...AWS_SSO_PermissionSet_CustomerManagedPolicyReference) *AWS_SSO_PermissionSet {
	t.CustomerManagedPolicyReferences = cfz.SV(v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_SSO_PermissionSet) Set__Description(v cfz.Expression[string]) *AWS_SSO_PermissionSet {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_SSO_PermissionSet) SetV__Description(v string) *AWS_SSO_PermissionSet {
	t.Description = cfz.V(v)
	return t
}

// Set__InlinePolicy updates property "InlinePolicy".
func (t *AWS_SSO_PermissionSet) Set__InlinePolicy(v cfz.Expression[json.RawMessage]) *AWS_SSO_PermissionSet {
	t.InlinePolicy = v
	return t
}

// SetV__InlinePolicy updates property "InlinePolicy".
func (t *AWS_SSO_PermissionSet) SetV__InlinePolicy(v json.RawMessage) *AWS_SSO_PermissionSet {
	t.InlinePolicy = cfz.V(v)
	return t
}

// Set__InstanceArn updates property "InstanceArn".
func (t *AWS_SSO_PermissionSet) Set__InstanceArn(v cfz.Expression[string]) *AWS_SSO_PermissionSet {
	t.InstanceArn = v
	return t
}

// SetV__InstanceArn updates property "InstanceArn".
func (t *AWS_SSO_PermissionSet) SetV__InstanceArn(v string) *AWS_SSO_PermissionSet {
	t.InstanceArn = cfz.V(v)
	return t
}

// Set__ManagedPolicies updates property "ManagedPolicies".
func (t *AWS_SSO_PermissionSet) Set__ManagedPolicies(v cfz.ExpressionSlice[string]) *AWS_SSO_PermissionSet {
	t.ManagedPolicies = v
	return t
}

// SetS__ManagedPolicies updates property "ManagedPolicies".
func (t *AWS_SSO_PermissionSet) SetS__ManagedPolicies(v ...cfz.Expression[string]) *AWS_SSO_PermissionSet {
	t.ManagedPolicies = cfz.S(v...)
	return t
}

// SetSV__ManagedPolicies updates property "ManagedPolicies".
func (t *AWS_SSO_PermissionSet) SetSV__ManagedPolicies(v ...string) *AWS_SSO_PermissionSet {
	t.ManagedPolicies = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_SSO_PermissionSet) Set__Name(v cfz.Expression[string]) *AWS_SSO_PermissionSet {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_SSO_PermissionSet) SetV__Name(v string) *AWS_SSO_PermissionSet {
	t.Name = cfz.V(v)
	return t
}

// Set__PermissionsBoundary updates property "PermissionsBoundary".
func (t *AWS_SSO_PermissionSet) Set__PermissionsBoundary(v cfz.Expression[AWS_SSO_PermissionSet_PermissionsBoundary]) *AWS_SSO_PermissionSet {
	t.PermissionsBoundary = v
	return t
}

// SetV__PermissionsBoundary updates property "PermissionsBoundary".
func (t *AWS_SSO_PermissionSet) SetV__PermissionsBoundary(v AWS_SSO_PermissionSet_PermissionsBoundary) *AWS_SSO_PermissionSet {
	t.PermissionsBoundary = cfz.V(v)
	return t
}

// Set__RelayStateType updates property "RelayStateType".
func (t *AWS_SSO_PermissionSet) Set__RelayStateType(v cfz.Expression[string]) *AWS_SSO_PermissionSet {
	t.RelayStateType = v
	return t
}

// SetV__RelayStateType updates property "RelayStateType".
func (t *AWS_SSO_PermissionSet) SetV__RelayStateType(v string) *AWS_SSO_PermissionSet {
	t.RelayStateType = cfz.V(v)
	return t
}

// Set__SessionDuration updates property "SessionDuration".
func (t *AWS_SSO_PermissionSet) Set__SessionDuration(v cfz.Expression[string]) *AWS_SSO_PermissionSet {
	t.SessionDuration = v
	return t
}

// SetV__SessionDuration updates property "SessionDuration".
func (t *AWS_SSO_PermissionSet) SetV__SessionDuration(v string) *AWS_SSO_PermissionSet {
	t.SessionDuration = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_SSO_PermissionSet) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_SSO_PermissionSet {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_SSO_PermissionSet) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_SSO_PermissionSet {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_SSO_PermissionSet) SetSV__Tags(v ...cfz.Tag) *AWS_SSO_PermissionSet {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_SSO_PermissionSet) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__PermissionSetArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PermissionSetArn
func (t *AWS_SSO_PermissionSet) GetAtt__PermissionSetArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SSO_PermissionSet__AttributesMap.PermissionSetArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_SSO_PermissionSet) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PermissionSetArn returns a conventionally configured output for an attribute of this resource.
// Attribute: PermissionSetArn
func (t *AWS_SSO_PermissionSet) GetConventionalOutputAtt__PermissionSetArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPermissionSetArn", t.GetAtt__PermissionSetArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_SSO_PermissionSet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_SSO_PermissionSet

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

func (t *AWS_SSO_PermissionSet) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
