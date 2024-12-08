// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_cognito

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__Type is the CloudFormation type for AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType.
	AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__Type = "AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType"
)

var (
	// AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__PropertiesMap reports all the CloudFormation properties for AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType.
	AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__PropertiesMap = struct {
		Actions     string
		EventFilter string
	}{
		Actions:     "Actions",
		EventFilter: "EventFilter",
	}

	// AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__PropertiesSlice reports all the CloudFormation properties for AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType.
	AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__PropertiesSlice = []string{
		AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__PropertiesMap.Actions,
		AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__PropertiesMap.EventFilter,
	}
)

// AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType is a binding for AWS::Cognito::UserPoolRiskConfigurationAttachment.CompromisedCredentialsRiskConfigurationType.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html
type AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType struct {
	// Actions is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-actions
	Actions cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsType] `json:"Actions,omitempty"`

	// EventFilter is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype-eventfilter
	EventFilter cfz.ExpressionSlice[string] `json:"EventFilter,omitempty"`
}

// New__AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType initializes a new AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType.
func New__AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType() AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType {
	return AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType{}
}

// GetType returns the CloudFormation type.
func (AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) GetType() string {
	return AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType__Type
}

// Set__Actions updates property "Actions".
func (t AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) Set__Actions(v cfz.Expression[AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsType]) AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType {
	t.Actions = v
	return t
}

// SetV__Actions updates property "Actions".
func (t AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) SetV__Actions(v AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsActionsType) AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType {
	t.Actions = cfz.V(v)
	return t
}

// Set__EventFilter updates property "EventFilter".
func (t AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) Set__EventFilter(v cfz.ExpressionSlice[string]) AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType {
	t.EventFilter = v
	return t
}

// SetS__EventFilter updates property "EventFilter".
func (t AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) SetS__EventFilter(v ...cfz.Expression[string]) AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType {
	t.EventFilter = cfz.S(v...)
	return t
}

// SetSV__EventFilter updates property "EventFilter".
func (t AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType) SetSV__EventFilter(v ...string) AWS_Cognito_UserPoolRiskConfigurationAttachment_CompromisedCredentialsRiskConfigurationType {
	t.EventFilter = cfz.SV(v...)
	return t
}
