// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_route53resolver

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Route53Resolver_ResolverDNSSECConfig)(nil)
	_ cfz.Resource                   = (*AWS_Route53Resolver_ResolverDNSSECConfig)(nil)
)

const (
	// AWS_Route53Resolver_ResolverDNSSECConfig__Type is the CloudFormation type for AWS::Route53Resolver::ResolverDNSSECConfig.
	AWS_Route53Resolver_ResolverDNSSECConfig__Type = "AWS::Route53Resolver::ResolverDNSSECConfig"
)

var (
	// AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap reports all the CloudFormation attributes for AWS::Route53Resolver::ResolverDNSSECConfig.
	AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap = struct {
		Id               string
		OwnerId          string
		ValidationStatus string
	}{
		Id:               "Id",
		OwnerId:          "OwnerId",
		ValidationStatus: "ValidationStatus",
	}

	// AWS_Route53Resolver_ResolverDNSSECConfig__AttributesSlice reports all the CloudFormation attributes for AWS::Route53Resolver::ResolverDNSSECConfig.
	AWS_Route53Resolver_ResolverDNSSECConfig__AttributesSlice = []string{
		AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap.Id,
		AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap.OwnerId,
		AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap.ValidationStatus,
	}
)

var (
	// AWS_Route53Resolver_ResolverDNSSECConfig__PropertiesMap reports all the CloudFormation properties for AWS::Route53Resolver::ResolverDNSSECConfig.
	AWS_Route53Resolver_ResolverDNSSECConfig__PropertiesMap = struct {
		ResourceId string
	}{
		ResourceId: "ResourceId",
	}

	// AWS_Route53Resolver_ResolverDNSSECConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Route53Resolver::ResolverDNSSECConfig.
	AWS_Route53Resolver_ResolverDNSSECConfig__PropertiesSlice = []string{
		AWS_Route53Resolver_ResolverDNSSECConfig__PropertiesMap.ResourceId,
	}
)

// AWS_Route53Resolver_ResolverDNSSECConfig is a binding for AWS::Route53Resolver::ResolverDNSSECConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverdnssecconfig.html
type AWS_Route53Resolver_ResolverDNSSECConfig struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ResourceId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverdnssecconfig.html#cfn-route53resolver-resolverdnssecconfig-resourceid
	ResourceId cfz.Expression[string] `json:"ResourceId,omitempty"`
}

// New__AWS_Route53Resolver_ResolverDNSSECConfig initializes a new *AWS_Route53Resolver_ResolverDNSSECConfig.
func New__AWS_Route53Resolver_ResolverDNSSECConfig(logicalName string) *AWS_Route53Resolver_ResolverDNSSECConfig {
	return &AWS_Route53Resolver_ResolverDNSSECConfig{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Route53Resolver_ResolverDNSSECConfig) GetType() string {
	return AWS_Route53Resolver_ResolverDNSSECConfig__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Set__LogicalName(v string) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Set__RequestedOutputs(v []cfz.Output) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Add__RequestedOutputs(v ...cfz.Output) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ResourceId updates property "ResourceId".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Set__ResourceId(v cfz.Expression[string]) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.ResourceId = v
	return t
}

// SetV__ResourceId updates property "ResourceId".
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) SetV__ResourceId(v string) *AWS_Route53Resolver_ResolverDNSSECConfig {
	t.ResourceId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Id returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Id
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetAtt__Id() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap.Id))
}

// GetAtt__OwnerId returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: OwnerId
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetAtt__OwnerId() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap.OwnerId))
}

// GetAtt__ValidationStatus returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ValidationStatus
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetAtt__ValidationStatus() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Route53Resolver_ResolverDNSSECConfig__AttributesMap.ValidationStatus))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Id returns a conventionally configured output for an attribute of this resource.
// Attribute: Id
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetConventionalOutputAtt__Id(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttId", t.GetAtt__Id())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__OwnerId returns a conventionally configured output for an attribute of this resource.
// Attribute: OwnerId
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetConventionalOutputAtt__OwnerId(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttOwnerId", t.GetAtt__OwnerId())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ValidationStatus returns a conventionally configured output for an attribute of this resource.
// Attribute: ValidationStatus
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) GetConventionalOutputAtt__ValidationStatus(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttValidationStatus", t.GetAtt__ValidationStatus())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Route53Resolver_ResolverDNSSECConfig) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Route53Resolver_ResolverDNSSECConfig

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

func (t *AWS_Route53Resolver_ResolverDNSSECConfig) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
