// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotevents

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTEvents_DetectorModel_Sqs__Type is the CloudFormation type for AWS::IoTEvents::DetectorModel.Sqs.
	AWS_IoTEvents_DetectorModel_Sqs__Type = "AWS::IoTEvents::DetectorModel.Sqs"
)

var (
	// AWS_IoTEvents_DetectorModel_Sqs__PropertiesMap reports all the CloudFormation properties for AWS::IoTEvents::DetectorModel.Sqs.
	AWS_IoTEvents_DetectorModel_Sqs__PropertiesMap = struct {
		Payload   string
		QueueUrl  string
		UseBase64 string
	}{
		Payload:   "Payload",
		QueueUrl:  "QueueUrl",
		UseBase64: "UseBase64",
	}

	// AWS_IoTEvents_DetectorModel_Sqs__PropertiesSlice reports all the CloudFormation properties for AWS::IoTEvents::DetectorModel.Sqs.
	AWS_IoTEvents_DetectorModel_Sqs__PropertiesSlice = []string{
		AWS_IoTEvents_DetectorModel_Sqs__PropertiesMap.Payload,
		AWS_IoTEvents_DetectorModel_Sqs__PropertiesMap.QueueUrl,
		AWS_IoTEvents_DetectorModel_Sqs__PropertiesMap.UseBase64,
	}
)

// AWS_IoTEvents_DetectorModel_Sqs is a binding for AWS::IoTEvents::DetectorModel.Sqs.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sqs.html
type AWS_IoTEvents_DetectorModel_Sqs struct {
	// Payload is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sqs.html#cfn-iotevents-detectormodel-sqs-payload
	Payload cfz.Expression[AWS_IoTEvents_DetectorModel_Payload] `json:"Payload,omitempty"`

	// QueueUrl is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sqs.html#cfn-iotevents-detectormodel-sqs-queueurl
	QueueUrl cfz.Expression[string] `json:"QueueUrl,omitempty"`

	// UseBase64 is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-sqs.html#cfn-iotevents-detectormodel-sqs-usebase64
	UseBase64 cfz.Expression[bool] `json:"UseBase64,omitempty"`
}

// New__AWS_IoTEvents_DetectorModel_Sqs initializes a new AWS_IoTEvents_DetectorModel_Sqs.
func New__AWS_IoTEvents_DetectorModel_Sqs() AWS_IoTEvents_DetectorModel_Sqs {
	return AWS_IoTEvents_DetectorModel_Sqs{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTEvents_DetectorModel_Sqs) GetType() string {
	return AWS_IoTEvents_DetectorModel_Sqs__Type
}

// Set__Payload updates property "Payload".
func (t AWS_IoTEvents_DetectorModel_Sqs) Set__Payload(v cfz.Expression[AWS_IoTEvents_DetectorModel_Payload]) AWS_IoTEvents_DetectorModel_Sqs {
	t.Payload = v
	return t
}

// SetV__Payload updates property "Payload".
func (t AWS_IoTEvents_DetectorModel_Sqs) SetV__Payload(v AWS_IoTEvents_DetectorModel_Payload) AWS_IoTEvents_DetectorModel_Sqs {
	t.Payload = cfz.V(v)
	return t
}

// Set__QueueUrl updates property "QueueUrl".
func (t AWS_IoTEvents_DetectorModel_Sqs) Set__QueueUrl(v cfz.Expression[string]) AWS_IoTEvents_DetectorModel_Sqs {
	t.QueueUrl = v
	return t
}

// SetV__QueueUrl updates property "QueueUrl".
func (t AWS_IoTEvents_DetectorModel_Sqs) SetV__QueueUrl(v string) AWS_IoTEvents_DetectorModel_Sqs {
	t.QueueUrl = cfz.V(v)
	return t
}

// Set__UseBase64 updates property "UseBase64".
func (t AWS_IoTEvents_DetectorModel_Sqs) Set__UseBase64(v cfz.Expression[bool]) AWS_IoTEvents_DetectorModel_Sqs {
	t.UseBase64 = v
	return t
}

// SetV__UseBase64 updates property "UseBase64".
func (t AWS_IoTEvents_DetectorModel_Sqs) SetV__UseBase64(v bool) AWS_IoTEvents_DetectorModel_Sqs {
	t.UseBase64 = cfz.V(v)
	return t
}
