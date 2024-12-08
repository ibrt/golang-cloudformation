// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotanalytics

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IoTAnalytics_Channel)(nil)
	_ cfz.Resource                   = (*AWS_IoTAnalytics_Channel)(nil)
)

const (
	// AWS_IoTAnalytics_Channel__Type is the CloudFormation type for AWS::IoTAnalytics::Channel.
	AWS_IoTAnalytics_Channel__Type = "AWS::IoTAnalytics::Channel"
)

var (
	// AWS_IoTAnalytics_Channel__AttributesMap reports all the CloudFormation attributes for AWS::IoTAnalytics::Channel.
	AWS_IoTAnalytics_Channel__AttributesMap = struct {
		Id string
	}{
		Id: "Id",
	}

	// AWS_IoTAnalytics_Channel__AttributesSlice reports all the CloudFormation attributes for AWS::IoTAnalytics::Channel.
	AWS_IoTAnalytics_Channel__AttributesSlice = []string{
		AWS_IoTAnalytics_Channel__AttributesMap.Id,
	}
)

var (
	// AWS_IoTAnalytics_Channel__PropertiesMap reports all the CloudFormation properties for AWS::IoTAnalytics::Channel.
	AWS_IoTAnalytics_Channel__PropertiesMap = struct {
		ChannelName     string
		ChannelStorage  string
		RetentionPeriod string
		Tags            string
	}{
		ChannelName:     "ChannelName",
		ChannelStorage:  "ChannelStorage",
		RetentionPeriod: "RetentionPeriod",
		Tags:            "Tags",
	}

	// AWS_IoTAnalytics_Channel__PropertiesSlice reports all the CloudFormation properties for AWS::IoTAnalytics::Channel.
	AWS_IoTAnalytics_Channel__PropertiesSlice = []string{
		AWS_IoTAnalytics_Channel__PropertiesMap.ChannelName,
		AWS_IoTAnalytics_Channel__PropertiesMap.ChannelStorage,
		AWS_IoTAnalytics_Channel__PropertiesMap.RetentionPeriod,
		AWS_IoTAnalytics_Channel__PropertiesMap.Tags,
	}
)

// AWS_IoTAnalytics_Channel is a binding for AWS::IoTAnalytics::Channel.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html
type AWS_IoTAnalytics_Channel struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ChannelName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html#cfn-iotanalytics-channel-channelname
	ChannelName cfz.Expression[string] `json:"ChannelName,omitempty"`

	// ChannelStorage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html#cfn-iotanalytics-channel-channelstorage
	ChannelStorage cfz.Expression[AWS_IoTAnalytics_Channel_ChannelStorage] `json:"ChannelStorage,omitempty"`

	// RetentionPeriod is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html#cfn-iotanalytics-channel-retentionperiod
	RetentionPeriod cfz.Expression[AWS_IoTAnalytics_Channel_RetentionPeriod] `json:"RetentionPeriod,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html#cfn-iotanalytics-channel-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_IoTAnalytics_Channel initializes a new *AWS_IoTAnalytics_Channel.
func New__AWS_IoTAnalytics_Channel(logicalName string) *AWS_IoTAnalytics_Channel {
	return &AWS_IoTAnalytics_Channel{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IoTAnalytics_Channel) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IoTAnalytics_Channel) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IoTAnalytics_Channel) GetType() string {
	return AWS_IoTAnalytics_Channel__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IoTAnalytics_Channel) Set__LogicalName(v string) *AWS_IoTAnalytics_Channel {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IoTAnalytics_Channel) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IoTAnalytics_Channel {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IoTAnalytics_Channel) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IoTAnalytics_Channel {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IoTAnalytics_Channel) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IoTAnalytics_Channel {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IoTAnalytics_Channel) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IoTAnalytics_Channel {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IoTAnalytics_Channel) Set__RequestedOutputs(v []cfz.Output) *AWS_IoTAnalytics_Channel {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IoTAnalytics_Channel) Add__RequestedOutputs(v ...cfz.Output) *AWS_IoTAnalytics_Channel {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ChannelName updates property "ChannelName".
func (t *AWS_IoTAnalytics_Channel) Set__ChannelName(v cfz.Expression[string]) *AWS_IoTAnalytics_Channel {
	t.ChannelName = v
	return t
}

// SetV__ChannelName updates property "ChannelName".
func (t *AWS_IoTAnalytics_Channel) SetV__ChannelName(v string) *AWS_IoTAnalytics_Channel {
	t.ChannelName = cfz.V(v)
	return t
}

// Set__ChannelStorage updates property "ChannelStorage".
func (t *AWS_IoTAnalytics_Channel) Set__ChannelStorage(v cfz.Expression[AWS_IoTAnalytics_Channel_ChannelStorage]) *AWS_IoTAnalytics_Channel {
	t.ChannelStorage = v
	return t
}

// SetV__ChannelStorage updates property "ChannelStorage".
func (t *AWS_IoTAnalytics_Channel) SetV__ChannelStorage(v AWS_IoTAnalytics_Channel_ChannelStorage) *AWS_IoTAnalytics_Channel {
	t.ChannelStorage = cfz.V(v)
	return t
}

// Set__RetentionPeriod updates property "RetentionPeriod".
func (t *AWS_IoTAnalytics_Channel) Set__RetentionPeriod(v cfz.Expression[AWS_IoTAnalytics_Channel_RetentionPeriod]) *AWS_IoTAnalytics_Channel {
	t.RetentionPeriod = v
	return t
}

// SetV__RetentionPeriod updates property "RetentionPeriod".
func (t *AWS_IoTAnalytics_Channel) SetV__RetentionPeriod(v AWS_IoTAnalytics_Channel_RetentionPeriod) *AWS_IoTAnalytics_Channel {
	t.RetentionPeriod = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IoTAnalytics_Channel) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IoTAnalytics_Channel {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IoTAnalytics_Channel) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IoTAnalytics_Channel {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IoTAnalytics_Channel) SetSV__Tags(v ...cfz.Tag) *AWS_IoTAnalytics_Channel {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IoTAnalytics_Channel) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_IoTAnalytics_Channel) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTAnalytics_Channel__AttributesMap.Id))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IoTAnalytics_Channel) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_IoTAnalytics_Channel) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IoTAnalytics_Channel) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IoTAnalytics_Channel

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

func (t *AWS_IoTAnalytics_Channel) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
