// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_databrew

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_DataBrew_Schedule)(nil)
	_ cfz.Resource                   = (*AWS_DataBrew_Schedule)(nil)
)

const (
	// AWS_DataBrew_Schedule__Type is the CloudFormation type for AWS::DataBrew::Schedule.
	AWS_DataBrew_Schedule__Type = "AWS::DataBrew::Schedule"
)

var (
	// AWS_DataBrew_Schedule__PropertiesMap reports all the CloudFormation properties for AWS::DataBrew::Schedule.
	AWS_DataBrew_Schedule__PropertiesMap = struct {
		CronExpression string
		JobNames       string
		Name           string
		Tags           string
	}{
		CronExpression: "CronExpression",
		JobNames:       "JobNames",
		Name:           "Name",
		Tags:           "Tags",
	}

	// AWS_DataBrew_Schedule__PropertiesSlice reports all the CloudFormation properties for AWS::DataBrew::Schedule.
	AWS_DataBrew_Schedule__PropertiesSlice = []string{
		AWS_DataBrew_Schedule__PropertiesMap.CronExpression,
		AWS_DataBrew_Schedule__PropertiesMap.JobNames,
		AWS_DataBrew_Schedule__PropertiesMap.Name,
		AWS_DataBrew_Schedule__PropertiesMap.Tags,
	}
)

// AWS_DataBrew_Schedule is a binding for AWS::DataBrew::Schedule.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html
type AWS_DataBrew_Schedule struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CronExpression is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-cronexpression
	CronExpression cfz.Expression[string] `json:"CronExpression,omitempty"`

	// JobNames is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-jobnames
	JobNames cfz.ExpressionSlice[string] `json:"JobNames,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_DataBrew_Schedule initializes a new *AWS_DataBrew_Schedule.
func New__AWS_DataBrew_Schedule(logicalName string) *AWS_DataBrew_Schedule {
	return &AWS_DataBrew_Schedule{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_DataBrew_Schedule) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_DataBrew_Schedule) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_DataBrew_Schedule) GetType() string {
	return AWS_DataBrew_Schedule__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_DataBrew_Schedule) Set__LogicalName(v string) *AWS_DataBrew_Schedule {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_DataBrew_Schedule) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_DataBrew_Schedule {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_DataBrew_Schedule) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_DataBrew_Schedule {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_DataBrew_Schedule) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_DataBrew_Schedule {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_DataBrew_Schedule) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_DataBrew_Schedule {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_DataBrew_Schedule) Set__RequestedOutputs(v []cfz.Output) *AWS_DataBrew_Schedule {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_DataBrew_Schedule) Add__RequestedOutputs(v ...cfz.Output) *AWS_DataBrew_Schedule {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CronExpression updates property "CronExpression".
func (t *AWS_DataBrew_Schedule) Set__CronExpression(v cfz.Expression[string]) *AWS_DataBrew_Schedule {
	t.CronExpression = v
	return t
}

// SetV__CronExpression updates property "CronExpression".
func (t *AWS_DataBrew_Schedule) SetV__CronExpression(v string) *AWS_DataBrew_Schedule {
	t.CronExpression = cfz.V(v)
	return t
}

// Set__JobNames updates property "JobNames".
func (t *AWS_DataBrew_Schedule) Set__JobNames(v cfz.ExpressionSlice[string]) *AWS_DataBrew_Schedule {
	t.JobNames = v
	return t
}

// SetS__JobNames updates property "JobNames".
func (t *AWS_DataBrew_Schedule) SetS__JobNames(v ...cfz.Expression[string]) *AWS_DataBrew_Schedule {
	t.JobNames = cfz.S(v...)
	return t
}

// SetSV__JobNames updates property "JobNames".
func (t *AWS_DataBrew_Schedule) SetSV__JobNames(v ...string) *AWS_DataBrew_Schedule {
	t.JobNames = cfz.SV(v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_DataBrew_Schedule) Set__Name(v cfz.Expression[string]) *AWS_DataBrew_Schedule {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_DataBrew_Schedule) SetV__Name(v string) *AWS_DataBrew_Schedule {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_DataBrew_Schedule) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_DataBrew_Schedule {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_DataBrew_Schedule) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_DataBrew_Schedule {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_DataBrew_Schedule) SetSV__Tags(v ...cfz.Tag) *AWS_DataBrew_Schedule {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_DataBrew_Schedule) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_DataBrew_Schedule) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_DataBrew_Schedule) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_DataBrew_Schedule

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

func (t *AWS_DataBrew_Schedule) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
