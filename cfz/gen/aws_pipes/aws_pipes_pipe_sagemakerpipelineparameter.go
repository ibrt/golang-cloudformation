// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_SageMakerPipelineParameter__Type is the CloudFormation type for AWS::Pipes::Pipe.SageMakerPipelineParameter.
	AWS_Pipes_Pipe_SageMakerPipelineParameter__Type = "AWS::Pipes::Pipe.SageMakerPipelineParameter"
)

var (
	// AWS_Pipes_Pipe_SageMakerPipelineParameter__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.SageMakerPipelineParameter.
	AWS_Pipes_Pipe_SageMakerPipelineParameter__PropertiesMap = struct {
		Name  string
		Value string
	}{
		Name:  "Name",
		Value: "Value",
	}

	// AWS_Pipes_Pipe_SageMakerPipelineParameter__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.SageMakerPipelineParameter.
	AWS_Pipes_Pipe_SageMakerPipelineParameter__PropertiesSlice = []string{
		AWS_Pipes_Pipe_SageMakerPipelineParameter__PropertiesMap.Name,
		AWS_Pipes_Pipe_SageMakerPipelineParameter__PropertiesMap.Value,
	}
)

// AWS_Pipes_Pipe_SageMakerPipelineParameter is a binding for AWS::Pipes::Pipe.SageMakerPipelineParameter.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-sagemakerpipelineparameter.html
type AWS_Pipes_Pipe_SageMakerPipelineParameter struct {
	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-sagemakerpipelineparameter.html#cfn-pipes-pipe-sagemakerpipelineparameter-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// Value is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-sagemakerpipelineparameter.html#cfn-pipes-pipe-sagemakerpipelineparameter-value
	Value cfz.Expression[string] `json:"Value,omitempty"`
}

// New__AWS_Pipes_Pipe_SageMakerPipelineParameter initializes a new AWS_Pipes_Pipe_SageMakerPipelineParameter.
func New__AWS_Pipes_Pipe_SageMakerPipelineParameter() AWS_Pipes_Pipe_SageMakerPipelineParameter {
	return AWS_Pipes_Pipe_SageMakerPipelineParameter{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_SageMakerPipelineParameter) GetType() string {
	return AWS_Pipes_Pipe_SageMakerPipelineParameter__Type
}

// Set__Name updates property "Name".
func (t AWS_Pipes_Pipe_SageMakerPipelineParameter) Set__Name(v cfz.Expression[string]) AWS_Pipes_Pipe_SageMakerPipelineParameter {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Pipes_Pipe_SageMakerPipelineParameter) SetV__Name(v string) AWS_Pipes_Pipe_SageMakerPipelineParameter {
	t.Name = cfz.V(v)
	return t
}

// Set__Value updates property "Value".
func (t AWS_Pipes_Pipe_SageMakerPipelineParameter) Set__Value(v cfz.Expression[string]) AWS_Pipes_Pipe_SageMakerPipelineParameter {
	t.Value = v
	return t
}

// SetV__Value updates property "Value".
func (t AWS_Pipes_Pipe_SageMakerPipelineParameter) SetV__Value(v string) AWS_Pipes_Pipe_SageMakerPipelineParameter {
	t.Value = cfz.V(v)
	return t
}
