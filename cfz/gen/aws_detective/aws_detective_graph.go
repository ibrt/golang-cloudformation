// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_detective

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Detective_Graph)(nil)
	_ cfz.Resource                   = (*AWS_Detective_Graph)(nil)
)

const (
	// AWS_Detective_Graph__Type is the CloudFormation type for AWS::Detective::Graph.
	AWS_Detective_Graph__Type = "AWS::Detective::Graph"
)

var (
	// AWS_Detective_Graph__AttributesMap reports all the CloudFormation attributes for AWS::Detective::Graph.
	AWS_Detective_Graph__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Detective_Graph__AttributesSlice reports all the CloudFormation attributes for AWS::Detective::Graph.
	AWS_Detective_Graph__AttributesSlice = []string{
		AWS_Detective_Graph__AttributesMap.Arn,
	}
)

var (
	// AWS_Detective_Graph__PropertiesMap reports all the CloudFormation properties for AWS::Detective::Graph.
	AWS_Detective_Graph__PropertiesMap = struct {
		AutoEnableMembers string
		Tags              string
	}{
		AutoEnableMembers: "AutoEnableMembers",
		Tags:              "Tags",
	}

	// AWS_Detective_Graph__PropertiesSlice reports all the CloudFormation properties for AWS::Detective::Graph.
	AWS_Detective_Graph__PropertiesSlice = []string{
		AWS_Detective_Graph__PropertiesMap.AutoEnableMembers,
		AWS_Detective_Graph__PropertiesMap.Tags,
	}
)

// AWS_Detective_Graph is a binding for AWS::Detective::Graph.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html
type AWS_Detective_Graph struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AutoEnableMembers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html#cfn-detective-graph-autoenablemembers
	AutoEnableMembers cfz.Expression[bool] `json:"AutoEnableMembers,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-graph.html#cfn-detective-graph-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Detective_Graph initializes a new *AWS_Detective_Graph.
func New__AWS_Detective_Graph(logicalName string) *AWS_Detective_Graph {
	return &AWS_Detective_Graph{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Detective_Graph) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Detective_Graph) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Detective_Graph) GetType() string {
	return AWS_Detective_Graph__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Detective_Graph) Set__LogicalName(v string) *AWS_Detective_Graph {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Detective_Graph) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Detective_Graph {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Detective_Graph) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Detective_Graph {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Detective_Graph) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Detective_Graph {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Detective_Graph) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Detective_Graph {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Detective_Graph) Set__RequestedOutputs(v []cfz.Output) *AWS_Detective_Graph {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Detective_Graph) Add__RequestedOutputs(v ...cfz.Output) *AWS_Detective_Graph {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AutoEnableMembers updates property "AutoEnableMembers".
func (t *AWS_Detective_Graph) Set__AutoEnableMembers(v cfz.Expression[bool]) *AWS_Detective_Graph {
	t.AutoEnableMembers = v
	return t
}

// SetV__AutoEnableMembers updates property "AutoEnableMembers".
func (t *AWS_Detective_Graph) SetV__AutoEnableMembers(v bool) *AWS_Detective_Graph {
	t.AutoEnableMembers = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Detective_Graph) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Detective_Graph {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Detective_Graph) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Detective_Graph {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Detective_Graph) SetSV__Tags(v ...cfz.Tag) *AWS_Detective_Graph {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Detective_Graph) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Detective_Graph) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Detective_Graph__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Detective_Graph) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Detective_Graph) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Detective_Graph) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Detective_Graph

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

func (t *AWS_Detective_Graph) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
