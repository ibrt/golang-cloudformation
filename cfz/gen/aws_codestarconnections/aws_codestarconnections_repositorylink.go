// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_codestarconnections

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_CodeStarConnections_RepositoryLink)(nil)
	_ cfz.Resource                   = (*AWS_CodeStarConnections_RepositoryLink)(nil)
)

const (
	// AWS_CodeStarConnections_RepositoryLink__Type is the CloudFormation type for AWS::CodeStarConnections::RepositoryLink.
	AWS_CodeStarConnections_RepositoryLink__Type = "AWS::CodeStarConnections::RepositoryLink"
)

var (
	// AWS_CodeStarConnections_RepositoryLink__AttributesMap reports all the CloudFormation attributes for AWS::CodeStarConnections::RepositoryLink.
	AWS_CodeStarConnections_RepositoryLink__AttributesMap = struct {
		ProviderType      string
		RepositoryLinkArn string
		RepositoryLinkId  string
	}{
		ProviderType:      "ProviderType",
		RepositoryLinkArn: "RepositoryLinkArn",
		RepositoryLinkId:  "RepositoryLinkId",
	}

	// AWS_CodeStarConnections_RepositoryLink__AttributesSlice reports all the CloudFormation attributes for AWS::CodeStarConnections::RepositoryLink.
	AWS_CodeStarConnections_RepositoryLink__AttributesSlice = []string{
		AWS_CodeStarConnections_RepositoryLink__AttributesMap.ProviderType,
		AWS_CodeStarConnections_RepositoryLink__AttributesMap.RepositoryLinkArn,
		AWS_CodeStarConnections_RepositoryLink__AttributesMap.RepositoryLinkId,
	}
)

var (
	// AWS_CodeStarConnections_RepositoryLink__PropertiesMap reports all the CloudFormation properties for AWS::CodeStarConnections::RepositoryLink.
	AWS_CodeStarConnections_RepositoryLink__PropertiesMap = struct {
		ConnectionArn    string
		EncryptionKeyArn string
		OwnerId          string
		RepositoryName   string
		Tags             string
	}{
		ConnectionArn:    "ConnectionArn",
		EncryptionKeyArn: "EncryptionKeyArn",
		OwnerId:          "OwnerId",
		RepositoryName:   "RepositoryName",
		Tags:             "Tags",
	}

	// AWS_CodeStarConnections_RepositoryLink__PropertiesSlice reports all the CloudFormation properties for AWS::CodeStarConnections::RepositoryLink.
	AWS_CodeStarConnections_RepositoryLink__PropertiesSlice = []string{
		AWS_CodeStarConnections_RepositoryLink__PropertiesMap.ConnectionArn,
		AWS_CodeStarConnections_RepositoryLink__PropertiesMap.EncryptionKeyArn,
		AWS_CodeStarConnections_RepositoryLink__PropertiesMap.OwnerId,
		AWS_CodeStarConnections_RepositoryLink__PropertiesMap.RepositoryName,
		AWS_CodeStarConnections_RepositoryLink__PropertiesMap.Tags,
	}
)

// AWS_CodeStarConnections_RepositoryLink is a binding for AWS::CodeStarConnections::RepositoryLink.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html
type AWS_CodeStarConnections_RepositoryLink struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ConnectionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-connectionarn
	ConnectionArn cfz.Expression[string] `json:"ConnectionArn,omitempty"`

	// EncryptionKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-encryptionkeyarn
	EncryptionKeyArn cfz.Expression[string] `json:"EncryptionKeyArn,omitempty"`

	// OwnerId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-ownerid
	OwnerId cfz.Expression[string] `json:"OwnerId,omitempty"`

	// RepositoryName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-repositoryname
	RepositoryName cfz.Expression[string] `json:"RepositoryName,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarconnections-repositorylink.html#cfn-codestarconnections-repositorylink-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_CodeStarConnections_RepositoryLink initializes a new *AWS_CodeStarConnections_RepositoryLink.
func New__AWS_CodeStarConnections_RepositoryLink(logicalName string) *AWS_CodeStarConnections_RepositoryLink {
	return &AWS_CodeStarConnections_RepositoryLink{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_CodeStarConnections_RepositoryLink) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_CodeStarConnections_RepositoryLink) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_CodeStarConnections_RepositoryLink) GetType() string {
	return AWS_CodeStarConnections_RepositoryLink__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__LogicalName(v string) *AWS_CodeStarConnections_RepositoryLink {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_CodeStarConnections_RepositoryLink {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_CodeStarConnections_RepositoryLink) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_CodeStarConnections_RepositoryLink {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_CodeStarConnections_RepositoryLink {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_CodeStarConnections_RepositoryLink {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__RequestedOutputs(v []cfz.Output) *AWS_CodeStarConnections_RepositoryLink {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_CodeStarConnections_RepositoryLink) Add__RequestedOutputs(v ...cfz.Output) *AWS_CodeStarConnections_RepositoryLink {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ConnectionArn updates property "ConnectionArn".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__ConnectionArn(v cfz.Expression[string]) *AWS_CodeStarConnections_RepositoryLink {
	t.ConnectionArn = v
	return t
}

// SetV__ConnectionArn updates property "ConnectionArn".
func (t *AWS_CodeStarConnections_RepositoryLink) SetV__ConnectionArn(v string) *AWS_CodeStarConnections_RepositoryLink {
	t.ConnectionArn = cfz.V(v)
	return t
}

// Set__EncryptionKeyArn updates property "EncryptionKeyArn".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__EncryptionKeyArn(v cfz.Expression[string]) *AWS_CodeStarConnections_RepositoryLink {
	t.EncryptionKeyArn = v
	return t
}

// SetV__EncryptionKeyArn updates property "EncryptionKeyArn".
func (t *AWS_CodeStarConnections_RepositoryLink) SetV__EncryptionKeyArn(v string) *AWS_CodeStarConnections_RepositoryLink {
	t.EncryptionKeyArn = cfz.V(v)
	return t
}

// Set__OwnerId updates property "OwnerId".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__OwnerId(v cfz.Expression[string]) *AWS_CodeStarConnections_RepositoryLink {
	t.OwnerId = v
	return t
}

// SetV__OwnerId updates property "OwnerId".
func (t *AWS_CodeStarConnections_RepositoryLink) SetV__OwnerId(v string) *AWS_CodeStarConnections_RepositoryLink {
	t.OwnerId = cfz.V(v)
	return t
}

// Set__RepositoryName updates property "RepositoryName".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__RepositoryName(v cfz.Expression[string]) *AWS_CodeStarConnections_RepositoryLink {
	t.RepositoryName = v
	return t
}

// SetV__RepositoryName updates property "RepositoryName".
func (t *AWS_CodeStarConnections_RepositoryLink) SetV__RepositoryName(v string) *AWS_CodeStarConnections_RepositoryLink {
	t.RepositoryName = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_CodeStarConnections_RepositoryLink) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_CodeStarConnections_RepositoryLink {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_CodeStarConnections_RepositoryLink) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_CodeStarConnections_RepositoryLink {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_CodeStarConnections_RepositoryLink) SetSV__Tags(v ...cfz.Tag) *AWS_CodeStarConnections_RepositoryLink {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_CodeStarConnections_RepositoryLink) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__ProviderType returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ProviderType
func (t *AWS_CodeStarConnections_RepositoryLink) GetAtt__ProviderType() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CodeStarConnections_RepositoryLink__AttributesMap.ProviderType))
}

// GetAtt__RepositoryLinkArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RepositoryLinkArn
func (t *AWS_CodeStarConnections_RepositoryLink) GetAtt__RepositoryLinkArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CodeStarConnections_RepositoryLink__AttributesMap.RepositoryLinkArn))
}

// GetAtt__RepositoryLinkId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: RepositoryLinkId
func (t *AWS_CodeStarConnections_RepositoryLink) GetAtt__RepositoryLinkId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_CodeStarConnections_RepositoryLink__AttributesMap.RepositoryLinkId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_CodeStarConnections_RepositoryLink) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ProviderType returns a conventionally configured output for an attribute of this resource.
// Attribute: ProviderType
func (t *AWS_CodeStarConnections_RepositoryLink) GetConventionalOutputAtt__ProviderType(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttProviderType", t.GetAtt__ProviderType())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RepositoryLinkArn returns a conventionally configured output for an attribute of this resource.
// Attribute: RepositoryLinkArn
func (t *AWS_CodeStarConnections_RepositoryLink) GetConventionalOutputAtt__RepositoryLinkArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRepositoryLinkArn", t.GetAtt__RepositoryLinkArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__RepositoryLinkId returns a conventionally configured output for an attribute of this resource.
// Attribute: RepositoryLinkId
func (t *AWS_CodeStarConnections_RepositoryLink) GetConventionalOutputAtt__RepositoryLinkId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttRepositoryLinkId", t.GetAtt__RepositoryLinkId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_CodeStarConnections_RepositoryLink) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_CodeStarConnections_RepositoryLink

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

func (t *AWS_CodeStarConnections_RepositoryLink) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
