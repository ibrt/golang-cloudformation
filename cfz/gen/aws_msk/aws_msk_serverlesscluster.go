// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_msk

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MSK_ServerlessCluster)(nil)
	_ cfz.Resource                   = (*AWS_MSK_ServerlessCluster)(nil)
)

const (
	// AWS_MSK_ServerlessCluster__Type is the CloudFormation type for AWS::MSK::ServerlessCluster.
	AWS_MSK_ServerlessCluster__Type = "AWS::MSK::ServerlessCluster"
)

var (
	// AWS_MSK_ServerlessCluster__AttributesMap reports all the CloudFormation attributes for AWS::MSK::ServerlessCluster.
	AWS_MSK_ServerlessCluster__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_MSK_ServerlessCluster__AttributesSlice reports all the CloudFormation attributes for AWS::MSK::ServerlessCluster.
	AWS_MSK_ServerlessCluster__AttributesSlice = []string{
		AWS_MSK_ServerlessCluster__AttributesMap.Arn,
	}
)

var (
	// AWS_MSK_ServerlessCluster__PropertiesMap reports all the CloudFormation properties for AWS::MSK::ServerlessCluster.
	AWS_MSK_ServerlessCluster__PropertiesMap = struct {
		ClientAuthentication string
		ClusterName          string
		Tags                 string
		VpcConfigs           string
	}{
		ClientAuthentication: "ClientAuthentication",
		ClusterName:          "ClusterName",
		Tags:                 "Tags",
		VpcConfigs:           "VpcConfigs",
	}

	// AWS_MSK_ServerlessCluster__PropertiesSlice reports all the CloudFormation properties for AWS::MSK::ServerlessCluster.
	AWS_MSK_ServerlessCluster__PropertiesSlice = []string{
		AWS_MSK_ServerlessCluster__PropertiesMap.ClientAuthentication,
		AWS_MSK_ServerlessCluster__PropertiesMap.ClusterName,
		AWS_MSK_ServerlessCluster__PropertiesMap.Tags,
		AWS_MSK_ServerlessCluster__PropertiesMap.VpcConfigs,
	}
)

// AWS_MSK_ServerlessCluster is a binding for AWS::MSK::ServerlessCluster.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html
type AWS_MSK_ServerlessCluster struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ClientAuthentication is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-clientauthentication
	ClientAuthentication cfz.Expression[AWS_MSK_ServerlessCluster_ClientAuthentication] `json:"ClientAuthentication,omitempty"`

	// ClusterName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-clustername
	ClusterName cfz.Expression[string] `json:"ClusterName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-tags
	Tags cfz.ExpressionMap[string] `json:"Tags,omitempty"`

	// VpcConfigs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-msk-serverlesscluster.html#cfn-msk-serverlesscluster-vpcconfigs
	VpcConfigs cfz.ExpressionSlice[AWS_MSK_ServerlessCluster_VpcConfig] `json:"VpcConfigs,omitempty"`
}

// New__AWS_MSK_ServerlessCluster initializes a new *AWS_MSK_ServerlessCluster.
func New__AWS_MSK_ServerlessCluster(logicalName string) *AWS_MSK_ServerlessCluster {
	return &AWS_MSK_ServerlessCluster{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MSK_ServerlessCluster) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MSK_ServerlessCluster) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MSK_ServerlessCluster) GetType() string {
	return AWS_MSK_ServerlessCluster__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MSK_ServerlessCluster) Set__LogicalName(v string) *AWS_MSK_ServerlessCluster {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MSK_ServerlessCluster) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MSK_ServerlessCluster {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MSK_ServerlessCluster) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MSK_ServerlessCluster {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MSK_ServerlessCluster) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MSK_ServerlessCluster {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MSK_ServerlessCluster) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MSK_ServerlessCluster {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MSK_ServerlessCluster) Set__RequestedOutputs(v []cfz.Output) *AWS_MSK_ServerlessCluster {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MSK_ServerlessCluster) Add__RequestedOutputs(v ...cfz.Output) *AWS_MSK_ServerlessCluster {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ClientAuthentication updates property "ClientAuthentication".
func (t *AWS_MSK_ServerlessCluster) Set__ClientAuthentication(v cfz.Expression[AWS_MSK_ServerlessCluster_ClientAuthentication]) *AWS_MSK_ServerlessCluster {
	t.ClientAuthentication = v
	return t
}

// SetV__ClientAuthentication updates property "ClientAuthentication".
func (t *AWS_MSK_ServerlessCluster) SetV__ClientAuthentication(v AWS_MSK_ServerlessCluster_ClientAuthentication) *AWS_MSK_ServerlessCluster {
	t.ClientAuthentication = cfz.V(v)
	return t
}

// Set__ClusterName updates property "ClusterName".
func (t *AWS_MSK_ServerlessCluster) Set__ClusterName(v cfz.Expression[string]) *AWS_MSK_ServerlessCluster {
	t.ClusterName = v
	return t
}

// SetV__ClusterName updates property "ClusterName".
func (t *AWS_MSK_ServerlessCluster) SetV__ClusterName(v string) *AWS_MSK_ServerlessCluster {
	t.ClusterName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MSK_ServerlessCluster) Set__Tags(v cfz.ExpressionMap[string]) *AWS_MSK_ServerlessCluster {
	t.Tags = v
	return t
}

// SetM__Tags updates property "Tags".
func (t *AWS_MSK_ServerlessCluster) SetM__Tags(v ...map[string]cfz.Expression[string]) *AWS_MSK_ServerlessCluster {
	t.Tags = cfz.M(v...)
	return t
}

// SetMV__Tags updates property "Tags".
func (t *AWS_MSK_ServerlessCluster) SetMV__Tags(v ...map[string]string) *AWS_MSK_ServerlessCluster {
	t.Tags = cfz.MV(v...)
	return t
}

// Set__VpcConfigs updates property "VpcConfigs".
func (t *AWS_MSK_ServerlessCluster) Set__VpcConfigs(v cfz.ExpressionSlice[AWS_MSK_ServerlessCluster_VpcConfig]) *AWS_MSK_ServerlessCluster {
	t.VpcConfigs = v
	return t
}

// SetS__VpcConfigs updates property "VpcConfigs".
func (t *AWS_MSK_ServerlessCluster) SetS__VpcConfigs(v ...cfz.Expression[AWS_MSK_ServerlessCluster_VpcConfig]) *AWS_MSK_ServerlessCluster {
	t.VpcConfigs = cfz.S(v...)
	return t
}

// SetSV__VpcConfigs updates property "VpcConfigs".
func (t *AWS_MSK_ServerlessCluster) SetSV__VpcConfigs(v ...AWS_MSK_ServerlessCluster_VpcConfig) *AWS_MSK_ServerlessCluster {
	t.VpcConfigs = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MSK_ServerlessCluster) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_MSK_ServerlessCluster) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MSK_ServerlessCluster__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MSK_ServerlessCluster) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_MSK_ServerlessCluster) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MSK_ServerlessCluster) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MSK_ServerlessCluster

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

func (t *AWS_MSK_ServerlessCluster) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
