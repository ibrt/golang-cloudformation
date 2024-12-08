// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_SingleMeasureMapping__Type is the CloudFormation type for AWS::Pipes::Pipe.SingleMeasureMapping.
	AWS_Pipes_Pipe_SingleMeasureMapping__Type = "AWS::Pipes::Pipe.SingleMeasureMapping"
)

var (
	// AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.SingleMeasureMapping.
	AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesMap = struct {
		MeasureName      string
		MeasureValue     string
		MeasureValueType string
	}{
		MeasureName:      "MeasureName",
		MeasureValue:     "MeasureValue",
		MeasureValueType: "MeasureValueType",
	}

	// AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.SingleMeasureMapping.
	AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesSlice = []string{
		AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesMap.MeasureName,
		AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesMap.MeasureValue,
		AWS_Pipes_Pipe_SingleMeasureMapping__PropertiesMap.MeasureValueType,
	}
)

// AWS_Pipes_Pipe_SingleMeasureMapping is a binding for AWS::Pipes::Pipe.SingleMeasureMapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html
type AWS_Pipes_Pipe_SingleMeasureMapping struct {
	// MeasureName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurename
	MeasureName cfz.Expression[string] `json:"MeasureName,omitempty"`

	// MeasureValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurevalue
	MeasureValue cfz.Expression[string] `json:"MeasureValue,omitempty"`

	// MeasureValueType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-singlemeasuremapping.html#cfn-pipes-pipe-singlemeasuremapping-measurevaluetype
	MeasureValueType cfz.Expression[string] `json:"MeasureValueType,omitempty"`
}

// New__AWS_Pipes_Pipe_SingleMeasureMapping initializes a new AWS_Pipes_Pipe_SingleMeasureMapping.
func New__AWS_Pipes_Pipe_SingleMeasureMapping() AWS_Pipes_Pipe_SingleMeasureMapping {
	return AWS_Pipes_Pipe_SingleMeasureMapping{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_SingleMeasureMapping) GetType() string {
	return AWS_Pipes_Pipe_SingleMeasureMapping__Type
}

// Set__MeasureName updates property "MeasureName".
func (t AWS_Pipes_Pipe_SingleMeasureMapping) Set__MeasureName(v cfz.Expression[string]) AWS_Pipes_Pipe_SingleMeasureMapping {
	t.MeasureName = v
	return t
}

// SetV__MeasureName updates property "MeasureName".
func (t AWS_Pipes_Pipe_SingleMeasureMapping) SetV__MeasureName(v string) AWS_Pipes_Pipe_SingleMeasureMapping {
	t.MeasureName = cfz.V(v)
	return t
}

// Set__MeasureValue updates property "MeasureValue".
func (t AWS_Pipes_Pipe_SingleMeasureMapping) Set__MeasureValue(v cfz.Expression[string]) AWS_Pipes_Pipe_SingleMeasureMapping {
	t.MeasureValue = v
	return t
}

// SetV__MeasureValue updates property "MeasureValue".
func (t AWS_Pipes_Pipe_SingleMeasureMapping) SetV__MeasureValue(v string) AWS_Pipes_Pipe_SingleMeasureMapping {
	t.MeasureValue = cfz.V(v)
	return t
}

// Set__MeasureValueType updates property "MeasureValueType".
func (t AWS_Pipes_Pipe_SingleMeasureMapping) Set__MeasureValueType(v cfz.Expression[string]) AWS_Pipes_Pipe_SingleMeasureMapping {
	t.MeasureValueType = v
	return t
}

// SetV__MeasureValueType updates property "MeasureValueType".
func (t AWS_Pipes_Pipe_SingleMeasureMapping) SetV__MeasureValueType(v string) AWS_Pipes_Pipe_SingleMeasureMapping {
	t.MeasureValueType = cfz.V(v)
	return t
}
