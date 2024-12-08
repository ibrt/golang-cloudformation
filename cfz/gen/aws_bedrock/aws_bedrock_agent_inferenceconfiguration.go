// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_bedrock

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Bedrock_Agent_InferenceConfiguration__Type is the CloudFormation type for AWS::Bedrock::Agent.InferenceConfiguration.
	AWS_Bedrock_Agent_InferenceConfiguration__Type = "AWS::Bedrock::Agent.InferenceConfiguration"
)

var (
	// AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap reports all the CloudFormation properties for AWS::Bedrock::Agent.InferenceConfiguration.
	AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap = struct {
		MaximumLength string
		StopSequences string
		Temperature   string
		TopK          string
		TopP          string
	}{
		MaximumLength: "MaximumLength",
		StopSequences: "StopSequences",
		Temperature:   "Temperature",
		TopK:          "TopK",
		TopP:          "TopP",
	}

	// AWS_Bedrock_Agent_InferenceConfiguration__PropertiesSlice reports all the CloudFormation properties for AWS::Bedrock::Agent.InferenceConfiguration.
	AWS_Bedrock_Agent_InferenceConfiguration__PropertiesSlice = []string{
		AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap.MaximumLength,
		AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap.StopSequences,
		AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap.Temperature,
		AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap.TopK,
		AWS_Bedrock_Agent_InferenceConfiguration__PropertiesMap.TopP,
	}
)

// AWS_Bedrock_Agent_InferenceConfiguration is a binding for AWS::Bedrock::Agent.InferenceConfiguration.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html
type AWS_Bedrock_Agent_InferenceConfiguration struct {
	// MaximumLength is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-maximumlength
	MaximumLength cfz.Expression[float64] `json:"MaximumLength,omitempty"`

	// StopSequences is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-stopsequences
	StopSequences cfz.ExpressionSlice[string] `json:"StopSequences,omitempty"`

	// Temperature is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-temperature
	Temperature cfz.Expression[float64] `json:"Temperature,omitempty"`

	// TopK is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-topk
	TopK cfz.Expression[float64] `json:"TopK,omitempty"`

	// TopP is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-topp
	TopP cfz.Expression[float64] `json:"TopP,omitempty"`
}

// New__AWS_Bedrock_Agent_InferenceConfiguration initializes a new AWS_Bedrock_Agent_InferenceConfiguration.
func New__AWS_Bedrock_Agent_InferenceConfiguration() AWS_Bedrock_Agent_InferenceConfiguration {
	return AWS_Bedrock_Agent_InferenceConfiguration{}
}

// GetType returns the CloudFormation type.
func (AWS_Bedrock_Agent_InferenceConfiguration) GetType() string {
	return AWS_Bedrock_Agent_InferenceConfiguration__Type
}

// Set__MaximumLength updates property "MaximumLength".
func (t AWS_Bedrock_Agent_InferenceConfiguration) Set__MaximumLength(v cfz.Expression[float64]) AWS_Bedrock_Agent_InferenceConfiguration {
	t.MaximumLength = v
	return t
}

// SetV__MaximumLength updates property "MaximumLength".
func (t AWS_Bedrock_Agent_InferenceConfiguration) SetV__MaximumLength(v float64) AWS_Bedrock_Agent_InferenceConfiguration {
	t.MaximumLength = cfz.V(v)
	return t
}

// Set__StopSequences updates property "StopSequences".
func (t AWS_Bedrock_Agent_InferenceConfiguration) Set__StopSequences(v cfz.ExpressionSlice[string]) AWS_Bedrock_Agent_InferenceConfiguration {
	t.StopSequences = v
	return t
}

// SetS__StopSequences updates property "StopSequences".
func (t AWS_Bedrock_Agent_InferenceConfiguration) SetS__StopSequences(v ...cfz.Expression[string]) AWS_Bedrock_Agent_InferenceConfiguration {
	t.StopSequences = cfz.S(v...)
	return t
}

// SetSV__StopSequences updates property "StopSequences".
func (t AWS_Bedrock_Agent_InferenceConfiguration) SetSV__StopSequences(v ...string) AWS_Bedrock_Agent_InferenceConfiguration {
	t.StopSequences = cfz.SV(v...)
	return t
}

// Set__Temperature updates property "Temperature".
func (t AWS_Bedrock_Agent_InferenceConfiguration) Set__Temperature(v cfz.Expression[float64]) AWS_Bedrock_Agent_InferenceConfiguration {
	t.Temperature = v
	return t
}

// SetV__Temperature updates property "Temperature".
func (t AWS_Bedrock_Agent_InferenceConfiguration) SetV__Temperature(v float64) AWS_Bedrock_Agent_InferenceConfiguration {
	t.Temperature = cfz.V(v)
	return t
}

// Set__TopK updates property "TopK".
func (t AWS_Bedrock_Agent_InferenceConfiguration) Set__TopK(v cfz.Expression[float64]) AWS_Bedrock_Agent_InferenceConfiguration {
	t.TopK = v
	return t
}

// SetV__TopK updates property "TopK".
func (t AWS_Bedrock_Agent_InferenceConfiguration) SetV__TopK(v float64) AWS_Bedrock_Agent_InferenceConfiguration {
	t.TopK = cfz.V(v)
	return t
}

// Set__TopP updates property "TopP".
func (t AWS_Bedrock_Agent_InferenceConfiguration) Set__TopP(v cfz.Expression[float64]) AWS_Bedrock_Agent_InferenceConfiguration {
	t.TopP = v
	return t
}

// SetV__TopP updates property "TopP".
func (t AWS_Bedrock_Agent_InferenceConfiguration) SetV__TopP(v float64) AWS_Bedrock_Agent_InferenceConfiguration {
	t.TopP = cfz.V(v)
	return t
}
