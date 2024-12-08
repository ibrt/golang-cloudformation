// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_appflow

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__Type is the CloudFormation type for AWS::AppFlow::ConnectorProfile.CustomAuthCredentials.
	AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__Type = "AWS::AppFlow::ConnectorProfile.CustomAuthCredentials"
)

var (
	// AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__PropertiesMap reports all the CloudFormation properties for AWS::AppFlow::ConnectorProfile.CustomAuthCredentials.
	AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__PropertiesMap = struct {
		CredentialsMap           string
		CustomAuthenticationType string
	}{
		CredentialsMap:           "CredentialsMap",
		CustomAuthenticationType: "CustomAuthenticationType",
	}

	// AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__PropertiesSlice reports all the CloudFormation properties for AWS::AppFlow::ConnectorProfile.CustomAuthCredentials.
	AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__PropertiesSlice = []string{
		AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__PropertiesMap.CredentialsMap,
		AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__PropertiesMap.CustomAuthenticationType,
	}
)

// AWS_AppFlow_ConnectorProfile_CustomAuthCredentials is a binding for AWS::AppFlow::ConnectorProfile.CustomAuthCredentials.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html
type AWS_AppFlow_ConnectorProfile_CustomAuthCredentials struct {
	// CredentialsMap is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-credentialsmap
	CredentialsMap cfz.ExpressionMap[string] `json:"CredentialsMap,omitempty"`

	// CustomAuthenticationType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customauthcredentials.html#cfn-appflow-connectorprofile-customauthcredentials-customauthenticationtype
	CustomAuthenticationType cfz.Expression[string] `json:"CustomAuthenticationType,omitempty"`
}

// New__AWS_AppFlow_ConnectorProfile_CustomAuthCredentials initializes a new AWS_AppFlow_ConnectorProfile_CustomAuthCredentials.
func New__AWS_AppFlow_ConnectorProfile_CustomAuthCredentials() AWS_AppFlow_ConnectorProfile_CustomAuthCredentials {
	return AWS_AppFlow_ConnectorProfile_CustomAuthCredentials{}
}

// GetType returns the CloudFormation type.
func (AWS_AppFlow_ConnectorProfile_CustomAuthCredentials) GetType() string {
	return AWS_AppFlow_ConnectorProfile_CustomAuthCredentials__Type
}

// Set__CredentialsMap updates property "CredentialsMap".
func (t AWS_AppFlow_ConnectorProfile_CustomAuthCredentials) Set__CredentialsMap(v cfz.ExpressionMap[string]) AWS_AppFlow_ConnectorProfile_CustomAuthCredentials {
	t.CredentialsMap = v
	return t
}

// SetM__CredentialsMap updates property "CredentialsMap".
func (t AWS_AppFlow_ConnectorProfile_CustomAuthCredentials) SetM__CredentialsMap(v ...map[string]cfz.Expression[string]) AWS_AppFlow_ConnectorProfile_CustomAuthCredentials {
	t.CredentialsMap = cfz.M(v...)
	return t
}

// SetMV__CredentialsMap updates property "CredentialsMap".
func (t AWS_AppFlow_ConnectorProfile_CustomAuthCredentials) SetMV__CredentialsMap(v ...map[string]string) AWS_AppFlow_ConnectorProfile_CustomAuthCredentials {
	t.CredentialsMap = cfz.MV(v...)
	return t
}

// Set__CustomAuthenticationType updates property "CustomAuthenticationType".
func (t AWS_AppFlow_ConnectorProfile_CustomAuthCredentials) Set__CustomAuthenticationType(v cfz.Expression[string]) AWS_AppFlow_ConnectorProfile_CustomAuthCredentials {
	t.CustomAuthenticationType = v
	return t
}

// SetV__CustomAuthenticationType updates property "CustomAuthenticationType".
func (t AWS_AppFlow_ConnectorProfile_CustomAuthCredentials) SetV__CustomAuthenticationType(v string) AWS_AppFlow_ConnectorProfile_CustomAuthCredentials {
	t.CustomAuthenticationType = cfz.V(v)
	return t
}
