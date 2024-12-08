// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_TransitGatewayVpcAttachment)(nil)
	_ cfz.Resource                   = (*AWS_EC2_TransitGatewayVpcAttachment)(nil)
)

const (
	// AWS_EC2_TransitGatewayVpcAttachment__Type is the CloudFormation type for AWS::EC2::TransitGatewayVpcAttachment.
	AWS_EC2_TransitGatewayVpcAttachment__Type = "AWS::EC2::TransitGatewayVpcAttachment"
)

var (
	// AWS_EC2_TransitGatewayVpcAttachment__AttributesMap reports all the CloudFormation attributes for AWS::EC2::TransitGatewayVpcAttachment.
	AWS_EC2_TransitGatewayVpcAttachment__AttributesMap = struct {
		Id string
	}{
		Id: "Id",
	}

	// AWS_EC2_TransitGatewayVpcAttachment__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::TransitGatewayVpcAttachment.
	AWS_EC2_TransitGatewayVpcAttachment__AttributesSlice = []string{
		AWS_EC2_TransitGatewayVpcAttachment__AttributesMap.Id,
	}
)

var (
	// AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap reports all the CloudFormation properties for AWS::EC2::TransitGatewayVpcAttachment.
	AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap = struct {
		AddSubnetIds     string
		Options          string
		RemoveSubnetIds  string
		SubnetIds        string
		Tags             string
		TransitGatewayId string
		VpcId            string
	}{
		AddSubnetIds:     "AddSubnetIds",
		Options:          "Options",
		RemoveSubnetIds:  "RemoveSubnetIds",
		SubnetIds:        "SubnetIds",
		Tags:             "Tags",
		TransitGatewayId: "TransitGatewayId",
		VpcId:            "VpcId",
	}

	// AWS_EC2_TransitGatewayVpcAttachment__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::TransitGatewayVpcAttachment.
	AWS_EC2_TransitGatewayVpcAttachment__PropertiesSlice = []string{
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.AddSubnetIds,
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.Options,
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.RemoveSubnetIds,
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.SubnetIds,
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.Tags,
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.TransitGatewayId,
		AWS_EC2_TransitGatewayVpcAttachment__PropertiesMap.VpcId,
	}
)

// AWS_EC2_TransitGatewayVpcAttachment is a binding for AWS::EC2::TransitGatewayVpcAttachment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html
type AWS_EC2_TransitGatewayVpcAttachment struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AddSubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-addsubnetids
	AddSubnetIds cfz.ExpressionSlice[string] `json:"AddSubnetIds,omitempty"`

	// Options is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-options
	Options cfz.Expression[AWS_EC2_TransitGatewayVpcAttachment_Options] `json:"Options,omitempty"`

	// RemoveSubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-removesubnetids
	RemoveSubnetIds cfz.ExpressionSlice[string] `json:"RemoveSubnetIds,omitempty"`

	// SubnetIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-subnetids
	SubnetIds cfz.ExpressionSlice[string] `json:"SubnetIds,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TransitGatewayId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-transitgatewayid
	TransitGatewayId cfz.Expression[string] `json:"TransitGatewayId,omitempty"`

	// VpcId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-vpcid
	VpcId cfz.Expression[string] `json:"VpcId,omitempty"`
}

// New__AWS_EC2_TransitGatewayVpcAttachment initializes a new *AWS_EC2_TransitGatewayVpcAttachment.
func New__AWS_EC2_TransitGatewayVpcAttachment(logicalName string) *AWS_EC2_TransitGatewayVpcAttachment {
	return &AWS_EC2_TransitGatewayVpcAttachment{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_TransitGatewayVpcAttachment) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayVpcAttachment) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_TransitGatewayVpcAttachment) GetType() string {
	return AWS_EC2_TransitGatewayVpcAttachment__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__LogicalName(v string) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_TransitGatewayVpcAttachment {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AddSubnetIds updates property "AddSubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__AddSubnetIds(v cfz.ExpressionSlice[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.AddSubnetIds = v
	return t
}

// SetS__AddSubnetIds updates property "AddSubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetS__AddSubnetIds(v ...cfz.Expression[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.AddSubnetIds = cfz.S(v...)
	return t
}

// SetSV__AddSubnetIds updates property "AddSubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetSV__AddSubnetIds(v ...string) *AWS_EC2_TransitGatewayVpcAttachment {
	t.AddSubnetIds = cfz.SV(v...)
	return t
}

// Set__Options updates property "Options".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__Options(v cfz.Expression[AWS_EC2_TransitGatewayVpcAttachment_Options]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.Options = v
	return t
}

// SetV__Options updates property "Options".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetV__Options(v AWS_EC2_TransitGatewayVpcAttachment_Options) *AWS_EC2_TransitGatewayVpcAttachment {
	t.Options = cfz.V(v)
	return t
}

// Set__RemoveSubnetIds updates property "RemoveSubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__RemoveSubnetIds(v cfz.ExpressionSlice[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.RemoveSubnetIds = v
	return t
}

// SetS__RemoveSubnetIds updates property "RemoveSubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetS__RemoveSubnetIds(v ...cfz.Expression[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.RemoveSubnetIds = cfz.S(v...)
	return t
}

// SetSV__RemoveSubnetIds updates property "RemoveSubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetSV__RemoveSubnetIds(v ...string) *AWS_EC2_TransitGatewayVpcAttachment {
	t.RemoveSubnetIds = cfz.SV(v...)
	return t
}

// Set__SubnetIds updates property "SubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__SubnetIds(v cfz.ExpressionSlice[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.SubnetIds = v
	return t
}

// SetS__SubnetIds updates property "SubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetS__SubnetIds(v ...cfz.Expression[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.SubnetIds = cfz.S(v...)
	return t
}

// SetSV__SubnetIds updates property "SubnetIds".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetSV__SubnetIds(v ...string) *AWS_EC2_TransitGatewayVpcAttachment {
	t.SubnetIds = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetSV__Tags(v ...cfz.Tag) *AWS_EC2_TransitGatewayVpcAttachment {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TransitGatewayId updates property "TransitGatewayId".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__TransitGatewayId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.TransitGatewayId = v
	return t
}

// SetV__TransitGatewayId updates property "TransitGatewayId".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetV__TransitGatewayId(v string) *AWS_EC2_TransitGatewayVpcAttachment {
	t.TransitGatewayId = cfz.V(v)
	return t
}

// Set__VpcId updates property "VpcId".
func (t *AWS_EC2_TransitGatewayVpcAttachment) Set__VpcId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayVpcAttachment {
	t.VpcId = v
	return t
}

// SetV__VpcId updates property "VpcId".
func (t *AWS_EC2_TransitGatewayVpcAttachment) SetV__VpcId(v string) *AWS_EC2_TransitGatewayVpcAttachment {
	t.VpcId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_TransitGatewayVpcAttachment) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_EC2_TransitGatewayVpcAttachment) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayVpcAttachment__AttributesMap.Id))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_TransitGatewayVpcAttachment) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_EC2_TransitGatewayVpcAttachment) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayVpcAttachment) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_TransitGatewayVpcAttachment

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

func (t *AWS_EC2_TransitGatewayVpcAttachment) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
