// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_guardduty

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_GuardDuty_Filter)(nil)
	_ cfz.Resource                   = (*AWS_GuardDuty_Filter)(nil)
)

const (
	// AWS_GuardDuty_Filter__Type is the CloudFormation type for AWS::GuardDuty::Filter.
	AWS_GuardDuty_Filter__Type = "AWS::GuardDuty::Filter"
)

var (
	// AWS_GuardDuty_Filter__PropertiesMap reports all the CloudFormation properties for AWS::GuardDuty::Filter.
	AWS_GuardDuty_Filter__PropertiesMap = struct {
		Action          string
		Description     string
		DetectorId      string
		FindingCriteria string
		Name            string
		Rank            string
		Tags            string
	}{
		Action:          "Action",
		Description:     "Description",
		DetectorId:      "DetectorId",
		FindingCriteria: "FindingCriteria",
		Name:            "Name",
		Rank:            "Rank",
		Tags:            "Tags",
	}

	// AWS_GuardDuty_Filter__PropertiesSlice reports all the CloudFormation properties for AWS::GuardDuty::Filter.
	AWS_GuardDuty_Filter__PropertiesSlice = []string{
		AWS_GuardDuty_Filter__PropertiesMap.Action,
		AWS_GuardDuty_Filter__PropertiesMap.Description,
		AWS_GuardDuty_Filter__PropertiesMap.DetectorId,
		AWS_GuardDuty_Filter__PropertiesMap.FindingCriteria,
		AWS_GuardDuty_Filter__PropertiesMap.Name,
		AWS_GuardDuty_Filter__PropertiesMap.Rank,
		AWS_GuardDuty_Filter__PropertiesMap.Tags,
	}
)

// AWS_GuardDuty_Filter is a binding for AWS::GuardDuty::Filter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html
type AWS_GuardDuty_Filter struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Action is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-action
	Action cfz.Expression[string] `json:"Action,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DetectorId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-detectorid
	DetectorId cfz.Expression[string] `json:"DetectorId,omitempty"`

	// FindingCriteria is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-findingcriteria
	FindingCriteria cfz.Expression[AWS_GuardDuty_Filter_FindingCriteria] `json:"FindingCriteria,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Rank is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-rank
	Rank cfz.Expression[int32] `json:"Rank,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-tags
	Tags cfz.ExpressionSlice[AWS_GuardDuty_Filter_TagItem] `json:"Tags,omitempty"`
}

// New__AWS_GuardDuty_Filter initializes a new *AWS_GuardDuty_Filter.
func New__AWS_GuardDuty_Filter(logicalName string) *AWS_GuardDuty_Filter {
	return &AWS_GuardDuty_Filter{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_GuardDuty_Filter) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_GuardDuty_Filter) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_GuardDuty_Filter) GetType() string {
	return AWS_GuardDuty_Filter__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_GuardDuty_Filter) Set__LogicalName(v string) *AWS_GuardDuty_Filter {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_GuardDuty_Filter) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_GuardDuty_Filter {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_GuardDuty_Filter) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_GuardDuty_Filter {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_GuardDuty_Filter) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_GuardDuty_Filter {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_GuardDuty_Filter) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_GuardDuty_Filter {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_GuardDuty_Filter) Set__RequestedOutputs(v []cfz.Output) *AWS_GuardDuty_Filter {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_GuardDuty_Filter) Add__RequestedOutputs(v ...cfz.Output) *AWS_GuardDuty_Filter {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Action updates property "Action".
func (t *AWS_GuardDuty_Filter) Set__Action(v cfz.Expression[string]) *AWS_GuardDuty_Filter {
	t.Action = v
	return t
}

// SetV__Action updates property "Action".
func (t *AWS_GuardDuty_Filter) SetV__Action(v string) *AWS_GuardDuty_Filter {
	t.Action = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_GuardDuty_Filter) Set__Description(v cfz.Expression[string]) *AWS_GuardDuty_Filter {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_GuardDuty_Filter) SetV__Description(v string) *AWS_GuardDuty_Filter {
	t.Description = cfz.V(v)
	return t
}

// Set__DetectorId updates property "DetectorId".
func (t *AWS_GuardDuty_Filter) Set__DetectorId(v cfz.Expression[string]) *AWS_GuardDuty_Filter {
	t.DetectorId = v
	return t
}

// SetV__DetectorId updates property "DetectorId".
func (t *AWS_GuardDuty_Filter) SetV__DetectorId(v string) *AWS_GuardDuty_Filter {
	t.DetectorId = cfz.V(v)
	return t
}

// Set__FindingCriteria updates property "FindingCriteria".
func (t *AWS_GuardDuty_Filter) Set__FindingCriteria(v cfz.Expression[AWS_GuardDuty_Filter_FindingCriteria]) *AWS_GuardDuty_Filter {
	t.FindingCriteria = v
	return t
}

// SetV__FindingCriteria updates property "FindingCriteria".
func (t *AWS_GuardDuty_Filter) SetV__FindingCriteria(v AWS_GuardDuty_Filter_FindingCriteria) *AWS_GuardDuty_Filter {
	t.FindingCriteria = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_GuardDuty_Filter) Set__Name(v cfz.Expression[string]) *AWS_GuardDuty_Filter {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_GuardDuty_Filter) SetV__Name(v string) *AWS_GuardDuty_Filter {
	t.Name = cfz.V(v)
	return t
}

// Set__Rank updates property "Rank".
func (t *AWS_GuardDuty_Filter) Set__Rank(v cfz.Expression[int32]) *AWS_GuardDuty_Filter {
	t.Rank = v
	return t
}

// SetV__Rank updates property "Rank".
func (t *AWS_GuardDuty_Filter) SetV__Rank(v int32) *AWS_GuardDuty_Filter {
	t.Rank = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_GuardDuty_Filter) Set__Tags(v cfz.ExpressionSlice[AWS_GuardDuty_Filter_TagItem]) *AWS_GuardDuty_Filter {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_GuardDuty_Filter) SetS__Tags(v ...cfz.Expression[AWS_GuardDuty_Filter_TagItem]) *AWS_GuardDuty_Filter {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_GuardDuty_Filter) SetSV__Tags(v ...AWS_GuardDuty_Filter_TagItem) *AWS_GuardDuty_Filter {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_GuardDuty_Filter) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_GuardDuty_Filter) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_GuardDuty_Filter) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_GuardDuty_Filter

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

func (t *AWS_GuardDuty_Filter) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
