// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amplifyuibuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AmplifyUIBuilder_Form_ValueMapping__Type is the CloudFormation type for AWS::AmplifyUIBuilder::Form.ValueMapping.
	AWS_AmplifyUIBuilder_Form_ValueMapping__Type = "AWS::AmplifyUIBuilder::Form.ValueMapping"
)

var (
	// AWS_AmplifyUIBuilder_Form_ValueMapping__PropertiesMap reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Form.ValueMapping.
	AWS_AmplifyUIBuilder_Form_ValueMapping__PropertiesMap = struct {
		DisplayValue string
		Value        string
	}{
		DisplayValue: "DisplayValue",
		Value:        "Value",
	}

	// AWS_AmplifyUIBuilder_Form_ValueMapping__PropertiesSlice reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Form.ValueMapping.
	AWS_AmplifyUIBuilder_Form_ValueMapping__PropertiesSlice = []string{
		AWS_AmplifyUIBuilder_Form_ValueMapping__PropertiesMap.DisplayValue,
		AWS_AmplifyUIBuilder_Form_ValueMapping__PropertiesMap.Value,
	}
)

// AWS_AmplifyUIBuilder_Form_ValueMapping is a binding for AWS::AmplifyUIBuilder::Form.ValueMapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemapping.html
type AWS_AmplifyUIBuilder_Form_ValueMapping struct {
	// DisplayValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemapping.html#cfn-amplifyuibuilder-form-valuemapping-displayvalue
	DisplayValue cfz.Expression[AWS_AmplifyUIBuilder_Form_FormInputValueProperty] `json:"DisplayValue,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-valuemapping.html#cfn-amplifyuibuilder-form-valuemapping-value
	Value cfz.Expression[AWS_AmplifyUIBuilder_Form_FormInputValueProperty] `json:"Value,omitempty"`
}

// New__AWS_AmplifyUIBuilder_Form_ValueMapping initializes a new AWS_AmplifyUIBuilder_Form_ValueMapping.
func New__AWS_AmplifyUIBuilder_Form_ValueMapping() AWS_AmplifyUIBuilder_Form_ValueMapping {
	return AWS_AmplifyUIBuilder_Form_ValueMapping{}
}

// GetType returns the CloudFormation type.
func (AWS_AmplifyUIBuilder_Form_ValueMapping) GetType() string {
	return AWS_AmplifyUIBuilder_Form_ValueMapping__Type
}

// Set__DisplayValue updates property "DisplayValue".
func (t AWS_AmplifyUIBuilder_Form_ValueMapping) Set__DisplayValue(v cfz.Expression[AWS_AmplifyUIBuilder_Form_FormInputValueProperty]) AWS_AmplifyUIBuilder_Form_ValueMapping {
	t.DisplayValue = v
	return t
}

// SetV__DisplayValue updates property "DisplayValue".
func (t AWS_AmplifyUIBuilder_Form_ValueMapping) SetV__DisplayValue(v AWS_AmplifyUIBuilder_Form_FormInputValueProperty) AWS_AmplifyUIBuilder_Form_ValueMapping {
	t.DisplayValue = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_AmplifyUIBuilder_Form_ValueMapping) Set__Value(v cfz.Expression[AWS_AmplifyUIBuilder_Form_FormInputValueProperty]) AWS_AmplifyUIBuilder_Form_ValueMapping {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_AmplifyUIBuilder_Form_ValueMapping) SetV__Value(v AWS_AmplifyUIBuilder_Form_FormInputValueProperty) AWS_AmplifyUIBuilder_Form_ValueMapping {
	t.Value = cfz.V(v)
	return t
}
