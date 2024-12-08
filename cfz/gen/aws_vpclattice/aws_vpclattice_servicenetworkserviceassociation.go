// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_vpclattice

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_VpcLattice_ServiceNetworkServiceAssociation)(nil)
	_ cfz.Resource                   = (*AWS_VpcLattice_ServiceNetworkServiceAssociation)(nil)
)

const (
	// AWS_VpcLattice_ServiceNetworkServiceAssociation__Type is the CloudFormation type for AWS::VpcLattice::ServiceNetworkServiceAssociation.
	AWS_VpcLattice_ServiceNetworkServiceAssociation__Type = "AWS::VpcLattice::ServiceNetworkServiceAssociation"
)

var (
	// AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap reports all the CloudFormation attributes for AWS::VpcLattice::ServiceNetworkServiceAssociation.
	AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap = struct {
		Arn                   string
		CreatedAt             string
		DnsEntry_DomainName   string
		DnsEntry_HostedZoneId string
		Id                    string
		ServiceArn            string
		ServiceId             string
		ServiceName           string
		ServiceNetworkArn     string
		ServiceNetworkId      string
		ServiceNetworkName    string
		Status                string
	}{
		Arn:                   "Arn",
		CreatedAt:             "CreatedAt",
		DnsEntry_DomainName:   "DnsEntry.DomainName",
		DnsEntry_HostedZoneId: "DnsEntry.HostedZoneId",
		Id:                    "Id",
		ServiceArn:            "ServiceArn",
		ServiceId:             "ServiceId",
		ServiceName:           "ServiceName",
		ServiceNetworkArn:     "ServiceNetworkArn",
		ServiceNetworkId:      "ServiceNetworkId",
		ServiceNetworkName:    "ServiceNetworkName",
		Status:                "Status",
	}

	// AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesSlice reports all the CloudFormation attributes for AWS::VpcLattice::ServiceNetworkServiceAssociation.
	AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesSlice = []string{
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.Arn,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.CreatedAt,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.DnsEntry_DomainName,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.DnsEntry_HostedZoneId,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.Id,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceArn,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceId,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceName,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceNetworkArn,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceNetworkId,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceNetworkName,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.Status,
	}
)

var (
	// AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesMap reports all the CloudFormation properties for AWS::VpcLattice::ServiceNetworkServiceAssociation.
	AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesMap = struct {
		DnsEntry                 string
		ServiceIdentifier        string
		ServiceNetworkIdentifier string
		Tags                     string
	}{
		DnsEntry:                 "DnsEntry",
		ServiceIdentifier:        "ServiceIdentifier",
		ServiceNetworkIdentifier: "ServiceNetworkIdentifier",
		Tags:                     "Tags",
	}

	// AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesSlice reports all the CloudFormation properties for AWS::VpcLattice::ServiceNetworkServiceAssociation.
	AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesSlice = []string{
		AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesMap.DnsEntry,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesMap.ServiceIdentifier,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesMap.ServiceNetworkIdentifier,
		AWS_VpcLattice_ServiceNetworkServiceAssociation__PropertiesMap.Tags,
	}
)

// AWS_VpcLattice_ServiceNetworkServiceAssociation is a binding for AWS::VpcLattice::ServiceNetworkServiceAssociation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetworkserviceassociation.html
type AWS_VpcLattice_ServiceNetworkServiceAssociation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// DnsEntry is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetworkserviceassociation.html#cfn-vpclattice-servicenetworkserviceassociation-dnsentry
	DnsEntry cfz.Expression[AWS_VpcLattice_ServiceNetworkServiceAssociation_DnsEntry] `json:"DnsEntry,omitempty"`

	// ServiceIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetworkserviceassociation.html#cfn-vpclattice-servicenetworkserviceassociation-serviceidentifier
	ServiceIdentifier cfz.Expression[string] `json:"ServiceIdentifier,omitempty"`

	// ServiceNetworkIdentifier is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetworkserviceassociation.html#cfn-vpclattice-servicenetworkserviceassociation-servicenetworkidentifier
	ServiceNetworkIdentifier cfz.Expression[string] `json:"ServiceNetworkIdentifier,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-servicenetworkserviceassociation.html#cfn-vpclattice-servicenetworkserviceassociation-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`
}

// New__AWS_VpcLattice_ServiceNetworkServiceAssociation initializes a new *AWS_VpcLattice_ServiceNetworkServiceAssociation.
func New__AWS_VpcLattice_ServiceNetworkServiceAssociation(logicalName string) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	return &AWS_VpcLattice_ServiceNetworkServiceAssociation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_VpcLattice_ServiceNetworkServiceAssociation) GetType() string {
	return AWS_VpcLattice_ServiceNetworkServiceAssociation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__LogicalName(v string) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__RequestedOutputs(v []cfz.Output) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Add__RequestedOutputs(v ...cfz.Output) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__DnsEntry updates property "DnsEntry".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__DnsEntry(v cfz.Expression[AWS_VpcLattice_ServiceNetworkServiceAssociation_DnsEntry]) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.DnsEntry = v
	return t
}

// SetV__DnsEntry updates property "DnsEntry".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) SetV__DnsEntry(v AWS_VpcLattice_ServiceNetworkServiceAssociation_DnsEntry) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.DnsEntry = cfz.V(v)
	return t
}

// Set__ServiceIdentifier updates property "ServiceIdentifier".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__ServiceIdentifier(v cfz.Expression[string]) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.ServiceIdentifier = v
	return t
}

// SetV__ServiceIdentifier updates property "ServiceIdentifier".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) SetV__ServiceIdentifier(v string) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.ServiceIdentifier = cfz.V(v)
	return t
}

// Set__ServiceNetworkIdentifier updates property "ServiceNetworkIdentifier".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__ServiceNetworkIdentifier(v cfz.Expression[string]) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.ServiceNetworkIdentifier = v
	return t
}

// SetV__ServiceNetworkIdentifier updates property "ServiceNetworkIdentifier".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) SetV__ServiceNetworkIdentifier(v string) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.ServiceNetworkIdentifier = cfz.V(v)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) SetSV__Tags(v ...cfz.Tag) *AWS_VpcLattice_ServiceNetworkServiceAssociation {
	t.Tags = cfz.SV(v...)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.Arn))
}

// GetAtt__CreatedAt returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: CreatedAt
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__CreatedAt() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.CreatedAt))
}

// GetAtt__DnsEntry_DomainName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DnsEntry.DomainName
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__DnsEntry_DomainName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.DnsEntry_DomainName))
}

// GetAtt__DnsEntry_HostedZoneId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: DnsEntry.HostedZoneId
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__DnsEntry_HostedZoneId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.DnsEntry_HostedZoneId))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.Id))
}

// GetAtt__ServiceArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ServiceArn
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__ServiceArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceArn))
}

// GetAtt__ServiceId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ServiceId
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__ServiceId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceId))
}

// GetAtt__ServiceName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ServiceName
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__ServiceName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceName))
}

// GetAtt__ServiceNetworkArn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ServiceNetworkArn
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__ServiceNetworkArn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceNetworkArn))
}

// GetAtt__ServiceNetworkId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ServiceNetworkId
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__ServiceNetworkId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceNetworkId))
}

// GetAtt__ServiceNetworkName returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ServiceNetworkName
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__ServiceNetworkName() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.ServiceNetworkName))
}

// GetAtt__Status returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Status
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetAtt__Status() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_VpcLattice_ServiceNetworkServiceAssociation__AttributesMap.Status))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__CreatedAt returns a conventionally configured output for an attribute of this resource.
// Attribute: CreatedAt
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__CreatedAt(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttCreatedAt", t.GetAtt__CreatedAt())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DnsEntry_DomainName returns a conventionally configured output for an attribute of this resource.
// Attribute: DnsEntry.DomainName
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__DnsEntry_DomainName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDnsEntryDomainName", t.GetAtt__DnsEntry_DomainName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__DnsEntry_HostedZoneId returns a conventionally configured output for an attribute of this resource.
// Attribute: DnsEntry.HostedZoneId
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__DnsEntry_HostedZoneId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttDnsEntryHostedZoneId", t.GetAtt__DnsEntry_HostedZoneId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ServiceArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ServiceArn
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__ServiceArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttServiceArn", t.GetAtt__ServiceArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ServiceId returns a conventionally configured output for an attribute of this resource.
// Attribute: ServiceId
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__ServiceId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttServiceId", t.GetAtt__ServiceId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ServiceName returns a conventionally configured output for an attribute of this resource.
// Attribute: ServiceName
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__ServiceName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttServiceName", t.GetAtt__ServiceName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ServiceNetworkArn returns a conventionally configured output for an attribute of this resource.
// Attribute: ServiceNetworkArn
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__ServiceNetworkArn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttServiceNetworkArn", t.GetAtt__ServiceNetworkArn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ServiceNetworkId returns a conventionally configured output for an attribute of this resource.
// Attribute: ServiceNetworkId
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__ServiceNetworkId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttServiceNetworkId", t.GetAtt__ServiceNetworkId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ServiceNetworkName returns a conventionally configured output for an attribute of this resource.
// Attribute: ServiceNetworkName
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__ServiceNetworkName(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttServiceNetworkName", t.GetAtt__ServiceNetworkName())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Status returns a conventionally configured output for an attribute of this resource.
// Attribute: Status
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) GetConventionalOutputAtt__Status(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttStatus", t.GetAtt__Status())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_VpcLattice_ServiceNetworkServiceAssociation

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

func (t *AWS_VpcLattice_ServiceNetworkServiceAssociation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
