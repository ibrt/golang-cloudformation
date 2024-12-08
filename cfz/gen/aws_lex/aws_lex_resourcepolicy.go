// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_lex

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Lex_ResourcePolicy)(nil)
	_ cfz.Resource                   = (*AWS_Lex_ResourcePolicy)(nil)
)

const (
	// AWS_Lex_ResourcePolicy__Type is the CloudFormation type for AWS::Lex::ResourcePolicy.
	AWS_Lex_ResourcePolicy__Type = "AWS::Lex::ResourcePolicy"
)

var (
	// AWS_Lex_ResourcePolicy__AttributesMap reports all the CloudFormation attributes for AWS::Lex::ResourcePolicy.
	AWS_Lex_ResourcePolicy__AttributesMap = struct {
		Id         string
		RevisionId string
	}{
		Id:         "Id",
		RevisionId: "RevisionId",
	}

	// AWS_Lex_ResourcePolicy__AttributesSlice reports all the CloudFormation attributes for AWS::Lex::ResourcePolicy.
	AWS_Lex_ResourcePolicy__AttributesSlice = []string{
		AWS_Lex_ResourcePolicy__AttributesMap.Id,
		AWS_Lex_ResourcePolicy__AttributesMap.RevisionId,
	}
)

var (
	// AWS_Lex_ResourcePolicy__PropertiesMap reports all the CloudFormation properties for AWS::Lex::ResourcePolicy.
	AWS_Lex_ResourcePolicy__PropertiesMap = struct {
		Policy      string
		ResourceArn string
	}{
		Policy:      "Policy",
		ResourceArn: "ResourceArn",
	}

	// AWS_Lex_ResourcePolicy__PropertiesSlice reports all the CloudFormation properties for AWS::Lex::ResourcePolicy.
	AWS_Lex_ResourcePolicy__PropertiesSlice = []string{
		AWS_Lex_ResourcePolicy__PropertiesMap.Policy,
		AWS_Lex_ResourcePolicy__PropertiesMap.ResourceArn,
	}
)

// AWS_Lex_ResourcePolicy is a binding for AWS::Lex::ResourcePolicy.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-resourcepolicy.html
type AWS_Lex_ResourcePolicy struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Policy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-resourcepolicy.html#cfn-lex-resourcepolicy-policy
	Policy cfz.Expression[json.RawMessage] `json:"Policy,omitempty"`

	// ResourceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lex-resourcepolicy.html#cfn-lex-resourcepolicy-resourcearn
	ResourceArn cfz.Expression[string] `json:"ResourceArn,omitempty"`
}

// New__AWS_Lex_ResourcePolicy initializes a new *AWS_Lex_ResourcePolicy.
func New__AWS_Lex_ResourcePolicy(logicalName string) *AWS_Lex_ResourcePolicy {
	return &AWS_Lex_ResourcePolicy{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Lex_ResourcePolicy) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Lex_ResourcePolicy) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Lex_ResourcePolicy) GetType() string {
	return AWS_Lex_ResourcePolicy__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Lex_ResourcePolicy) Set__LogicalName(v string) *AWS_Lex_ResourcePolicy {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Lex_ResourcePolicy) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Lex_ResourcePolicy {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Lex_ResourcePolicy) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Lex_ResourcePolicy {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Lex_ResourcePolicy) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Lex_ResourcePolicy {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Lex_ResourcePolicy) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Lex_ResourcePolicy {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Lex_ResourcePolicy) Set__RequestedOutputs(v []cfz.Output) *AWS_Lex_ResourcePolicy {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Lex_ResourcePolicy) Add__RequestedOutputs(v ...cfz.Output) *AWS_Lex_ResourcePolicy {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Policy updates property "Policy".
func (t *AWS_Lex_ResourcePolicy) Set__Policy(v cfz.Expression[json.RawMessage]) *AWS_Lex_ResourcePolicy {
	t.Policy = v
	return t
}

// SetV__Policy updates property "Policy".
func (t *AWS_Lex_ResourcePolicy) SetV__Policy(v json.RawMessage) *AWS_Lex_ResourcePolicy {
	t.Policy = cfz.V(v)
	return t
}

// Set__ResourceArn updates property "ResourceArn".
func (t *AWS_Lex_ResourcePolicy) Set__ResourceArn(v cfz.Expression[string]) *AWS_Lex_ResourcePolicy {
	t.ResourceArn = v
	return t
}

// SetV__ResourceArn updates property "ResourceArn".
func (t *AWS_Lex_ResourcePolicy) SetV__ResourceArn(v string) *AWS_Lex_ResourcePolicy {
	t.ResourceArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Lex_ResourcePolicy) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Lex_ResourcePolicy) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lex_ResourcePolicy__AttributesMap.Id))
}

// GetAtt__RevisionId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RevisionId
func (t *AWS_Lex_ResourcePolicy) GetAtt__RevisionId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Lex_ResourcePolicy__AttributesMap.RevisionId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Lex_ResourcePolicy) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Lex_ResourcePolicy) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RevisionId returns a conventionally configured output for an attribute of this resource.
// Attribute: RevisionId
func (t *AWS_Lex_ResourcePolicy) GetConventionalOutputAtt__RevisionId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRevisionId", t.GetAtt__RevisionId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Lex_ResourcePolicy) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Lex_ResourcePolicy

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

func (t *AWS_Lex_ResourcePolicy) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
