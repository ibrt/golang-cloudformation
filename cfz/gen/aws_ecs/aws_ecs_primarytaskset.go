// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ecs

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_ECS_PrimaryTaskSet)(nil)
	_ cfz.Resource                   = (*AWS_ECS_PrimaryTaskSet)(nil)
)

const (
	// AWS_ECS_PrimaryTaskSet__Type is the CloudFormation type for AWS::ECS::PrimaryTaskSet.
	AWS_ECS_PrimaryTaskSet__Type = "AWS::ECS::PrimaryTaskSet"
)

var (
	// AWS_ECS_PrimaryTaskSet__PropertiesMap reports all the CloudFormation properties for AWS::ECS::PrimaryTaskSet.
	AWS_ECS_PrimaryTaskSet__PropertiesMap = struct {
		Cluster   string
		Service   string
		TaskSetId string
	}{
		Cluster:   "Cluster",
		Service:   "Service",
		TaskSetId: "TaskSetId",
	}

	// AWS_ECS_PrimaryTaskSet__PropertiesSlice reports all the CloudFormation properties for AWS::ECS::PrimaryTaskSet.
	AWS_ECS_PrimaryTaskSet__PropertiesSlice = []string{
		AWS_ECS_PrimaryTaskSet__PropertiesMap.Cluster,
		AWS_ECS_PrimaryTaskSet__PropertiesMap.Service,
		AWS_ECS_PrimaryTaskSet__PropertiesMap.TaskSetId,
	}
)

// AWS_ECS_PrimaryTaskSet is a binding for AWS::ECS::PrimaryTaskSet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-primarytaskset.html
type AWS_ECS_PrimaryTaskSet struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Cluster is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-primarytaskset.html#cfn-ecs-primarytaskset-cluster
	Cluster cfz.Expression[string] `json:"Cluster,omitempty"`

	// Service is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-primarytaskset.html#cfn-ecs-primarytaskset-service
	Service cfz.Expression[string] `json:"Service,omitempty"`

	// TaskSetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-primarytaskset.html#cfn-ecs-primarytaskset-tasksetid
	TaskSetId cfz.Expression[string] `json:"TaskSetId,omitempty"`
}

// New__AWS_ECS_PrimaryTaskSet initializes a new *AWS_ECS_PrimaryTaskSet.
func New__AWS_ECS_PrimaryTaskSet(logicalName string) *AWS_ECS_PrimaryTaskSet {
	return &AWS_ECS_PrimaryTaskSet{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_ECS_PrimaryTaskSet) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_ECS_PrimaryTaskSet) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_ECS_PrimaryTaskSet) GetType() string {
	return AWS_ECS_PrimaryTaskSet__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_ECS_PrimaryTaskSet) Set__LogicalName(v string) *AWS_ECS_PrimaryTaskSet {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_ECS_PrimaryTaskSet) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_ECS_PrimaryTaskSet {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_ECS_PrimaryTaskSet) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_ECS_PrimaryTaskSet {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_ECS_PrimaryTaskSet) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_ECS_PrimaryTaskSet {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_ECS_PrimaryTaskSet) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_ECS_PrimaryTaskSet {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_ECS_PrimaryTaskSet) Set__RequestedOutputs(v []cfz.Output) *AWS_ECS_PrimaryTaskSet {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_ECS_PrimaryTaskSet) Add__RequestedOutputs(v ...cfz.Output) *AWS_ECS_PrimaryTaskSet {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Cluster updates property "Cluster".
func (t *AWS_ECS_PrimaryTaskSet) Set__Cluster(v cfz.Expression[string]) *AWS_ECS_PrimaryTaskSet {
	t.Cluster = v
	return t
}

// SetV__Cluster updates property "Cluster".
func (t *AWS_ECS_PrimaryTaskSet) SetV__Cluster(v string) *AWS_ECS_PrimaryTaskSet {
	t.Cluster = cfz.V(v)
	return t
}

// Set__Service updates property "Service".
func (t *AWS_ECS_PrimaryTaskSet) Set__Service(v cfz.Expression[string]) *AWS_ECS_PrimaryTaskSet {
	t.Service = v
	return t
}

// SetV__Service updates property "Service".
func (t *AWS_ECS_PrimaryTaskSet) SetV__Service(v string) *AWS_ECS_PrimaryTaskSet {
	t.Service = cfz.V(v)
	return t
}

// Set__TaskSetId updates property "TaskSetId".
func (t *AWS_ECS_PrimaryTaskSet) Set__TaskSetId(v cfz.Expression[string]) *AWS_ECS_PrimaryTaskSet {
	t.TaskSetId = v
	return t
}

// SetV__TaskSetId updates property "TaskSetId".
func (t *AWS_ECS_PrimaryTaskSet) SetV__TaskSetId(v string) *AWS_ECS_PrimaryTaskSet {
	t.TaskSetId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_ECS_PrimaryTaskSet) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_ECS_PrimaryTaskSet) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_ECS_PrimaryTaskSet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_ECS_PrimaryTaskSet

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

func (t *AWS_ECS_PrimaryTaskSet) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
