// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_wafv2

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_WAFv2_LoggingConfiguration)(nil)
	_ cfz.Resource                   = (*AWS_WAFv2_LoggingConfiguration)(nil)
)

const (
	// AWS_WAFv2_LoggingConfiguration__Type is the CloudFormation type for AWS::WAFv2::LoggingConfiguration.
	AWS_WAFv2_LoggingConfiguration__Type = "AWS::WAFv2::LoggingConfiguration"
)

var (
	// AWS_WAFv2_LoggingConfiguration__AttributesMap reports all the CloudFormation attributes for AWS::WAFv2::LoggingConfiguration.
	AWS_WAFv2_LoggingConfiguration__AttributesMap = struct {
		ManagedByFirewallManager string
	}{
		ManagedByFirewallManager: "ManagedByFirewallManager",
	}

	// AWS_WAFv2_LoggingConfiguration__AttributesSlice reports all the CloudFormation attributes for AWS::WAFv2::LoggingConfiguration.
	AWS_WAFv2_LoggingConfiguration__AttributesSlice = []string{
		AWS_WAFv2_LoggingConfiguration__AttributesMap.ManagedByFirewallManager,
	}
)

var (
	// AWS_WAFv2_LoggingConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::WAFv2::LoggingConfiguration.
	AWS_WAFv2_LoggingConfiguration__PropertiesMap = struct {
		LogDestinationConfigs string
		LoggingFilter         string
		RedactedFields        string
		ResourceArn           string
	}{
		LogDestinationConfigs: "LogDestinationConfigs",
		LoggingFilter:         "LoggingFilter",
		RedactedFields:        "RedactedFields",
		ResourceArn:           "ResourceArn",
	}

	// AWS_WAFv2_LoggingConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::WAFv2::LoggingConfiguration.
	AWS_WAFv2_LoggingConfiguration__PropertiesSlice = []string{
		AWS_WAFv2_LoggingConfiguration__PropertiesMap.LogDestinationConfigs,
		AWS_WAFv2_LoggingConfiguration__PropertiesMap.LoggingFilter,
		AWS_WAFv2_LoggingConfiguration__PropertiesMap.RedactedFields,
		AWS_WAFv2_LoggingConfiguration__PropertiesMap.ResourceArn,
	}
)

// AWS_WAFv2_LoggingConfiguration is a binding for AWS::WAFv2::LoggingConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-loggingconfiguration.html
type AWS_WAFv2_LoggingConfiguration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// LogDestinationConfigs is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-loggingconfiguration.html#cfn-wafv2-loggingconfiguration-logdestinationconfigs
	LogDestinationConfigs cfz.ExpressionSlice[string] `json:"LogDestinationConfigs,omitempty"`

	// LoggingFilter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-loggingconfiguration.html#cfn-wafv2-loggingconfiguration-loggingfilter
	LoggingFilter cfz.Expression[AWS_WAFv2_LoggingConfiguration_LoggingFilter] `json:"LoggingFilter,omitempty"`

	// RedactedFields is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-loggingconfiguration.html#cfn-wafv2-loggingconfiguration-redactedfields
	RedactedFields cfz.ExpressionSlice[AWS_WAFv2_LoggingConfiguration_FieldToMatch] `json:"RedactedFields,omitempty"`

	// ResourceArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-loggingconfiguration.html#cfn-wafv2-loggingconfiguration-resourcearn
	ResourceArn cfz.Expression[string] `json:"ResourceArn,omitempty"`
}

// New__AWS_WAFv2_LoggingConfiguration initializes a new *AWS_WAFv2_LoggingConfiguration.
func New__AWS_WAFv2_LoggingConfiguration(logicalName string) *AWS_WAFv2_LoggingConfiguration {
	return &AWS_WAFv2_LoggingConfiguration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_WAFv2_LoggingConfiguration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_WAFv2_LoggingConfiguration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_WAFv2_LoggingConfiguration) GetType() string {
	return AWS_WAFv2_LoggingConfiguration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_WAFv2_LoggingConfiguration) Set__LogicalName(v string) *AWS_WAFv2_LoggingConfiguration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_WAFv2_LoggingConfiguration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_WAFv2_LoggingConfiguration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_WAFv2_LoggingConfiguration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_WAFv2_LoggingConfiguration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_WAFv2_LoggingConfiguration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_WAFv2_LoggingConfiguration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_WAFv2_LoggingConfiguration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_WAFv2_LoggingConfiguration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_WAFv2_LoggingConfiguration) Set__RequestedOutputs(v []cfz.Output) *AWS_WAFv2_LoggingConfiguration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_WAFv2_LoggingConfiguration) Add__RequestedOutputs(v ...cfz.Output) *AWS_WAFv2_LoggingConfiguration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__LogDestinationConfigs updates property "LogDestinationConfigs".
func (t *AWS_WAFv2_LoggingConfiguration) Set__LogDestinationConfigs(v cfz.ExpressionSlice[string]) *AWS_WAFv2_LoggingConfiguration {
	t.LogDestinationConfigs = v
	return t
}

// SetS__LogDestinationConfigs updates property "LogDestinationConfigs".
func (t *AWS_WAFv2_LoggingConfiguration) SetS__LogDestinationConfigs(v ...cfz.Expression[string]) *AWS_WAFv2_LoggingConfiguration {
	t.LogDestinationConfigs = cfz.S(v...)
	return t
}

// SetSV__LogDestinationConfigs updates property "LogDestinationConfigs".
func (t *AWS_WAFv2_LoggingConfiguration) SetSV__LogDestinationConfigs(v ...string) *AWS_WAFv2_LoggingConfiguration {
	t.LogDestinationConfigs = cfz.SV(v...)
	return t
}

// Set__LoggingFilter updates property "LoggingFilter".
func (t *AWS_WAFv2_LoggingConfiguration) Set__LoggingFilter(v cfz.Expression[AWS_WAFv2_LoggingConfiguration_LoggingFilter]) *AWS_WAFv2_LoggingConfiguration {
	t.LoggingFilter = v
	return t
}

// SetV__LoggingFilter updates property "LoggingFilter".
func (t *AWS_WAFv2_LoggingConfiguration) SetV__LoggingFilter(v AWS_WAFv2_LoggingConfiguration_LoggingFilter) *AWS_WAFv2_LoggingConfiguration {
	t.LoggingFilter = cfz.V(v)
	return t
}

// Set__RedactedFields updates property "RedactedFields".
func (t *AWS_WAFv2_LoggingConfiguration) Set__RedactedFields(v cfz.ExpressionSlice[AWS_WAFv2_LoggingConfiguration_FieldToMatch]) *AWS_WAFv2_LoggingConfiguration {
	t.RedactedFields = v
	return t
}

// SetS__RedactedFields updates property "RedactedFields".
func (t *AWS_WAFv2_LoggingConfiguration) SetS__RedactedFields(v ...cfz.Expression[AWS_WAFv2_LoggingConfiguration_FieldToMatch]) *AWS_WAFv2_LoggingConfiguration {
	t.RedactedFields = cfz.S(v...)
	return t
}

// SetSV__RedactedFields updates property "RedactedFields".
func (t *AWS_WAFv2_LoggingConfiguration) SetSV__RedactedFields(v ...AWS_WAFv2_LoggingConfiguration_FieldToMatch) *AWS_WAFv2_LoggingConfiguration {
	t.RedactedFields = cfz.SV(v...)
	return t
}

// Set__ResourceArn updates property "ResourceArn".
func (t *AWS_WAFv2_LoggingConfiguration) Set__ResourceArn(v cfz.Expression[string]) *AWS_WAFv2_LoggingConfiguration {
	t.ResourceArn = v
	return t
}

// SetV__ResourceArn updates property "ResourceArn".
func (t *AWS_WAFv2_LoggingConfiguration) SetV__ResourceArn(v string) *AWS_WAFv2_LoggingConfiguration {
	t.ResourceArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_WAFv2_LoggingConfiguration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__ManagedByFirewallManager returns a $cfz.Expression[bool] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: ManagedByFirewallManager
func (t *AWS_WAFv2_LoggingConfiguration) GetAtt__ManagedByFirewallManager() cfz.Expression[bool] {
	return cfz.GetAtt[bool](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_WAFv2_LoggingConfiguration__AttributesMap.ManagedByFirewallManager))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_WAFv2_LoggingConfiguration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__ManagedByFirewallManager returns a conventionally configured output for an attribute of this resource.
// Attribute: ManagedByFirewallManager
func (t *AWS_WAFv2_LoggingConfiguration) GetConventionalOutputAtt__ManagedByFirewallManager(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttManagedByFirewallManager", t.GetAtt__ManagedByFirewallManager())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_WAFv2_LoggingConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_WAFv2_LoggingConfiguration

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

func (t *AWS_WAFv2_LoggingConfiguration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
