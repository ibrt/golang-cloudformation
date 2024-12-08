// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codebuild

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_CodeBuild_Fleet)(nil)
	_ cfz.Resource                   = (*AWS_CodeBuild_Fleet)(nil)
)

const (
	// AWS_CodeBuild_Fleet__Type is the CloudFormation type for AWS::CodeBuild::Fleet.
	AWS_CodeBuild_Fleet__Type = "AWS::CodeBuild::Fleet"
)

var (
	// AWS_CodeBuild_Fleet__AttributesMap reports all the CloudFormation attributes for AWS::CodeBuild::Fleet.
	AWS_CodeBuild_Fleet__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_CodeBuild_Fleet__AttributesSlice reports all the CloudFormation attributes for AWS::CodeBuild::Fleet.
	AWS_CodeBuild_Fleet__AttributesSlice = []string{
		AWS_CodeBuild_Fleet__AttributesMap.Arn,
	}
)

var (
	// AWS_CodeBuild_Fleet__PropertiesMap reports all the CloudFormation properties for AWS::CodeBuild::Fleet.
	AWS_CodeBuild_Fleet__PropertiesMap = struct {
		BaseCapacity            string
		ComputeConfiguration    string
		ComputeType             string
		EnvironmentType         string
		FleetProxyConfiguration string
		FleetServiceRole        string
		FleetVpcConfig          string
		ImageId                 string
		Name                    string
		OverflowBehavior        string
		ScalingConfiguration    string
		Tags                    string
	}{
		BaseCapacity:            "BaseCapacity",
		ComputeConfiguration:    "ComputeConfiguration",
		ComputeType:             "ComputeType",
		EnvironmentType:         "EnvironmentType",
		FleetProxyConfiguration: "FleetProxyConfiguration",
		FleetServiceRole:        "FleetServiceRole",
		FleetVpcConfig:          "FleetVpcConfig",
		ImageId:                 "ImageId",
		Name:                    "Name",
		OverflowBehavior:        "OverflowBehavior",
		ScalingConfiguration:    "ScalingConfiguration",
		Tags:                    "Tags",
	}

	// AWS_CodeBuild_Fleet__PropertiesSlice reports all the CloudFormation properties for AWS::CodeBuild::Fleet.
	AWS_CodeBuild_Fleet__PropertiesSlice = []string{
		AWS_CodeBuild_Fleet__PropertiesMap.BaseCapacity,
		AWS_CodeBuild_Fleet__PropertiesMap.ComputeConfiguration,
		AWS_CodeBuild_Fleet__PropertiesMap.ComputeType,
		AWS_CodeBuild_Fleet__PropertiesMap.EnvironmentType,
		AWS_CodeBuild_Fleet__PropertiesMap.FleetProxyConfiguration,
		AWS_CodeBuild_Fleet__PropertiesMap.FleetServiceRole,
		AWS_CodeBuild_Fleet__PropertiesMap.FleetVpcConfig,
		AWS_CodeBuild_Fleet__PropertiesMap.ImageId,
		AWS_CodeBuild_Fleet__PropertiesMap.Name,
		AWS_CodeBuild_Fleet__PropertiesMap.OverflowBehavior,
		AWS_CodeBuild_Fleet__PropertiesMap.ScalingConfiguration,
		AWS_CodeBuild_Fleet__PropertiesMap.Tags,
	}
)

// AWS_CodeBuild_Fleet is a binding for AWS::CodeBuild::Fleet.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html
type AWS_CodeBuild_Fleet struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// BaseCapacity is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-basecapacity
	BaseCapacity cfz.Expression[int32] `json:"BaseCapacity,omitempty"`

	// ComputeConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-computeconfiguration
	ComputeConfiguration cfz.Expression[AWS_CodeBuild_Fleet_ComputeConfiguration] `json:"ComputeConfiguration,omitempty"`

	// ComputeType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-computetype
	ComputeType cfz.Expression[string] `json:"ComputeType,omitempty"`

	// EnvironmentType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-environmenttype
	EnvironmentType cfz.Expression[string] `json:"EnvironmentType,omitempty"`

	// FleetProxyConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-fleetproxyconfiguration
	FleetProxyConfiguration cfz.Expression[AWS_CodeBuild_Fleet_ProxyConfiguration] `json:"FleetProxyConfiguration,omitempty"`

	// FleetServiceRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-fleetservicerole
	FleetServiceRole cfz.Expression[string] `json:"FleetServiceRole,omitempty"`

	// FleetVpcConfig is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-fleetvpcconfig
	FleetVpcConfig cfz.Expression[AWS_CodeBuild_Fleet_VpcConfig] `json:"FleetVpcConfig,omitempty"`

	// ImageId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-imageid
	ImageId cfz.Expression[string] `json:"ImageId,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// OverflowBehavior is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-overflowbehavior
	OverflowBehavior cfz.Expression[string] `json:"OverflowBehavior,omitempty"`

	// ScalingConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-scalingconfiguration
	ScalingConfiguration cfz.Expression[AWS_CodeBuild_Fleet_ScalingConfigurationInput] `json:"ScalingConfiguration,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html#cfn-codebuild-fleet-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_CodeBuild_Fleet initializes a new *AWS_CodeBuild_Fleet.
func New__AWS_CodeBuild_Fleet(logicalName string) *AWS_CodeBuild_Fleet {
	return &AWS_CodeBuild_Fleet{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_CodeBuild_Fleet) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_CodeBuild_Fleet) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_CodeBuild_Fleet) GetType() string {
	return AWS_CodeBuild_Fleet__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_CodeBuild_Fleet) Set__LogicalName(v string) *AWS_CodeBuild_Fleet {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_CodeBuild_Fleet) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_CodeBuild_Fleet {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_CodeBuild_Fleet) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_CodeBuild_Fleet {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_CodeBuild_Fleet) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_CodeBuild_Fleet {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_CodeBuild_Fleet) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_CodeBuild_Fleet {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_CodeBuild_Fleet) Set__RequestedOutputs(v []cfz.Output) *AWS_CodeBuild_Fleet {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_CodeBuild_Fleet) Add__RequestedOutputs(v ...cfz.Output) *AWS_CodeBuild_Fleet {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__BaseCapacity updates property "BaseCapacity".
func (t *AWS_CodeBuild_Fleet) Set__BaseCapacity(v cfz.Expression[int32]) *AWS_CodeBuild_Fleet {
	t.BaseCapacity = v
	return t
}

// SetV__BaseCapacity updates property "BaseCapacity".
func (t *AWS_CodeBuild_Fleet) SetV__BaseCapacity(v int32) *AWS_CodeBuild_Fleet {
	t.BaseCapacity = cfz.V(v)
	return t
}

// Set__ComputeConfiguration updates property "ComputeConfiguration".
func (t *AWS_CodeBuild_Fleet) Set__ComputeConfiguration(v cfz.Expression[AWS_CodeBuild_Fleet_ComputeConfiguration]) *AWS_CodeBuild_Fleet {
	t.ComputeConfiguration = v
	return t
}

// SetV__ComputeConfiguration updates property "ComputeConfiguration".
func (t *AWS_CodeBuild_Fleet) SetV__ComputeConfiguration(v AWS_CodeBuild_Fleet_ComputeConfiguration) *AWS_CodeBuild_Fleet {
	t.ComputeConfiguration = cfz.V(v)
	return t
}

// Set__ComputeType updates property "ComputeType".
func (t *AWS_CodeBuild_Fleet) Set__ComputeType(v cfz.Expression[string]) *AWS_CodeBuild_Fleet {
	t.ComputeType = v
	return t
}

// SetV__ComputeType updates property "ComputeType".
func (t *AWS_CodeBuild_Fleet) SetV__ComputeType(v string) *AWS_CodeBuild_Fleet {
	t.ComputeType = cfz.V(v)
	return t
}

// Set__EnvironmentType updates property "EnvironmentType".
func (t *AWS_CodeBuild_Fleet) Set__EnvironmentType(v cfz.Expression[string]) *AWS_CodeBuild_Fleet {
	t.EnvironmentType = v
	return t
}

// SetV__EnvironmentType updates property "EnvironmentType".
func (t *AWS_CodeBuild_Fleet) SetV__EnvironmentType(v string) *AWS_CodeBuild_Fleet {
	t.EnvironmentType = cfz.V(v)
	return t
}

// Set__FleetProxyConfiguration updates property "FleetProxyConfiguration".
func (t *AWS_CodeBuild_Fleet) Set__FleetProxyConfiguration(v cfz.Expression[AWS_CodeBuild_Fleet_ProxyConfiguration]) *AWS_CodeBuild_Fleet {
	t.FleetProxyConfiguration = v
	return t
}

// SetV__FleetProxyConfiguration updates property "FleetProxyConfiguration".
func (t *AWS_CodeBuild_Fleet) SetV__FleetProxyConfiguration(v AWS_CodeBuild_Fleet_ProxyConfiguration) *AWS_CodeBuild_Fleet {
	t.FleetProxyConfiguration = cfz.V(v)
	return t
}

// Set__FleetServiceRole updates property "FleetServiceRole".
func (t *AWS_CodeBuild_Fleet) Set__FleetServiceRole(v cfz.Expression[string]) *AWS_CodeBuild_Fleet {
	t.FleetServiceRole = v
	return t
}

// SetV__FleetServiceRole updates property "FleetServiceRole".
func (t *AWS_CodeBuild_Fleet) SetV__FleetServiceRole(v string) *AWS_CodeBuild_Fleet {
	t.FleetServiceRole = cfz.V(v)
	return t
}

// Set__FleetVpcConfig updates property "FleetVpcConfig".
func (t *AWS_CodeBuild_Fleet) Set__FleetVpcConfig(v cfz.Expression[AWS_CodeBuild_Fleet_VpcConfig]) *AWS_CodeBuild_Fleet {
	t.FleetVpcConfig = v
	return t
}

// SetV__FleetVpcConfig updates property "FleetVpcConfig".
func (t *AWS_CodeBuild_Fleet) SetV__FleetVpcConfig(v AWS_CodeBuild_Fleet_VpcConfig) *AWS_CodeBuild_Fleet {
	t.FleetVpcConfig = cfz.V(v)
	return t
}

// Set__ImageId updates property "ImageId".
func (t *AWS_CodeBuild_Fleet) Set__ImageId(v cfz.Expression[string]) *AWS_CodeBuild_Fleet {
	t.ImageId = v
	return t
}

// SetV__ImageId updates property "ImageId".
func (t *AWS_CodeBuild_Fleet) SetV__ImageId(v string) *AWS_CodeBuild_Fleet {
	t.ImageId = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_CodeBuild_Fleet) Set__Name(v cfz.Expression[string]) *AWS_CodeBuild_Fleet {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_CodeBuild_Fleet) SetV__Name(v string) *AWS_CodeBuild_Fleet {
	t.Name = cfz.V(v)
	return t
}

// Set__OverflowBehavior updates property "OverflowBehavior".
func (t *AWS_CodeBuild_Fleet) Set__OverflowBehavior(v cfz.Expression[string]) *AWS_CodeBuild_Fleet {
	t.OverflowBehavior = v
	return t
}

// SetV__OverflowBehavior updates property "OverflowBehavior".
func (t *AWS_CodeBuild_Fleet) SetV__OverflowBehavior(v string) *AWS_CodeBuild_Fleet {
	t.OverflowBehavior = cfz.V(v)
	return t
}

// Set__ScalingConfiguration updates property "ScalingConfiguration".
func (t *AWS_CodeBuild_Fleet) Set__ScalingConfiguration(v cfz.Expression[AWS_CodeBuild_Fleet_ScalingConfigurationInput]) *AWS_CodeBuild_Fleet {
	t.ScalingConfiguration = v
	return t
}

// SetV__ScalingConfiguration updates property "ScalingConfiguration".
func (t *AWS_CodeBuild_Fleet) SetV__ScalingConfiguration(v AWS_CodeBuild_Fleet_ScalingConfigurationInput) *AWS_CodeBuild_Fleet {
	t.ScalingConfiguration = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_CodeBuild_Fleet) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_CodeBuild_Fleet {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_CodeBuild_Fleet) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_CodeBuild_Fleet {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_CodeBuild_Fleet) SetSV__Tags(v ...cfz.Tag) *AWS_CodeBuild_Fleet {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_CodeBuild_Fleet) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_CodeBuild_Fleet) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CodeBuild_Fleet__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_CodeBuild_Fleet) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_CodeBuild_Fleet) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_CodeBuild_Fleet) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_CodeBuild_Fleet

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

func (t *AWS_CodeBuild_Fleet) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
