// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_batch

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Batch_JobQueue)(nil)
	_ cfz.Resource                   = (*AWS_Batch_JobQueue)(nil)
)

const (
	// AWS_Batch_JobQueue__Type is the CloudFormation type for AWS::Batch::JobQueue.
	AWS_Batch_JobQueue__Type = "AWS::Batch::JobQueue"
)

var (
	// AWS_Batch_JobQueue__AttributesMap reports all the CloudFormation attributes for AWS::Batch::JobQueue.
	AWS_Batch_JobQueue__AttributesMap = struct {
		JobQueueArn string
	}{
		JobQueueArn: "JobQueueArn",
	}

	// AWS_Batch_JobQueue__AttributesSlice reports all the CloudFormation attributes for AWS::Batch::JobQueue.
	AWS_Batch_JobQueue__AttributesSlice = []string{
		AWS_Batch_JobQueue__AttributesMap.JobQueueArn,
	}
)

var (
	// AWS_Batch_JobQueue__PropertiesMap reports all the CloudFormation properties for AWS::Batch::JobQueue.
	AWS_Batch_JobQueue__PropertiesMap = struct {
		ComputeEnvironmentOrder  string
		JobQueueName             string
		JobStateTimeLimitActions string
		Priority                 string
		SchedulingPolicyArn      string
		State                    string
		Tags                     string
	}{
		ComputeEnvironmentOrder:  "ComputeEnvironmentOrder",
		JobQueueName:             "JobQueueName",
		JobStateTimeLimitActions: "JobStateTimeLimitActions",
		Priority:                 "Priority",
		SchedulingPolicyArn:      "SchedulingPolicyArn",
		State:                    "State",
		Tags:                     "Tags",
	}

	// AWS_Batch_JobQueue__PropertiesSlice reports all the CloudFormation properties for AWS::Batch::JobQueue.
	AWS_Batch_JobQueue__PropertiesSlice = []string{
		AWS_Batch_JobQueue__PropertiesMap.ComputeEnvironmentOrder,
		AWS_Batch_JobQueue__PropertiesMap.JobQueueName,
		AWS_Batch_JobQueue__PropertiesMap.JobStateTimeLimitActions,
		AWS_Batch_JobQueue__PropertiesMap.Priority,
		AWS_Batch_JobQueue__PropertiesMap.SchedulingPolicyArn,
		AWS_Batch_JobQueue__PropertiesMap.State,
		AWS_Batch_JobQueue__PropertiesMap.Tags,
	}
)

// AWS_Batch_JobQueue is a binding for AWS::Batch::JobQueue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html
type AWS_Batch_JobQueue struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ComputeEnvironmentOrder is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-computeenvironmentorder
	ComputeEnvironmentOrder cfz.ExpressionSlice[AWS_Batch_JobQueue_ComputeEnvironmentOrder] `json:"ComputeEnvironmentOrder,omitempty"`

	// JobQueueName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-jobqueuename
	JobQueueName cfz.Expression[string] `json:"JobQueueName,omitempty"`

	// JobStateTimeLimitActions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-jobstatetimelimitactions
	JobStateTimeLimitActions cfz.ExpressionSlice[AWS_Batch_JobQueue_JobStateTimeLimitAction] `json:"JobStateTimeLimitActions,omitempty"`

	// Priority is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-priority
	Priority cfz.Expression[int32] `json:"Priority,omitempty"`

	// SchedulingPolicyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-schedulingpolicyarn
	SchedulingPolicyArn cfz.Expression[string] `json:"SchedulingPolicyArn,omitempty"`

	// State is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-state
	State cfz.Expression[string] `json:"State,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html#cfn-batch-jobqueue-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`
}

// New__AWS_Batch_JobQueue initializes a new *AWS_Batch_JobQueue.
func New__AWS_Batch_JobQueue(logicalName string) *AWS_Batch_JobQueue {
	return &AWS_Batch_JobQueue{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Batch_JobQueue) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Batch_JobQueue) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Batch_JobQueue) GetType() string {
	return AWS_Batch_JobQueue__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Batch_JobQueue) Set__LogicalName(v string) *AWS_Batch_JobQueue {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Batch_JobQueue) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Batch_JobQueue {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Batch_JobQueue) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Batch_JobQueue {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Batch_JobQueue) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Batch_JobQueue {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Batch_JobQueue) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Batch_JobQueue {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Batch_JobQueue) Set__RequestedOutputs(v []cfz.Output) *AWS_Batch_JobQueue {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Batch_JobQueue) Add__RequestedOutputs(v ...cfz.Output) *AWS_Batch_JobQueue {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ComputeEnvironmentOrder updates property "ComputeEnvironmentOrder".
func (t *AWS_Batch_JobQueue) Set__ComputeEnvironmentOrder(v cfz.ExpressionSlice[AWS_Batch_JobQueue_ComputeEnvironmentOrder]) *AWS_Batch_JobQueue {
	t.ComputeEnvironmentOrder = v
	return t
}

// SetS__ComputeEnvironmentOrder updates property "ComputeEnvironmentOrder".
func (t *AWS_Batch_JobQueue) SetS__ComputeEnvironmentOrder(v ...cfz.Expression[AWS_Batch_JobQueue_ComputeEnvironmentOrder]) *AWS_Batch_JobQueue {
	t.ComputeEnvironmentOrder = cfz.S(v...)
	return t
}

// SetSV__ComputeEnvironmentOrder updates property "ComputeEnvironmentOrder".
func (t *AWS_Batch_JobQueue) SetSV__ComputeEnvironmentOrder(v ...AWS_Batch_JobQueue_ComputeEnvironmentOrder) *AWS_Batch_JobQueue {
	t.ComputeEnvironmentOrder = cfz.SV(v...)
	return t
}

// Set__JobQueueName updates property "JobQueueName".
func (t *AWS_Batch_JobQueue) Set__JobQueueName(v cfz.Expression[string]) *AWS_Batch_JobQueue {
	t.JobQueueName = v
	return t
}

// SetV__JobQueueName updates property "JobQueueName".
func (t *AWS_Batch_JobQueue) SetV__JobQueueName(v string) *AWS_Batch_JobQueue {
	t.JobQueueName = cfz.V(v)
	return t
}

// Set__JobStateTimeLimitActions updates property "JobStateTimeLimitActions".
func (t *AWS_Batch_JobQueue) Set__JobStateTimeLimitActions(v cfz.ExpressionSlice[AWS_Batch_JobQueue_JobStateTimeLimitAction]) *AWS_Batch_JobQueue {
	t.JobStateTimeLimitActions = v
	return t
}

// SetS__JobStateTimeLimitActions updates property "JobStateTimeLimitActions".
func (t *AWS_Batch_JobQueue) SetS__JobStateTimeLimitActions(v ...cfz.Expression[AWS_Batch_JobQueue_JobStateTimeLimitAction]) *AWS_Batch_JobQueue {
	t.JobStateTimeLimitActions = cfz.S(v...)
	return t
}

// SetSV__JobStateTimeLimitActions updates property "JobStateTimeLimitActions".
func (t *AWS_Batch_JobQueue) SetSV__JobStateTimeLimitActions(v ...AWS_Batch_JobQueue_JobStateTimeLimitAction) *AWS_Batch_JobQueue {
	t.JobStateTimeLimitActions = cfz.SV(v...)
	return t
}

// Set__Priority updates property "Priority".
func (t *AWS_Batch_JobQueue) Set__Priority(v cfz.Expression[int32]) *AWS_Batch_JobQueue {
	t.Priority = v
	return t
}

// SetV__Priority updates property "Priority".
func (t *AWS_Batch_JobQueue) SetV__Priority(v int32) *AWS_Batch_JobQueue {
	t.Priority = cfz.V(v)
	return t
}

// Set__SchedulingPolicyArn updates property "SchedulingPolicyArn".
func (t *AWS_Batch_JobQueue) Set__SchedulingPolicyArn(v cfz.Expression[string]) *AWS_Batch_JobQueue {
	t.SchedulingPolicyArn = v
	return t
}

// SetV__SchedulingPolicyArn updates property "SchedulingPolicyArn".
func (t *AWS_Batch_JobQueue) SetV__SchedulingPolicyArn(v string) *AWS_Batch_JobQueue {
	t.SchedulingPolicyArn = cfz.V(v)
	return t
}

// Set__State updates property "State".
func (t *AWS_Batch_JobQueue) Set__State(v cfz.Expression[string]) *AWS_Batch_JobQueue {
	t.State = v
	return t
}

// SetV__State updates property "State".
func (t *AWS_Batch_JobQueue) SetV__State(v string) *AWS_Batch_JobQueue {
	t.State = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Batch_JobQueue) Set__Tags(v cfz.ExpressionMap[string]) *AWS_Batch_JobQueue {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_Batch_JobQueue) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_Batch_JobQueue {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_Batch_JobQueue) SetMV__Tags(v ...map[string]string) *AWS_Batch_JobQueue {
	t.Tags = cfz.MV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Batch_JobQueue) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__JobQueueArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: JobQueueArn
func (t *AWS_Batch_JobQueue) GetAtt__JobQueueArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Batch_JobQueue__AttributesMap.JobQueueArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Batch_JobQueue) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__JobQueueArn returns a conventionally configured output for an attribute of this resource.
// Attribute: JobQueueArn
func (t *AWS_Batch_JobQueue) GetConventionalOutputAtt__JobQueueArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttJobQueueArn", t.GetAtt__JobQueueArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Batch_JobQueue) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Batch_JobQueue

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

func (t *AWS_Batch_JobQueue) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
