// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_iotfleetwise

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_IoTFleetWise_SignalCatalog_Sensor__Type is the CloudFormation type for AWS::IoTFleetWise::SignalCatalog.Sensor.
	AWS_IoTFleetWise_SignalCatalog_Sensor__Type = "AWS::IoTFleetWise::SignalCatalog.Sensor"
)

var (
	// AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap reports all the CloudFormation properties for AWS::IoTFleetWise::SignalCatalog.Sensor.
	AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap = struct {
		AllowedValues      string
		DataType           string
		Description        string
		FullyQualifiedName string
		Max                string
		Min                string
		Unit               string
	}{
		AllowedValues:      "AllowedValues",
		DataType:           "DataType",
		Description:        "Description",
		FullyQualifiedName: "FullyQualifiedName",
		Max:                "Max",
		Min:                "Min",
		Unit:               "Unit",
	}

	// AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesSlice reports all the CloudFormation properties for AWS::IoTFleetWise::SignalCatalog.Sensor.
	AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesSlice = []string{
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.AllowedValues,
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.DataType,
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.Description,
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.FullyQualifiedName,
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.Max,
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.Min,
		AWS_IoTFleetWise_SignalCatalog_Sensor__PropertiesMap.Unit,
	}
)

// AWS_IoTFleetWise_SignalCatalog_Sensor is a binding for AWS::IoTFleetWise::SignalCatalog.Sensor.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html
type AWS_IoTFleetWise_SignalCatalog_Sensor struct {
	// AllowedValues is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-allowedvalues
	AllowedValues cfz.ExpressionSlice[string] `json:"AllowedValues,omitempty"`

	// DataType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-datatype
	DataType cfz.Expression[string] `json:"DataType,omitempty"`

	// Description is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-description
	Description cfz.Expression[string] `json:"Description,omitempty"`

	// FullyQualifiedName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-fullyqualifiedname
	FullyQualifiedName cfz.Expression[string] `json:"FullyQualifiedName,omitempty"`

	// Max is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-max
	Max cfz.Expression[float64] `json:"Max,omitempty"`

	// Min is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-min
	Min cfz.Expression[float64] `json:"Min,omitempty"`

	// Unit is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-signalcatalog-sensor.html#cfn-iotfleetwise-signalcatalog-sensor-unit
	Unit cfz.Expression[string] `json:"Unit,omitempty"`
}

// New__AWS_IoTFleetWise_SignalCatalog_Sensor initializes a new AWS_IoTFleetWise_SignalCatalog_Sensor.
func New__AWS_IoTFleetWise_SignalCatalog_Sensor() AWS_IoTFleetWise_SignalCatalog_Sensor {
	return AWS_IoTFleetWise_SignalCatalog_Sensor{}
}

// GetType returns the CloudFormation type.
func (AWS_IoTFleetWise_SignalCatalog_Sensor) GetType() string {
	return AWS_IoTFleetWise_SignalCatalog_Sensor__Type
}

// Set__AllowedValues updates property "AllowedValues".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__AllowedValues(v cfz.ExpressionSlice[string]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.AllowedValues = v
	return t
}

// SetS__AllowedValues updates property "AllowedValues".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetS__AllowedValues(v ...cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.AllowedValues = cfz.S(v...)
	return t
}

// SetSV__AllowedValues updates property "AllowedValues".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetSV__AllowedValues(v ...string) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.AllowedValues = cfz.SV(v...)
	return t
}

// Set__DataType updates property "DataType".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__DataType(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.DataType = v
	return t
}

// SetV__DataType updates property "DataType".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetV__DataType(v string) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.DataType = cfz.V(v)
	return t
}

// Set__Description updates property "Description".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__Description(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Description = v
	return t
}

// SetV__Description updates property "Description".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetV__Description(v string) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Description = cfz.V(v)
	return t
}

// Set__FullyQualifiedName updates property "FullyQualifiedName".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__FullyQualifiedName(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.FullyQualifiedName = v
	return t
}

// SetV__FullyQualifiedName updates property "FullyQualifiedName".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetV__FullyQualifiedName(v string) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.FullyQualifiedName = cfz.V(v)
	return t
}

// Set__Max updates property "Max".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__Max(v cfz.Expression[float64]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Max = v
	return t
}

// SetV__Max updates property "Max".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetV__Max(v float64) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Max = cfz.V(v)
	return t
}

// Set__Min updates property "Min".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__Min(v cfz.Expression[float64]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Min = v
	return t
}

// SetV__Min updates property "Min".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetV__Min(v float64) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Min = cfz.V(v)
	return t
}

// Set__Unit updates property "Unit".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) Set__Unit(v cfz.Expression[string]) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Unit = v
	return t
}

// SetV__Unit updates property "Unit".
func (t AWS_IoTFleetWise_SignalCatalog_Sensor) SetV__Unit(v string) AWS_IoTFleetWise_SignalCatalog_Sensor {
	t.Unit = cfz.V(v)
	return t
}
