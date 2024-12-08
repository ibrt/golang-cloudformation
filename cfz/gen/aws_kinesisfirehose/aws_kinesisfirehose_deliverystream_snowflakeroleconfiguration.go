// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisfirehose

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__Type is the CloudFormation type for AWS::KinesisFirehose::DeliveryStream.SnowflakeRoleConfiguration.
	AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__Type = "AWS::KinesisFirehose::DeliveryStream.SnowflakeRoleConfiguration"
)

var (
	// AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::KinesisFirehose::DeliveryStream.SnowflakeRoleConfiguration.
	AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__PropertiesMap = struct {
		Enabled       string
		SnowflakeRole string
	}{
		Enabled:       "Enabled",
		SnowflakeRole: "SnowflakeRole",
	}

	// AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisFirehose::DeliveryStream.SnowflakeRoleConfiguration.
	AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__PropertiesSlice = []string{
		AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__PropertiesMap.Enabled,
		AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__PropertiesMap.SnowflakeRole,
	}
)

// AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration is a binding for AWS::KinesisFirehose::DeliveryStream.SnowflakeRoleConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.html
type AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration struct {
	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakeroleconfiguration-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`

	// SnowflakeRole is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakeroleconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakeroleconfiguration-snowflakerole
	SnowflakeRole cfz.Expression[string] `json:"SnowflakeRole,omitempty"`
}

// New__AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration initializes a new AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration.
func New__AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration() AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration {
	return AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration) GetType() string {
	return AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration__Type
}

// Set__Enabled updates property "Enabled".
func (t AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration) Set__Enabled(v cfz.Expression[bool]) AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration) SetV__Enabled(v bool) AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration {
	t.Enabled = cfz.V(v)
	return t
}

// Set__SnowflakeRole updates property "SnowflakeRole".
func (t AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration) Set__SnowflakeRole(v cfz.Expression[string]) AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration {
	t.SnowflakeRole = v
	return t
}

// SetV__SnowflakeRole updates property "SnowflakeRole".
func (t AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration) SetV__SnowflakeRole(v string) AWS_KinesisFirehose_DeliveryStream_SnowflakeRoleConfiguration {
	t.SnowflakeRole = cfz.V(v)
	return t
}
