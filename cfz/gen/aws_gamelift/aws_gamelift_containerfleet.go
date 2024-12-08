// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_gamelift

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_GameLift_ContainerFleet)(nil)
	_ cfz.Resource                   = (*AWS_GameLift_ContainerFleet)(nil)
)

const (
	// AWS_GameLift_ContainerFleet__Type is the CloudFormation type for AWS::GameLift::ContainerFleet.
	AWS_GameLift_ContainerFleet__Type = "AWS::GameLift::ContainerFleet"
)

var (
	// AWS_GameLift_ContainerFleet__AttributesMap reports all the CloudFormation attributes for AWS::GameLift::ContainerFleet.
	AWS_GameLift_ContainerFleet__AttributesMap = struct {
		CreationTime                                string
		DeploymentDetails                           string
		DeploymentDetails_LatestDeploymentId        string
		FleetArn                                    string
		FleetId                                     string
		GameServerContainerGroupDefinitionArn       string
		MaximumGameServerContainerGroupsPerInstance string
		PerInstanceContainerGroupDefinitionArn      string
		Status                                      string
	}{
		CreationTime:                          "CreationTime",
		DeploymentDetails:                     "DeploymentDetails",
		DeploymentDetails_LatestDeploymentId:  "DeploymentDetails.LatestDeploymentId",
		FleetArn:                              "FleetArn",
		FleetId:                               "FleetId",
		GameServerContainerGroupDefinitionArn: "GameServerContainerGroupDefinitionArn",
		MaximumGameServerContainerGroupsPerInstance: "MaximumGameServerContainerGroupsPerInstance",
		PerInstanceContainerGroupDefinitionArn:      "PerInstanceContainerGroupDefinitionArn",
		Status:                                      "Status",
	}

	// AWS_GameLift_ContainerFleet__AttributesSlice reports all the CloudFormation attributes for AWS::GameLift::ContainerFleet.
	AWS_GameLift_ContainerFleet__AttributesSlice = []string{
		AWS_GameLift_ContainerFleet__AttributesMap.CreationTime,
		AWS_GameLift_ContainerFleet__AttributesMap.DeploymentDetails,
		AWS_GameLift_ContainerFleet__AttributesMap.DeploymentDetails_LatestDeploymentId,
		AWS_GameLift_ContainerFleet__AttributesMap.FleetArn,
		AWS_GameLift_ContainerFleet__AttributesMap.FleetId,
		AWS_GameLift_ContainerFleet__AttributesMap.GameServerContainerGroupDefinitionArn,
		AWS_GameLift_ContainerFleet__AttributesMap.MaximumGameServerContainerGroupsPerInstance,
		AWS_GameLift_ContainerFleet__AttributesMap.PerInstanceContainerGroupDefinitionArn,
		AWS_GameLift_ContainerFleet__AttributesMap.Status,
	}
)

var (
	// AWS_GameLift_ContainerFleet__PropertiesMap reports all the CloudFormation properties for AWS::GameLift::ContainerFleet.
	AWS_GameLift_ContainerFleet__PropertiesMap = struct {
		BillingType                             string
		DeploymentConfiguration                 string
		Description                             string
		FleetRoleArn                            string
		GameServerContainerGroupDefinitionName  string
		GameServerContainerGroupsPerInstance    string
		GameSessionCreationLimitPolicy          string
		InstanceConnectionPortRange             string
		InstanceInboundPermissions              string
		InstanceType                            string
		Locations                               string
		LogConfiguration                        string
		MetricGroups                            string
		NewGameSessionProtectionPolicy          string
		PerInstanceContainerGroupDefinitionName string
		ScalingPolicies                         string
		Tags                                    string
	}{
		BillingType:                             "BillingType",
		DeploymentConfiguration:                 "DeploymentConfiguration",
		Description:                             "Description",
		FleetRoleArn:                            "FleetRoleArn",
		GameServerContainerGroupDefinitionName:  "GameServerContainerGroupDefinitionName",
		GameServerContainerGroupsPerInstance:    "GameServerContainerGroupsPerInstance",
		GameSessionCreationLimitPolicy:          "GameSessionCreationLimitPolicy",
		InstanceConnectionPortRange:             "InstanceConnectionPortRange",
		InstanceInboundPermissions:              "InstanceInboundPermissions",
		InstanceType:                            "InstanceType",
		Locations:                               "Locations",
		LogConfiguration:                        "LogConfiguration",
		MetricGroups:                            "MetricGroups",
		NewGameSessionProtectionPolicy:          "NewGameSessionProtectionPolicy",
		PerInstanceContainerGroupDefinitionName: "PerInstanceContainerGroupDefinitionName",
		ScalingPolicies:                         "ScalingPolicies",
		Tags:                                    "Tags",
	}

	// AWS_GameLift_ContainerFleet__PropertiesSlice reports all the CloudFormation properties for AWS::GameLift::ContainerFleet.
	AWS_GameLift_ContainerFleet__PropertiesSlice = []string{
		AWS_GameLift_ContainerFleet__PropertiesMap.BillingType,
		AWS_GameLift_ContainerFleet__PropertiesMap.DeploymentConfiguration,
		AWS_GameLift_ContainerFleet__PropertiesMap.Description,
		AWS_GameLift_ContainerFleet__PropertiesMap.FleetRoleArn,
		AWS_GameLift_ContainerFleet__PropertiesMap.GameServerContainerGroupDefinitionName,
		AWS_GameLift_ContainerFleet__PropertiesMap.GameServerContainerGroupsPerInstance,
		AWS_GameLift_ContainerFleet__PropertiesMap.GameSessionCreationLimitPolicy,
		AWS_GameLift_ContainerFleet__PropertiesMap.InstanceConnectionPortRange,
		AWS_GameLift_ContainerFleet__PropertiesMap.InstanceInboundPermissions,
		AWS_GameLift_ContainerFleet__PropertiesMap.InstanceType,
		AWS_GameLift_ContainerFleet__PropertiesMap.Locations,
		AWS_GameLift_ContainerFleet__PropertiesMap.LogConfiguration,
		AWS_GameLift_ContainerFleet__PropertiesMap.MetricGroups,
		AWS_GameLift_ContainerFleet__PropertiesMap.NewGameSessionProtectionPolicy,
		AWS_GameLift_ContainerFleet__PropertiesMap.PerInstanceContainerGroupDefinitionName,
		AWS_GameLift_ContainerFleet__PropertiesMap.ScalingPolicies,
		AWS_GameLift_ContainerFleet__PropertiesMap.Tags,
	}
)

// AWS_GameLift_ContainerFleet is a binding for AWS::GameLift::ContainerFleet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html
type AWS_GameLift_ContainerFleet struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// BillingType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-billingtype
	BillingType cfz.Expression[string] `json:"BillingType,omitempty"`

	// DeploymentConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-deploymentconfiguration
	DeploymentConfiguration cfz.Expression[AWS_GameLift_ContainerFleet_DeploymentConfiguration] `json:"DeploymentConfiguration,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// FleetRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-fleetrolearn
	FleetRoleArn cfz.Expression[string] `json:"FleetRoleArn,omitempty"`

	// GameServerContainerGroupDefinitionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gameservercontainergroupdefinitionname
	GameServerContainerGroupDefinitionName cfz.Expression[string] `json:"GameServerContainerGroupDefinitionName,omitempty"`

	// GameServerContainerGroupsPerInstance is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gameservercontainergroupsperinstance
	GameServerContainerGroupsPerInstance cfz.Expression[int32] `json:"GameServerContainerGroupsPerInstance,omitempty"`

	// GameSessionCreationLimitPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-gamesessioncreationlimitpolicy
	GameSessionCreationLimitPolicy cfz.Expression[AWS_GameLift_ContainerFleet_GameSessionCreationLimitPolicy] `json:"GameSessionCreationLimitPolicy,omitempty"`

	// InstanceConnectionPortRange is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instanceconnectionportrange
	InstanceConnectionPortRange cfz.Expression[AWS_GameLift_ContainerFleet_ConnectionPortRange] `json:"InstanceConnectionPortRange,omitempty"`

	// InstanceInboundPermissions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instanceinboundpermissions
	InstanceInboundPermissions cfz.ExpressionSlice[AWS_GameLift_ContainerFleet_IpPermission] `json:"InstanceInboundPermissions,omitempty"`

	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`

	// Locations is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-locations
	Locations cfz.ExpressionSlice[AWS_GameLift_ContainerFleet_LocationConfiguration] `json:"Locations,omitempty"`

	// LogConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-logconfiguration
	LogConfiguration cfz.Expression[AWS_GameLift_ContainerFleet_LogConfiguration] `json:"LogConfiguration,omitempty"`

	// MetricGroups is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-metricgroups
	MetricGroups cfz.ExpressionSlice[string] `json:"MetricGroups,omitempty"`

	// NewGameSessionProtectionPolicy is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-newgamesessionprotectionpolicy
	NewGameSessionProtectionPolicy cfz.Expression[string] `json:"NewGameSessionProtectionPolicy,omitempty"`

	// PerInstanceContainerGroupDefinitionName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-perinstancecontainergroupdefinitionname
	PerInstanceContainerGroupDefinitionName cfz.Expression[string] `json:"PerInstanceContainerGroupDefinitionName,omitempty"`

	// ScalingPolicies is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-scalingpolicies
	ScalingPolicies cfz.ExpressionSlice[AWS_GameLift_ContainerFleet_ScalingPolicy] `json:"ScalingPolicies,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html#cfn-gamelift-containerfleet-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_GameLift_ContainerFleet initializes a new *AWS_GameLift_ContainerFleet.
func New__AWS_GameLift_ContainerFleet(logicalName string) *AWS_GameLift_ContainerFleet {
	return &AWS_GameLift_ContainerFleet{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_GameLift_ContainerFleet) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_GameLift_ContainerFleet) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_GameLift_ContainerFleet) GetType() string {
	return AWS_GameLift_ContainerFleet__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_GameLift_ContainerFleet) Set__LogicalName(v string) *AWS_GameLift_ContainerFleet {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_GameLift_ContainerFleet) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_GameLift_ContainerFleet {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_GameLift_ContainerFleet) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_GameLift_ContainerFleet {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_GameLift_ContainerFleet) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_GameLift_ContainerFleet {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_GameLift_ContainerFleet) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_GameLift_ContainerFleet {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_GameLift_ContainerFleet) Set__RequestedOutputs(v []cfz.Output) *AWS_GameLift_ContainerFleet {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_GameLift_ContainerFleet) Add__RequestedOutputs(v ...cfz.Output) *AWS_GameLift_ContainerFleet {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__BillingType updates property "BillingType".
func (t *AWS_GameLift_ContainerFleet) Set__BillingType(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.BillingType = v
	return t
}

// SetV__BillingType updates property "BillingType".
func (t *AWS_GameLift_ContainerFleet) SetV__BillingType(v string) *AWS_GameLift_ContainerFleet {
	t.BillingType = cfz.V(v)
	return t
}

// Set__DeploymentConfiguration updates property "DeploymentConfiguration".
func (t *AWS_GameLift_ContainerFleet) Set__DeploymentConfiguration(v cfz.Expression[AWS_GameLift_ContainerFleet_DeploymentConfiguration]) *AWS_GameLift_ContainerFleet {
	t.DeploymentConfiguration = v
	return t
}

// SetV__DeploymentConfiguration updates property "DeploymentConfiguration".
func (t *AWS_GameLift_ContainerFleet) SetV__DeploymentConfiguration(v AWS_GameLift_ContainerFleet_DeploymentConfiguration) *AWS_GameLift_ContainerFleet {
	t.DeploymentConfiguration = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_GameLift_ContainerFleet) Set__Description(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_GameLift_ContainerFleet) SetV__Description(v string) *AWS_GameLift_ContainerFleet {
	t.Description = cfz.V(v)
	return t
}

// Set__FleetRoleArn updates property "FleetRoleArn".
func (t *AWS_GameLift_ContainerFleet) Set__FleetRoleArn(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.FleetRoleArn = v
	return t
}

// SetV__FleetRoleArn updates property "FleetRoleArn".
func (t *AWS_GameLift_ContainerFleet) SetV__FleetRoleArn(v string) *AWS_GameLift_ContainerFleet {
	t.FleetRoleArn = cfz.V(v)
	return t
}

// Set__GameServerContainerGroupDefinitionName updates property "GameServerContainerGroupDefinitionName".
func (t *AWS_GameLift_ContainerFleet) Set__GameServerContainerGroupDefinitionName(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.GameServerContainerGroupDefinitionName = v
	return t
}

// SetV__GameServerContainerGroupDefinitionName updates property "GameServerContainerGroupDefinitionName".
func (t *AWS_GameLift_ContainerFleet) SetV__GameServerContainerGroupDefinitionName(v string) *AWS_GameLift_ContainerFleet {
	t.GameServerContainerGroupDefinitionName = cfz.V(v)
	return t
}

// Set__GameServerContainerGroupsPerInstance updates property "GameServerContainerGroupsPerInstance".
func (t *AWS_GameLift_ContainerFleet) Set__GameServerContainerGroupsPerInstance(v cfz.Expression[int32]) *AWS_GameLift_ContainerFleet {
	t.GameServerContainerGroupsPerInstance = v
	return t
}

// SetV__GameServerContainerGroupsPerInstance updates property "GameServerContainerGroupsPerInstance".
func (t *AWS_GameLift_ContainerFleet) SetV__GameServerContainerGroupsPerInstance(v int32) *AWS_GameLift_ContainerFleet {
	t.GameServerContainerGroupsPerInstance = cfz.V(v)
	return t
}

// Set__GameSessionCreationLimitPolicy updates property "GameSessionCreationLimitPolicy".
func (t *AWS_GameLift_ContainerFleet) Set__GameSessionCreationLimitPolicy(v cfz.Expression[AWS_GameLift_ContainerFleet_GameSessionCreationLimitPolicy]) *AWS_GameLift_ContainerFleet {
	t.GameSessionCreationLimitPolicy = v
	return t
}

// SetV__GameSessionCreationLimitPolicy updates property "GameSessionCreationLimitPolicy".
func (t *AWS_GameLift_ContainerFleet) SetV__GameSessionCreationLimitPolicy(v AWS_GameLift_ContainerFleet_GameSessionCreationLimitPolicy) *AWS_GameLift_ContainerFleet {
	t.GameSessionCreationLimitPolicy = cfz.V(v)
	return t
}

// Set__InstanceConnectionPortRange updates property "InstanceConnectionPortRange".
func (t *AWS_GameLift_ContainerFleet) Set__InstanceConnectionPortRange(v cfz.Expression[AWS_GameLift_ContainerFleet_ConnectionPortRange]) *AWS_GameLift_ContainerFleet {
	t.InstanceConnectionPortRange = v
	return t
}

// SetV__InstanceConnectionPortRange updates property "InstanceConnectionPortRange".
func (t *AWS_GameLift_ContainerFleet) SetV__InstanceConnectionPortRange(v AWS_GameLift_ContainerFleet_ConnectionPortRange) *AWS_GameLift_ContainerFleet {
	t.InstanceConnectionPortRange = cfz.V(v)
	return t
}

// Set__InstanceInboundPermissions updates property "InstanceInboundPermissions".
func (t *AWS_GameLift_ContainerFleet) Set__InstanceInboundPermissions(v cfz.ExpressionSlice[AWS_GameLift_ContainerFleet_IpPermission]) *AWS_GameLift_ContainerFleet {
	t.InstanceInboundPermissions = v
	return t
}

// SetS__InstanceInboundPermissions updates property "InstanceInboundPermissions".
func (t *AWS_GameLift_ContainerFleet) SetS__InstanceInboundPermissions(v ...cfz.Expression[AWS_GameLift_ContainerFleet_IpPermission]) *AWS_GameLift_ContainerFleet {
	t.InstanceInboundPermissions = cfz.S(v...)
	return t
}

// SetSV__InstanceInboundPermissions updates property "InstanceInboundPermissions".
func (t *AWS_GameLift_ContainerFleet) SetSV__InstanceInboundPermissions(v ...AWS_GameLift_ContainerFleet_IpPermission) *AWS_GameLift_ContainerFleet {
	t.InstanceInboundPermissions = cfz.SV(v...)
	return t
}

// Set__InstanceType updates property "InstanceType".
func (t *AWS_GameLift_ContainerFleet) Set__InstanceType(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t *AWS_GameLift_ContainerFleet) SetV__InstanceType(v string) *AWS_GameLift_ContainerFleet {
	t.InstanceType = cfz.V(v)
	return t
}

// Set__Locations updates property "Locations".
func (t *AWS_GameLift_ContainerFleet) Set__Locations(v cfz.ExpressionSlice[AWS_GameLift_ContainerFleet_LocationConfiguration]) *AWS_GameLift_ContainerFleet {
	t.Locations = v
	return t
}

// SetS__Locations updates property "Locations".
func (t *AWS_GameLift_ContainerFleet) SetS__Locations(v ...cfz.Expression[AWS_GameLift_ContainerFleet_LocationConfiguration]) *AWS_GameLift_ContainerFleet {
	t.Locations = cfz.S(v...)
	return t
}

// SetSV__Locations updates property "Locations".
func (t *AWS_GameLift_ContainerFleet) SetSV__Locations(v ...AWS_GameLift_ContainerFleet_LocationConfiguration) *AWS_GameLift_ContainerFleet {
	t.Locations = cfz.SV(v...)
	return t
}

// Set__LogConfiguration updates property "LogConfiguration".
func (t *AWS_GameLift_ContainerFleet) Set__LogConfiguration(v cfz.Expression[AWS_GameLift_ContainerFleet_LogConfiguration]) *AWS_GameLift_ContainerFleet {
	t.LogConfiguration = v
	return t
}

// SetV__LogConfiguration updates property "LogConfiguration".
func (t *AWS_GameLift_ContainerFleet) SetV__LogConfiguration(v AWS_GameLift_ContainerFleet_LogConfiguration) *AWS_GameLift_ContainerFleet {
	t.LogConfiguration = cfz.V(v)
	return t
}

// Set__MetricGroups updates property "MetricGroups".
func (t *AWS_GameLift_ContainerFleet) Set__MetricGroups(v cfz.ExpressionSlice[string]) *AWS_GameLift_ContainerFleet {
	t.MetricGroups = v
	return t
}

// SetS__MetricGroups updates property "MetricGroups".
func (t *AWS_GameLift_ContainerFleet) SetS__MetricGroups(v ...cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.MetricGroups = cfz.S(v...)
	return t
}

// SetSV__MetricGroups updates property "MetricGroups".
func (t *AWS_GameLift_ContainerFleet) SetSV__MetricGroups(v ...string) *AWS_GameLift_ContainerFleet {
	t.MetricGroups = cfz.SV(v...)
	return t
}

// Set__NewGameSessionProtectionPolicy updates property "NewGameSessionProtectionPolicy".
func (t *AWS_GameLift_ContainerFleet) Set__NewGameSessionProtectionPolicy(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.NewGameSessionProtectionPolicy = v
	return t
}

// SetV__NewGameSessionProtectionPolicy updates property "NewGameSessionProtectionPolicy".
func (t *AWS_GameLift_ContainerFleet) SetV__NewGameSessionProtectionPolicy(v string) *AWS_GameLift_ContainerFleet {
	t.NewGameSessionProtectionPolicy = cfz.V(v)
	return t
}

// Set__PerInstanceContainerGroupDefinitionName updates property "PerInstanceContainerGroupDefinitionName".
func (t *AWS_GameLift_ContainerFleet) Set__PerInstanceContainerGroupDefinitionName(v cfz.Expression[string]) *AWS_GameLift_ContainerFleet {
	t.PerInstanceContainerGroupDefinitionName = v
	return t
}

// SetV__PerInstanceContainerGroupDefinitionName updates property "PerInstanceContainerGroupDefinitionName".
func (t *AWS_GameLift_ContainerFleet) SetV__PerInstanceContainerGroupDefinitionName(v string) *AWS_GameLift_ContainerFleet {
	t.PerInstanceContainerGroupDefinitionName = cfz.V(v)
	return t
}

// Set__ScalingPolicies updates property "ScalingPolicies".
func (t *AWS_GameLift_ContainerFleet) Set__ScalingPolicies(v cfz.ExpressionSlice[AWS_GameLift_ContainerFleet_ScalingPolicy]) *AWS_GameLift_ContainerFleet {
	t.ScalingPolicies = v
	return t
}

// SetS__ScalingPolicies updates property "ScalingPolicies".
func (t *AWS_GameLift_ContainerFleet) SetS__ScalingPolicies(v ...cfz.Expression[AWS_GameLift_ContainerFleet_ScalingPolicy]) *AWS_GameLift_ContainerFleet {
	t.ScalingPolicies = cfz.S(v...)
	return t
}

// SetSV__ScalingPolicies updates property "ScalingPolicies".
func (t *AWS_GameLift_ContainerFleet) SetSV__ScalingPolicies(v ...AWS_GameLift_ContainerFleet_ScalingPolicy) *AWS_GameLift_ContainerFleet {
	t.ScalingPolicies = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_GameLift_ContainerFleet) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_GameLift_ContainerFleet {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_GameLift_ContainerFleet) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_GameLift_ContainerFleet {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_GameLift_ContainerFleet) SetSV__Tags(v ...cfz.Tag) *AWS_GameLift_ContainerFleet {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_GameLift_ContainerFleet) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreationTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTime
func (t *AWS_GameLift_ContainerFleet) GetAtt__CreationTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.CreationTime))
}

// GetAtt__DeploymentDetails returns a $cfz.Expression[AWS_GameLift_ContainerFleet_DeploymentDetails] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DeploymentDetails
func (t *AWS_GameLift_ContainerFleet) GetAtt__DeploymentDetails() cfz.Expression[AWS_GameLift_ContainerFleet_DeploymentDetails] {
	return cfz.GetAtt[AWS_GameLift_ContainerFleet_DeploymentDetails](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.DeploymentDetails))
}

// GetAtt__DeploymentDetails_LatestDeploymentId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DeploymentDetails.LatestDeploymentId
func (t *AWS_GameLift_ContainerFleet) GetAtt__DeploymentDetails_LatestDeploymentId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.DeploymentDetails_LatestDeploymentId))
}

// GetAtt__FleetArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FleetArn
func (t *AWS_GameLift_ContainerFleet) GetAtt__FleetArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.FleetArn))
}

// GetAtt__FleetId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: FleetId
func (t *AWS_GameLift_ContainerFleet) GetAtt__FleetId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.FleetId))
}

// GetAtt__GameServerContainerGroupDefinitionArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: GameServerContainerGroupDefinitionArn
func (t *AWS_GameLift_ContainerFleet) GetAtt__GameServerContainerGroupDefinitionArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.GameServerContainerGroupDefinitionArn))
}

// GetAtt__MaximumGameServerContainerGroupsPerInstance returns a $cfz.Expression[int32] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: MaximumGameServerContainerGroupsPerInstance
func (t *AWS_GameLift_ContainerFleet) GetAtt__MaximumGameServerContainerGroupsPerInstance() cfz.Expression[int32] {
	return cfz.GetAtt[int32](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.MaximumGameServerContainerGroupsPerInstance))
}

// GetAtt__PerInstanceContainerGroupDefinitionArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: PerInstanceContainerGroupDefinitionArn
func (t *AWS_GameLift_ContainerFleet) GetAtt__PerInstanceContainerGroupDefinitionArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.PerInstanceContainerGroupDefinitionArn))
}

// GetAtt__Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_GameLift_ContainerFleet) GetAtt__Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_GameLift_ContainerFleet__AttributesMap.Status))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTime
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__CreationTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTime", t.GetAtt__CreationTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DeploymentDetails returns a conventionally configured output for an attribute of this resource.
// Attribute: DeploymentDetails
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__DeploymentDetails(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDeploymentDetails", t.GetAtt__DeploymentDetails())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DeploymentDetails_LatestDeploymentId returns a conventionally configured output for an attribute of this resource.
// Attribute: DeploymentDetails.LatestDeploymentId
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__DeploymentDetails_LatestDeploymentId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDeploymentDetailsLatestDeploymentId", t.GetAtt__DeploymentDetails_LatestDeploymentId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FleetArn returns a conventionally configured output for an attribute of this resource.
// Attribute: FleetArn
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__FleetArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFleetArn", t.GetAtt__FleetArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__FleetId returns a conventionally configured output for an attribute of this resource.
// Attribute: FleetId
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__FleetId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttFleetId", t.GetAtt__FleetId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__GameServerContainerGroupDefinitionArn returns a conventionally configured output for an attribute of this resource.
// Attribute: GameServerContainerGroupDefinitionArn
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__GameServerContainerGroupDefinitionArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttGameServerContainerGroupDefinitionArn", t.GetAtt__GameServerContainerGroupDefinitionArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__MaximumGameServerContainerGroupsPerInstance returns a conventionally configured output for an attribute of this resource.
// Attribute: MaximumGameServerContainerGroupsPerInstance
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__MaximumGameServerContainerGroupsPerInstance(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttMaximumGameServerContainerGroupsPerInstance", t.GetAtt__MaximumGameServerContainerGroupsPerInstance())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__PerInstanceContainerGroupDefinitionArn returns a conventionally configured output for an attribute of this resource.
// Attribute: PerInstanceContainerGroupDefinitionArn
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__PerInstanceContainerGroupDefinitionArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttPerInstanceContainerGroupDefinitionArn", t.GetAtt__PerInstanceContainerGroupDefinitionArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_GameLift_ContainerFleet) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_GameLift_ContainerFleet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_GameLift_ContainerFleet

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

func (t *AWS_GameLift_ContainerFleet) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
