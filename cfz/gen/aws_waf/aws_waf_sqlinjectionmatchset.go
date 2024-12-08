// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_waf

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_WAF_SqlInjectionMatchSet)(nil)
	_ cfz.Resource                   = (*AWS_WAF_SqlInjectionMatchSet)(nil)
)

const (
	// AWS_WAF_SqlInjectionMatchSet__Type is the CloudFormation type for AWS::WAF::SqlInjectionMatchSet.
	AWS_WAF_SqlInjectionMatchSet__Type = "AWS::WAF::SqlInjectionMatchSet"
)

var (
	// AWS_WAF_SqlInjectionMatchSet__PropertiesMap reports all the CloudFormation properties for AWS::WAF::SqlInjectionMatchSet.
	AWS_WAF_SqlInjectionMatchSet__PropertiesMap = struct {
		Name                    string
		SqlInjectionMatchTuples string
	}{
		Name:                    "Name",
		SqlInjectionMatchTuples: "SqlInjectionMatchTuples",
	}

	// AWS_WAF_SqlInjectionMatchSet__PropertiesSlice reports all the CloudFormation properties for AWS::WAF::SqlInjectionMatchSet.
	AWS_WAF_SqlInjectionMatchSet__PropertiesSlice = []string{
		AWS_WAF_SqlInjectionMatchSet__PropertiesMap.Name,
		AWS_WAF_SqlInjectionMatchSet__PropertiesMap.SqlInjectionMatchTuples,
	}
)

// AWS_WAF_SqlInjectionMatchSet is a binding for AWS::WAF::SqlInjectionMatchSet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html
type AWS_WAF_SqlInjectionMatchSet struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// SqlInjectionMatchTuples is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples
	SqlInjectionMatchTuples cfz.ExpressionSlice[AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple] `json:"SqlInjectionMatchTuples,omitempty"`
}

// New__AWS_WAF_SqlInjectionMatchSet initializes a new *AWS_WAF_SqlInjectionMatchSet.
func New__AWS_WAF_SqlInjectionMatchSet(logicalName string) *AWS_WAF_SqlInjectionMatchSet {
	return &AWS_WAF_SqlInjectionMatchSet{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_WAF_SqlInjectionMatchSet) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_WAF_SqlInjectionMatchSet) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_WAF_SqlInjectionMatchSet) GetType() string {
	return AWS_WAF_SqlInjectionMatchSet__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__LogicalName(v string) *AWS_WAF_SqlInjectionMatchSet {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_WAF_SqlInjectionMatchSet {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_WAF_SqlInjectionMatchSet) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_WAF_SqlInjectionMatchSet {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_WAF_SqlInjectionMatchSet {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_WAF_SqlInjectionMatchSet {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__RequestedOutputs(v []cfz.Output) *AWS_WAF_SqlInjectionMatchSet {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_WAF_SqlInjectionMatchSet) Add__RequestedOutputs(v ...cfz.Output) *AWS_WAF_SqlInjectionMatchSet {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__Name(v cfz.Expression[string]) *AWS_WAF_SqlInjectionMatchSet {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_WAF_SqlInjectionMatchSet) SetV__Name(v string) *AWS_WAF_SqlInjectionMatchSet {
	t.Name = cfz.V(v)
	return t
}

// Set__SqlInjectionMatchTuples updates property "SqlInjectionMatchTuples".
func (t *AWS_WAF_SqlInjectionMatchSet) Set__SqlInjectionMatchTuples(v cfz.ExpressionSlice[AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple]) *AWS_WAF_SqlInjectionMatchSet {
	t.SqlInjectionMatchTuples = v
	return t
}

// SetS__SqlInjectionMatchTuples updates property "SqlInjectionMatchTuples".
func (t *AWS_WAF_SqlInjectionMatchSet) SetS__SqlInjectionMatchTuples(v ...cfz.Expression[AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple]) *AWS_WAF_SqlInjectionMatchSet {
	t.SqlInjectionMatchTuples = cfz.S(v...)
	return t
}

// SetSV__SqlInjectionMatchTuples updates property "SqlInjectionMatchTuples".
func (t *AWS_WAF_SqlInjectionMatchSet) SetSV__SqlInjectionMatchTuples(v ...AWS_WAF_SqlInjectionMatchSet_SqlInjectionMatchTuple) *AWS_WAF_SqlInjectionMatchSet {
	t.SqlInjectionMatchTuples = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_WAF_SqlInjectionMatchSet) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_WAF_SqlInjectionMatchSet) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_WAF_SqlInjectionMatchSet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_WAF_SqlInjectionMatchSet

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

func (t *AWS_WAF_SqlInjectionMatchSet) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
