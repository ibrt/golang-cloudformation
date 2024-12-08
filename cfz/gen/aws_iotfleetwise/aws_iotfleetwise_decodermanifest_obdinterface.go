// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotfleetwise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTFleetWise_DecoderManifest_ObdInterface__Type is the CloudFormation type for AWS::IoTFleetWise::DecoderManifest.ObdInterface.
	AWS_IoTFleetWise_DecoderManifest_ObdInterface__Type = "AWS::IoTFleetWise::DecoderManifest.ObdInterface"
)

var (
	// AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap reports all the CloudFormation properties for AWS::IoTFleetWise::DecoderManifest.ObdInterface.
	AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap = struct {
		DtcRequestIntervalSeconds string
		HasTransmissionEcu        string
		Name                      string
		ObdStandard               string
		PidRequestIntervalSeconds string
		RequestMessageId          string
		UseExtendedIds            string
	}{
		DtcRequestIntervalSeconds: "DtcRequestIntervalSeconds",
		HasTransmissionEcu:        "HasTransmissionEcu",
		Name:                      "Name",
		ObdStandard:               "ObdStandard",
		PidRequestIntervalSeconds: "PidRequestIntervalSeconds",
		RequestMessageId:          "RequestMessageId",
		UseExtendedIds:            "UseExtendedIds",
	}

	// AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesSlice reports all the CloudFormation properties for AWS::IoTFleetWise::DecoderManifest.ObdInterface.
	AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesSlice = []string{
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.DtcRequestIntervalSeconds,
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.HasTransmissionEcu,
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.Name,
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.ObdStandard,
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.PidRequestIntervalSeconds,
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.RequestMessageId,
		AWS_IoTFleetWise_DecoderManifest_ObdInterface__PropertiesMap.UseExtendedIds,
	}
)

// AWS_IoTFleetWise_DecoderManifest_ObdInterface is a binding for AWS::IoTFleetWise::DecoderManifest.ObdInterface.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html
type AWS_IoTFleetWise_DecoderManifest_ObdInterface struct {
	// DtcRequestIntervalSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-dtcrequestintervalseconds
	DtcRequestIntervalSeconds cfz.Expression[string] `json:"DtcRequestIntervalSeconds,omitempty"`

	// HasTransmissionEcu is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-hastransmissionecu
	HasTransmissionEcu cfz.Expression[string] `json:"HasTransmissionEcu,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ObdStandard is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-obdstandard
	ObdStandard cfz.Expression[string] `json:"ObdStandard,omitempty"`

	// PidRequestIntervalSeconds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-pidrequestintervalseconds
	PidRequestIntervalSeconds cfz.Expression[string] `json:"PidRequestIntervalSeconds,omitempty"`

	// RequestMessageId is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-requestmessageid
	RequestMessageId cfz.Expression[string] `json:"RequestMessageId,omitempty"`

	// UseExtendedIds is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-useextendedids
	UseExtendedIds cfz.Expression[string] `json:"UseExtendedIds,omitempty"`
}

// New__AWS_IoTFleetWise_DecoderManifest_ObdInterface initializes a new AWS_IoTFleetWise_DecoderManifest_ObdInterface.
func New__AWS_IoTFleetWise_DecoderManifest_ObdInterface() AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	return AWS_IoTFleetWise_DecoderManifest_ObdInterface{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTFleetWise_DecoderManifest_ObdInterface) GetType() string {
	return AWS_IoTFleetWise_DecoderManifest_ObdInterface__Type
}

// Set__DtcRequestIntervalSeconds updates property "DtcRequestIntervalSeconds".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__DtcRequestIntervalSeconds(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.DtcRequestIntervalSeconds = v
	return t
}

// SetV__DtcRequestIntervalSeconds updates property "DtcRequestIntervalSeconds".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__DtcRequestIntervalSeconds(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.DtcRequestIntervalSeconds = cfz.V(v)
	return t
}

// Set__HasTransmissionEcu updates property "HasTransmissionEcu".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__HasTransmissionEcu(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.HasTransmissionEcu = v
	return t
}

// SetV__HasTransmissionEcu updates property "HasTransmissionEcu".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__HasTransmissionEcu(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.HasTransmissionEcu = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__Name(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__Name(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.Name = cfz.V(v)
	return t
}

// Set__ObdStandard updates property "ObdStandard".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__ObdStandard(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.ObdStandard = v
	return t
}

// SetV__ObdStandard updates property "ObdStandard".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__ObdStandard(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.ObdStandard = cfz.V(v)
	return t
}

// Set__PidRequestIntervalSeconds updates property "PidRequestIntervalSeconds".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__PidRequestIntervalSeconds(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.PidRequestIntervalSeconds = v
	return t
}

// SetV__PidRequestIntervalSeconds updates property "PidRequestIntervalSeconds".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__PidRequestIntervalSeconds(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.PidRequestIntervalSeconds = cfz.V(v)
	return t
}

// Set__RequestMessageId updates property "RequestMessageId".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__RequestMessageId(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.RequestMessageId = v
	return t
}

// SetV__RequestMessageId updates property "RequestMessageId".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__RequestMessageId(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.RequestMessageId = cfz.V(v)
	return t
}

// Set__UseExtendedIds updates property "UseExtendedIds".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) Set__UseExtendedIds(v cfz.Expression[string]) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.UseExtendedIds = v
	return t
}

// SetV__UseExtendedIds updates property "UseExtendedIds".
func (t AWS_IoTFleetWise_DecoderManifest_ObdInterface) SetV__UseExtendedIds(v string) AWS_IoTFleetWise_DecoderManifest_ObdInterface {
	t.UseExtendedIds = cfz.V(v)
	return t
}
