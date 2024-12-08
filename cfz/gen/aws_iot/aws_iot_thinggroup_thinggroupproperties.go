// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iot

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoT_ThingGroup_ThingGroupProperties__Type is the CloudFormation type for AWS::IoT::ThingGroup.ThingGroupProperties.
	AWS_IoT_ThingGroup_ThingGroupProperties__Type = "AWS::IoT::ThingGroup.ThingGroupProperties"
)

var (
	// AWS_IoT_ThingGroup_ThingGroupProperties__PropertiesMap reports all the CloudFormation properties for AWS::IoT::ThingGroup.ThingGroupProperties.
	AWS_IoT_ThingGroup_ThingGroupProperties__PropertiesMap = struct {
		AttributePayload      string
		ThingGroupDescription string
	}{
		AttributePayload:      "AttributePayload",
		ThingGroupDescription: "ThingGroupDescription",
	}

	// AWS_IoT_ThingGroup_ThingGroupProperties__PropertiesSlice reports all the CloudFormation properties for AWS::IoT::ThingGroup.ThingGroupProperties.
	AWS_IoT_ThingGroup_ThingGroupProperties__PropertiesSlice = []string{
		AWS_IoT_ThingGroup_ThingGroupProperties__PropertiesMap.AttributePayload,
		AWS_IoT_ThingGroup_ThingGroupProperties__PropertiesMap.ThingGroupDescription,
	}
)

// AWS_IoT_ThingGroup_ThingGroupProperties is a binding for AWS::IoT::ThingGroup.ThingGroupProperties.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-thinggroupproperties.html
type AWS_IoT_ThingGroup_ThingGroupProperties struct {
	// AttributePayload is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-thinggroupproperties.html#cfn-iot-thinggroup-thinggroupproperties-attributepayload
	AttributePayload cfz.Expression[AWS_IoT_ThingGroup_AttributePayload] `json:"AttributePayload,omitempty"`

	// ThingGroupDescription is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thinggroup-thinggroupproperties.html#cfn-iot-thinggroup-thinggroupproperties-thinggroupdescription
	ThingGroupDescription cfz.Expression[string] `json:"ThingGroupDescription,omitempty"`
}

// New__AWS_IoT_ThingGroup_ThingGroupProperties initializes a new AWS_IoT_ThingGroup_ThingGroupProperties.
func New__AWS_IoT_ThingGroup_ThingGroupProperties() AWS_IoT_ThingGroup_ThingGroupProperties {
	return AWS_IoT_ThingGroup_ThingGroupProperties{}
}

// GetType returns the CloudFormation type.
func (AWS_IoT_ThingGroup_ThingGroupProperties) GetType() string {
	return AWS_IoT_ThingGroup_ThingGroupProperties__Type
}

// Set__AttributePayload updates property "AttributePayload".
func (t AWS_IoT_ThingGroup_ThingGroupProperties) Set__AttributePayload(v cfz.Expression[AWS_IoT_ThingGroup_AttributePayload]) AWS_IoT_ThingGroup_ThingGroupProperties {
	t.AttributePayload = v
	return t
}

// SetV__AttributePayload updates property "AttributePayload".
func (t AWS_IoT_ThingGroup_ThingGroupProperties) SetV__AttributePayload(v AWS_IoT_ThingGroup_AttributePayload) AWS_IoT_ThingGroup_ThingGroupProperties {
	t.AttributePayload = cfz.V(v)
	return t
}

// Set__ThingGroupDescription updates property "ThingGroupDescription".
func (t AWS_IoT_ThingGroup_ThingGroupProperties) Set__ThingGroupDescription(v cfz.Expression[string]) AWS_IoT_ThingGroup_ThingGroupProperties {
	t.ThingGroupDescription = v
	return t
}

// SetV__ThingGroupDescription updates property "ThingGroupDescription".
func (t AWS_IoT_ThingGroup_ThingGroupProperties) SetV__ThingGroupDescription(v string) AWS_IoT_ThingGroup_ThingGroupProperties {
	t.ThingGroupDescription = cfz.V(v)
	return t
}
