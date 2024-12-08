// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_SecurityGroupIngress)(nil)
	_ cfz.Resource                   = (*AWS_EC2_SecurityGroupIngress)(nil)
)

const (
	// AWS_EC2_SecurityGroupIngress__Type is the CloudFormation type for AWS::EC2::SecurityGroupIngress.
	AWS_EC2_SecurityGroupIngress__Type = "AWS::EC2::SecurityGroupIngress"
)

var (
	// AWS_EC2_SecurityGroupIngress__AttributesMap reports all the CloudFormation attributes for AWS::EC2::SecurityGroupIngress.
	AWS_EC2_SecurityGroupIngress__AttributesMap = struct {
		Id string
	}{
		Id: "Id",
	}

	// AWS_EC2_SecurityGroupIngress__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::SecurityGroupIngress.
	AWS_EC2_SecurityGroupIngress__AttributesSlice = []string{
		AWS_EC2_SecurityGroupIngress__AttributesMap.Id,
	}
)

var (
	// AWS_EC2_SecurityGroupIngress__PropertiesMap reports all the CloudFormation properties for AWS::EC2::SecurityGroupIngress.
	AWS_EC2_SecurityGroupIngress__PropertiesMap = struct {
		CidrIp                     string
		CidrIpv6                   string
		Description                string
		FromPort                   string
		GroupId                    string
		GroupName                  string
		IpProtocol                 string
		SourcePrefixListId         string
		SourceSecurityGroupId      string
		SourceSecurityGroupName    string
		SourceSecurityGroupOwnerId string
		ToPort                     string
	}{
		CidrIp:                     "CidrIp",
		CidrIpv6:                   "CidrIpv6",
		Description:                "Description",
		FromPort:                   "FromPort",
		GroupId:                    "GroupId",
		GroupName:                  "GroupName",
		IpProtocol:                 "IpProtocol",
		SourcePrefixListId:         "SourcePrefixListId",
		SourceSecurityGroupId:      "SourceSecurityGroupId",
		SourceSecurityGroupName:    "SourceSecurityGroupName",
		SourceSecurityGroupOwnerId: "SourceSecurityGroupOwnerId",
		ToPort:                     "ToPort",
	}

	// AWS_EC2_SecurityGroupIngress__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::SecurityGroupIngress.
	AWS_EC2_SecurityGroupIngress__PropertiesSlice = []string{
		AWS_EC2_SecurityGroupIngress__PropertiesMap.CidrIp,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.CidrIpv6,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.Description,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.FromPort,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.GroupId,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.GroupName,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.IpProtocol,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.SourcePrefixListId,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.SourceSecurityGroupId,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.SourceSecurityGroupName,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.SourceSecurityGroupOwnerId,
		AWS_EC2_SecurityGroupIngress__PropertiesMap.ToPort,
	}
)

// AWS_EC2_SecurityGroupIngress is a binding for AWS::EC2::SecurityGroupIngress.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html
type AWS_EC2_SecurityGroupIngress struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CidrIp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-cidrip
	CidrIp cfz.Expression[string] `json:"CidrIp,omitempty"`

	// CidrIpv6 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-cidripv6
	CidrIpv6 cfz.Expression[string] `json:"CidrIpv6,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// FromPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-fromport
	FromPort cfz.Expression[int32] `json:"FromPort,omitempty"`

	// GroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-groupid
	GroupId cfz.Expression[string] `json:"GroupId,omitempty"`

	// GroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-groupname
	GroupName cfz.Expression[string] `json:"GroupName,omitempty"`

	// IpProtocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-ipprotocol
	IpProtocol cfz.Expression[string] `json:"IpProtocol,omitempty"`

	// SourcePrefixListId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-sourceprefixlistid
	SourcePrefixListId cfz.Expression[string] `json:"SourcePrefixListId,omitempty"`

	// SourceSecurityGroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-sourcesecuritygroupid
	SourceSecurityGroupId cfz.Expression[string] `json:"SourceSecurityGroupId,omitempty"`

	// SourceSecurityGroupName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-sourcesecuritygroupname
	SourceSecurityGroupName cfz.Expression[string] `json:"SourceSecurityGroupName,omitempty"`

	// SourceSecurityGroupOwnerId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-sourcesecuritygroupownerid
	SourceSecurityGroupOwnerId cfz.Expression[string] `json:"SourceSecurityGroupOwnerId,omitempty"`

	// ToPort is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupingress.html#cfn-ec2-securitygroupingress-toport
	ToPort cfz.Expression[int32] `json:"ToPort,omitempty"`
}

// New__AWS_EC2_SecurityGroupIngress initializes a new *AWS_EC2_SecurityGroupIngress.
func New__AWS_EC2_SecurityGroupIngress(logicalName string) *AWS_EC2_SecurityGroupIngress {
	return &AWS_EC2_SecurityGroupIngress{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_SecurityGroupIngress) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_SecurityGroupIngress) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_SecurityGroupIngress) GetType() string {
	return AWS_EC2_SecurityGroupIngress__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_SecurityGroupIngress) Set__LogicalName(v string) *AWS_EC2_SecurityGroupIngress {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_SecurityGroupIngress) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_SecurityGroupIngress {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_SecurityGroupIngress) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_SecurityGroupIngress {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_SecurityGroupIngress) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_SecurityGroupIngress {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_SecurityGroupIngress) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_SecurityGroupIngress {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_SecurityGroupIngress) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_SecurityGroupIngress {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_SecurityGroupIngress) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_SecurityGroupIngress {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CidrIp updates property "CidrIp".
func (t *AWS_EC2_SecurityGroupIngress) Set__CidrIp(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.CidrIp = v
	return t
}

// SetV__CidrIp updates property "CidrIp".
func (t *AWS_EC2_SecurityGroupIngress) SetV__CidrIp(v string) *AWS_EC2_SecurityGroupIngress {
	t.CidrIp = cfz.V(v)
	return t
}

// Set__CidrIpv6 updates property "CidrIpv6".
func (t *AWS_EC2_SecurityGroupIngress) Set__CidrIpv6(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.CidrIpv6 = v
	return t
}

// SetV__CidrIpv6 updates property "CidrIpv6".
func (t *AWS_EC2_SecurityGroupIngress) SetV__CidrIpv6(v string) *AWS_EC2_SecurityGroupIngress {
	t.CidrIpv6 = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_EC2_SecurityGroupIngress) Set__Description(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_EC2_SecurityGroupIngress) SetV__Description(v string) *AWS_EC2_SecurityGroupIngress {
	t.Description = cfz.V(v)
	return t
}

// Set__FromPort updates property "FromPort".
func (t *AWS_EC2_SecurityGroupIngress) Set__FromPort(v cfz.Expression[int32]) *AWS_EC2_SecurityGroupIngress {
	t.FromPort = v
	return t
}

// SetV__FromPort updates property "FromPort".
func (t *AWS_EC2_SecurityGroupIngress) SetV__FromPort(v int32) *AWS_EC2_SecurityGroupIngress {
	t.FromPort = cfz.V(v)
	return t
}

// Set__GroupId updates property "GroupId".
func (t *AWS_EC2_SecurityGroupIngress) Set__GroupId(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.GroupId = v
	return t
}

// SetV__GroupId updates property "GroupId".
func (t *AWS_EC2_SecurityGroupIngress) SetV__GroupId(v string) *AWS_EC2_SecurityGroupIngress {
	t.GroupId = cfz.V(v)
	return t
}

// Set__GroupName updates property "GroupName".
func (t *AWS_EC2_SecurityGroupIngress) Set__GroupName(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.GroupName = v
	return t
}

// SetV__GroupName updates property "GroupName".
func (t *AWS_EC2_SecurityGroupIngress) SetV__GroupName(v string) *AWS_EC2_SecurityGroupIngress {
	t.GroupName = cfz.V(v)
	return t
}

// Set__IpProtocol updates property "IpProtocol".
func (t *AWS_EC2_SecurityGroupIngress) Set__IpProtocol(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.IpProtocol = v
	return t
}

// SetV__IpProtocol updates property "IpProtocol".
func (t *AWS_EC2_SecurityGroupIngress) SetV__IpProtocol(v string) *AWS_EC2_SecurityGroupIngress {
	t.IpProtocol = cfz.V(v)
	return t
}

// Set__SourcePrefixListId updates property "SourcePrefixListId".
func (t *AWS_EC2_SecurityGroupIngress) Set__SourcePrefixListId(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.SourcePrefixListId = v
	return t
}

// SetV__SourcePrefixListId updates property "SourcePrefixListId".
func (t *AWS_EC2_SecurityGroupIngress) SetV__SourcePrefixListId(v string) *AWS_EC2_SecurityGroupIngress {
	t.SourcePrefixListId = cfz.V(v)
	return t
}

// Set__SourceSecurityGroupId updates property "SourceSecurityGroupId".
func (t *AWS_EC2_SecurityGroupIngress) Set__SourceSecurityGroupId(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.SourceSecurityGroupId = v
	return t
}

// SetV__SourceSecurityGroupId updates property "SourceSecurityGroupId".
func (t *AWS_EC2_SecurityGroupIngress) SetV__SourceSecurityGroupId(v string) *AWS_EC2_SecurityGroupIngress {
	t.SourceSecurityGroupId = cfz.V(v)
	return t
}

// Set__SourceSecurityGroupName updates property "SourceSecurityGroupName".
func (t *AWS_EC2_SecurityGroupIngress) Set__SourceSecurityGroupName(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.SourceSecurityGroupName = v
	return t
}

// SetV__SourceSecurityGroupName updates property "SourceSecurityGroupName".
func (t *AWS_EC2_SecurityGroupIngress) SetV__SourceSecurityGroupName(v string) *AWS_EC2_SecurityGroupIngress {
	t.SourceSecurityGroupName = cfz.V(v)
	return t
}

// Set__SourceSecurityGroupOwnerId updates property "SourceSecurityGroupOwnerId".
func (t *AWS_EC2_SecurityGroupIngress) Set__SourceSecurityGroupOwnerId(v cfz.Expression[string]) *AWS_EC2_SecurityGroupIngress {
	t.SourceSecurityGroupOwnerId = v
	return t
}

// SetV__SourceSecurityGroupOwnerId updates property "SourceSecurityGroupOwnerId".
func (t *AWS_EC2_SecurityGroupIngress) SetV__SourceSecurityGroupOwnerId(v string) *AWS_EC2_SecurityGroupIngress {
	t.SourceSecurityGroupOwnerId = cfz.V(v)
	return t
}

// Set__ToPort updates property "ToPort".
func (t *AWS_EC2_SecurityGroupIngress) Set__ToPort(v cfz.Expression[int32]) *AWS_EC2_SecurityGroupIngress {
	t.ToPort = v
	return t
}

// SetV__ToPort updates property "ToPort".
func (t *AWS_EC2_SecurityGroupIngress) SetV__ToPort(v int32) *AWS_EC2_SecurityGroupIngress {
	t.ToPort = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_SecurityGroupIngress) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_EC2_SecurityGroupIngress) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_SecurityGroupIngress__AttributesMap.Id))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_SecurityGroupIngress) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_EC2_SecurityGroupIngress) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_SecurityGroupIngress) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_SecurityGroupIngress

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

func (t *AWS_EC2_SecurityGroupIngress) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
