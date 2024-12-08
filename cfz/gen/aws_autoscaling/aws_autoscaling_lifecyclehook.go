// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_autoscaling

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AutoScaling_LifecycleHook)(nil)
	_ cfz.Resource                   = (*AWS_AutoScaling_LifecycleHook)(nil)
)

const (
	// AWS_AutoScaling_LifecycleHook__Type is the CloudFormation type for AWS::AutoScaling::LifecycleHook.
	AWS_AutoScaling_LifecycleHook__Type = "AWS::AutoScaling::LifecycleHook"
)

var (
	// AWS_AutoScaling_LifecycleHook__PropertiesMap reports all the CloudFormation properties for AWS::AutoScaling::LifecycleHook.
	AWS_AutoScaling_LifecycleHook__PropertiesMap = struct {
		AutoScalingGroupName  string
		DefaultResult         string
		HeartbeatTimeout      string
		LifecycleHookName     string
		LifecycleTransition   string
		NotificationMetadata  string
		NotificationTargetARN string
		RoleARN               string
	}{
		AutoScalingGroupName:  "AutoScalingGroupName",
		DefaultResult:         "DefaultResult",
		HeartbeatTimeout:      "HeartbeatTimeout",
		LifecycleHookName:     "LifecycleHookName",
		LifecycleTransition:   "LifecycleTransition",
		NotificationMetadata:  "NotificationMetadata",
		NotificationTargetARN: "NotificationTargetARN",
		RoleARN:               "RoleARN",
	}

	// AWS_AutoScaling_LifecycleHook__PropertiesSlice reports all the CloudFormation properties for AWS::AutoScaling::LifecycleHook.
	AWS_AutoScaling_LifecycleHook__PropertiesSlice = []string{
		AWS_AutoScaling_LifecycleHook__PropertiesMap.AutoScalingGroupName,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.DefaultResult,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.HeartbeatTimeout,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.LifecycleHookName,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.LifecycleTransition,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.NotificationMetadata,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.NotificationTargetARN,
		AWS_AutoScaling_LifecycleHook__PropertiesMap.RoleARN,
	}
)

// AWS_AutoScaling_LifecycleHook is a binding for AWS::AutoScaling::LifecycleHook.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html
type AWS_AutoScaling_LifecycleHook struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AutoScalingGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-autoscalinggroupname
	AutoScalingGroupName cfz.Expression[string] `json:"AutoScalingGroupName,omitempty"`

	// DefaultResult is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-defaultresult
	DefaultResult cfz.Expression[string] `json:"DefaultResult,omitempty"`

	// HeartbeatTimeout is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-heartbeattimeout
	HeartbeatTimeout cfz.Expression[int32] `json:"HeartbeatTimeout,omitempty"`

	// LifecycleHookName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-lifecyclehookname
	LifecycleHookName cfz.Expression[string] `json:"LifecycleHookName,omitempty"`

	// LifecycleTransition is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-lifecycletransition
	LifecycleTransition cfz.Expression[string] `json:"LifecycleTransition,omitempty"`

	// NotificationMetadata is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-notificationmetadata
	NotificationMetadata cfz.Expression[string] `json:"NotificationMetadata,omitempty"`

	// NotificationTargetARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-notificationtargetarn
	NotificationTargetARN cfz.Expression[string] `json:"NotificationTargetARN,omitempty"`

	// RoleARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html#cfn-autoscaling-lifecyclehook-rolearn
	RoleARN cfz.Expression[string] `json:"RoleARN,omitempty"`
}

// New__AWS_AutoScaling_LifecycleHook initializes a new *AWS_AutoScaling_LifecycleHook.
func New__AWS_AutoScaling_LifecycleHook(logicalName string) *AWS_AutoScaling_LifecycleHook {
	return &AWS_AutoScaling_LifecycleHook{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AutoScaling_LifecycleHook) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AutoScaling_LifecycleHook) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AutoScaling_LifecycleHook) GetType() string {
	return AWS_AutoScaling_LifecycleHook__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AutoScaling_LifecycleHook) Set__LogicalName(v string) *AWS_AutoScaling_LifecycleHook {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AutoScaling_LifecycleHook) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AutoScaling_LifecycleHook {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AutoScaling_LifecycleHook) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AutoScaling_LifecycleHook {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AutoScaling_LifecycleHook) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AutoScaling_LifecycleHook {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AutoScaling_LifecycleHook) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AutoScaling_LifecycleHook {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AutoScaling_LifecycleHook) Set__RequestedOutputs(v []cfz.Output) *AWS_AutoScaling_LifecycleHook {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AutoScaling_LifecycleHook) Add__RequestedOutputs(v ...cfz.Output) *AWS_AutoScaling_LifecycleHook {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AutoScalingGroupName updates property "AutoScalingGroupName".
func (t *AWS_AutoScaling_LifecycleHook) Set__AutoScalingGroupName(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.AutoScalingGroupName = v
	return t
}

// SetV__AutoScalingGroupName updates property "AutoScalingGroupName".
func (t *AWS_AutoScaling_LifecycleHook) SetV__AutoScalingGroupName(v string) *AWS_AutoScaling_LifecycleHook {
	t.AutoScalingGroupName = cfz.V(v)
	return t
}

// Set__DefaultResult updates property "DefaultResult".
func (t *AWS_AutoScaling_LifecycleHook) Set__DefaultResult(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.DefaultResult = v
	return t
}

// SetV__DefaultResult updates property "DefaultResult".
func (t *AWS_AutoScaling_LifecycleHook) SetV__DefaultResult(v string) *AWS_AutoScaling_LifecycleHook {
	t.DefaultResult = cfz.V(v)
	return t
}

// Set__HeartbeatTimeout updates property "HeartbeatTimeout".
func (t *AWS_AutoScaling_LifecycleHook) Set__HeartbeatTimeout(v cfz.Expression[int32]) *AWS_AutoScaling_LifecycleHook {
	t.HeartbeatTimeout = v
	return t
}

// SetV__HeartbeatTimeout updates property "HeartbeatTimeout".
func (t *AWS_AutoScaling_LifecycleHook) SetV__HeartbeatTimeout(v int32) *AWS_AutoScaling_LifecycleHook {
	t.HeartbeatTimeout = cfz.V(v)
	return t
}

// Set__LifecycleHookName updates property "LifecycleHookName".
func (t *AWS_AutoScaling_LifecycleHook) Set__LifecycleHookName(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.LifecycleHookName = v
	return t
}

// SetV__LifecycleHookName updates property "LifecycleHookName".
func (t *AWS_AutoScaling_LifecycleHook) SetV__LifecycleHookName(v string) *AWS_AutoScaling_LifecycleHook {
	t.LifecycleHookName = cfz.V(v)
	return t
}

// Set__LifecycleTransition updates property "LifecycleTransition".
func (t *AWS_AutoScaling_LifecycleHook) Set__LifecycleTransition(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.LifecycleTransition = v
	return t
}

// SetV__LifecycleTransition updates property "LifecycleTransition".
func (t *AWS_AutoScaling_LifecycleHook) SetV__LifecycleTransition(v string) *AWS_AutoScaling_LifecycleHook {
	t.LifecycleTransition = cfz.V(v)
	return t
}

// Set__NotificationMetadata updates property "NotificationMetadata".
func (t *AWS_AutoScaling_LifecycleHook) Set__NotificationMetadata(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.NotificationMetadata = v
	return t
}

// SetV__NotificationMetadata updates property "NotificationMetadata".
func (t *AWS_AutoScaling_LifecycleHook) SetV__NotificationMetadata(v string) *AWS_AutoScaling_LifecycleHook {
	t.NotificationMetadata = cfz.V(v)
	return t
}

// Set__NotificationTargetARN updates property "NotificationTargetARN".
func (t *AWS_AutoScaling_LifecycleHook) Set__NotificationTargetARN(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.NotificationTargetARN = v
	return t
}

// SetV__NotificationTargetARN updates property "NotificationTargetARN".
func (t *AWS_AutoScaling_LifecycleHook) SetV__NotificationTargetARN(v string) *AWS_AutoScaling_LifecycleHook {
	t.NotificationTargetARN = cfz.V(v)
	return t
}

// Set__RoleARN updates property "RoleARN".
func (t *AWS_AutoScaling_LifecycleHook) Set__RoleARN(v cfz.Expression[string]) *AWS_AutoScaling_LifecycleHook {
	t.RoleARN = v
	return t
}

// SetV__RoleARN updates property "RoleARN".
func (t *AWS_AutoScaling_LifecycleHook) SetV__RoleARN(v string) *AWS_AutoScaling_LifecycleHook {
	t.RoleARN = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AutoScaling_LifecycleHook) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AutoScaling_LifecycleHook) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AutoScaling_LifecycleHook) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AutoScaling_LifecycleHook

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

func (t *AWS_AutoScaling_LifecycleHook) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
