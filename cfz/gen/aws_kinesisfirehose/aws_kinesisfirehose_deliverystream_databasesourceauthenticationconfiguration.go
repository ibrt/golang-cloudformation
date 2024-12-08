// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisfirehose

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__Type is the CloudFormation type for AWS::KinesisFirehose::DeliveryStream.DatabaseSourceAuthenticationConfiguration.
	AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__Type = "AWS::KinesisFirehose::DeliveryStream.DatabaseSourceAuthenticationConfiguration"
)

var (
	// AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::KinesisFirehose::DeliveryStream.DatabaseSourceAuthenticationConfiguration.
	AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__PropertiesMap = struct {
		SecretsManagerConfiguration string
	}{
		SecretsManagerConfiguration: "SecretsManagerConfiguration",
	}

	// AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisFirehose::DeliveryStream.DatabaseSourceAuthenticationConfiguration.
	AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__PropertiesSlice = []string{
		AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__PropertiesMap.SecretsManagerConfiguration,
	}
)

// AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration is a binding for AWS::KinesisFirehose::DeliveryStream.DatabaseSourceAuthenticationConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration.html
type AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration struct {
	// SecretsManagerConfiguration is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration.html#cfn-kinesisfirehose-deliverystream-databasesourceauthenticationconfiguration-secretsmanagerconfiguration
	SecretsManagerConfiguration cfz.Expression[AWS_KinesisFirehose_DeliveryStream_SecretsManagerConfiguration] `json:"SecretsManagerConfiguration,omitempty"`
}

// New__AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration initializes a new AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration.
func New__AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration() AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration {
	return AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration) GetType() string {
	return AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration__Type
}

// Set__SecretsManagerConfiguration updates property "SecretsManagerConfiguration".
func (t AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration) Set__SecretsManagerConfiguration(v cfz.Expression[AWS_KinesisFirehose_DeliveryStream_SecretsManagerConfiguration]) AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration {
	t.SecretsManagerConfiguration = v
	return t
}

// SetV__SecretsManagerConfiguration updates property "SecretsManagerConfiguration".
func (t AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration) SetV__SecretsManagerConfiguration(v AWS_KinesisFirehose_DeliveryStream_SecretsManagerConfiguration) AWS_KinesisFirehose_DeliveryStream_DatabaseSourceAuthenticationConfiguration {
	t.SecretsManagerConfiguration = cfz.V(v)
	return t
}
