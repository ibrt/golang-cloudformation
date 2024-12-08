// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_sagemaker

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_SageMaker_Space)(nil)
	_ cfz.Resource                   = (*AWS_SageMaker_Space)(nil)
)

const (
	// AWS_SageMaker_Space__Type is the CloudFormation type for AWS::SageMaker::Space.
	AWS_SageMaker_Space__Type = "AWS::SageMaker::Space"
)

var (
	// AWS_SageMaker_Space__AttributesMap reports all the CloudFormation attributes for AWS::SageMaker::Space.
	AWS_SageMaker_Space__AttributesMap = struct {
		SpaceArn string
		Url      string
	}{
		SpaceArn: "SpaceArn",
		Url:      "Url",
	}

	// AWS_SageMaker_Space__AttributesSlice reports all the CloudFormation attributes for AWS::SageMaker::Space.
	AWS_SageMaker_Space__AttributesSlice = []string{
		AWS_SageMaker_Space__AttributesMap.SpaceArn,
		AWS_SageMaker_Space__AttributesMap.Url,
	}
)

var (
	// AWS_SageMaker_Space__PropertiesMap reports all the CloudFormation properties for AWS::SageMaker::Space.
	AWS_SageMaker_Space__PropertiesMap = struct {
		DomainId             string
		OwnershipSettings    string
		SpaceDisplayName     string
		SpaceName            string
		SpaceSettings        string
		SpaceSharingSettings string
		Tags                 string
	}{
		DomainId:             "DomainId",
		OwnershipSettings:    "OwnershipSettings",
		SpaceDisplayName:     "SpaceDisplayName",
		SpaceName:            "SpaceName",
		SpaceSettings:        "SpaceSettings",
		SpaceSharingSettings: "SpaceSharingSettings",
		Tags:                 "Tags",
	}

	// AWS_SageMaker_Space__PropertiesSlice reports all the CloudFormation properties for AWS::SageMaker::Space.
	AWS_SageMaker_Space__PropertiesSlice = []string{
		AWS_SageMaker_Space__PropertiesMap.DomainId,
		AWS_SageMaker_Space__PropertiesMap.OwnershipSettings,
		AWS_SageMaker_Space__PropertiesMap.SpaceDisplayName,
		AWS_SageMaker_Space__PropertiesMap.SpaceName,
		AWS_SageMaker_Space__PropertiesMap.SpaceSettings,
		AWS_SageMaker_Space__PropertiesMap.SpaceSharingSettings,
		AWS_SageMaker_Space__PropertiesMap.Tags,
	}
)

// AWS_SageMaker_Space is a binding for AWS::SageMaker::Space.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html
type AWS_SageMaker_Space struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DomainId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-domainid
	DomainId cfz.Expression[string] `json:"DomainId,omitempty"`

	// OwnershipSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-ownershipsettings
	OwnershipSettings cfz.Expression[AWS_SageMaker_Space_OwnershipSettings] `json:"OwnershipSettings,omitempty"`

	// SpaceDisplayName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacedisplayname
	SpaceDisplayName cfz.Expression[string] `json:"SpaceDisplayName,omitempty"`

	// SpaceName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacename
	SpaceName cfz.Expression[string] `json:"SpaceName,omitempty"`

	// SpaceSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacesettings
	SpaceSettings cfz.Expression[AWS_SageMaker_Space_SpaceSettings] `json:"SpaceSettings,omitempty"`

	// SpaceSharingSettings is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-spacesharingsettings
	SpaceSharingSettings cfz.Expression[AWS_SageMaker_Space_SpaceSharingSettings] `json:"SpaceSharingSettings,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-space.html#cfn-sagemaker-space-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_SageMaker_Space initializes a new *AWS_SageMaker_Space.
func New__AWS_SageMaker_Space(logicalName string) *AWS_SageMaker_Space {
	return &AWS_SageMaker_Space{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_SageMaker_Space) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_SageMaker_Space) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_SageMaker_Space) GetType() string {
	return AWS_SageMaker_Space__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_SageMaker_Space) Set__LogicalName(v string) *AWS_SageMaker_Space {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_SageMaker_Space) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_SageMaker_Space {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_SageMaker_Space) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_SageMaker_Space {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_SageMaker_Space) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_SageMaker_Space {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_SageMaker_Space) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_SageMaker_Space {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_SageMaker_Space) Set__RequestedOutputs(v []cfz.Output) *AWS_SageMaker_Space {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_SageMaker_Space) Add__RequestedOutputs(v ...cfz.Output) *AWS_SageMaker_Space {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DomainId updates property "DomainId".
func (t *AWS_SageMaker_Space) Set__DomainId(v cfz.Expression[string]) *AWS_SageMaker_Space {
	t.DomainId = v
	return t
}

// SetV__DomainId updates property "DomainId".
func (t *AWS_SageMaker_Space) SetV__DomainId(v string) *AWS_SageMaker_Space {
	t.DomainId = cfz.V(v)
	return t
}

// Set__OwnershipSettings updates property "OwnershipSettings".
func (t *AWS_SageMaker_Space) Set__OwnershipSettings(v cfz.Expression[AWS_SageMaker_Space_OwnershipSettings]) *AWS_SageMaker_Space {
	t.OwnershipSettings = v
	return t
}

// SetV__OwnershipSettings updates property "OwnershipSettings".
func (t *AWS_SageMaker_Space) SetV__OwnershipSettings(v AWS_SageMaker_Space_OwnershipSettings) *AWS_SageMaker_Space {
	t.OwnershipSettings = cfz.V(v)
	return t
}

// Set__SpaceDisplayName updates property "SpaceDisplayName".
func (t *AWS_SageMaker_Space) Set__SpaceDisplayName(v cfz.Expression[string]) *AWS_SageMaker_Space {
	t.SpaceDisplayName = v
	return t
}

// SetV__SpaceDisplayName updates property "SpaceDisplayName".
func (t *AWS_SageMaker_Space) SetV__SpaceDisplayName(v string) *AWS_SageMaker_Space {
	t.SpaceDisplayName = cfz.V(v)
	return t
}

// Set__SpaceName updates property "SpaceName".
func (t *AWS_SageMaker_Space) Set__SpaceName(v cfz.Expression[string]) *AWS_SageMaker_Space {
	t.SpaceName = v
	return t
}

// SetV__SpaceName updates property "SpaceName".
func (t *AWS_SageMaker_Space) SetV__SpaceName(v string) *AWS_SageMaker_Space {
	t.SpaceName = cfz.V(v)
	return t
}

// Set__SpaceSettings updates property "SpaceSettings".
func (t *AWS_SageMaker_Space) Set__SpaceSettings(v cfz.Expression[AWS_SageMaker_Space_SpaceSettings]) *AWS_SageMaker_Space {
	t.SpaceSettings = v
	return t
}

// SetV__SpaceSettings updates property "SpaceSettings".
func (t *AWS_SageMaker_Space) SetV__SpaceSettings(v AWS_SageMaker_Space_SpaceSettings) *AWS_SageMaker_Space {
	t.SpaceSettings = cfz.V(v)
	return t
}

// Set__SpaceSharingSettings updates property "SpaceSharingSettings".
func (t *AWS_SageMaker_Space) Set__SpaceSharingSettings(v cfz.Expression[AWS_SageMaker_Space_SpaceSharingSettings]) *AWS_SageMaker_Space {
	t.SpaceSharingSettings = v
	return t
}

// SetV__SpaceSharingSettings updates property "SpaceSharingSettings".
func (t *AWS_SageMaker_Space) SetV__SpaceSharingSettings(v AWS_SageMaker_Space_SpaceSharingSettings) *AWS_SageMaker_Space {
	t.SpaceSharingSettings = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_SageMaker_Space) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_SageMaker_Space {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_SageMaker_Space) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_SageMaker_Space {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_SageMaker_Space) SetSV__Tags(v ...cfz.Tag) *AWS_SageMaker_Space {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_SageMaker_Space) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__SpaceArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: SpaceArn
func (t *AWS_SageMaker_Space) GetAtt__SpaceArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_Space__AttributesMap.SpaceArn))
}

// GetAtt__Url returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Url
func (t *AWS_SageMaker_Space) GetAtt__Url() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_SageMaker_Space__AttributesMap.Url))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_SageMaker_Space) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__SpaceArn returns a conventionally configured output for an attribute of this resource.
// Attribute: SpaceArn
func (t *AWS_SageMaker_Space) GetConventionalOutputAtt__SpaceArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttSpaceArn", t.GetAtt__SpaceArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Url returns a conventionally configured output for an attribute of this resource.
// Attribute: Url
func (t *AWS_SageMaker_Space) GetConventionalOutputAtt__Url(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttUrl", t.GetAtt__Url())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_SageMaker_Space) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_SageMaker_Space

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

func (t *AWS_SageMaker_Space) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
