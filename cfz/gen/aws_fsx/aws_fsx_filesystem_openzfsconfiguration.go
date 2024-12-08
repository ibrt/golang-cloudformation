// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_fsx

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_FSx_FileSystem_OpenZFSConfiguration__Type is the CloudFormation type for AWS::FSx::FileSystem.OpenZFSConfiguration.
	AWS_FSx_FileSystem_OpenZFSConfiguration__Type = "AWS::FSx::FileSystem.OpenZFSConfiguration"
)

var (
	// AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::FSx::FileSystem.OpenZFSConfiguration.
	AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap = struct {
		AutomaticBackupRetentionDays  string
		CopyTagsToBackups             string
		CopyTagsToVolumes             string
		DailyAutomaticBackupStartTime string
		DeploymentType                string
		DiskIopsConfiguration         string
		EndpointIpAddressRange        string
		Options                       string
		PreferredSubnetId             string
		ReadCacheConfiguration        string
		RootVolumeConfiguration       string
		RouteTableIds                 string
		ThroughputCapacity            string
		WeeklyMaintenanceStartTime    string
	}{
		AutomaticBackupRetentionDays:  "AutomaticBackupRetentionDays",
		CopyTagsToBackups:             "CopyTagsToBackups",
		CopyTagsToVolumes:             "CopyTagsToVolumes",
		DailyAutomaticBackupStartTime: "DailyAutomaticBackupStartTime",
		DeploymentType:                "DeploymentType",
		DiskIopsConfiguration:         "DiskIopsConfiguration",
		EndpointIpAddressRange:        "EndpointIpAddressRange",
		Options:                       "Options",
		PreferredSubnetId:             "PreferredSubnetId",
		ReadCacheConfiguration:        "ReadCacheConfiguration",
		RootVolumeConfiguration:       "RootVolumeConfiguration",
		RouteTableIds:                 "RouteTableIds",
		ThroughputCapacity:            "ThroughputCapacity",
		WeeklyMaintenanceStartTime:    "WeeklyMaintenanceStartTime",
	}

	// AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::FSx::FileSystem.OpenZFSConfiguration.
	AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesSlice = []string{
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.AutomaticBackupRetentionDays,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.CopyTagsToBackups,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.CopyTagsToVolumes,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.DailyAutomaticBackupStartTime,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.DeploymentType,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.DiskIopsConfiguration,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.EndpointIpAddressRange,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.Options,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.PreferredSubnetId,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.ReadCacheConfiguration,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.RootVolumeConfiguration,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.RouteTableIds,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.ThroughputCapacity,
		AWS_FSx_FileSystem_OpenZFSConfiguration__PropertiesMap.WeeklyMaintenanceStartTime,
	}
)

// AWS_FSx_FileSystem_OpenZFSConfiguration is a binding for AWS::FSx::FileSystem.OpenZFSConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html
type AWS_FSx_FileSystem_OpenZFSConfiguration struct {
	// AutomaticBackupRetentionDays is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-automaticbackupretentiondays
	AutomaticBackupRetentionDays cfz.Expression[int32] `json:"AutomaticBackupRetentionDays,omitempty"`

	// CopyTagsToBackups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-copytagstobackups
	CopyTagsToBackups cfz.Expression[bool] `json:"CopyTagsToBackups,omitempty"`

	// CopyTagsToVolumes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-copytagstovolumes
	CopyTagsToVolumes cfz.Expression[bool] `json:"CopyTagsToVolumes,omitempty"`

	// DailyAutomaticBackupStartTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-dailyautomaticbackupstarttime
	DailyAutomaticBackupStartTime cfz.Expression[string] `json:"DailyAutomaticBackupStartTime,omitempty"`

	// DeploymentType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-deploymenttype
	DeploymentType cfz.Expression[string] `json:"DeploymentType,omitempty"`

	// DiskIopsConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-diskiopsconfiguration
	DiskIopsConfiguration cfz.Expression[AWS_FSx_FileSystem_DiskIopsConfiguration] `json:"DiskIopsConfiguration,omitempty"`

	// EndpointIpAddressRange is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-endpointipaddressrange
	EndpointIpAddressRange cfz.Expression[string] `json:"EndpointIpAddressRange,omitempty"`

	// Options is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-options
	Options cfz.ExpressionSlice[string] `json:"Options,omitempty"`

	// PreferredSubnetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-preferredsubnetid
	PreferredSubnetId cfz.Expression[string] `json:"PreferredSubnetId,omitempty"`

	// ReadCacheConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-readcacheconfiguration
	ReadCacheConfiguration cfz.Expression[AWS_FSx_FileSystem_ReadCacheConfiguration] `json:"ReadCacheConfiguration,omitempty"`

	// RootVolumeConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-rootvolumeconfiguration
	RootVolumeConfiguration cfz.Expression[AWS_FSx_FileSystem_RootVolumeConfiguration] `json:"RootVolumeConfiguration,omitempty"`

	// RouteTableIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-routetableids
	RouteTableIds cfz.ExpressionSlice[string] `json:"RouteTableIds,omitempty"`

	// ThroughputCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-throughputcapacity
	ThroughputCapacity cfz.Expression[int32] `json:"ThroughputCapacity,omitempty"`

	// WeeklyMaintenanceStartTime is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-openzfsconfiguration.html#cfn-fsx-filesystem-openzfsconfiguration-weeklymaintenancestarttime
	WeeklyMaintenanceStartTime cfz.Expression[string] `json:"WeeklyMaintenanceStartTime,omitempty"`
}

// New__AWS_FSx_FileSystem_OpenZFSConfiguration initializes a new AWS_FSx_FileSystem_OpenZFSConfiguration.
func New__AWS_FSx_FileSystem_OpenZFSConfiguration() AWS_FSx_FileSystem_OpenZFSConfiguration {
	return AWS_FSx_FileSystem_OpenZFSConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_FSx_FileSystem_OpenZFSConfiguration) GetType() string {
	return AWS_FSx_FileSystem_OpenZFSConfiguration__Type
}

// Set__AutomaticBackupRetentionDays updates property "AutomaticBackupRetentionDays".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__AutomaticBackupRetentionDays(v cfz.Expression[int32]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.AutomaticBackupRetentionDays = v
	return t
}

// SetV__AutomaticBackupRetentionDays updates property "AutomaticBackupRetentionDays".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__AutomaticBackupRetentionDays(v int32) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.AutomaticBackupRetentionDays = cfz.V(v)
	return t
}

// Set__CopyTagsToBackups updates property "CopyTagsToBackups".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__CopyTagsToBackups(v cfz.Expression[bool]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.CopyTagsToBackups = v
	return t
}

// SetV__CopyTagsToBackups updates property "CopyTagsToBackups".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__CopyTagsToBackups(v bool) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.CopyTagsToBackups = cfz.V(v)
	return t
}

// Set__CopyTagsToVolumes updates property "CopyTagsToVolumes".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__CopyTagsToVolumes(v cfz.Expression[bool]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.CopyTagsToVolumes = v
	return t
}

// SetV__CopyTagsToVolumes updates property "CopyTagsToVolumes".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__CopyTagsToVolumes(v bool) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.CopyTagsToVolumes = cfz.V(v)
	return t
}

// Set__DailyAutomaticBackupStartTime updates property "DailyAutomaticBackupStartTime".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__DailyAutomaticBackupStartTime(v cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.DailyAutomaticBackupStartTime = v
	return t
}

// SetV__DailyAutomaticBackupStartTime updates property "DailyAutomaticBackupStartTime".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__DailyAutomaticBackupStartTime(v string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.DailyAutomaticBackupStartTime = cfz.V(v)
	return t
}

// Set__DeploymentType updates property "DeploymentType".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__DeploymentType(v cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.DeploymentType = v
	return t
}

// SetV__DeploymentType updates property "DeploymentType".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__DeploymentType(v string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.DeploymentType = cfz.V(v)
	return t
}

// Set__DiskIopsConfiguration updates property "DiskIopsConfiguration".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__DiskIopsConfiguration(v cfz.Expression[AWS_FSx_FileSystem_DiskIopsConfiguration]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.DiskIopsConfiguration = v
	return t
}

// SetV__DiskIopsConfiguration updates property "DiskIopsConfiguration".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__DiskIopsConfiguration(v AWS_FSx_FileSystem_DiskIopsConfiguration) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.DiskIopsConfiguration = cfz.V(v)
	return t
}

// Set__EndpointIpAddressRange updates property "EndpointIpAddressRange".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__EndpointIpAddressRange(v cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.EndpointIpAddressRange = v
	return t
}

// SetV__EndpointIpAddressRange updates property "EndpointIpAddressRange".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__EndpointIpAddressRange(v string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.EndpointIpAddressRange = cfz.V(v)
	return t
}

// Set__Options updates property "Options".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__Options(v cfz.ExpressionSlice[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.Options = v
	return t
}

// SetS__Options updates property "Options".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetS__Options(v ...cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.Options = cfz.S(v...)
	return t
}

// SetSV__Options updates property "Options".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetSV__Options(v ...string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.Options = cfz.SV(v...)
	return t
}

// Set__PreferredSubnetId updates property "PreferredSubnetId".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__PreferredSubnetId(v cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.PreferredSubnetId = v
	return t
}

// SetV__PreferredSubnetId updates property "PreferredSubnetId".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__PreferredSubnetId(v string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.PreferredSubnetId = cfz.V(v)
	return t
}

// Set__ReadCacheConfiguration updates property "ReadCacheConfiguration".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__ReadCacheConfiguration(v cfz.Expression[AWS_FSx_FileSystem_ReadCacheConfiguration]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.ReadCacheConfiguration = v
	return t
}

// SetV__ReadCacheConfiguration updates property "ReadCacheConfiguration".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__ReadCacheConfiguration(v AWS_FSx_FileSystem_ReadCacheConfiguration) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.ReadCacheConfiguration = cfz.V(v)
	return t
}

// Set__RootVolumeConfiguration updates property "RootVolumeConfiguration".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__RootVolumeConfiguration(v cfz.Expression[AWS_FSx_FileSystem_RootVolumeConfiguration]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.RootVolumeConfiguration = v
	return t
}

// SetV__RootVolumeConfiguration updates property "RootVolumeConfiguration".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__RootVolumeConfiguration(v AWS_FSx_FileSystem_RootVolumeConfiguration) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.RootVolumeConfiguration = cfz.V(v)
	return t
}

// Set__RouteTableIds updates property "RouteTableIds".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__RouteTableIds(v cfz.ExpressionSlice[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.RouteTableIds = v
	return t
}

// SetS__RouteTableIds updates property "RouteTableIds".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetS__RouteTableIds(v ...cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.RouteTableIds = cfz.S(v...)
	return t
}

// SetSV__RouteTableIds updates property "RouteTableIds".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetSV__RouteTableIds(v ...string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.RouteTableIds = cfz.SV(v...)
	return t
}

// Set__ThroughputCapacity updates property "ThroughputCapacity".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__ThroughputCapacity(v cfz.Expression[int32]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.ThroughputCapacity = v
	return t
}

// SetV__ThroughputCapacity updates property "ThroughputCapacity".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__ThroughputCapacity(v int32) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.ThroughputCapacity = cfz.V(v)
	return t
}

// Set__WeeklyMaintenanceStartTime updates property "WeeklyMaintenanceStartTime".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) Set__WeeklyMaintenanceStartTime(v cfz.Expression[string]) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.WeeklyMaintenanceStartTime = v
	return t
}

// SetV__WeeklyMaintenanceStartTime updates property "WeeklyMaintenanceStartTime".
func (t AWS_FSx_FileSystem_OpenZFSConfiguration) SetV__WeeklyMaintenanceStartTime(v string) AWS_FSx_FileSystem_OpenZFSConfiguration {
	t.WeeklyMaintenanceStartTime = cfz.V(v)
	return t
}
