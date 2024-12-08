// Code generated by "golang-cloudformation". DO NOT EDIT.

package aws_pipes

import (
	"github.com/ibrt/golang-cloudformation/cfz"
)

const (
	// AWS_Pipes_Pipe_EcsContainerOverride__Type is the CloudFormation type for AWS::Pipes::Pipe.EcsContainerOverride.
	AWS_Pipes_Pipe_EcsContainerOverride__Type = "AWS::Pipes::Pipe.EcsContainerOverride"
)

var (
	// AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap reports all the CloudFormation properties for AWS::Pipes::Pipe.EcsContainerOverride.
	AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap = struct {
		Command              string
		Cpu                  string
		Environment          string
		EnvironmentFiles     string
		Memory               string
		MemoryReservation    string
		Name                 string
		ResourceRequirements string
	}{
		Command:              "Command",
		Cpu:                  "Cpu",
		Environment:          "Environment",
		EnvironmentFiles:     "EnvironmentFiles",
		Memory:               "Memory",
		MemoryReservation:    "MemoryReservation",
		Name:                 "Name",
		ResourceRequirements: "ResourceRequirements",
	}

	// AWS_Pipes_Pipe_EcsContainerOverride__PropertiesSlice reports all the CloudFormation properties for AWS::Pipes::Pipe.EcsContainerOverride.
	AWS_Pipes_Pipe_EcsContainerOverride__PropertiesSlice = []string{
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.Command,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.Cpu,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.Environment,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.EnvironmentFiles,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.Memory,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.MemoryReservation,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.Name,
		AWS_Pipes_Pipe_EcsContainerOverride__PropertiesMap.ResourceRequirements,
	}
)

// AWS_Pipes_Pipe_EcsContainerOverride is a binding for AWS::Pipes::Pipe.EcsContainerOverride.
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html
type AWS_Pipes_Pipe_EcsContainerOverride struct {
	// Command is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-command
	Command cfz.ExpressionSlice[string] `json:"Command,omitempty"`

	// Cpu is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-cpu
	Cpu cfz.Expression[int32] `json:"Cpu,omitempty"`

	// Environment is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-environment
	Environment cfz.ExpressionSlice[AWS_Pipes_Pipe_EcsEnvironmentVariable] `json:"Environment,omitempty"`

	// EnvironmentFiles is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-environmentfiles
	EnvironmentFiles cfz.ExpressionSlice[AWS_Pipes_Pipe_EcsEnvironmentFile] `json:"EnvironmentFiles,omitempty"`

	// Memory is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-memory
	Memory cfz.Expression[int32] `json:"Memory,omitempty"`

	// MemoryReservation is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-memoryreservation
	MemoryReservation cfz.Expression[int32] `json:"MemoryReservation,omitempty"`

	// Name is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-name
	Name cfz.Expression[string] `json:"Name,omitempty"`

	// ResourceRequirements is a property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecscontaineroverride.html#cfn-pipes-pipe-ecscontaineroverride-resourcerequirements
	ResourceRequirements cfz.ExpressionSlice[AWS_Pipes_Pipe_EcsResourceRequirement] `json:"ResourceRequirements,omitempty"`
}

// New__AWS_Pipes_Pipe_EcsContainerOverride initializes a new AWS_Pipes_Pipe_EcsContainerOverride.
func New__AWS_Pipes_Pipe_EcsContainerOverride() AWS_Pipes_Pipe_EcsContainerOverride {
	return AWS_Pipes_Pipe_EcsContainerOverride{}
}

// GetType returns the CloudFormation type.
func (AWS_Pipes_Pipe_EcsContainerOverride) GetType() string {
	return AWS_Pipes_Pipe_EcsContainerOverride__Type
}

// Set__Command updates property "Command".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__Command(v cfz.ExpressionSlice[string]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Command = v
	return t
}

// SetS__Command updates property "Command".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetS__Command(v ...cfz.Expression[string]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Command = cfz.S(v...)
	return t
}

// SetSV__Command updates property "Command".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetSV__Command(v ...string) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Command = cfz.SV(v...)
	return t
}

// Set__Cpu updates property "Cpu".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__Cpu(v cfz.Expression[int32]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Cpu = v
	return t
}

// SetV__Cpu updates property "Cpu".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetV__Cpu(v int32) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Cpu = cfz.V(v)
	return t
}

// Set__Environment updates property "Environment".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__Environment(v cfz.ExpressionSlice[AWS_Pipes_Pipe_EcsEnvironmentVariable]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Environment = v
	return t
}

// SetS__Environment updates property "Environment".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetS__Environment(v ...cfz.Expression[AWS_Pipes_Pipe_EcsEnvironmentVariable]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Environment = cfz.S(v...)
	return t
}

// SetSV__Environment updates property "Environment".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetSV__Environment(v ...AWS_Pipes_Pipe_EcsEnvironmentVariable) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Environment = cfz.SV(v...)
	return t
}

// Set__EnvironmentFiles updates property "EnvironmentFiles".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__EnvironmentFiles(v cfz.ExpressionSlice[AWS_Pipes_Pipe_EcsEnvironmentFile]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.EnvironmentFiles = v
	return t
}

// SetS__EnvironmentFiles updates property "EnvironmentFiles".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetS__EnvironmentFiles(v ...cfz.Expression[AWS_Pipes_Pipe_EcsEnvironmentFile]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.EnvironmentFiles = cfz.S(v...)
	return t
}

// SetSV__EnvironmentFiles updates property "EnvironmentFiles".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetSV__EnvironmentFiles(v ...AWS_Pipes_Pipe_EcsEnvironmentFile) AWS_Pipes_Pipe_EcsContainerOverride {
	t.EnvironmentFiles = cfz.SV(v...)
	return t
}

// Set__Memory updates property "Memory".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__Memory(v cfz.Expression[int32]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Memory = v
	return t
}

// SetV__Memory updates property "Memory".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetV__Memory(v int32) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Memory = cfz.V(v)
	return t
}

// Set__MemoryReservation updates property "MemoryReservation".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__MemoryReservation(v cfz.Expression[int32]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.MemoryReservation = v
	return t
}

// SetV__MemoryReservation updates property "MemoryReservation".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetV__MemoryReservation(v int32) AWS_Pipes_Pipe_EcsContainerOverride {
	t.MemoryReservation = cfz.V(v)
	return t
}

// Set__Name updates property "Name".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__Name(v cfz.Expression[string]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Name = v
	return t
}

// SetV__Name updates property "Name".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetV__Name(v string) AWS_Pipes_Pipe_EcsContainerOverride {
	t.Name = cfz.V(v)
	return t
}

// Set__ResourceRequirements updates property "ResourceRequirements".
func (t AWS_Pipes_Pipe_EcsContainerOverride) Set__ResourceRequirements(v cfz.ExpressionSlice[AWS_Pipes_Pipe_EcsResourceRequirement]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.ResourceRequirements = v
	return t
}

// SetS__ResourceRequirements updates property "ResourceRequirements".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetS__ResourceRequirements(v ...cfz.Expression[AWS_Pipes_Pipe_EcsResourceRequirement]) AWS_Pipes_Pipe_EcsContainerOverride {
	t.ResourceRequirements = cfz.S(v...)
	return t
}

// SetSV__ResourceRequirements updates property "ResourceRequirements".
func (t AWS_Pipes_Pipe_EcsContainerOverride) SetSV__ResourceRequirements(v ...AWS_Pipes_Pipe_EcsResourceRequirement) AWS_Pipes_Pipe_EcsContainerOverride {
	t.ResourceRequirements = cfz.SV(v...)
	return t
}
