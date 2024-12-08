// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_medialive

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MediaLive_Channel)(nil)
	_ cfz.Resource                   = (*AWS_MediaLive_Channel)(nil)
)

const (
	// AWS_MediaLive_Channel__Type is the CloudFormation type for AWS::MediaLive::Channel.
	AWS_MediaLive_Channel__Type = "AWS::MediaLive::Channel"
)

var (
	// AWS_MediaLive_Channel__AttributesMap reports all the CloudFormation attributes for AWS::MediaLive::Channel.
	AWS_MediaLive_Channel__AttributesMap = struct {
		Arn    string
		Inputs string
	}{
		Arn:    "Arn",
		Inputs: "Inputs",
	}

	// AWS_MediaLive_Channel__AttributesSlice reports all the CloudFormation attributes for AWS::MediaLive::Channel.
	AWS_MediaLive_Channel__AttributesSlice = []string{
		AWS_MediaLive_Channel__AttributesMap.Arn,
		AWS_MediaLive_Channel__AttributesMap.Inputs,
	}
)

var (
	// AWS_MediaLive_Channel__PropertiesMap reports all the CloudFormation properties for AWS::MediaLive::Channel.
	AWS_MediaLive_Channel__PropertiesMap = struct {
		AnywhereSettings      string
		CdiInputSpecification string
		ChannelClass          string
		Destinations          string
		EncoderSettings       string
		InputAttachments      string
		InputSpecification    string
		LogLevel              string
		Maintenance           string
		Name                  string
		RoleArn               string
		Tags                  string
		Vpc                   string
	}{
		AnywhereSettings:      "AnywhereSettings",
		CdiInputSpecification: "CdiInputSpecification",
		ChannelClass:          "ChannelClass",
		Destinations:          "Destinations",
		EncoderSettings:       "EncoderSettings",
		InputAttachments:      "InputAttachments",
		InputSpecification:    "InputSpecification",
		LogLevel:              "LogLevel",
		Maintenance:           "Maintenance",
		Name:                  "Name",
		RoleArn:               "RoleArn",
		Tags:                  "Tags",
		Vpc:                   "Vpc",
	}

	// AWS_MediaLive_Channel__PropertiesSlice reports all the CloudFormation properties for AWS::MediaLive::Channel.
	AWS_MediaLive_Channel__PropertiesSlice = []string{
		AWS_MediaLive_Channel__PropertiesMap.AnywhereSettings,
		AWS_MediaLive_Channel__PropertiesMap.CdiInputSpecification,
		AWS_MediaLive_Channel__PropertiesMap.ChannelClass,
		AWS_MediaLive_Channel__PropertiesMap.Destinations,
		AWS_MediaLive_Channel__PropertiesMap.EncoderSettings,
		AWS_MediaLive_Channel__PropertiesMap.InputAttachments,
		AWS_MediaLive_Channel__PropertiesMap.InputSpecification,
		AWS_MediaLive_Channel__PropertiesMap.LogLevel,
		AWS_MediaLive_Channel__PropertiesMap.Maintenance,
		AWS_MediaLive_Channel__PropertiesMap.Name,
		AWS_MediaLive_Channel__PropertiesMap.RoleArn,
		AWS_MediaLive_Channel__PropertiesMap.Tags,
		AWS_MediaLive_Channel__PropertiesMap.Vpc,
	}
)

// AWS_MediaLive_Channel is a binding for AWS::MediaLive::Channel.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html
type AWS_MediaLive_Channel struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AnywhereSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-anywheresettings
	AnywhereSettings cfz.Expression[AWS_MediaLive_Channel_AnywhereSettings] `json:"AnywhereSettings,omitempty"`

	// CdiInputSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-cdiinputspecification
	CdiInputSpecification cfz.Expression[AWS_MediaLive_Channel_CdiInputSpecification] `json:"CdiInputSpecification,omitempty"`

	// ChannelClass is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-channelclass
	ChannelClass cfz.Expression[string] `json:"ChannelClass,omitempty"`

	// Destinations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-destinations
	Destinations cfz.ExpressionSlice[AWS_MediaLive_Channel_OutputDestination] `json:"Destinations,omitempty"`

	// EncoderSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-encodersettings
	EncoderSettings cfz.Expression[AWS_MediaLive_Channel_EncoderSettings] `json:"EncoderSettings,omitempty"`

	// InputAttachments is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-inputattachments
	InputAttachments cfz.ExpressionSlice[AWS_MediaLive_Channel_InputAttachment] `json:"InputAttachments,omitempty"`

	// InputSpecification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-inputspecification
	InputSpecification cfz.Expression[AWS_MediaLive_Channel_InputSpecification] `json:"InputSpecification,omitempty"`

	// LogLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-loglevel
	LogLevel cfz.Expression[string] `json:"LogLevel,omitempty"`

	// Maintenance is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-maintenance
	Maintenance cfz.Expression[AWS_MediaLive_Channel_MaintenanceCreateSettings] `json:"Maintenance,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// RoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-rolearn
	RoleArn cfz.Expression[string] `json:"RoleArn,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-tags
	Tags cfz.Expression[json.RawMessage] `json:"Tags,omitempty"`

	// Vpc is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-vpc
	Vpc cfz.Expression[AWS_MediaLive_Channel_VpcOutputSettings] `json:"Vpc,omitempty"`
}

// New__AWS_MediaLive_Channel initializes a new *AWS_MediaLive_Channel.
func New__AWS_MediaLive_Channel(logicalName string) *AWS_MediaLive_Channel {
	return &AWS_MediaLive_Channel{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MediaLive_Channel) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MediaLive_Channel) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MediaLive_Channel) GetType() string {
	return AWS_MediaLive_Channel__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MediaLive_Channel) Set__LogicalName(v string) *AWS_MediaLive_Channel {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MediaLive_Channel) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MediaLive_Channel {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MediaLive_Channel) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MediaLive_Channel {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MediaLive_Channel) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MediaLive_Channel {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MediaLive_Channel) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MediaLive_Channel {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MediaLive_Channel) Set__RequestedOutputs(v []cfz.Output) *AWS_MediaLive_Channel {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MediaLive_Channel) Add__RequestedOutputs(v ...cfz.Output) *AWS_MediaLive_Channel {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AnywhereSettings updates property "AnywhereSettings".
func (t *AWS_MediaLive_Channel) Set__AnywhereSettings(v cfz.Expression[AWS_MediaLive_Channel_AnywhereSettings]) *AWS_MediaLive_Channel {
	t.AnywhereSettings = v
	return t
}

// SetV__AnywhereSettings updates property "AnywhereSettings".
func (t *AWS_MediaLive_Channel) SetV__AnywhereSettings(v AWS_MediaLive_Channel_AnywhereSettings) *AWS_MediaLive_Channel {
	t.AnywhereSettings = cfz.V(v)
	return t
}

// Set__CdiInputSpecification updates property "CdiInputSpecification".
func (t *AWS_MediaLive_Channel) Set__CdiInputSpecification(v cfz.Expression[AWS_MediaLive_Channel_CdiInputSpecification]) *AWS_MediaLive_Channel {
	t.CdiInputSpecification = v
	return t
}

// SetV__CdiInputSpecification updates property "CdiInputSpecification".
func (t *AWS_MediaLive_Channel) SetV__CdiInputSpecification(v AWS_MediaLive_Channel_CdiInputSpecification) *AWS_MediaLive_Channel {
	t.CdiInputSpecification = cfz.V(v)
	return t
}

// Set__ChannelClass updates property "ChannelClass".
func (t *AWS_MediaLive_Channel) Set__ChannelClass(v cfz.Expression[string]) *AWS_MediaLive_Channel {
	t.ChannelClass = v
	return t
}

// SetV__ChannelClass updates property "ChannelClass".
func (t *AWS_MediaLive_Channel) SetV__ChannelClass(v string) *AWS_MediaLive_Channel {
	t.ChannelClass = cfz.V(v)
	return t
}

// Set__Destinations updates property "Destinations".
func (t *AWS_MediaLive_Channel) Set__Destinations(v cfz.ExpressionSlice[AWS_MediaLive_Channel_OutputDestination]) *AWS_MediaLive_Channel {
	t.Destinations = v
	return t
}

// SetS__Destinations updates property "Destinations".
func (t *AWS_MediaLive_Channel) SetS__Destinations(v ...cfz.Expression[AWS_MediaLive_Channel_OutputDestination]) *AWS_MediaLive_Channel {
	t.Destinations = cfz.S(v...)
	return t
}

// SetSV__Destinations updates property "Destinations".
func (t *AWS_MediaLive_Channel) SetSV__Destinations(v ...AWS_MediaLive_Channel_OutputDestination) *AWS_MediaLive_Channel {
	t.Destinations = cfz.SV(v...)
	return t
}

// Set__EncoderSettings updates property "EncoderSettings".
func (t *AWS_MediaLive_Channel) Set__EncoderSettings(v cfz.Expression[AWS_MediaLive_Channel_EncoderSettings]) *AWS_MediaLive_Channel {
	t.EncoderSettings = v
	return t
}

// SetV__EncoderSettings updates property "EncoderSettings".
func (t *AWS_MediaLive_Channel) SetV__EncoderSettings(v AWS_MediaLive_Channel_EncoderSettings) *AWS_MediaLive_Channel {
	t.EncoderSettings = cfz.V(v)
	return t
}

// Set__InputAttachments updates property "InputAttachments".
func (t *AWS_MediaLive_Channel) Set__InputAttachments(v cfz.ExpressionSlice[AWS_MediaLive_Channel_InputAttachment]) *AWS_MediaLive_Channel {
	t.InputAttachments = v
	return t
}

// SetS__InputAttachments updates property "InputAttachments".
func (t *AWS_MediaLive_Channel) SetS__InputAttachments(v ...cfz.Expression[AWS_MediaLive_Channel_InputAttachment]) *AWS_MediaLive_Channel {
	t.InputAttachments = cfz.S(v...)
	return t
}

// SetSV__InputAttachments updates property "InputAttachments".
func (t *AWS_MediaLive_Channel) SetSV__InputAttachments(v ...AWS_MediaLive_Channel_InputAttachment) *AWS_MediaLive_Channel {
	t.InputAttachments = cfz.SV(v...)
	return t
}

// Set__InputSpecification updates property "InputSpecification".
func (t *AWS_MediaLive_Channel) Set__InputSpecification(v cfz.Expression[AWS_MediaLive_Channel_InputSpecification]) *AWS_MediaLive_Channel {
	t.InputSpecification = v
	return t
}

// SetV__InputSpecification updates property "InputSpecification".
func (t *AWS_MediaLive_Channel) SetV__InputSpecification(v AWS_MediaLive_Channel_InputSpecification) *AWS_MediaLive_Channel {
	t.InputSpecification = cfz.V(v)
	return t
}

// Set__LogLevel updates property "LogLevel".
func (t *AWS_MediaLive_Channel) Set__LogLevel(v cfz.Expression[string]) *AWS_MediaLive_Channel {
	t.LogLevel = v
	return t
}

// SetV__LogLevel updates property "LogLevel".
func (t *AWS_MediaLive_Channel) SetV__LogLevel(v string) *AWS_MediaLive_Channel {
	t.LogLevel = cfz.V(v)
	return t
}

// Set__Maintenance updates property "Maintenance".
func (t *AWS_MediaLive_Channel) Set__Maintenance(v cfz.Expression[AWS_MediaLive_Channel_MaintenanceCreateSettings]) *AWS_MediaLive_Channel {
	t.Maintenance = v
	return t
}

// SetV__Maintenance updates property "Maintenance".
func (t *AWS_MediaLive_Channel) SetV__Maintenance(v AWS_MediaLive_Channel_MaintenanceCreateSettings) *AWS_MediaLive_Channel {
	t.Maintenance = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_MediaLive_Channel) Set__Name(v cfz.Expression[string]) *AWS_MediaLive_Channel {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_MediaLive_Channel) SetV__Name(v string) *AWS_MediaLive_Channel {
	t.Name = cfz.V(v)
	return t
}

// Set__RoleArn updates property "RoleArn".
func (t *AWS_MediaLive_Channel) Set__RoleArn(v cfz.Expression[string]) *AWS_MediaLive_Channel {
	t.RoleArn = v
	return t
}

// SetV__RoleArn updates property "RoleArn".
func (t *AWS_MediaLive_Channel) SetV__RoleArn(v string) *AWS_MediaLive_Channel {
	t.RoleArn = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MediaLive_Channel) Set__Tags(v cfz.Expression[json.RawMessage]) *AWS_MediaLive_Channel {
	t.Tags = v
	return t
}

// SetV__Tags updates property "Tags".
func (t *AWS_MediaLive_Channel) SetV__Tags(v json.RawMessage) *AWS_MediaLive_Channel {
	t.Tags = cfz.V(v)
	return t
}

// Set__Vpc updates property "Vpc".
func (t *AWS_MediaLive_Channel) Set__Vpc(v cfz.Expression[AWS_MediaLive_Channel_VpcOutputSettings]) *AWS_MediaLive_Channel {
	t.Vpc = v
	return t
}

// SetV__Vpc updates property "Vpc".
func (t *AWS_MediaLive_Channel) SetV__Vpc(v AWS_MediaLive_Channel_VpcOutputSettings) *AWS_MediaLive_Channel {
	t.Vpc = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MediaLive_Channel) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_MediaLive_Channel) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaLive_Channel__AttributesMap.Arn))
}

// GetAttSlice__Inputs returns a $cfz.ExpressionSlice[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Inputs
func (t *AWS_MediaLive_Channel) GetAttSlice__Inputs() cfz.ExpressionSlice[string] {
	return cfz.GetAttSlice[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaLive_Channel__AttributesMap.Inputs))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MediaLive_Channel) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_MediaLive_Channel) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Inputs returns a conventionally configured output for an attribute of this resource.
// Attribute: Inputs
func (t *AWS_MediaLive_Channel) GetConventionalOutputAtt__Inputs(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttInputs", t.GetAttSlice__Inputs())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MediaLive_Channel) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MediaLive_Channel

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

func (t *AWS_MediaLive_Channel) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
