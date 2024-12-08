// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_BatchContainerOverrides__Type is the CloudFormation type for AWS::Pipes::Pipe.BatchContainerOverrides.
	AWS_Pipes_Pipe_BatchContainerOverrides__Type = "AWS::Pipes::Pipe.BatchContainerOverrides"
)

var (
	// AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.BatchContainerOverrides.
	AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesMap = struct {
		Command              string
		Environment          string
		InstanceType         string
		ResourceRequirements string
	}{
		Command:              "Command",
		Environment:          "Environment",
		InstanceType:         "InstanceType",
		ResourceRequirements: "ResourceRequirements",
	}

	// AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.BatchContainerOverrides.
	AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesSlice = []string{
		AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesMap.Command,
		AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesMap.Environment,
		AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesMap.InstanceType,
		AWS_Pipes_Pipe_BatchContainerOverrides__PropertiesMap.ResourceRequirements,
	}
)

// AWS_Pipes_Pipe_BatchContainerOverrides is a binding for AWS::Pipes::Pipe.BatchContainerOverrides.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html
type AWS_Pipes_Pipe_BatchContainerOverrides struct {
	// Command is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-command
	Command cfz.ExpressionSlice[string] `json:"Command,omitempty"`

	// Environment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-environment
	Environment cfz.ExpressionSlice[AWS_Pipes_Pipe_BatchEnvironmentVariable] `json:"Environment,omitempty"`

	// InstanceType is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-instancetype
	InstanceType cfz.Expression[string] `json:"InstanceType,omitempty"`

	// ResourceRequirements is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchcontaineroverrides.html#cfn-pipes-pipe-batchcontaineroverrides-resourcerequirements
	ResourceRequirements cfz.ExpressionSlice[AWS_Pipes_Pipe_BatchResourceRequirement] `json:"ResourceRequirements,omitempty"`
}

// New__AWS_Pipes_Pipe_BatchContainerOverrides initializes a new AWS_Pipes_Pipe_BatchContainerOverrides.
func New__AWS_Pipes_Pipe_BatchContainerOverrides() AWS_Pipes_Pipe_BatchContainerOverrides {
	return AWS_Pipes_Pipe_BatchContainerOverrides{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_BatchContainerOverrides) GetType() string {
	return AWS_Pipes_Pipe_BatchContainerOverrides__Type
}

// Set__Command updates property "Command".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) Set__Command(v cfz.ExpressionSlice[string]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.Command = v
	return t
}

// SetS__Command updates property "Command".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetS__Command(v ...cfz.Expression[string]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.Command = cfz.S(v...)
	return t
}

// SetSV__Command updates property "Command".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetSV__Command(v ...string) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.Command = cfz.SV(v...)
	return t
}

// Set__Environment updates property "Environment".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) Set__Environment(v cfz.ExpressionSlice[AWS_Pipes_Pipe_BatchEnvironmentVariable]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.Environment = v
	return t
}

// SetS__Environment updates property "Environment".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetS__Environment(v ...cfz.Expression[AWS_Pipes_Pipe_BatchEnvironmentVariable]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.Environment = cfz.S(v...)
	return t
}

// SetSV__Environment updates property "Environment".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetSV__Environment(v ...AWS_Pipes_Pipe_BatchEnvironmentVariable) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.Environment = cfz.SV(v...)
	return t
}

// Set__InstanceType updates property "InstanceType".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) Set__InstanceType(v cfz.Expression[string]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.InstanceType = v
	return t
}

// SetV__InstanceType updates property "InstanceType".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetV__InstanceType(v string) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.InstanceType = cfz.V(v)
	return t
}

// Set__ResourceRequirements updates property "ResourceRequirements".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) Set__ResourceRequirements(v cfz.ExpressionSlice[AWS_Pipes_Pipe_BatchResourceRequirement]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.ResourceRequirements = v
	return t
}

// SetS__ResourceRequirements updates property "ResourceRequirements".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetS__ResourceRequirements(v ...cfz.Expression[AWS_Pipes_Pipe_BatchResourceRequirement]) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.ResourceRequirements = cfz.S(v...)
	return t
}

// SetSV__ResourceRequirements updates property "ResourceRequirements".
func (t AWS_Pipes_Pipe_BatchContainerOverrides) SetSV__ResourceRequirements(v ...AWS_Pipes_Pipe_BatchResourceRequirement) AWS_Pipes_Pipe_BatchContainerOverrides {
	t.ResourceRequirements = cfz.SV(v...)
	return t
}
