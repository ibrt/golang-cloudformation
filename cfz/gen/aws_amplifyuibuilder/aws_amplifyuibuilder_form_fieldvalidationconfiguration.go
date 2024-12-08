// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_amplifyuibuilder

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__Type is the CloudFormation type for AWS::AmplifyUIBuilder::Form.FieldValidationConfiguration.
	AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__Type = "AWS::AmplifyUIBuilder::Form.FieldValidationConfiguration"
)

var (
	// AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Form.FieldValidationConfiguration.
	AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesMap = struct {
		NumValues         string
		StrValues         string
		Type              string
		ValidationMessage string
	}{
		NumValues:         "NumValues",
		StrValues:         "StrValues",
		Type:              "Type",
		ValidationMessage: "ValidationMessage",
	}

	// AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::AmplifyUIBuilder::Form.FieldValidationConfiguration.
	AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesSlice = []string{
		AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesMap.NumValues,
		AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesMap.StrValues,
		AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesMap.Type,
		AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__PropertiesMap.ValidationMessage,
	}
)

// AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration is a binding for AWS::AmplifyUIBuilder::Form.FieldValidationConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html
type AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration struct {
	// NumValues is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-numvalues
	NumValues cfz.ExpressionSlice[float64] `json:"NumValues,omitempty"`

	// StrValues is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-strvalues
	StrValues cfz.ExpressionSlice[string] `json:"StrValues,omitempty"`

	// Type is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-type
	Type cfz.Expression[string] `json:"Type,omitempty"`

	// ValidationMessage is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-validationmessage
	ValidationMessage cfz.Expression[string] `json:"ValidationMessage,omitempty"`
}

// New__AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration initializes a new AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration.
func New__AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration() AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	return AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) GetType() string {
	return AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration__Type
}

// Set__NumValues updates property "NumValues".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) Set__NumValues(v cfz.ExpressionSlice[float64]) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.NumValues = v
	return t
}

// SetS__NumValues updates property "NumValues".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) SetS__NumValues(v ...cfz.Expression[float64]) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.NumValues = cfz.S(v...)
	return t
}

// SetSV__NumValues updates property "NumValues".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) SetSV__NumValues(v ...float64) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.NumValues = cfz.SV(v...)
	return t
}

// Set__StrValues updates property "StrValues".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) Set__StrValues(v cfz.ExpressionSlice[string]) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.StrValues = v
	return t
}

// SetS__StrValues updates property "StrValues".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) SetS__StrValues(v ...cfz.Expression[string]) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.StrValues = cfz.S(v...)
	return t
}

// SetSV__StrValues updates property "StrValues".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) SetSV__StrValues(v ...string) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.StrValues = cfz.SV(v...)
	return t
}

// Set__Type updates property "Type".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) Set__Type(v cfz.Expression[string]) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.Type = v
	return t
}

// SetV__Type updates property "Type".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) SetV__Type(v string) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.Type = cfz.V(v)
	return t
}

// Set__ValidationMessage updates property "ValidationMessage".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) Set__ValidationMessage(v cfz.Expression[string]) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.ValidationMessage = v
	return t
}

// SetV__ValidationMessage updates property "ValidationMessage".
func (t AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration) SetV__ValidationMessage(v string) AWS_AmplifyUIBuilder_Form_FieldValidationConfiguration {
	t.ValidationMessage = cfz.V(v)
	return t
}
