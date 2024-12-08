// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_mediaconnect

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__Type is the CloudFormation type for AWS::MediaConnect::BridgeOutput.BridgeNetworkOutput.
	AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__Type = "AWS::MediaConnect::BridgeOutput.BridgeNetworkOutput"
)

var (
	// AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap reports all the CloudFormation properties for AWS::MediaConnect::BridgeOutput.BridgeNetworkOutput.
	AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap = struct {
		IpAddress   string
		NetworkName string
		Port        string
		Protocol    string
		Ttl         string
	}{
		IpAddress:   "IpAddress",
		NetworkName: "NetworkName",
		Port:        "Port",
		Protocol:    "Protocol",
		Ttl:         "Ttl",
	}

	// AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesSlice reports all the CloudFormation properties for AWS::MediaConnect::BridgeOutput.BridgeNetworkOutput.
	AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesSlice = []string{
		AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap.IpAddress,
		AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap.NetworkName,
		AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap.Port,
		AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap.Protocol,
		AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__PropertiesMap.Ttl,
	}
)

// AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput is a binding for AWS::MediaConnect::BridgeOutput.BridgeNetworkOutput.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.html
type AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput struct {
	// IpAddress is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.html#cfn-mediaconnect-bridgeoutput-bridgenetworkoutput-ipaddress
	IpAddress cfz.Expression[string] `json:"IpAddress,omitempty"`

	// NetworkName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.html#cfn-mediaconnect-bridgeoutput-bridgenetworkoutput-networkname
	NetworkName cfz.Expression[string] `json:"NetworkName,omitempty"`

	// Port is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.html#cfn-mediaconnect-bridgeoutput-bridgenetworkoutput-port
	Port cfz.Expression[int32] `json:"Port,omitempty"`

	// Protocol is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.html#cfn-mediaconnect-bridgeoutput-bridgenetworkoutput-protocol
	Protocol cfz.Expression[string] `json:"Protocol,omitempty"`

	// Ttl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.html#cfn-mediaconnect-bridgeoutput-bridgenetworkoutput-ttl
	Ttl cfz.Expression[int32] `json:"Ttl,omitempty"`
}

// New__AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput initializes a new AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput.
func New__AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput() AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	return AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput{}
}

// GetType returns the CloudFormation type.
func (AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) GetType() string {
	return AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput__Type
}

// Set__IpAddress updates property "IpAddress".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) Set__IpAddress(v cfz.Expression[string]) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.IpAddress = v
	return t
}

// SetV__IpAddress updates property "IpAddress".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) SetV__IpAddress(v string) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.IpAddress = cfz.V(v)
	return t
}

// Set__NetworkName updates property "NetworkName".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) Set__NetworkName(v cfz.Expression[string]) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.NetworkName = v
	return t
}

// SetV__NetworkName updates property "NetworkName".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) SetV__NetworkName(v string) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.NetworkName = cfz.V(v)
	return t
}

// Set__Port updates property "Port".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) Set__Port(v cfz.Expression[int32]) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.Port = v
	return t
}

// SetV__Port updates property "Port".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) SetV__Port(v int32) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.Port = cfz.V(v)
	return t
}

// Set__Protocol updates property "Protocol".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) Set__Protocol(v cfz.Expression[string]) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.Protocol = v
	return t
}

// SetV__Protocol updates property "Protocol".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) SetV__Protocol(v string) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.Protocol = cfz.V(v)
	return t
}

// Set__Ttl updates property "Ttl".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) Set__Ttl(v cfz.Expression[int32]) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.Ttl = v
	return t
}

// SetV__Ttl updates property "Ttl".
func (t AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput) SetV__Ttl(v int32) AWS_MediaConnect_BridgeOutput_BridgeNetworkOutput {
	t.Ttl = cfz.V(v)
	return t
}
