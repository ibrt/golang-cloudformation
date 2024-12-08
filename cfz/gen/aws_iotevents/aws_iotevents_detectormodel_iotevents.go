// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotevents

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTEvents_DetectorModel_IotEvents__Type is the CloudFormation type for AWS::IoTEvents::DetectorModel.IotEvents.
	AWS_IoTEvents_DetectorModel_IotEvents__Type = "AWS::IoTEvents::DetectorModel.IotEvents"
)

var (
	// AWS_IoTEvents_DetectorModel_IotEvents__PropertiesMap reports all the CloudFormation properties for AWS::IoTEvents::DetectorModel.IotEvents.
	AWS_IoTEvents_DetectorModel_IotEvents__PropertiesMap = struct {
		InputName string
		Payload   string
	}{
		InputName: "InputName",
		Payload:   "Payload",
	}

	// AWS_IoTEvents_DetectorModel_IotEvents__PropertiesSlice reports all the CloudFormation properties for AWS::IoTEvents::DetectorModel.IotEvents.
	AWS_IoTEvents_DetectorModel_IotEvents__PropertiesSlice = []string{
		AWS_IoTEvents_DetectorModel_IotEvents__PropertiesMap.InputName,
		AWS_IoTEvents_DetectorModel_IotEvents__PropertiesMap.Payload,
	}
)

// AWS_IoTEvents_DetectorModel_IotEvents is a binding for AWS::IoTEvents::DetectorModel.IotEvents.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html
type AWS_IoTEvents_DetectorModel_IotEvents struct {
	// InputName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html#cfn-iotevents-detectormodel-iotevents-inputname
	InputName cfz.Expression[string] `json:"InputName,omitempty"`

	// Payload is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-iotevents.html#cfn-iotevents-detectormodel-iotevents-payload
	Payload cfz.Expression[AWS_IoTEvents_DetectorModel_Payload] `json:"Payload,omitempty"`
}

// New__AWS_IoTEvents_DetectorModel_IotEvents initializes a new AWS_IoTEvents_DetectorModel_IotEvents.
func New__AWS_IoTEvents_DetectorModel_IotEvents() AWS_IoTEvents_DetectorModel_IotEvents {
	return AWS_IoTEvents_DetectorModel_IotEvents{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTEvents_DetectorModel_IotEvents) GetType() string {
	return AWS_IoTEvents_DetectorModel_IotEvents__Type
}

// Set__InputName updates property "InputName".
func (t AWS_IoTEvents_DetectorModel_IotEvents) Set__InputName(v cfz.Expression[string]) AWS_IoTEvents_DetectorModel_IotEvents {
	t.InputName = v
	return t
}

// SetV__InputName updates property "InputName".
func (t AWS_IoTEvents_DetectorModel_IotEvents) SetV__InputName(v string) AWS_IoTEvents_DetectorModel_IotEvents {
	t.InputName = cfz.V(v)
	return t
}

// Set__Payload updates property "Payload".
func (t AWS_IoTEvents_DetectorModel_IotEvents) Set__Payload(v cfz.Expression[AWS_IoTEvents_DetectorModel_Payload]) AWS_IoTEvents_DetectorModel_IotEvents {
	t.Payload = v
	return t
}

// SetV__Payload updates property "Payload".
func (t AWS_IoTEvents_DetectorModel_IotEvents) SetV__Payload(v AWS_IoTEvents_DetectorModel_Payload) AWS_IoTEvents_DetectorModel_IotEvents {
	t.Payload = cfz.V(v)
	return t
}
