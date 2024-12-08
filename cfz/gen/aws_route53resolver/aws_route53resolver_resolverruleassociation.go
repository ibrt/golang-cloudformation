// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53resolver

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Route53Resolver_ResolverRuleAssociation)(nil)
	_ cfz.Resource                   = (*AWS_Route53Resolver_ResolverRuleAssociation)(nil)
)

const (
	// AWS_Route53Resolver_ResolverRuleAssociation__Type is the CloudFormation type for AWS::Route53Resolver::ResolverRuleAssociation.
	AWS_Route53Resolver_ResolverRuleAssociation__Type = "AWS::Route53Resolver::ResolverRuleAssociation"
)

var (
	// AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap reports all the CloudFormation attributes for AWS::Route53Resolver::ResolverRuleAssociation.
	AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap = struct {
		Name                      string
		ResolverRuleAssociationId string
		ResolverRuleId            string
		VPCId                     string
	}{
		Name:                      "Name",
		ResolverRuleAssociationId: "ResolverRuleAssociationId",
		ResolverRuleId:            "ResolverRuleId",
		VPCId:                     "VPCId",
	}

	// AWS_Route53Resolver_ResolverRuleAssociation__AttributesSlice reports all the CloudFormation attributes for AWS::Route53Resolver::ResolverRuleAssociation.
	AWS_Route53Resolver_ResolverRuleAssociation__AttributesSlice = []string{
		AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.Name,
		AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.ResolverRuleAssociationId,
		AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.ResolverRuleId,
		AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.VPCId,
	}
)

var (
	// AWS_Route53Resolver_ResolverRuleAssociation__PropertiesMap reports all the CloudFormation properties for AWS::Route53Resolver::ResolverRuleAssociation.
	AWS_Route53Resolver_ResolverRuleAssociation__PropertiesMap = struct {
		Name           string
		ResolverRuleId string
		VPCId          string
	}{
		Name:           "Name",
		ResolverRuleId: "ResolverRuleId",
		VPCId:          "VPCId",
	}

	// AWS_Route53Resolver_ResolverRuleAssociation__PropertiesSlice reports all the CloudFormation properties for AWS::Route53Resolver::ResolverRuleAssociation.
	AWS_Route53Resolver_ResolverRuleAssociation__PropertiesSlice = []string{
		AWS_Route53Resolver_ResolverRuleAssociation__PropertiesMap.Name,
		AWS_Route53Resolver_ResolverRuleAssociation__PropertiesMap.ResolverRuleId,
		AWS_Route53Resolver_ResolverRuleAssociation__PropertiesMap.VPCId,
	}
)

// AWS_Route53Resolver_ResolverRuleAssociation is a binding for AWS::Route53Resolver::ResolverRuleAssociation.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html
type AWS_Route53Resolver_ResolverRuleAssociation struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ResolverRuleId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-resolverruleid
	ResolverRuleId cfz.Expression[string] `json:"ResolverRuleId,omitempty"`

	// VPCId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html#cfn-route53resolver-resolverruleassociation-vpcid
	VPCId cfz.Expression[string] `json:"VPCId,omitempty"`
}

// New__AWS_Route53Resolver_ResolverRuleAssociation initializes a new *AWS_Route53Resolver_ResolverRuleAssociation.
func New__AWS_Route53Resolver_ResolverRuleAssociation(logicalName string) *AWS_Route53Resolver_ResolverRuleAssociation {
	return &AWS_Route53Resolver_ResolverRuleAssociation{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Route53Resolver_ResolverRuleAssociation) GetType() string {
	return AWS_Route53Resolver_ResolverRuleAssociation__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__LogicalName(v string) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__RequestedOutputs(v []cfz.Output) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Add__RequestedOutputs(v ...cfz.Output) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__Name updates property "Name".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__Name(v cfz.Expression[string]) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) SetV__Name(v string) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.Name = cfz.V(v)
	return t
}

// Set__ResolverRuleId updates property "ResolverRuleId".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__ResolverRuleId(v cfz.Expression[string]) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.ResolverRuleId = v
	return t
}

// SetV__ResolverRuleId updates property "ResolverRuleId".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) SetV__ResolverRuleId(v string) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.ResolverRuleId = cfz.V(v)
	return t
}

// Set__VPCId updates property "VPCId".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Set__VPCId(v cfz.Expression[string]) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.VPCId = v
	return t
}

// SetV__VPCId updates property "VPCId".
func (t *AWS_Route53Resolver_ResolverRuleAssociation) SetV__VPCId(v string) *AWS_Route53Resolver_ResolverRuleAssociation {
	t.VPCId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Route53Resolver_ResolverRuleAssociation) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Name returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Name
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetAtt__Name() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.Name))
}

// GetAtt__ResolverRuleAssociationId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ResolverRuleAssociationId
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetAtt__ResolverRuleAssociationId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.ResolverRuleAssociationId))
}

// GetAtt__ResolverRuleId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ResolverRuleId
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetAtt__ResolverRuleId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.ResolverRuleId))
}

// GetAtt__VPCId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: VPCId
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetAtt__VPCId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverRuleAssociation__AttributesMap.VPCId))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Name returns a conventionally configured output for an attribute of this resource.
// Attribute: Name
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetConventionalOutputAtt__Name(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttName", t.GetAtt__Name())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ResolverRuleAssociationId returns a conventionally configured output for an attribute of this resource.
// Attribute: ResolverRuleAssociationId
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetConventionalOutputAtt__ResolverRuleAssociationId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttResolverRuleAssociationId", t.GetAtt__ResolverRuleAssociationId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ResolverRuleId returns a conventionally configured output for an attribute of this resource.
// Attribute: ResolverRuleId
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetConventionalOutputAtt__ResolverRuleId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttResolverRuleId", t.GetAtt__ResolverRuleId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__VPCId returns a conventionally configured output for an attribute of this resource.
// Attribute: VPCId
func (t *AWS_Route53Resolver_ResolverRuleAssociation) GetConventionalOutputAtt__VPCId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttVPCId", t.GetAtt__VPCId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Route53Resolver_ResolverRuleAssociation) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Route53Resolver_ResolverRuleAssociation

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

func (t *AWS_Route53Resolver_ResolverRuleAssociation) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
