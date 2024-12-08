// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appsync

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_AppSync_DomainName)(nil)
	_ cfz.Resource                   = (*AWS_AppSync_DomainName)(nil)
)

const (
	// AWS_AppSync_DomainName__Type is the CloudFormation type for AWS::AppSync::DomainName.
	AWS_AppSync_DomainName__Type = "AWS::AppSync::DomainName"
)

var (
	// AWS_AppSync_DomainName__AttributesMap reports all the CloudFormation attributes for AWS::AppSync::DomainName.
	AWS_AppSync_DomainName__AttributesMap = struct {
		AppSyncDomainName string
		DomainName        string
		HostedZoneId      string
	}{
		AppSyncDomainName: "AppSyncDomainName",
		DomainName:        "DomainName",
		HostedZoneId:      "HostedZoneId",
	}

	// AWS_AppSync_DomainName__AttributesSlice reports all the CloudFormation attributes for AWS::AppSync::DomainName.
	AWS_AppSync_DomainName__AttributesSlice = []string{
		AWS_AppSync_DomainName__AttributesMap.AppSyncDomainName,
		AWS_AppSync_DomainName__AttributesMap.DomainName,
		AWS_AppSync_DomainName__AttributesMap.HostedZoneId,
	}
)

var (
	// AWS_AppSync_DomainName__PropertiesMap reports all the CloudFormation properties for AWS::AppSync::DomainName.
	AWS_AppSync_DomainName__PropertiesMap = struct {
		CertificateArn string
		Description    string
		DomainName     string
	}{
		CertificateArn: "CertificateArn",
		Description:    "Description",
		DomainName:     "DomainName",
	}

	// AWS_AppSync_DomainName__PropertiesSlice reports all the CloudFormation properties for AWS::AppSync::DomainName.
	AWS_AppSync_DomainName__PropertiesSlice = []string{
		AWS_AppSync_DomainName__PropertiesMap.CertificateArn,
		AWS_AppSync_DomainName__PropertiesMap.Description,
		AWS_AppSync_DomainName__PropertiesMap.DomainName,
	}
)

// AWS_AppSync_DomainName is a binding for AWS::AppSync::DomainName.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html
type AWS_AppSync_DomainName struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// CertificateArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html#cfn-appsync-domainname-certificatearn
	CertificateArn cfz.Expression[string] `json:"CertificateArn,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html#cfn-appsync-domainname-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// DomainName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-domainname.html#cfn-appsync-domainname-domainname
	DomainName cfz.Expression[string] `json:"DomainName,omitempty"`
}

// New__AWS_AppSync_DomainName initializes a new *AWS_AppSync_DomainName.
func New__AWS_AppSync_DomainName(logicalName string) *AWS_AppSync_DomainName {
	return &AWS_AppSync_DomainName{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_AppSync_DomainName) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_AppSync_DomainName) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_AppSync_DomainName) GetType() string {
	return AWS_AppSync_DomainName__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_AppSync_DomainName) Set__LogicalName(v string) *AWS_AppSync_DomainName {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_AppSync_DomainName) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_AppSync_DomainName {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_AppSync_DomainName) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_AppSync_DomainName {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_AppSync_DomainName) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_AppSync_DomainName {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_AppSync_DomainName) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_AppSync_DomainName {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_AppSync_DomainName) Set__RequestedOutputs(v []cfz.Output) *AWS_AppSync_DomainName {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_AppSync_DomainName) Add__RequestedOutputs(v ...cfz.Output) *AWS_AppSync_DomainName {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__CertificateArn updates property "CertificateArn".
func (t *AWS_AppSync_DomainName) Set__CertificateArn(v cfz.Expression[string]) *AWS_AppSync_DomainName {
	t.CertificateArn = v
	return t
}

// SetV__CertificateArn updates property "CertificateArn".
func (t *AWS_AppSync_DomainName) SetV__CertificateArn(v string) *AWS_AppSync_DomainName {
	t.CertificateArn = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t *AWS_AppSync_DomainName) Set__Description(v cfz.Expression[string]) *AWS_AppSync_DomainName {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t *AWS_AppSync_DomainName) SetV__Description(v string) *AWS_AppSync_DomainName {
	t.Description = cfz.V(v)
	return t
}

// Set__DomainName updates property "DomainName".
func (t *AWS_AppSync_DomainName) Set__DomainName(v cfz.Expression[string]) *AWS_AppSync_DomainName {
	t.DomainName = v
	return t
}

// SetV__DomainName updates property "DomainName".
func (t *AWS_AppSync_DomainName) SetV__DomainName(v string) *AWS_AppSync_DomainName {
	t.DomainName = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_AppSync_DomainName) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__AppSyncDomainName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: AppSyncDomainName
func (t *AWS_AppSync_DomainName) GetAtt__AppSyncDomainName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppSync_DomainName__AttributesMap.AppSyncDomainName))
}

// GetAtt__DomainName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DomainName
func (t *AWS_AppSync_DomainName) GetAtt__DomainName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppSync_DomainName__AttributesMap.DomainName))
}

// GetAtt__HostedZoneId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: HostedZoneId
func (t *AWS_AppSync_DomainName) GetAtt__HostedZoneId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_AppSync_DomainName__AttributesMap.HostedZoneId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_AppSync_DomainName) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__AppSyncDomainName returns a conventionally configured output for an attribute of this resource.
// Attribute: AppSyncDomainName
func (t *AWS_AppSync_DomainName) GetConventionalOutputAtt__AppSyncDomainName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttAppSyncDomainName", t.GetAtt__AppSyncDomainName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DomainName returns a conventionally configured output for an attribute of this resource.
// Attribute: DomainName
func (t *AWS_AppSync_DomainName) GetConventionalOutputAtt__DomainName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDomainName", t.GetAtt__DomainName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__HostedZoneId returns a conventionally configured output for an attribute of this resource.
// Attribute: HostedZoneId
func (t *AWS_AppSync_DomainName) GetConventionalOutputAtt__HostedZoneId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttHostedZoneId", t.GetAtt__HostedZoneId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_AppSync_DomainName) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_AppSync_DomainName

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

func (t *AWS_AppSync_DomainName) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
