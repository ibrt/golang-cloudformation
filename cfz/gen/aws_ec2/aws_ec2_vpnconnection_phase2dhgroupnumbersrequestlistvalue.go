// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_ec2

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__Type is the CloudFormation type for AWS::EC2::VPNConnection.Phase2DHGroupNumbersRequestListValue.
	AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__Type = "AWS::EC2::VPNConnection.Phase2DHGroupNumbersRequestListValue"
)

var (
	// AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__PropertiesMap reports all the CloudFormation properties for AWS::EC2::VPNConnection.Phase2DHGroupNumbersRequestListValue.
	AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__PropertiesMap = struct {
		Value string
	}{
		Value: "Value",
	}

	// AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__PropertiesSlice reports all the CloudFormation properties for AWS::EC2::VPNConnection.Phase2DHGroupNumbersRequestListValue.
	AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__PropertiesSlice = []string{
		AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__PropertiesMap.Value,
	}
)

// AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue is a binding for AWS::EC2::VPNConnection.Phase2DHGroupNumbersRequestListValue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue.html
type AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue struct {
	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue.html#cfn-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue-value
	Value cfz.Expression[int32] `json:"Value,omitempty"`
}

// New__AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue initializes a new AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue.
func New__AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue() AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue {
	return AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue{}
}

// GetType returns the CloudFormation type.
func (AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue) GetType() string {
	return AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue__Type
}

// Set__Value updates property "Value".
func (t AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue) Set__Value(v cfz.Expression[int32]) AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue) SetV__Value(v int32) AWS_EC2_VPNConnection_Phase2DHGroupNumbersRequestListValue {
	t.Value = cfz.V(v)
	return t
}
