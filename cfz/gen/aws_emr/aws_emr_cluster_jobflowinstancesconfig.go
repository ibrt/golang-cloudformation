// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_emr

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EMR_Cluster_JobFlowInstancesConfig__Type is the CloudFormation type for AWS::EMR::Cluster.JobFlowInstancesConfig.
	AWS_EMR_Cluster_JobFlowInstancesConfig__Type = "AWS::EMR::Cluster.JobFlowInstancesConfig"
)

var (
	// AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap reports all the CloudFormation properties for AWS::EMR::Cluster.JobFlowInstancesConfig.
	AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap = struct {
		AdditionalMasterSecurityGroups string
		AdditionalSlaveSecurityGroups  string
		CoreInstanceFleet              string
		CoreInstanceGroup              string
		Ec2KeyName                     string
		Ec2SubnetId                    string
		Ec2SubnetIds                   string
		EmrManagedMasterSecurityGroup  string
		EmrManagedSlaveSecurityGroup   string
		HadoopVersion                  string
		KeepJobFlowAliveWhenNoSteps    string
		MasterInstanceFleet            string
		MasterInstanceGroup            string
		Placement                      string
		ServiceAccessSecurityGroup     string
		TaskInstanceFleets             string
		TaskInstanceGroups             string
		TerminationProtected           string
		UnhealthyNodeReplacement       string
	}{
		AdditionalMasterSecurityGroups: "AdditionalMasterSecurityGroups",
		AdditionalSlaveSecurityGroups:  "AdditionalSlaveSecurityGroups",
		CoreInstanceFleet:              "CoreInstanceFleet",
		CoreInstanceGroup:              "CoreInstanceGroup",
		Ec2KeyName:                     "Ec2KeyName",
		Ec2SubnetId:                    "Ec2SubnetId",
		Ec2SubnetIds:                   "Ec2SubnetIds",
		EmrManagedMasterSecurityGroup:  "EmrManagedMasterSecurityGroup",
		EmrManagedSlaveSecurityGroup:   "EmrManagedSlaveSecurityGroup",
		HadoopVersion:                  "HadoopVersion",
		KeepJobFlowAliveWhenNoSteps:    "KeepJobFlowAliveWhenNoSteps",
		MasterInstanceFleet:            "MasterInstanceFleet",
		MasterInstanceGroup:            "MasterInstanceGroup",
		Placement:                      "Placement",
		ServiceAccessSecurityGroup:     "ServiceAccessSecurityGroup",
		TaskInstanceFleets:             "TaskInstanceFleets",
		TaskInstanceGroups:             "TaskInstanceGroups",
		TerminationProtected:           "TerminationProtected",
		UnhealthyNodeReplacement:       "UnhealthyNodeReplacement",
	}

	// AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesSlice reports all the CloudFormation properties for AWS::EMR::Cluster.JobFlowInstancesConfig.
	AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesSlice = []string{
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.AdditionalMasterSecurityGroups,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.AdditionalSlaveSecurityGroups,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.CoreInstanceFleet,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.CoreInstanceGroup,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.Ec2KeyName,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.Ec2SubnetId,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.Ec2SubnetIds,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.EmrManagedMasterSecurityGroup,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.EmrManagedSlaveSecurityGroup,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.HadoopVersion,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.KeepJobFlowAliveWhenNoSteps,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.MasterInstanceFleet,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.MasterInstanceGroup,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.Placement,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.ServiceAccessSecurityGroup,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.TaskInstanceFleets,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.TaskInstanceGroups,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.TerminationProtected,
		AWS_EMR_Cluster_JobFlowInstancesConfig__PropertiesMap.UnhealthyNodeReplacement,
	}
)

// AWS_EMR_Cluster_JobFlowInstancesConfig is a binding for AWS::EMR::Cluster.JobFlowInstancesConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html
type AWS_EMR_Cluster_JobFlowInstancesConfig struct {
	// AdditionalMasterSecurityGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-additionalmastersecuritygroups
	AdditionalMasterSecurityGroups cfz.ExpressionSlice[string] `json:"AdditionalMasterSecurityGroups,omitempty"`

	// AdditionalSlaveSecurityGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-additionalslavesecuritygroups
	AdditionalSlaveSecurityGroups cfz.ExpressionSlice[string] `json:"AdditionalSlaveSecurityGroups,omitempty"`

	// CoreInstanceFleet is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-coreinstancefleet
	CoreInstanceFleet cfz.Expression[AWS_EMR_Cluster_InstanceFleetConfig] `json:"CoreInstanceFleet,omitempty"`

	// CoreInstanceGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-coreinstancegroup
	CoreInstanceGroup cfz.Expression[AWS_EMR_Cluster_InstanceGroupConfig] `json:"CoreInstanceGroup,omitempty"`

	// Ec2KeyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-ec2keyname
	Ec2KeyName cfz.Expression[string] `json:"Ec2KeyName,omitempty"`

	// Ec2SubnetId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-ec2subnetid
	Ec2SubnetId cfz.Expression[string] `json:"Ec2SubnetId,omitempty"`

	// Ec2SubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-ec2subnetids
	Ec2SubnetIds cfz.ExpressionSlice[string] `json:"Ec2SubnetIds,omitempty"`

	// EmrManagedMasterSecurityGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-emrmanagedmastersecuritygroup
	EmrManagedMasterSecurityGroup cfz.Expression[string] `json:"EmrManagedMasterSecurityGroup,omitempty"`

	// EmrManagedSlaveSecurityGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-emrmanagedslavesecuritygroup
	EmrManagedSlaveSecurityGroup cfz.Expression[string] `json:"EmrManagedSlaveSecurityGroup,omitempty"`

	// HadoopVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-hadoopversion
	HadoopVersion cfz.Expression[string] `json:"HadoopVersion,omitempty"`

	// KeepJobFlowAliveWhenNoSteps is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-keepjobflowalivewhennosteps
	KeepJobFlowAliveWhenNoSteps cfz.Expression[bool] `json:"KeepJobFlowAliveWhenNoSteps,omitempty"`

	// MasterInstanceFleet is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-masterinstancefleet
	MasterInstanceFleet cfz.Expression[AWS_EMR_Cluster_InstanceFleetConfig] `json:"MasterInstanceFleet,omitempty"`

	// MasterInstanceGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-masterinstancegroup
	MasterInstanceGroup cfz.Expression[AWS_EMR_Cluster_InstanceGroupConfig] `json:"MasterInstanceGroup,omitempty"`

	// Placement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-placement
	Placement cfz.Expression[AWS_EMR_Cluster_PlacementType] `json:"Placement,omitempty"`

	// ServiceAccessSecurityGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-serviceaccesssecuritygroup
	ServiceAccessSecurityGroup cfz.Expression[string] `json:"ServiceAccessSecurityGroup,omitempty"`

	// TaskInstanceFleets is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-taskinstancefleets
	TaskInstanceFleets cfz.ExpressionSlice[AWS_EMR_Cluster_InstanceFleetConfig] `json:"TaskInstanceFleets,omitempty"`

	// TaskInstanceGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-taskinstancegroups
	TaskInstanceGroups cfz.ExpressionSlice[AWS_EMR_Cluster_InstanceGroupConfig] `json:"TaskInstanceGroups,omitempty"`

	// TerminationProtected is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-terminationprotected
	TerminationProtected cfz.Expression[bool] `json:"TerminationProtected,omitempty"`

	// UnhealthyNodeReplacement is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-unhealthynodereplacement
	UnhealthyNodeReplacement cfz.Expression[bool] `json:"UnhealthyNodeReplacement,omitempty"`
}

// New__AWS_EMR_Cluster_JobFlowInstancesConfig initializes a new AWS_EMR_Cluster_JobFlowInstancesConfig.
func New__AWS_EMR_Cluster_JobFlowInstancesConfig() AWS_EMR_Cluster_JobFlowInstancesConfig {
	return AWS_EMR_Cluster_JobFlowInstancesConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_EMR_Cluster_JobFlowInstancesConfig) GetType() string {
	return AWS_EMR_Cluster_JobFlowInstancesConfig__Type
}

// Set__AdditionalMasterSecurityGroups updates property "AdditionalMasterSecurityGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__AdditionalMasterSecurityGroups(v cfz.ExpressionSlice[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.AdditionalMasterSecurityGroups = v
	return t
}

// SetS__AdditionalMasterSecurityGroups updates property "AdditionalMasterSecurityGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetS__AdditionalMasterSecurityGroups(v ...cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.AdditionalMasterSecurityGroups = cfz.S(v...)
	return t
}

// SetSV__AdditionalMasterSecurityGroups updates property "AdditionalMasterSecurityGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetSV__AdditionalMasterSecurityGroups(v ...string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.AdditionalMasterSecurityGroups = cfz.SV(v...)
	return t
}

// Set__AdditionalSlaveSecurityGroups updates property "AdditionalSlaveSecurityGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__AdditionalSlaveSecurityGroups(v cfz.ExpressionSlice[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.AdditionalSlaveSecurityGroups = v
	return t
}

// SetS__AdditionalSlaveSecurityGroups updates property "AdditionalSlaveSecurityGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetS__AdditionalSlaveSecurityGroups(v ...cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.AdditionalSlaveSecurityGroups = cfz.S(v...)
	return t
}

// SetSV__AdditionalSlaveSecurityGroups updates property "AdditionalSlaveSecurityGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetSV__AdditionalSlaveSecurityGroups(v ...string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.AdditionalSlaveSecurityGroups = cfz.SV(v...)
	return t
}

// Set__CoreInstanceFleet updates property "CoreInstanceFleet".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__CoreInstanceFleet(v cfz.Expression[AWS_EMR_Cluster_InstanceFleetConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.CoreInstanceFleet = v
	return t
}

// SetV__CoreInstanceFleet updates property "CoreInstanceFleet".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__CoreInstanceFleet(v AWS_EMR_Cluster_InstanceFleetConfig) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.CoreInstanceFleet = cfz.V(v)
	return t
}

// Set__CoreInstanceGroup updates property "CoreInstanceGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__CoreInstanceGroup(v cfz.Expression[AWS_EMR_Cluster_InstanceGroupConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.CoreInstanceGroup = v
	return t
}

// SetV__CoreInstanceGroup updates property "CoreInstanceGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__CoreInstanceGroup(v AWS_EMR_Cluster_InstanceGroupConfig) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.CoreInstanceGroup = cfz.V(v)
	return t
}

// Set__Ec2KeyName updates property "Ec2KeyName".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__Ec2KeyName(v cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2KeyName = v
	return t
}

// SetV__Ec2KeyName updates property "Ec2KeyName".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__Ec2KeyName(v string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2KeyName = cfz.V(v)
	return t
}

// Set__Ec2SubnetId updates property "Ec2SubnetId".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__Ec2SubnetId(v cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2SubnetId = v
	return t
}

// SetV__Ec2SubnetId updates property "Ec2SubnetId".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__Ec2SubnetId(v string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2SubnetId = cfz.V(v)
	return t
}

// Set__Ec2SubnetIds updates property "Ec2SubnetIds".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__Ec2SubnetIds(v cfz.ExpressionSlice[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2SubnetIds = v
	return t
}

// SetS__Ec2SubnetIds updates property "Ec2SubnetIds".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetS__Ec2SubnetIds(v ...cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2SubnetIds = cfz.S(v...)
	return t
}

// SetSV__Ec2SubnetIds updates property "Ec2SubnetIds".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetSV__Ec2SubnetIds(v ...string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Ec2SubnetIds = cfz.SV(v...)
	return t
}

// Set__EmrManagedMasterSecurityGroup updates property "EmrManagedMasterSecurityGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__EmrManagedMasterSecurityGroup(v cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.EmrManagedMasterSecurityGroup = v
	return t
}

// SetV__EmrManagedMasterSecurityGroup updates property "EmrManagedMasterSecurityGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__EmrManagedMasterSecurityGroup(v string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.EmrManagedMasterSecurityGroup = cfz.V(v)
	return t
}

// Set__EmrManagedSlaveSecurityGroup updates property "EmrManagedSlaveSecurityGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__EmrManagedSlaveSecurityGroup(v cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.EmrManagedSlaveSecurityGroup = v
	return t
}

// SetV__EmrManagedSlaveSecurityGroup updates property "EmrManagedSlaveSecurityGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__EmrManagedSlaveSecurityGroup(v string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.EmrManagedSlaveSecurityGroup = cfz.V(v)
	return t
}

// Set__HadoopVersion updates property "HadoopVersion".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__HadoopVersion(v cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.HadoopVersion = v
	return t
}

// SetV__HadoopVersion updates property "HadoopVersion".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__HadoopVersion(v string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.HadoopVersion = cfz.V(v)
	return t
}

// Set__KeepJobFlowAliveWhenNoSteps updates property "KeepJobFlowAliveWhenNoSteps".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__KeepJobFlowAliveWhenNoSteps(v cfz.Expression[bool]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.KeepJobFlowAliveWhenNoSteps = v
	return t
}

// SetV__KeepJobFlowAliveWhenNoSteps updates property "KeepJobFlowAliveWhenNoSteps".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__KeepJobFlowAliveWhenNoSteps(v bool) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.KeepJobFlowAliveWhenNoSteps = cfz.V(v)
	return t
}

// Set__MasterInstanceFleet updates property "MasterInstanceFleet".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__MasterInstanceFleet(v cfz.Expression[AWS_EMR_Cluster_InstanceFleetConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.MasterInstanceFleet = v
	return t
}

// SetV__MasterInstanceFleet updates property "MasterInstanceFleet".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__MasterInstanceFleet(v AWS_EMR_Cluster_InstanceFleetConfig) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.MasterInstanceFleet = cfz.V(v)
	return t
}

// Set__MasterInstanceGroup updates property "MasterInstanceGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__MasterInstanceGroup(v cfz.Expression[AWS_EMR_Cluster_InstanceGroupConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.MasterInstanceGroup = v
	return t
}

// SetV__MasterInstanceGroup updates property "MasterInstanceGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__MasterInstanceGroup(v AWS_EMR_Cluster_InstanceGroupConfig) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.MasterInstanceGroup = cfz.V(v)
	return t
}

// Set__Placement updates property "Placement".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__Placement(v cfz.Expression[AWS_EMR_Cluster_PlacementType]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Placement = v
	return t
}

// SetV__Placement updates property "Placement".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__Placement(v AWS_EMR_Cluster_PlacementType) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.Placement = cfz.V(v)
	return t
}

// Set__ServiceAccessSecurityGroup updates property "ServiceAccessSecurityGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__ServiceAccessSecurityGroup(v cfz.Expression[string]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.ServiceAccessSecurityGroup = v
	return t
}

// SetV__ServiceAccessSecurityGroup updates property "ServiceAccessSecurityGroup".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__ServiceAccessSecurityGroup(v string) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.ServiceAccessSecurityGroup = cfz.V(v)
	return t
}

// Set__TaskInstanceFleets updates property "TaskInstanceFleets".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__TaskInstanceFleets(v cfz.ExpressionSlice[AWS_EMR_Cluster_InstanceFleetConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TaskInstanceFleets = v
	return t
}

// SetS__TaskInstanceFleets updates property "TaskInstanceFleets".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetS__TaskInstanceFleets(v ...cfz.Expression[AWS_EMR_Cluster_InstanceFleetConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TaskInstanceFleets = cfz.S(v...)
	return t
}

// SetSV__TaskInstanceFleets updates property "TaskInstanceFleets".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetSV__TaskInstanceFleets(v ...AWS_EMR_Cluster_InstanceFleetConfig) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TaskInstanceFleets = cfz.SV(v...)
	return t
}

// Set__TaskInstanceGroups updates property "TaskInstanceGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__TaskInstanceGroups(v cfz.ExpressionSlice[AWS_EMR_Cluster_InstanceGroupConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TaskInstanceGroups = v
	return t
}

// SetS__TaskInstanceGroups updates property "TaskInstanceGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetS__TaskInstanceGroups(v ...cfz.Expression[AWS_EMR_Cluster_InstanceGroupConfig]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TaskInstanceGroups = cfz.S(v...)
	return t
}

// SetSV__TaskInstanceGroups updates property "TaskInstanceGroups".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetSV__TaskInstanceGroups(v ...AWS_EMR_Cluster_InstanceGroupConfig) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TaskInstanceGroups = cfz.SV(v...)
	return t
}

// Set__TerminationProtected updates property "TerminationProtected".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__TerminationProtected(v cfz.Expression[bool]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TerminationProtected = v
	return t
}

// SetV__TerminationProtected updates property "TerminationProtected".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__TerminationProtected(v bool) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.TerminationProtected = cfz.V(v)
	return t
}

// Set__UnhealthyNodeReplacement updates property "UnhealthyNodeReplacement".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) Set__UnhealthyNodeReplacement(v cfz.Expression[bool]) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.UnhealthyNodeReplacement = v
	return t
}

// SetV__UnhealthyNodeReplacement updates property "UnhealthyNodeReplacement".
func (t AWS_EMR_Cluster_JobFlowInstancesConfig) SetV__UnhealthyNodeReplacement(v bool) AWS_EMR_Cluster_JobFlowInstancesConfig {
	t.UnhealthyNodeReplacement = cfz.V(v)
	return t
}
