// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_networkfirewall

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_NetworkFirewall_LoggingConfiguration)(nil)
	_ cfz.Resource                   = (*AWS_NetworkFirewall_LoggingConfiguration)(nil)
)

const (
	// AWS_NetworkFirewall_LoggingConfiguration__Type is the CloudFormation type for AWS::NetworkFirewall::LoggingConfiguration.
	AWS_NetworkFirewall_LoggingConfiguration__Type = "AWS::NetworkFirewall::LoggingConfiguration"
)

var (
	// AWS_NetworkFirewall_LoggingConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::NetworkFirewall::LoggingConfiguration.
	AWS_NetworkFirewall_LoggingConfiguration__PropertiesMap = struct {
		FirewallArn          string
		FirewallName         string
		LoggingConfiguration string
	}{
		FirewallArn:          "FirewallArn",
		FirewallName:         "FirewallName",
		LoggingConfiguration: "LoggingConfiguration",
	}

	// AWS_NetworkFirewall_LoggingConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::NetworkFirewall::LoggingConfiguration.
	AWS_NetworkFirewall_LoggingConfiguration__PropertiesSlice = []string{
		AWS_NetworkFirewall_LoggingConfiguration__PropertiesMap.FirewallArn,
		AWS_NetworkFirewall_LoggingConfiguration__PropertiesMap.FirewallName,
		AWS_NetworkFirewall_LoggingConfiguration__PropertiesMap.LoggingConfiguration,
	}
)

// AWS_NetworkFirewall_LoggingConfiguration is a binding for AWS::NetworkFirewall::LoggingConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html
type AWS_NetworkFirewall_LoggingConfiguration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// FirewallArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallarn
	FirewallArn cfz.Expression[string] `json:"FirewallArn,omitempty"`

	// FirewallName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-firewallname
	FirewallName cfz.Expression[string] `json:"FirewallName,omitempty"`

	// LoggingConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration
	LoggingConfiguration cfz.Expression[AWS_NetworkFirewall_LoggingConfiguration_LoggingConfiguration] `json:"LoggingConfiguration,omitempty"`
}

// New__AWS_NetworkFirewall_LoggingConfiguration initializes a new *AWS_NetworkFirewall_LoggingConfiguration.
func New__AWS_NetworkFirewall_LoggingConfiguration(logicalName string) *AWS_NetworkFirewall_LoggingConfiguration {
	return &AWS_NetworkFirewall_LoggingConfiguration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_NetworkFirewall_LoggingConfiguration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_NetworkFirewall_LoggingConfiguration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_NetworkFirewall_LoggingConfiguration) GetType() string {
	return AWS_NetworkFirewall_LoggingConfiguration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__LogicalName(v string) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__RequestedOutputs(v []cfz.Output) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Add__RequestedOutputs(v ...cfz.Output) *AWS_NetworkFirewall_LoggingConfiguration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__FirewallArn updates property "FirewallArn".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__FirewallArn(v cfz.Expression[string]) *AWS_NetworkFirewall_LoggingConfiguration {
	t.FirewallArn = v
	return t
}

// SetV__FirewallArn updates property "FirewallArn".
func (t *AWS_NetworkFirewall_LoggingConfiguration) SetV__FirewallArn(v string) *AWS_NetworkFirewall_LoggingConfiguration {
	t.FirewallArn = cfz.V(v)
	return t
}

// Set__FirewallName updates property "FirewallName".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__FirewallName(v cfz.Expression[string]) *AWS_NetworkFirewall_LoggingConfiguration {
	t.FirewallName = v
	return t
}

// SetV__FirewallName updates property "FirewallName".
func (t *AWS_NetworkFirewall_LoggingConfiguration) SetV__FirewallName(v string) *AWS_NetworkFirewall_LoggingConfiguration {
	t.FirewallName = cfz.V(v)
	return t
}

// Set__LoggingConfiguration updates property "LoggingConfiguration".
func (t *AWS_NetworkFirewall_LoggingConfiguration) Set__LoggingConfiguration(v cfz.Expression[AWS_NetworkFirewall_LoggingConfiguration_LoggingConfiguration]) *AWS_NetworkFirewall_LoggingConfiguration {
	t.LoggingConfiguration = v
	return t
}

// SetV__LoggingConfiguration updates property "LoggingConfiguration".
func (t *AWS_NetworkFirewall_LoggingConfiguration) SetV__LoggingConfiguration(v AWS_NetworkFirewall_LoggingConfiguration_LoggingConfiguration) *AWS_NetworkFirewall_LoggingConfiguration {
	t.LoggingConfiguration = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_NetworkFirewall_LoggingConfiguration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_NetworkFirewall_LoggingConfiguration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_NetworkFirewall_LoggingConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_NetworkFirewall_LoggingConfiguration

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

func (t *AWS_NetworkFirewall_LoggingConfiguration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
