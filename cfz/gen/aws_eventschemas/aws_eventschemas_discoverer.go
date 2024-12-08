// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_eventschemas

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EventSchemas_Discoverer)(nil)
	_ cfz.Resource                   = (*AWS_EventSchemas_Discoverer)(nil)
)

const (
	// AWS_EventSchemas_Discoverer__Type is the CloudFormation type for AWS::EventSchemas::Discoverer.
	AWS_EventSchemas_Discoverer__Type = "AWS::EventSchemas::Discoverer"
)

var (
	// AWS_EventSchemas_Discoverer__AttributesMap reports all the CloudFormation attributes for AWS::EventSchemas::Discoverer.
	AWS_EventSchemas_Discoverer__AttributesMap = struct {
		DiscovererArn string
		DiscovererId  string
		State         string
	}{
		DiscovererArn: "DiscovererArn",
		DiscovererId:  "DiscovererId",
		State:         "State",
	}

	// AWS_EventSchemas_Discoverer__AttributesSlice reports all the CloudFormation attributes for AWS::EventSchemas::Discoverer.
	AWS_EventSchemas_Discoverer__AttributesSlice = []string{
		AWS_EventSchemas_Discoverer__AttributesMap.DiscovererArn,
		AWS_EventSchemas_Discoverer__AttributesMap.DiscovererId,
		AWS_EventSchemas_Discoverer__AttributesMap.State,
	}
)

var (
	// AWS_EventSchemas_Discoverer__PropertiesMap reports all the CloudFormation properties for AWS::EventSchemas::Discoverer.
	AWS_EventSchemas_Discoverer__PropertiesMap = struct {
		CrossAccount string
		Description  string
		SourceArn    string
		Tags         string
	}{
		CrossAccount: "CrossAccount",
		Description:  "Description",
		SourceArn:    "SourceArn",
		Tags:         "Tags",
	}

	// AWS_EventSchemas_Discoverer__PropertiesSlice reports all the CloudFormation properties for AWS::EventSchemas::Discoverer.
	AWS_EventSchemas_Discoverer__PropertiesSlice = []string{
		AWS_EventSchemas_Discoverer__PropertiesMap.CrossAccount,
		AWS_EventSchemas_Discoverer__PropertiesMap.Description,
		AWS_EventSchemas_Discoverer__PropertiesMap.SourceArn,
		AWS_EventSchemas_Discoverer__PropertiesMap.Tags,
	}
)

// AWS_EventSchemas_Discoverer is a binding for AWS::EventSchemas::Discoverer.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html
type AWS_EventSchemas_Discoverer struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CrossAccount is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-crossaccount
	CrossAccount cfz.Expression[bool] `json:"CrossAccount,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// SourceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-sourcearn
	SourceArn cfz.Expression[string] `json:"SourceArn,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-tags
	Tags cfz.ExpressionSlice[AWS_EventSchemas_Discoverer_TagsEntry] `json:"Tags,omitempty"`
}

// New__AWS_EventSchemas_Discoverer initializes a new *AWS_EventSchemas_Discoverer.
func New__AWS_EventSchemas_Discoverer(logicalName string) *AWS_EventSchemas_Discoverer {
	return &AWS_EventSchemas_Discoverer{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EventSchemas_Discoverer) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EventSchemas_Discoverer) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EventSchemas_Discoverer) GetType() string {
	return AWS_EventSchemas_Discoverer__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EventSchemas_Discoverer) Set__LogicalName(v string) *AWS_EventSchemas_Discoverer {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EventSchemas_Discoverer) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EventSchemas_Discoverer {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EventSchemas_Discoverer) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EventSchemas_Discoverer {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EventSchemas_Discoverer) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EventSchemas_Discoverer {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EventSchemas_Discoverer) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EventSchemas_Discoverer {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EventSchemas_Discoverer) Set__RequestedOutputs(v []cfz.Output) *AWS_EventSchemas_Discoverer {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EventSchemas_Discoverer) Add__RequestedOutputs(v ...cfz.Output) *AWS_EventSchemas_Discoverer {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CrossAccount updates property "CrossAccount".
func (t *AWS_EventSchemas_Discoverer) Set__CrossAccount(v cfz.Expression[bool]) *AWS_EventSchemas_Discoverer {
	t.CrossAccount = v
	return t
}

// SetV__CrossAccount updates property "CrossAccount".
func (t *AWS_EventSchemas_Discoverer) SetV__CrossAccount(v bool) *AWS_EventSchemas_Discoverer {
	t.CrossAccount = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_EventSchemas_Discoverer) Set__Description(v cfz.Expression[string]) *AWS_EventSchemas_Discoverer {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_EventSchemas_Discoverer) SetV__Description(v string) *AWS_EventSchemas_Discoverer {
	t.Description = cfz.V(v)
	return t
}

// Set__SourceArn updates property "SourceArn".
func (t *AWS_EventSchemas_Discoverer) Set__SourceArn(v cfz.Expression[string]) *AWS_EventSchemas_Discoverer {
	t.SourceArn = v
	return t
}

// SetV__SourceArn updates property "SourceArn".
func (t *AWS_EventSchemas_Discoverer) SetV__SourceArn(v string) *AWS_EventSchemas_Discoverer {
	t.SourceArn = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EventSchemas_Discoverer) Set__Tags(v cfz.ExpressionSlice[AWS_EventSchemas_Discoverer_TagsEntry]) *AWS_EventSchemas_Discoverer {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EventSchemas_Discoverer) SetS__Tags(v ...cfz.Expression[AWS_EventSchemas_Discoverer_TagsEntry]) *AWS_EventSchemas_Discoverer {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EventSchemas_Discoverer) SetSV__Tags(v ...AWS_EventSchemas_Discoverer_TagsEntry) *AWS_EventSchemas_Discoverer {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EventSchemas_Discoverer) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DiscovererArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DiscovererArn
func (t *AWS_EventSchemas_Discoverer) GetAtt__DiscovererArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Discoverer__AttributesMap.DiscovererArn))
}

// GetAtt__DiscovererId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DiscovererId
func (t *AWS_EventSchemas_Discoverer) GetAtt__DiscovererId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Discoverer__AttributesMap.DiscovererId))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_EventSchemas_Discoverer) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EventSchemas_Discoverer__AttributesMap.State))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EventSchemas_Discoverer) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DiscovererArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DiscovererArn
func (t *AWS_EventSchemas_Discoverer) GetConventionalOutputAtt__DiscovererArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDiscovererArn", t.GetAtt__DiscovererArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DiscovererId returns a conventionally configured output for an attribute of this resource.
// Attribute: DiscovererId
func (t *AWS_EventSchemas_Discoverer) GetConventionalOutputAtt__DiscovererId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDiscovererId", t.GetAtt__DiscovererId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_EventSchemas_Discoverer) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EventSchemas_Discoverer) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EventSchemas_Discoverer

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

func (t *AWS_EventSchemas_Discoverer) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
