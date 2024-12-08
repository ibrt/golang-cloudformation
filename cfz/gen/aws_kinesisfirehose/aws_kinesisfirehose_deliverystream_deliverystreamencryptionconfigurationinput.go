// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_kinesisfirehose

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__Type is the CloudFormation type for AWS::KinesisFirehose::DeliveryStream.DeliveryStreamEncryptionConfigurationInput.
	AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__Type = "AWS::KinesisFirehose::DeliveryStream.DeliveryStreamEncryptionConfigurationInput"
)

var (
	// AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__PropertiesMap reports all the CloudFormation properties for AWS::KinesisFirehose::DeliveryStream.DeliveryStreamEncryptionConfigurationInput.
	AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__PropertiesMap = struct {
		KeyARN  string
		KeyType string
	}{
		KeyARN:  "KeyARN",
		KeyType: "KeyType",
	}

	// AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__PropertiesSlice reports all the CloudFormation properties for AWS::KinesisFirehose::DeliveryStream.DeliveryStreamEncryptionConfigurationInput.
	AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__PropertiesSlice = []string{
		AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__PropertiesMap.KeyARN,
		AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__PropertiesMap.KeyType,
	}
)

// AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput is a binding for AWS::KinesisFirehose::DeliveryStream.DeliveryStreamEncryptionConfigurationInput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput.html
type AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput struct {
	// KeyARN is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput-keyarn
	KeyARN cfz.Expression[string] `json:"KeyARN,omitempty"`

	// KeyType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput-keytype
	KeyType cfz.Expression[string] `json:"KeyType,omitempty"`
}

// New__AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput initializes a new AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput.
func New__AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput() AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput {
	return AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput{}
}

// GetType returns the CloudFormation type.
func (AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput) GetType() string {
	return AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput__Type
}

// Set__KeyARN updates property "KeyARN".
func (t AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput) Set__KeyARN(v cfz.Expression[string]) AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput {
	t.KeyARN = v
	return t
}

// SetV__KeyARN updates property "KeyARN".
func (t AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput) SetV__KeyARN(v string) AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput {
	t.KeyARN = cfz.V(v)
	return t
}

// Set__KeyType updates property "KeyType".
func (t AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput) Set__KeyType(v cfz.Expression[string]) AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput {
	t.KeyType = v
	return t
}

// SetV__KeyType updates property "KeyType".
func (t AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput) SetV__KeyType(v string) AWS_KinesisFirehose_DeliveryStream_DeliveryStreamEncryptionConfigurationInput {
	t.KeyType = cfz.V(v)
	return t
}
