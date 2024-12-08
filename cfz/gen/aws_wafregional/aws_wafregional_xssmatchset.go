// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafregional

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_WAFRegional_XssMatchSet)(nil)
	_ cfz.Resource                   = (*AWS_WAFRegional_XssMatchSet)(nil)
)

const (
	// AWS_WAFRegional_XssMatchSet__Type is the CloudFormation type for AWS::WAFRegional::XssMatchSet.
	AWS_WAFRegional_XssMatchSet__Type = "AWS::WAFRegional::XssMatchSet"
)

var (
	// AWS_WAFRegional_XssMatchSet__PropertiesMap reports all the CloudFormation properties for AWS::WAFRegional::XssMatchSet.
	AWS_WAFRegional_XssMatchSet__PropertiesMap = struct {
		Name           string
		XssMatchTuples string
	}{
		Name:           "Name",
		XssMatchTuples: "XssMatchTuples",
	}

	// AWS_WAFRegional_XssMatchSet__PropertiesSlice reports all the CloudFormation properties for AWS::WAFRegional::XssMatchSet.
	AWS_WAFRegional_XssMatchSet__PropertiesSlice = []string{
		AWS_WAFRegional_XssMatchSet__PropertiesMap.Name,
		AWS_WAFRegional_XssMatchSet__PropertiesMap.XssMatchTuples,
	}
)

// AWS_WAFRegional_XssMatchSet is a binding for AWS::WAFRegional::XssMatchSet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html
type AWS_WAFRegional_XssMatchSet struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// XssMatchTuples is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html#cfn-wafregional-xssmatchset-xssmatchtuples
	XssMatchTuples cfz.ExpressionSlice[AWS_WAFRegional_XssMatchSet_XssMatchTuple] `json:"XssMatchTuples,omitempty"`
}

// New__AWS_WAFRegional_XssMatchSet initializes a new *AWS_WAFRegional_XssMatchSet.
func New__AWS_WAFRegional_XssMatchSet(logicalName string) *AWS_WAFRegional_XssMatchSet {
	return &AWS_WAFRegional_XssMatchSet{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_WAFRegional_XssMatchSet) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_WAFRegional_XssMatchSet) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_WAFRegional_XssMatchSet) GetType() string {
	return AWS_WAFRegional_XssMatchSet__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_WAFRegional_XssMatchSet) Set__LogicalName(v string) *AWS_WAFRegional_XssMatchSet {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_WAFRegional_XssMatchSet) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_WAFRegional_XssMatchSet {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_WAFRegional_XssMatchSet) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_WAFRegional_XssMatchSet {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_WAFRegional_XssMatchSet) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_WAFRegional_XssMatchSet {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_WAFRegional_XssMatchSet) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_WAFRegional_XssMatchSet {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_WAFRegional_XssMatchSet) Set__RequestedOutputs(v []cfz.Output) *AWS_WAFRegional_XssMatchSet {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_WAFRegional_XssMatchSet) Add__RequestedOutputs(v ...cfz.Output) *AWS_WAFRegional_XssMatchSet {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_WAFRegional_XssMatchSet) Set__Name(v cfz.Expression[string]) *AWS_WAFRegional_XssMatchSet {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_WAFRegional_XssMatchSet) SetV__Name(v string) *AWS_WAFRegional_XssMatchSet {
	t.Name = cfz.V(v)
	return t
}

// Set__XssMatchTuples updates property "XssMatchTuples".
func (t *AWS_WAFRegional_XssMatchSet) Set__XssMatchTuples(v cfz.ExpressionSlice[AWS_WAFRegional_XssMatchSet_XssMatchTuple]) *AWS_WAFRegional_XssMatchSet {
	t.XssMatchTuples = v
	return t
}

// SetS__XssMatchTuples updates property "XssMatchTuples".
func (t *AWS_WAFRegional_XssMatchSet) SetS__XssMatchTuples(v ...cfz.Expression[AWS_WAFRegional_XssMatchSet_XssMatchTuple]) *AWS_WAFRegional_XssMatchSet {
	t.XssMatchTuples = cfz.S(v...)
	return t
}

// SetSV__XssMatchTuples updates property "XssMatchTuples".
func (t *AWS_WAFRegional_XssMatchSet) SetSV__XssMatchTuples(v ...AWS_WAFRegional_XssMatchSet_XssMatchTuple) *AWS_WAFRegional_XssMatchSet {
	t.XssMatchTuples = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_WAFRegional_XssMatchSet) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_WAFRegional_XssMatchSet) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_WAFRegional_XssMatchSet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_WAFRegional_XssMatchSet

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

func (t *AWS_WAFRegional_XssMatchSet) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
