// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_workspacesweb

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_WorkSpacesWeb_TrustStore)(nil)
	_ cfz.Resource                   = (*AWS_WorkSpacesWeb_TrustStore)(nil)
)

const (
	// AWS_WorkSpacesWeb_TrustStore__Type is the CloudFormation type for AWS::WorkSpacesWeb::TrustStore.
	AWS_WorkSpacesWeb_TrustStore__Type = "AWS::WorkSpacesWeb::TrustStore"
)

var (
	// AWS_WorkSpacesWeb_TrustStore__AttributesMap reports all the CloudFormation attributes for AWS::WorkSpacesWeb::TrustStore.
	AWS_WorkSpacesWeb_TrustStore__AttributesMap = struct {
		AssociatedPortalArns string
		TrustStoreArn        string
	}{
		AssociatedPortalArns: "AssociatedPortalArns",
		TrustStoreArn:        "TrustStoreArn",
	}

	// AWS_WorkSpacesWeb_TrustStore__AttributesSlice reports all the CloudFormation attributes for AWS::WorkSpacesWeb::TrustStore.
	AWS_WorkSpacesWeb_TrustStore__AttributesSlice = []string{
		AWS_WorkSpacesWeb_TrustStore__AttributesMap.AssociatedPortalArns,
		AWS_WorkSpacesWeb_TrustStore__AttributesMap.TrustStoreArn,
	}
)

var (
	// AWS_WorkSpacesWeb_TrustStore__PropertiesMap reports all the CloudFormation properties for AWS::WorkSpacesWeb::TrustStore.
	AWS_WorkSpacesWeb_TrustStore__PropertiesMap = struct {
		CertificateList string
		Tags            string
	}{
		CertificateList: "CertificateList",
		Tags:            "Tags",
	}

	// AWS_WorkSpacesWeb_TrustStore__PropertiesSlice reports all the CloudFormation properties for AWS::WorkSpacesWeb::TrustStore.
	AWS_WorkSpacesWeb_TrustStore__PropertiesSlice = []string{
		AWS_WorkSpacesWeb_TrustStore__PropertiesMap.CertificateList,
		AWS_WorkSpacesWeb_TrustStore__PropertiesMap.Tags,
	}
)

// AWS_WorkSpacesWeb_TrustStore is a binding for AWS::WorkSpacesWeb::TrustStore.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-truststore.html
type AWS_WorkSpacesWeb_TrustStore struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CertificateList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-truststore.html#cfn-workspacesweb-truststore-certificatelist
	CertificateList cfz.ExpressionSlice[string] `json:"CertificateList,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-truststore.html#cfn-workspacesweb-truststore-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_WorkSpacesWeb_TrustStore initializes a new *AWS_WorkSpacesWeb_TrustStore.
func New__AWS_WorkSpacesWeb_TrustStore(logicalName string) *AWS_WorkSpacesWeb_TrustStore {
	return &AWS_WorkSpacesWeb_TrustStore{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_WorkSpacesWeb_TrustStore) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_WorkSpacesWeb_TrustStore) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_WorkSpacesWeb_TrustStore) GetType() string {
	return AWS_WorkSpacesWeb_TrustStore__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__LogicalName(v string) *AWS_WorkSpacesWeb_TrustStore {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_WorkSpacesWeb_TrustStore {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_WorkSpacesWeb_TrustStore) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_WorkSpacesWeb_TrustStore {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_WorkSpacesWeb_TrustStore {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_WorkSpacesWeb_TrustStore {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__RequestedOutputs(v []cfz.Output) *AWS_WorkSpacesWeb_TrustStore {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_WorkSpacesWeb_TrustStore) Add__RequestedOutputs(v ...cfz.Output) *AWS_WorkSpacesWeb_TrustStore {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CertificateList updates property "CertificateList".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__CertificateList(v cfz.ExpressionSlice[string]) *AWS_WorkSpacesWeb_TrustStore {
	t.CertificateList = v
	return t
}

// SetS__CertificateList updates property "CertificateList".
func (t *AWS_WorkSpacesWeb_TrustStore) SetS__CertificateList(v ...cfz.Expression[string]) *AWS_WorkSpacesWeb_TrustStore {
	t.CertificateList = cfz.S(v...)
	return t
}

// SetSV__CertificateList updates property "CertificateList".
func (t *AWS_WorkSpacesWeb_TrustStore) SetSV__CertificateList(v ...string) *AWS_WorkSpacesWeb_TrustStore {
	t.CertificateList = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_WorkSpacesWeb_TrustStore) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_WorkSpacesWeb_TrustStore {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_WorkSpacesWeb_TrustStore) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_WorkSpacesWeb_TrustStore {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_WorkSpacesWeb_TrustStore) SetSV__Tags(v ...cfz.Tag) *AWS_WorkSpacesWeb_TrustStore {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_WorkSpacesWeb_TrustStore) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAttSlice__AssociatedPortalArns returns a $cfz.ExpressionSlice[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AssociatedPortalArns
func (t *AWS_WorkSpacesWeb_TrustStore) GetAttSlice__AssociatedPortalArns() cfz.ExpressionSlice[string] {
	return cfz.GetAttSlice[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesWeb_TrustStore__AttributesMap.AssociatedPortalArns))
}

// GetAtt__TrustStoreArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: TrustStoreArn
func (t *AWS_WorkSpacesWeb_TrustStore) GetAtt__TrustStoreArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WorkSpacesWeb_TrustStore__AttributesMap.TrustStoreArn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_WorkSpacesWeb_TrustStore) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AssociatedPortalArns returns a conventionally configured output for an attribute of this resource.
// Attribute: AssociatedPortalArns
func (t *AWS_WorkSpacesWeb_TrustStore) GetConventionalOutputAtt__AssociatedPortalArns(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAssociatedPortalArns", t.GetAttSlice__AssociatedPortalArns())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__TrustStoreArn returns a conventionally configured output for an attribute of this resource.
// Attribute: TrustStoreArn
func (t *AWS_WorkSpacesWeb_TrustStore) GetConventionalOutputAtt__TrustStoreArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttTrustStoreArn", t.GetAtt__TrustStoreArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_WorkSpacesWeb_TrustStore) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_WorkSpacesWeb_TrustStore

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

func (t *AWS_WorkSpacesWeb_TrustStore) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
