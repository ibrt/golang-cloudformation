// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_MultiMeasureAttributeMapping__Type is the CloudFormation type for AWS::Pipes::Pipe.MultiMeasureAttributeMapping.
	AWS_Pipes_Pipe_MultiMeasureAttributeMapping__Type = "AWS::Pipes::Pipe.MultiMeasureAttributeMapping"
)

var (
	// AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.MultiMeasureAttributeMapping.
	AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesMap = struct {
		MeasureValue              string
		MeasureValueType          string
		MultiMeasureAttributeName string
	}{
		MeasureValue:              "MeasureValue",
		MeasureValueType:          "MeasureValueType",
		MultiMeasureAttributeName: "MultiMeasureAttributeName",
	}

	// AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.MultiMeasureAttributeMapping.
	AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesSlice = []string{
		AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesMap.MeasureValue,
		AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesMap.MeasureValueType,
		AWS_Pipes_Pipe_MultiMeasureAttributeMapping__PropertiesMap.MultiMeasureAttributeName,
	}
)

// AWS_Pipes_Pipe_MultiMeasureAttributeMapping is a binding for AWS::Pipes::Pipe.MultiMeasureAttributeMapping.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html
type AWS_Pipes_Pipe_MultiMeasureAttributeMapping struct {
	// MeasureValue is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html#cfn-pipes-pipe-multimeasureattributemapping-measurevalue
	MeasureValue cfz.Expression[string] `json:"MeasureValue,omitempty"`

	// MeasureValueType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html#cfn-pipes-pipe-multimeasureattributemapping-measurevaluetype
	MeasureValueType cfz.Expression[string] `json:"MeasureValueType,omitempty"`

	// MultiMeasureAttributeName is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-multimeasureattributemapping.html#cfn-pipes-pipe-multimeasureattributemapping-multimeasureattributename
	MultiMeasureAttributeName cfz.Expression[string] `json:"MultiMeasureAttributeName,omitempty"`
}

// New__AWS_Pipes_Pipe_MultiMeasureAttributeMapping initializes a new AWS_Pipes_Pipe_MultiMeasureAttributeMapping.
func New__AWS_Pipes_Pipe_MultiMeasureAttributeMapping() AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	return AWS_Pipes_Pipe_MultiMeasureAttributeMapping{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_MultiMeasureAttributeMapping) GetType() string {
	return AWS_Pipes_Pipe_MultiMeasureAttributeMapping__Type
}

// Set__MeasureValue updates property "MeasureValue".
func (t AWS_Pipes_Pipe_MultiMeasureAttributeMapping) Set__MeasureValue(v cfz.Expression[string]) AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	t.MeasureValue = v
	return t
}

// SetV__MeasureValue updates property "MeasureValue".
func (t AWS_Pipes_Pipe_MultiMeasureAttributeMapping) SetV__MeasureValue(v string) AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	t.MeasureValue = cfz.V(v)
	return t
}

// Set__MeasureValueType updates property "MeasureValueType".
func (t AWS_Pipes_Pipe_MultiMeasureAttributeMapping) Set__MeasureValueType(v cfz.Expression[string]) AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	t.MeasureValueType = v
	return t
}

// SetV__MeasureValueType updates property "MeasureValueType".
func (t AWS_Pipes_Pipe_MultiMeasureAttributeMapping) SetV__MeasureValueType(v string) AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	t.MeasureValueType = cfz.V(v)
	return t
}

// Set__MultiMeasureAttributeName updates property "MultiMeasureAttributeName".
func (t AWS_Pipes_Pipe_MultiMeasureAttributeMapping) Set__MultiMeasureAttributeName(v cfz.Expression[string]) AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	t.MultiMeasureAttributeName = v
	return t
}

// SetV__MultiMeasureAttributeName updates property "MultiMeasureAttributeName".
func (t AWS_Pipes_Pipe_MultiMeasureAttributeMapping) SetV__MultiMeasureAttributeName(v string) AWS_Pipes_Pipe_MultiMeasureAttributeMapping {
	t.MultiMeasureAttributeName = cfz.V(v)
	return t
}
