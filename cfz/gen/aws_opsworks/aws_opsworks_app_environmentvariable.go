// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_opsworks

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_OpsWorks_App_EnvironmentVariable__Type is the CloudFormation type for AWS::OpsWorks::App.EnvironmentVariable.
	AWS_OpsWorks_App_EnvironmentVariable__Type = "AWS::OpsWorks::App.EnvironmentVariable"
)

var (
	// AWS_OpsWorks_App_EnvironmentVariable__PropertiesMap reports all the CloudFormation properties for AWS::OpsWorks::App.EnvironmentVariable.
	AWS_OpsWorks_App_EnvironmentVariable__PropertiesMap = struct {
		Key    string
		Secure string
		Value  string
	}{
		Key:    "Key",
		Secure: "Secure",
		Value:  "Value",
	}

	// AWS_OpsWorks_App_EnvironmentVariable__PropertiesSlice reports all the CloudFormation properties for AWS::OpsWorks::App.EnvironmentVariable.
	AWS_OpsWorks_App_EnvironmentVariable__PropertiesSlice = []string{
		AWS_OpsWorks_App_EnvironmentVariable__PropertiesMap.Key,
		AWS_OpsWorks_App_EnvironmentVariable__PropertiesMap.Secure,
		AWS_OpsWorks_App_EnvironmentVariable__PropertiesMap.Value,
	}
)

// AWS_OpsWorks_App_EnvironmentVariable is a binding for AWS::OpsWorks::App.EnvironmentVariable.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html
type AWS_OpsWorks_App_EnvironmentVariable struct {
	// Key is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-key
	Key cfz.Expression[string] `json:"Key,omitempty"`

	// Secure is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#cfn-opsworks-app-environment-secure
	Secure cfz.Expression[bool] `json:"Secure,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-environment.html#value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_OpsWorks_App_EnvironmentVariable initializes a new AWS_OpsWorks_App_EnvironmentVariable.
func New__AWS_OpsWorks_App_EnvironmentVariable() AWS_OpsWorks_App_EnvironmentVariable {
	return AWS_OpsWorks_App_EnvironmentVariable{}
}

// GetType returns the CloudFormation type.
func (AWS_OpsWorks_App_EnvironmentVariable) GetType() string {
	return AWS_OpsWorks_App_EnvironmentVariable__Type
}

// Set__Key updates property "Key".
func (t AWS_OpsWorks_App_EnvironmentVariable) Set__Key(v cfz.Expression[string]) AWS_OpsWorks_App_EnvironmentVariable {
	t.Key = v
	return t
}

// SetV__Key updates property "Key".
func (t AWS_OpsWorks_App_EnvironmentVariable) SetV__Key(v string) AWS_OpsWorks_App_EnvironmentVariable {
	t.Key = cfz.V(v)
	return t
}

// Set__Secure updates property "Secure".
func (t AWS_OpsWorks_App_EnvironmentVariable) Set__Secure(v cfz.Expression[bool]) AWS_OpsWorks_App_EnvironmentVariable {
	t.Secure = v
	return t
}

// SetV__Secure updates property "Secure".
func (t AWS_OpsWorks_App_EnvironmentVariable) SetV__Secure(v bool) AWS_OpsWorks_App_EnvironmentVariable {
	t.Secure = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_OpsWorks_App_EnvironmentVariable) Set__Value(v cfz.Expression[string]) AWS_OpsWorks_App_EnvironmentVariable {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_OpsWorks_App_EnvironmentVariable) SetV__Value(v string) AWS_OpsWorks_App_EnvironmentVariable {
	t.Value = cfz.V(v)
	return t
}
