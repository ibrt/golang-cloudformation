// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_Instance_CreditSpecification__Type is the CloudFormation type for AWS::EC2::Instance.CreditSpecification.
	AWS_EC2_Instance_CreditSpecification__Type = "AWS::EC2::Instance.CreditSpecification"
)

var (
	// AWS_EC2_Instance_CreditSpecification__PropertiesMap reports all the CloudFormation properties for AWS::EC2::Instance.CreditSpecification.
	AWS_EC2_Instance_CreditSpecification__PropertiesMap = struct {
		CPUCredits string
	}{
		CPUCredits: "CPUCredits",
	}

	// AWS_EC2_Instance_CreditSpecification__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::Instance.CreditSpecification.
	AWS_EC2_Instance_CreditSpecification__PropertiesSlice = []string{
		AWS_EC2_Instance_CreditSpecification__PropertiesMap.CPUCredits,
	}
)

// AWS_EC2_Instance_CreditSpecification is a binding for AWS::EC2::Instance.CreditSpecification.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-creditspecification.html
type AWS_EC2_Instance_CreditSpecification struct {
	// CPUCredits is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-creditspecification.html#cfn-ec2-instance-creditspecification-cpucredits
	CPUCredits cfz.Expression[string] `json:"CPUCredits,omitempty"`
}

// New__AWS_EC2_Instance_CreditSpecification initializes a new AWS_EC2_Instance_CreditSpecification.
func New__AWS_EC2_Instance_CreditSpecification() AWS_EC2_Instance_CreditSpecification {
	return AWS_EC2_Instance_CreditSpecification{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_Instance_CreditSpecification) GetType() string {
	return AWS_EC2_Instance_CreditSpecification__Type
}

// Set__CPUCredits updates property "CPUCredits".
func (t AWS_EC2_Instance_CreditSpecification) Set__CPUCredits(v cfz.Expression[string]) AWS_EC2_Instance_CreditSpecification {
	t.CPUCredits = v
	return t
}

// SetV__CPUCredits updates property "CPUCredits".
func (t AWS_EC2_Instance_CreditSpecification) SetV__CPUCredits(v string) AWS_EC2_Instance_CreditSpecification {
	t.CPUCredits = cfz.V(v)
	return t
}
