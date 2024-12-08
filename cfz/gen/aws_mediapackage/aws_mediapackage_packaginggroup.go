// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediapackage

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_MediaPackage_PackagingGroup)(nil)
	_ cfz.Resource                   = (*AWS_MediaPackage_PackagingGroup)(nil)
)

const (
	// AWS_MediaPackage_PackagingGroup__Type is the CloudFormation type for AWS::MediaPackage::PackagingGroup.
	AWS_MediaPackage_PackagingGroup__Type = "AWS::MediaPackage::PackagingGroup"
)

var (
	// AWS_MediaPackage_PackagingGroup__AttributesMap reports all the CloudFormation attributes for AWS::MediaPackage::PackagingGroup.
	AWS_MediaPackage_PackagingGroup__AttributesMap = struct {
		Arn        string
		DomainName string
	}{
		Arn:        "Arn",
		DomainName: "DomainName",
	}

	// AWS_MediaPackage_PackagingGroup__AttributesSlice reports all the CloudFormation attributes for AWS::MediaPackage::PackagingGroup.
	AWS_MediaPackage_PackagingGroup__AttributesSlice = []string{
		AWS_MediaPackage_PackagingGroup__AttributesMap.Arn,
		AWS_MediaPackage_PackagingGroup__AttributesMap.DomainName,
	}
)

var (
	// AWS_MediaPackage_PackagingGroup__PropertiesMap reports all the CloudFormation properties for AWS::MediaPackage::PackagingGroup.
	AWS_MediaPackage_PackagingGroup__PropertiesMap = struct {
		Authorization    string
		EgressAccessLogs string
		Id               string
		Tags             string
	}{
		Authorization:    "Authorization",
		EgressAccessLogs: "EgressAccessLogs",
		Id:               "Id",
		Tags:             "Tags",
	}

	// AWS_MediaPackage_PackagingGroup__PropertiesSlice reports all the CloudFormation properties for AWS::MediaPackage::PackagingGroup.
	AWS_MediaPackage_PackagingGroup__PropertiesSlice = []string{
		AWS_MediaPackage_PackagingGroup__PropertiesMap.Authorization,
		AWS_MediaPackage_PackagingGroup__PropertiesMap.EgressAccessLogs,
		AWS_MediaPackage_PackagingGroup__PropertiesMap.Id,
		AWS_MediaPackage_PackagingGroup__PropertiesMap.Tags,
	}
)

// AWS_MediaPackage_PackagingGroup is a binding for AWS::MediaPackage::PackagingGroup.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html
type AWS_MediaPackage_PackagingGroup struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Authorization is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-authorization
	Authorization cfz.Expression[AWS_MediaPackage_PackagingGroup_Authorization] `json:"Authorization,omitempty"`

	// EgressAccessLogs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-egressaccesslogs
	EgressAccessLogs cfz.Expression[AWS_MediaPackage_PackagingGroup_LogConfiguration] `json:"EgressAccessLogs,omitempty"`

	// Id is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-id
	Id cfz.Expression[string] `json:"Id,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html#cfn-mediapackage-packaginggroup-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_MediaPackage_PackagingGroup initializes a new *AWS_MediaPackage_PackagingGroup.
func New__AWS_MediaPackage_PackagingGroup(logicalName string) *AWS_MediaPackage_PackagingGroup {
	return &AWS_MediaPackage_PackagingGroup{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_MediaPackage_PackagingGroup) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_MediaPackage_PackagingGroup) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_MediaPackage_PackagingGroup) GetType() string {
	return AWS_MediaPackage_PackagingGroup__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_MediaPackage_PackagingGroup) Set__LogicalName(v string) *AWS_MediaPackage_PackagingGroup {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_MediaPackage_PackagingGroup) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_MediaPackage_PackagingGroup {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_MediaPackage_PackagingGroup) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_MediaPackage_PackagingGroup {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_MediaPackage_PackagingGroup) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_MediaPackage_PackagingGroup {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_MediaPackage_PackagingGroup) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_MediaPackage_PackagingGroup {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_MediaPackage_PackagingGroup) Set__RequestedOutputs(v []cfz.Output) *AWS_MediaPackage_PackagingGroup {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_MediaPackage_PackagingGroup) Add__RequestedOutputs(v ...cfz.Output) *AWS_MediaPackage_PackagingGroup {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Authorization updates property "Authorization".
func (t *AWS_MediaPackage_PackagingGroup) Set__Authorization(v cfz.Expression[AWS_MediaPackage_PackagingGroup_Authorization]) *AWS_MediaPackage_PackagingGroup {
	t.Authorization = v
	return t
}

// SetV__Authorization updates property "Authorization".
func (t *AWS_MediaPackage_PackagingGroup) SetV__Authorization(v AWS_MediaPackage_PackagingGroup_Authorization) *AWS_MediaPackage_PackagingGroup {
	t.Authorization = cfz.V(v)
	return t
}

// Set__EgressAccessLogs updates property "EgressAccessLogs".
func (t *AWS_MediaPackage_PackagingGroup) Set__EgressAccessLogs(v cfz.Expression[AWS_MediaPackage_PackagingGroup_LogConfiguration]) *AWS_MediaPackage_PackagingGroup {
	t.EgressAccessLogs = v
	return t
}

// SetV__EgressAccessLogs updates property "EgressAccessLogs".
func (t *AWS_MediaPackage_PackagingGroup) SetV__EgressAccessLogs(v AWS_MediaPackage_PackagingGroup_LogConfiguration) *AWS_MediaPackage_PackagingGroup {
	t.EgressAccessLogs = cfz.V(v)
	return t
}

// Set__Id updates property "Id".
func (t *AWS_MediaPackage_PackagingGroup) Set__Id(v cfz.Expression[string]) *AWS_MediaPackage_PackagingGroup {
	t.Id = v
	return t
}

// SetV__Id updates property "Id".
func (t *AWS_MediaPackage_PackagingGroup) SetV__Id(v string) *AWS_MediaPackage_PackagingGroup {
	t.Id = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_MediaPackage_PackagingGroup) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_MediaPackage_PackagingGroup {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_MediaPackage_PackagingGroup) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_MediaPackage_PackagingGroup {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_MediaPackage_PackagingGroup) SetSV__Tags(v ...cfz.Tag) *AWS_MediaPackage_PackagingGroup {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_MediaPackage_PackagingGroup) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_MediaPackage_PackagingGroup) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaPackage_PackagingGroup__AttributesMap.Arn))
}

// GetAtt__DomainName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DomainName
func (t *AWS_MediaPackage_PackagingGroup) GetAtt__DomainName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_MediaPackage_PackagingGroup__AttributesMap.DomainName))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_MediaPackage_PackagingGroup) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_MediaPackage_PackagingGroup) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DomainName returns a conventionally configured output for an attribute of this resource.
// Attribute: DomainName
func (t *AWS_MediaPackage_PackagingGroup) GetConventionalOutputAtt__DomainName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDomainName", t.GetAtt__DomainName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_MediaPackage_PackagingGroup) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_MediaPackage_PackagingGroup

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

func (t *AWS_MediaPackage_PackagingGroup) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
