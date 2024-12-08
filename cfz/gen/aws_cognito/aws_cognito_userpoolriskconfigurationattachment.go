// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cognito

import (
	"encoding/json"
	"github.com/ibrt/golang-cloudformation/cfz"
)

var (
	_ cfz.ResourcePartialLogicalName = (*AWS_Cognito_UserPoolRiskConfigurationAttachment)(nil)
	_ cfz.Resource                   = (*AWS_Cognito_UserPoolRiskConfigurationAttachment)(nil)
)

const (
	// AWS_Cognito_UserPoolRiskConfigurationAttachment__Type is the CloudFormation type for AWS::Cognito::UserPoolRiskConfigurationAttachment.
	AWS_Cognito_UserPoolRiskConfigurationAttachment__Type = "AWS::Cognito::UserPoolRiskConfigurationAttachment"
)

var (
	// AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap reports all the CloudFormation properties for AWS::Cognito::UserPoolRiskConfigurationAttachment.
	AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap = struct {
		AccountTakeoverRiskConfiguration        string
		ClientId                                string
		CompromisedCredentialsRiskConfiguration string
		RiskExceptionConfiguration              string
		UserPoolId                              string
	}{
		AccountTakeoverRiskConfiguration:        "AccountTakeoverRiskConfiguration",
		ClientId:                                "ClientId",
		CompromisedCredentialsRiskConfiguration: "CompromisedCredentialsRiskConfiguration",
		RiskExceptionConfiguration:              "RiskExceptionConfiguration",
		UserPoolId:                              "UserPoolId",
	}

	// AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesSlice reports all the CloudFormation properties for AWS::Cognito::UserPoolRiskConfigurationAttachment.
	AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesSlice = []string{
		AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap.AccountTakeoverRiskConfiguration,
		AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap.ClientId,
		AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap.CompromisedCredentialsRiskConfiguration,
		AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap.RiskExceptionConfiguration,
		AWS_Cognito_UserPoolRiskConfigurationAttachment__PropertiesMap.UserPoolId,
	}
)

// AWS_Cognito_UserPoolRiskConfigurationAttachment is a binding for AWS::Cognito::UserPoolRiskConfigurationAttachment.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html
type AWS_Cognito_UserPoolRiskConfigurationAttachment struct { // __LogicalName is the CloudFormation logical name for this resource in the template.
	__LogicalName string `json:"-"`

	// __DependsOn indicates which resources must be created before this one.
	__DependsOn []cfz.ResourcePartialLogicalName `json:"-"`

	// __DeletionPolicy indicates the deletion behavior of CloudFormation for this resource.
	__DeletionPolicy cfz.ResourceDeletionPolicy `json:"-"`

	// __UpdateReplacePolicy indicates the update replace behavior of CloudFormation for this resource.
	__UpdateReplacePolicy cfz.ResourceUpdateReplacePolicy `json:"-"`

	// __RequestedOutputs allows to group outputs together with resources, for later inclusion in a template.
	__RequestedOutputs []cfz.Output

	// AccountTakeoverRiskConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-accounttakeoverriskconfiguration
	AccountTakeoverRiskConfiguration cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_AccountTakeoverRiskConfigurationType] `json:"AccountTakeoverRiskConfiguration,omitempty"`

	// ClientId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-clientid
	ClientId cfz.Expression[string] `json:"ClientId,omitempty"`

	// CompromisedCredentialsRiskConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfiguration
	CompromisedCredentialsRiskConfiguration cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType] `json:"CompromisedCredentialsRiskConfiguration,omitempty"`

	// RiskExceptionConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-riskexceptionconfiguration
	RiskExceptionConfiguration cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_RiskExceptionConfigurationType] `json:"RiskExceptionConfiguration,omitempty"`

	// UserPoolId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html#cfn-cognito-userpoolriskconfigurationattachment-userpoolid
	UserPoolId cfz.Expression[string] `json:"UserPoolId,omitempty"`
}

// New__AWS_Cognito_UserPoolRiskConfigurationAttachment initializes a new *AWS_Cognito_UserPoolRiskConfigurationAttachment.
func New__AWS_Cognito_UserPoolRiskConfigurationAttachment(logicalName string) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	return &AWS_Cognito_UserPoolRiskConfigurationAttachment{
		__LogicalName: logicalName,
	}
}

// GetLogicalName returns the CloudFormation logical name for this resource in the template.
// It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) GetResourceLogicalName() string {
	return t.__LogicalName
}

// GetRequestedOutputs returns the requested outputs for this resource in the template.
// It implements the cfz.Resource interface.
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) GetRequestedOutputs() []cfz.Output {
	return t.__RequestedOutputs
}

// GetType returns the CloudFormation type.
// It implements the cfz.Resource interface.
func (*AWS_Cognito_UserPoolRiskConfigurationAttachment) GetType() string {
	return AWS_Cognito_UserPoolRiskConfigurationAttachment__Type
}

// Set__LogicalName updates field "__LogicalName".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__LogicalName(v string) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__LogicalName = v
	return t
}

// Set__DependsOn updates (replaces) field "__DependsOn".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__DependsOn(v []cfz.ResourcePartialLogicalName) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__DependsOn = v
	return t
}

// Add__DependsOn updates (appends to) field "__DependsOn".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Add__DependsOn(v ...cfz.ResourcePartialLogicalName) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__DependsOn = append(t.__DependsOn, v...)
	return t
}

// Set__DeletionPolicy updates field "__DeletionPolicy".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__DeletionPolicy(v cfz.ResourceDeletionPolicy) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__DeletionPolicy = v
	return t
}

// Set__UpdateReplacePolicy updates field "__UpdateReplacePolicy".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__UpdateReplacePolicy(v cfz.ResourceUpdateReplacePolicy) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__UpdateReplacePolicy = v
	return t
}

// Set__RequestedOutputs updates (replaces) field "__RequestedOutputs".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__RequestedOutputs(v []cfz.Output) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__RequestedOutputs = v
	return t
}

// Add__RequestedOutputs updates (appends to) field "__RequestedOutputs".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Add__RequestedOutputs(v ...cfz.Output) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.__RequestedOutputs = append(t.__RequestedOutputs, v...)
	return t
}

// Set__AccountTakeoverRiskConfiguration updates property "AccountTakeoverRiskConfiguration".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__AccountTakeoverRiskConfiguration(v cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_AccountTakeoverRiskConfigurationType]) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.AccountTakeoverRiskConfiguration = v
	return t
}

// SetV__AccountTakeoverRiskConfiguration updates property "AccountTakeoverRiskConfiguration".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) SetV__AccountTakeoverRiskConfiguration(v AWS_Cognito_UserPoolRiskConfigurationAttachment_AccountTakeoverRiskConfigurationType) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.AccountTakeoverRiskConfiguration = cfz.V(v)
	return t
}

// Set__ClientId updates property "ClientId".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__ClientId(v cfz.Expression[string]) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.ClientId = v
	return t
}

// SetV__ClientId updates property "ClientId".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) SetV__ClientId(v string) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.ClientId = cfz.V(v)
	return t
}

// Set__CompromisedCredentialsRiskConfiguration updates property "CompromisedCredentialsRiskConfiguration".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__CompromisedCredentialsRiskConfiguration(v cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType]) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.CompromisedCredentialsRiskConfiguration = v
	return t
}

// SetV__CompromisedCredentialsRiskConfiguration updates property "CompromisedCredentialsRiskConfiguration".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) SetV__CompromisedCredentialsRiskConfiguration(v AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.CompromisedCredentialsRiskConfiguration = cfz.V(v)
	return t
}

// Set__RiskExceptionConfiguration updates property "RiskExceptionConfiguration".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__RiskExceptionConfiguration(v cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_RiskExceptionConfigurationType]) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.RiskExceptionConfiguration = v
	return t
}

// SetV__RiskExceptionConfiguration updates property "RiskExceptionConfiguration".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) SetV__RiskExceptionConfiguration(v AWS_Cognito_UserPoolRiskConfigurationAttachment_RiskExceptionConfigurationType) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.RiskExceptionConfiguration = cfz.V(v)
	return t
}

// Set__UserPoolId updates property "UserPoolId".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Set__UserPoolId(v cfz.Expression[string]) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.UserPoolId = v
	return t
}

// SetV__UserPoolId updates property "UserPoolId".
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) SetV__UserPoolId(v string) *AWS_Cognito_UserPoolRiskConfigurationAttachment {
	t.UserPoolId = cfz.V(v)
	return t
}

// Ref returns a cfz.Expression[string] that resolves to the Ref intrinsic function for this resource.
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) Ref() cfz.Expression[string] {
	return cfz.Ref(cfz.V(t.GetResourceLogicalName()))
}

// GetConventionalOutputRef returns a conventionally configured output for the Ref of this resource.
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) GetConventionalOutputRef(isExported bool) cfz.Output {
	o := cfz.NewOutput(t.GetResourceLogicalName()+"Ref", t.Ref())
	if isExported {
		o.SetConventionalExportName()
	}
	return o
}

// MarshalJSON implements the cfz.Resource interface.
func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`null`), nil
	}

	type Properties AWS_Cognito_UserPoolRiskConfigurationAttachment

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

func (t *AWS_Cognito_UserPoolRiskConfigurationAttachment) getDependsOn() []string {
	dependsOn := make([]string, 0, len(t.__DependsOn))

	for _, r := range t.__DependsOn {
		dependsOn = append(dependsOn, r.GetResourceLogicalName())
	}

	return dependsOn
}
