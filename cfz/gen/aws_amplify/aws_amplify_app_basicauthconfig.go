// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amplify

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Amplify_App_BasicAuthConfig__Type is the CloudFormation type for AWS::Amplify::App.BasicAuthConfig.
	AWS_Amplify_App_BasicAuthConfig__Type = "AWS::Amplify::App.BasicAuthConfig"
)

var (
	// AWS_Amplify_App_BasicAuthConfig__PropertiesMap reports all the CloudFormation properties for AWS::Amplify::App.BasicAuthConfig.
	AWS_Amplify_App_BasicAuthConfig__PropertiesMap = struct {
		EnableBasicAuth string
		Password        string
		Username        string
	}{
		EnableBasicAuth: "EnableBasicAuth",
		Password:        "Password",
		Username:        "Username",
	}

	// AWS_Amplify_App_BasicAuthConfig__PropertiesSlice reports all the CloudFormation properties for AWS::Amplify::App.BasicAuthConfig.
	AWS_Amplify_App_BasicAuthConfig__PropertiesSlice = []string{
		AWS_Amplify_App_BasicAuthConfig__PropertiesMap.EnableBasicAuth,
		AWS_Amplify_App_BasicAuthConfig__PropertiesMap.Password,
		AWS_Amplify_App_BasicAuthConfig__PropertiesMap.Username,
	}
)

// AWS_Amplify_App_BasicAuthConfig is a binding for AWS::Amplify::App.BasicAuthConfig.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html
type AWS_Amplify_App_BasicAuthConfig struct {
	// EnableBasicAuth is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html#cfn-amplify-app-basicauthconfig-enablebasicauth
	EnableBasicAuth cfz.Expression[bool] `json:"EnableBasicAuth,omitempty"`

	// Password is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html#cfn-amplify-app-basicauthconfig-password
	Password cfz.Expression[string] `json:"Password,omitempty"`

	// Username is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-app-basicauthconfig.html#cfn-amplify-app-basicauthconfig-username
	Username cfz.Expression[string] `json:"Username,omitempty"`
}

// New__AWS_Amplify_App_BasicAuthConfig initializes a new AWS_Amplify_App_BasicAuthConfig.
func New__AWS_Amplify_App_BasicAuthConfig() AWS_Amplify_App_BasicAuthConfig {
	return AWS_Amplify_App_BasicAuthConfig{}
}

// GetType returns the CloudFormation type.
func (AWS_Amplify_App_BasicAuthConfig) GetType() string {
	return AWS_Amplify_App_BasicAuthConfig__Type
}

// Set__EnableBasicAuth updates property "EnableBasicAuth".
func (t AWS_Amplify_App_BasicAuthConfig) Set__EnableBasicAuth(v cfz.Expression[bool]) AWS_Amplify_App_BasicAuthConfig {
	t.EnableBasicAuth = v
	return t
}

// SetV__EnableBasicAuth updates property "EnableBasicAuth".
func (t AWS_Amplify_App_BasicAuthConfig) SetV__EnableBasicAuth(v bool) AWS_Amplify_App_BasicAuthConfig {
	t.EnableBasicAuth = cfz.V(v)
	return t
}

// Set__Password updates property "Password".
func (t AWS_Amplify_App_BasicAuthConfig) Set__Password(v cfz.Expression[string]) AWS_Amplify_App_BasicAuthConfig {
	t.Password = v
	return t
}

// SetV__Password updates property "Password".
func (t AWS_Amplify_App_BasicAuthConfig) SetV__Password(v string) AWS_Amplify_App_BasicAuthConfig {
	t.Password = cfz.V(v)
	return t
}

// Set__Username updates property "Username".
func (t AWS_Amplify_App_BasicAuthConfig) Set__Username(v cfz.Expression[string]) AWS_Amplify_App_BasicAuthConfig {
	t.Username = v
	return t
}

// SetV__Username updates property "Username".
func (t AWS_Amplify_App_BasicAuthConfig) SetV__Username(v string) AWS_Amplify_App_BasicAuthConfig {
	t.Username = cfz.V(v)
	return t
}
