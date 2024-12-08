// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_logs

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Logs_SubscriptionFilter)(nil)
	_ cfz.Resource                   = (*AWS_Logs_SubscriptionFilter)(nil)
)

const (
	// AWS_Logs_SubscriptionFilter__Type is the CloudFormation type for AWS::Logs::SubscriptionFilter.
	AWS_Logs_SubscriptionFilter__Type = "AWS::Logs::SubscriptionFilter"
)

var (
	// AWS_Logs_SubscriptionFilter__PropertiesMap reports all the CloudFormation properties for AWS::Logs::SubscriptionFilter.
	AWS_Logs_SubscriptionFilter__PropertiesMap = struct {
		DestinationArn string
		Distribution   string
		FilterName     string
		FilterPattern  string
		LogGroupName   string
		RoleArn        string
	}{
		DestinationArn: "DestinationArn",
		Distribution:   "Distribution",
		FilterName:     "FilterName",
		FilterPattern:  "FilterPattern",
		LogGroupName:   "LogGroupName",
		RoleArn:        "RoleArn",
	}

	// AWS_Logs_SubscriptionFilter__PropertiesSlice reports all the CloudFormation properties for AWS::Logs::SubscriptionFilter.
	AWS_Logs_SubscriptionFilter__PropertiesSlice = []string{
		AWS_Logs_SubscriptionFilter__PropertiesMap.DestinationArn,
		AWS_Logs_SubscriptionFilter__PropertiesMap.Distribution,
		AWS_Logs_SubscriptionFilter__PropertiesMap.FilterName,
		AWS_Logs_SubscriptionFilter__PropertiesMap.FilterPattern,
		AWS_Logs_SubscriptionFilter__PropertiesMap.LogGroupName,
		AWS_Logs_SubscriptionFilter__PropertiesMap.RoleArn,
	}
)

// AWS_Logs_SubscriptionFilter is a binding for AWS::Logs::SubscriptionFilter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
type AWS_Logs_SubscriptionFilter struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DestinationArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-destinationarn
	DestinationArn cfz.Expression[string] `json:"DestinationArn,omitempty"`

	// Distribution is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-distribution
	Distribution cfz.Expression[string] `json:"Distribution,omitempty"`

	// FilterName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-filtername
	FilterName cfz.Expression[string] `json:"FilterName,omitempty"`

	// FilterPattern is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-filterpattern
	FilterPattern cfz.Expression[string] `json:"FilterPattern,omitempty"`

	// LogGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-loggroupname
	LogGroupName cfz.Expression[string] `json:"LogGroupName,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-logs-subscriptionfilter-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`
}

// New__AWS_Logs_SubscriptionFilter initializes a new *AWS_Logs_SubscriptionFilter.
func New__AWS_Logs_SubscriptionFilter(logicalName string) *AWS_Logs_SubscriptionFilter {
	return &AWS_Logs_SubscriptionFilter{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Logs_SubscriptionFilter) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Logs_SubscriptionFilter) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Logs_SubscriptionFilter) GetType() string {
	return AWS_Logs_SubscriptionFilter__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Logs_SubscriptionFilter) Set__LogicalName(v string) *AWS_Logs_SubscriptionFilter {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Logs_SubscriptionFilter) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Logs_SubscriptionFilter {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Logs_SubscriptionFilter) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Logs_SubscriptionFilter {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Logs_SubscriptionFilter) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Logs_SubscriptionFilter {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Logs_SubscriptionFilter) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Logs_SubscriptionFilter {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Logs_SubscriptionFilter) Set__RequestedOutputs(v []cfz.Output) *AWS_Logs_SubscriptionFilter {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Logs_SubscriptionFilter) Add__RequestedOutputs(v ...cfz.Output) *AWS_Logs_SubscriptionFilter {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DestinationArn updates property "DestinationArn".
func (t *AWS_Logs_SubscriptionFilter) Set__DestinationArn(v cfz.Expression[string]) *AWS_Logs_SubscriptionFilter {
	t.DestinationArn = v
	return t
}

// SetV__DestinationArn updates property "DestinationArn".
func (t *AWS_Logs_SubscriptionFilter) SetV__DestinationArn(v string) *AWS_Logs_SubscriptionFilter {
	t.DestinationArn = cfz.V(v)
	return t
}

// Set__Distribution updates property "Distribution".
func (t *AWS_Logs_SubscriptionFilter) Set__Distribution(v cfz.Expression[string]) *AWS_Logs_SubscriptionFilter {
	t.Distribution = v
	return t
}

// SetV__Distribution updates property "Distribution".
func (t *AWS_Logs_SubscriptionFilter) SetV__Distribution(v string) *AWS_Logs_SubscriptionFilter {
	t.Distribution = cfz.V(v)
	return t
}

// Set__FilterName updates property "FilterName".
func (t *AWS_Logs_SubscriptionFilter) Set__FilterName(v cfz.Expression[string]) *AWS_Logs_SubscriptionFilter {
	t.FilterName = v
	return t
}

// SetV__FilterName updates property "FilterName".
func (t *AWS_Logs_SubscriptionFilter) SetV__FilterName(v string) *AWS_Logs_SubscriptionFilter {
	t.FilterName = cfz.V(v)
	return t
}

// Set__FilterPattern updates property "FilterPattern".
func (t *AWS_Logs_SubscriptionFilter) Set__FilterPattern(v cfz.Expression[string]) *AWS_Logs_SubscriptionFilter {
	t.FilterPattern = v
	return t
}

// SetV__FilterPattern updates property "FilterPattern".
func (t *AWS_Logs_SubscriptionFilter) SetV__FilterPattern(v string) *AWS_Logs_SubscriptionFilter {
	t.FilterPattern = cfz.V(v)
	return t
}

// Set__LogGroupName updates property "LogGroupName".
func (t *AWS_Logs_SubscriptionFilter) Set__LogGroupName(v cfz.Expression[string]) *AWS_Logs_SubscriptionFilter {
	t.LogGroupName = v
	return t
}

// SetV__LogGroupName updates property "LogGroupName".
func (t *AWS_Logs_SubscriptionFilter) SetV__LogGroupName(v string) *AWS_Logs_SubscriptionFilter {
	t.LogGroupName = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_Logs_SubscriptionFilter) Set__RoleArn(v cfz.Expression[string]) *AWS_Logs_SubscriptionFilter {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_Logs_SubscriptionFilter) SetV__RoleArn(v string) *AWS_Logs_SubscriptionFilter {
	t.RoleArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Logs_SubscriptionFilter) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Logs_SubscriptionFilter) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Logs_SubscriptionFilter) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Logs_SubscriptionFilter

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

func (t *AWS_Logs_SubscriptionFilter) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
