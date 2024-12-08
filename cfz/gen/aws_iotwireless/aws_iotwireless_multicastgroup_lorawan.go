// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotwireless

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTWireless_MulticastGroup_LoRaWAN__Type is the CloudFormation type for AWS::IoTWireless::MulticastGroup.LoRaWAN.
	AWS_IoTWireless_MulticastGroup_LoRaWAN__Type = "AWS::IoTWireless::MulticastGroup.LoRaWAN"
)

var (
	// AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesMap reports all the CloudFormation properties for AWS::IoTWireless::MulticastGroup.LoRaWAN.
	AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesMap = struct {
		DlClass                  string
		NumberOfDevicesInGroup   string
		NumberOfDevicesRequested string
		RfRegion                 string
	}{
		DlClass:                  "DlClass",
		NumberOfDevicesInGroup:   "NumberOfDevicesInGroup",
		NumberOfDevicesRequested: "NumberOfDevicesRequested",
		RfRegion:                 "RfRegion",
	}

	// AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesSlice reports all the CloudFormation properties for AWS::IoTWireless::MulticastGroup.LoRaWAN.
	AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesSlice = []string{
		AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesMap.DlClass,
		AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesMap.NumberOfDevicesInGroup,
		AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesMap.NumberOfDevicesRequested,
		AWS_IoTWireless_MulticastGroup_LoRaWAN__PropertiesMap.RfRegion,
	}
)

// AWS_IoTWireless_MulticastGroup_LoRaWAN is a binding for AWS::IoTWireless::MulticastGroup.LoRaWAN.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html
type AWS_IoTWireless_MulticastGroup_LoRaWAN struct {
	// DlClass is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-dlclass
	DlClass cfz.Expression[string] `json:"DlClass,omitempty"`

	// NumberOfDevicesInGroup is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-numberofdevicesingroup
	NumberOfDevicesInGroup cfz.Expression[int32] `json:"NumberOfDevicesInGroup,omitempty"`

	// NumberOfDevicesRequested is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-numberofdevicesrequested
	NumberOfDevicesRequested cfz.Expression[int32] `json:"NumberOfDevicesRequested,omitempty"`

	// RfRegion is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-multicastgroup-lorawan.html#cfn-iotwireless-multicastgroup-lorawan-rfregion
	RfRegion cfz.Expression[string] `json:"RfRegion,omitempty"`
}

// New__AWS_IoTWireless_MulticastGroup_LoRaWAN initializes a new AWS_IoTWireless_MulticastGroup_LoRaWAN.
func New__AWS_IoTWireless_MulticastGroup_LoRaWAN() AWS_IoTWireless_MulticastGroup_LoRaWAN {
	return AWS_IoTWireless_MulticastGroup_LoRaWAN{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTWireless_MulticastGroup_LoRaWAN) GetType() string {
	return AWS_IoTWireless_MulticastGroup_LoRaWAN__Type
}

// Set__DlClass updates property "DlClass".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) Set__DlClass(v cfz.Expression[string]) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.DlClass = v
	return t
}

// SetV__DlClass updates property "DlClass".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) SetV__DlClass(v string) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.DlClass = cfz.V(v)
	return t
}

// Set__NumberOfDevicesInGroup updates property "NumberOfDevicesInGroup".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) Set__NumberOfDevicesInGroup(v cfz.Expression[int32]) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.NumberOfDevicesInGroup = v
	return t
}

// SetV__NumberOfDevicesInGroup updates property "NumberOfDevicesInGroup".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) SetV__NumberOfDevicesInGroup(v int32) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.NumberOfDevicesInGroup = cfz.V(v)
	return t
}

// Set__NumberOfDevicesRequested updates property "NumberOfDevicesRequested".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) Set__NumberOfDevicesRequested(v cfz.Expression[int32]) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.NumberOfDevicesRequested = v
	return t
}

// SetV__NumberOfDevicesRequested updates property "NumberOfDevicesRequested".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) SetV__NumberOfDevicesRequested(v int32) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.NumberOfDevicesRequested = cfz.V(v)
	return t
}

// Set__RfRegion updates property "RfRegion".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) Set__RfRegion(v cfz.Expression[string]) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.RfRegion = v
	return t
}

// SetV__RfRegion updates property "RfRegion".
func (t AWS_IoTWireless_MulticastGroup_LoRaWAN) SetV__RfRegion(v string) AWS_IoTWireless_MulticastGroup_LoRaWAN {
	t.RfRegion = cfz.V(v)
	return t
}
