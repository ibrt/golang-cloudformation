// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_efs

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EFS_FileSystem)(nil)
	_ cfz.Resource                   = (*AWS_EFS_FileSystem)(nil)
)

const (
	// AWS_EFS_FileSystem__Type is the CloudFormation type for AWS::EFS::FileSystem.
	AWS_EFS_FileSystem__Type = "AWS::EFS::FileSystem"
)

var (
	// AWS_EFS_FileSystem__AttributesMap reports all the CloudFormation attributes for AWS::EFS::FileSystem.
	AWS_EFS_FileSystem__AttributesMap = struct {
		Arn          string
		FileSystemId string
	}{
		Arn:          "Arn",
		FileSystemId: "FileSystemId",
	}

	// AWS_EFS_FileSystem__AttributesSlice reports all the CloudFormation attributes for AWS::EFS::FileSystem.
	AWS_EFS_FileSystem__AttributesSlice = []string{
		AWS_EFS_FileSystem__AttributesMap.Arn,
		AWS_EFS_FileSystem__AttributesMap.FileSystemId,
	}
)

var (
	// AWS_EFS_FileSystem__PropertiesMap reports all the CloudFormation properties for AWS::EFS::FileSystem.
	AWS_EFS_FileSystem__PropertiesMap = struct {
		AvailabilityZoneName           string
		BackupPolicy                   string
		BypassPolicyLockoutSafetyCheck string
		Encrypted                      string
		FileSystemPolicy               string
		FileSystemProtection           string
		FileSystemTags                 string
		KmsKeyId                       string
		LifecyclePolicies              string
		PerformanceMode                string
		ProvisionedThroughputInMibps   string
		ReplicationConfiguration       string
		ThroughputMode                 string
	}{
		AvailabilityZoneName:           "AvailabilityZoneName",
		BackupPolicy:                   "BackupPolicy",
		BypassPolicyLockoutSafetyCheck: "BypassPolicyLockoutSafetyCheck",
		Encrypted:                      "Encrypted",
		FileSystemPolicy:               "FileSystemPolicy",
		FileSystemProtection:           "FileSystemProtection",
		FileSystemTags:                 "FileSystemTags",
		KmsKeyId:                       "KmsKeyId",
		LifecyclePolicies:              "LifecyclePolicies",
		PerformanceMode:                "PerformanceMode",
		ProvisionedThroughputInMibps:   "ProvisionedThroughputInMibps",
		ReplicationConfiguration:       "ReplicationConfiguration",
		ThroughputMode:                 "ThroughputMode",
	}

	// AWS_EFS_FileSystem__PropertiesSlice reports all the CloudFormation properties for AWS::EFS::FileSystem.
	AWS_EFS_FileSystem__PropertiesSlice = []string{
		AWS_EFS_FileSystem__PropertiesMap.AvailabilityZoneName,
		AWS_EFS_FileSystem__PropertiesMap.BackupPolicy,
		AWS_EFS_FileSystem__PropertiesMap.BypassPolicyLockoutSafetyCheck,
		AWS_EFS_FileSystem__PropertiesMap.Encrypted,
		AWS_EFS_FileSystem__PropertiesMap.FileSystemPolicy,
		AWS_EFS_FileSystem__PropertiesMap.FileSystemProtection,
		AWS_EFS_FileSystem__PropertiesMap.FileSystemTags,
		AWS_EFS_FileSystem__PropertiesMap.KmsKeyId,
		AWS_EFS_FileSystem__PropertiesMap.LifecyclePolicies,
		AWS_EFS_FileSystem__PropertiesMap.PerformanceMode,
		AWS_EFS_FileSystem__PropertiesMap.ProvisionedThroughputInMibps,
		AWS_EFS_FileSystem__PropertiesMap.ReplicationConfiguration,
		AWS_EFS_FileSystem__PropertiesMap.ThroughputMode,
	}
)

// AWS_EFS_FileSystem is a binding for AWS::EFS::FileSystem.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
type AWS_EFS_FileSystem struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AvailabilityZoneName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-availabilityzonename
	AvailabilityZoneName cfz.Expression[string] `json:"AvailabilityZoneName,omitempty"`

	// BackupPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-backuppolicy
	BackupPolicy cfz.Expression[AWS_EFS_FileSystem_BackupPolicy] `json:"BackupPolicy,omitempty"`

	// BypassPolicyLockoutSafetyCheck is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-bypasspolicylockoutsafetycheck
	BypassPolicyLockoutSafetyCheck cfz.Expression[bool] `json:"BypassPolicyLockoutSafetyCheck,omitempty"`

	// Encrypted is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-encrypted
	Encrypted cfz.Expression[bool] `json:"Encrypted,omitempty"`

	// FileSystemPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystempolicy
	FileSystemPolicy cfz.Expression[json.RawMessage] `json:"FileSystemPolicy,omitempty"`

	// FileSystemProtection is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemprotection
	FileSystemProtection cfz.Expression[AWS_EFS_FileSystem_FileSystemProtection] `json:"FileSystemProtection,omitempty"`

	// FileSystemTags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemtags
	FileSystemTags cfz.ExpressionSlice[AWS_EFS_FileSystem_ElasticFileSystemTag] `json:"FileSystemTags,omitempty"`

	// KmsKeyId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-kmskeyid
	KmsKeyId cfz.Expression[string] `json:"KmsKeyId,omitempty"`

	// LifecyclePolicies is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-lifecyclepolicies
	LifecyclePolicies cfz.ExpressionSlice[AWS_EFS_FileSystem_LifecyclePolicy] `json:"LifecyclePolicies,omitempty"`

	// PerformanceMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-performancemode
	PerformanceMode cfz.Expression[string] `json:"PerformanceMode,omitempty"`

	// ProvisionedThroughputInMibps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-provisionedthroughputinmibps
	ProvisionedThroughputInMibps cfz.Expression[float64] `json:"ProvisionedThroughputInMibps,omitempty"`

	// ReplicationConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-replicationconfiguration
	ReplicationConfiguration cfz.Expression[AWS_EFS_FileSystem_ReplicationConfiguration] `json:"ReplicationConfiguration,omitempty"`

	// ThroughputMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-throughputmode
	ThroughputMode cfz.Expression[string] `json:"ThroughputMode,omitempty"`
}

// New__AWS_EFS_FileSystem initializes a new *AWS_EFS_FileSystem.
func New__AWS_EFS_FileSystem(logicalName string) *AWS_EFS_FileSystem {
	return &AWS_EFS_FileSystem{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EFS_FileSystem) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EFS_FileSystem) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EFS_FileSystem) GetType() string {
	return AWS_EFS_FileSystem__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EFS_FileSystem) Set__LogicalName(v string) *AWS_EFS_FileSystem {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EFS_FileSystem) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EFS_FileSystem {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EFS_FileSystem) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EFS_FileSystem {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EFS_FileSystem) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EFS_FileSystem {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EFS_FileSystem) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EFS_FileSystem {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EFS_FileSystem) Set__RequestedOutputs(v []cfz.Output) *AWS_EFS_FileSystem {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EFS_FileSystem) Add__RequestedOutputs(v ...cfz.Output) *AWS_EFS_FileSystem {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AvailabilityZoneName updates property "AvailabilityZoneName".
func (t *AWS_EFS_FileSystem) Set__AvailabilityZoneName(v cfz.Expression[string]) *AWS_EFS_FileSystem {
	t.AvailabilityZoneName = v
	return t
}

// SetV__AvailabilityZoneName updates property "AvailabilityZoneName".
func (t *AWS_EFS_FileSystem) SetV__AvailabilityZoneName(v string) *AWS_EFS_FileSystem {
	t.AvailabilityZoneName = cfz.V(v)
	return t
}

// Set__BackupPolicy updates property "BackupPolicy".
func (t *AWS_EFS_FileSystem) Set__BackupPolicy(v cfz.Expression[AWS_EFS_FileSystem_BackupPolicy]) *AWS_EFS_FileSystem {
	t.BackupPolicy = v
	return t
}

// SetV__BackupPolicy updates property "BackupPolicy".
func (t *AWS_EFS_FileSystem) SetV__BackupPolicy(v AWS_EFS_FileSystem_BackupPolicy) *AWS_EFS_FileSystem {
	t.BackupPolicy = cfz.V(v)
	return t
}

// Set__BypassPolicyLockoutSafetyCheck updates property "BypassPolicyLockoutSafetyCheck".
func (t *AWS_EFS_FileSystem) Set__BypassPolicyLockoutSafetyCheck(v cfz.Expression[bool]) *AWS_EFS_FileSystem {
	t.BypassPolicyLockoutSafetyCheck = v
	return t
}

// SetV__BypassPolicyLockoutSafetyCheck updates property "BypassPolicyLockoutSafetyCheck".
func (t *AWS_EFS_FileSystem) SetV__BypassPolicyLockoutSafetyCheck(v bool) *AWS_EFS_FileSystem {
	t.BypassPolicyLockoutSafetyCheck = cfz.V(v)
	return t
}

// Set__Encrypted updates property "Encrypted".
func (t *AWS_EFS_FileSystem) Set__Encrypted(v cfz.Expression[bool]) *AWS_EFS_FileSystem {
	t.Encrypted = v
	return t
}

// SetV__Encrypted updates property "Encrypted".
func (t *AWS_EFS_FileSystem) SetV__Encrypted(v bool) *AWS_EFS_FileSystem {
	t.Encrypted = cfz.V(v)
	return t
}

// Set__FileSystemPolicy updates property "FileSystemPolicy".
func (t *AWS_EFS_FileSystem) Set__FileSystemPolicy(v cfz.Expression[json.RawMessage]) *AWS_EFS_FileSystem {
	t.FileSystemPolicy = v
	return t
}

// SetV__FileSystemPolicy updates property "FileSystemPolicy".
func (t *AWS_EFS_FileSystem) SetV__FileSystemPolicy(v json.RawMessage) *AWS_EFS_FileSystem {
	t.FileSystemPolicy = cfz.V(v)
	return t
}

// Set__FileSystemProtection updates property "FileSystemProtection".
func (t *AWS_EFS_FileSystem) Set__FileSystemProtection(v cfz.Expression[AWS_EFS_FileSystem_FileSystemProtection]) *AWS_EFS_FileSystem {
	t.FileSystemProtection = v
	return t
}

// SetV__FileSystemProtection updates property "FileSystemProtection".
func (t *AWS_EFS_FileSystem) SetV__FileSystemProtection(v AWS_EFS_FileSystem_FileSystemProtection) *AWS_EFS_FileSystem {
	t.FileSystemProtection = cfz.V(v)
	return t
}

// Set__FileSystemTags updates property "FileSystemTags".
func (t *AWS_EFS_FileSystem) Set__FileSystemTags(v cfz.ExpressionSlice[AWS_EFS_FileSystem_ElasticFileSystemTag]) *AWS_EFS_FileSystem {
	t.FileSystemTags = v
	return t
}

// SetS__FileSystemTags updates property "FileSystemTags".
func (t *AWS_EFS_FileSystem) SetS__FileSystemTags(v ...cfz.Expression[AWS_EFS_FileSystem_ElasticFileSystemTag]) *AWS_EFS_FileSystem {
	t.FileSystemTags = cfz.S(v...)
	return t
}

// SetSV__FileSystemTags updates property "FileSystemTags".
func (t *AWS_EFS_FileSystem) SetSV__FileSystemTags(v ...AWS_EFS_FileSystem_ElasticFileSystemTag) *AWS_EFS_FileSystem {
	t.FileSystemTags = cfz.SV(v...)
	return t
}

// Set__KmsKeyId updates property "KmsKeyId".
func (t *AWS_EFS_FileSystem) Set__KmsKeyId(v cfz.Expression[string]) *AWS_EFS_FileSystem {
	t.KmsKeyId = v
	return t
}

// SetV__KmsKeyId updates property "KmsKeyId".
func (t *AWS_EFS_FileSystem) SetV__KmsKeyId(v string) *AWS_EFS_FileSystem {
	t.KmsKeyId = cfz.V(v)
	return t
}

// Set__LifecyclePolicies updates property "LifecyclePolicies".
func (t *AWS_EFS_FileSystem) Set__LifecyclePolicies(v cfz.ExpressionSlice[AWS_EFS_FileSystem_LifecyclePolicy]) *AWS_EFS_FileSystem {
	t.LifecyclePolicies = v
	return t
}

// SetS__LifecyclePolicies updates property "LifecyclePolicies".
func (t *AWS_EFS_FileSystem) SetS__LifecyclePolicies(v ...cfz.Expression[AWS_EFS_FileSystem_LifecyclePolicy]) *AWS_EFS_FileSystem {
	t.LifecyclePolicies = cfz.S(v...)
	return t
}

// SetSV__LifecyclePolicies updates property "LifecyclePolicies".
func (t *AWS_EFS_FileSystem) SetSV__LifecyclePolicies(v ...AWS_EFS_FileSystem_LifecyclePolicy) *AWS_EFS_FileSystem {
	t.LifecyclePolicies = cfz.SV(v...)
	return t
}

// Set__PerformanceMode updates property "PerformanceMode".
func (t *AWS_EFS_FileSystem) Set__PerformanceMode(v cfz.Expression[string]) *AWS_EFS_FileSystem {
	t.PerformanceMode = v
	return t
}

// SetV__PerformanceMode updates property "PerformanceMode".
func (t *AWS_EFS_FileSystem) SetV__PerformanceMode(v string) *AWS_EFS_FileSystem {
	t.PerformanceMode = cfz.V(v)
	return t
}

// Set__ProvisionedThroughputInMibps updates property "ProvisionedThroughputInMibps".
func (t *AWS_EFS_FileSystem) Set__ProvisionedThroughputInMibps(v cfz.Expression[float64]) *AWS_EFS_FileSystem {
	t.ProvisionedThroughputInMibps = v
	return t
}

// SetV__ProvisionedThroughputInMibps updates property "ProvisionedThroughputInMibps".
func (t *AWS_EFS_FileSystem) SetV__ProvisionedThroughputInMibps(v float64) *AWS_EFS_FileSystem {
	t.ProvisionedThroughputInMibps = cfz.V(v)
	return t
}

// Set__ReplicationConfiguration updates property "ReplicationConfiguration".
func (t *AWS_EFS_FileSystem) Set__ReplicationConfiguration(v cfz.Expression[AWS_EFS_FileSystem_ReplicationConfiguration]) *AWS_EFS_FileSystem {
	t.ReplicationConfiguration = v
	return t
}

// SetV__ReplicationConfiguration updates property "ReplicationConfiguration".
func (t *AWS_EFS_FileSystem) SetV__ReplicationConfiguration(v AWS_EFS_FileSystem_ReplicationConfiguration) *AWS_EFS_FileSystem {
	t.ReplicationConfiguration = cfz.V(v)
	return t
}

// Set__ThroughputMode updates property "ThroughputMode".
func (t *AWS_EFS_FileSystem) Set__ThroughputMode(v cfz.Expression[string]) *AWS_EFS_FileSystem {
	t.ThroughputMode = v
	return t
}

// SetV__ThroughputMode updates property "ThroughputMode".
func (t *AWS_EFS_FileSystem) SetV__ThroughputMode(v string) *AWS_EFS_FileSystem {
	t.ThroughputMode = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EFS_FileSystem) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_EFS_FileSystem) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EFS_FileSystem__AttributesMap.Arn))
}

// GetAtt__FileSystemId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FileSystemId
func (t *AWS_EFS_FileSystem) GetAtt__FileSystemId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EFS_FileSystem__AttributesMap.FileSystemId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EFS_FileSystem) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_EFS_FileSystem) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FileSystemId returns a conventionally configured output for an attribute of this resource.
// Attribute: FileSystemId
func (t *AWS_EFS_FileSystem) GetConventionalOutputAtt__FileSystemId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFileSystemId", t.GetAtt__FileSystemId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EFS_FileSystem) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EFS_FileSystem

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

func (t *AWS_EFS_FileSystem) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
