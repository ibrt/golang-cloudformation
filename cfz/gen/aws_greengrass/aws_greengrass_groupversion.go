// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_greengrass

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Greengrass_GroupVersion)(nil)
	_ cfz.Resource                   = (*AWS_Greengrass_GroupVersion)(nil)
)

const (
	// AWS_Greengrass_GroupVersion__Type is the CloudFormation type for AWS::Greengrass::GroupVersion.
	AWS_Greengrass_GroupVersion__Type = "AWS::Greengrass::GroupVersion"
)

var (
	// AWS_Greengrass_GroupVersion__PropertiesMap reports all the CloudFormation properties for AWS::Greengrass::GroupVersion.
	AWS_Greengrass_GroupVersion__PropertiesMap = struct {
		ConnectorDefinitionVersionArn    string
		CoreDefinitionVersionArn         string
		DeviceDefinitionVersionArn       string
		FunctionDefinitionVersionArn     string
		GroupId                          string
		LoggerDefinitionVersionArn       string
		ResourceDefinitionVersionArn     string
		SubscriptionDefinitionVersionArn string
	}{
		ConnectorDefinitionVersionArn:    "ConnectorDefinitionVersionArn",
		CoreDefinitionVersionArn:         "CoreDefinitionVersionArn",
		DeviceDefinitionVersionArn:       "DeviceDefinitionVersionArn",
		FunctionDefinitionVersionArn:     "FunctionDefinitionVersionArn",
		GroupId:                          "GroupId",
		LoggerDefinitionVersionArn:       "LoggerDefinitionVersionArn",
		ResourceDefinitionVersionArn:     "ResourceDefinitionVersionArn",
		SubscriptionDefinitionVersionArn: "SubscriptionDefinitionVersionArn",
	}

	// AWS_Greengrass_GroupVersion__PropertiesSlice reports all the CloudFormation properties for AWS::Greengrass::GroupVersion.
	AWS_Greengrass_GroupVersion__PropertiesSlice = []string{
		AWS_Greengrass_GroupVersion__PropertiesMap.ConnectorDefinitionVersionArn,
		AWS_Greengrass_GroupVersion__PropertiesMap.CoreDefinitionVersionArn,
		AWS_Greengrass_GroupVersion__PropertiesMap.DeviceDefinitionVersionArn,
		AWS_Greengrass_GroupVersion__PropertiesMap.FunctionDefinitionVersionArn,
		AWS_Greengrass_GroupVersion__PropertiesMap.GroupId,
		AWS_Greengrass_GroupVersion__PropertiesMap.LoggerDefinitionVersionArn,
		AWS_Greengrass_GroupVersion__PropertiesMap.ResourceDefinitionVersionArn,
		AWS_Greengrass_GroupVersion__PropertiesMap.SubscriptionDefinitionVersionArn,
	}
)

// AWS_Greengrass_GroupVersion is a binding for AWS::Greengrass::GroupVersion.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html
type AWS_Greengrass_GroupVersion struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ConnectorDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-connectordefinitionversionarn
	ConnectorDefinitionVersionArn cfz.Expression[string] `json:"ConnectorDefinitionVersionArn,omitempty"`

	// CoreDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-coredefinitionversionarn
	CoreDefinitionVersionArn cfz.Expression[string] `json:"CoreDefinitionVersionArn,omitempty"`

	// DeviceDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-devicedefinitionversionarn
	DeviceDefinitionVersionArn cfz.Expression[string] `json:"DeviceDefinitionVersionArn,omitempty"`

	// FunctionDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-functiondefinitionversionarn
	FunctionDefinitionVersionArn cfz.Expression[string] `json:"FunctionDefinitionVersionArn,omitempty"`

	// GroupId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-groupid
	GroupId cfz.Expression[string] `json:"GroupId,omitempty"`

	// LoggerDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-loggerdefinitionversionarn
	LoggerDefinitionVersionArn cfz.Expression[string] `json:"LoggerDefinitionVersionArn,omitempty"`

	// ResourceDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-resourcedefinitionversionarn
	ResourceDefinitionVersionArn cfz.Expression[string] `json:"ResourceDefinitionVersionArn,omitempty"`

	// SubscriptionDefinitionVersionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-groupversion.html#cfn-greengrass-groupversion-subscriptiondefinitionversionarn
	SubscriptionDefinitionVersionArn cfz.Expression[string] `json:"SubscriptionDefinitionVersionArn,omitempty"`
}

// New__AWS_Greengrass_GroupVersion initializes a new *AWS_Greengrass_GroupVersion.
func New__AWS_Greengrass_GroupVersion(logicalName string) *AWS_Greengrass_GroupVersion {
	return &AWS_Greengrass_GroupVersion{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Greengrass_GroupVersion) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Greengrass_GroupVersion) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Greengrass_GroupVersion) GetType() string {
	return AWS_Greengrass_GroupVersion__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Greengrass_GroupVersion) Set__LogicalName(v string) *AWS_Greengrass_GroupVersion {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Greengrass_GroupVersion) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Greengrass_GroupVersion {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Greengrass_GroupVersion) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Greengrass_GroupVersion {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Greengrass_GroupVersion) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Greengrass_GroupVersion {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Greengrass_GroupVersion) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Greengrass_GroupVersion {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Greengrass_GroupVersion) Set__RequestedOutputs(v []cfz.Output) *AWS_Greengrass_GroupVersion {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Greengrass_GroupVersion) Add__RequestedOutputs(v ...cfz.Output) *AWS_Greengrass_GroupVersion {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ConnectorDefinitionVersionArn updates property "ConnectorDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__ConnectorDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.ConnectorDefinitionVersionArn = v
	return t
}

// SetV__ConnectorDefinitionVersionArn updates property "ConnectorDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__ConnectorDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.ConnectorDefinitionVersionArn = cfz.V(v)
	return t
}

// Set__CoreDefinitionVersionArn updates property "CoreDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__CoreDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.CoreDefinitionVersionArn = v
	return t
}

// SetV__CoreDefinitionVersionArn updates property "CoreDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__CoreDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.CoreDefinitionVersionArn = cfz.V(v)
	return t
}

// Set__DeviceDefinitionVersionArn updates property "DeviceDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__DeviceDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.DeviceDefinitionVersionArn = v
	return t
}

// SetV__DeviceDefinitionVersionArn updates property "DeviceDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__DeviceDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.DeviceDefinitionVersionArn = cfz.V(v)
	return t
}

// Set__FunctionDefinitionVersionArn updates property "FunctionDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__FunctionDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.FunctionDefinitionVersionArn = v
	return t
}

// SetV__FunctionDefinitionVersionArn updates property "FunctionDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__FunctionDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.FunctionDefinitionVersionArn = cfz.V(v)
	return t
}

// Set__GroupId updates property "GroupId".
func (t *AWS_Greengrass_GroupVersion) Set__GroupId(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.GroupId = v
	return t
}

// SetV__GroupId updates property "GroupId".
func (t *AWS_Greengrass_GroupVersion) SetV__GroupId(v string) *AWS_Greengrass_GroupVersion {
	t.GroupId = cfz.V(v)
	return t
}

// Set__LoggerDefinitionVersionArn updates property "LoggerDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__LoggerDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.LoggerDefinitionVersionArn = v
	return t
}

// SetV__LoggerDefinitionVersionArn updates property "LoggerDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__LoggerDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.LoggerDefinitionVersionArn = cfz.V(v)
	return t
}

// Set__ResourceDefinitionVersionArn updates property "ResourceDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__ResourceDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.ResourceDefinitionVersionArn = v
	return t
}

// SetV__ResourceDefinitionVersionArn updates property "ResourceDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__ResourceDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.ResourceDefinitionVersionArn = cfz.V(v)
	return t
}

// Set__SubscriptionDefinitionVersionArn updates property "SubscriptionDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) Set__SubscriptionDefinitionVersionArn(v cfz.Expression[string]) *AWS_Greengrass_GroupVersion {
	t.SubscriptionDefinitionVersionArn = v
	return t
}

// SetV__SubscriptionDefinitionVersionArn updates property "SubscriptionDefinitionVersionArn".
func (t *AWS_Greengrass_GroupVersion) SetV__SubscriptionDefinitionVersionArn(v string) *AWS_Greengrass_GroupVersion {
	t.SubscriptionDefinitionVersionArn = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Greengrass_GroupVersion) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Greengrass_GroupVersion) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Greengrass_GroupVersion) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Greengrass_GroupVersion

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

func (t *AWS_Greengrass_GroupVersion) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
