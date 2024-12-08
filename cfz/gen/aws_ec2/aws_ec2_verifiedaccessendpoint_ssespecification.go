// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_VerifiedAccessEndpoint_SseSpecification__Type is the CloudFormation type for AWS::EC2::VerifiedAccessEndpoint.SseSpecification.
	AWS_EC2_VerifiedAccessEndpoint_SseSpecification__Type = "AWS::EC2::VerifiedAccessEndpoint.SseSpecification"
)

var (
	// AWS_EC2_VerifiedAccessEndpoint_SseSpecification__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VerifiedAccessEndpoint.SseSpecification.
	AWS_EC2_VerifiedAccessEndpoint_SseSpecification__PropertiesMap = struct {
		CustomerManagedKeyEnabled string
		KmsKeyArn                 string
	}{
		CustomerManagedKeyEnabled: "CustomerManagedKeyEnabled",
		KmsKeyArn:                 "KmsKeyArn",
	}

	// AWS_EC2_VerifiedAccessEndpoint_SseSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VerifiedAccessEndpoint.SseSpecification.
	AWS_EC2_VerifiedAccessEndpoint_SseSpecification__PropertiesSlice = []string{
		AWS_EC2_VerifiedAccessEndpoint_SseSpecification__PropertiesMap.CustomerManagedKeyEnabled,
		AWS_EC2_VerifiedAccessEndpoint_SseSpecification__PropertiesMap.KmsKeyArn,
	}
)

// AWS_EC2_VerifiedAccessEndpoint_SseSpecification is a binding for AWS::EC2::VerifiedAccessEndpoint.SseSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html
type AWS_EC2_VerifiedAccessEndpoint_SseSpecification struct {
	// CustomerManagedKeyEnabled is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html#cfn-ec2-verifiedaccessendpoint-ssespecification-customermanagedkeyenabled
	CustomerManagedKeyEnabled cfz.Expression[bool] `json:"CustomerManagedKeyEnabled,omitempty"`

	// KmsKeyArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html#cfn-ec2-verifiedaccessendpoint-ssespecification-kmskeyarn
	KmsKeyArn cfz.Expression[string] `json:"KmsKeyArn,omitempty"`
}

// New__AWS_EC2_VerifiedAccessEndpoint_SseSpecification initializes a new AWS_EC2_VerifiedAccessEndpoint_SseSpecification.
func New__AWS_EC2_VerifiedAccessEndpoint_SseSpecification() AWS_EC2_VerifiedAccessEndpoint_SseSpecification {
	return AWS_EC2_VerifiedAccessEndpoint_SseSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_VerifiedAccessEndpoint_SseSpecification) GetType() string {
	return AWS_EC2_VerifiedAccessEndpoint_SseSpecification__Type
}

// Set__CustomerManagedKeyEnabled updates property "CustomerManagedKeyEnabled".
func (t AWS_EC2_VerifiedAccessEndpoint_SseSpecification) Set__CustomerManagedKeyEnabled(v cfz.Expression[bool]) AWS_EC2_VerifiedAccessEndpoint_SseSpecification {
	t.CustomerManagedKeyEnabled = v
	return t
}

// SetV__CustomerManagedKeyEnabled updates property "CustomerManagedKeyEnabled".
func (t AWS_EC2_VerifiedAccessEndpoint_SseSpecification) SetV__CustomerManagedKeyEnabled(v bool) AWS_EC2_VerifiedAccessEndpoint_SseSpecification {
	t.CustomerManagedKeyEnabled = cfz.V(v)
	return t
}

// Set__KmsKeyArn updates property "KmsKeyArn".
func (t AWS_EC2_VerifiedAccessEndpoint_SseSpecification) Set__KmsKeyArn(v cfz.Expression[string]) AWS_EC2_VerifiedAccessEndpoint_SseSpecification {
	t.KmsKeyArn = v
	return t
}

// SetV__KmsKeyArn updates property "KmsKeyArn".
func (t AWS_EC2_VerifiedAccessEndpoint_SseSpecification) SetV__KmsKeyArn(v string) AWS_EC2_VerifiedAccessEndpoint_SseSpecification {
	t.KmsKeyArn = cfz.V(v)
	return t
}
