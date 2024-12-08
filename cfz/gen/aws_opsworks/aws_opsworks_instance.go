// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opsworks

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_OpsWorks_Instance)(nil)
	_ cfz.Resource                   = (*AWS_OpsWorks_Instance)(nil)
)

const (
	// AWS_OpsWorks_Instance__Type is the CloudFormation type for AWS::OpsWorks::Instance.
	AWS_OpsWorks_Instance__Type = "AWS::OpsWorks::Instance"
)

var (
	// AWS_OpsWorks_Instance__AttributesMap reports all the CloudFormation attributes for AWS::OpsWorks::Instance.
	AWS_OpsWorks_Instance__AttributesMap = struct {
		AvailabilityZone string
		PrivateDnsName   string
		PrivateIp        string
		PublicDnsName    string
		PublicIp         string
	}{
		AvailabilityZone: "AvailabilityZone",
		PrivateDnsName:   "PrivateDnsName",
		PrivateIp:        "PrivateIp",
		PublicDnsName:    "PublicDnsName",
		PublicIp:         "PublicIp",
	}

	// AWS_OpsWorks_Instance__AttributesSlice reports all the CloudFormation attributes for AWS::OpsWorks::Instance.
	AWS_OpsWorks_Instance__AttributesSlice = []string{
		AWS_OpsWorks_Instance__AttributesMap.AvailabilityZone,
		AWS_OpsWorks_Instance__AttributesMap.PrivateDnsName,
		AWS_OpsWorks_Instance__AttributesMap.PrivateIp,
		AWS_OpsWorks_Instance__AttributesMap.PublicDnsName,
		AWS_OpsWorks_Instance__AttributesMap.PublicIp,
	}
)

var (
	// AWS_OpsWorks_Instance__PropertiesMap reports all the CloudFormation properties for AWS::OpsWorks::Instance.
	AWS_OpsWorks_Instance__PropertiesMap = struct {
		AgentVersion         string
		AmiId                string
		Architecture         string
		AutoScalingType      string
		AvailabilityZone     string
		BlockDeviceMappings  string
		EbsOptimized         string
		ElasticIps           string
		Hostname             string
		InstallUpdatesOnBoot string
		InstanceType         string
		LayerIds             string
		Os                   string
		RootDeviceType       string
		SshKeyName           string
		StackId              string
		SubnetId             string
		Tenancy              string
		TimeBasedAutoScaling string
		VirtualizationType   string
		Volumes              string
	}{
		AgentVersion:         "AgentVersion",
		AmiId:                "AmiId",
		Architecture:         "Architecture",
		AutoScalingType:      "AutoScalingType",
		AvailabilityZone:     "AvailabilityZone",
		BlockDeviceMappings:  "BlockDeviceMappings",
		EbsOptimized:         "EbsOptimized",
		ElasticIps:           "ElasticIps",
		Hostname:             "Hostname",
		InstallUpdatesOnBoot: "InstallUpdatesOnBoot",
		InstanceType:         "InstanceType",
		LayerIds:             "LayerIds",
		Os:                   "Os",
		RootDeviceType:       "RootDeviceType",
		SshKeyName:           "SshKeyName",
		StackId:              "StackId",
		SubnetId:             "SubnetId",
		Tenancy:              "Tenancy",
		TimeBasedAutoScaling: "TimeBasedAutoScaling",
		VirtualizationType:   "VirtualizationType",
		Volumes:              "Volumes",
	}

	// AWS_OpsWorks_Instance__PropertiesSlice reports all the CloudFormation properties for AWS::OpsWorks::Instance.
	AWS_OpsWorks_Instance__PropertiesSlice = []string{
		AWS_OpsWorks_Instance__PropertiesMap.AgentVersion,
		AWS_OpsWorks_Instance__PropertiesMap.AmiId,
		AWS_OpsWorks_Instance__PropertiesMap.Architecture,
		AWS_OpsWorks_Instance__PropertiesMap.AutoScalingType,
		AWS_OpsWorks_Instance__PropertiesMap.AvailabilityZone,
		AWS_OpsWorks_Instance__PropertiesMap.BlockDeviceMappings,
		AWS_OpsWorks_Instance__PropertiesMap.EbsOptimized,
		AWS_OpsWorks_Instance__PropertiesMap.ElasticIps,
		AWS_OpsWorks_Instance__PropertiesMap.Hostname,
		AWS_OpsWorks_Instance__PropertiesMap.InstallUpdatesOnBoot,
		AWS_OpsWorks_Instance__PropertiesMap.InstanceType,
		AWS_OpsWorks_Instance__PropertiesMap.LayerIds,
		AWS_OpsWorks_Instance__PropertiesMap.Os,
		AWS_OpsWorks_Instance__PropertiesMap.RootDeviceType,
		AWS_OpsWorks_Instance__PropertiesMap.SshKeyName,
		AWS_OpsWorks_Instance__PropertiesMap.StackId,
		AWS_OpsWorks_Instance__PropertiesMap.SubnetId,
		AWS_OpsWorks_Instance__PropertiesMap.Tenancy,
		AWS_OpsWorks_Instance__PropertiesMap.TimeBasedAutoScaling,
		AWS_OpsWorks_Instance__PropertiesMap.VirtualizationType,
		AWS_OpsWorks_Instance__PropertiesMap.Volumes,
	}
)

// AWS_OpsWorks_Instance is a binding for AWS::OpsWorks::Instance.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
type AWS_OpsWorks_Instance struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AgentVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-agentversion
	AgentVersion cfz.Expression[string] `json:"AgentVersion,omitempty"`

	// AmiId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-amiid
	AmiId cfz.Expression[string] `json:"AmiId,omitempty"`

	// Architecture is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-architecture
	Architecture cfz.Expression[string] `json:"Architecture,omitempty"`

	// AutoScalingType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-autoscalingtype
	AutoScalingType cfz.Expression[string] `json:"AutoScalingType,omitempty"`

	// AvailabilityZone is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-availabilityzone
	AvailabilityZone cfz.Expression[string] `json:"AvailabilityZone,omitempty"`

	// BlockDeviceMappings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-blockdevicemappings
	BlockDeviceMappings cfz.ExpressionSlice[AWS_OpsWorks_Instance_BlockDeviceMapping] `json:"BlockDeviceMappings,omitempty"`

	// EbsOptimized is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-ebsoptimized
	EbsOptimized cfz.Expression[bool] `json:"EbsOptimized,omitempty"`

	// ElasticIps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-elasticips
	ElasticIps cfz.ExpressionSlice[string] `json:"ElasticIps,omitempty"`

	// Hostname is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-hostname
	Hostname cfz.Expression[string] `json:"Hostname,omitempty"`

	// InstallUpdatesOnBoot is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-installupdatesonboot
	InstallUpdatesOnBoot cfz.Expression[bool] `json:"InstallUpdatesOnBoot,omitempty"`

	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`

	// LayerIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-layerids
	LayerIds cfz.ExpressionSlice[string] `json:"LayerIds,omitempty"`

	// Os is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-os
	Os cfz.Expression[string] `json:"Os,omitempty"`

	// RootDeviceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-rootdevicetype
	RootDeviceType cfz.Expression[string] `json:"RootDeviceType,omitempty"`

	// SshKeyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-sshkeyname
	SshKeyName cfz.Expression[string] `json:"SshKeyName,omitempty"`

	// StackId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-stackid
	StackId cfz.Expression[string] `json:"StackId,omitempty"`

	// SubnetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-subnetid
	SubnetId cfz.Expression[string] `json:"SubnetId,omitempty"`

	// Tenancy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-tenancy
	Tenancy cfz.Expression[string] `json:"Tenancy,omitempty"`

	// TimeBasedAutoScaling is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-timebasedautoscaling
	TimeBasedAutoScaling cfz.Expression[AWS_OpsWorks_Instance_TimeBasedAutoScaling] `json:"TimeBasedAutoScaling,omitempty"`

	// VirtualizationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-virtualizationtype
	VirtualizationType cfz.Expression[string] `json:"VirtualizationType,omitempty"`

	// Volumes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-volumes
	Volumes cfz.ExpressionSlice[string] `json:"Volumes,omitempty"`
}

// New__AWS_OpsWorks_Instance initializes a new *AWS_OpsWorks_Instance.
func New__AWS_OpsWorks_Instance(logicalName string) *AWS_OpsWorks_Instance {
	return &AWS_OpsWorks_Instance{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_OpsWorks_Instance) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_OpsWorks_Instance) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_OpsWorks_Instance) GetType() string {
	return AWS_OpsWorks_Instance__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_OpsWorks_Instance) Set__LogicalName(v string) *AWS_OpsWorks_Instance {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_OpsWorks_Instance) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_OpsWorks_Instance {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_OpsWorks_Instance) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_OpsWorks_Instance {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_OpsWorks_Instance) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_OpsWorks_Instance {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_OpsWorks_Instance) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_OpsWorks_Instance {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_OpsWorks_Instance) Set__RequestedOutputs(v []cfz.Output) *AWS_OpsWorks_Instance {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_OpsWorks_Instance) Add__RequestedOutputs(v ...cfz.Output) *AWS_OpsWorks_Instance {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AgentVersion updates property "AgentVersion".
func (t *AWS_OpsWorks_Instance) Set__AgentVersion(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.AgentVersion = v
	return t
}

// SetV__AgentVersion updates property "AgentVersion".
func (t *AWS_OpsWorks_Instance) SetV__AgentVersion(v string) *AWS_OpsWorks_Instance {
	t.AgentVersion = cfz.V(v)
	return t
}

// Set__AmiId updates property "AmiId".
func (t *AWS_OpsWorks_Instance) Set__AmiId(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.AmiId = v
	return t
}

// SetV__AmiId updates property "AmiId".
func (t *AWS_OpsWorks_Instance) SetV__AmiId(v string) *AWS_OpsWorks_Instance {
	t.AmiId = cfz.V(v)
	return t
}

// Set__Architecture updates property "Architecture".
func (t *AWS_OpsWorks_Instance) Set__Architecture(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.Architecture = v
	return t
}

// SetV__Architecture updates property "Architecture".
func (t *AWS_OpsWorks_Instance) SetV__Architecture(v string) *AWS_OpsWorks_Instance {
	t.Architecture = cfz.V(v)
	return t
}

// Set__AutoScalingType updates property "AutoScalingType".
func (t *AWS_OpsWorks_Instance) Set__AutoScalingType(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.AutoScalingType = v
	return t
}

// SetV__AutoScalingType updates property "AutoScalingType".
func (t *AWS_OpsWorks_Instance) SetV__AutoScalingType(v string) *AWS_OpsWorks_Instance {
	t.AutoScalingType = cfz.V(v)
	return t
}

// Set__AvailabilityZone updates property "AvailabilityZone".
func (t *AWS_OpsWorks_Instance) Set__AvailabilityZone(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.AvailabilityZone = v
	return t
}

// SetV__AvailabilityZone updates property "AvailabilityZone".
func (t *AWS_OpsWorks_Instance) SetV__AvailabilityZone(v string) *AWS_OpsWorks_Instance {
	t.AvailabilityZone = cfz.V(v)
	return t
}

// Set__BlockDeviceMappings updates property "BlockDeviceMappings".
func (t *AWS_OpsWorks_Instance) Set__BlockDeviceMappings(v cfz.ExpressionSlice[AWS_OpsWorks_Instance_BlockDeviceMapping]) *AWS_OpsWorks_Instance {
	t.BlockDeviceMappings = v
	return t
}

// SetS__BlockDeviceMappings updates property "BlockDeviceMappings".
func (t *AWS_OpsWorks_Instance) SetS__BlockDeviceMappings(v ...cfz.Expression[AWS_OpsWorks_Instance_BlockDeviceMapping]) *AWS_OpsWorks_Instance {
	t.BlockDeviceMappings = cfz.S(v...)
	return t
}

// SetSV__BlockDeviceMappings updates property "BlockDeviceMappings".
func (t *AWS_OpsWorks_Instance) SetSV__BlockDeviceMappings(v ...AWS_OpsWorks_Instance_BlockDeviceMapping) *AWS_OpsWorks_Instance {
	t.BlockDeviceMappings = cfz.SV(v...)
	return t
}

// Set__EbsOptimized updates property "EbsOptimized".
func (t *AWS_OpsWorks_Instance) Set__EbsOptimized(v cfz.Expression[bool]) *AWS_OpsWorks_Instance {
	t.EbsOptimized = v
	return t
}

// SetV__EbsOptimized updates property "EbsOptimized".
func (t *AWS_OpsWorks_Instance) SetV__EbsOptimized(v bool) *AWS_OpsWorks_Instance {
	t.EbsOptimized = cfz.V(v)
	return t
}

// Set__ElasticIps updates property "ElasticIps".
func (t *AWS_OpsWorks_Instance) Set__ElasticIps(v cfz.ExpressionSlice[string]) *AWS_OpsWorks_Instance {
	t.ElasticIps = v
	return t
}

// SetS__ElasticIps updates property "ElasticIps".
func (t *AWS_OpsWorks_Instance) SetS__ElasticIps(v ...cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.ElasticIps = cfz.S(v...)
	return t
}

// SetSV__ElasticIps updates property "ElasticIps".
func (t *AWS_OpsWorks_Instance) SetSV__ElasticIps(v ...string) *AWS_OpsWorks_Instance {
	t.ElasticIps = cfz.SV(v...)
	return t
}

// Set__Hostname updates property "Hostname".
func (t *AWS_OpsWorks_Instance) Set__Hostname(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.Hostname = v
	return t
}

// SetV__Hostname updates property "Hostname".
func (t *AWS_OpsWorks_Instance) SetV__Hostname(v string) *AWS_OpsWorks_Instance {
	t.Hostname = cfz.V(v)
	return t
}

// Set__InstallUpdatesOnBoot updates property "InstallUpdatesOnBoot".
func (t *AWS_OpsWorks_Instance) Set__InstallUpdatesOnBoot(v cfz.Expression[bool]) *AWS_OpsWorks_Instance {
	t.InstallUpdatesOnBoot = v
	return t
}

// SetV__InstallUpdatesOnBoot updates property "InstallUpdatesOnBoot".
func (t *AWS_OpsWorks_Instance) SetV__InstallUpdatesOnBoot(v bool) *AWS_OpsWorks_Instance {
	t.InstallUpdatesOnBoot = cfz.V(v)
	return t
}

// Set__InstanceType updates property "InstanceType".
func (t *AWS_OpsWorks_Instance) Set__InstanceType(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t *AWS_OpsWorks_Instance) SetV__InstanceType(v string) *AWS_OpsWorks_Instance {
	t.InstanceType = cfz.V(v)
	return t
}

// Set__LayerIds updates property "LayerIds".
func (t *AWS_OpsWorks_Instance) Set__LayerIds(v cfz.ExpressionSlice[string]) *AWS_OpsWorks_Instance {
	t.LayerIds = v
	return t
}

// SetS__LayerIds updates property "LayerIds".
func (t *AWS_OpsWorks_Instance) SetS__LayerIds(v ...cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.LayerIds = cfz.S(v...)
	return t
}

// SetSV__LayerIds updates property "LayerIds".
func (t *AWS_OpsWorks_Instance) SetSV__LayerIds(v ...string) *AWS_OpsWorks_Instance {
	t.LayerIds = cfz.SV(v...)
	return t
}

// Set__Os updates property "Os".
func (t *AWS_OpsWorks_Instance) Set__Os(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.Os = v
	return t
}

// SetV__Os updates property "Os".
func (t *AWS_OpsWorks_Instance) SetV__Os(v string) *AWS_OpsWorks_Instance {
	t.Os = cfz.V(v)
	return t
}

// Set__RootDeviceType updates property "RootDeviceType".
func (t *AWS_OpsWorks_Instance) Set__RootDeviceType(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.RootDeviceType = v
	return t
}

// SetV__RootDeviceType updates property "RootDeviceType".
func (t *AWS_OpsWorks_Instance) SetV__RootDeviceType(v string) *AWS_OpsWorks_Instance {
	t.RootDeviceType = cfz.V(v)
	return t
}

// Set__SshKeyName updates property "SshKeyName".
func (t *AWS_OpsWorks_Instance) Set__SshKeyName(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.SshKeyName = v
	return t
}

// SetV__SshKeyName updates property "SshKeyName".
func (t *AWS_OpsWorks_Instance) SetV__SshKeyName(v string) *AWS_OpsWorks_Instance {
	t.SshKeyName = cfz.V(v)
	return t
}

// Set__StackId updates property "StackId".
func (t *AWS_OpsWorks_Instance) Set__StackId(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.StackId = v
	return t
}

// SetV__StackId updates property "StackId".
func (t *AWS_OpsWorks_Instance) SetV__StackId(v string) *AWS_OpsWorks_Instance {
	t.StackId = cfz.V(v)
	return t
}

// Set__SubnetId updates property "SubnetId".
func (t *AWS_OpsWorks_Instance) Set__SubnetId(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.SubnetId = v
	return t
}

// SetV__SubnetId updates property "SubnetId".
func (t *AWS_OpsWorks_Instance) SetV__SubnetId(v string) *AWS_OpsWorks_Instance {
	t.SubnetId = cfz.V(v)
	return t
}

// Set__Tenancy updates property "Tenancy".
func (t *AWS_OpsWorks_Instance) Set__Tenancy(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.Tenancy = v
	return t
}

// SetV__Tenancy updates property "Tenancy".
func (t *AWS_OpsWorks_Instance) SetV__Tenancy(v string) *AWS_OpsWorks_Instance {
	t.Tenancy = cfz.V(v)
	return t
}

// Set__TimeBasedAutoScaling updates property "TimeBasedAutoScaling".
func (t *AWS_OpsWorks_Instance) Set__TimeBasedAutoScaling(v cfz.Expression[AWS_OpsWorks_Instance_TimeBasedAutoScaling]) *AWS_OpsWorks_Instance {
	t.TimeBasedAutoScaling = v
	return t
}

// SetV__TimeBasedAutoScaling updates property "TimeBasedAutoScaling".
func (t *AWS_OpsWorks_Instance) SetV__TimeBasedAutoScaling(v AWS_OpsWorks_Instance_TimeBasedAutoScaling) *AWS_OpsWorks_Instance {
	t.TimeBasedAutoScaling = cfz.V(v)
	return t
}

// Set__VirtualizationType updates property "VirtualizationType".
func (t *AWS_OpsWorks_Instance) Set__VirtualizationType(v cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.VirtualizationType = v
	return t
}

// SetV__VirtualizationType updates property "VirtualizationType".
func (t *AWS_OpsWorks_Instance) SetV__VirtualizationType(v string) *AWS_OpsWorks_Instance {
	t.VirtualizationType = cfz.V(v)
	return t
}

// Set__Volumes updates property "Volumes".
func (t *AWS_OpsWorks_Instance) Set__Volumes(v cfz.ExpressionSlice[string]) *AWS_OpsWorks_Instance {
	t.Volumes = v
	return t
}

// SetS__Volumes updates property "Volumes".
func (t *AWS_OpsWorks_Instance) SetS__Volumes(v ...cfz.Expression[string]) *AWS_OpsWorks_Instance {
	t.Volumes = cfz.S(v...)
	return t
}

// SetSV__Volumes updates property "Volumes".
func (t *AWS_OpsWorks_Instance) SetSV__Volumes(v ...string) *AWS_OpsWorks_Instance {
	t.Volumes = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_OpsWorks_Instance) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AvailabilityZone returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AvailabilityZone
func (t *AWS_OpsWorks_Instance) GetAtt__AvailabilityZone() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_OpsWorks_Instance__AttributesMap.AvailabilityZone))
}

// GetAtt__PrivateDnsName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PrivateDnsName
func (t *AWS_OpsWorks_Instance) GetAtt__PrivateDnsName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_OpsWorks_Instance__AttributesMap.PrivateDnsName))
}

// GetAtt__PrivateIp returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PrivateIp
func (t *AWS_OpsWorks_Instance) GetAtt__PrivateIp() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_OpsWorks_Instance__AttributesMap.PrivateIp))
}

// GetAtt__PublicDnsName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PublicDnsName
func (t *AWS_OpsWorks_Instance) GetAtt__PublicDnsName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_OpsWorks_Instance__AttributesMap.PublicDnsName))
}

// GetAtt__PublicIp returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PublicIp
func (t *AWS_OpsWorks_Instance) GetAtt__PublicIp() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_OpsWorks_Instance__AttributesMap.PublicIp))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_OpsWorks_Instance) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AvailabilityZone returns a conventionally configured output for an attribute of this resource.
// Attribute: AvailabilityZone
func (t *AWS_OpsWorks_Instance) GetConventionalOutputAtt__AvailabilityZone(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAvailabilityZone", t.GetAtt__AvailabilityZone())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PrivateDnsName returns a conventionally configured output for an attribute of this resource.
// Attribute: PrivateDnsName
func (t *AWS_OpsWorks_Instance) GetConventionalOutputAtt__PrivateDnsName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPrivateDnsName", t.GetAtt__PrivateDnsName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PrivateIp returns a conventionally configured output for an attribute of this resource.
// Attribute: PrivateIp
func (t *AWS_OpsWorks_Instance) GetConventionalOutputAtt__PrivateIp(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPrivateIp", t.GetAtt__PrivateIp())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PublicDnsName returns a conventionally configured output for an attribute of this resource.
// Attribute: PublicDnsName
func (t *AWS_OpsWorks_Instance) GetConventionalOutputAtt__PublicDnsName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPublicDnsName", t.GetAtt__PublicDnsName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PublicIp returns a conventionally configured output for an attribute of this resource.
// Attribute: PublicIp
func (t *AWS_OpsWorks_Instance) GetConventionalOutputAtt__PublicIp(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPublicIp", t.GetAtt__PublicIp())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_OpsWorks_Instance) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_OpsWorks_Instance

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

func (t *AWS_OpsWorks_Instance) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
