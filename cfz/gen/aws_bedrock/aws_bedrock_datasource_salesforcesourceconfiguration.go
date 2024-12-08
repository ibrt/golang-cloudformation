// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_DataSource_SalesforceSourceConfiguration__Type is the CloudFormation type for AWS::Bedrock::DataSource.SalesforceSourceConfiguration.
	AWS_Bedrock_DataSource_SalesforceSourceConfiguration__Type = "AWS::Bedrock::DataSource.SalesforceSourceConfiguration"
)

var (
	// AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::DataSource.SalesforceSourceConfiguration.
	AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesMap = struct {
		AuthType             string
		CredentialsSecretArn string
		HostUrl              string
	}{
		AuthType:             "AuthType",
		CredentialsSecretArn: "CredentialsSecretArn",
		HostUrl:              "HostUrl",
	}

	// AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::DataSource.SalesforceSourceConfiguration.
	AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesMap.AuthType,
		AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesMap.CredentialsSecretArn,
		AWS_Bedrock_DataSource_SalesforceSourceConfiguration__PropertiesMap.HostUrl,
	}
)

// AWS_Bedrock_DataSource_SalesforceSourceConfiguration is a binding for AWS::Bedrock::DataSource.SalesforceSourceConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html
type AWS_Bedrock_DataSource_SalesforceSourceConfiguration struct {
	// AuthType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html#cfn-bedrock-datasource-salesforcesourceconfiguration-authtype
	AuthType cfz.Expression[string] `json:"AuthType,omitempty"`

	// CredentialsSecretArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html#cfn-bedrock-datasource-salesforcesourceconfiguration-credentialssecretarn
	CredentialsSecretArn cfz.Expression[string] `json:"CredentialsSecretArn,omitempty"`

	// HostUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcesourceconfiguration.html#cfn-bedrock-datasource-salesforcesourceconfiguration-hosturl
	HostUrl cfz.Expression[string] `json:"HostUrl,omitempty"`
}

// New__AWS_Bedrock_DataSource_SalesforceSourceConfiguration initializes a new AWS_Bedrock_DataSource_SalesforceSourceConfiguration.
func New__AWS_Bedrock_DataSource_SalesforceSourceConfiguration() AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	return AWS_Bedrock_DataSource_SalesforceSourceConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_DataSource_SalesforceSourceConfiguration) GetType() string {
	return AWS_Bedrock_DataSource_SalesforceSourceConfiguration__Type
}

// Set__AuthType updates property "AuthType".
func (t AWS_Bedrock_DataSource_SalesforceSourceConfiguration) Set__AuthType(v cfz.Expression[string]) AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	t.AuthType = v
	return t
}

// SetV__AuthType updates property "AuthType".
func (t AWS_Bedrock_DataSource_SalesforceSourceConfiguration) SetV__AuthType(v string) AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	t.AuthType = cfz.V(v)
	return t
}

// Set__CredentialsSecretArn updates property "CredentialsSecretArn".
func (t AWS_Bedrock_DataSource_SalesforceSourceConfiguration) Set__CredentialsSecretArn(v cfz.Expression[string]) AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	t.CredentialsSecretArn = v
	return t
}

// SetV__CredentialsSecretArn updates property "CredentialsSecretArn".
func (t AWS_Bedrock_DataSource_SalesforceSourceConfiguration) SetV__CredentialsSecretArn(v string) AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	t.CredentialsSecretArn = cfz.V(v)
	return t
}

// Set__HostUrl updates property "HostUrl".
func (t AWS_Bedrock_DataSource_SalesforceSourceConfiguration) Set__HostUrl(v cfz.Expression[string]) AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	t.HostUrl = v
	return t
}

// SetV__HostUrl updates property "HostUrl".
func (t AWS_Bedrock_DataSource_SalesforceSourceConfiguration) SetV__HostUrl(v string) AWS_Bedrock_DataSource_SalesforceSourceConfiguration {
	t.HostUrl = cfz.V(v)
	return t
}
