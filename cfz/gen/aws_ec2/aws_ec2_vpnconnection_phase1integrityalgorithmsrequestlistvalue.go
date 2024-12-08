// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__Type is the CloudFormation type for AWS::EC2::VPNConnection.Phase1IntegrityAlgorithmsRequestListValue.
	AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__Type = "AWS::EC2::VPNConnection.Phase1IntegrityAlgorithmsRequestListValue"
)

var (
	// AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VPNConnection.Phase1IntegrityAlgorithmsRequestListValue.
	AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__PropertiesMap = struct {
		Value string
	}{
		Value: "Value",
	}

	// AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VPNConnection.Phase1IntegrityAlgorithmsRequestListValue.
	AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__PropertiesSlice = []string{
		AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__PropertiesMap.Value,
	}
)

// AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue is a binding for AWS::EC2::VPNConnection.Phase1IntegrityAlgorithmsRequestListValue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue.html
type AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue struct {
	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue.html#cfn-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue initializes a new AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue.
func New__AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue() AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue {
	return AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue) GetType() string {
	return AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue__Type
}

// Set__Value updates property "Value".
func (t AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue) Set__Value(v cfz.Expression[string]) AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue) SetV__Value(v string) AWS_EC2_VPNConnection_Phase1IntegrityAlgorithmsRequestListValue {
	t.Value = cfz.V(v)
	return t
}
