// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_budgets

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Budgets_Budget)(nil)
	_ cfz.Resource                   = (*AWS_Budgets_Budget)(nil)
)

const (
	// AWS_Budgets_Budget__Type is the CloudFormation type for AWS::Budgets::Budget.
	AWS_Budgets_Budget__Type = "AWS::Budgets::Budget"
)

var (
	// AWS_Budgets_Budget__PropertiesMap reports all the CloudFormation properties for AWS::Budgets::Budget.
	AWS_Budgets_Budget__PropertiesMap = struct {
		Budget                       string
		NotificationsWithSubscribers string
		ResourceTags                 string
	}{
		Budget:                       "Budget",
		NotificationsWithSubscribers: "NotificationsWithSubscribers",
		ResourceTags:                 "ResourceTags",
	}

	// AWS_Budgets_Budget__PropertiesSlice reports all the CloudFormation properties for AWS::Budgets::Budget.
	AWS_Budgets_Budget__PropertiesSlice = []string{
		AWS_Budgets_Budget__PropertiesMap.Budget,
		AWS_Budgets_Budget__PropertiesMap.NotificationsWithSubscribers,
		AWS_Budgets_Budget__PropertiesMap.ResourceTags,
	}
)

// AWS_Budgets_Budget is a binding for AWS::Budgets::Budget.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html
type AWS_Budgets_Budget struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Budget is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
	Budget cfz.Expression[AWS_Budgets_Budget_BudgetData] `json:"Budget,omitempty"`

	// NotificationsWithSubscribers is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
	NotificationsWithSubscribers cfz.ExpressionSlice[AWS_Budgets_Budget_NotificationWithSubscribers] `json:"NotificationsWithSubscribers,omitempty"`

	// ResourceTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-resourcetags
	ResourceTags cfz.ExpressionSlice[AWS_Budgets_Budget_ResourceTag] `json:"ResourceTags,omitempty"`
}

// New__AWS_Budgets_Budget initializes a new *AWS_Budgets_Budget.
func New__AWS_Budgets_Budget(logicalName string) *AWS_Budgets_Budget {
	return &AWS_Budgets_Budget{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Budgets_Budget) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Budgets_Budget) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Budgets_Budget) GetType() string {
	return AWS_Budgets_Budget__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Budgets_Budget) Set__LogicalName(v string) *AWS_Budgets_Budget {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Budgets_Budget) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Budgets_Budget {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Budgets_Budget) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Budgets_Budget {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Budgets_Budget) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Budgets_Budget {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Budgets_Budget) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Budgets_Budget {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Budgets_Budget) Set__RequestedOutputs(v []cfz.Output) *AWS_Budgets_Budget {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Budgets_Budget) Add__RequestedOutputs(v ...cfz.Output) *AWS_Budgets_Budget {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Budget updates property "Budget".
func (t *AWS_Budgets_Budget) Set__Budget(v cfz.Expression[AWS_Budgets_Budget_BudgetData]) *AWS_Budgets_Budget {
	t.Budget = v
	return t
}

// SetV__Budget updates property "Budget".
func (t *AWS_Budgets_Budget) SetV__Budget(v AWS_Budgets_Budget_BudgetData) *AWS_Budgets_Budget {
	t.Budget = cfz.V(v)
	return t
}

// Set__NotificationsWithSubscribers updates property "NotificationsWithSubscribers".
func (t *AWS_Budgets_Budget) Set__NotificationsWithSubscribers(v cfz.ExpressionSlice[AWS_Budgets_Budget_NotificationWithSubscribers]) *AWS_Budgets_Budget {
	t.NotificationsWithSubscribers = v
	return t
}

// SetS__NotificationsWithSubscribers updates property "NotificationsWithSubscribers".
func (t *AWS_Budgets_Budget) SetS__NotificationsWithSubscribers(v ...cfz.Expression[AWS_Budgets_Budget_NotificationWithSubscribers]) *AWS_Budgets_Budget {
	t.NotificationsWithSubscribers = cfz.S(v...)
	return t
}

// SetSV__NotificationsWithSubscribers updates property "NotificationsWithSubscribers".
func (t *AWS_Budgets_Budget) SetSV__NotificationsWithSubscribers(v ...AWS_Budgets_Budget_NotificationWithSubscribers) *AWS_Budgets_Budget {
	t.NotificationsWithSubscribers = cfz.SV(v...)
	return t
}

// Set__ResourceTags updates property "ResourceTags".
func (t *AWS_Budgets_Budget) Set__ResourceTags(v cfz.ExpressionSlice[AWS_Budgets_Budget_ResourceTag]) *AWS_Budgets_Budget {
	t.ResourceTags = v
	return t
}

// SetS__ResourceTags updates property "ResourceTags".
func (t *AWS_Budgets_Budget) SetS__ResourceTags(v ...cfz.Expression[AWS_Budgets_Budget_ResourceTag]) *AWS_Budgets_Budget {
	t.ResourceTags = cfz.S(v...)
	return t
}

// SetSV__ResourceTags updates property "ResourceTags".
func (t *AWS_Budgets_Budget) SetSV__ResourceTags(v ...AWS_Budgets_Budget_ResourceTag) *AWS_Budgets_Budget {
	t.ResourceTags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Budgets_Budget) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Budgets_Budget) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Budgets_Budget) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Budgets_Budget

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

func (t *AWS_Budgets_Budget) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
