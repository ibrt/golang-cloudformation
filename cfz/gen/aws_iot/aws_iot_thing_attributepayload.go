// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_Thing_AttributePayload__Type is the CloudFormation type for AWS::IoT::Thing.AttributePayload.
	AWS_IoT_Thing_AttributePayload__Type = "AWS::IoT::Thing.AttributePayload"
)

var (
	// AWS_IoT_Thing_AttributePayload__PropertiesMap reports all the CloudFormation properties for AWS::IoT::Thing.AttributePayload.
	AWS_IoT_Thing_AttributePayload__PropertiesMap = struct {
		Attributes string
	}{
		Attributes: "Attributes",
	}

	// AWS_IoT_Thing_AttributePayload__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::Thing.AttributePayload.
	AWS_IoT_Thing_AttributePayload__PropertiesSlice = []string{
		AWS_IoT_Thing_AttributePayload__PropertiesMap.Attributes,
	}
)

// AWS_IoT_Thing_AttributePayload is a binding for AWS::IoT::Thing.AttributePayload.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html
type AWS_IoT_Thing_AttributePayload struct {
	// Attributes is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html#cfn-iot-thing-attributepayload-attributes
	Attributes cfz.ExpressionMap[string] `json:"Attributes,omitempty"`
}

// New__AWS_IoT_Thing_AttributePayload initializes a new AWS_IoT_Thing_AttributePayload.
func New__AWS_IoT_Thing_AttributePayload() AWS_IoT_Thing_AttributePayload {
	return AWS_IoT_Thing_AttributePayload{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_Thing_AttributePayload) GetType() string {
	return AWS_IoT_Thing_AttributePayload__Type
}

// Set__Attributes updates property "Attributes".
func (t AWS_IoT_Thing_AttributePayload) Set__Attributes(v cfz.ExpressionMap[string]) AWS_IoT_Thing_AttributePayload {
	t.Attributes = v
	return t
}

// SetM__Attributes updates property "Attributes".
func (t AWS_IoT_Thing_AttributePayload) SetM__Attributes(v ...map[string]cfz.Expression[string]) AWS_IoT_Thing_AttributePayload {
	t.Attributes = cfz.M(v...)
	return t
}

// SetMV__Attributes updates property "Attributes".
func (t AWS_IoT_Thing_AttributePayload) SetMV__Attributes(v ...map[string]string) AWS_IoT_Thing_AttributePayload {
	t.Attributes = cfz.MV(v...)
	return t
}
