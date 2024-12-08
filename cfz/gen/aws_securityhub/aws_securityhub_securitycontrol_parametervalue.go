// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_securityhub

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_SecurityHub_SecurityControl_ParameterValue__Type is the CloudFormation type for AWS::SecurityHub::SecurityControl.ParameterValue.
	AWS_SecurityHub_SecurityControl_ParameterValue__Type = "AWS::SecurityHub::SecurityControl.ParameterValue"
)

var (
	// AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap reports all the CloudFormation properties for AWS::SecurityHub::SecurityControl.ParameterValue.
	AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap = struct {
		Boolean     string
		Double      string
		Enum        string
		EnumList    string
		Integer     string
		IntegerList string
		String      string
		StringList  string
	}{
		Boolean:     "Boolean",
		Double:      "Double",
		Enum:        "Enum",
		EnumList:    "EnumList",
		Integer:     "Integer",
		IntegerList: "IntegerList",
		String:      "String",
		StringList:  "StringList",
	}

	// AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesSlice reports all the CloudFormation properties for AWS::SecurityHub::SecurityControl.ParameterValue.
	AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesSlice = []string{
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.Boolean,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.Double,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.Enum,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.EnumList,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.Integer,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.IntegerList,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.String,
		AWS_SecurityHub_SecurityControl_ParameterValue__PropertiesMap.StringList,
	}
)

// AWS_SecurityHub_SecurityControl_ParameterValue is a binding for AWS::SecurityHub::SecurityControl.ParameterValue.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html
type AWS_SecurityHub_SecurityControl_ParameterValue struct {
	// Boolean is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-boolean
	Boolean cfz.Expression[bool] `json:"Boolean,omitempty"`

	// Double is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-double
	Double cfz.Expression[float64] `json:"Double,omitempty"`

	// Enum is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-enum
	Enum cfz.Expression[string] `json:"Enum,omitempty"`

	// EnumList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-enumlist
	EnumList cfz.ExpressionSlice[string] `json:"EnumList,omitempty"`

	// Integer is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-integer
	Integer cfz.Expression[int32] `json:"Integer,omitempty"`

	// IntegerList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-integerlist
	IntegerList cfz.ExpressionSlice[int32] `json:"IntegerList,omitempty"`

	// String is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-string
	String cfz.Expression[string] `json:"String,omitempty"`

	// StringList is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-stringlist
	StringList cfz.ExpressionSlice[string] `json:"StringList,omitempty"`
}

// New__AWS_SecurityHub_SecurityControl_ParameterValue initializes a new AWS_SecurityHub_SecurityControl_ParameterValue.
func New__AWS_SecurityHub_SecurityControl_ParameterValue() AWS_SecurityHub_SecurityControl_ParameterValue {
	return AWS_SecurityHub_SecurityControl_ParameterValue{}
}

// GetType returns the CloudFormation type.
func (AWS_SecurityHub_SecurityControl_ParameterValue) GetType() string {
	return AWS_SecurityHub_SecurityControl_ParameterValue__Type
}

// Set__Boolean updates property "Boolean".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__Boolean(v cfz.Expression[bool]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Boolean = v
	return t
}

// SetV__Boolean updates property "Boolean".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetV__Boolean(v bool) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Boolean = cfz.V(v)
	return t
}

// Set__Double updates property "Double".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__Double(v cfz.Expression[float64]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Double = v
	return t
}

// SetV__Double updates property "Double".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetV__Double(v float64) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Double = cfz.V(v)
	return t
}

// Set__Enum updates property "Enum".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__Enum(v cfz.Expression[string]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Enum = v
	return t
}

// SetV__Enum updates property "Enum".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetV__Enum(v string) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Enum = cfz.V(v)
	return t
}

// Set__EnumList updates property "EnumList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__EnumList(v cfz.ExpressionSlice[string]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.EnumList = v
	return t
}

// SetS__EnumList updates property "EnumList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetS__EnumList(v ...cfz.Expression[string]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.EnumList = cfz.S(v...)
	return t
}

// SetSV__EnumList updates property "EnumList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetSV__EnumList(v ...string) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.EnumList = cfz.SV(v...)
	return t
}

// Set__Integer updates property "Integer".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__Integer(v cfz.Expression[int32]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Integer = v
	return t
}

// SetV__Integer updates property "Integer".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetV__Integer(v int32) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.Integer = cfz.V(v)
	return t
}

// Set__IntegerList updates property "IntegerList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__IntegerList(v cfz.ExpressionSlice[int32]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.IntegerList = v
	return t
}

// SetS__IntegerList updates property "IntegerList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetS__IntegerList(v ...cfz.Expression[int32]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.IntegerList = cfz.S(v...)
	return t
}

// SetSV__IntegerList updates property "IntegerList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetSV__IntegerList(v ...int32) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.IntegerList = cfz.SV(v...)
	return t
}

// Set__String updates property "String".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__String(v cfz.Expression[string]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.String = v
	return t
}

// SetV__String updates property "String".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetV__String(v string) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.String = cfz.V(v)
	return t
}

// Set__StringList updates property "StringList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) Set__StringList(v cfz.ExpressionSlice[string]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.StringList = v
	return t
}

// SetS__StringList updates property "StringList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetS__StringList(v ...cfz.Expression[string]) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.StringList = cfz.S(v...)
	return t
}

// SetSV__StringList updates property "StringList".
func (t AWS_SecurityHub_SecurityControl_ParameterValue) SetSV__StringList(v ...string) AWS_SecurityHub_SecurityControl_ParameterValue {
	t.StringList = cfz.SV(v...)
	return t
}
