// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EMR_WALWorkspace)(nil)
	_ cfz.Resource                   = (*AWS_EMR_WALWorkspace)(nil)
)

const (
	// AWS_EMR_WALWorkspace__Type is the CloudFormation type for AWS::EMR::WALWorkspace.
	AWS_EMR_WALWorkspace__Type = "AWS::EMR::WALWorkspace"
)

var (
	// AWS_EMR_WALWorkspace__PropertiesMap reports all the CloudFormation properties for AWS::EMR::WALWorkspace.
	AWS_EMR_WALWorkspace__PropertiesMap = struct {
		Tags             string
		WALWorkspaceName string
	}{
		Tags:             "Tags",
		WALWorkspaceName: "WALWorkspaceName",
	}

	// AWS_EMR_WALWorkspace__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::WALWorkspace.
	AWS_EMR_WALWorkspace__PropertiesSlice = []string{
		AWS_EMR_WALWorkspace__PropertiesMap.Tags,
		AWS_EMR_WALWorkspace__PropertiesMap.WALWorkspaceName,
	}
)

// AWS_EMR_WALWorkspace is a binding for AWS::EMR::WALWorkspace.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-walworkspace.html
type AWS_EMR_WALWorkspace struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-walworkspace.html#cfn-emr-walworkspace-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// WALWorkspaceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-walworkspace.html#cfn-emr-walworkspace-walworkspacename
	WALWorkspaceName cfz.Expression[string] `json:"WALWorkspaceName,omitempty"`
}

// New__AWS_EMR_WALWorkspace initializes a new *AWS_EMR_WALWorkspace.
func New__AWS_EMR_WALWorkspace(logicalName string) *AWS_EMR_WALWorkspace {
	return &AWS_EMR_WALWorkspace{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EMR_WALWorkspace) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EMR_WALWorkspace) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EMR_WALWorkspace) GetType() string {
	return AWS_EMR_WALWorkspace__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EMR_WALWorkspace) Set__LogicalName(v string) *AWS_EMR_WALWorkspace {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EMR_WALWorkspace) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EMR_WALWorkspace {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EMR_WALWorkspace) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EMR_WALWorkspace {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EMR_WALWorkspace) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EMR_WALWorkspace {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EMR_WALWorkspace) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EMR_WALWorkspace {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EMR_WALWorkspace) Set__RequestedOutputs(v []cfz.Output) *AWS_EMR_WALWorkspace {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EMR_WALWorkspace) Add__RequestedOutputs(v ...cfz.Output) *AWS_EMR_WALWorkspace {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EMR_WALWorkspace) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EMR_WALWorkspace {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EMR_WALWorkspace) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EMR_WALWorkspace {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EMR_WALWorkspace) SetSV__Tags(v ...cfz.Tag) *AWS_EMR_WALWorkspace {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__WALWorkspaceName updates property "WALWorkspaceName".
func (t *AWS_EMR_WALWorkspace) Set__WALWorkspaceName(v cfz.Expression[string]) *AWS_EMR_WALWorkspace {
	t.WALWorkspaceName = v
	return t
}

// SetV__WALWorkspaceName updates property "WALWorkspaceName".
func (t *AWS_EMR_WALWorkspace) SetV__WALWorkspaceName(v string) *AWS_EMR_WALWorkspace {
	t.WALWorkspaceName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EMR_WALWorkspace) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EMR_WALWorkspace) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EMR_WALWorkspace) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EMR_WALWorkspace

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

func (t *AWS_EMR_WALWorkspace) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
