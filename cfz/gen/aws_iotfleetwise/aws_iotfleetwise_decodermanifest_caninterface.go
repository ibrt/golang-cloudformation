// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotfleetwise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTFleetWise_DecoderManifest_CanInterface__Type is the CloudFormation type for AWS::IoTFleetWise::DecoderManifest.CanInterface.
	AWS_IoTFleetWise_DecoderManifest_CanInterface__Type = "AWS::IoTFleetWise::DecoderManifest.CanInterface"
)

var (
	// AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesMap reports all the CloudFormation properties for AWS::IoTFleetWise::DecoderManifest.CanInterface.
	AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesMap = struct {
		Name            string
		ProtocolName    string
		ProtocolVersion string
	}{
		Name:            "Name",
		ProtocolName:    "ProtocolName",
		ProtocolVersion: "ProtocolVersion",
	}

	// AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesSlice reports all the CloudFormation properties for AWS::IoTFleetWise::DecoderManifest.CanInterface.
	AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesSlice = []string{
		AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesMap.Name,
		AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesMap.ProtocolName,
		AWS_IoTFleetWise_DecoderManifest_CanInterface__PropertiesMap.ProtocolVersion,
	}
)

// AWS_IoTFleetWise_DecoderManifest_CanInterface is a binding for AWS::IoTFleetWise::DecoderManifest.CanInterface.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html
type AWS_IoTFleetWise_DecoderManifest_CanInterface struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ProtocolName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-protocolname
	ProtocolName cfz.Expression[string] `json:"ProtocolName,omitempty"`

	// ProtocolVersion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-protocolversion
	ProtocolVersion cfz.Expression[string] `json:"ProtocolVersion,omitempty"`
}

// New__AWS_IoTFleetWise_DecoderManifest_CanInterface initializes a new AWS_IoTFleetWise_DecoderManifest_CanInterface.
func New__AWS_IoTFleetWise_DecoderManifest_CanInterface() AWS_IoTFleetWise_DecoderManifest_CanInterface {
	return AWS_IoTFleetWise_DecoderManifest_CanInterface{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTFleetWise_DecoderManifest_CanInterface) GetType() string {
	return AWS_IoTFleetWise_DecoderManifest_CanInterface__Type
}

// Set__Name updates property "Name".
func (t AWS_IoTFleetWise_DecoderManifest_CanInterface) Set__Name(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_CanInterface {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_IoTFleetWise_DecoderManifest_CanInterface) SetV__Name(v string) AWS_IoTFleetWise_DecoderManifest_CanInterface {
	t.Name = cfz.V(v)
	return t
}

// Set__ProtocolName updates property "ProtocolName".
func (t AWS_IoTFleetWise_DecoderManifest_CanInterface) Set__ProtocolName(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_CanInterface {
	t.ProtocolName = v
	return t
}

// SetV__ProtocolName updates property "ProtocolName".
func (t AWS_IoTFleetWise_DecoderManifest_CanInterface) SetV__ProtocolName(v string) AWS_IoTFleetWise_DecoderManifest_CanInterface {
	t.ProtocolName = cfz.V(v)
	return t
}

// Set__ProtocolVersion updates property "ProtocolVersion".
func (t AWS_IoTFleetWise_DecoderManifest_CanInterface) Set__ProtocolVersion(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_CanInterface {
	t.ProtocolVersion = v
	return t
}

// SetV__ProtocolVersion updates property "ProtocolVersion".
func (t AWS_IoTFleetWise_DecoderManifest_CanInterface) SetV__ProtocolVersion(v string) AWS_IoTFleetWise_DecoderManifest_CanInterface {
	t.ProtocolVersion = cfz.V(v)
	return t
}
