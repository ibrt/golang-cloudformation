// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_rds

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_RDS_DBProxyEndpoint)(nil)
	_ cfz.Resource                   = (*AWS_RDS_DBProxyEndpoint)(nil)
)

const (
	// AWS_RDS_DBProxyEndpoint__Type is the CloudFormation type for AWS::RDS::DBProxyEndpoint.
	AWS_RDS_DBProxyEndpoint__Type = "AWS::RDS::DBProxyEndpoint"
)

var (
	// AWS_RDS_DBProxyEndpoint__AttributesMap reports all the CloudFormation attributes for AWS::RDS::DBProxyEndpoint.
	AWS_RDS_DBProxyEndpoint__AttributesMap = struct {
		DBProxyEndpointArn string
		Endpoint           string
		IsDefault          string
		VpcId              string
	}{
		DBProxyEndpointArn: "DBProxyEndpointArn",
		Endpoint:           "Endpoint",
		IsDefault:          "IsDefault",
		VpcId:              "VpcId",
	}

	// AWS_RDS_DBProxyEndpoint__AttributesSlice reports all the CloudFormation attributes for AWS::RDS::DBProxyEndpoint.
	AWS_RDS_DBProxyEndpoint__AttributesSlice = []string{
		AWS_RDS_DBProxyEndpoint__AttributesMap.DBProxyEndpointArn,
		AWS_RDS_DBProxyEndpoint__AttributesMap.Endpoint,
		AWS_RDS_DBProxyEndpoint__AttributesMap.IsDefault,
		AWS_RDS_DBProxyEndpoint__AttributesMap.VpcId,
	}
)

var (
	// AWS_RDS_DBProxyEndpoint__PropertiesMap reports all the CloudFormation properties for AWS::RDS::DBProxyEndpoint.
	AWS_RDS_DBProxyEndpoint__PropertiesMap = struct {
		DBProxyEndpointName string
		DBProxyName         string
		Tags                string
		TargetRole          string
		VpcSecurityGroupIds string
		VpcSubnetIds        string
	}{
		DBProxyEndpointName: "DBProxyEndpointName",
		DBProxyName:         "DBProxyName",
		Tags:                "Tags",
		TargetRole:          "TargetRole",
		VpcSecurityGroupIds: "VpcSecurityGroupIds",
		VpcSubnetIds:        "VpcSubnetIds",
	}

	// AWS_RDS_DBProxyEndpoint__PropertiesSlice reports all the CloudFormation properties for AWS::RDS::DBProxyEndpoint.
	AWS_RDS_DBProxyEndpoint__PropertiesSlice = []string{
		AWS_RDS_DBProxyEndpoint__PropertiesMap.DBProxyEndpointName,
		AWS_RDS_DBProxyEndpoint__PropertiesMap.DBProxyName,
		AWS_RDS_DBProxyEndpoint__PropertiesMap.Tags,
		AWS_RDS_DBProxyEndpoint__PropertiesMap.TargetRole,
		AWS_RDS_DBProxyEndpoint__PropertiesMap.VpcSecurityGroupIds,
		AWS_RDS_DBProxyEndpoint__PropertiesMap.VpcSubnetIds,
	}
)

// AWS_RDS_DBProxyEndpoint is a binding for AWS::RDS::DBProxyEndpoint.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html
type AWS_RDS_DBProxyEndpoint struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DBProxyEndpointName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-dbproxyendpointname
	DBProxyEndpointName cfz.Expression[string] `json:"DBProxyEndpointName,omitempty"`

	// DBProxyName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-dbproxyname
	DBProxyName cfz.Expression[string] `json:"DBProxyName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-tags
	Tags cfz.ExpressionSlice[AWS_RDS_DBProxyEndpoint_TagFormat] `json:"Tags,omitempty"`

	// TargetRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-targetrole
	TargetRole cfz.Expression[string] `json:"TargetRole,omitempty"`

	// VpcSecurityGroupIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-vpcsecuritygroupids
	VpcSecurityGroupIds cfz.ExpressionSlice[string] `json:"VpcSecurityGroupIds,omitempty"`

	// VpcSubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-vpcsubnetids
	VpcSubnetIds cfz.ExpressionSlice[string] `json:"VpcSubnetIds,omitempty"`
}

// New__AWS_RDS_DBProxyEndpoint initializes a new *AWS_RDS_DBProxyEndpoint.
func New__AWS_RDS_DBProxyEndpoint(logicalName string) *AWS_RDS_DBProxyEndpoint {
	return &AWS_RDS_DBProxyEndpoint{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_RDS_DBProxyEndpoint) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_RDS_DBProxyEndpoint) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_RDS_DBProxyEndpoint) GetType() string {
	return AWS_RDS_DBProxyEndpoint__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_RDS_DBProxyEndpoint) Set__LogicalName(v string) *AWS_RDS_DBProxyEndpoint {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_RDS_DBProxyEndpoint) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_RDS_DBProxyEndpoint {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_RDS_DBProxyEndpoint) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_RDS_DBProxyEndpoint {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_RDS_DBProxyEndpoint) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_RDS_DBProxyEndpoint {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_RDS_DBProxyEndpoint) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_RDS_DBProxyEndpoint {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_RDS_DBProxyEndpoint) Set__RequestedOutputs(v []cfz.Output) *AWS_RDS_DBProxyEndpoint {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_RDS_DBProxyEndpoint) Add__RequestedOutputs(v ...cfz.Output) *AWS_RDS_DBProxyEndpoint {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DBProxyEndpointName updates property "DBProxyEndpointName".
func (t *AWS_RDS_DBProxyEndpoint) Set__DBProxyEndpointName(v cfz.Expression[string]) *AWS_RDS_DBProxyEndpoint {
	t.DBProxyEndpointName = v
	return t
}

// SetV__DBProxyEndpointName updates property "DBProxyEndpointName".
func (t *AWS_RDS_DBProxyEndpoint) SetV__DBProxyEndpointName(v string) *AWS_RDS_DBProxyEndpoint {
	t.DBProxyEndpointName = cfz.V(v)
	return t
}

// Set__DBProxyName updates property "DBProxyName".
func (t *AWS_RDS_DBProxyEndpoint) Set__DBProxyName(v cfz.Expression[string]) *AWS_RDS_DBProxyEndpoint {
	t.DBProxyName = v
	return t
}

// SetV__DBProxyName updates property "DBProxyName".
func (t *AWS_RDS_DBProxyEndpoint) SetV__DBProxyName(v string) *AWS_RDS_DBProxyEndpoint {
	t.DBProxyName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_RDS_DBProxyEndpoint) Set__Tags(v cfz.ExpressionSlice[AWS_RDS_DBProxyEndpoint_TagFormat]) *AWS_RDS_DBProxyEndpoint {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_RDS_DBProxyEndpoint) SetS__Tags(v ...cfz.Expression[AWS_RDS_DBProxyEndpoint_TagFormat]) *AWS_RDS_DBProxyEndpoint {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_RDS_DBProxyEndpoint) SetSV__Tags(v ...AWS_RDS_DBProxyEndpoint_TagFormat) *AWS_RDS_DBProxyEndpoint {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TargetRole updates property "TargetRole".
func (t *AWS_RDS_DBProxyEndpoint) Set__TargetRole(v cfz.Expression[string]) *AWS_RDS_DBProxyEndpoint {
	t.TargetRole = v
	return t
}

// SetV__TargetRole updates property "TargetRole".
func (t *AWS_RDS_DBProxyEndpoint) SetV__TargetRole(v string) *AWS_RDS_DBProxyEndpoint {
	t.TargetRole = cfz.V(v)
	return t
}

// Set__VpcSecurityGroupIds updates property "VpcSecurityGroupIds".
func (t *AWS_RDS_DBProxyEndpoint) Set__VpcSecurityGroupIds(v cfz.ExpressionSlice[string]) *AWS_RDS_DBProxyEndpoint {
	t.VpcSecurityGroupIds = v
	return t
}

// SetS__VpcSecurityGroupIds updates property "VpcSecurityGroupIds".
func (t *AWS_RDS_DBProxyEndpoint) SetS__VpcSecurityGroupIds(v ...cfz.Expression[string]) *AWS_RDS_DBProxyEndpoint {
	t.VpcSecurityGroupIds = cfz.S(v...)
	return t
}

// SetSV__VpcSecurityGroupIds updates property "VpcSecurityGroupIds".
func (t *AWS_RDS_DBProxyEndpoint) SetSV__VpcSecurityGroupIds(v ...string) *AWS_RDS_DBProxyEndpoint {
	t.VpcSecurityGroupIds = cfz.SV(v...)
	return t
}

// Set__VpcSubnetIds updates property "VpcSubnetIds".
func (t *AWS_RDS_DBProxyEndpoint) Set__VpcSubnetIds(v cfz.ExpressionSlice[string]) *AWS_RDS_DBProxyEndpoint {
	t.VpcSubnetIds = v
	return t
}

// SetS__VpcSubnetIds updates property "VpcSubnetIds".
func (t *AWS_RDS_DBProxyEndpoint) SetS__VpcSubnetIds(v ...cfz.Expression[string]) *AWS_RDS_DBProxyEndpoint {
	t.VpcSubnetIds = cfz.S(v...)
	return t
}

// SetSV__VpcSubnetIds updates property "VpcSubnetIds".
func (t *AWS_RDS_DBProxyEndpoint) SetSV__VpcSubnetIds(v ...string) *AWS_RDS_DBProxyEndpoint {
	t.VpcSubnetIds = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_RDS_DBProxyEndpoint) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__DBProxyEndpointArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DBProxyEndpointArn
func (t *AWS_RDS_DBProxyEndpoint) GetAtt__DBProxyEndpointArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_RDS_DBProxyEndpoint__AttributesMap.DBProxyEndpointArn))
}

// GetAtt__Endpoint returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Endpoint
func (t *AWS_RDS_DBProxyEndpoint) GetAtt__Endpoint() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_RDS_DBProxyEndpoint__AttributesMap.Endpoint))
}

// GetAtt__IsDefault returns a $cfz.Expression[bool] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: IsDefault
func (t *AWS_RDS_DBProxyEndpoint) GetAtt__IsDefault() cfz.Expression[bool] {
	return cfz.GetAtt[bool](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_RDS_DBProxyEndpoint__AttributesMap.IsDefault))
}

// GetAtt__VpcId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: VpcId
func (t *AWS_RDS_DBProxyEndpoint) GetAtt__VpcId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_RDS_DBProxyEndpoint__AttributesMap.VpcId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_RDS_DBProxyEndpoint) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DBProxyEndpointArn returns a conventionally configured output for an attribute of this resource.
// Attribute: DBProxyEndpointArn
func (t *AWS_RDS_DBProxyEndpoint) GetConventionalOutputAtt__DBProxyEndpointArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDBProxyEndpointArn", t.GetAtt__DBProxyEndpointArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Endpoint returns a conventionally configured output for an attribute of this resource.
// Attribute: Endpoint
func (t *AWS_RDS_DBProxyEndpoint) GetConventionalOutputAtt__Endpoint(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttEndpoint", t.GetAtt__Endpoint())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__IsDefault returns a conventionally configured output for an attribute of this resource.
// Attribute: IsDefault
func (t *AWS_RDS_DBProxyEndpoint) GetConventionalOutputAtt__IsDefault(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttIsDefault", t.GetAtt__IsDefault())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__VpcId returns a conventionally configured output for an attribute of this resource.
// Attribute: VpcId
func (t *AWS_RDS_DBProxyEndpoint) GetConventionalOutputAtt__VpcId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVpcId", t.GetAtt__VpcId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_RDS_DBProxyEndpoint) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_RDS_DBProxyEndpoint

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

func (t *AWS_RDS_DBProxyEndpoint) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
