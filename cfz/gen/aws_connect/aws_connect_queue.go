// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_connect

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Connect_Queue)(nil)
	_ cfz.Resource                   = (*AWS_Connect_Queue)(nil)
)

const (
	// AWS_Connect_Queue__Type is the CloudFormation type for AWS::Connect::Queue.
	AWS_Connect_Queue__Type = "AWS::Connect::Queue"
)

var (
	// AWS_Connect_Queue__AttributesMap reports all the CloudFormation attributes for AWS::Connect::Queue.
	AWS_Connect_Queue__AttributesMap = struct {
		QueueArn string
		Type     string
	}{
		QueueArn: "QueueArn",
		Type:     "Type",
	}

	// AWS_Connect_Queue__AttributesSlice reports all the CloudFormation attributes for AWS::Connect::Queue.
	AWS_Connect_Queue__AttributesSlice = []string{
		AWS_Connect_Queue__AttributesMap.QueueArn,
		AWS_Connect_Queue__AttributesMap.Type,
	}
)

var (
	// AWS_Connect_Queue__PropertiesMap reports all the CloudFormation properties for AWS::Connect::Queue.
	AWS_Connect_Queue__PropertiesMap = struct {
		Description          string
		HoursOfOperationArn  string
		InstanceArn          string
		MaxContacts          string
		Name                 string
		OutboundCallerConfig string
		QuickConnectArns     string
		Status               string
		Tags                 string
	}{
		Description:          "Description",
		HoursOfOperationArn:  "HoursOfOperationArn",
		InstanceArn:          "InstanceArn",
		MaxContacts:          "MaxContacts",
		Name:                 "Name",
		OutboundCallerConfig: "OutboundCallerConfig",
		QuickConnectArns:     "QuickConnectArns",
		Status:               "Status",
		Tags:                 "Tags",
	}

	// AWS_Connect_Queue__PropertiesSlice reports all the CloudFormation properties for AWS::Connect::Queue.
	AWS_Connect_Queue__PropertiesSlice = []string{
		AWS_Connect_Queue__PropertiesMap.Description,
		AWS_Connect_Queue__PropertiesMap.HoursOfOperationArn,
		AWS_Connect_Queue__PropertiesMap.InstanceArn,
		AWS_Connect_Queue__PropertiesMap.MaxContacts,
		AWS_Connect_Queue__PropertiesMap.Name,
		AWS_Connect_Queue__PropertiesMap.OutboundCallerConfig,
		AWS_Connect_Queue__PropertiesMap.QuickConnectArns,
		AWS_Connect_Queue__PropertiesMap.Status,
		AWS_Connect_Queue__PropertiesMap.Tags,
	}
)

// AWS_Connect_Queue is a binding for AWS::Connect::Queue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html
type AWS_Connect_Queue struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// HoursOfOperationArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-hoursofoperationarn
	HoursOfOperationArn cfz.Expression[string] `json:"HoursOfOperationArn,omitempty"`

	// InstanceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-instancearn
	InstanceArn cfz.Expression[string] `json:"InstanceArn,omitempty"`

	// MaxContacts is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-maxcontacts
	MaxContacts cfz.Expression[int32] `json:"MaxContacts,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// OutboundCallerConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-outboundcallerconfig
	OutboundCallerConfig cfz.Expression[AWS_Connect_Queue_OutboundCallerConfig] `json:"OutboundCallerConfig,omitempty"`

	// QuickConnectArns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-quickconnectarns
	QuickConnectArns cfz.ExpressionSlice[string] `json:"QuickConnectArns,omitempty"`

	// Status is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-status
	Status cfz.Expression[string] `json:"Status,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-queue.html#cfn-connect-queue-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_Connect_Queue initializes a new *AWS_Connect_Queue.
func New__AWS_Connect_Queue(logicalName string) *AWS_Connect_Queue {
	return &AWS_Connect_Queue{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Connect_Queue) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Connect_Queue) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Connect_Queue) GetType() string {
	return AWS_Connect_Queue__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Connect_Queue) Set__LogicalName(v string) *AWS_Connect_Queue {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Connect_Queue) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Connect_Queue {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Connect_Queue) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Connect_Queue {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Connect_Queue) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Connect_Queue {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Connect_Queue) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Connect_Queue {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Connect_Queue) Set__RequestedOutputs(v []cfz.Output) *AWS_Connect_Queue {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Connect_Queue) Add__RequestedOutputs(v ...cfz.Output) *AWS_Connect_Queue {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_Connect_Queue) Set__Description(v cfz.Expression[string]) *AWS_Connect_Queue {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_Connect_Queue) SetV__Description(v string) *AWS_Connect_Queue {
	t.Description = cfz.V(v)
	return t
}

// Set__HoursOfOperationArn updates property "HoursOfOperationArn".
func (t *AWS_Connect_Queue) Set__HoursOfOperationArn(v cfz.Expression[string]) *AWS_Connect_Queue {
	t.HoursOfOperationArn = v
	return t
}

// SetV__HoursOfOperationArn updates property "HoursOfOperationArn".
func (t *AWS_Connect_Queue) SetV__HoursOfOperationArn(v string) *AWS_Connect_Queue {
	t.HoursOfOperationArn = cfz.V(v)
	return t
}

// Set__InstanceArn updates property "InstanceArn".
func (t *AWS_Connect_Queue) Set__InstanceArn(v cfz.Expression[string]) *AWS_Connect_Queue {
	t.InstanceArn = v
	return t
}

// SetV__InstanceArn updates property "InstanceArn".
func (t *AWS_Connect_Queue) SetV__InstanceArn(v string) *AWS_Connect_Queue {
	t.InstanceArn = cfz.V(v)
	return t
}

// Set__MaxContacts updates property "MaxContacts".
func (t *AWS_Connect_Queue) Set__MaxContacts(v cfz.Expression[int32]) *AWS_Connect_Queue {
	t.MaxContacts = v
	return t
}

// SetV__MaxContacts updates property "MaxContacts".
func (t *AWS_Connect_Queue) SetV__MaxContacts(v int32) *AWS_Connect_Queue {
	t.MaxContacts = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Connect_Queue) Set__Name(v cfz.Expression[string]) *AWS_Connect_Queue {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Connect_Queue) SetV__Name(v string) *AWS_Connect_Queue {
	t.Name = cfz.V(v)
	return t
}

// Set__OutboundCallerConfig updates property "OutboundCallerConfig".
func (t *AWS_Connect_Queue) Set__OutboundCallerConfig(v cfz.Expression[AWS_Connect_Queue_OutboundCallerConfig]) *AWS_Connect_Queue {
	t.OutboundCallerConfig = v
	return t
}

// SetV__OutboundCallerConfig updates property "OutboundCallerConfig".
func (t *AWS_Connect_Queue) SetV__OutboundCallerConfig(v AWS_Connect_Queue_OutboundCallerConfig) *AWS_Connect_Queue {
	t.OutboundCallerConfig = cfz.V(v)
	return t
}

// Set__QuickConnectArns updates property "QuickConnectArns".
func (t *AWS_Connect_Queue) Set__QuickConnectArns(v cfz.ExpressionSlice[string]) *AWS_Connect_Queue {
	t.QuickConnectArns = v
	return t
}

// SetS__QuickConnectArns updates property "QuickConnectArns".
func (t *AWS_Connect_Queue) SetS__QuickConnectArns(v ...cfz.Expression[string]) *AWS_Connect_Queue {
	t.QuickConnectArns = cfz.S(v...)
	return t
}

// SetSV__QuickConnectArns updates property "QuickConnectArns".
func (t *AWS_Connect_Queue) SetSV__QuickConnectArns(v ...string) *AWS_Connect_Queue {
	t.QuickConnectArns = cfz.SV(v...)
	return t
}

// Set__Status updates property "Status".
func (t *AWS_Connect_Queue) Set__Status(v cfz.Expression[string]) *AWS_Connect_Queue {
	t.Status = v
	return t
}

// SetV__Status updates property "Status".
func (t *AWS_Connect_Queue) SetV__Status(v string) *AWS_Connect_Queue {
	t.Status = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Connect_Queue) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Connect_Queue {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Connect_Queue) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Connect_Queue {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Connect_Queue) SetSV__Tags(v ...cfz.Tag) *AWS_Connect_Queue {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Connect_Queue) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__QueueArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: QueueArn
func (t *AWS_Connect_Queue) GetAtt__QueueArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Connect_Queue__AttributesMap.QueueArn))
}

// GetAtt__Type returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Type
func (t *AWS_Connect_Queue) GetAtt__Type() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Connect_Queue__AttributesMap.Type))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Connect_Queue) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__QueueArn returns a conventionally configured output for an attribute of this resource.
// Attribute: QueueArn
func (t *AWS_Connect_Queue) GetConventionalOutputAtt__QueueArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttQueueArn", t.GetAtt__QueueArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Type returns a conventionally configured output for an attribute of this resource.
// Attribute: Type
func (t *AWS_Connect_Queue) GetConventionalOutputAtt__Type(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttType", t.GetAtt__Type())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Connect_Queue) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Connect_Queue

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

func (t *AWS_Connect_Queue) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
