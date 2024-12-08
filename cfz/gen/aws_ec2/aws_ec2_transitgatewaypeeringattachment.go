// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_EC2_TransitGatewayPeeringAttachment)(nil)
	_ cfz.Resource                   = (*AWS_EC2_TransitGatewayPeeringAttachment)(nil)
)

const (
	// AWS_EC2_TransitGatewayPeeringAttachment__Type is the CloudFormation type for AWS::EC2::TransitGatewayPeeringAttachment.
	AWS_EC2_TransitGatewayPeeringAttachment__Type = "AWS::EC2::TransitGatewayPeeringAttachment"
)

var (
	// AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap reports all the CloudFormation attributes for AWS::EC2::TransitGatewayPeeringAttachment.
	AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap = struct {
		CreationTime               string
		State                      string
		Status                     string
		Status_Code                string
		Status_Message             string
		TransitGatewayAttachmentId string
	}{
		CreationTime:               "CreationTime",
		State:                      "State",
		Status:                     "Status",
		Status_Code:                "Status.Code",
		Status_Message:             "Status.Message",
		TransitGatewayAttachmentId: "TransitGatewayAttachmentId",
	}

	// AWS_EC2_TransitGatewayPeeringAttachment__AttributesSlice reports all the CloudFormation attributes for AWS::EC2::TransitGatewayPeeringAttachment.
	AWS_EC2_TransitGatewayPeeringAttachment__AttributesSlice = []string{
		AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.CreationTime,
		AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.State,
		AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.Status,
		AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.Status_Code,
		AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.Status_Message,
		AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.TransitGatewayAttachmentId,
	}
)

var (
	// AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap reports all the CloudFormation properties for AWS::EC2::TransitGatewayPeeringAttachment.
	AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap = struct {
		PeerAccountId        string
		PeerRegion           string
		PeerTransitGatewayId string
		Tags                 string
		TransitGatewayId     string
	}{
		PeerAccountId:        "PeerAccountId",
		PeerRegion:           "PeerRegion",
		PeerTransitGatewayId: "PeerTransitGatewayId",
		Tags:                 "Tags",
		TransitGatewayId:     "TransitGatewayId",
	}

	// AWS_EC2_TransitGatewayPeeringAttachment__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::TransitGatewayPeeringAttachment.
	AWS_EC2_TransitGatewayPeeringAttachment__PropertiesSlice = []string{
		AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap.PeerAccountId,
		AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap.PeerRegion,
		AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap.PeerTransitGatewayId,
		AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap.Tags,
		AWS_EC2_TransitGatewayPeeringAttachment__PropertiesMap.TransitGatewayId,
	}
)

// AWS_EC2_TransitGatewayPeeringAttachment is a binding for AWS::EC2::TransitGatewayPeeringAttachment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html
type AWS_EC2_TransitGatewayPeeringAttachment struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// PeerAccountId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-peeraccountid
	PeerAccountId cfz.Expression[string] `json:"PeerAccountId,omitempty"`

	// PeerRegion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-peerregion
	PeerRegion cfz.Expression[string] `json:"PeerRegion,omitempty"`

	// PeerTransitGatewayId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-peertransitgatewayid
	PeerTransitGatewayId cfz.Expression[string] `json:"PeerTransitGatewayId,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TransitGatewayId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html#cfn-ec2-transitgatewaypeeringattachment-transitgatewayid
	TransitGatewayId cfz.Expression[string] `json:"TransitGatewayId,omitempty"`
}

// New__AWS_EC2_TransitGatewayPeeringAttachment initializes a new *AWS_EC2_TransitGatewayPeeringAttachment.
func New__AWS_EC2_TransitGatewayPeeringAttachment(logicalName string) *AWS_EC2_TransitGatewayPeeringAttachment {
	return &AWS_EC2_TransitGatewayPeeringAttachment{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_EC2_TransitGatewayPeeringAttachment) GetType() string {
	return AWS_EC2_TransitGatewayPeeringAttachment__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__LogicalName(v string) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__RequestedOutputs(v []cfz.Output) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Add__RequestedOutputs(v ...cfz.Output) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__PeerAccountId updates property "PeerAccountId".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__PeerAccountId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.PeerAccountId = v
	return t
}

// SetV__PeerAccountId updates property "PeerAccountId".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) SetV__PeerAccountId(v string) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.PeerAccountId = cfz.V(v)
	return t
}

// Set__PeerRegion updates property "PeerRegion".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__PeerRegion(v cfz.Expression[string]) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.PeerRegion = v
	return t
}

// SetV__PeerRegion updates property "PeerRegion".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) SetV__PeerRegion(v string) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.PeerRegion = cfz.V(v)
	return t
}

// Set__PeerTransitGatewayId updates property "PeerTransitGatewayId".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__PeerTransitGatewayId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.PeerTransitGatewayId = v
	return t
}

// SetV__PeerTransitGatewayId updates property "PeerTransitGatewayId".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) SetV__PeerTransitGatewayId(v string) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.PeerTransitGatewayId = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) SetSV__Tags(v ...cfz.Tag) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TransitGatewayId updates property "TransitGatewayId".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Set__TransitGatewayId(v cfz.Expression[string]) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.TransitGatewayId = v
	return t
}

// SetV__TransitGatewayId updates property "TransitGatewayId".
func (t *AWS_EC2_TransitGatewayPeeringAttachment) SetV__TransitGatewayId(v string) *AWS_EC2_TransitGatewayPeeringAttachment {
	t.TransitGatewayId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_EC2_TransitGatewayPeeringAttachment) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__CreationTime returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreationTime
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetAtt__CreationTime() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.CreationTime))
}

// GetAtt__State returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: State
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetAtt__State() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.State))
}

// GetAtt__Status returns a $cfz.Expression[AWS_EC2_TransitGatewayPeeringAttachment_PeeringAttachmentStatus] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetAtt__Status() cfz.Expression[AWS_EC2_TransitGatewayPeeringAttachment_PeeringAttachmentStatus] {
	return cfz.GetAtt[AWS_EC2_TransitGatewayPeeringAttachment_PeeringAttachmentStatus](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.Status))
}

// GetAtt__Status_Code returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status.Code
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetAtt__Status_Code() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.Status_Code))
}

// GetAtt__Status_Message returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status.Message
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetAtt__Status_Message() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.Status_Message))
}

// GetAtt__TransitGatewayAttachmentId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TransitGatewayAttachmentId
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetAtt__TransitGatewayAttachmentId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_EC2_TransitGatewayPeeringAttachment__AttributesMap.TransitGatewayAttachmentId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreationTime returns a conventionally configured output for an attribute of this resource.
// Attribute: CreationTime
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputAtt__CreationTime(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreationTime", t.GetAtt__CreationTime())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__State returns a conventionally configured output for an attribute of this resource.
// Attribute: State
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputAtt__State(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttState", t.GetAtt__State())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status_Code returns a conventionally configured output for an attribute of this resource.
// Attribute: Status.Code
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputAtt__Status_Code(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatusCode", t.GetAtt__Status_Code())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status_Message returns a conventionally configured output for an attribute of this resource.
// Attribute: Status.Message
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputAtt__Status_Message(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatusMessage", t.GetAtt__Status_Message())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TransitGatewayAttachmentId returns a conventionally configured output for an attribute of this resource.
// Attribute: TransitGatewayAttachmentId
func (t *AWS_EC2_TransitGatewayPeeringAttachment) GetConventionalOutputAtt__TransitGatewayAttachmentId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTransitGatewayAttachmentId", t.GetAtt__TransitGatewayAttachmentId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_EC2_TransitGatewayPeeringAttachment) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_EC2_TransitGatewayPeeringAttachment

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

func (t *AWS_EC2_TransitGatewayPeeringAttachment) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
