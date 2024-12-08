// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_redshift

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Redshift_ClusterSecurityGroupIngress)(nil)
	_ cfz.Resource                   = (*AWS_Redshift_ClusterSecurityGroupIngress)(nil)
)

const (
	// AWS_Redshift_ClusterSecurityGroupIngress__Type is the CloudFormation type for AWS::Redshift::ClusterSecurityGroupIngress.
	AWS_Redshift_ClusterSecurityGroupIngress__Type = "AWS::Redshift::ClusterSecurityGroupIngress"
)

var (
	// AWS_Redshift_ClusterSecurityGroupIngress__PropertiesMap reports all the CloudFormation properties for AWS::Redshift::ClusterSecurityGroupIngress.
	AWS_Redshift_ClusterSecurityGroupIngress__PropertiesMap = struct {
		CIDRIP                   string
		ClusterSecurityGroupName string
		EC2SecurityGroupName     string
		EC2SecurityGroupOwnerId  string
	}{
		CIDRIP:                   "CIDRIP",
		ClusterSecurityGroupName: "ClusterSecurityGroupName",
		EC2SecurityGroupName:     "EC2SecurityGroupName",
		EC2SecurityGroupOwnerId:  "EC2SecurityGroupOwnerId",
	}

	// AWS_Redshift_ClusterSecurityGroupIngress__PropertiesSlice reports all the CloudFormation properties for AWS::Redshift::ClusterSecurityGroupIngress.
	AWS_Redshift_ClusterSecurityGroupIngress__PropertiesSlice = []string{
		AWS_Redshift_ClusterSecurityGroupIngress__PropertiesMap.CIDRIP,
		AWS_Redshift_ClusterSecurityGroupIngress__PropertiesMap.ClusterSecurityGroupName,
		AWS_Redshift_ClusterSecurityGroupIngress__PropertiesMap.EC2SecurityGroupName,
		AWS_Redshift_ClusterSecurityGroupIngress__PropertiesMap.EC2SecurityGroupOwnerId,
	}
)

// AWS_Redshift_ClusterSecurityGroupIngress is a binding for AWS::Redshift::ClusterSecurityGroupIngress.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html
type AWS_Redshift_ClusterSecurityGroupIngress struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CIDRIP is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-cidrip
	CIDRIP cfz.Expression[string] `json:"CIDRIP,omitempty"`

	// ClusterSecurityGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-clustersecuritygroupname
	ClusterSecurityGroupName cfz.Expression[string] `json:"ClusterSecurityGroupName,omitempty"`

	// EC2SecurityGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupname
	EC2SecurityGroupName cfz.Expression[string] `json:"EC2SecurityGroupName,omitempty"`

	// EC2SecurityGroupOwnerId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroupingress.html#cfn-redshift-clustersecuritygroupingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerId cfz.Expression[string] `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// New__AWS_Redshift_ClusterSecurityGroupIngress initializes a new *AWS_Redshift_ClusterSecurityGroupIngress.
func New__AWS_Redshift_ClusterSecurityGroupIngress(logicalName string) *AWS_Redshift_ClusterSecurityGroupIngress {
	return &AWS_Redshift_ClusterSecurityGroupIngress{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Redshift_ClusterSecurityGroupIngress) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Redshift_ClusterSecurityGroupIngress) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Redshift_ClusterSecurityGroupIngress) GetType() string {
	return AWS_Redshift_ClusterSecurityGroupIngress__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__LogicalName(v string) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__RequestedOutputs(v []cfz.Output) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Add__RequestedOutputs(v ...cfz.Output) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CIDRIP updates property "CIDRIP".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__CIDRIP(v cfz.Expression[string]) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.CIDRIP = v
	return t
}

// SetV__CIDRIP updates property "CIDRIP".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) SetV__CIDRIP(v string) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.CIDRIP = cfz.V(v)
	return t
}

// Set__ClusterSecurityGroupName updates property "ClusterSecurityGroupName".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__ClusterSecurityGroupName(v cfz.Expression[string]) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.ClusterSecurityGroupName = v
	return t
}

// SetV__ClusterSecurityGroupName updates property "ClusterSecurityGroupName".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) SetV__ClusterSecurityGroupName(v string) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.ClusterSecurityGroupName = cfz.V(v)
	return t
}

// Set__EC2SecurityGroupName updates property "EC2SecurityGroupName".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__EC2SecurityGroupName(v cfz.Expression[string]) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.EC2SecurityGroupName = v
	return t
}

// SetV__EC2SecurityGroupName updates property "EC2SecurityGroupName".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) SetV__EC2SecurityGroupName(v string) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.EC2SecurityGroupName = cfz.V(v)
	return t
}

// Set__EC2SecurityGroupOwnerId updates property "EC2SecurityGroupOwnerId".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Set__EC2SecurityGroupOwnerId(v cfz.Expression[string]) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.EC2SecurityGroupOwnerId = v
	return t
}

// SetV__EC2SecurityGroupOwnerId updates property "EC2SecurityGroupOwnerId".
func (t *AWS_Redshift_ClusterSecurityGroupIngress) SetV__EC2SecurityGroupOwnerId(v string) *AWS_Redshift_ClusterSecurityGroupIngress {
	t.EC2SecurityGroupOwnerId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Redshift_ClusterSecurityGroupIngress) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Redshift_ClusterSecurityGroupIngress) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Redshift_ClusterSecurityGroupIngress) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Redshift_ClusterSecurityGroupIngress

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

func (t *AWS_Redshift_ClusterSecurityGroupIngress) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
