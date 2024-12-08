// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotfleetwise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTFleetWise_SignalCatalog_Actuator__Type is the CloudFormation type for AWS::IoTFleetWise::SignalCatalog.Actuator.
	AWS_IoTFleetWise_SignalCatalog_Actuator__Type = "AWS::IoTFleetWise::SignalCatalog.Actuator"
)

var (
	// AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap reports all the CloudFormation properties for AWS::IoTFleetWise::SignalCatalog.Actuator.
	AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap = struct {
		AllowedValues      string
		AssignedValue      string
		DataType           string
		Description        string
		FullyQualifiedName string
		Max                string
		Min                string
		Unit               string
	}{
		AllowedValues:      "AllowedValues",
		AssignedValue:      "AssignedValue",
		DataType:           "DataType",
		Description:        "Description",
		FullyQualifiedName: "FullyQualifiedName",
		Max:                "Max",
		Min:                "Min",
		Unit:               "Unit",
	}

	// AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesSlice reports all the CloudFormation properties for AWS::IoTFleetWise::SignalCatalog.Actuator.
	AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesSlice = []string{
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.AllowedValues,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.AssignedValue,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.DataType,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.Description,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.FullyQualifiedName,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.Max,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.Min,
		AWS_IoTFleetWise_SignalCatalog_Actuator__PropertiesMap.Unit,
	}
)

// AWS_IoTFleetWise_SignalCatalog_Actuator is a binding for AWS::IoTFleetWise::SignalCatalog.Actuator.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html
type AWS_IoTFleetWise_SignalCatalog_Actuator struct {
	// AllowedValues is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-allowedvalues
	AllowedValues cfz.ExpressionSlice[string] `json:"AllowedValues,omitempty"`

	// AssignedValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-assignedvalue
	AssignedValue cfz.Expression[string] `json:"AssignedValue,omitempty"`

	// DataType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-datatype
	DataType cfz.Expression[string] `json:"DataType,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// FullyQualifiedName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-fullyqualifiedname
	FullyQualifiedName cfz.Expression[string] `json:"FullyQualifiedName,omitempty"`

	// Max is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-max
	Max cfz.Expression[float64] `json:"Max,omitempty"`

	// Min is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-min
	Min cfz.Expression[float64] `json:"Min,omitempty"`

	// Unit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-actuator.html#cfn-iotfleetwise-signalcatalog-actuator-unit
	Unit cfz.Expression[string] `json:"Unit,omitempty"`
}

// New__AWS_IoTFleetWise_SignalCatalog_Actuator initializes a new AWS_IoTFleetWise_SignalCatalog_Actuator.
func New__AWS_IoTFleetWise_SignalCatalog_Actuator() AWS_IoTFleetWise_SignalCatalog_Actuator {
	return AWS_IoTFleetWise_SignalCatalog_Actuator{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTFleetWise_SignalCatalog_Actuator) GetType() string {
	return AWS_IoTFleetWise_SignalCatalog_Actuator__Type
}

// Set__AllowedValues updates property "AllowedValues".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__AllowedValues(v cfz.ExpressionSlice[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.AllowedValues = v
	return t
}

// SetS__AllowedValues updates property "AllowedValues".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetS__AllowedValues(v ...cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.AllowedValues = cfz.S(v...)
	return t
}

// SetSV__AllowedValues updates property "AllowedValues".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetSV__AllowedValues(v ...string) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.AllowedValues = cfz.SV(v...)
	return t
}

// Set__AssignedValue updates property "AssignedValue".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__AssignedValue(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.AssignedValue = v
	return t
}

// SetV__AssignedValue updates property "AssignedValue".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__AssignedValue(v string) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.AssignedValue = cfz.V(v)
	return t
}

// Set__DataType updates property "DataType".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__DataType(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.DataType = v
	return t
}

// SetV__DataType updates property "DataType".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__DataType(v string) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.DataType = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__Description(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__Description(v string) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Description = cfz.V(v)
	return t
}

// Set__FullyQualifiedName updates property "FullyQualifiedName".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__FullyQualifiedName(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.FullyQualifiedName = v
	return t
}

// SetV__FullyQualifiedName updates property "FullyQualifiedName".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__FullyQualifiedName(v string) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.FullyQualifiedName = cfz.V(v)
	return t
}

// Set__Max updates property "Max".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__Max(v cfz.Expression[float64]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Max = v
	return t
}

// SetV__Max updates property "Max".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__Max(v float64) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Max = cfz.V(v)
	return t
}

// Set__Min updates property "Min".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__Min(v cfz.Expression[float64]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Min = v
	return t
}

// SetV__Min updates property "Min".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__Min(v float64) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Min = cfz.V(v)
	return t
}

// Set__Unit updates property "Unit".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) Set__Unit(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Unit = v
	return t
}

// SetV__Unit updates property "Unit".
func (t AWS_IoTFleetWise_SignalCatalog_Actuator) SetV__Unit(v string) AWS_IoTFleetWise_SignalCatalog_Actuator {
	t.Unit = cfz.V(v)
	return t
}
