// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_eks

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EKS_Nodegroup)(nil)
	_ cfz.Resource                   = (*AWS_EKS_Nodegroup)(nil)
)

const (
	// AWS_EKS_Nodegroup__Type is the CloudFormation type for AWS::EKS::Nodegroup.
	AWS_EKS_Nodegroup__Type = "AWS::EKS::Nodegroup"
)

var (
	// AWS_EKS_Nodegroup__AttributesMap reports all the CloudFormation attributes for AWS::EKS::Nodegroup.
	AWS_EKS_Nodegroup__AttributesMap = struct {
		Arn           string
		ClusterName   string
		Id            string
		NodegroupName string
	}{
		Arn:           "Arn",
		ClusterName:   "ClusterName",
		Id:            "Id",
		NodegroupName: "NodegroupName",
	}

	// AWS_EKS_Nodegroup__AttributesSlice reports all the CloudFormation attributes for AWS::EKS::Nodegroup.
	AWS_EKS_Nodegroup__AttributesSlice = []string{
		AWS_EKS_Nodegroup__AttributesMap.Arn,
		AWS_EKS_Nodegroup__AttributesMap.ClusterName,
		AWS_EKS_Nodegroup__AttributesMap.Id,
		AWS_EKS_Nodegroup__AttributesMap.NodegroupName,
	}
)

var (
	// AWS_EKS_Nodegroup__PropertiesMap reports all the CloudFormation properties for AWS::EKS::Nodegroup.
	AWS_EKS_Nodegroup__PropertiesMap = struct {
		AmiType            string
		CapacityType       string
		ClusterName        string
		DiskSize           string
		ForceUpdateEnabled string
		InstanceTypes      string
		Labels             string
		LaunchTemplate     string
		NodeRepairConfig   string
		NodeRole           string
		NodegroupName      string
		ReleaseVersion     string
		RemoteAccess       string
		ScalingConfig      string
		Subnets            string
		Tags               string
		Taints             string
		UpdateConfig       string
		Version            string
	}{
		AmiType:            "AmiType",
		CapacityType:       "CapacityType",
		ClusterName:        "ClusterName",
		DiskSize:           "DiskSize",
		ForceUpdateEnabled: "ForceUpdateEnabled",
		InstanceTypes:      "InstanceTypes",
		Labels:             "Labels",
		LaunchTemplate:     "LaunchTemplate",
		NodeRepairConfig:   "NodeRepairConfig",
		NodeRole:           "NodeRole",
		NodegroupName:      "NodegroupName",
		ReleaseVersion:     "ReleaseVersion",
		RemoteAccess:       "RemoteAccess",
		ScalingConfig:      "ScalingConfig",
		Subnets:            "Subnets",
		Tags:               "Tags",
		Taints:             "Taints",
		UpdateConfig:       "UpdateConfig",
		Version:            "Version",
	}

	// AWS_EKS_Nodegroup__PropertiesSlice reports all the CloudFormation properties for AWS::EKS::Nodegroup.
	AWS_EKS_Nodegroup__PropertiesSlice = []string{
		AWS_EKS_Nodegroup__PropertiesMap.AmiType,
		AWS_EKS_Nodegroup__PropertiesMap.CapacityType,
		AWS_EKS_Nodegroup__PropertiesMap.ClusterName,
		AWS_EKS_Nodegroup__PropertiesMap.DiskSize,
		AWS_EKS_Nodegroup__PropertiesMap.ForceUpdateEnabled,
		AWS_EKS_Nodegroup__PropertiesMap.InstanceTypes,
		AWS_EKS_Nodegroup__PropertiesMap.Labels,
		AWS_EKS_Nodegroup__PropertiesMap.LaunchTemplate,
		AWS_EKS_Nodegroup__PropertiesMap.NodeRepairConfig,
		AWS_EKS_Nodegroup__PropertiesMap.NodeRole,
		AWS_EKS_Nodegroup__PropertiesMap.NodegroupName,
		AWS_EKS_Nodegroup__PropertiesMap.ReleaseVersion,
		AWS_EKS_Nodegroup__PropertiesMap.RemoteAccess,
		AWS_EKS_Nodegroup__PropertiesMap.ScalingConfig,
		AWS_EKS_Nodegroup__PropertiesMap.Subnets,
		AWS_EKS_Nodegroup__PropertiesMap.Tags,
		AWS_EKS_Nodegroup__PropertiesMap.Taints,
		AWS_EKS_Nodegroup__PropertiesMap.UpdateConfig,
		AWS_EKS_Nodegroup__PropertiesMap.Version,
	}
)

// AWS_EKS_Nodegroup is a binding for AWS::EKS::Nodegroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html
type AWS_EKS_Nodegroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AmiType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-amitype
	AmiType cfz.Expression[string] `json:"AmiType,omitempty"`

	// CapacityType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-capacitytype
	CapacityType cfz.Expression[string] `json:"CapacityType,omitempty"`

	// ClusterName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-clustername
	ClusterName cfz.Expression[string] `json:"ClusterName,omitempty"`

	// DiskSize is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-disksize
	DiskSize cfz.Expression[int32] `json:"DiskSize,omitempty"`

	// ForceUpdateEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-forceupdateenabled
	ForceUpdateEnabled cfz.Expression[bool] `json:"ForceUpdateEnabled,omitempty"`

	// InstanceTypes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-instancetypes
	InstanceTypes cfz.ExpressionSlice[string] `json:"InstanceTypes,omitempty"`

	// Labels is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-labels
	Labels cfz.ExpressionMap[string] `json:"Labels,omitempty"`

	// LaunchTemplate is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-launchtemplate
	LaunchTemplate cfz.Expression[AWS_EKS_Nodegroup_LaunchTemplateSpecification] `json:"LaunchTemplate,omitempty"`

	// NodeRepairConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-noderepairconfig
	NodeRepairConfig cfz.Expression[AWS_EKS_Nodegroup_NodeRepairConfig] `json:"NodeRepairConfig,omitempty"`

	// NodeRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-noderole
	NodeRole cfz.Expression[string] `json:"NodeRole,omitempty"`

	// NodegroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-nodegroupname
	NodegroupName cfz.Expression[string] `json:"NodegroupName,omitempty"`

	// ReleaseVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-releaseversion
	ReleaseVersion cfz.Expression[string] `json:"ReleaseVersion,omitempty"`

	// RemoteAccess is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-remoteaccess
	RemoteAccess cfz.Expression[AWS_EKS_Nodegroup_RemoteAccess] `json:"RemoteAccess,omitempty"`

	// ScalingConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-scalingconfig
	ScalingConfig cfz.Expression[AWS_EKS_Nodegroup_ScalingConfig] `json:"ScalingConfig,omitempty"`

	// Subnets is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-subnets
	Subnets cfz.ExpressionSlice[string] `json:"Subnets,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`

	// Taints is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-taints
	Taints cfz.ExpressionSlice[AWS_EKS_Nodegroup_Taint] `json:"Taints,omitempty"`

	// UpdateConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-updateconfig
	UpdateConfig cfz.Expression[AWS_EKS_Nodegroup_UpdateConfig] `json:"UpdateConfig,omitempty"`

	// Version is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-nodegroup.html#cfn-eks-nodegroup-version
	Version cfz.Expression[string] `json:"Version,omitempty"`
}

// New__AWS_EKS_Nodegroup initializes a new *AWS_EKS_Nodegroup.
func New__AWS_EKS_Nodegroup(logicalName string) *AWS_EKS_Nodegroup {
	return &AWS_EKS_Nodegroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EKS_Nodegroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EKS_Nodegroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EKS_Nodegroup) GetType() string {
	return AWS_EKS_Nodegroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EKS_Nodegroup) Set__LogicalName(v string) *AWS_EKS_Nodegroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EKS_Nodegroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EKS_Nodegroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EKS_Nodegroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EKS_Nodegroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EKS_Nodegroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EKS_Nodegroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EKS_Nodegroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EKS_Nodegroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EKS_Nodegroup) Set__RequestedOutputs(v []cfz.Output) *AWS_EKS_Nodegroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EKS_Nodegroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_EKS_Nodegroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AmiType updates property "AmiType".
func (t *AWS_EKS_Nodegroup) Set__AmiType(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.AmiType = v
	return t
}

// SetV__AmiType updates property "AmiType".
func (t *AWS_EKS_Nodegroup) SetV__AmiType(v string) *AWS_EKS_Nodegroup {
	t.AmiType = cfz.V(v)
	return t
}

// Set__CapacityType updates property "CapacityType".
func (t *AWS_EKS_Nodegroup) Set__CapacityType(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.CapacityType = v
	return t
}

// SetV__CapacityType updates property "CapacityType".
func (t *AWS_EKS_Nodegroup) SetV__CapacityType(v string) *AWS_EKS_Nodegroup {
	t.CapacityType = cfz.V(v)
	return t
}

// Set__ClusterName updates property "ClusterName".
func (t *AWS_EKS_Nodegroup) Set__ClusterName(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.ClusterName = v
	return t
}

// SetV__ClusterName updates property "ClusterName".
func (t *AWS_EKS_Nodegroup) SetV__ClusterName(v string) *AWS_EKS_Nodegroup {
	t.ClusterName = cfz.V(v)
	return t
}

// Set__DiskSize updates property "DiskSize".
func (t *AWS_EKS_Nodegroup) Set__DiskSize(v cfz.Expression[int32]) *AWS_EKS_Nodegroup {
	t.DiskSize = v
	return t
}

// SetV__DiskSize updates property "DiskSize".
func (t *AWS_EKS_Nodegroup) SetV__DiskSize(v int32) *AWS_EKS_Nodegroup {
	t.DiskSize = cfz.V(v)
	return t
}

// Set__ForceUpdateEnabled updates property "ForceUpdateEnabled".
func (t *AWS_EKS_Nodegroup) Set__ForceUpdateEnabled(v cfz.Expression[bool]) *AWS_EKS_Nodegroup {
	t.ForceUpdateEnabled = v
	return t
}

// SetV__ForceUpdateEnabled updates property "ForceUpdateEnabled".
func (t *AWS_EKS_Nodegroup) SetV__ForceUpdateEnabled(v bool) *AWS_EKS_Nodegroup {
	t.ForceUpdateEnabled = cfz.V(v)
	return t
}

// Set__InstanceTypes updates property "InstanceTypes".
func (t *AWS_EKS_Nodegroup) Set__InstanceTypes(v cfz.ExpressionSlice[string]) *AWS_EKS_Nodegroup {
	t.InstanceTypes = v
	return t
}

// SetS__InstanceTypes updates property "InstanceTypes".
func (t *AWS_EKS_Nodegroup) SetS__InstanceTypes(v ...cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.InstanceTypes = cfz.S(v...)
	return t
}

// SetSV__InstanceTypes updates property "InstanceTypes".
func (t *AWS_EKS_Nodegroup) SetSV__InstanceTypes(v ...string) *AWS_EKS_Nodegroup {
	t.InstanceTypes = cfz.SV(v...)
	return t
}

// Set__Labels updates property "Labels".
func (t *AWS_EKS_Nodegroup) Set__Labels(v cfz.ExpressionMap[string]) *AWS_EKS_Nodegroup {
	t.Labels = v
	return t
}

// SetM__Labels updates property "Labels".
func (t *AWS_EKS_Nodegroup) SetM__Labels(v ...map[string]cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.Labels = cfz.M(v...)
	return t
}

// SetMV__Labels updates property "Labels".
func (t *AWS_EKS_Nodegroup) SetMV__Labels(v ...map[string]string) *AWS_EKS_Nodegroup {
	t.Labels = cfz.MV(v...)
	return t
}

// Set__LaunchTemplate updates property "LaunchTemplate".
func (t *AWS_EKS_Nodegroup) Set__LaunchTemplate(v cfz.Expression[AWS_EKS_Nodegroup_LaunchTemplateSpecification]) *AWS_EKS_Nodegroup {
	t.LaunchTemplate = v
	return t
}

// SetV__LaunchTemplate updates property "LaunchTemplate".
func (t *AWS_EKS_Nodegroup) SetV__LaunchTemplate(v AWS_EKS_Nodegroup_LaunchTemplateSpecification) *AWS_EKS_Nodegroup {
	t.LaunchTemplate = cfz.V(v)
	return t
}

// Set__NodeRepairConfig updates property "NodeRepairConfig".
func (t *AWS_EKS_Nodegroup) Set__NodeRepairConfig(v cfz.Expression[AWS_EKS_Nodegroup_NodeRepairConfig]) *AWS_EKS_Nodegroup {
	t.NodeRepairConfig = v
	return t
}

// SetV__NodeRepairConfig updates property "NodeRepairConfig".
func (t *AWS_EKS_Nodegroup) SetV__NodeRepairConfig(v AWS_EKS_Nodegroup_NodeRepairConfig) *AWS_EKS_Nodegroup {
	t.NodeRepairConfig = cfz.V(v)
	return t
}

// Set__NodeRole updates property "NodeRole".
func (t *AWS_EKS_Nodegroup) Set__NodeRole(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.NodeRole = v
	return t
}

// SetV__NodeRole updates property "NodeRole".
func (t *AWS_EKS_Nodegroup) SetV__NodeRole(v string) *AWS_EKS_Nodegroup {
	t.NodeRole = cfz.V(v)
	return t
}

// Set__NodegroupName updates property "NodegroupName".
func (t *AWS_EKS_Nodegroup) Set__NodegroupName(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.NodegroupName = v
	return t
}

// SetV__NodegroupName updates property "NodegroupName".
func (t *AWS_EKS_Nodegroup) SetV__NodegroupName(v string) *AWS_EKS_Nodegroup {
	t.NodegroupName = cfz.V(v)
	return t
}

// Set__ReleaseVersion updates property "ReleaseVersion".
func (t *AWS_EKS_Nodegroup) Set__ReleaseVersion(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.ReleaseVersion = v
	return t
}

// SetV__ReleaseVersion updates property "ReleaseVersion".
func (t *AWS_EKS_Nodegroup) SetV__ReleaseVersion(v string) *AWS_EKS_Nodegroup {
	t.ReleaseVersion = cfz.V(v)
	return t
}

// Set__RemoteAccess updates property "RemoteAccess".
func (t *AWS_EKS_Nodegroup) Set__RemoteAccess(v cfz.Expression[AWS_EKS_Nodegroup_RemoteAccess]) *AWS_EKS_Nodegroup {
	t.RemoteAccess = v
	return t
}

// SetV__RemoteAccess updates property "RemoteAccess".
func (t *AWS_EKS_Nodegroup) SetV__RemoteAccess(v AWS_EKS_Nodegroup_RemoteAccess) *AWS_EKS_Nodegroup {
	t.RemoteAccess = cfz.V(v)
	return t
}

// Set__ScalingConfig updates property "ScalingConfig".
func (t *AWS_EKS_Nodegroup) Set__ScalingConfig(v cfz.Expression[AWS_EKS_Nodegroup_ScalingConfig]) *AWS_EKS_Nodegroup {
	t.ScalingConfig = v
	return t
}

// SetV__ScalingConfig updates property "ScalingConfig".
func (t *AWS_EKS_Nodegroup) SetV__ScalingConfig(v AWS_EKS_Nodegroup_ScalingConfig) *AWS_EKS_Nodegroup {
	t.ScalingConfig = cfz.V(v)
	return t
}

// Set__Subnets updates property "Subnets".
func (t *AWS_EKS_Nodegroup) Set__Subnets(v cfz.ExpressionSlice[string]) *AWS_EKS_Nodegroup {
	t.Subnets = v
	return t
}

// SetS__Subnets updates property "Subnets".
func (t *AWS_EKS_Nodegroup) SetS__Subnets(v ...cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.Subnets = cfz.S(v...)
	return t
}

// SetSV__Subnets updates property "Subnets".
func (t *AWS_EKS_Nodegroup) SetSV__Subnets(v ...string) *AWS_EKS_Nodegroup {
	t.Subnets = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EKS_Nodegroup) Set__Tags(v cfz.ExpressionMap[string]) *AWS_EKS_Nodegroup {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_EKS_Nodegroup) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_EKS_Nodegroup) SetMV__Tags(v ...map[string]string) *AWS_EKS_Nodegroup {
	t.Tags = cfz.MV(v...)
	return t
}

// Set__Taints updates property "Taints".
func (t *AWS_EKS_Nodegroup) Set__Taints(v cfz.ExpressionSlice[AWS_EKS_Nodegroup_Taint]) *AWS_EKS_Nodegroup {
	t.Taints = v
	return t
}

// SetS__Taints updates property "Taints".
func (t *AWS_EKS_Nodegroup) SetS__Taints(v ...cfz.Expression[AWS_EKS_Nodegroup_Taint]) *AWS_EKS_Nodegroup {
	t.Taints = cfz.S(v...)
	return t
}

// SetSV__Taints updates property "Taints".
func (t *AWS_EKS_Nodegroup) SetSV__Taints(v ...AWS_EKS_Nodegroup_Taint) *AWS_EKS_Nodegroup {
	t.Taints = cfz.SV(v...)
	return t
}

// Set__UpdateConfig updates property "UpdateConfig".
func (t *AWS_EKS_Nodegroup) Set__UpdateConfig(v cfz.Expression[AWS_EKS_Nodegroup_UpdateConfig]) *AWS_EKS_Nodegroup {
	t.UpdateConfig = v
	return t
}

// SetV__UpdateConfig updates property "UpdateConfig".
func (t *AWS_EKS_Nodegroup) SetV__UpdateConfig(v AWS_EKS_Nodegroup_UpdateConfig) *AWS_EKS_Nodegroup {
	t.UpdateConfig = cfz.V(v)
	return t
}

// Set__Version updates property "Version".
func (t *AWS_EKS_Nodegroup) Set__Version(v cfz.Expression[string]) *AWS_EKS_Nodegroup {
	t.Version = v
	return t
}

// SetV__Version updates property "Version".
func (t *AWS_EKS_Nodegroup) SetV__Version(v string) *AWS_EKS_Nodegroup {
	t.Version = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EKS_Nodegroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_EKS_Nodegroup) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EKS_Nodegroup__AttributesMap.Arn))
}

// GetAtt__ClusterName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ClusterName
func (t *AWS_EKS_Nodegroup) GetAtt__ClusterName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EKS_Nodegroup__AttributesMap.ClusterName))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_EKS_Nodegroup) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EKS_Nodegroup__AttributesMap.Id))
}

// GetAtt__NodegroupName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: NodegroupName
func (t *AWS_EKS_Nodegroup) GetAtt__NodegroupName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EKS_Nodegroup__AttributesMap.NodegroupName))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EKS_Nodegroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_EKS_Nodegroup) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ClusterName returns a conventionally configured output for an attribute of this resource.
// Attribute: ClusterName
func (t *AWS_EKS_Nodegroup) GetConventionalOutputAtt__ClusterName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttClusterName", t.GetAtt__ClusterName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_EKS_Nodegroup) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__NodegroupName returns a conventionally configured output for an attribute of this resource.
// Attribute: NodegroupName
func (t *AWS_EKS_Nodegroup) GetConventionalOutputAtt__NodegroupName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttNodegroupName", t.GetAtt__NodegroupName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EKS_Nodegroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EKS_Nodegroup

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

func (t *AWS_EKS_Nodegroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
