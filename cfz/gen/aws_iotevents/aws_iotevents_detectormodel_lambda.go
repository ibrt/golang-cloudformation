// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotevents

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTEvents_DetectorModel_Lambda__Type is the CloudFormation type for AWS::IoTEvents::DetectorModel.Lambda.
	AWS_IoTEvents_DetectorModel_Lambda__Type = "AWS::IoTEvents::DetectorModel.Lambda"
)

var (
	// AWS_IoTEvents_DetectorModel_Lambda__PropertiesMap reports all the CloudFormation properties for AWS::IoTEvents::DetectorModel.Lambda.
	AWS_IoTEvents_DetectorModel_Lambda__PropertiesMap = struct {
		FunctionArn string
		Payload     string
	}{
		FunctionArn: "FunctionArn",
		Payload:     "Payload",
	}

	// AWS_IoTEvents_DetectorModel_Lambda__PropertiesSlice reports all the CloudFormation properties for AWS::IoTEvents::DetectorModel.Lambda.
	AWS_IoTEvents_DetectorModel_Lambda__PropertiesSlice = []string{
		AWS_IoTEvents_DetectorModel_Lambda__PropertiesMap.FunctionArn,
		AWS_IoTEvents_DetectorModel_Lambda__PropertiesMap.Payload,
	}
)

// AWS_IoTEvents_DetectorModel_Lambda is a binding for AWS::IoTEvents::DetectorModel.Lambda.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-lambda.html
type AWS_IoTEvents_DetectorModel_Lambda struct {
	// FunctionArn is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-lambda.html#cfn-iotevents-detectormodel-lambda-functionarn
	FunctionArn cfz.Expression[string] `json:"FunctionArn,omitempty"`

	// Payload is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-lambda.html#cfn-iotevents-detectormodel-lambda-payload
	Payload cfz.Expression[AWS_IoTEvents_DetectorModel_Payload] `json:"Payload,omitempty"`
}

// New__AWS_IoTEvents_DetectorModel_Lambda initializes a new AWS_IoTEvents_DetectorModel_Lambda.
func New__AWS_IoTEvents_DetectorModel_Lambda() AWS_IoTEvents_DetectorModel_Lambda {
	return AWS_IoTEvents_DetectorModel_Lambda{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTEvents_DetectorModel_Lambda) GetType() string {
	return AWS_IoTEvents_DetectorModel_Lambda__Type
}

// Set__FunctionArn updates property "FunctionArn".
func (t AWS_IoTEvents_DetectorModel_Lambda) Set__FunctionArn(v cfz.Expression[string]) AWS_IoTEvents_DetectorModel_Lambda {
	t.FunctionArn = v
	return t
}

// SetV__FunctionArn updates property "FunctionArn".
func (t AWS_IoTEvents_DetectorModel_Lambda) SetV__FunctionArn(v string) AWS_IoTEvents_DetectorModel_Lambda {
	t.FunctionArn = cfz.V(v)
	return t
}

// Set__Payload updates property "Payload".
func (t AWS_IoTEvents_DetectorModel_Lambda) Set__Payload(v cfz.Expression[AWS_IoTEvents_DetectorModel_Payload]) AWS_IoTEvents_DetectorModel_Lambda {
	t.Payload = v
	return t
}

// SetV__Payload updates property "Payload".
func (t AWS_IoTEvents_DetectorModel_Lambda) SetV__Payload(v AWS_IoTEvents_DetectorModel_Payload) AWS_IoTEvents_DetectorModel_Lambda {
	t.Payload = cfz.V(v)
	return t
}
