// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_acmpca

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_ACMPCA_Certificate_OtherName__Type is the CloudFormation type for AWS::ACMPCA::Certificate.OtherName.
	AWS_ACMPCA_Certificate_OtherName__Type = "AWS::ACMPCA::Certificate.OtherName"
)

var (
	// AWS_ACMPCA_Certificate_OtherName__PropertiesMap reports all the CloudFormation properties for AWS::ACMPCA::Certificate.OtherName.
	AWS_ACMPCA_Certificate_OtherName__PropertiesMap = struct {
		TypeId string
		Value  string
	}{
		TypeId: "TypeId",
		Value:  "Value",
	}

	// AWS_ACMPCA_Certificate_OtherName__PropertiesSlice reports all the CloudFormation properties for AWS::ACMPCA::Certificate.OtherName.
	AWS_ACMPCA_Certificate_OtherName__PropertiesSlice = []string{
		AWS_ACMPCA_Certificate_OtherName__PropertiesMap.TypeId,
		AWS_ACMPCA_Certificate_OtherName__PropertiesMap.Value,
	}
)

// AWS_ACMPCA_Certificate_OtherName is a binding for AWS::ACMPCA::Certificate.OtherName.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-othername.html
type AWS_ACMPCA_Certificate_OtherName struct {
	// TypeId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-othername.html#cfn-acmpca-certificate-othername-typeid
	TypeId cfz.Expression[string] `json:"TypeId,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-othername.html#cfn-acmpca-certificate-othername-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_ACMPCA_Certificate_OtherName initializes a new AWS_ACMPCA_Certificate_OtherName.
func New__AWS_ACMPCA_Certificate_OtherName() AWS_ACMPCA_Certificate_OtherName {
	return AWS_ACMPCA_Certificate_OtherName{}
}

// GetType returns the CloudFormation type.
func (AWS_ACMPCA_Certificate_OtherName) GetType() string {
	return AWS_ACMPCA_Certificate_OtherName__Type
}

// Set__TypeId updates property "TypeId".
func (t AWS_ACMPCA_Certificate_OtherName) Set__TypeId(v cfz.Expression[string]) AWS_ACMPCA_Certificate_OtherName {
	t.TypeId = v
	return t
}

// SetV__TypeId updates property "TypeId".
func (t AWS_ACMPCA_Certificate_OtherName) SetV__TypeId(v string) AWS_ACMPCA_Certificate_OtherName {
	t.TypeId = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_ACMPCA_Certificate_OtherName) Set__Value(v cfz.Expression[string]) AWS_ACMPCA_Certificate_OtherName {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_ACMPCA_Certificate_OtherName) SetV__Value(v string) AWS_ACMPCA_Certificate_OtherName {
	t.Value = cfz.V(v)
	return t
}
