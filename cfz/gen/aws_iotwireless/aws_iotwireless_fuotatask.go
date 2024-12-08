// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotwireless

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_IoTWireless_FuotaTask)(nil)
	_ cfz.Resource                   = (*AWS_IoTWireless_FuotaTask)(nil)
)

const (
	// AWS_IoTWireless_FuotaTask__Type is the CloudFormation type for AWS::IoTWireless::FuotaTask.
	AWS_IoTWireless_FuotaTask__Type = "AWS::IoTWireless::FuotaTask"
)

var (
	// AWS_IoTWireless_FuotaTask__AttributesMap reports all the CloudFormation attributes for AWS::IoTWireless::FuotaTask.
	AWS_IoTWireless_FuotaTask__AttributesMap = struct {
		Arn               string
		FuotaTaskStatus   string
		Id                string
		LoRaWAN_StartTime string
	}{
		Arn:               "Arn",
		FuotaTaskStatus:   "FuotaTaskStatus",
		Id:                "Id",
		LoRaWAN_StartTime: "LoRaWAN.StartTime",
	}

	// AWS_IoTWireless_FuotaTask__AttributesSlice reports all the CloudFormation attributes for AWS::IoTWireless::FuotaTask.
	AWS_IoTWireless_FuotaTask__AttributesSlice = []string{
		AWS_IoTWireless_FuotaTask__AttributesMap.Arn,
		AWS_IoTWireless_FuotaTask__AttributesMap.FuotaTaskStatus,
		AWS_IoTWireless_FuotaTask__AttributesMap.Id,
		AWS_IoTWireless_FuotaTask__AttributesMap.LoRaWAN_StartTime,
	}
)

var (
	// AWS_IoTWireless_FuotaTask__PropertiesMap reports all the CloudFormation properties for AWS::IoTWireless::FuotaTask.
	AWS_IoTWireless_FuotaTask__PropertiesMap = struct {
		AssociateMulticastGroup    string
		AssociateWirelessDevice    string
		Description                string
		DisassociateMulticastGroup string
		DisassociateWirelessDevice string
		FirmwareUpdateImage        string
		FirmwareUpdateRole         string
		LoRaWAN                    string
		Name                       string
		Tags                       string
	}{
		AssociateMulticastGroup:    "AssociateMulticastGroup",
		AssociateWirelessDevice:    "AssociateWirelessDevice",
		Description:                "Description",
		DisassociateMulticastGroup: "DisassociateMulticastGroup",
		DisassociateWirelessDevice: "DisassociateWirelessDevice",
		FirmwareUpdateImage:        "FirmwareUpdateImage",
		FirmwareUpdateRole:         "FirmwareUpdateRole",
		LoRaWAN:                    "LoRaWAN",
		Name:                       "Name",
		Tags:                       "Tags",
	}

	// AWS_IoTWireless_FuotaTask__PropertiesSlice reports all the CloudFormation properties for AWS::IoTWireless::FuotaTask.
	AWS_IoTWireless_FuotaTask__PropertiesSlice = []string{
		AWS_IoTWireless_FuotaTask__PropertiesMap.AssociateMulticastGroup,
		AWS_IoTWireless_FuotaTask__PropertiesMap.AssociateWirelessDevice,
		AWS_IoTWireless_FuotaTask__PropertiesMap.Description,
		AWS_IoTWireless_FuotaTask__PropertiesMap.DisassociateMulticastGroup,
		AWS_IoTWireless_FuotaTask__PropertiesMap.DisassociateWirelessDevice,
		AWS_IoTWireless_FuotaTask__PropertiesMap.FirmwareUpdateImage,
		AWS_IoTWireless_FuotaTask__PropertiesMap.FirmwareUpdateRole,
		AWS_IoTWireless_FuotaTask__PropertiesMap.LoRaWAN,
		AWS_IoTWireless_FuotaTask__PropertiesMap.Name,
		AWS_IoTWireless_FuotaTask__PropertiesMap.Tags,
	}
)

// AWS_IoTWireless_FuotaTask is a binding for AWS::IoTWireless::FuotaTask.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html
type AWS_IoTWireless_FuotaTask struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AssociateMulticastGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-associatemulticastgroup
	AssociateMulticastGroup cfz.Expression[string] `json:"AssociateMulticastGroup,omitempty"`

	// AssociateWirelessDevice is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-associatewirelessdevice
	AssociateWirelessDevice cfz.Expression[string] `json:"AssociateWirelessDevice,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DisassociateMulticastGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-disassociatemulticastgroup
	DisassociateMulticastGroup cfz.Expression[string] `json:"DisassociateMulticastGroup,omitempty"`

	// DisassociateWirelessDevice is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-disassociatewirelessdevice
	DisassociateWirelessDevice cfz.Expression[string] `json:"DisassociateWirelessDevice,omitempty"`

	// FirmwareUpdateImage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-firmwareupdateimage
	FirmwareUpdateImage cfz.Expression[string] `json:"FirmwareUpdateImage,omitempty"`

	// FirmwareUpdateRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-firmwareupdaterole
	FirmwareUpdateRole cfz.Expression[string] `json:"FirmwareUpdateRole,omitempty"`

	// LoRaWAN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-lorawan
	LoRaWAN cfz.Expression[AWS_IoTWireless_FuotaTask_LoRaWAN] `json:"LoRaWAN,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_IoTWireless_FuotaTask initializes a new *AWS_IoTWireless_FuotaTask.
func New__AWS_IoTWireless_FuotaTask(logicalName string) *AWS_IoTWireless_FuotaTask {
	return &AWS_IoTWireless_FuotaTask{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_IoTWireless_FuotaTask) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_IoTWireless_FuotaTask) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_IoTWireless_FuotaTask) GetType() string {
	return AWS_IoTWireless_FuotaTask__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_IoTWireless_FuotaTask) Set__LogicalName(v string) *AWS_IoTWireless_FuotaTask {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_IoTWireless_FuotaTask) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_IoTWireless_FuotaTask {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_IoTWireless_FuotaTask) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_IoTWireless_FuotaTask {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_IoTWireless_FuotaTask) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_IoTWireless_FuotaTask {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_IoTWireless_FuotaTask) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_IoTWireless_FuotaTask {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_IoTWireless_FuotaTask) Set__RequestedOutputs(v []cfz.Output) *AWS_IoTWireless_FuotaTask {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_IoTWireless_FuotaTask) Add__RequestedOutputs(v ...cfz.Output) *AWS_IoTWireless_FuotaTask {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AssociateMulticastGroup updates property "AssociateMulticastGroup".
func (t *AWS_IoTWireless_FuotaTask) Set__AssociateMulticastGroup(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.AssociateMulticastGroup = v
	return t
}

// SetV__AssociateMulticastGroup updates property "AssociateMulticastGroup".
func (t *AWS_IoTWireless_FuotaTask) SetV__AssociateMulticastGroup(v string) *AWS_IoTWireless_FuotaTask {
	t.AssociateMulticastGroup = cfz.V(v)
	return t
}

// Set__AssociateWirelessDevice updates property "AssociateWirelessDevice".
func (t *AWS_IoTWireless_FuotaTask) Set__AssociateWirelessDevice(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.AssociateWirelessDevice = v
	return t
}

// SetV__AssociateWirelessDevice updates property "AssociateWirelessDevice".
func (t *AWS_IoTWireless_FuotaTask) SetV__AssociateWirelessDevice(v string) *AWS_IoTWireless_FuotaTask {
	t.AssociateWirelessDevice = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_IoTWireless_FuotaTask) Set__Description(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_IoTWireless_FuotaTask) SetV__Description(v string) *AWS_IoTWireless_FuotaTask {
	t.Description = cfz.V(v)
	return t
}

// Set__DisassociateMulticastGroup updates property "DisassociateMulticastGroup".
func (t *AWS_IoTWireless_FuotaTask) Set__DisassociateMulticastGroup(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.DisassociateMulticastGroup = v
	return t
}

// SetV__DisassociateMulticastGroup updates property "DisassociateMulticastGroup".
func (t *AWS_IoTWireless_FuotaTask) SetV__DisassociateMulticastGroup(v string) *AWS_IoTWireless_FuotaTask {
	t.DisassociateMulticastGroup = cfz.V(v)
	return t
}

// Set__DisassociateWirelessDevice updates property "DisassociateWirelessDevice".
func (t *AWS_IoTWireless_FuotaTask) Set__DisassociateWirelessDevice(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.DisassociateWirelessDevice = v
	return t
}

// SetV__DisassociateWirelessDevice updates property "DisassociateWirelessDevice".
func (t *AWS_IoTWireless_FuotaTask) SetV__DisassociateWirelessDevice(v string) *AWS_IoTWireless_FuotaTask {
	t.DisassociateWirelessDevice = cfz.V(v)
	return t
}

// Set__FirmwareUpdateImage updates property "FirmwareUpdateImage".
func (t *AWS_IoTWireless_FuotaTask) Set__FirmwareUpdateImage(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.FirmwareUpdateImage = v
	return t
}

// SetV__FirmwareUpdateImage updates property "FirmwareUpdateImage".
func (t *AWS_IoTWireless_FuotaTask) SetV__FirmwareUpdateImage(v string) *AWS_IoTWireless_FuotaTask {
	t.FirmwareUpdateImage = cfz.V(v)
	return t
}

// Set__FirmwareUpdateRole updates property "FirmwareUpdateRole".
func (t *AWS_IoTWireless_FuotaTask) Set__FirmwareUpdateRole(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.FirmwareUpdateRole = v
	return t
}

// SetV__FirmwareUpdateRole updates property "FirmwareUpdateRole".
func (t *AWS_IoTWireless_FuotaTask) SetV__FirmwareUpdateRole(v string) *AWS_IoTWireless_FuotaTask {
	t.FirmwareUpdateRole = cfz.V(v)
	return t
}

// Set__LoRaWAN updates property "LoRaWAN".
func (t *AWS_IoTWireless_FuotaTask) Set__LoRaWAN(v cfz.Expression[AWS_IoTWireless_FuotaTask_LoRaWAN]) *AWS_IoTWireless_FuotaTask {
	t.LoRaWAN = v
	return t
}

// SetV__LoRaWAN updates property "LoRaWAN".
func (t *AWS_IoTWireless_FuotaTask) SetV__LoRaWAN(v AWS_IoTWireless_FuotaTask_LoRaWAN) *AWS_IoTWireless_FuotaTask {
	t.LoRaWAN = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_IoTWireless_FuotaTask) Set__Name(v cfz.Expression[string]) *AWS_IoTWireless_FuotaTask {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_IoTWireless_FuotaTask) SetV__Name(v string) *AWS_IoTWireless_FuotaTask {
	t.Name = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_IoTWireless_FuotaTask) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_IoTWireless_FuotaTask {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_IoTWireless_FuotaTask) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_IoTWireless_FuotaTask {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_IoTWireless_FuotaTask) SetSV__Tags(v ...cfz.Tag) *AWS_IoTWireless_FuotaTask {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_IoTWireless_FuotaTask) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_IoTWireless_FuotaTask) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTWireless_FuotaTask__AttributesMap.Arn))
}

// GetAtt__FuotaTaskStatus returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FuotaTaskStatus
func (t *AWS_IoTWireless_FuotaTask) GetAtt__FuotaTaskStatus() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTWireless_FuotaTask__AttributesMap.FuotaTaskStatus))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_IoTWireless_FuotaTask) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTWireless_FuotaTask__AttributesMap.Id))
}

// GetAtt__LoRaWAN_StartTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: LoRaWAN.StartTime
func (t *AWS_IoTWireless_FuotaTask) GetAtt__LoRaWAN_StartTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_IoTWireless_FuotaTask__AttributesMap.LoRaWAN_StartTime))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_IoTWireless_FuotaTask) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_IoTWireless_FuotaTask) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FuotaTaskStatus returns a conventionally configured output for an attribute of this resource.
// Attribute: FuotaTaskStatus
func (t *AWS_IoTWireless_FuotaTask) GetConventionalOutputAtt__FuotaTaskStatus(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFuotaTaskStatus", t.GetAtt__FuotaTaskStatus())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_IoTWireless_FuotaTask) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__LoRaWAN_StartTime returns a conventionally configured output for an attribute of this resource.
// Attribute: LoRaWAN.StartTime
func (t *AWS_IoTWireless_FuotaTask) GetConventionalOutputAtt__LoRaWAN_StartTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttLoRaWANStartTime", t.GetAtt__LoRaWAN_StartTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_IoTWireless_FuotaTask) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_IoTWireless_FuotaTask

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

func (t *AWS_IoTWireless_FuotaTask) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
