// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_detective

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Detective_MemberInvitation)(nil)
	_ cfz.Resource                   = (*AWS_Detective_MemberInvitation)(nil)
)

const (
	// AWS_Detective_MemberInvitation__Type is the CloudFormation type for AWS::Detective::MemberInvitation.
	AWS_Detective_MemberInvitation__Type = "AWS::Detective::MemberInvitation"
)

var (
	// AWS_Detective_MemberInvitation__PropertiesMap reports all the CloudFormation properties for AWS::Detective::MemberInvitation.
	AWS_Detective_MemberInvitation__PropertiesMap = struct {
		DisableEmailNotification string
		GraphArn                 string
		MemberEmailAddress       string
		MemberId                 string
		Message                  string
	}{
		DisableEmailNotification: "DisableEmailNotification",
		GraphArn:                 "GraphArn",
		MemberEmailAddress:       "MemberEmailAddress",
		MemberId:                 "MemberId",
		Message:                  "Message",
	}

	// AWS_Detective_MemberInvitation__PropertiesSlice reports all the CloudFormation properties for AWS::Detective::MemberInvitation.
	AWS_Detective_MemberInvitation__PropertiesSlice = []string{
		AWS_Detective_MemberInvitation__PropertiesMap.DisableEmailNotification,
		AWS_Detective_MemberInvitation__PropertiesMap.GraphArn,
		AWS_Detective_MemberInvitation__PropertiesMap.MemberEmailAddress,
		AWS_Detective_MemberInvitation__PropertiesMap.MemberId,
		AWS_Detective_MemberInvitation__PropertiesMap.Message,
	}
)

// AWS_Detective_MemberInvitation is a binding for AWS::Detective::MemberInvitation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html
type AWS_Detective_MemberInvitation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DisableEmailNotification is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-disableemailnotification
	DisableEmailNotification cfz.Expression[bool] `json:"DisableEmailNotification,omitempty"`

	// GraphArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-grapharn
	GraphArn cfz.Expression[string] `json:"GraphArn,omitempty"`

	// MemberEmailAddress is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberemailaddress
	MemberEmailAddress cfz.Expression[string] `json:"MemberEmailAddress,omitempty"`

	// MemberId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-memberid
	MemberId cfz.Expression[string] `json:"MemberId,omitempty"`

	// Message is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-detective-memberinvitation.html#cfn-detective-memberinvitation-message
	Message cfz.Expression[string] `json:"Message,omitempty"`
}

// New__AWS_Detective_MemberInvitation initializes a new *AWS_Detective_MemberInvitation.
func New__AWS_Detective_MemberInvitation(logicalName string) *AWS_Detective_MemberInvitation {
	return &AWS_Detective_MemberInvitation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Detective_MemberInvitation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Detective_MemberInvitation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Detective_MemberInvitation) GetType() string {
	return AWS_Detective_MemberInvitation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Detective_MemberInvitation) Set__LogicalName(v string) *AWS_Detective_MemberInvitation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Detective_MemberInvitation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Detective_MemberInvitation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Detective_MemberInvitation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Detective_MemberInvitation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Detective_MemberInvitation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Detective_MemberInvitation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Detective_MemberInvitation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Detective_MemberInvitation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Detective_MemberInvitation) Set__RequestedOutputs(v []cfz.Output) *AWS_Detective_MemberInvitation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Detective_MemberInvitation) Add__RequestedOutputs(v ...cfz.Output) *AWS_Detective_MemberInvitation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DisableEmailNotification updates property "DisableEmailNotification".
func (t *AWS_Detective_MemberInvitation) Set__DisableEmailNotification(v cfz.Expression[bool]) *AWS_Detective_MemberInvitation {
	t.DisableEmailNotification = v
	return t
}

// SetV__DisableEmailNotification updates property "DisableEmailNotification".
func (t *AWS_Detective_MemberInvitation) SetV__DisableEmailNotification(v bool) *AWS_Detective_MemberInvitation {
	t.DisableEmailNotification = cfz.V(v)
	return t
}

// Set__GraphArn updates property "GraphArn".
func (t *AWS_Detective_MemberInvitation) Set__GraphArn(v cfz.Expression[string]) *AWS_Detective_MemberInvitation {
	t.GraphArn = v
	return t
}

// SetV__GraphArn updates property "GraphArn".
func (t *AWS_Detective_MemberInvitation) SetV__GraphArn(v string) *AWS_Detective_MemberInvitation {
	t.GraphArn = cfz.V(v)
	return t
}

// Set__MemberEmailAddress updates property "MemberEmailAddress".
func (t *AWS_Detective_MemberInvitation) Set__MemberEmailAddress(v cfz.Expression[string]) *AWS_Detective_MemberInvitation {
	t.MemberEmailAddress = v
	return t
}

// SetV__MemberEmailAddress updates property "MemberEmailAddress".
func (t *AWS_Detective_MemberInvitation) SetV__MemberEmailAddress(v string) *AWS_Detective_MemberInvitation {
	t.MemberEmailAddress = cfz.V(v)
	return t
}

// Set__MemberId updates property "MemberId".
func (t *AWS_Detective_MemberInvitation) Set__MemberId(v cfz.Expression[string]) *AWS_Detective_MemberInvitation {
	t.MemberId = v
	return t
}

// SetV__MemberId updates property "MemberId".
func (t *AWS_Detective_MemberInvitation) SetV__MemberId(v string) *AWS_Detective_MemberInvitation {
	t.MemberId = cfz.V(v)
	return t
}

// Set__Message updates property "Message".
func (t *AWS_Detective_MemberInvitation) Set__Message(v cfz.Expression[string]) *AWS_Detective_MemberInvitation {
	t.Message = v
	return t
}

// SetV__Message updates property "Message".
func (t *AWS_Detective_MemberInvitation) SetV__Message(v string) *AWS_Detective_MemberInvitation {
	t.Message = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Detective_MemberInvitation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Detective_MemberInvitation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Detective_MemberInvitation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Detective_MemberInvitation

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

func (t *AWS_Detective_MemberInvitation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
