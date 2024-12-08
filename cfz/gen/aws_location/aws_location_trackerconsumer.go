// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_location

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Location_TrackerConsumer)(nil)
	_ cfz.Resource                   = (*AWS_Location_TrackerConsumer)(nil)
)

const (
	// AWS_Location_TrackerConsumer__Type is the CloudFormation type for AWS::Location::TrackerConsumer.
	AWS_Location_TrackerConsumer__Type = "AWS::Location::TrackerConsumer"
)

var (
	// AWS_Location_TrackerConsumer__PropertiesMap reports all the CloudFormation properties for AWS::Location::TrackerConsumer.
	AWS_Location_TrackerConsumer__PropertiesMap = struct {
		ConsumerArn string
		TrackerName string
	}{
		ConsumerArn: "ConsumerArn",
		TrackerName: "TrackerName",
	}

	// AWS_Location_TrackerConsumer__PropertiesSlice reports all the CloudFormation properties for AWS::Location::TrackerConsumer.
	AWS_Location_TrackerConsumer__PropertiesSlice = []string{
		AWS_Location_TrackerConsumer__PropertiesMap.ConsumerArn,
		AWS_Location_TrackerConsumer__PropertiesMap.TrackerName,
	}
)

// AWS_Location_TrackerConsumer is a binding for AWS::Location::TrackerConsumer.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html
type AWS_Location_TrackerConsumer struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ConsumerArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-consumerarn
	ConsumerArn cfz.Expression[string] `json:"ConsumerArn,omitempty"`

	// TrackerName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-trackername
	TrackerName cfz.Expression[string] `json:"TrackerName,omitempty"`
}

// New__AWS_Location_TrackerConsumer initializes a new *AWS_Location_TrackerConsumer.
func New__AWS_Location_TrackerConsumer(logicalName string) *AWS_Location_TrackerConsumer {
	return &AWS_Location_TrackerConsumer{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Location_TrackerConsumer) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Location_TrackerConsumer) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Location_TrackerConsumer) GetType() string {
	return AWS_Location_TrackerConsumer__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Location_TrackerConsumer) Set__LogicalName(v string) *AWS_Location_TrackerConsumer {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Location_TrackerConsumer) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Location_TrackerConsumer {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Location_TrackerConsumer) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Location_TrackerConsumer {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Location_TrackerConsumer) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Location_TrackerConsumer {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Location_TrackerConsumer) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Location_TrackerConsumer {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Location_TrackerConsumer) Set__RequestedOutputs(v []cfz.Output) *AWS_Location_TrackerConsumer {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Location_TrackerConsumer) Add__RequestedOutputs(v ...cfz.Output) *AWS_Location_TrackerConsumer {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ConsumerArn updates property "ConsumerArn".
func (t *AWS_Location_TrackerConsumer) Set__ConsumerArn(v cfz.Expression[string]) *AWS_Location_TrackerConsumer {
	t.ConsumerArn = v
	return t
}

// SetV__ConsumerArn updates property "ConsumerArn".
func (t *AWS_Location_TrackerConsumer) SetV__ConsumerArn(v string) *AWS_Location_TrackerConsumer {
	t.ConsumerArn = cfz.V(v)
	return t
}

// Set__TrackerName updates property "TrackerName".
func (t *AWS_Location_TrackerConsumer) Set__TrackerName(v cfz.Expression[string]) *AWS_Location_TrackerConsumer {
	t.TrackerName = v
	return t
}

// SetV__TrackerName updates property "TrackerName".
func (t *AWS_Location_TrackerConsumer) SetV__TrackerName(v string) *AWS_Location_TrackerConsumer {
	t.TrackerName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Location_TrackerConsumer) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Location_TrackerConsumer) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Location_TrackerConsumer) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Location_TrackerConsumer

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

func (t *AWS_Location_TrackerConsumer) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
