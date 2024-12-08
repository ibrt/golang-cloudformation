// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_workspacesthinclient

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_WorkSpacesThinClient_Environment)(nil)
	_ cfz.Resource                   = (*AWS_WorkSpacesThinClient_Environment)(nil)
)

const (
	// AWS_WorkSpacesThinClient_Environment__Type is the CloudFormation type for AWS::WorkSpacesThinClient::Environment.
	AWS_WorkSpacesThinClient_Environment__Type = "AWS::WorkSpacesThinClient::Environment"
)

var (
	// AWS_WorkSpacesThinClient_Environment__AttributesMap reports all the CloudFormation attributes for AWS::WorkSpacesThinClient::Environment.
	AWS_WorkSpacesThinClient_Environment__AttributesMap = struct {
		ActivationCode              string
		Arn                         string
		CreatedAt                   string
		DesktopType                 string
		Id                          string
		PendingSoftwareSetId        string
		PendingSoftwareSetVersion   string
		RegisteredDevicesCount      string
		SoftwareSetComplianceStatus string
		UpdatedAt                   string
	}{
		ActivationCode:              "ActivationCode",
		Arn:                         "Arn",
		CreatedAt:                   "CreatedAt",
		DesktopType:                 "DesktopType",
		Id:                          "Id",
		PendingSoftwareSetId:        "PendingSoftwareSetId",
		PendingSoftwareSetVersion:   "PendingSoftwareSetVersion",
		RegisteredDevicesCount:      "RegisteredDevicesCount",
		SoftwareSetComplianceStatus: "SoftwareSetComplianceStatus",
		UpdatedAt:                   "UpdatedAt",
	}

	// AWS_WorkSpacesThinClient_Environment__AttributesSlice reports all the CloudFormation attributes for AWS::WorkSpacesThinClient::Environment.
	AWS_WorkSpacesThinClient_Environment__AttributesSlice = []string{
		AWS_WorkSpacesThinClient_Environment__AttributesMap.ActivationCode,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.Arn,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.CreatedAt,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.DesktopType,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.Id,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.PendingSoftwareSetId,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.PendingSoftwareSetVersion,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.RegisteredDevicesCount,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.SoftwareSetComplianceStatus,
		AWS_WorkSpacesThinClient_Environment__AttributesMap.UpdatedAt,
	}
)

var (
	// AWS_WorkSpacesThinClient_Environment__PropertiesMap reports all the CloudFormation properties for AWS::WorkSpacesThinClient::Environment.
	AWS_WorkSpacesThinClient_Environment__PropertiesMap = struct {
		DesiredSoftwareSetId      string
		DesktopArn                string
		DesktopEndpoint           string
		DeviceCreationTags        string
		KmsKeyArn                 string
		MaintenanceWindow         string
		Name                      string
		SoftwareSetUpdateMode     string
		SoftwareSetUpdateSchedule string
		Tags                      string
	}{
		DesiredSoftwareSetId:      "DesiredSoftwareSetId",
		DesktopArn:                "DesktopArn",
		DesktopEndpoint:           "DesktopEndpoint",
		DeviceCreationTags:        "DeviceCreationTags",
		KmsKeyArn:                 "KmsKeyArn",
		MaintenanceWindow:         "MaintenanceWindow",
		Name:                      "Name",
		SoftwareSetUpdateMode:     "SoftwareSetUpdateMode",
		SoftwareSetUpdateSchedule: "SoftwareSetUpdateSchedule",
		Tags:                      "Tags",
	}

	// AWS_WorkSpacesThinClient_Environment__PropertiesSlice reports all the CloudFormation properties for AWS::WorkSpacesThinClient::Environment.
	AWS_WorkSpacesThinClient_Environment__PropertiesSlice = []string{
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.DesiredSoftwareSetId,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.DesktopArn,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.DesktopEndpoint,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.DeviceCreationTags,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.KmsKeyArn,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.MaintenanceWindow,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.Name,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.SoftwareSetUpdateMode,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.SoftwareSetUpdateSchedule,
		AWS_WorkSpacesThinClient_Environment__PropertiesMap.Tags,
	}
)

// AWS_WorkSpacesThinClient_Environment is a binding for AWS::WorkSpacesThinClient::Environment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html
type AWS_WorkSpacesThinClient_Environment struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DesiredSoftwareSetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-desiredsoftwaresetid
	DesiredSoftwareSetId cfz.Expression[string] `json:"DesiredSoftwareSetId,omitempty"`

	// DesktopArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-desktoparn
	DesktopArn cfz.Expression[string] `json:"DesktopArn,omitempty"`

	// DesktopEndpoint is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-desktopendpoint
	DesktopEndpoint cfz.Expression[string] `json:"DesktopEndpoint,omitempty"`

	// DeviceCreationTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-devicecreationtags
	DeviceCreationTags cfz.ExpressionSlice[cfz.Tag] `json:"DeviceCreationTags,omitempty"`

	// KmsKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-kmskeyarn
	KmsKeyArn cfz.Expression[string] `json:"KmsKeyArn,omitempty"`

	// MaintenanceWindow is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-maintenancewindow
	MaintenanceWindow cfz.Expression[AWS_WorkSpacesThinClient_Environment_MaintenanceWindow] `json:"MaintenanceWindow,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// SoftwareSetUpdateMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-softwaresetupdatemode
	SoftwareSetUpdateMode cfz.Expression[string] `json:"SoftwareSetUpdateMode,omitempty"`

	// SoftwareSetUpdateSchedule is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-softwaresetupdateschedule
	SoftwareSetUpdateSchedule cfz.Expression[string] `json:"SoftwareSetUpdateSchedule,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesthinclient-environment.html#cfn-workspacesthinclient-environment-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_WorkSpacesThinClient_Environment initializes a new *AWS_WorkSpacesThinClient_Environment.
func New__AWS_WorkSpacesThinClient_Environment(logicalName string) *AWS_WorkSpacesThinClient_Environment {
	return &AWS_WorkSpacesThinClient_Environment{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_WorkSpacesThinClient_Environment) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_WorkSpacesThinClient_Environment) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_WorkSpacesThinClient_Environment) GetType() string {
	return AWS_WorkSpacesThinClient_Environment__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_WorkSpacesThinClient_Environment) Set__LogicalName(v string) *AWS_WorkSpacesThinClient_Environment {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_WorkSpacesThinClient_Environment) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_WorkSpacesThinClient_Environment {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_WorkSpacesThinClient_Environment) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_WorkSpacesThinClient_Environment {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_WorkSpacesThinClient_Environment) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_WorkSpacesThinClient_Environment {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_WorkSpacesThinClient_Environment) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_WorkSpacesThinClient_Environment {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_WorkSpacesThinClient_Environment) Set__RequestedOutputs(v []cfz.Output) *AWS_WorkSpacesThinClient_Environment {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_WorkSpacesThinClient_Environment) Add__RequestedOutputs(v ...cfz.Output) *AWS_WorkSpacesThinClient_Environment {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DesiredSoftwareSetId updates property "DesiredSoftwareSetId".
func (t *AWS_WorkSpacesThinClient_Environment) Set__DesiredSoftwareSetId(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.DesiredSoftwareSetId = v
	return t
}

// SetV__DesiredSoftwareSetId updates property "DesiredSoftwareSetId".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__DesiredSoftwareSetId(v string) *AWS_WorkSpacesThinClient_Environment {
	t.DesiredSoftwareSetId = cfz.V(v)
	return t
}

// Set__DesktopArn updates property "DesktopArn".
func (t *AWS_WorkSpacesThinClient_Environment) Set__DesktopArn(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.DesktopArn = v
	return t
}

// SetV__DesktopArn updates property "DesktopArn".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__DesktopArn(v string) *AWS_WorkSpacesThinClient_Environment {
	t.DesktopArn = cfz.V(v)
	return t
}

// Set__DesktopEndpoint updates property "DesktopEndpoint".
func (t *AWS_WorkSpacesThinClient_Environment) Set__DesktopEndpoint(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.DesktopEndpoint = v
	return t
}

// SetV__DesktopEndpoint updates property "DesktopEndpoint".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__DesktopEndpoint(v string) *AWS_WorkSpacesThinClient_Environment {
	t.DesktopEndpoint = cfz.V(v)
	return t
}

// Set__DeviceCreationTags updates property "DeviceCreationTags".
func (t *AWS_WorkSpacesThinClient_Environment) Set__DeviceCreationTags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_WorkSpacesThinClient_Environment {
	t.DeviceCreationTags = v
	return t
}

// SetS__DeviceCreationTags updates property "DeviceCreationTags".
func (t *AWS_WorkSpacesThinClient_Environment) SetS__DeviceCreationTags(v ...cfz.Expression[cfz.Tag]) *AWS_WorkSpacesThinClient_Environment {
	t.DeviceCreationTags = cfz.S(v...)
	return t
}

// SetSV__DeviceCreationTags updates property "DeviceCreationTags".
func (t *AWS_WorkSpacesThinClient_Environment) SetSV__DeviceCreationTags(v ...cfz.Tag) *AWS_WorkSpacesThinClient_Environment {
	t.DeviceCreationTags = cfz.SV(v...)
	return t
}

// Set__KmsKeyArn updates property "KmsKeyArn".
func (t *AWS_WorkSpacesThinClient_Environment) Set__KmsKeyArn(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.KmsKeyArn = v
	return t
}

// SetV__KmsKeyArn updates property "KmsKeyArn".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__KmsKeyArn(v string) *AWS_WorkSpacesThinClient_Environment {
	t.KmsKeyArn = cfz.V(v)
	return t
}

// Set__MaintenanceWindow updates property "MaintenanceWindow".
func (t *AWS_WorkSpacesThinClient_Environment) Set__MaintenanceWindow(v cfz.Expression[AWS_WorkSpacesThinClient_Environment_MaintenanceWindow]) *AWS_WorkSpacesThinClient_Environment {
	t.MaintenanceWindow = v
	return t
}

// SetV__MaintenanceWindow updates property "MaintenanceWindow".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__MaintenanceWindow(v AWS_WorkSpacesThinClient_Environment_MaintenanceWindow) *AWS_WorkSpacesThinClient_Environment {
	t.MaintenanceWindow = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_WorkSpacesThinClient_Environment) Set__Name(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__Name(v string) *AWS_WorkSpacesThinClient_Environment {
	t.Name = cfz.V(v)
	return t
}

// Set__SoftwareSetUpdateMode updates property "SoftwareSetUpdateMode".
func (t *AWS_WorkSpacesThinClient_Environment) Set__SoftwareSetUpdateMode(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.SoftwareSetUpdateMode = v
	return t
}

// SetV__SoftwareSetUpdateMode updates property "SoftwareSetUpdateMode".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__SoftwareSetUpdateMode(v string) *AWS_WorkSpacesThinClient_Environment {
	t.SoftwareSetUpdateMode = cfz.V(v)
	return t
}

// Set__SoftwareSetUpdateSchedule updates property "SoftwareSetUpdateSchedule".
func (t *AWS_WorkSpacesThinClient_Environment) Set__SoftwareSetUpdateSchedule(v cfz.Expression[string]) *AWS_WorkSpacesThinClient_Environment {
	t.SoftwareSetUpdateSchedule = v
	return t
}

// SetV__SoftwareSetUpdateSchedule updates property "SoftwareSetUpdateSchedule".
func (t *AWS_WorkSpacesThinClient_Environment) SetV__SoftwareSetUpdateSchedule(v string) *AWS_WorkSpacesThinClient_Environment {
	t.SoftwareSetUpdateSchedule = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_WorkSpacesThinClient_Environment) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_WorkSpacesThinClient_Environment {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_WorkSpacesThinClient_Environment) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_WorkSpacesThinClient_Environment {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_WorkSpacesThinClient_Environment) SetSV__Tags(v ...cfz.Tag) *AWS_WorkSpacesThinClient_Environment {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_WorkSpacesThinClient_Environment) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__ActivationCode returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ActivationCode
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__ActivationCode() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.ActivationCode))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.Arn))
}

// GetAtt__CreatedAt returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedAt
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__CreatedAt() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.CreatedAt))
}

// GetAtt__DesktopType returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DesktopType
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__DesktopType() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.DesktopType))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.Id))
}

// GetAtt__PendingSoftwareSetId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PendingSoftwareSetId
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__PendingSoftwareSetId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.PendingSoftwareSetId))
}

// GetAtt__PendingSoftwareSetVersion returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PendingSoftwareSetVersion
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__PendingSoftwareSetVersion() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.PendingSoftwareSetVersion))
}

// GetAtt__RegisteredDevicesCount returns a $cfz.Expression[int32] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RegisteredDevicesCount
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__RegisteredDevicesCount() cfz.Expression[int32] {
	return cfz.GetAtt[int32](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.RegisteredDevicesCount))
}

// GetAtt__SoftwareSetComplianceStatus returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SoftwareSetComplianceStatus
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__SoftwareSetComplianceStatus() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.SoftwareSetComplianceStatus))
}

// GetAtt__UpdatedAt returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: UpdatedAt
func (t *AWS_WorkSpacesThinClient_Environment) GetAtt__UpdatedAt() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesThinClient_Environment__AttributesMap.UpdatedAt))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ActivationCode returns a conventionally configured output for an attribute of this resource.
// Attribute: ActivationCode
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__ActivationCode(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttActivationCode", t.GetAtt__ActivationCode())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedAt returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedAt
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__CreatedAt(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedAt", t.GetAtt__CreatedAt())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DesktopType returns a conventionally configured output for an attribute of this resource.
// Attribute: DesktopType
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__DesktopType(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDesktopType", t.GetAtt__DesktopType())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PendingSoftwareSetId returns a conventionally configured output for an attribute of this resource.
// Attribute: PendingSoftwareSetId
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__PendingSoftwareSetId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPendingSoftwareSetId", t.GetAtt__PendingSoftwareSetId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PendingSoftwareSetVersion returns a conventionally configured output for an attribute of this resource.
// Attribute: PendingSoftwareSetVersion
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__PendingSoftwareSetVersion(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPendingSoftwareSetVersion", t.GetAtt__PendingSoftwareSetVersion())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RegisteredDevicesCount returns a conventionally configured output for an attribute of this resource.
// Attribute: RegisteredDevicesCount
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__RegisteredDevicesCount(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRegisteredDevicesCount", t.GetAtt__RegisteredDevicesCount())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SoftwareSetComplianceStatus returns a conventionally configured output for an attribute of this resource.
// Attribute: SoftwareSetComplianceStatus
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__SoftwareSetComplianceStatus(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSoftwareSetComplianceStatus", t.GetAtt__SoftwareSetComplianceStatus())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__UpdatedAt returns a conventionally configured output for an attribute of this resource.
// Attribute: UpdatedAt
func (t *AWS_WorkSpacesThinClient_Environment) GetConventionalOutputAtt__UpdatedAt(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttUpdatedAt", t.GetAtt__UpdatedAt())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_WorkSpacesThinClient_Environment) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_WorkSpacesThinClient_Environment

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

func (t *AWS_WorkSpacesThinClient_Environment) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
