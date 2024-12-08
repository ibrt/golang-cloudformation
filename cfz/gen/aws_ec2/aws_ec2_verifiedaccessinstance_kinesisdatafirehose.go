// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__Type is the CloudFormation type for AWS::EC2::VerifiedAccessInstance.KinesisDataFirehose.
	AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__Type = "AWS::EC2::VerifiedAccessInstance.KinesisDataFirehose"
)

var (
	// AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VerifiedAccessInstance.KinesisDataFirehose.
	AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__PropertiesMap = struct {
		DeliveryStream string
		Enabled        string
	}{
		DeliveryStream: "DeliveryStream",
		Enabled:        "Enabled",
	}

	// AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VerifiedAccessInstance.KinesisDataFirehose.
	AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__PropertiesSlice = []string{
		AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__PropertiesMap.DeliveryStream,
		AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__PropertiesMap.Enabled,
	}
)

// AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose is a binding for AWS::EC2::VerifiedAccessInstance.KinesisDataFirehose.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.html
type AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose struct {
	// DeliveryStream is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.html#cfn-ec2-verifiedaccessinstance-kinesisdatafirehose-deliverystream
	DeliveryStream cfz.Expression[string] `json:"DeliveryStream,omitempty"`

	// Enabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessinstance-kinesisdatafirehose.html#cfn-ec2-verifiedaccessinstance-kinesisdatafirehose-enabled
	Enabled cfz.Expression[bool] `json:"Enabled,omitempty"`
}

// New__AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose initializes a new AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose.
func New__AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose() AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose {
	return AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose) GetType() string {
	return AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose__Type
}

// Set__DeliveryStream updates property "DeliveryStream".
func (t AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose) Set__DeliveryStream(v cfz.Expression[string]) AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose {
	t.DeliveryStream = v
	return t
}

// SetV__DeliveryStream updates property "DeliveryStream".
func (t AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose) SetV__DeliveryStream(v string) AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose {
	t.DeliveryStream = cfz.V(v)
	return t
}

// Set__Enabled updates property "Enabled".
func (t AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose) Set__Enabled(v cfz.Expression[bool]) AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose {
	t.Enabled = v
	return t
}

// SetV__Enabled updates property "Enabled".
func (t AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose) SetV__Enabled(v bool) AWS_EC2_VerifiedAccessInstance_KinesisDataFirehose {
	t.Enabled = cfz.V(v)
	return t
}
