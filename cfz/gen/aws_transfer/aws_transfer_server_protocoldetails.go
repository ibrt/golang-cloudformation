// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_transfer

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Transfer_Server_ProtocolDetails__Type is the CloudFormation type for AWS::Transfer::Server.ProtocolDetails.
	AWS_Transfer_Server_ProtocolDetails__Type = "AWS::Transfer::Server.ProtocolDetails"
)

var (
	// AWS_Transfer_Server_ProtocolDetails__PropertiesMap reports all the CloudFormation properties for AWS::Transfer::Server.ProtocolDetails.
	AWS_Transfer_Server_ProtocolDetails__PropertiesMap = struct {
		As2Transports            string
		PassiveIp                string
		SetStatOption            string
		TlsSessionResumptionMode string
	}{
		As2Transports:            "As2Transports",
		PassiveIp:                "PassiveIp",
		SetStatOption:            "SetStatOption",
		TlsSessionResumptionMode: "TlsSessionResumptionMode",
	}

	// AWS_Transfer_Server_ProtocolDetails__PropertiesSlice reports all the CloudFormation properties for AWS::Transfer::Server.ProtocolDetails.
	AWS_Transfer_Server_ProtocolDetails__PropertiesSlice = []string{
		AWS_Transfer_Server_ProtocolDetails__PropertiesMap.As2Transports,
		AWS_Transfer_Server_ProtocolDetails__PropertiesMap.PassiveIp,
		AWS_Transfer_Server_ProtocolDetails__PropertiesMap.SetStatOption,
		AWS_Transfer_Server_ProtocolDetails__PropertiesMap.TlsSessionResumptionMode,
	}
)

// AWS_Transfer_Server_ProtocolDetails is a binding for AWS::Transfer::Server.ProtocolDetails.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html
type AWS_Transfer_Server_ProtocolDetails struct {
	// As2Transports is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-as2transports
	As2Transports cfz.ExpressionSlice[string] `json:"As2Transports,omitempty"`

	// PassiveIp is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-passiveip
	PassiveIp cfz.Expression[string] `json:"PassiveIp,omitempty"`

	// SetStatOption is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-setstatoption
	SetStatOption cfz.Expression[string] `json:"SetStatOption,omitempty"`

	// TlsSessionResumptionMode is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-protocoldetails.html#cfn-transfer-server-protocoldetails-tlssessionresumptionmode
	TlsSessionResumptionMode cfz.Expression[string] `json:"TlsSessionResumptionMode,omitempty"`
}

// New__AWS_Transfer_Server_ProtocolDetails initializes a new AWS_Transfer_Server_ProtocolDetails.
func New__AWS_Transfer_Server_ProtocolDetails() AWS_Transfer_Server_ProtocolDetails {
	return AWS_Transfer_Server_ProtocolDetails{}
}

// GetType returns the CloudFormation type.
func (AWS_Transfer_Server_ProtocolDetails) GetType() string {
	return AWS_Transfer_Server_ProtocolDetails__Type
}

// Set__As2Transports updates property "As2Transports".
func (t AWS_Transfer_Server_ProtocolDetails) Set__As2Transports(v cfz.ExpressionSlice[string]) AWS_Transfer_Server_ProtocolDetails {
	t.As2Transports = v
	return t
}

// SetS__As2Transports updates property "As2Transports".
func (t AWS_Transfer_Server_ProtocolDetails) SetS__As2Transports(v ...cfz.Expression[string]) AWS_Transfer_Server_ProtocolDetails {
	t.As2Transports = cfz.S(v...)
	return t
}

// SetSV__As2Transports updates property "As2Transports".
func (t AWS_Transfer_Server_ProtocolDetails) SetSV__As2Transports(v ...string) AWS_Transfer_Server_ProtocolDetails {
	t.As2Transports = cfz.SV(v...)
	return t
}

// Set__PassiveIp updates property "PassiveIp".
func (t AWS_Transfer_Server_ProtocolDetails) Set__PassiveIp(v cfz.Expression[string]) AWS_Transfer_Server_ProtocolDetails {
	t.PassiveIp = v
	return t
}

// SetV__PassiveIp updates property "PassiveIp".
func (t AWS_Transfer_Server_ProtocolDetails) SetV__PassiveIp(v string) AWS_Transfer_Server_ProtocolDetails {
	t.PassiveIp = cfz.V(v)
	return t
}

// Set__SetStatOption updates property "SetStatOption".
func (t AWS_Transfer_Server_ProtocolDetails) Set__SetStatOption(v cfz.Expression[string]) AWS_Transfer_Server_ProtocolDetails {
	t.SetStatOption = v
	return t
}

// SetV__SetStatOption updates property "SetStatOption".
func (t AWS_Transfer_Server_ProtocolDetails) SetV__SetStatOption(v string) AWS_Transfer_Server_ProtocolDetails {
	t.SetStatOption = cfz.V(v)
	return t
}

// Set__TlsSessionResumptionMode updates property "TlsSessionResumptionMode".
func (t AWS_Transfer_Server_ProtocolDetails) Set__TlsSessionResumptionMode(v cfz.Expression[string]) AWS_Transfer_Server_ProtocolDetails {
	t.TlsSessionResumptionMode = v
	return t
}

// SetV__TlsSessionResumptionMode updates property "TlsSessionResumptionMode".
func (t AWS_Transfer_Server_ProtocolDetails) SetV__TlsSessionResumptionMode(v string) AWS_Transfer_Server_ProtocolDetails {
	t.TlsSessionResumptionMode = cfz.V(v)
	return t
}
