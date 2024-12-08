// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_chatbot

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Chatbot_MicrosoftTeamsChannelConfiguration)(nil)
	_ cfz.Resource                   = (*AWS_Chatbot_MicrosoftTeamsChannelConfiguration)(nil)
)

const (
	// AWS_Chatbot_MicrosoftTeamsChannelConfiguration__Type is the CloudFormation type for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
	AWS_Chatbot_MicrosoftTeamsChannelConfiguration__Type = "AWS::Chatbot::MicrosoftTeamsChannelConfiguration"
)

var (
	// AWS_Chatbot_MicrosoftTeamsChannelConfiguration__AttributesMap reports all the CloudFormation attributes for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
	AWS_Chatbot_MicrosoftTeamsChannelConfiguration__AttributesMap = struct {
		Arn string
	}{
		Arn: "Arn",
	}

	// AWS_Chatbot_MicrosoftTeamsChannelConfiguration__AttributesSlice reports all the CloudFormation attributes for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
	AWS_Chatbot_MicrosoftTeamsChannelConfiguration__AttributesSlice = []string{
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__AttributesMap.Arn,
	}
)

var (
	// AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
	AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap = struct {
		ConfigurationName         string
		CustomizationResourceArns string
		GuardrailPolicies         string
		IamRoleArn                string
		LoggingLevel              string
		SnsTopicArns              string
		Tags                      string
		TeamId                    string
		TeamsChannelId            string
		TeamsTenantId             string
		UserRoleRequired          string
	}{
		ConfigurationName:         "ConfigurationName",
		CustomizationResourceArns: "CustomizationResourceArns",
		GuardrailPolicies:         "GuardrailPolicies",
		IamRoleArn:                "IamRoleArn",
		LoggingLevel:              "LoggingLevel",
		SnsTopicArns:              "SnsTopicArns",
		Tags:                      "Tags",
		TeamId:                    "TeamId",
		TeamsChannelId:            "TeamsChannelId",
		TeamsTenantId:             "TeamsTenantId",
		UserRoleRequired:          "UserRoleRequired",
	}

	// AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
	AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesSlice = []string{
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.ConfigurationName,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.CustomizationResourceArns,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.GuardrailPolicies,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.IamRoleArn,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.LoggingLevel,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.SnsTopicArns,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.Tags,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.TeamId,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.TeamsChannelId,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.TeamsTenantId,
		AWS_Chatbot_MicrosoftTeamsChannelConfiguration__PropertiesMap.UserRoleRequired,
	}
)

// AWS_Chatbot_MicrosoftTeamsChannelConfiguration is a binding for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html
type AWS_Chatbot_MicrosoftTeamsChannelConfiguration struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// ConfigurationName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-configurationname
	ConfigurationName cfz.Expression[string] `json:"ConfigurationName,omitempty"`

	// CustomizationResourceArns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-customizationresourcearns
	CustomizationResourceArns cfz.ExpressionSlice[string] `json:"CustomizationResourceArns,omitempty"`

	// GuardrailPolicies is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-guardrailpolicies
	GuardrailPolicies cfz.ExpressionSlice[string] `json:"GuardrailPolicies,omitempty"`

	// IamRoleArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-iamrolearn
	IamRoleArn cfz.Expression[string] `json:"IamRoleArn,omitempty"`

	// LoggingLevel is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-logginglevel
	LoggingLevel cfz.Expression[string] `json:"LoggingLevel,omitempty"`

	// SnsTopicArns is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-snstopicarns
	SnsTopicArns cfz.ExpressionSlice[string] `json:"SnsTopicArns,omitempty"`

	// Tags is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-tags
	Tags cfz.ExpressionSlice[cfz.Tag] `json:"Tags,omitempty"`

	// TeamId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamid
	TeamId cfz.Expression[string] `json:"TeamId,omitempty"`

	// TeamsChannelId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamschannelid
	TeamsChannelId cfz.Expression[string] `json:"TeamsChannelId,omitempty"`

	// TeamsTenantId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-teamstenantid
	TeamsTenantId cfz.Expression[string] `json:"TeamsTenantId,omitempty"`

	// UserRoleRequired is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-microsoftteamschannelconfiguration.html#cfn-chatbot-microsoftteamschannelconfiguration-userrolerequired
	UserRoleRequired cfz.Expression[bool] `json:"UserRoleRequired,omitempty"`
}

// New__AWS_Chatbot_MicrosoftTeamsChannelConfiguration initializes a new *AWS_Chatbot_MicrosoftTeamsChannelConfiguration.
func New__AWS_Chatbot_MicrosoftTeamsChannelConfiguration(logicalName string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	return &AWS_Chatbot_MicrosoftTeamsChannelConfiguration{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Chatbot_MicrosoftTeamsChannelConfiguration) GetType() string {
	return AWS_Chatbot_MicrosoftTeamsChannelConfiguration__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__LogicalName(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__RequestedOutputs(v []cfz.Output) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Add__RequestedOutputs(v ...cfz.Output) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__ConfigurationName updates property "ConfigurationName".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__ConfigurationName(v cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.ConfigurationName = v
	return t
}

// SetV__ConfigurationName updates property "ConfigurationName".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__ConfigurationName(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.ConfigurationName = cfz.V(v)
	return t
}

// Set__CustomizationResourceArns updates property "CustomizationResourceArns".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__CustomizationResourceArns(v cfz.ExpressionSlice[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.CustomizationResourceArns = v
	return t
}

// SetS__CustomizationResourceArns updates property "CustomizationResourceArns".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetS__CustomizationResourceArns(v ...cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.CustomizationResourceArns = cfz.S(v...)
	return t
}

// SetSV__CustomizationResourceArns updates property "CustomizationResourceArns".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetSV__CustomizationResourceArns(v ...string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.CustomizationResourceArns = cfz.SV(v...)
	return t
}

// Set__GuardrailPolicies updates property "GuardrailPolicies".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__GuardrailPolicies(v cfz.ExpressionSlice[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.GuardrailPolicies = v
	return t
}

// SetS__GuardrailPolicies updates property "GuardrailPolicies".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetS__GuardrailPolicies(v ...cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.GuardrailPolicies = cfz.S(v...)
	return t
}

// SetSV__GuardrailPolicies updates property "GuardrailPolicies".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetSV__GuardrailPolicies(v ...string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.GuardrailPolicies = cfz.SV(v...)
	return t
}

// Set__IamRoleArn updates property "IamRoleArn".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__IamRoleArn(v cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.IamRoleArn = v
	return t
}

// SetV__IamRoleArn updates property "IamRoleArn".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__IamRoleArn(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.IamRoleArn = cfz.V(v)
	return t
}

// Set__LoggingLevel updates property "LoggingLevel".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__LoggingLevel(v cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.LoggingLevel = v
	return t
}

// SetV__LoggingLevel updates property "LoggingLevel".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__LoggingLevel(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.LoggingLevel = cfz.V(v)
	return t
}

// Set__SnsTopicArns updates property "SnsTopicArns".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__SnsTopicArns(v cfz.ExpressionSlice[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.SnsTopicArns = v
	return t
}

// SetS__SnsTopicArns updates property "SnsTopicArns".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetS__SnsTopicArns(v ...cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.SnsTopicArns = cfz.S(v...)
	return t
}

// SetSV__SnsTopicArns updates property "SnsTopicArns".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetSV__SnsTopicArns(v ...string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.SnsTopicArns = cfz.SV(v...)
	return t
}

// Set__Tags updates property "Tags".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__Tags(v cfz.ExpressionSlice[cfz.Tag]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.Tags = v
	return t
}

// SetS__Tags updates property "Tags".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetS__Tags(v ...cfz.Expression[cfz.Tag]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.Tags = cfz.S(v...)
	return t
}

// SetSV__Tags updates property "Tags".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetSV__Tags(v ...cfz.Tag) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.Tags = cfz.SV(v...)
	return t
}

// Set__TeamId updates property "TeamId".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__TeamId(v cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.TeamId = v
	return t
}

// SetV__TeamId updates property "TeamId".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__TeamId(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.TeamId = cfz.V(v)
	return t
}

// Set__TeamsChannelId updates property "TeamsChannelId".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__TeamsChannelId(v cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.TeamsChannelId = v
	return t
}

// SetV__TeamsChannelId updates property "TeamsChannelId".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__TeamsChannelId(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.TeamsChannelId = cfz.V(v)
	return t
}

// Set__TeamsTenantId updates property "TeamsTenantId".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__TeamsTenantId(v cfz.Expression[string]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.TeamsTenantId = v
	return t
}

// SetV__TeamsTenantId updates property "TeamsTenantId".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__TeamsTenantId(v string) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.TeamsTenantId = cfz.V(v)
	return t
}

// Set__UserRoleRequired updates property "UserRoleRequired".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Set__UserRoleRequired(v cfz.Expression[bool]) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.UserRoleRequired = v
	return t
}

// SetV__UserRoleRequired updates property "UserRoleRequired".
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) SetV__UserRoleRequired(v bool) *AWS_Chatbot_MicrosoftTeamsChannelConfiguration {
	t.UserRoleRequired = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetAtt__Arn returns a $cfz.Expression[string] that resolves to the FN::GetAtt intrinsic function for this resource.
// Attribute: Arn
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) GetAtt__Arn() cfz.Expression[string] {
	return cfz.GetAtt[string](cfz.V(t.GetResourceLogicalName()), cfz.V(AWS_Chatbot_MicrosoftTeamsChannelConfiguration__AttributesMap.Arn))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// GetConventionalOutputAtt__Arn returns a conventionally configured output for an attribute of this resource.
// Attribute: Arn
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) GetConventionalOutputAtt__Arn(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"AttArn", t.GetAtt__Arn())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Chatbot_MicrosoftTeamsChannelConfiguration

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

func (t *AWS_Chatbot_MicrosoftTeamsChannelConfiguration) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
