// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__Type is the CloudFormation type for AWS::EC2::VPNConnection.Phase1EncryptionAlgorithmsRequestListValue.
	AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__Type = "AWS::EC2::VPNConnection.Phase1EncryptionAlgorithmsRequestListValue"
)

var (
	// AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VPNConnection.Phase1EncryptionAlgorithmsRequestListValue.
	AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__PropertiesMap = struct {
		Value string
	}{
		Value: "Value",
	}

	// AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VPNConnection.Phase1EncryptionAlgorithmsRequestListValue.
	AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__PropertiesSlice = []string{
		AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__PropertiesMap.Value,
	}
)

// AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue is a binding for AWS::EC2::VPNConnection.Phase1EncryptionAlgorithmsRequestListValue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue.html
type AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue struct {
	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue.html#cfn-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue initializes a new AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue.
func New__AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue() AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue {
	return AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue) GetType() string {
	return AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue__Type
}

// Set__Value updates property "Value".
func (t AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue) Set__Value(v cfz.Expression[string]) AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue) SetV__Value(v string) AWS_EC2_VPNConnection_Phase1EncryptionAlgorithmsRequestListValue {
	t.Value = cfz.V(v)
	return t
}
