// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediatailor

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__Type is the CloudFormation type for AWS::MediaTailor::SourceLocation.SecretsManagerAccessTokenConfiguration.
	AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__Type = "AWS::MediaTailor::SourceLocation.SecretsManagerAccessTokenConfiguration"
)

var (
	// AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::MediaTailor::SourceLocation.SecretsManagerAccessTokenConfiguration.
	AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesMap = struct {
		HeaderName      string
		SecretArn       string
		SecretStringKey string
	}{
		HeaderName:      "HeaderName",
		SecretArn:       "SecretArn",
		SecretStringKey: "SecretStringKey",
	}

	// AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::MediaTailor::SourceLocation.SecretsManagerAccessTokenConfiguration.
	AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesSlice = []string{
		AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesMap.HeaderName,
		AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesMap.SecretArn,
		AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__PropertiesMap.SecretStringKey,
	}
)

// AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration is a binding for AWS::MediaTailor::SourceLocation.SecretsManagerAccessTokenConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html
type AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration struct {
	// HeaderName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html#cfn-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration-headername
	HeaderName cfz.Expression[string] `json:"HeaderName,omitempty"`

	// SecretArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html#cfn-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration-secretarn
	SecretArn cfz.Expression[string] `json:"SecretArn,omitempty"`

	// SecretStringKey is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html#cfn-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration-secretstringkey
	SecretStringKey cfz.Expression[string] `json:"SecretStringKey,omitempty"`
}

// New__AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration initializes a new AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration.
func New__AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration() AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	return AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) GetType() string {
	return AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration__Type
}

// Set__HeaderName updates property "HeaderName".
func (t AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) Set__HeaderName(v cfz.Expression[string]) AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	t.HeaderName = v
	return t
}

// SetV__HeaderName updates property "HeaderName".
func (t AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) SetV__HeaderName(v string) AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	t.HeaderName = cfz.V(v)
	return t
}

// Set__SecretArn updates property "SecretArn".
func (t AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) Set__SecretArn(v cfz.Expression[string]) AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	t.SecretArn = v
	return t
}

// SetV__SecretArn updates property "SecretArn".
func (t AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) SetV__SecretArn(v string) AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	t.SecretArn = cfz.V(v)
	return t
}

// Set__SecretStringKey updates property "SecretStringKey".
func (t AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) Set__SecretStringKey(v cfz.Expression[string]) AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	t.SecretStringKey = v
	return t
}

// SetV__SecretStringKey updates property "SecretStringKey".
func (t AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration) SetV__SecretStringKey(v string) AWS_MediaTailor_SourceLocation_SecretsManagerAccessTokenConfiguration {
	t.SecretStringKey = cfz.V(v)
	return t
}
